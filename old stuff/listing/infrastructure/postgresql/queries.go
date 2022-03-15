package postgresql

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/sss-eda/lemi025/instrument"
	"github.com/sss-eda/lemi025/listing"
)

// GetInstrumentByID TODO
func GetInstrumentByID(
	db *pgx.Conn,
) func(context.Context, instrument.ID) (*listing.Instrument, error) {
	return func(ctx context.Context, id instrument.ID) (*listing.Instrument, error) {
		rows, err := db.Query(ctx, "SELECT * FROM instruments WHERE id=$1", id)
		if err != nil {
			return nil, nil
		}
		defer rows.Close()

		var instrument listing.Instrument
		for rows.Next() {
			err = rows.Scan(&instrument)
			if err != nil {
				return nil, err
			}
		}

		if rows.Err() != nil {
			return nil, rows.Err()
		}

		return &instrument, nil
	}
}
