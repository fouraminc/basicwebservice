package main

import (
	"fmt"
	"net/http"
	"webservice/controllers"
)

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}

func startWebServer(port, numberOfRetries int) (int, error) {

	fmt.Println("Starting Server on port", port)

	fmt.Println("Server started on", port)

	return 3030, nil
}
