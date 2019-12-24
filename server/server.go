package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"
	"sync"
	"time"

	u "github.com/d3z41k/jsonrpc-server-boilerplate/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

// NewRouter return HTTP handler that implements the main server routers
func NewRouter() http.Handler {
	router := chi.NewRouter()

	helloManager := NewHelloManager()

	server := rpc.NewServer()
	server.Register(helloManager)

	helloHandler := &Handler{
		rpcServer: server,
	}

	// Set up our middleware with sane defaults
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// Set up root handler
	router.Handle("/rpc", helloHandler)

	return router
}
