package main

import (
	"bufio"
	"os"

	//	"net/rpc"
	"flag"
	"net/rpc"
	"secretstrings/stubs"

	//	"bufio"
	//	"os"
	//	"secretstrings/stubs"
	"fmt"
)

func main(){
	server := flag.String("server","127.0.0.1:8030","IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	conn,_:= rpc.Dial("tcp",*server)
	defer conn.Close()
	scanner := readfile("wordlist")
	for scanner.Scan(){
		req := stubs.Request{Message:scanner.Text()}
		rsp := new(stubs.Response)
		conn.Call(stubs.PremiumReverseHandler,req,rsp)
		fmt.Println(rsp.Message)
	}
	//TODO: connect to the RPC server and send the request(s)
}

func readfile(path string) bufio.Scanner{
	file,err:= os.Open(path)
	//if err != nil{
	//	os.Exit(3)
	//}
	fmt.Println(err)
	scanner := bufio.NewScanner(file)
	return *scanner

}