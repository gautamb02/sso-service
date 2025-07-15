package user

import "go.mongodb.org/mongo-driver/mongo"

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepositoryI {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) RegisterUser() {

}
