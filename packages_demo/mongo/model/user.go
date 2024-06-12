package model

import (
	"context"
	"fmt"
	"mongo/db"
	"time"

	"github.com/jaswdr/faker/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Age       int                `bson:"age,omitempty,truncate"`
	Email     string             `bson:"email"`
	Address   Address            `bson:"address,inline"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
}

type Address struct {
	Street string
	City   string
}

type UserCollection struct {
	Collection *mongo.Collection
}

func NewUserCollection(db db.MongoDB) *UserCollection {
	return &UserCollection{
		Collection: db.Client.Database("demo").Collection("users"),
	}
}

func (u *UserCollection) CreateUser(user User) error {
	r, err := u.Collection.InsertOne(context.Background(), user)
	fmt.Printf("r: %v\n", r)
	return err
}

func (u *UserCollection) CreateUsers(users []interface{}) error {
	r, err := u.Collection.InsertMany(context.Background(), users)
	fmt.Printf("Documents inserted: %v\n", len(r.InsertedIDs))
	return err
}

func NewFakeUser(fake faker.Faker) User {
	address := Address{
		City:   fake.Address().City(),
		Street: fake.Address().StreetAddress(),
	}
	return User{
		Name:      fake.Person().Name(),
		Age:       fake.RandomNumber(2),
		Email:     fake.Person().Contact().Email,
		Address:   address,
		CreatedAt: time.Now(),
	}
}

func (u *UserCollection) ReplaceUser(name string, user User) error {
	filter := bson.M{"name": name}
	replacement := user
	r, err := u.Collection.ReplaceOne(context.Background(), filter, replacement)
	fmt.Printf("Matched %v documents and replaced %v documents.\n", r.MatchedCount, r.ModifiedCount)
	return err
}

func (u *UserCollection) UpdateUser(name string) error {
	filter := bson.M{"name": name}
	update := bson.M{"$set": bson.M{"created_at": time.Now()}}
	r, err := u.Collection.UpdateOne(context.Background(), filter, update)
	fmt.Printf("Matched %v documents and updated %v documents.\n", r.MatchedCount, r.ModifiedCount)
	return err
}

func (u *UserCollection) DeleteUser(name string) error {
	filter := bson.M{"name": name}
	r, err := u.Collection.DeleteOne(context.Background(), filter)
	fmt.Printf("Delete %v documents.\n", r.DeletedCount)
	return err
}

func (u *UserCollection) FindUser(name string) (User, error) {
	filter := bson.M{"name": name}
	var r User
	err := u.Collection.FindOne(context.Background(), filter).Decode(&r)
	fmt.Println(r)
	return r, err
}
