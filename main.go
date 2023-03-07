package main

import "Go_ChatGPT/router"

func main() {

	route := router.InitRouter()
	route.Run(":12233")
}
