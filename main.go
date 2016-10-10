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
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Invalid: %s", err)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	server.InitTemplates()
	getConfig()
	inf := server.InfluxServ{}
	inf.Server = viper.GetString("InfluxServer")
	server.SetServer(inf)
	fmt.Printf("Using influx %s", viper.GetString("InfluxServer"))
	router.HandleFunc("/", server.ServePage)
	http.ListenAndServe(":8080", router)
}
