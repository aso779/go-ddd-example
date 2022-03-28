package connection

import (
	"database/sql"
	"fmt"
	"github.com/aso779/go-ddd-example/application/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"go.uber.org/zap"
	"time"
)

const (
	retryTimeout      = 3000 * time.Millisecond
	maxConnectAttempt = 10
)

type ConnSet struct {
	conf      *config.Config
	log       *zap.Logger
	readPool  *bun.DB
	writePool *bun.DB
}

var set *ConnSet

func NewPGConnSet(
	conf *config.Config,
	log *zap.Logger,
) *ConnSet {
	if set == nil {
		set = &ConnSet{
			conf: conf,
			log:  log,
		}
	}

	return set
}

type DbLogger struct {
}

func (r *ConnSet) ReadPool() *bun.DB {
	readConnAttempts := 0

start:
	if r.readPool == nil {
		r.readPool = r.connect(r.conf.Postgres.Read)
	}
	if readConnAttempts > maxConnectAttempt {
		panic("can't connect to read db")
	}

	if err := Ping(r.readPool); err != nil {
		r.log.Info("try to ping read conn", zap.Int("attempt", readConnAttempts), zap.Error(err))
		r.readPool = nil
		readConnAttempts += 1
		time.Sleep(retryTimeout)
		goto start
	}

	return r.readPool
}

// WritePool get write connection pool
func (r *ConnSet) WritePool() *bun.DB {
	writeConnAttempts := 0
start:
	if r.writePool == nil {
		r.writePool = r.connect(r.conf.Postgres.Write)
	}
	if writeConnAttempts > maxConnectAttempt {
		// do now allow goroitune count growing
		panic("can't connect to write db")
	}

	if err := Ping(r.writePool); err != nil {
		r.log.Info("try to ping read conn", zap.Int("attempt", writeConnAttempts), zap.Error(err))
		r.writePool = nil
		writeConnAttempts += 1
		time.Sleep(retryTimeout)
		goto start
	}

	return r.writePool
}

func Ping(pool *bun.DB) error {
	if pool == nil {
		//TODO err
	}
	if _, err := pool.Exec("SELECT 1;"); err != nil {
		return err
	}

	return nil
}

func (r *ConnSet) connect(config config.Postgres) *bun.DB {
	conn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(fmt.Sprintf("%s:%d", config.Host, config.Port)),
		//pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		pgdriver.WithInsecure(true),
		pgdriver.WithUser(config.User),
		pgdriver.WithPassword(config.Password),
		pgdriver.WithDatabase(config.Database),
		pgdriver.WithApplicationName("books"),
		pgdriver.WithTimeout(5*time.Second),
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(5*time.Second),
		pgdriver.WithWriteTimeout(5*time.Second),
	)

	db := bun.NewDB(sql.OpenDB(conn), pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	return db
}
