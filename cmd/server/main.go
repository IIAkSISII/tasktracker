package main

import (
	"fmt"
	"github.com/IIAkSISII/tasktracker/internal/config"
	"github.com/IIAkSISII/tasktracker/internal/database"
	"github.com/IIAkSISII/tasktracker/internal/logger"
	boardRepository "github.com/IIAkSISII/tasktracker/internal/repository/board"
	projectRepository "github.com/IIAkSISII/tasktracker/internal/repository/project"
	taskRepository "github.com/IIAkSISII/tasktracker/internal/repository/task"
	ticketRepository "github.com/IIAkSISII/tasktracker/internal/repository/ticket"
	userReposity "github.com/IIAkSISII/tasktracker/internal/repository/user"
	boardSevice "github.com/IIAkSISII/tasktracker/internal/service/board"
	projectService "github.com/IIAkSISII/tasktracker/internal/service/project"
	taskService "github.com/IIAkSISII/tasktracker/internal/service/task"
	ticketService "github.com/IIAkSISII/tasktracker/internal/service/ticket"
	userService "github.com/IIAkSISII/tasktracker/internal/service/user"
	"github.com/IIAkSISII/tasktracker/internal/transport/http/boardHandler"
	"github.com/IIAkSISII/tasktracker/internal/transport/http/projectHandler"
	"github.com/IIAkSISII/tasktracker/internal/transport/http/taskHandler"
	"github.com/IIAkSISII/tasktracker/internal/transport/http/ticketHandler"
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

	logger, err := logger.NewLogrusLogger(cfg.Logger.Level, cfg.Logger.Format)
	if err != nil {
		fmt.Println("Error initializing logger:", err)
	}

	uRepo := userReposity.NewUserRepository(db, logger)
	uSvc := userService.NewUserService(uRepo, logger)
	uHandler := userHandler.NewUserHandler(uSvc, logger)

	pRepo := projectRepository.NewProjectRepository(db, logger)
	pSvc := projectService.NewProjectService(pRepo, logger)
	pHandler := projectHandler.NewProjectHandler(pSvc, logger)

	bRepo := boardRepository.NewBoardRepository(db, logger)
	bSvc := boardSevice.NewBoardService(bRepo, logger)
	bHandler := boardHandler.NewBoardHandler(bSvc, logger)

	tRepo := ticketRepository.NewTicketRepository(db, logger)
	tSvc := ticketService.NewTicketService(tRepo, logger)
	tHandler := ticketHandler.NewTicketHandler(tSvc, logger)

	taskRepo := taskRepository.NewTaskRepository(db, logger)
	taskSvc := taskService.NewTicketService(taskRepo, logger)
	taskHandler := taskHandler.NewTaskHandler(taskSvc, logger)

	router := mux.NewRouter()
	router.Use(middlewares.CorsMiddleware)
	router.Use(middlewares.LoggerMiddleware(logger))

	uHandler.ConfigureRoutes(router)
	pHandler.ConfigureRoutes(router)
	bHandler.ConfigureRoutes(router)
	tHandler.ConfigureRoutes(router)
	taskHandler.ConfigureRoutes(router)

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
