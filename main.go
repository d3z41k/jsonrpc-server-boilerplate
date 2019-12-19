package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"os/signal"
	"strconv"
	"sync"

	u "github.com/d3z41k/jsonrpc-server-boilerplate/utils"
	"github.com/joho/godotenv"
)

var acceptIps = []string{"127.0.0.1", "192.168.1.1"}

type HttpConn struct {
	in  io.Reader
	out io.Writer
}

func (c *HttpConn) Read(p []byte) (n int, err error)  { return c.in.Read(p) }
func (c *HttpConn) Write(d []byte) (n int, err error) { return c.out.Write(d) }
func (c *HttpConn) Close() error                      { return nil }

type Handler struct {
	rpcServer *rpc.Server
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("rpc auth: ", r.Header.Get("X-Auth"))

	ip, _ := u.GetIP(r)

	if !u.Contains(acceptIps, ip) {
		w.WriteHeader(403)
		return
	}

	serverCodec := jsonrpc.NewServerCodec(&HttpConn{
		in:  r.Body,
		out: w,
	})
	w.Header().Set("Content-type", "application/json")
	err := h.rpcServer.ServeRequest(serverCodec)
	if err != nil {
		log.Printf("Error while serving JSON request: %v", err)
		http.Error(w, `{"error":"cant serve request"}`, 500)
	} else {
		w.WriteHeader(200)
	}
}

///////////////

type Name struct {
	Name string
	Age  int
}

type HelloManager struct {
	mu sync.RWMutex
}

func NewHelloManager() *HelloManager {
	return &HelloManager{
		mu: sync.RWMutex{},
	}
}
func (hm *HelloManager) Hello(in *Name, out *string) error {
	fmt.Println("call Hello", in)
	// hm.mu.Lock()
	// defer hm.mu.Unlock()

	*out = "Hello " + in.Name + ", your age is " + strconv.Itoa(in.Age) + "."
	return nil
}

//////////////

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rpcPort := os.Getenv("RPC_PORT")

	helloManager := NewHelloManager()

	server := rpc.NewServer()
	server.Register(helloManager)

	helloHandler := &Handler{
		rpcServer: server,
	}
	http.Handle("/rpc", helloHandler)

	fmt.Println("Starting RPC server at :" + rpcPort)
	http.ListenAndServe(":"+rpcPort, nil)

	// Wait for an interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
