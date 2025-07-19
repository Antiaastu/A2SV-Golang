package data
import(
	"context"
	"errors"
	"time"
	"task_manager_5/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var taskCollection *mongo.Collection

func InitMongoDB(ctx context.Context, uri string) error{
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil{
		return err
	}
	taskCollection = client.Database("task_manager").Collection("tasks")
	return nil
}

func CloseMongoDB(){
	if client != nil{
		_ = client.Disconnect(context.TODO())
	}
}

func GetTasks() ([]models.Task, error){
	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	cursor, err := taskCollection.Find(ctx, bson.M{})
	if err != nil{
		return nil, err
	}
	//the cursor will be closed after reading in which it prevents memory leaks
	defer cursor.Close(ctx) 

	var tasks []models.Task
	for cursor.Next(ctx){
		var task models.Task
		if err := cursor.Decode(&task); err != nil{
			return nil, err
		}
		tasks = append(tasks,task)
	}
	return tasks, nil
}

func GetTaskById(id string) (models.Task, error){
	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return models.Task{}, err
	}

	var task models.Task
	err = taskCollection.FindOne(ctx, bson.M{"_id":objID}).Decode(&task)
	if err == mongo.ErrNoDocuments{
		return models.Task{}, errors.New("task not found")
	}
	return task, err
}

func CreateTask(task models.Task) (models.Task,error){
	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	task.ID = primitive.NewObjectID()
	_, err := taskCollection.InsertOne(ctx, task)
	return task, err
}

func UpdateTask(id string, updatedTask models.Task) (models.Task, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return models.Task{}, err
	}

	update := bson.M{
		"$set": bson.M{
			"title": updatedTask.Title,
			"description": updatedTask.Description,
			"due_date": updatedTask.DueDate,
			"status": updatedTask.Status,
		},
	}

	result := taskCollection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil{
		return models.Task{}, result.Err()
	}
	var task models.Task
	err = result.Decode(&task)
	return task, err
}

func DeleteTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return err
	}
	res, err := taskCollection.DeleteOne(ctx,bson.M{"_id":objID})
	if err != nil{
		return err
	}
	if res.DeletedCount == 0{
		return errors.New("task not found")
	}
	return nil
}