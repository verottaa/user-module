package repository

import (
	"errors"
	database_module "github.com/verottaa/database-module"
	"github.com/verottaa/user-module/entity"
	"go.mongodb.org/mongo-driver/bson"
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

func decodeBson(bsonObject interface{}, target interface{}) {
	m, _ := bson.Marshal(bsonObject)
	_ = bson.Unmarshal(m, target)
}

func decodeUserFromBson(bsonUser interface{}) (*entity.User, error) {
	var user entity.User
	err := user.GetUserFromBson(bsonUser)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r repository) Find(id primitive.ObjectID) (*entity.User, error) {
	filter := entity.UserFilter{
		Id: id,
	}
	code, bsonUser := database_module.FindOne(r.GetCollectionName(), filter)
	switch code {
	case database_module.FOUND:
		return decodeUserFromBson(bsonUser)
	case database_module.ERROR:
		return nil, errors.New("unexpected error with database connection")
	default:
		return nil, errors.New("unexpected code")
	}
}

func (r repository) FindAll() ([]*entity.User, error) {
	filter := entity.UserFilter{}
	code, bsonUsers := database_module.FindMany(r.GetCollectionName(), filter)
	switch code {
	case database_module.FOUND_ANY:
		var users []*entity.User
		for _, bsonUser := range bsonUsers {
			user, err := decodeUserFromBson(bsonUser)
			if err != nil {
				return nil, err
			}
			users = append(users, user)
		}
		return users, nil
	case database_module.ERROR:
		return nil, errors.New("unexpected error with database connection")
	default:
		return nil, errors.New("unexpected code")
	}
}

func (r repository) Update(user *entity.User) error {
	filter:= entity.UserFilter{
		Id: user.Id,
	}
	code := database_module.UpdateOne(r.GetCollectionName(), filter, user.ToUpdateObjectData())
	switch code {
	case database_module.UPDATED:
		return nil
	case database_module.NOT_FOUND:
		return errors.New("object didn't found in database")
	case database_module.ERROR:
		return errors.New("unexpected error with database connection")
	default:
		return errors.New("unexpected code")
	}
}

func (r repository) Store(user *entity.User) (primitive.ObjectID, error) {
	user.Id = database_module.GenerateObjectID()
	code, bsonId := database_module.PushOne(r.GetCollectionName(), user)
	switch code {
	case database_module.CREATED:
		var id primitive.ObjectID
		decodeBson(bsonId, &id)

		if user.Id.String() == id.String() {
			return id, nil
		}
		return primitive.ObjectID{}, errors.New("validation didn't pass")
	case database_module.ERROR:
		return primitive.ObjectID{}, errors.New("unexpected error with database connection")
	default:
		return primitive.ObjectID{}, errors.New("unexpected code")

	}
}

func (r repository) Delete(id primitive.ObjectID) error {
	filter := entity.UserFilter{
		Id: id,
	}
	code := database_module.DeleteOne(r.GetCollectionName(), filter)
	switch code {
	case database_module.DELETED:
		return nil
	case database_module.NOT_FOUND:
		return errors.New("object didn't found in database")
	case database_module.ERROR:
		return errors.New("unexpected error with database connection")
	default:
		return errors.New("unexpected code")
	}
}

func (r repository) DeleteMany(filter entity.UserFilter) (int64, error) {
	code, quantity := database_module.DeleteMany(r.GetCollectionName(), filter)
	switch code {
	case database_module.DELETED:
		return quantity, nil
	case database_module.NOT_FOUND:
		return quantity, errors.New("object didn't found in database")
	case database_module.ERROR:
		return quantity, errors.New("unexpected error with database connection")
	default:
		return quantity, errors.New("unexpected code")
	}
}

func (r repository) DeleteAll() (int64, error) {
	filter := entity.UserFilter{}
	return r.DeleteMany(filter)
}
