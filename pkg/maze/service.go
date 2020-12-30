package maze

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MazeService implements yor service methods.
type MazeService interface {
	Create(ctx context.Context, maze Maze) (Maze, error)
	Get(ctx context.Context, mazeID string) (Maze, error)
	Delete(ctx context.Context, mazeID string) error
	Update(ctx context.Context, maze Maze) (Maze, error)
}

type stubMazeService struct {
	client   *mongo.Client
	database string
}

//NewMazeService return a new instance of the service.
func NewMazeService() (s MazeService) {
	//this could be on a config
	client, err := mongo.NewClient(options.Client().ApplyURI(
		"mongodb+srv://admin:admin@appetizer.rtdgx.mongodb.net/Appetizer?retryWrites=true&w=majority",
	))

	if err != nil {
		panic(err)
	}
	s = &stubMazeService{client, "Appetizer"}
	return s
}

//Create create method
func (ma *stubMazeService) Create(ctx context.Context, maze Maze) (res Maze, err error) {
	if err = ma.client.Connect(ctx); err != nil {
		return maze, err
	}
	defer ma.client.Disconnect(ctx)
	var mres *mongo.InsertOneResult
	mres, err = ma.client.Database(ma.database).Collection("Mazes").InsertOne(ctx, maze)
	maze.ID = mres
	return maze, err
}

// Implement the business logic of Create
func (ma *stubMazeService) Get(ctx context.Context, mazeID string) (m Maze, err error) {
	if err = ma.client.Connect(ctx); err != nil {
		return m, err
	}
	defer ma.client.Disconnect(ctx)

	var mres *mongo.SingleResult
	filter := bson.M{"_id": mazeID}
	collection := ma.client.Database(ma.database).Collection("Mazes")
	mres = collection.FindOne(ctx, filter)
	if mres.Err() != nil {
		return m, mres.Err()
	}

	err = mres.Decode(&m)
	return m, err
}

// Implement the business logic of Delete
func (ma *stubMazeService) Delete(ctx context.Context, mazeID string) (err error) {
	if err = ma.client.Connect(ctx); err != nil {
		return err
	}
	defer ma.client.Disconnect(ctx)
	filter := bson.M{"_id": mazeID}
	collection := ma.client.Database(ma.database).Collection("Mazes")
	if _, err = collection.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return
}

// Implement the business logic of Update
func (ma *stubMazeService) Update(ctx context.Context, maze Maze) (m Maze, err error) {
	if err = ma.client.Connect(ctx); err != nil {
		return maze, err
	}
	defer ma.client.Disconnect(ctx)
	filter := bson.M{"_id": maze.ID}
	var mres *mongo.UpdateResult
	collection := ma.client.Database(ma.database).Collection("Mazes")
	mres, err = collection.UpdateOne(ctx, filter, maze)
	maze.ID = mres
	return maze, err
}
