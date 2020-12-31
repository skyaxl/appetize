package spots

import (
	"context"

	"github.com/skyaxl/synack/pkg/maze/spots/spotsdomain"
	"github.com/skyaxl/synack/pkg/mongoconnector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//StubSpotService Stub for sport service
type StubSpotService struct {
	Connector mongoconnector.MongoConnector
}

//NewSpotService return a new instance of the service.
func NewSpotService() (s *StubSpotService) {
	//this could be on a config
	client, err := mongo.NewClient(options.Client().ApplyURI(
		"mongodb+srv://admin:admin@appetizer.rtdgx.mongodb.net/Appetizer?retryWrites=true&w=majority",
	))

	if err != nil {
		panic(err)
	}
	s = &StubSpotService{mongoconnector.NewDefault(
		"synack",
		"Spots",
		client,
	)}
	return
}

//Create create method
func (ma *StubSpotService) Create(ctx context.Context, spot spotsdomain.Spot) (res spotsdomain.Spot, err error) {
	var collection mongoconnector.MongoCollection
	if collection, err = ma.Connector.Connect(ctx); err != nil {
		return spot, err
	}
	var mres *mongo.InsertOneResult
	if mres, err = collection.InsertOne(ctx, spot); err != nil {
		return spot, err
	}

	spot.ID = mres.InsertedID.(primitive.ObjectID)
	return spot, err
}

//Get get stub by id
func (ma *StubSpotService) Get(ctx context.Context, spotID string) (spot spotsdomain.Spot, err error) {
	var collection mongoconnector.MongoCollection
	if collection, err = ma.Connector.Connect(ctx); err != nil {
		return spot, err
	}

	var mres *mongo.SingleResult
	var id primitive.ObjectID
	if id, err = primitive.ObjectIDFromHex(spotID); err != nil {
		return spot, err
	}

	filter := bson.M{"_id": id}
	mres = collection.FindOne(ctx, filter)
	if mres.Err() != nil {
		return spot, mres.Err()
	}
	spot = spotsdomain.Spot{}
	err = mres.Decode(&spot)
	return spot, err
}

//Delete Spot
func (ma *StubSpotService) Delete(ctx context.Context, spotID string) (err error) {
	var collection mongoconnector.MongoCollection
	if collection, err = ma.Connector.Connect(ctx); err != nil {
		return err
	}
	//defer ma.client.Disconnect(ctx)
	var id primitive.ObjectID
	if id, err = primitive.ObjectIDFromHex(spotID); err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	if _, err = collection.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return
}

//GetAll get all
func (ma *StubSpotService) GetAll(ctx context.Context, pagination spotsdomain.Pagination) (res spotsdomain.GetAllResponse, err error) {
	var collection mongoconnector.MongoCollection
	if collection, err = ma.Connector.Connect(ctx); err != nil {
		return res, err
	}
	filter := bson.M{}
	var cursor *mongo.Cursor
	cursor, err = collection.Find(ctx, filter, options.Find().SetLimit(pagination.Limit).SetSkip(pagination.Limit*(pagination.Page-1)))
	if err != nil {
		return res, err
	}
	pagination.Total = int64(cursor.RemainingBatchLength())
	res = spotsdomain.GetAllResponse{
		Pagination: pagination,
		Spots:      []spotsdomain.Spot{},
	}
	cursor.All(ctx, &res.Spots)
	return res, err
}

//Update Spot
func (ma *StubSpotService) Update(ctx context.Context, spot spotsdomain.Spot) (m spotsdomain.Spot, err error) {
	var collection mongoconnector.MongoCollection
	if collection, err = ma.Connector.Connect(ctx); err != nil {
		return m, err
	}
	//defer ma.client.Disconnect(ctx)
	filter := bson.M{"_id": spot.ID}
	_, err = collection.ReplaceOne(ctx, filter, spot)
	return spot, err
}
