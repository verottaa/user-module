package repository

import (
	"errors"
	"fmt"
	database_module "github.com/verottaa/database-module"
	"github.com/verottaa/user-module/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
)

var once sync.Once
var destroyCh = make(chan bool)
var getCollectionNameCh = make(chan chan string)

var repositoryInstance *repository

type repository struct {
	collectionName string
}

func GetRepository() Repository {
	once.Do(func() {
		repositoryInstance = createRepository()
		go func() {
			for
			{
				select {
				case ch := <-getCollectionNameCh:
					ch <- repositoryInstance.collectionName
				case <-destroyCh:
					return
				}
			}
		}()
	})

	return repositoryInstance
}

func createRepository() *repository {
	return new(repository)
}

func (r repository) Destroy() {
	destroyCh <- true
	close(destroyCh)
	repositoryInstance = nil
}

func (r repository) GetCollectionName() string {
	resCh := make(chan string)
	defer close(resCh)
	getCollectionNameCh <- resCh
	return <-resCh
}

func (r repository) Find(id primitive.ObjectID) (*entity.User, error) {
	var filter = entity.UserFilter{
		Id: id,
	}
	code, bsonUser := database_module.FindOne(r.GetCollectionName(), filter)
	switch code {
	case database_module.FOUND:
		return decodeUserFromBson(bsonUser)
	case database_module.ERROR:
		return nil, errors.New("unexpected error with database connection")
	default:
		return nil, nil
	}
}

func decodeUserFromBson(bsonUser interface{}) (*entity.User, error) {
	var user entity.User
	err := user.GetUserFromBson(bsonUser)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &user, nil
}

func (r repository) FindAll() ([]*entity.User, error) {
	panic("implement me")
}

func (r repository) Update(user *entity.User) error {
	panic("implement me")
}

func (r repository) Store(user *entity.User) (primitive.ObjectID, error) {
	panic("implement me")
}

func (r repository) Delete(user *entity.User) error {
	panic("implement me")
}

func (r repository) DeleteAll() error {
	panic("implement me")
}
