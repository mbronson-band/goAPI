package main

import (
	"encoding/json"
	"fmt"
	"log"

	"database/sql"
	"net/http"

	_ "github.com/snowflakedb/gosnowflake"

	"github.com/gorilla/mux"
	"os"
	`"github.com/spf13/viper"`

	crud "goAPI/crud"
)

type DbString struct{
	DBConfig string 'mapstructure:"DBConfig"'
	DBSource string 'mapstructure:"DBSource"'
} 

func allArticles(w http.ResponseWriter, r *http.Request) {
	

}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	
}

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/", allArticles).Methods("GET")
	http.HandleFunc("/", testPostArticles).Methods("POST")

	log.Fatal(http.ListenAndServe(":8082", myRouter))

}


func LoadConfig(path string) (DbString, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
    viper.SetConfigType("env")

	viper.AutomaticEnv()

    err = viper.ReadInConfig()
    if err != nil {
        return
    }

    err = viper.Unmarshal(&config)
    return
}

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
        log.Fatal("cannot load config:", err)
    }

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("error opening DB (%s)", err)
	}

	handleRequests()
}
