package RPC_Client

import (
	"Go_RPC/Common"
	"fmt"
	"log"
	"net/rpc"
)

func Client(address string, port string) {
	rpcClient, e := rpc.DialHTTP("tcp", address+":"+port)
	if e != nil{
		log.Fatal("error connecting to RPC server: ",e)
	}
	args:= Common.Args{StringOne:"Hello, ",StringTwo:" World!"}
	var resultString string
	e = rpcClient.Call("StringOps.Concat", args, &resultString)
	if e != nil {
		log.Fatal("RPC error:", e)
	}

	fmt.Println("Result string is: ",resultString)
}
