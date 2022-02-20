package postgres

// // GetInstrumentByID TODO
// func GetInstrumentByID(
// 	db *pgx.Conn,
// ) func(context.Context, instrument.ID) (instrument.Instrument, error) {
// 	return func(ctx context.Context, id instrument.ID) (instrument.Instrument, error) {
// 		rows, err := db.Query(ctx, "SELECT * FROM instruments WHERE id=$1", id)
// 		if err != nil {
// 			return instrument.Instrument{}, nil
// 		}
// 		rows.
// 		for _, row := rows
// 	}
// }
