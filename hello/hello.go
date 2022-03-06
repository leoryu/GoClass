package hello

import (
	"fmt"

	restfulv3 "github.com/emicklei/go-restful/v3"
)

func Hello(req *restfulv3.Request, rep *restfulv3.Response) {
	// name := req.QueryParameter("name")
	// if name == "leo" {
	// 	name = "Leo"
	// } else if name == "lora" {
	// 	name = "Lora"
	// } else {
	// 	name = "someone"
	// }

	// switch name {
	// case "leo":
	// 	name = "Leo"
	// case "lora":
	// 	name = "Lora"
	// default:
	// 	name = "someone"

	// }
	// response := fmt.Sprintf("Hello %s", name)

	fmt.Println("this is head")
	defer fmt.Println("this is defer func")
	names := req.QueryParameters("name")

	var response string
	for i, name := range names {
		response += fmt.Sprintf("Hello %d.%s \n", i+1, name)
	}
	fmt.Fprintln(rep.ResponseWriter, response)
	fmt.Println("this is tail")
}
