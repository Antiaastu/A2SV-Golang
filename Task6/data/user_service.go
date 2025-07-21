package data

import(
	"context"
	"errors"
	"time"
	"Task6/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func InitUserCollection(){
	userCollection = client.Database("task_manager_6").Collection("users")
}

func Register(user models.User) error{
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	//check if username already exists
	count, _ := userCollection.CountDocuments(ctx, bson.M{"username":user.Username})
	if count > 0{
		return errors.New("username already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil{
		return err
	}
	user.Password = string(hashed)
	_, err = userCollection.InsertOne(ctx, user)
	return err
}

func Login(username, password string) (models.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"username":username}).Decode(&user)

	if err != nil{
		return models.User{}, errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	if err != nil{
		return models.User{}, errors.New("invalid username or password")
	}
	return user, nil
}