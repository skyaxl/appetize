package quadrants

import (
	"context"
	"errors"
	"testing"

	"github.com/skyaxl/synack/mocks"
	"github.com/skyaxl/synack/pkg/maze/quadrants/quadrantsdomain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestStubQuadrantService_Create_FailedConnection(t *testing.T) {
	connector := new(mocks.MongoConnector)
	connector.On("Connect", mock.Anything).Return(nil, errors.New("test"))
	svc := &StubQuadrantService{
		Connector: connector,
	}
	_, err := svc.Create(context.TODO(), quadrantsdomain.Quadrant{})
	assert.Equal(t, err, errors.New("test"))

}

func TestStubQuadrantService_Create_FailedInsertOne(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	quadrant := quadrantsdomain.Quadrant{}
	collection.On("InsertOne", mock.Anything, quadrant).Return(nil, errors.New("test"))
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubQuadrantService{
		Connector: connector,
	}
	_, err := svc.Create(context.TODO(), quadrantsdomain.Quadrant{})
	assert.Equal(t, err, errors.New("test"))

}

func TestStubQuadrantService_Create_Ok(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	quadrant := quadrantsdomain.Quadrant{}
	id := primitive.NewObjectID()
	collection.On("InsertOne", mock.Anything, quadrant).Return(&mongo.InsertOneResult{InsertedID: id}, nil)
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubQuadrantService{
		Connector: connector,
	}
	res, err := svc.Create(context.TODO(), quadrantsdomain.Quadrant{})
	assert.Nil(t, err)
	assert.Equal(t, id, res.ID)
}

func TestStubQuadrantService_Get_FailedConnection(t *testing.T) {
	connector := new(mocks.MongoConnector)
	connector.On("Connect", mock.Anything).Return(nil, errors.New("test"))
	svc := &StubQuadrantService{
		Connector: connector,
	}
	_, err := svc.Get(context.TODO(), "id")
	assert.Equal(t, err, errors.New("test"))

}

func TestStubQuadrantService_Get_Ok_NotFound(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	id := primitive.NewObjectID()
	filter := bson.M{"_id": id}
	collection.On("FindOne", mock.Anything, filter).Return(&mongo.SingleResult{})
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubQuadrantService{
		Connector: connector,
	}
	_, err := svc.Get(context.TODO(), id.Hex())
	assert.Equal(t, err, errors.New("mongo: no documents in result"))
}

func TestStubQuadrantService_Delete_FailedConnection(t *testing.T) {
	connector := new(mocks.MongoConnector)
	connector.On("Connect", mock.Anything).Return(nil, errors.New("test"))
	svc := &StubQuadrantService{
		Connector: connector,
	}
	err := svc.Delete(context.TODO(), "id")
	assert.Equal(t, err, errors.New("test"))

}

func TestStubQuadrantService_Delete_FailedDeleteOne(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	id := primitive.NewObjectID()
	filter := bson.M{"_id": id}
	collection.On("DeleteOne", mock.Anything, filter).Return(nil, errors.New("test"))
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubQuadrantService{
		Connector: connector,
	}
	err := svc.Delete(context.TODO(), id.Hex())
	assert.Equal(t, err, errors.New("test"))

}

func TestStubQuadrantService_Delete_Ok(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	id := primitive.NewObjectID()
	filter := bson.M{"_id": id}
	collection.On("DeleteOne", mock.Anything, filter).Return(nil, nil)
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubQuadrantService{
		Connector: connector,
	}
	err := svc.Delete(context.TODO(), id.Hex())
	assert.Nil(t, err)
}

func TestStubQuadrantService_Update_FailedConnection(t *testing.T) {
	connector := new(mocks.MongoConnector)
	connector.On("Connect", mock.Anything).Return(nil, errors.New("test"))
	svc := &StubQuadrantService{
		Connector: connector,
	}
	_, err := svc.Update(context.TODO(), quadrantsdomain.Quadrant{})
	assert.Equal(t, err, errors.New("test"))

}

func TestStubQuadrantService_Update_FailedInsertOne(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	id := primitive.NewObjectID()
	filter := bson.M{"_id": id}
	quadrant := quadrantsdomain.Quadrant{ID: id}
	collection.On("ReplaceOne", mock.Anything, filter, quadrant).Return(nil, errors.New("test"))
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubQuadrantService{
		Connector: connector,
	}
	_, err := svc.Update(context.TODO(), quadrant)
	assert.Equal(t, err, errors.New("test"))

}

func TestStubQuadrantService_Update_Ok(t *testing.T) {
	connector := new(mocks.MongoConnector)
	collection := new(mocks.MongoCollection)
	id := primitive.NewObjectID()
	filter := bson.M{"_id": id}
	quadrant := quadrantsdomain.Quadrant{ID: id}
	collection.On("ReplaceOne", mock.Anything, filter, quadrant).Return(nil, nil)
	connector.On("Connect", mock.Anything).Return(collection, nil)
	svc := &StubQuadrantService{
		Connector: connector,
	}
	_, err := svc.Update(context.TODO(), quadrant)
	assert.Nil(t, err)
}
