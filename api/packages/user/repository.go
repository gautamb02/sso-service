package user

import (
	"context"
	"fmt"

	"github.com/gautamb02/sso-service/shared"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepositoryI {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) RegisterUser(user *UserDetail, ctx context.Context) (int64, error) {
	if user == nil {
		return 0, fmt.Errorf("user empty data") // or an appropriate error
	}
	collection := r.db.Collection(shared.COLLECTION_USERS)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return 0, err
	}

	return 0, nil

}

func (r *UserRepository) CheckIfEmailExist(email string, ctx context.Context) (int64, error) {
	collection := r.db.Collection(shared.COLLECTION_USERS)

	filter := bson.M{"email": email}

	var result bson.M
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		}
		return 0, err
	}

	return 1, nil
}
