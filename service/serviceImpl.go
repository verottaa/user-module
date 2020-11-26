package service

import (
	"github.com/verottaa/user-module/entity"
	"github.com/verottaa/user-module/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
)

var once sync.Once
var destroyCh = make(chan bool)

var serviceInstance *service

type service struct {
	repo repository.Repository
}

func GetService() Service {
	once.Do(func() {
		serviceInstance = createService()
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

	return serviceInstance
}

func createService() *service {
	var service = new(service)
	service.repo = repository.GetRepository()
	return service
}

func (s service) Destroy() {
	destroyCh <- true
	close(destroyCh)
	serviceInstance = nil
}

func (s service) Find(id primitive.ObjectID) (*entity.User, error) {
	return s.repo.Find(id)
}

func (s service) FindAll() ([]*entity.User, error) {
	return s.repo.FindAll()
}

func (s service) Update(user *entity.User) error {
	return s.repo.Update(user)
}

func (s service) Store(user *entity.User) (primitive.ObjectID, error) {
	return s.repo.Store(user)
}

func (s service) Delete(user *entity.User) error {
	return s.repo.Delete(user)
}

func (s service) DeleteAll() error {
	return s.repo.DeleteAll()
}
