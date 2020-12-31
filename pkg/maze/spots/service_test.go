package spots

import (
	"context"
	"errors"
	"testing"

	"github.com/skyaxl/synack/mocks"
	"github.com/skyaxl/synack/pkg/maze/spots/spotsdomain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestStubSpotService_Create_FailedConnection(t *testing.T) {
	connector := new(mocks.MongoConnector)
	connector.On("Connect", mock.Anything).Return(nil, errors.New("test"))
	svc := &StubSpotService{
		Connector: connector,
	}
	_, err := svc.Create(context.TODO(), spotsdomain.Spot{})
	assert.Equal(t, err, errors.New("test"))

}

func TestStubSpotService_Create_FailedInsertOne(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	spot := spotsdomain.Spot{}
	collection.On("InsertOne", mock.Anything, spot).Return(nil, errors.New("test"))
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubSpotService{
		Connector: connector,
	}
	_, err := svc.Create(context.TODO(), spotsdomain.Spot{})
	assert.Equal(t, err, errors.New("test"))

}

func TestStubSpotService_Create_Ok(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	spot := spotsdomain.Spot{}
	id := primitive.NewObjectID()
	collection.On("InsertOne", mock.Anything, spot).Return(&mongo.InsertOneResult{InsertedID: id}, nil)
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubSpotService{
		Connector: connector,
	}
	res, err := svc.Create(context.TODO(), spotsdomain.Spot{})
	assert.Nil(t, err)
	assert.Equal(t, id, res.ID)
}

func TestStubSpotService_Get_FailedConnection(t *testing.T) {
	connector := new(mocks.MongoConnector)
	connector.On("Connect", mock.Anything).Return(nil, errors.New("test"))
	svc := &StubSpotService{
		Connector: connector,
	}
	_, err := svc.Get(context.TODO(), "id")
	assert.Equal(t, err, errors.New("test"))

}

func TestStubSpotService_Get_Ok_NotFound(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	id := primitive.NewObjectID()
	filter := bson.M{"_id": id}
	collection.On("FindOne", mock.Anything, filter).Return(&mongo.SingleResult{})
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubSpotService{
		Connector: connector,
	}
	_, err := svc.Get(context.TODO(), id.Hex())
	assert.Equal(t, err, errors.New("mongo: no documents in result"))
}

func TestStubSpotService_Delete_FailedConnection(t *testing.T) {
	connector := new(mocks.MongoConnector)
	connector.On("Connect", mock.Anything).Return(nil, errors.New("test"))
	svc := &StubSpotService{
		Connector: connector,
	}
	err := svc.Delete(context.TODO(), "id")
	assert.Equal(t, err, errors.New("test"))

}

func TestStubSpotService_Delete_FailedDeleteOne(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	id := primitive.NewObjectID()
	filter := bson.M{"_id": id}
	collection.On("DeleteOne", mock.Anything, filter).Return(nil, errors.New("test"))
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubSpotService{
		Connector: connector,
	}
	err := svc.Delete(context.TODO(), id.Hex())
	assert.Equal(t, err, errors.New("test"))

}

func TestStubSpotService_Delete_Ok(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	id := primitive.NewObjectID()
	filter := bson.M{"_id": id}
	collection.On("DeleteOne", mock.Anything, filter).Return(nil, nil)
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubSpotService{
		Connector: connector,
	}
	err := svc.Delete(context.TODO(), id.Hex())
	assert.Nil(t, err)
}

func TestStubSpotService_Update_FailedConnection(t *testing.T) {
	connector := new(mocks.MongoConnector)
	connector.On("Connect", mock.Anything).Return(nil, errors.New("test"))
	svc := &StubSpotService{
		Connector: connector,
	}
	_, err := svc.Update(context.TODO(), spotsdomain.Spot{})
	assert.Equal(t, err, errors.New("test"))

}

func TestStubSpotService_Update_FailedInsertOne(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	id := primitive.NewObjectID()
	filter := bson.M{"_id": id}
	spot := spotsdomain.Spot{ID: id}
	collection.On("ReplaceOne", mock.Anything, filter, spot).Return(nil, errors.New("test"))
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubSpotService{
		Connector: connector,
	}
	_, err := svc.Update(context.TODO(), spot)
	assert.Equal(t, err, errors.New("test"))

}

func TestStubSpotService_Update_Ok(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	id := primitive.NewObjectID()
	filter := bson.M{"_id": id}
	spot := spotsdomain.Spot{ID: id}
	collection.On("ReplaceOne", mock.Anything, filter, spot).Return(nil, nil)
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubSpotService{
		Connector: connector,
	}
	_, err := svc.Update(context.TODO(), spot)
	assert.Nil(t, err)
}
