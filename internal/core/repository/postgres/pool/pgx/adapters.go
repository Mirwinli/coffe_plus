package core_postgres_pgx

import (
	"context"
	"errors"
	"fmt"

	core_postgres_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type pgxRows struct {
	pgx.Rows
}

type pgxRow struct {
	pgx.Row
}

func (r pgxRow) Scan(dest ...any) error {
	if err := r.Row.Scan(dest...); err != nil {
		return mapErrors(err)
	}

	return nil
}

type pgxCommandTag struct {
	pgconn.CommandTag
}

type pgxBatchResults struct {
	pgx.BatchResults
}

type pgxTx struct {
	tx pgx.Tx
}

func (p *pgxTx) Commit(ctx context.Context) error {
	return p.tx.Commit(ctx)
}

func (p *pgxTx) Rollback(ctx context.Context) error {
	return p.tx.Rollback(ctx)
}

func (p *pgxTx) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return p.tx.SendBatch(ctx, b)
}

func (p *pgxTx) Query(ctx context.Context, sql string, args ...any) (core_postgres_pool.Rows, error) {
	rows, err := p.tx.Query(ctx, sql, args...)
	if err != nil {
		return nil, mapErrors(err)
	}

	return &pgxRows{rows}, nil
}

func (p *pgxTx) QueryRow(ctx context.Context, sql string, args ...any) core_postgres_pool.Row {
	row := p.tx.QueryRow(ctx, sql, args...)

	return &pgxRow{row}
}

func (p *pgxTx) Exec(ctx context.Context, sql string, args ...any) (core_postgres_pool.CommandTag, error) {
	cmd, err := p.tx.Exec(ctx, sql, args...)
	if err != nil {
		return nil, mapErrors(err)
	}

	return &pgxCommandTag{cmd}, nil
}

func mapErrors(err error) error {
	const (
		pgxViolatesUniqueErrorCode = "23505"
		pgxForeignKeyViolation     = "23503"
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return core_postgres_pool.ErrNoRows
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == pgxViolatesUniqueErrorCode {
			return fmt.Errorf(
				"%v: %w",
				err,
				core_postgres_pool.ErrViolatesUnique,
			)
		}
		if pgErr.Code == pgxForeignKeyViolation {
			return fmt.Errorf(
				"%v: %w",
				err,
				core_postgres_pool.ErrForeignKeyViolation,
			)
		}
	}

	return fmt.Errorf(
		"%v: %w",
		err,
		core_postgres_pool.ErrUnknown,
	)
}
