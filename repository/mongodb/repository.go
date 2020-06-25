package mongodb

import (
	// jsonserializer "MinifyURL/serializer/json"
	"MinifyURL/shortener"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	// "io/ioutil"
	// "log"
	// "os"
)

const redirectCollectionName = "redirects"

type mongoDBRepository struct{
	client *mongo.Client
	database string
	timeout time.Duration
}
type mongoClient struct{
	fileName string
	filePath string
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client,error){
	ctx,cancel := context.WithTimeout(context.Background(),time.Duration(mongoTimeout))
	defer cancel()
	client,err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil{
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil{
		return nil, err
	}
	return client,nil
}

func NewMongoRepository(mongoURL, mongoDB string, timeout int) (shortener.RedirectRepository,error){
	repo := &mongoDBRepository{
		timeout: time.Duration(timeout) * time.Second,
		database: mongoDB,
	} 
	client,err := newMongoClient(mongoURL,timeout)
	if err != nil{
		return nil, err
	}
	repo.client = client
	return repo,nil
}

func (r *mongoDBRepository) Find(code string) (*shortener.Redirect, error){
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	redirect := &shortener.Redirect{}
	collection := r.client.Database(r.database).Collection(redirectCollectionName)
	filter := bson.M{"code":code}
	err := collection.FindOne(ctx,filter).Decode(&redirect)
	if err != nil{
		return nil, err
	}
	return redirect,nil
}


func (r *mongoDBRepository) Store(redirect *shortener.Redirect) error{
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(redirectCollectionName)
	_, err := collection.InsertOne(ctx, bson.M{
		"code" : redirect.Code,
		"url" : redirect.URL,
		"createdAt" : redirect.CreatedAt,
	})
	if err != nil{
		return err
	}
	return nil
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}