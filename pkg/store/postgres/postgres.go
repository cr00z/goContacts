package postgres

import (
	"context"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

const pingTimeout = 5 * time.Second

type Store struct {
	Pool *pgxpool.Pool
}

type Settings struct {
	Hostname  string
	Port      int
	Database  string
	Username  string
	Password  string
	SSLEnable bool
}

func init() {
	if err := rewriteDefaultEnv(); err != nil {
		panic(err)
	}
}

func rewriteDefaultEnv() error {
	defaults := map[string]string{
		"PGHOST":     "localhost",
		"PGPORT":     "5432",
		"PGDATABASE": "postgres",
		"PGUSER":     "postgres",
		"PGPASSWORD": "password",
		"PGSSLMODE":  "disable",
	}
	for envKey, envVal := range defaults {
		if len(os.Getenv(envKey)) == 0 {
			if err := os.Setenv(envKey, envVal); err != nil {
				return errors.WithStack(err)
			}
		}
	}
	return nil
}

//func (s Settings) toDSN() string {
//	sslMode := "disable"
//	if s.SSLEnable {
//		sslMode = "enable"
//	}
//	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
//		s.Username, s.Password, s.Hostname, s.Port, s.Database, sslMode)
//}

func (s Settings) toConnStr() string {
	var args []string

	if len(s.Hostname) > 0 {
		args = append(args, "host="+s.Hostname)
	}

	if s.Port > 0 {
		args = append(args, "port="+strconv.Itoa(s.Port))
	}

	if len(s.Database) > 0 {
		args = append(args, "dbname="+s.Database)
	}

	if len(s.Username) > 0 {
		args = append(args, "user="+s.Username)
	}

	if len(s.Password) > 0 {
		args = append(args, "password="+s.Password)
	}

	sslMode := "sslmode=disable"
	if s.SSLEnable {
		sslMode = "sslmode=enable"
	}
	args = append(args, sslMode)

	return strings.Join(args, " ")
}

func New(settings Settings) (*Store, error) {
	config, err := pgxpool.ParseConfig(settings.toConnStr())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()
	if err = pool.Ping(ctx); err != nil {
		return nil, errors.WithStack(err)
	}

	return &Store{pool}, nil
}
