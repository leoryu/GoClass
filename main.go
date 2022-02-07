package main

import (
	"net/http"

	restfulv3 "github.com/emicklei/go-restful/v3"
	"github.com/leoryu/GoLesson/hello"
)

func main() {
	ws := new(restfulv3.WebService)
	ws.Route(ws.GET("/").To(hello.Hello))
	restfulv3.Add(ws)

	http.ListenAndServe(":8080", nil)
}
