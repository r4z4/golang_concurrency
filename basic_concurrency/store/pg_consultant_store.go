package store

import (
	"context"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

const driverName = "pgx"

type PgConsultantStore struct {
	dbUrl string
	dbx   *sqlx.DB
}

func NewPostgresMoviesStore(dbUrl string) *PgConsultantStore {
	return &PgConsultantStore{
		dbUrl: dbUrl,
	}
}

func (s *PgConsultantStore) connect(ctx context.Context) error {
	dbx, err := sqlx.ConnectContext(ctx, driverName, s.dbUrl)
	if err != nil {
		return err
	}

	s.dbx = dbx
	return nil
}

func (s *PgConsultantStore) close() error {
	return s.dbx.Close()
}
