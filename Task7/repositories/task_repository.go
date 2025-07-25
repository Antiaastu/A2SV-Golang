package repositories

import (
	"Task7/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type taskRepository struct {
	collection *mongo.Collection
}

func NewMongoTaskRepo(c *mongo.Collection) *taskRepository {
	return &taskRepository{collection: c}
}

func (r *taskRepository) Create(ctx context.Context, task domain.Task) (domain.Task, error) {
	res, err := r.collection.InsertOne(ctx, bson.M{
		"_id":         primitive.NewObjectID(),
		"title":       task.Title,
		"description": task.Description,
		"status":      task.Status,
		"due_date":    task.DueDate,
		"owner_id":    task.OwnerID,
	})
	if err != nil {
		return domain.Task{}, err
	}
	task.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return task, nil
}

func (r *taskRepository) GetAll(ctx context.Context, ownerID string) ([]domain.Task, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"owner_id": ownerID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []domain.Task
	for cursor.Next(ctx) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}

		task := domain.Task{
			ID:          doc["_id"].(primitive.ObjectID).Hex(),
			Title:       doc["title"].(string),
			Description: doc["description"].(string),
			Status:      doc["status"].(string),
			DueDate:     doc["due_date"].(primitive.DateTime).Time(),
			OwnerID:     doc["owner_id"].(string),
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) GetByID(ctx context.Context, taskID string) (domain.Task, error) {
	oid, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return domain.Task{}, err
	}
	var doc bson.M
	err = r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Task{}, errors.New("task not found")
		}
		return domain.Task{}, err
	}
	task := domain.Task{
		ID:          doc["_id"].(primitive.ObjectID).Hex(),
		Title:       doc["title"].(string),
		Description: doc["description"].(string),
		Status:      doc["status"].(string),
		DueDate:     doc["due_date"].(primitive.DateTime).Time().UTC(),
		OwnerID:     doc["owner_id"].(string),
	}
	return task, nil
}

func (r *taskRepository) Update(ctx context.Context, task domain.Task) (domain.Task, error) {
	oid, err := primitive.ObjectIDFromHex(task.ID)
	if err != nil {
		return domain.Task{}, err
	}
	update := bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
			"status":      task.Status,
			"due_date":    task.DueDate,
			"owner_id":    task.OwnerID,
		},
	}
	res := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": oid}, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	var doc bson.M
	if err := res.Decode(&doc); err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Task{}, errors.New("task not found")
		}
		return domain.Task{}, err
	}
	t := domain.Task{
		ID:          doc["_id"].(primitive.ObjectID).Hex(),
		Title:       doc["title"].(string),
		Description: doc["description"].(string),
		Status:      doc["status"].(string),
		DueDate:     doc["due_date"].(primitive.DateTime).Time().UTC(),
		OwnerID:     doc["owner_id"].(string),
	}
	return t, nil
}

func (r *taskRepository) Delete(ctx context.Context, taskID string) error {
	oid, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return err
	}
	res, err := r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
