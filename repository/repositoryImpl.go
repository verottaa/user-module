package repository

import (
	"github.com/verottaa/user-module/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
)

var once sync.Once
var destroyCh = make(chan bool)

var repositoryInstance *repository

type repository struct {
}

func GetRepository() Repository {
	once.Do(func() {
		repositoryInstance = createRepository()
		go func() {
			for
			{
				select {
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

func (r repository) Find(id primitive.ObjectID) (*entity.User, error) {
	panic("implement me")
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
