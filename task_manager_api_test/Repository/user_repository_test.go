package Repository

import (
	"task_manager_api_test/Domain"
	"task_manager_api_test/Repository/repositoryMocks"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// this is repository level test. to just test the business logic with out depending on
// the datbase.
func TestUserRepositoryInterface(t *testing.T){

	t.Run("StoreUser Success", func (t *testing.T){
		mockRepo := new(repositoryMocks.UserRepoSitoryMock)
		testUser := &Domain.User{
			Email:"ed@gmail.com",
			Password:"123",
			Role:"admin",
		}

		mockRepo.On("StoreUser", testUser).Return(nil)

		err := mockRepo.StoreUser(testUser)
		assert.NoError(t,err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("CheckEmailAvailablity Found", func (t *testing.T){
		mockRepo := new(repositoryMocks.UserRepoSitoryMock)

		testEmail := "ed@gmail.com"
		testUser := Domain.User{
			Email:"ed@gmail.com",
			Password:"123",
			Role:"admin",
		}
		mockRepo.On("CheckEmailAvailablity", testEmail). Return (testUser , nil)
		user , err := mockRepo.CheckEmailAvailablity(testEmail)
		assert.NoError(t,err)
		assert.Equal(t, testUser, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("CheckEmailAvailablity NotFound", func(t *testing.T){
		mockRepo := new(repositoryMocks.UserRepoSitoryMock)
		testEmail := "notfoundemail@gmail.com"

		testUser := Domain.User{}

		mockRepo.On("CheckEmailAvailablity", testEmail).Return(testUser, mongo.ErrNoDocuments)

		user , err := mockRepo.CheckEmailAvailablity(testEmail)
		assert.Error(t,err)
		assert.Equal(t,user,testUser)
		mockRepo.AssertExpectations(t)
	})
}





func asMongoCollection (m *repositoryMocks.MongoCollectionMock) *mongo.Collection{
	return (*mongo.Collection)(unsafe.Pointer(m))
}
// this is database level testing to test the repository implementation with mongoDB driver mock.
func TestUserRepositoryMongo(t *testing.T) {
	t.Run("StoreUser Success", func(t *testing.T) {
		mockCollection := repositoryMocks.NewMongoCollectionMock()
		repo := NewUserRepo(asMongoCollection(mockCollection))
		
		testUser := &Domain.User{
			ID:       primitive.NewObjectID(),
			Email:    "ed@gmail.com",
			Password: "123",
			Role:     "admin",
		}

		mockCollection.On(
			"InsertOne", 
			mock.AnythingOfType("*context.timerCtx"),
			testUser, 
			mock.AnythingOfType("[]*options.InsertOneOptions"),
		).Return(&mongo.InsertOneResult{InsertedID: testUser.ID}, nil)

		err := repo.StoreUser(testUser)
		assert.NoError(t, err)
		mockCollection.AssertExpectations(t)
	})

	t.Run("CheckEmailAvailability found", func(t *testing.T) {
		mockCollection := new(repositoryMocks.MongoCollectionMock)
		repo := NewUserRepo(asMongoCollection(mockCollection))
		
		testEmail := "ed@gmail.com"
		expectedUser := Domain.User{Email: testEmail} // Note: Not pointer

		mockCollection.On(
			"FindOne",
			mock.AnythingOfType("*context.timerCtx"), // Fix typo: was "timeCtx"
			bson.M{"email": testEmail},
			mock.AnythingOfType("[]*options.FindOneOptions"), // Add asterisk
		).Return(expectedUser, nil) // Return value, not pointer

		user, err := repo.CheckEmailAvailablity(testEmail)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user) // Compare values
	})
}