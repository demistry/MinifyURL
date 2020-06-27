package main

import (
	h "MinifyURL/api"
	mr "MinifyURL/repository/mongodb"
	"MinifyURL/shortener"
	"os/signal"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"

	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"
) 

func main() {
	err := godotenv.Load("env_variables.env")
	if err != nil{
		fmt.Println("Error loading environment files ", err.Error())
		panic("Could not load env file")
	}
	app := fiber.New()
	repo := createRepo()
	service := shortener.NewRedirectService(repo)
	handler := h.NewHandler(service)
	handler.GET(app)
	handler.POST(app)
 
	errs := make(chan error, 2)

	go func(){
		port := setupPort()
		fmt.Printf("Listening on port %s",port)
		errs <- app.Listen(port)
	}()

	go func(){
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <- c)
	}()

	fmt.Printf("Terminating server due to %s", <- errs)
}

func setupPort() string{
	port := "8080"
	envport,ok := os.LookupEnv("PORT")
	if ok{
		port = envport
	}
	return fmt.Sprintf(":%s",port)
}

func createRepo() shortener.RedirectRepository{
	mongoURL,_ := os.LookupEnv("MONGO_URL")
	mongoDBName,_ := os.LookupEnv("MONGO_DB")
	envTime,_ := os.LookupEnv("MONGO_TIMEOUT")
	mongoTimeout,_ := strconv.Atoi(envTime)
	repo,err := mr.NewMongoRepository(mongoURL, mongoDBName, mongoTimeout + 20)
	if err != nil{
		log.Fatal(err)
	}
	return repo
}