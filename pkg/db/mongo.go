package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strux_api/internal/config"
	"sync"
)

type ErrClientInstanceAlreadyCreated struct {
}

func (e *ErrClientInstanceAlreadyCreated) Error() string {
	return "Single client instance already created."
}

type Client struct {
	Options *options.ClientOptions
	Ctx     context.Context
	client  *mongo.Client
}

// Connect Connects to a database and returns connections
func (c Client) Connect() (*mongo.Client, context.Context, error) {
	var err error
	c.client, err = mongo.Connect(c.Ctx, c.Options)
	if err != nil {
		return nil, nil, err
	}
	err = c.client.Ping(c.Ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	return c.client, c.Ctx, err
}

var lock = &sync.Mutex{}
var singleInstance *Client

// GetMongoClient A singleton that creates or returns a merge to connect to the database.
func GetMongoClient() (*Client, error) {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			ctx := context.TODO()
			clientOptions := options.Client().ApplyURI(config.MongoUrl)
			singleInstance = &Client{
				Options: clientOptions,
				Ctx:     ctx,
			}
		} else {
			return nil, &ErrClientInstanceAlreadyCreated{}
		}
	} else {
		return singleInstance, nil
	}
	return singleInstance, nil
}

type DatabaseOperation struct {
	DbName         string
	CollectionName string
	Client         *mongo.Client
	Ctx            context.Context
}

func (do *DatabaseOperation) InsertOne(schema interface{}) (*mongo.InsertOneResult, error) {
	collection := do.Client.Database(do.DbName).Collection(do.CollectionName)
	res, err := collection.InsertOne(do.Ctx, schema)
	return res, err
}

func (do *DatabaseOperation) DropDatabase() error {
	db := do.Client.Database(do.DbName)
	err := db.Drop(do.Ctx)
	return err
}

func (do *DatabaseOperation) DropCollection() error {
	collection := do.Client.Database(do.DbName).Collection(do.CollectionName)
	err := collection.Drop(do.Ctx)
	return err
}

func (do *DatabaseOperation) FindOneByValue(colName string, value string, result interface{}) error {
	collection := do.Client.Database(do.DbName).Collection(do.CollectionName)
	filter := bson.D{{colName, value}}
	err := collection.FindOne(do.Ctx, filter).Decode(result)
	return err
}
