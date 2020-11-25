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
	Delete(user *entity.User) error
	DeleteAll() error
}

type Destroyable interface {
	Destroy()
}
