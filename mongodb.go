package Sample

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

func newMongoDB()(*mongo.Client,error){
	client,err := mongo.NewClient(options.Client().ApplyURI("<mongo.Client_URI>"))
	if err != nil {
		return nil,errors.Wrap(err,"mongo db client")
	}
	return client,nil
}

func newMongoDBDIProvider()func() (*mongo.Client, error){
	var cl *mongo.Client
	var err error
	var mu sync.Mutex
	return func() (*mongo.Client, error) {
		mu.Lock()
		defer mu.Unlock()
		if cl == nil {
			cl,err = newMongoDB()
		}

		return cl,err
	}
}
