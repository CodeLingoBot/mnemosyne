package main

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/piotrkowalczuk/mnemosyne"
)

const (
	postgresTableNamePlaceholder = "%%TABLE_NAME%%"
)

type postgresStorage struct {
	db        *sql.DB
	tableName string
	generator mnemosyne.RandomBytesGenerator
	monitor   *monitoring
}

func newPostgresStorage(tn string, db *sql.DB, m *monitoring) Storage {
	return &postgresStorage{
		db:        db,
		tableName: tn,
		generator: &mnemosyne.SystemRandomBytesGenerator{},
		monitor:   m,
	}
}

func initPostgresStorage(tn string, db *sql.DB, m *monitoring) func() (Storage, error) {
	return func() (Storage, error) {
		return newPostgresStorage(tn, db, m), nil
	}
}

// Create ...
func (ps *postgresStorage) Create(data map[string]string) (*mnemosyne.Session, error) {
	id, err := mnemosyne.NewIDRandom(ps.generator, "1")
	if err != nil {
		return nil, err
	}

	entity := &SessionEntity{
		ID:   id,
		Data: Data(data),
	}

	if err := ps.save(entity); err != nil {
		return nil, err
	}

	return newSessionFromSessionEntity(entity), nil
}

func (ps *postgresStorage) save(entity *SessionEntity) error {
	query := `
		INSERT INTO ` + ps.tableName + ` (id, data, expire_at)
		VALUES ($1, $2, NOW() + '30 minutes'::interval)
		RETURNING expire_at

	`
	field := metrics.Field{Key: "query", Value: query}

	encodedData, err := entity.Data.EncodeToJSON()
	if err != nil {
		ps.monitor.postgres.errors.With(field).Add(1)
		return err
	}

	err = ps.db.QueryRow(
		query,
		entity.ID,
		encodedData,
	).Scan(
		&entity.ExpireAt,
	)
	ps.monitor.postgres.queries.With(field).Add(1)

	return err
}

// Get ...
func (ps *postgresStorage) Get(id *mnemosyne.ID) (*mnemosyne.Session, error) {
	var data Data
	var expireAt time.Time
	query := `
		SELECT data, expire_at
		FROM ` + ps.tableName + `
		WHERE id = $1
		LIMIT 1
	`
	field := metrics.Field{Key: "query", Value: query}

	err := ps.db.QueryRow(query, id).Scan(
		&data,
		&expireAt,
	)
	if err != nil {
		ps.monitor.postgres.errors.With(field).Add(1)
		if err == sql.ErrNoRows {
			return nil, errSessionNotFound
		}
		return nil, err
	}

	return &mnemosyne.Session{
		Id:       id,
		Data:     data,
		ExpireAt: mnemosyne.TimeToTimestamp(expireAt),
	}, nil
}

// List satisfy Storage interface.
func (ps *postgresStorage) List(offset, limit int64, expiredAtFrom, expiredAtTo *time.Time) ([]*mnemosyne.Session, error) {
	if limit == 0 {
		return nil, errors.New("mnemosyned: cannot retrieve list of sessions, limit needs to be higher than 0")
	}

	args := []interface{}{offset, limit}
	query := "SELECT id, data, expire_at FROM " + ps.tableName

	switch {
	case expiredAtFrom != nil && expiredAtTo == nil:
		query += "expire_at > $3"
		args = append(args, expiredAtFrom)
	case expiredAtFrom == nil && expiredAtTo != nil:
		query += "expire_at < $3"
		args = append(args, expiredAtTo)
	case expiredAtFrom != nil && expiredAtTo != nil:
		query += "expire_at > $4 AND expire_at < $5"
		args = append(args, expiredAtFrom, expiredAtTo)
	}

	query += " OFFSET $1 LIMIT $2"

	field := metrics.Field{Key: "query", Value: query}

	rows, err := ps.db.Query(query, args...)
	if err != nil {
		ps.monitor.postgres.errors.With(field).Add(1)
		return nil, err
	}
	ps.monitor.postgres.queries.With(field).Add(1)

	sessions := make([]*mnemosyne.Session, 0, limit)
	for rows.Next() {
		var id mnemosyne.ID
		var data Data
		var expireAt time.Time

		err = rows.Scan(
			&id,
			&data,
			&expireAt,
		)
		if err != nil {
			ps.monitor.postgres.errors.With(field).Add(1)
			return nil, err
		}

		sessions = append(sessions, &mnemosyne.Session{
			Id:       &id,
			Data:     data,
			ExpireAt: mnemosyne.TimeToTimestamp(expireAt),
		})
	}
	if rows.Err() != nil {
		ps.monitor.postgres.errors.With(field).Add(1)
		return nil, rows.Err()
	}

	return sessions, nil
}

// Exists ...
func (ps *postgresStorage) Exists(id *mnemosyne.ID) (exists bool, err error) {
	query := `SELECT EXISTS(SELECT 1 FROM ` + ps.tableName + ` WHERE id = $1)`
	field := metrics.Field{Key: "query", Value: query}

	err = ps.db.QueryRow(query, id).Scan(
		&exists,
	)
	if err != nil {
		ps.monitor.postgres.errors.With(field).Add(1)
	}
	ps.monitor.postgres.queries.With(field).Add(1)

	return
}

