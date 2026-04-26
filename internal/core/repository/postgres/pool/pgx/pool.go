package core_postgres_pgx

import (
	"context"
	"fmt"
	"time"

	core_postgres_pool "github.com/Mirwinli/coffe_plus/internal/core/repository/postgres/pool"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool struct {
	*pgxpool.Pool
	opTimeout time.Duration
}

func NewPool(ctx context.Context, config Config) (*Pool, error) {
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	poolConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("parse pgxconfig: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("create new pgxpool: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pgxpool ping: %w", err)
	}

	return &Pool{
		pool,
		config.Timeout,
	}, nil
}

func (p *Pool) Query(
	ctx context.Context,
	sql string,
	args ...any,
) (core_postgres_pool.Rows, error) {
	rows, err := p.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	return pgxRows{rows}, nil
}

func (p *Pool) QueryRow(ctx context.Context, sql string, args ...any) core_postgres_pool.Row {
	row := p.Pool.QueryRow(ctx, sql, args...)

	return pgxRow{row}
}

func (p *Pool) Exec(ctx context.Context, sql string, args ...any) (core_postgres_pool.CommandTag, error) {
	cmdTag, err := p.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return nil, mapErrors(err)
	}

	return pgxCommandTag{cmdTag}, nil
}

func (p *Pool) OpTimeout() time.Duration {
	return p.opTimeout
}

func (p *Pool) Begin(ctx context.Context) (core_postgres_pool.Tx, error) {
	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return &pgxTx{tx}, nil
}

func (p *Pool) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return p.Pool.SendBatch(ctx, b)
}
