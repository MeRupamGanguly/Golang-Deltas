package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Status string

const (
	Pending  Status = "Pending"
	Approved Status = "Approved"
	Rejected Status = "Rejected"
)

type Artist struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}
type Release struct {
	PlatformId string   `json:"platform_id" bson:"platform_id"`
	Status     Status   `json:"status" bson:"status"`
	Artists    []Artist `json:"artists" bson:"artists"`
}
type Album struct {
	Id       string    `json:"id" bson:"_id"`
	Name     string    `json:"name" bson:"name"`
	LabelId  string    `json:"label_id" bson:"label_id"`
	Releases []Release `json:"releases" bson:"releases"`
}
type Song struct {
	Id      string `json:"id" bson:"_id"`
	Name    string `json:"name" bson:"name"`
	AlbumId string `json:"album_id" bson:"album_id"`
}

type Label struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

func main() {
	uri := "mongodb://localhost:27017"
	conn, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
		panic("err")
	}
	albumCol := conn.Database("MyDatabase").Collection("Album")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	rs := []Release{
		{PlatformId: bson.NewObjectID().Hex(), Status: Pending, Artists: []Artist{{Id: bson.NewObjectID().Hex(), Name: "Artist1"}, {Id: bson.NewObjectID().Hex(), Name: "Artist3"}}},
		{PlatformId: bson.NewObjectID().Hex(), Status: Approved, Artists: []Artist{{Id: bson.NewObjectID().Hex(), Name: "Artist2"}}},
	}
	l := Label{
		Id:   bson.NewObjectID().Hex(),
		Name: "L1",
	}
	a := Album{
		Id:       bson.NewObjectID().Hex(),
		Name:     "A1",
		LabelId:  l.Id,
		Releases: rs,
	}
	fmt.Println(a)
	rs2 := []Release{
		{PlatformId: bson.NewObjectID().Hex(), Status: Pending, Artists: []Artist{{Id: bson.NewObjectID().Hex(), Name: "Artist4"}, {Id: bson.NewObjectID().Hex(), Name: "Artist5"}}},
		{PlatformId: bson.NewObjectID().Hex(), Status: Approved, Artists: []Artist{{Id: bson.NewObjectID().Hex(), Name: "Artist3"}}},
	}
	l2 := Label{
		Id:   bson.NewObjectID().Hex(),
		Name: "L2",
	}
	a2 := Album{
		Id:       bson.NewObjectID().Hex(),
		Name:     "A2",
		LabelId:  l2.Id,
		Releases: rs2,
	}
	s1 := Song{
		Id:      bson.NewObjectID().Hex(),
		Name:    "A1 Song 1",
		AlbumId: a.Id,
	}
	s2 := Song{
		Id:      bson.NewObjectID().Hex(),
		Name:    "A1 Song 2",
		AlbumId: a.Id,
	}
	s3 := Song{
		Id:      bson.NewObjectID().Hex(),
		Name:    "A2 Song 1",
		AlbumId: a2.Id,
	}
	oid, err := albumCol.InsertOne(ctx, a)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(oid)
	oid, err = albumCol.InsertOne(ctx, a2)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(oid)
	songCol := conn.Database("MyDatabase").Collection("Song")
	oid, err = songCol.InsertOne(ctx, s1)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(oid)
	oid, err = songCol.InsertOne(ctx, s2)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(oid)
	oid, err = songCol.InsertOne(ctx, s3)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(oid)
}
