package main

import (
	"flag"
	"fmt"
	"os"

	"time"

	"github.com/piotrkowalczuk/mnemosyne/mnemosyned"
)

const version = "0.2.1"

type configuration struct {
	host      string
	port      int
	namespace string
	subsystem string
	logger    struct {
		adapter string
		format  string
		level   int
	}
	session struct {
		ttl time.Duration
		ttc time.Duration
	}
	monitoring struct {
		enabled bool
	}
	storage struct {
		engine   string
		postgres struct {
			address string
			table   string
		}
	}
	tls struct {
		enabled  bool
		certFile string
		keyFile  string
	}
}

func (c *configuration) init() {
	if c == nil {
		*c = configuration{}
	}

	flag.StringVar(&c.host, "host", "127.0.0.1", "host")
	flag.IntVar(&c.port, "port", 8080, "port")
	flag.StringVar(&c.namespace, "namespace", "", "namespace")
	flag.StringVar(&c.subsystem, "subsystem", "mnemosyne", "subsystem")
	flag.DurationVar(&c.session.ttl, "ttl", mnemosyned.DefaultTTL, "session time to live, after which session is deleted")
	flag.DurationVar(&c.session.ttc, "ttc", mnemosyned.DefaultTTC, "session time to cleanup, how offten cleanup will be performed")
	flag.StringVar(&c.logger.adapter, "log.adapter", loggerAdapterStdOut, "logger adapter")
	flag.StringVar(&c.logger.format, "log.format", loggerFormatJSON, "logger format")
	flag.IntVar(&c.logger.level, "log.level", 6, "logger level")
	flag.BoolVar(&c.monitoring.enabled, "monitoring", false, "toggle application monitoring")
	flag.StringVar(&c.storage.engine, "storage.engine", mnemosyned.StorageEnginePostgres, "storage engine") // TODO: change to in memory when implemented
	flag.StringVar(&c.storage.postgres.address, "storage.postgres.address", "postgres://postgres:postgres@postgres/postgres?sslmode=disable", "storage postgres connection string")
	flag.StringVar(&c.storage.postgres.table, "storage.postgres.table", "session", "storage postgres table name")
	flag.BoolVar(&c.tls.enabled, "tls", false, "tls enable flag")
	flag.StringVar(&c.tls.certFile, "tls.certfile", "", "path to tls cert file")
	flag.StringVar(&c.tls.keyFile, "tls.keyfile", "", "path to tls key file")
}

func (c *configuration) parse() {
	if !flag.Parsed() {
		ver := flag.Bool("version", false, "print version and exit")
		flag.Parse()
		if *ver {
			fmt.Printf("%s", version)
			os.Exit(0)
		}
	}
}
