package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"notesmanagement/internal/database"

	"notesmanagement/internal/handlers"

	"notesmanagement/internal/services"

	"notesmanagement/internal/repositories"

	"github.com/spf13/viper"

	"github.com/joho/godotenv"
)

type Config struct{
	Port string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main(){
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	viper.AutomaticEnv()

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	fmt.Println("DB_CONN =", config.DBConn)

	//handling confignya
	db, err := database.InitDB(config.DBConn)
	if err != nil{
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	noteRepo := repositories.NewNoteRepo(db)
	noteService := services.NewNoteService(noteRepo)
	noteHandler := handlers.NewNoteHandler(noteService)

	//routing note
	http.HandleFunc("/api/note", noteHandler.HandleNotes)
	http.HandleFunc("/api/note/", noteHandler.HandleNoteByID)

	//tes localhost dengan port dari .env
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		//panggil
		json.NewEncoder(w).Encode(map[string]string{
			"Status": "OK",
			"Message": "API RUNNING",
		})
	})
	fmt.Println("Server Localhost Running di: "+config.Port)

	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil{
		fmt.Println("gagal running server")
	}
}