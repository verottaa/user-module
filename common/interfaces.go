package common

import (
	"github.com/verottaa/user-module/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reader interface {
	Find(id primitive.ObjectID) (*entity.User, error)
	FindAll() ([]*entity.User, error)
}

type Writer interface {
	Update(user *entity.User) error
	Store(user *entity.User) (primitive.ObjectID, error)
	Delete(id primitive.ObjectID) error
	DeleteMany(filter entity.UserFilter) (int64, error)
	DeleteAll() (int64, error)
}

type Destroyable interface {
	Destroy()
}
