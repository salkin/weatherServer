package server

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/spf13/viper"
)

//CreateStat creates the influx snapshot every hour
func CreateStat(dir string, authToken string) {
	grafanaUser := viper.GetString("GrafanaUser")
	grafanaPassword := viper.GetString("GrafanaPassword")
	grafanaServer := viper.GetString("GrafanaServer")
	for {
		client := http.Client{}

		httpReq, err := http.NewRequest("GET", "http://"+grafanaUser+":"+grafanaPassword+grafanaServer+"/render/dashboard/db/weather", nil)

		httpReq.Header.Add("Authorization", "Bearer "+authToken)
		res, err := client.Do(httpReq)
		if err != nil {

		}
		file, err := os.Create(dir + "/static/daily.png")
		defer file.Close()
		io.Copy(file, res.Body)

		time.Sleep(time.Hour)
	}

}
