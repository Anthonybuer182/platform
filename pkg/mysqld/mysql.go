package mysqld

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"platform/cmd/order/config"
	"time"

	"golang.org/x/exp/slog"
)

const (
	_defaultConnAttempts = 3
	_defaultConnTimeout  = time.Second
)

type mysqldb struct {
	connAttempts int
	connTimeout  time.Duration
	maxIdleConns int
	maxOpenConns int
	db           *sql.DB
}

var _ DBEngine = (*mysqldb)(nil)

func NewMysqlDb(mysqlCfg config.Mysql) (DBEngine, error) {
	slog.Info("CONN", "connect string", mysqlCfg.URL)

	pg := &mysqldb{
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
		maxOpenConns: mysqlCfg.MaxOpenConns,
		maxIdleConns: mysqlCfg.MaxIdleConns,
	}

	var err error
	for pg.connAttempts > 0 {
		slog.Info(string(mysqlCfg.URL))
		pg.db, err = sql.Open("mysql", string(mysqlCfg.URL))
		slog.Info("pg.db===", pg.db)
		pg.db.SetConnMaxLifetime(time.Minute)
		pg.db.SetMaxIdleConns(int(pg.maxIdleConns))
		pg.db.SetMaxOpenConns(int(pg.maxOpenConns))
		if err != nil && pg.db.Ping() != nil {
			break
		}

		log.Printf("mysql is trying to connect, attempts left: %d", pg.connAttempts)

		time.Sleep(pg.connTimeout)

		pg.connAttempts--
	}

	slog.Info("📰 connected to mysql 🎉")

	return pg, nil
}

func (p *mysqldb) Configure(opts ...Options) DBEngine {
	for _, opt := range opts {
		opt(p)
	}

	return p
}

func (p *mysqldb) GetDB() *sql.DB {
	return p.db
}

func (p *mysqldb) Close() {
	if p.db != nil {
		p.db.Close()
	}
}

func (p *mysqldb) GetDbName() string {
	return "mysql"
}
