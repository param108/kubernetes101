package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"
)

var (
	DB     *gorm.DB
	server *http.Server
	done   chan int
	Version string
)

func main() {
	godotenv.Load()
	server = &http.Server{
		Addr: "0.0.0.0:8080",
	}

	done = make(chan int)
	handler := mux.NewRouter()
	handler.HandleFunc("/ping", pingHandler)
	handler.HandleFunc("/write", writeHandler).Methods("POST")
	handler.HandleFunc("/read", readHandler)
	handler.HandleFunc("/quit", quitHandler)
	server.Handler = handler
	setupDB()


	server.ListenAndServe()
	<-done
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("<html>ok-%s</html>", Version)))
}

func quitHandler(w http.ResponseWriter, r *http.Request) {
	go func() {
		time.Sleep(5000)
		server.Shutdown(context.Background())
	}()
	w.Write([]byte("<html>ok</html>"))
	close(done)

}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	line := &Line{}
	err := json.NewDecoder(r.Body).Decode(line)
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, err.Error(), 400)
		return
	}
	DB.Create(line)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\": \"true\"}"))
}

type ReadResponse struct {
	Success string `json:"success"`
	Lines   []Line `json:lines`
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	lines := []Line{}
	response := ReadResponse{}
	DB.Find(&lines)
	w.WriteHeader(http.StatusOK)
	response.Success = "true"
	response.Lines = lines
	b, _ := json.Marshal(response)
	w.Write(b)
}

func setupDB() {
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbUri)

	for {
		conn, err := gorm.Open("postgres", dbUri)
		if err != nil {
			fmt.Println(err)
		} else {
			conn.AutoMigrate(&Line{})
			DB = conn
			break
		}
		time.Sleep(2000)
	}

}
