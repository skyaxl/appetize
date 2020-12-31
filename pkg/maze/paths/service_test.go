package paths

import (
	"context"
	"errors"
	"testing"

	"github.com/skyaxl/synack/mocks"
	"github.com/skyaxl/synack/pkg/maze/paths/pathsdomain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestStubPathService_Create_FailedConnection(t *testing.T) {
	connector := new(mocks.MongoConnector)
	connector.On("Connect", mock.Anything).Return(nil, errors.New("test"))
	svc := &StubPathService{
		Connector: connector,
	}
	_, err := svc.Create(context.TODO(), pathsdomain.Path{})
	assert.Equal(t, err, errors.New("test"))

}

func TestStubPathService_Create_FailedInsertOne(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	path := pathsdomain.Path{}
	collection.On("InsertOne", mock.Anything, path).Return(nil, errors.New("test"))
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubPathService{
		Connector: connector,
	}
	_, err := svc.Create(context.TODO(), pathsdomain.Path{})
	assert.Equal(t, err, errors.New("test"))

}

func TestStubPathService_Create_Ok(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	path := pathsdomain.Path{}
	id := primitive.NewObjectID()
	collection.On("InsertOne", mock.Anything, path).Return(&mongo.InsertOneResult{InsertedID: id}, nil)
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubPathService{
		Connector: connector,
	}
	res, err := svc.Create(context.TODO(), pathsdomain.Path{})
	assert.Nil(t, err)
	assert.Equal(t, id, res.ID)
}

func TestStubPathService_Get_FailedConnection(t *testing.T) {
	connector := new(mocks.MongoConnector)
	connector.On("Connect", mock.Anything).Return(nil, errors.New("test"))
	svc := &StubPathService{
		Connector: connector,
	}
	_, err := svc.Get(context.TODO(), "id")
	assert.Equal(t, err, errors.New("test"))

}

func TestStubPathService_Get_Ok_NotFound(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	id := primitive.NewObjectID()
	filter := bson.M{"_id": id}
	collection.On("FindOne", mock.Anything, filter).Return(&mongo.SingleResult{})
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubPathService{
		Connector: connector,
	}
	_, err := svc.Get(context.TODO(), id.Hex())
	assert.Equal(t, err, errors.New("mongo: no documents in result"))
}

func TestStubPathService_Delete_FailedConnection(t *testing.T) {
	connector := new(mocks.MongoConnector)
	connector.On("Connect", mock.Anything).Return(nil, errors.New("test"))
	svc := &StubPathService{
		Connector: connector,
	}
	err := svc.Delete(context.TODO(), "id")
	assert.Equal(t, err, errors.New("test"))

}

func TestStubPathService_Delete_FailedDeleteOne(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	id := primitive.NewObjectID()
	filter := bson.M{"_id": id}
	collection.On("DeleteOne", mock.Anything, filter).Return(nil, errors.New("test"))
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubPathService{
		Connector: connector,
	}
	err := svc.Delete(context.TODO(), id.Hex())
	assert.Equal(t, err, errors.New("test"))

}

func TestStubPathService_Delete_Ok(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	id := primitive.NewObjectID()
	filter := bson.M{"_id": id}
	collection.On("DeleteOne", mock.Anything, filter).Return(nil, nil)
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubPathService{
		Connector: connector,
	}
	err := svc.Delete(context.TODO(), id.Hex())
	assert.Nil(t, err)
}

func TestStubPathService_Update_FailedConnection(t *testing.T) {
	connector := new(mocks.MongoConnector)
	connector.On("Connect", mock.Anything).Return(nil, errors.New("test"))
	svc := &StubPathService{
		Connector: connector,
	}
	_, err := svc.Update(context.TODO(), pathsdomain.Path{})
	assert.Equal(t, err, errors.New("test"))

}

func TestStubPathService_Update_FailedInsertOne(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	id := primitive.NewObjectID()
	filter := bson.M{"_id": id}
	path := pathsdomain.Path{ID: id}
	collection.On("ReplaceOne", mock.Anything, filter, path).Return(nil, errors.New("test"))
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubPathService{
		Connector: connector,
	}
	_, err := svc.Update(context.TODO(), path)
	assert.Equal(t, err, errors.New("test"))

}

func TestStubPathService_Update_Ok(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	id := primitive.NewObjectID()
	filter := bson.M{"_id": id}
	path := pathsdomain.Path{ID: id}
	collection.On("ReplaceOne", mock.Anything, filter, path).Return(nil, nil)
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubPathService{
		Connector: connector,
	}
	_, err := svc.Update(context.TODO(), path)
	assert.Nil(t, err)
}
