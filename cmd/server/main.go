package main

import (
	"fmt"
	"github.com/IIAkSISII/tasktracker/internal/config"
	"github.com/IIAkSISII/tasktracker/internal/database"
	userReposity "github.com/IIAkSISII/tasktracker/internal/repository/user"
	userService "github.com/IIAkSISII/tasktracker/internal/service/user"
	"github.com/IIAkSISII/tasktracker/internal/transport/http/userHandler"
	"github.com/IIAkSISII/tasktracker/internal/transport/middlewares"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title			Трекер задач
// @version		1.0
// @description	Разработка бекенда для отслеживания задач
// @contact.name	Антон
// @contact.url	https://github.com/IIAkSISII
// @contact.email	iiaksisii@gmail.com
// @host			localhost:8080
//
// @accept			json
// @produce		json text/plain
// @schemes		http https
func main() {
	cfg, err := config.NewConfig(config.GetConfigPath())
	if err != nil {
		fmt.Println("Error loading config:", err)
	}

	db, err := database.Connect(cfg.Database)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}
	defer db.Close()

	uRepo := userReposity.NewUserRepository(db)
	uSvc := userService.NewUserService(uRepo)
	uHandler := userHandler.NewUserHandler(uSvc)

	router := mux.NewRouter()
	router.Use(middlewares.CorsMiddleware)

	uHandler.ConfigureRoutes(router)

	router.HandleFunc("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	}).Methods(http.MethodGet)

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler())

	server := &http.Server{
		Addr:    cfg.Server.Host + ":" + cfg.Server.Port,
		Handler: router,
	}
	log.Printf("Starting server on %s", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
