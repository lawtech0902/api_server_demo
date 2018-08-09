package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_projects/api_server/router"
	"log"
	"net/http"
	"time"
)

/*
__author__ = 'lawtech'
__date__ = '2018/8/9 下午6:00'
*/

func main() {
	// create the Gin engine
	g := gin.New()

	// gin middlewares
	middlewares := []gin.HandlerFunc{}

	// router
	router.Load(g, middlewares...)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 2; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
