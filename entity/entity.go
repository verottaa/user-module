package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	Name       string             `json:"name" bson:"name"`
	Surname    string             `json:"surname" bson:"surname"`
	Patronymic string             `json:"patronymic" bson:"patronymic"`
	Email      string             `json:"email" bson:"email"`
	Password   string             `json:"password" bson:"password"`
	Type       primitive.ObjectID `json:"type" bson:"type"`
	Branch     primitive.ObjectID `json:"branch" bson:"branch"`
	Department primitive.ObjectID `json:"department" bson:"department"`
	Position   primitive.ObjectID `json:"position" bson:"position"`
}
