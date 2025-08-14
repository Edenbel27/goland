package repositoryMocks

import (
	"task_manager_api_test/Domain"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/stretchr/testify/mock"
)

type MongoCollectionMock struct {
	mock.Mock
	mongo.Collection
}

func NewMongoCollectionMock() *MongoCollectionMock {
	return &MongoCollectionMock{}
}
func (m *MongoCollectionMock) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document, opts)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MongoCollectionMock) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter, opts)

	if args.Get(0) == nil{
		return mongo.NewSingleResultFromDocument(nil, args.Error(1), nil)
	}
	return mongo.NewSingleResultFromDocument(args.Get(0).(Domain.User), args.Error(1), nil)
}