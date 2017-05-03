package postgres_test

import (
	"flag"
	"os"
	"testing"

	"time"

	_ "github.com/lib/pq"
	"github.com/piotrkowalczuk/mnemosyne/internal/service/postgres"
	"github.com/piotrkowalczuk/sklog"
)

var (
	testPostgresAddress string
)

func TestMain(m *testing.M) {
	flag.StringVar(&testPostgresAddress, "postgres.address", getStringEnvOr("MNEMOSYNED_POSTGRES_ADDRESS", "postgres://localhost/test?sslmode=disable"), "")
	flag.Parse()

	os.Exit(m.Run())
}

func getStringEnvOr(env, or string) string {
	if v := os.Getenv(env); v != "" {
		return v
	}
	return or
}

func TestInit_retry(t *testing.T) {
	_, err := postgres.Init(testPostgresAddress, postgres.Opts{
		Logger:  sklog.NewTestLogger(t),
		Timeout: 10 * time.Second,
		Retry:   1 * time.Microsecond,
	})
	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}
}

func TestInit_timeout(t *testing.T) {
	_, err := postgres.Init(testPostgresAddress, postgres.Opts{
		Logger:  sklog.NewTestLogger(t),
		Timeout: 1 * time.Microsecond,
		Retry:   1 * time.Microsecond,
	})
	if err != nil && err != postgres.Timeout {
		t.Fatalf("unexpected error: %s", err.Error())
	}
	if err == nil {
		t.Fatal("error expected, got nil")
	}
}
func TestInit_empty(t *testing.T) {
	_, err := postgres.Init(testPostgresAddress, postgres.Opts{
		Logger: sklog.NewTestLogger(t),
	})
	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}
}
