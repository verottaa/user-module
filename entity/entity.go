package entity

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

type UserFilter struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Surname    string             `json:"surname" bson:"surname,omitempty"`
	Patronymic string             `json:"patronymic" bson:"patronymic,omitempty"`
	Email      string             `json:"email" bson:"email,omitempty"`
	Password   string             `json:"password" bson:"password,omitempty"`
	Type       primitive.ObjectID `json:"type" bson:"type,omitempty"`
	Branch     primitive.ObjectID `json:"branch" bson:"branch,omitempty"`
	Department primitive.ObjectID `json:"department" bson:"department,omitempty"`
	Position   primitive.ObjectID `json:"position" bson:"position,omitempty"`
}

type ObjectCreatedDto struct {
	Id string `json:"id"`
}

func (u User) GetUserFromBson(data interface{}) error {
	m, _ := bson.Marshal(data)
	err := bson.Unmarshal(m, &u)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (u User) ToUpdateObjectData() bson.D {
	update := bson.D{{"$set",
		bson.D{
			{"name", u.Name},
			{"surname", u.Surname},
			{"patronymic", u.Patronymic},
			{"email", u.Email},
			{"password", u.Password},
			{"type", u.Type},
			{"branch", u.Branch},
			{"department", u.Department},
			{"position", u.Position},
		},
	}}
	return update
}

func (u User) ToUpdateObjectBasicData() bson.D {
	update := bson.D{{"$set",
		bson.D{
			{"name", u.Name},
			{"surname", u.Surname},
			{"patronymic", u.Patronymic},
		},
	}}
	return update
}

func (u User) ToUpdateObjectAuthenticationData() bson.D {
	update := bson.D{{"$set",
		bson.D{
			{"email", u.Email},
			{"password", u.Password},
		},
	}}
	return update
}

func (u User) ToUpdateObjectTypeData() bson.D {
	update := bson.D{{"$set",
		bson.D{
			{"type", u.Type},
		},
	}}
	return update
}

func (u User) ToUpdateObjectEmployeeData() bson.D {
	update := bson.D{{"$set",
		bson.D{
			{"branch", u.Branch},
			{"department", u.Department},
			{"position", u.Position},
		},
	}}
	return update
}