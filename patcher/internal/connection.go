package internal

import (
	"context"
	"database/sql"

	"github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
)

type IDb interface {
	Select(ctx context.Context, dest interface{}, sqlizer sqrl.Sqlizer) error
	Get(ctx context.Context, dest interface{}, sqlizer sqrl.Sqlizer) error
	Exec(ctx context.Context, sqlizer sqrl.Sqlizer) (sql.Result, error)
}

type DB struct {
	DB *sqlx.DB
}

func (db *DB) Get(ctx context.Context, dest interface{}, sqlizer sqrl.Sqlizer) error {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}
	sqlxDB := db.DB

	err = sqlxDB.GetContext(ctx, dest, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) Select(ctx context.Context, dest interface{}, sqlizer sqrl.Sqlizer) error {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}
	sqlxDB := db.DB

	err = sqlxDB.SelectContext(ctx, dest, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) Exec(ctx context.Context, sqlizer sqrl.Sqlizer) (sql.Result, error) {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return nil, err
	}
	sqlxDB := db.DB

	res, err := sqlxDB.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
