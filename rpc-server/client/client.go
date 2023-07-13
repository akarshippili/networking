package main

import (
	"fmt"
	"net/rpc"

	"github.com/akarshippili/networking/rpc-server/common"
)

func main() {
	client, error := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if error != nil {
		fmt.Println(error.Error())
	}

	args := common.Args{
		A: 1,
		B: 2000,
	}
	var ans int
	var ans2 float64

	error = client.Call("Math.Add", args, &ans)
	if error != nil {
		fmt.Println(error.Error())
	} else {
		fmt.Printf("ans: %v \n", ans)
	}

	error = client.Call("Math.Substract", args, &ans)
	if error != nil {
		fmt.Println(error.Error())
	} else {
		fmt.Printf("ans: %v \n", ans)
	}

	error = client.Call("Math.Multiply", args, &ans)
	if error != nil {
		fmt.Println(error.Error())
	} else {
		fmt.Printf("ans: %v \n", ans)
	}

	error = client.Call("Math.Divide", args, &ans2)
	if error != nil {
		fmt.Println(error.Error())
	} else {
		fmt.Printf("ans: %v \n", ans2)
	}
}
