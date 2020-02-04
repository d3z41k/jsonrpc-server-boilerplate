package server

import (
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"time"

	"github.com/d3z41k/jsonrpc-server-boilerplate/services"
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

	// // Loop over header names
	// for name, values := range r.Header {
	// 	// Loop over all values for the name.
	// 	for _, value := range values {
	// 		fmt.Println(name, value)
	// 	}
	// }

	ip, _ := u.GetIP(r)
	token := r.Header.Get("X-Auth")

	// Temp hardcode token
	if !u.Contains(acceptIps, ip) || token != os.Getenv("TOKEN_PASSWORD") {
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

// NewRouter return JSON-RPC handler over HTTP that implements the main server routers
func NewRouter() http.Handler {
	router := chi.NewRouter()

	server := rpc.NewServer()
	server.Register(&services.HelloService{})
	server.Register(&services.TradesService{})
	rpcHandler := &Handler{
		rpcServer: server,
	}

	// Set up our middleware with sane defaults
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// Set up root handler
	router.Handle("/rpc", rpcHandler)

	return router
}
