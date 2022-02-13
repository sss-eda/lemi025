package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/sss-eda/lemi025/internal/domain/entities/instrument"
)

// OnConfigRead - But this is specific to listing. And that's the way it should
// be. Decoupled from the command side.
func OnConfigRead(
	db *pgx.Conn,
) func(context.Context, instrument.ConfigReadEvent) error {
	return func(ctx context.Context, event instrument.ConfigReadEvent) error {
		// TODO: Do something with the CommandTag
		_, err := db.Exec(
			ctx,
			"INSERT INTO config(instrument_id,station_number) VALUES $1, $2",
			event.InstrumentID,
			event.NewStationNumber,
		)
		if err != nil {
			return err
		}
		return nil
	}
}

// OnTimeRead TODO
func OnTimeRead(
	db *pgx.Conn,
) func(context.Context, instrument.TimeReadEvent) error {
	return func(ctx context.Context, event instrument.TimeReadEvent) error {
		ctag, err := db.Exec(
			ctx,
			"INSERT INTO configs(instrument_id,station_number) VALUES $1, $2",
			event.AggregateID,
			event.Payload.Config.StationNumber,
		)
		if err != nil {
			return err
		}
		return nil
	}
}
