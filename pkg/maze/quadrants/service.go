package quadrants

import (
	"context"

	"github.com/skyaxl/synack/pkg/maze/quadrants/quadrantsdomain"
	"github.com/skyaxl/synack/pkg/mongoconnector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//StubQuadrantService Stub for sport service
type StubQuadrantService struct {
	Connector mongoconnector.MongoConnector
}

//NewQuadrantService return a new instance of the service.
func NewQuadrantService() (s *StubQuadrantService) {
	//this could be on a config
	client, err := mongo.NewClient(options.Client().ApplyURI(
		"mongodb+srv://admin:admin@appetizer.rtdgx.mongodb.net/Appetizer?retryWrites=true&w=majority",
	))

	if err != nil {
		panic(err)
	}
	s = &StubQuadrantService{mongoconnector.NewDefault(
		"synack",
		"Quadrants",
		client,
	)}
	return
}

//Create create method
func (ma *StubQuadrantService) Create(ctx context.Context, quadrant quadrantsdomain.Quadrant) (res quadrantsdomain.Quadrant, err error) {
	var collection mongoconnector.MongoCollection
	if collection, err = ma.Connector.Connect(ctx); err != nil {
		return quadrant, err
	}
	var mres *mongo.InsertOneResult
	if mres, err = collection.InsertOne(ctx, quadrant); err != nil {
		return quadrant, err
	}

	quadrant.ID = mres.InsertedID.(primitive.ObjectID)
	return quadrant, err
}

//Get get stub by id
func (ma *StubQuadrantService) Get(ctx context.Context, quadrantID string) (quadrant quadrantsdomain.Quadrant, err error) {
	var collection mongoconnector.MongoCollection
	if collection, err = ma.Connector.Connect(ctx); err != nil {
		return quadrant, err
	}

	var mres *mongo.SingleResult
	var id primitive.ObjectID
	if id, err = primitive.ObjectIDFromHex(quadrantID); err != nil {
		return quadrant, err
	}

	filter := bson.M{"_id": id}
	mres = collection.FindOne(ctx, filter)
	if mres.Err() != nil {
		return quadrant, mres.Err()
	}
	quadrant = quadrantsdomain.Quadrant{}
	err = mres.Decode(&quadrant)
	return quadrant, err
}

//GetAll get all
func (ma *StubQuadrantService) GetAll(ctx context.Context, pagination quadrantsdomain.Pagination) (res quadrantsdomain.GetAllResponse, err error) {
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
	res = quadrantsdomain.GetAllResponse{
		Pagination: pagination,
		Quadrants:  []quadrantsdomain.Quadrant{},
	}
	cursor.All(ctx, &res.Quadrants)
	return res, err
}

//Delete Quadrant
func (ma *StubQuadrantService) Delete(ctx context.Context, quadrantID string) (err error) {
	var collection mongoconnector.MongoCollection
	if collection, err = ma.Connector.Connect(ctx); err != nil {
		return err
	}
	//defer ma.client.Disconnect(ctx)
	var id primitive.ObjectID
	if id, err = primitive.ObjectIDFromHex(quadrantID); err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	if _, err = collection.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return
}

//Update Quadrant
func (ma *StubQuadrantService) Update(ctx context.Context, quadrant quadrantsdomain.Quadrant) (m quadrantsdomain.Quadrant, err error) {
	var collection mongoconnector.MongoCollection
	if collection, err = ma.Connector.Connect(ctx); err != nil {
		return m, err
	}
	//defer ma.client.Disconnect(ctx)
	filter := bson.M{"_id": quadrant.ID}
	_, err = collection.ReplaceOne(ctx, filter, quadrant)
	return quadrant, err
}
