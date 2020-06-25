package main

import (
	h "MinifyURL/api"
	mr "MinifyURL/repository/mongodb"
	"MinifyURL/shortener"
	// "os/signal"

	"github.com/gofiber/fiber"

	"fmt"
	"log"
	"os"
	"strconv"
) 

func main() {
	app := fiber.New()
	repo := createRepo()
	service := shortener.NewRedirectService(repo)
	handler := h.NewHandler(service)
	handler.GET(app)
	handler.POST(app)

	errs := make(chan error, 2)

	go func(){
		port := setupPort()
		fmt.Println("Listening on port %s",port)
		errs <- app.Listen(port)
	}()

	go func(){
		// c := make(chan os.Signal, 1)
		// signal.Notify(c, syscall.SIGINT)
		// errs <- fmt.Errorf("%s", <- c)
	}()

	fmt.Printf("Terminating server due to %s", <- errs)
}

func setupPort() string{
	port := "8080"
	if os.Getenv("PORT") != ""{
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s",port)
}

func createRepo() shortener.RedirectRepository{
	mongoURL := os.Getenv("MONGO_URL")
	mongoDBName := os.Getenv("MONGO_DB")
	mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
	repo,err := mr.NewMongoRepository(mongoURL, mongoDBName, mongoTimeout)
	if err != nil{
		log.Fatal(err)
	}
	return repo
}