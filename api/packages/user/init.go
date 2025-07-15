package user

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserModule(db *mongo.Database) *UserHandler {
	r := NewUserRepository(db)
	s := NewUserService(r)
	h := NewUserHandler(s)
	return h
}
