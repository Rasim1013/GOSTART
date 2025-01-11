package main

import (
	"fmt"
	"log"
	"my_crud/api"
	"my_crud/db"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	router := mux.NewRouter()

	// Загружаем .env файл
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	// Получаем параметры подключения
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbConfing := db.ConfigDB{
		IP:       dbHost,
		Port:     dbPort,
		DbName:   dbName,
		UserName: dbUser,
		Password: dbPassword,
	}

	mysqlConn, err := db.ConnectMySql(dbConfing)
	if err != nil {
		fmt.Printf("error: %w", err)
	}

	defer func() {
		err := mysqlConn.Close()
		if err != nil {
			fmt.Printf("Can't close pg connection: %s", err)
		}
	}()

	var (
		userMySqlStore = db.NewMysqlUserSotre(mysqlConn)
		userHandler    = api.NewUserHandler(userMySqlStore)
	)

	router.HandleFunc("/", api.TestHandler).Methods("GET")
	router.HandleFunc("/sms_one", api.SmsOneHandler).Methods("GET")
	router.HandleFunc("/sms_bulk", api.SmsBulkHandler).Methods("GET")
	router.HandleFunc("/users", userHandler.GetUsersHandler).Methods("GET")
	router.HandleFunc("/user/{id:[0-9]+}", userHandler.GetUserDetailHandler).Methods("GET")
	router.HandleFunc("/user_create", userHandler.CreateUsersHandler).Methods("POST")
	router.HandleFunc("/user_update/{id:[0-9]+}", userHandler.UpdateUsersHandler).Methods("PUT")
	router.HandleFunc("/user_delete/{id:[0-9]+}", userHandler.DeleteUsersHandler).Methods("DELETE")

	http.ListenAndServe(":1013", router)

}
