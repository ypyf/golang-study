package main

import (
	//"bytes"
	"fmt"
	"io"
	//"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	//"testing"

	"github.com/gorilla/mux"
)

// adapt HTTP connection to ReadWriteCloser
type HttpConn struct {
	in  io.Reader
	out io.Writer
}

func (c *HttpConn) Read(p []byte) (n int, err error)  { return c.in.Read(p) }
func (c *HttpConn) Write(d []byte) (n int, err error) { return c.out.Write(d) }
func (c *HttpConn) Close() error                      { return nil }

// our service
type Monitor struct{}

func (m *Monitor) StartwlanOk(n int, msg *string) error {
	*msg = fmt.Sprintf("your cake has been bacon (%d)", n)
	return nil
}

func InitJsonRpcServer(rcvr interface{}) {
	server := rpc.NewServer()
	server.Register(rcvr)

	listener, e := net.Listen("tcp", ":7498")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	defer listener.Close()

	RpcHandler := func(w http.ResponseWriter, r *http.Request) {
		serverCodec := jsonrpc.NewServerCodec(&HttpConn{in: r.Body, out: w})
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		err := server.ServeRequest(serverCodec)
		if err != nil {
			log.Printf("Error while serving JSON request: %v", err)
			http.Error(w, "Error while serving JSON request, details have been logged.", 500)
			return
		}
	}

	r := mux.NewRouter()
	r.HandleFunc("/rpc", RpcHandler)
	go http.Serve(listener, r)
}

func main() {
	server := rpc.NewServer()
	server.Register(&Monitor{})

	listener, e := net.Listen("tcp", ":7498")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	defer listener.Close()

	RpcHandler := func(w http.ResponseWriter, r *http.Request) {
		serverCodec := jsonrpc.NewServerCodec(&HttpConn{in: r.Body, out: w})
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		err := server.ServeRequest(serverCodec)
		if err != nil {
			log.Printf("Error while serving JSON request: %v", err)
			http.Error(w, "Error while serving JSON request, details have been logged.", 500)
			return
		}
	}

	r := mux.NewRouter()
	r.HandleFunc("/rpc", RpcHandler)
	http.Serve(listener, r)

	ch := make(chan int)
	<-ch
	/*
	    // Client Code
	    resp, err := http.Post("http://localhost:4321/bake-me-a-cake", "application/json", bytes.NewBufferString(
	       `{"jsonrpc":"2.0","id":1,"method":"CakeBaker.BakeIt","params":[10]}`,
	   ))
	   if err != nil {
	       panic(err)
	   }
	   defer resp.Body.Close()
	   b, err := ioutil.ReadAll(resp.Body)
	   if err != nil {
	       panic(err)
	   }

	   fmt.Printf("returned JSON: %s\n", string(b))
	*/
}
