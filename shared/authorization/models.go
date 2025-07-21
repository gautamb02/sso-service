package authorization

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserCreateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	Verified  bool   `json:"verified"`
}

type UserDetail struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	UserCreateRequest
}

type UserModelTypes interface {
	UserCreateRequest | UserDetail
}
