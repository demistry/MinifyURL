package mongodb

import (
	// jsonserializer "MinifyURL/serializer/json"
	"MinifyURL/shortener"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	// "io/ioutil"
	// "log"
	// "os"
)


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
		timeout: time.Duration(timeout),
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
	// byteOfData, err := ioutil.ReadFile(r.fileName)
	// if err != nil{
	// 	return nil,err
	// }
	
	return nil,nil
}


func (r *mongoDBRepository) Store(redirect *shortener.Redirect) error{
	// file,err := os.Create(r.filePath)
	// if err != nil{
	// 	log.Fatal("Failed to create file in repository")
	// 	return err
	// }
	// enc, err := r.serializer.Encode(redirect)
	// if err != nil{
	// 	return err
	// }
	// file.Write(enc)
	// defer file.Close()
	return nil
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}