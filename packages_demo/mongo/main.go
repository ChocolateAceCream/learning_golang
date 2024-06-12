package main

import (
	"fmt"
	"mongo/db"
	"mongo/model"

	"github.com/jaswdr/faker/v2"
)

func main() {
	var mongodb db.MongoDB
	mongodb.Init()
	defer mongodb.Disconnect()
	_, err := mongodb.Ping()
	if err != nil {
		fmt.Println("fail to ping mongodb: %w", err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	userCollection := model.NewUserCollection(mongodb)
	f := faker.New()
	fmt.Println(f)
	// insertOne demo
	// user := model.NewFakeUser(f)
	// err = userCollection.CreateUser(user)
	// if err != nil {
	// 	fmt.Println("insertOne user error: ", err)
	// }

	// batch insert demo
	// users := []interface{}{}
	// for i := 0; i < 100; i++ {
	// 	users = append(users, model.NewFakeUser(f))
	// }
	// err = userCollection.CreateUsers(users)
	// if err != nil {
	// 	fmt.Println("insertMany user error: ", err)
	// }

	// replaceOne demo
	u := model.NewFakeUser(f)
	err = userCollection.ReplaceUser("Lila Torp", u)
	if err != nil {
		fmt.Println("ReplaceUser error: ", err)
	}

	// updateOne demo
	err = userCollection.UpdateUser("Addie Mills")
	if err != nil {
		fmt.Println("updateOne error: ", err)
	}

	// deleteOne demo
	err = userCollection.DeleteUser("Brody VonRueden")
	if err != nil {
		fmt.Println("deleteOne error: ", err)
	}

	// findOne demo
	_, err = userCollection.FindUser("Ms. Zetta Runte I")
	if err != nil {
		fmt.Println("findOne error: ", err)
	}

	// Many to many demo
	relationshipCollection := model.NewRelationshipCollection(mongodb)
	follower, _ := userCollection.FindUser("Justyn Koepp")
	following, _ := userCollection.FindUser("Burnice Langworth")
	err = relationshipCollection.FollowUser(follower.ID, following.ID)
	if err != nil {
		fmt.Println("create relationship error: ", err)
	}

	// m2m get
	followers, err := relationshipCollection.GetFollowers(following.ID)
	fmt.Println("followers: ", followers)
}
