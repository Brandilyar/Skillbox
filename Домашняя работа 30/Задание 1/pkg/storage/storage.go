package storage

import (
	"context"
	"fmt"
	"http_service/pkg/user"
	"log"
	"sync"
	"sync/atomic"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserList struct {
	sync.Mutex
	ctx     context.Context
	Storage *mongo.Collection `json:"storage"`
	Maxid   int64             `json:"maxid"`
}
type FriendsList struct {
	Friends []*user.Friend `json:"friends"`
}

func (ul *UserList) AppendUser(u *user.User) error {
	ul.Lock()
	res, err := ul.Storage.InsertOne(ul.ctx, bson.D{primitive.E{Key: "id", Value: ul.Maxid}, primitive.E{Key: "name", Value: u.Name}, primitive.E{Key: "age", Value: u.Age}, primitive.E{Key: "friends", Value: u.Friends}})
	atomic.AddInt64(&ul.Maxid, 1)
	if err != nil {
		return err
	}
	fmt.Printf("inserted document with ID %v\n", res.InsertedID)
	ul.Unlock()
	return nil
}
func (ul *UserList) UpdateUser(u *user.User) error {
	ul.Lock()
	var updatedDocument bson.M
	filter := bson.D{{"id", u.Id}}
	update := bson.D{{"$set", bson.D{{"name", u.Name}, {"age", u.Age}, {"friends", u.Friends}}}}
	err := ul.Storage.FindOneAndUpdate(ul.ctx, filter, update).Decode(&updatedDocument)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			return err
		}
		log.Fatal(err)
	}
	fmt.Printf("updated document %v", updatedDocument)
	ul.Unlock()
	return nil
}
func (ul *UserList) DelUser(id int64) error {
	ul.Lock()
	cur, err := ul.Storage.Find(ul.ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		user := user.User{}
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		for i, one_user_id := range user.Friends {
			if one_user_id == id {
				user.Friends = append(user.Friends[:i], user.Friends[i+1:]...)
				var updatedDocument bson.M
				filter := bson.D{{"id", user.Id}}
				update := bson.D{{"$set", bson.D{{"friends", user.Friends}}}}
				err := ul.Storage.FindOneAndUpdate(ul.ctx, filter, update).Decode(&updatedDocument)
				if err != nil {
					// ErrNoDocuments means that the filter did not match any documents in
					// the collection.
					if err == mongo.ErrNoDocuments {
						return err
					}
					log.Fatal(err)
				}
				fmt.Printf("updated document %v", updatedDocument)

			}
		}
	}
	res, err := ul.Storage.DeleteOne(ul.ctx, bson.D{{"id", id}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	ul.Unlock()
	return nil
}
func (ul *UserList) GetUserFriend(id int64) ([]*user.Friend, error) {
	u := user.User{}
	friends := FriendsList{}
	err := ul.Storage.FindOne(ul.ctx, bson.D{{"id", id}}).Decode(&u)
	for _, friendid := range u.Friends {
		one_friend := user.Friend{}
		err = ul.Storage.FindOne(ul.ctx, bson.D{{"id", friendid}}).Decode(&one_friend)
		if err != nil {
			return nil, err
		}
		friends.Friends = append(friends.Friends, &one_friend)
	}
	if err != nil {
		return nil, err
	}
	return friends.Friends, nil
}
func (ul *UserList) GetUser(id int64) (*user.User, error) {
	user := user.User{}
	err := ul.Storage.FindOne(ul.ctx, bson.D{{"id", id}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func NewUserList(ctx context.Context, coll *mongo.Collection) *UserList {
	return &UserList{ctx: ctx, Storage: coll, Maxid: 1}
}
