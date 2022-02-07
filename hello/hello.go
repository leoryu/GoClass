package hello

import (
	"fmt"

	restfulv3 "github.com/emicklei/go-restful/v3"
)

func Hello(req *restfulv3.Request, rep *restfulv3.Response) {
	fmt.Fprintln(rep.ResponseWriter, "Hello world")
}
