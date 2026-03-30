package main

import "gotest/project/handlers"

//现在目录的入口文件

func main() {
	r := handlers.InitEngine()
	r.Run()
}
