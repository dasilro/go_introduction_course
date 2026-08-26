package main

import (
	"encoding/json"
	"fmt"

	// "github.com/ncostamagna/go-http-utils/meta"
	"github.com/ncostamagna/go-http-utils/meta"
	"github.com/ncostamagna/go-http-utils/response"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Mail      string `json:"mail"`
	Age       int    `json:"age"`
}

func main() {
	u := &User{FirstName: "Nahuel", LastName: "Costamagna", Mail: "nlcostamagna@gmail.com", Age: 32}
	print(response.OK("", u, nil))
	print(response.OK("response test", u, nil))

	m, _ := meta.New(1, 30, 100, "15")
	print(response.OK("response test", u, m))

	print(response.Created("", u, nil))
	print(response.Accepted("", u, nil))

	print(response.BadRequest("myError BadRequest"))
	print(response.NotFound("myError NotFound"))
	print(response.InternalServerError("myError InternalServerError"))

}

func print(v interface{}) {
	j, _ := json.Marshal(v)
	fmt.Println(string(j))
	fmt.Println()
}
