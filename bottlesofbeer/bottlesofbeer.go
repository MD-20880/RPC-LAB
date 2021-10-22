package main

import (
	"flag"
	"fmt"
	"net"
	"net/rpc"
	"os"
	//	"net/rpc"
	//	"net/rpc"
	//	"fmt"
	//	"time"
	//	"net"
)

//
var PassBottle = "Bottle.Pass"

type Request struct {
	Message int
}

type Response struct {
	Message int
}

//

var nextAddr string
var thisPort string

func handlingError(err error) {
	os.Exit(4)
}

type Bottle struct {
	nextInstance string
}

func (b *Bottle) Pass(request Request, response *Response) (err error) {

	//done:=make(chan *rpc.Call)
	number := request.Message
	if number <= 0 {
		return
	}
	fmt.Printf("%d bottles of beer on the wall, %d bottles of beer. Take one down, pass it around...\n", number, number)
	req := Request{Message: number - 1}
	res := new(Response)
	conn, err := rpc.Dial("tcp", nextAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(6)
	}
	response.Message = 1

	conn.Go(PassBottle, req, res, nil)
	conn.Close()
	return
}

func main() {
	flag.StringVar(&nextAddr, "next", "127.0.0.1:8040", "IP:Port string for next member of the round.")
	flag.StringVar(&thisPort, "this", "8030", "Port for this member")
	bottles := flag.Int("n", 0, "Bottles of Beer (launches song if not 0)")
	flag.Parse()
	var conn *rpc.Client

	req := Request{Message: *bottles}
	res := new(Response)

	if *bottles != 0 {
		go func() {
			for {
				dumiconn, err := rpc.Dial("tcp", nextAddr)
				if err == nil {
					conn = dumiconn
					break
				}
			}
			conn.Call(PassBottle, req, res)
			defer conn.Close()
		}()
	}

	//server part
	//TODO: Up to you from here! Remember, you'll need to both listen for
	rpc.Register(&Bottle{nextAddr})
	listener, _ := net.Listen("tcp", ":"+thisPort)
	defer listener.Close()
	rpc.Accept(listener)
}
