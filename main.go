package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/salkin/weatherServer/server"
	"github.com/spf13/viper"
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
	router := mux.NewRouter().StrictSlash(false)
	server.InitTemplates()
	getConfig()
	inf := server.InfluxServ{}
	inf.Server = viper.GetString("InfluxServer")
	server.SetServer(inf)
	go server.CreateStat()
	fmt.Printf("Using influx %s", viper.GetString("InfluxServer"))
	router.HandleFunc("/", server.ServePage)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(":8080", router)
}
