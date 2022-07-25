package main

import "finalProject/router"

func main() {
	r := router.StartApp()

	r.Run(":4001")
}
