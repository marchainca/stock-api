package stock

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
)

type Repository interface {
	SaveItems(ctx context.Context, items []Item) error
}

type repo struct{ db *sql.DB }

func NewRepo(db *sql.DB) Repository { return &repo{db} }

func (r *repo) SaveItems(ctx context.Context, items []Item) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO stock_items
		  (ticker, target_from, target_to, company, action, brokerage,
		   rating_from, rating_to, ts)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		ON CONFLICT (id) DO NOTHING
	`)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for _, it := range items {
		if _, err := stmt.ExecContext(ctx,
			it.Ticker,
			trimDollar(it.TargetFrom),
			trimDollar(it.TargetTo),
			it.Company,
			it.Action,
			it.Brokerage,
			it.RatingFrom,
			it.RatingTo,
			it.Time,
		); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

// helper simple para quitar el símbolo '$'
func trimDollar(s string) sql.NullFloat64 {
	f, _ := strconv.ParseFloat(strings.TrimPrefix(s, "$"), 64)
	return sql.NullFloat64{Float64: f, Valid: s != ""}
}
