package model

import (
	"context"
	"fmt"
	"mongo/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Relationship struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	FollowerID  primitive.ObjectID `bson:"follower_id,omitempty"`
	FollowingID primitive.ObjectID `bson:"following_id,omitempty"`
}

type RelationshipCollection struct {
	Collection *mongo.Collection
}

func NewRelationshipCollection(db db.MongoDB) *RelationshipCollection {
	return &RelationshipCollection{
		Collection: db.Client.Database("demo").Collection("relationships"),
	}
}

func (r *RelationshipCollection) FollowUser(followerID, followingID primitive.ObjectID) error {
	relationship := Relationship{
		FollowerID:  followerID,
		FollowingID: followingID,
	}

	result, err := r.Collection.InsertOne(context.Background(), relationship)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("inserted: %v\n", result.InsertedID)
	return nil
}

func (r *RelationshipCollection) GetFollowers(userID primitive.ObjectID) (followers []User, err error) {
	ctx := context.Background()
	cursor, err := r.Collection.Find(ctx, bson.M{"following_id": userID})
	if err != nil {
		return nil, fmt.Errorf("fail to find followers: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var u User
		if err := cursor.Decode(&u); err != nil {
			return nil, fmt.Errorf("fail to decode followers: %w", err)
		}
		followers = append(followers, u)
	}
	return
}
