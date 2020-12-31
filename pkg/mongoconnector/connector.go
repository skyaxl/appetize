package mongoconnector

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCollection interface {
	ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error)
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
}

type MongoConnector interface {
	Connect(ctx context.Context) (MongoCollection, error)
}

type Default struct {
	Database   string
	Collection string
	Client     *mongo.Client
}

//Connect connect with client and return a collection
func (ma *Default) Connect(ctx context.Context) (collection MongoCollection, err error) {
	if err = ma.Client.Ping(ctx, nil); err != nil {
		if err = ma.Client.Connect(ctx); err != nil {
			return nil, err
		}
	}
	collection = ma.Client.Database(ma.Database).Collection(ma.Collection)
	return
}

func NewDefault(database string, collection string, client *mongo.Client) *Default {
	return &Default{database, collection, client}
}
