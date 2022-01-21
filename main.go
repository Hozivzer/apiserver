package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/apiserver/config"
	"github.com/apiserver/model"
	"github.com/apiserver/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {

	pflag.Parse()

	//init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	//init db
	model.DB.Init()
	defer model.DB.Close()

	//set gin mode
	gin.SetMode(viper.GetString("runmode"))

	//create gin enine
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	//Routes
	router.Load(
		g,

		middlewares...,
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Print(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")

		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the route r")
}
