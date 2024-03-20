package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"stakeholders.xws.com/handler"
	"stakeholders.xws.com/model"
	"stakeholders.xws.com/repo"
	"stakeholders.xws.com/service"
)

func initDB() *gorm.DB {

	dsn := "user=postgres password=super dbname=stakeholders host=localhost port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Person{})
	database.AutoMigrate(&model.User{})
	return database
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	// Repositories
	personRepo := &repo.PersonRepository{DatabaseConnection: database}
	userRepo := &repo.UserRepository{DatabaseConnection: database}

	// Services
	personService := &service.PersonService{PersonRepo: personRepo}
	accountService := &service.AccountService{
		PersonRepo: personRepo,
		UserRepo:   userRepo,
	}

	//Handlers
	personHandler := &handler.PersonHandler{PersonService: personService}
	accountHandler := &handler.AccountHandler{AccountService: accountService}

	router := mux.NewRouter().StrictSlash(true)

	// Profile router
	router.HandleFunc("/profile/{userId}", personHandler.Get).Methods("GET")
	router.HandleFunc("/profile/{id}", personHandler.Update).Methods("PUT")

	// Account Router
	router.HandleFunc("/accounts", accountHandler.GetAll).Methods("GET")
	router.HandleFunc("/accounts", accountHandler.BlockOrUnblock).Methods("PUT")

	permittedHeaders := handlers.AllowedHeaders([]string{"Requested-With", "Content-Type", "Authorization"})
	permittedOrigins := handlers.AllowedOrigins([]string{"*"})
	permittedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8083", handlers.CORS(permittedHeaders, permittedOrigins, permittedMethods)(router)))
}
