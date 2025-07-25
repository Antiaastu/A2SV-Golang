package repositories

import (
	"Task7/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository implements the UserRepository interface for MongoDB.
type UserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepo(c *mongo.Collection) *UserRepository {
	return &UserRepository{collection: c}
}

func (r *UserRepository) Create(ctx context.Context, user domain.User) error {
	res, err := r.collection.InsertOne(ctx, bson.M{
		"username": user.Username,
		"password": user.Password,
		"role":     user.Role,
	})
	if err != nil {
		return err
	}
	user.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	var doc bson.M
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}
	user := domain.User{
		ID:       doc["_id"].(primitive.ObjectID).Hex(),
		Username: doc["username"].(string),
		Password: doc["password"].(string),
		Role:     doc["role"].(string),
	}
	return user, nil
}
