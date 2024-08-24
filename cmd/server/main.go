package main

import (
	"github.com/hungntth/be-go-ecom/internal/routers"
)

func main() {
	r := routers.NewRouter()

	r.Run(":8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

