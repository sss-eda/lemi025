package mongodb

import (
	"github.com/sss-eda/lemi025/listing"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetInstrumentByID TODO
func GetInstrumentByID(
	db mongo.Database,
) func(listing.InstrumentID) (*listing.Instrument, error) {
	return func(id listing.InstrumentID) (*listing.Instrument, error) {
		collection := db.Collection("instruments")
		collection.FindOne(ctx, nil, options.FindOneOptions{})
	}
}
