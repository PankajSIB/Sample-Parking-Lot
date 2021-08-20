package Sample

import "go.mongodb.org/mongo-driver/mongo"

type DIContainer struct {
	mongoClient func()(*mongo.Client,error)
}


func NewCommonDIContainer() *DIContainer{
	return &DIContainer{
		mongoClient: newMongoDBDIProvider(),
	}
}
func (dic *DIContainer) MongoDBClient() (*mongo.Client, error) {
	return dic.mongoClient()
}