package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/salkin/weatherServer/server"
	"github.com/spf13/viper"
	"net/http"
)

func getConfig() {
	viper.AddConfigPath("/opt/weatherServer")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Invalid: %s", err)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", server.ServeMain)

	http.ListenAndServe(":8080", router)
}
