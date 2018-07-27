package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"fmt"
	"github.com/gin-contrib/cors"
	"log"
	"net/http"
	"strconv"
	"time"
	"bytes"
	"io/ioutil"
)

// Create private data struct to hold config options.
type config struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	FromAccount     string    `yaml:"from_account"`
}

// Create a new config instance.
var (
	conf *config
)

var router *gin.Engine

// Read the config file from the current directory and marshal
// into the conf config struct.
func getConf() *config {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return conf
}

func main() {

	conf = getConf()

	gin.SetMode(gin.ReleaseMode)

	router = gin.New()
	router.Use(gin.Recovery())

	// allow all origins
	router.Use(cors.Default())

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.POST("/", func(c *gin.Context) {
		publicAddress := c.PostForm("public_address")
		log.Printf("Got address %s\n", publicAddress)

		data := "{\"jsonrpc\":\"2.0\",\"method\":\"personal_sendTransaction\",\"params\":[{\"from\":\""+ conf.FromAccount+"\",\"to\":\"" + publicAddress + "\",\"value\":\"0xde0b6b3a7640000\"}, \"\"],\"id\":0}"
		resp, err := http.Post("http://localhost:8545", "application/json", bytes.NewBuffer([]byte(data)))
		if err != nil {
			log.Panicf("%v", err)
		}

		b, err := ioutil.ReadAll(resp.Body)

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"status": string(b),
		})
	})

	log.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	log.Println("Faucet Running on http://" + conf.Hostname + ":" + strconv.Itoa(conf.Port))
	log.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	// Serve 'em...
	server := &http.Server{
		Addr:           ":" + strconv.Itoa(conf.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.SetKeepAlivesEnabled(false)
	server.ListenAndServe()

}
