package paths

import (
	"context"

	"github.com/skyaxl/synack/pkg/maze/paths/pathsdomain"
	"github.com/skyaxl/synack/pkg/mongoconnector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//StubPathService Stub for sport service
type StubPathService struct {
	Connector mongoconnector.MongoConnector
}

//NewPathService return a new instance of the service.
func NewPathService() (s *StubPathService) {
	//this could be on a config
	client, err := mongo.NewClient(options.Client().ApplyURI(
		"mongodb+srv://admin:admin@appetizer.rtdgx.mongodb.net/Appetizer?retryWrites=true&w=majority",
	))

	if err != nil {
		panic(err)
	}
	s = &StubPathService{mongoconnector.NewDefault(
		"synack",
		"Paths",
		client,
	)}
	return
}

//Create create method
func (ma *StubPathService) Create(ctx context.Context, path pathsdomain.Path) (res pathsdomain.Path, err error) {
	var collection mongoconnector.MongoCollection
	if collection, err = ma.Connector.Connect(ctx); err != nil {
		return path, err
	}
	var mres *mongo.InsertOneResult
	if mres, err = collection.InsertOne(ctx, path); err != nil {
		return path, err
	}

	path.ID = mres.InsertedID.(primitive.ObjectID)
	return path, err
}

//Get get stub by id
func (ma *StubPathService) Get(ctx context.Context, pathID string) (path pathsdomain.Path, err error) {
	var collection mongoconnector.MongoCollection
	if collection, err = ma.Connector.Connect(ctx); err != nil {
		return path, err
	}

	var mres *mongo.SingleResult
	var id primitive.ObjectID
	if id, err = primitive.ObjectIDFromHex(pathID); err != nil {
		return path, err
	}

	filter := bson.M{"_id": id}
	mres = collection.FindOne(ctx, filter)
	if mres.Err() != nil {
		return path, mres.Err()
	}
	path = pathsdomain.Path{}
	err = mres.Decode(&path)
	return path, err
}

//GetAll get all
func (ma *StubPathService) GetAll(ctx context.Context, pagination pathsdomain.Pagination) (res pathsdomain.GetAllResponse, err error) {
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
	res = pathsdomain.GetAllResponse{
		Pagination: pagination,
		Paths:      []pathsdomain.Path{},
	}
	cursor.All(ctx, &res.Paths)
	return res, err
}

//Delete Path
func (ma *StubPathService) Delete(ctx context.Context, pathID string) (err error) {
	var collection mongoconnector.MongoCollection
	if collection, err = ma.Connector.Connect(ctx); err != nil {
		return err
	}
	//defer ma.client.Disconnect(ctx)
	var id primitive.ObjectID
	if id, err = primitive.ObjectIDFromHex(pathID); err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	if _, err = collection.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return
}

//Update Path
func (ma *StubPathService) Update(ctx context.Context, path pathsdomain.Path) (m pathsdomain.Path, err error) {
	var collection mongoconnector.MongoCollection
	if collection, err = ma.Connector.Connect(ctx); err != nil {
		return m, err
	}
	//defer ma.client.Disconnect(ctx)
	filter := bson.M{"_id": path.ID}
	_, err = collection.ReplaceOne(ctx, filter, path)
	return path, err
}