// Abandon ...
func (ps *postgresStorage) Abandon(id *mnemosyne.ID) (bool, error) {
	query := `DELETE FROM ` + ps.tableName + ` WHERE id = $1`
	field := metrics.Field{Key: "query", Value: query}

	result, err := ps.db.Exec(query, id)
	if err != nil {
		ps.monitor.postgres.errors.With(field).Add(1)
		return false, err
	}

	ps.monitor.postgres.queries.With(field).Add(1)

	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if affected == 0 {
		return false, errSessionNotFound
	}

	return true, nil
}

// SetData ...
func (ps *postgresStorage) SetData(id *mnemosyne.ID, key, value string) (*mnemosyne.Session, error) {
	var dataEncoded []byte
	var err error

	entity := &SessionEntity{
		ID: id,
	}
	selectQuery := `
		SELECT data, expire_at
		FROM ` + ps.tableName + `
		WHERE id = $1
		FOR UPDATE
	`
	updateQuery := `
		UPDATE ` + ps.tableName + `
		SET
			data = $2
		WHERE id = $1
	`

	tx, err := ps.db.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.QueryRow(selectQuery, id).Scan(
		&dataEncoded,
		&entity.ExpireAt,
	)
	if err != nil {
		ps.monitor.postgres.errors.With(metrics.Field{Key: "query", Value: selectQuery}).Add(1)
		tx.Rollback()
		if err == sql.ErrNoRows {
			return nil, errSessionNotFound
		}
		return nil, err
	}
	ps.monitor.postgres.queries.With(metrics.Field{Key: "query", Value: selectQuery}).Add(1)

	entity.Data, err = DecodeSessionDataFromJSON(dataEncoded)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	entity.Data.Set(key, value)

	dataEncoded, err = entity.Data.EncodeToJSON()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec(updateQuery, id, dataEncoded)
	if err != nil {
		ps.monitor.postgres.errors.With(metrics.Field{Key: "query", Value: updateQuery}).Add(1)
		tx.Rollback()
		return nil, err
	}
	ps.monitor.postgres.queries.With(metrics.Field{Key: "query", Value: updateQuery}).Add(1)

	tx.Commit()

	return newSessionFromSessionEntity(entity), nil
}

// Delete
func (ps *postgresStorage) Delete(id *mnemosyne.ID, expiredAtFrom, expiredAtTo *time.Time) (int64, error) {
	if id == nil && expiredAtFrom == nil && expiredAtTo == nil {
		return 0, errors.New("mnemosyned: session cannot be deleted, no where parameter provided")
	}

	where, args := ps.where(id, expiredAtFrom, expiredAtTo)
	query := "DELETE FROM " + ps.tableName + " WHERE " + where
	field := metrics.Field{Key: "query", Value: query}

	result, err := ps.db.Exec(query, args...)
	if err != nil {
		ps.monitor.postgres.errors.With(field).Add(1)
		return 0, err
	}
	ps.monitor.postgres.queries.With(field).Add(1)

	return result.RowsAffected()
}

// Setup prepares storage for usage.
func (ps *postgresStorage) Setup() error {
	sql := strings.Replace(`
		CREATE TABLE IF NOT EXISTS %%TABLE_NAME%% (
			id character varying(255) PRIMARY KEY,
			data json NOT NULL,
			expire_at timestamp with time zone NOT NULL
		)
    `, postgresTableNamePlaceholder, ps.tableName, -1)

	_, err := ps.db.Exec(sql)

	return err
}

// TearDown ...
func (ps *postgresStorage) TearDown() error {
	sql := strings.Replace(`DROP TABLE %%TABLE_NAME%%`, postgresTableNamePlaceholder, ps.tableName, -1)

	_, err := ps.db.Exec(sql)

	return err
}

func (ps *postgresStorage) where(id *mnemosyne.ID, expiredAtFrom, expiredAtTo *time.Time) (string, []interface{}) {
	switch {
	case id != nil && expiredAtFrom == nil && expiredAtTo == nil:
		return "id = $1", []interface{}{id}
	case id == nil && expiredAtFrom != nil && expiredAtTo == nil:
		return "expire_at > $1", []interface{}{expiredAtFrom}
	case id == nil && expiredAtFrom == nil && expiredAtTo != nil:
		return "expire_at < $1", []interface{}{expiredAtTo}
	case id != nil && expiredAtFrom != nil && expiredAtTo == nil:
		return "id = $1 AND expire_at > $2", []interface{}{id, expiredAtFrom}
	case id != nil && expiredAtFrom == nil && expiredAtTo != nil:
		return "id = $1 AND expire_at < $2", []interface{}{id, expiredAtTo}
	case id == nil && expiredAtFrom != nil && expiredAtTo != nil:
		return "expire_at > $1 AND expire_at < $2", []interface{}{expiredAtFrom, expiredAtTo}
	case id != nil && expiredAtFrom != nil && expiredAtTo != nil:
		return "id = $1 AND expire_at > $2 AND expire_at < $3", []interface{}{id, expiredAtFrom, expiredAtTo}
	default:
		return "", nil
	}
}
