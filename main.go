package main

import (
	"log"
	"net/http"
	"trueconf-userapi/config"
	"trueconf-userapi/routers"
)

func main() {
	r := routers.ApiRouter()

	err := http.ListenAndServe(":"+config.PORT, r)
	if err != nil {
		log.Fatal(err)
	}
}
