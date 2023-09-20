package main

import "fmt"

type Server struct {
	options ServerOptions
}

type ServerOptions struct {
	maximumConnections int64
	id string
	tls bool
}

func defaultServerOptions() ServerOptions {
	return ServerOptions{
		maximumConnections: 10,
		id: "default",
		tls: false,
	}
}

type OptionFn = func(*ServerOptions);

func withTLS(serverOptions *ServerOptions) {
	serverOptions.tls = true;
}

func withMaximumConnections(n int64) OptionFn{
	return func (serverOptions *ServerOptions) {
		serverOptions.maximumConnections = n;
	}
}

func withID(id string) OptionFn {
	return func (serverOptions *ServerOptions) {
		serverOptions.id = id;
	}
}

func newServer(options ...OptionFn) Server {
	serverOptions := defaultServerOptions()
	for _, option := range options {
		option(&serverOptions)
	}
	return Server{
		options: serverOptions,
	}
}
func main(){
	server := newServer(withID("interview me"), withMaximumConnections(42), withTLS);
	fmt.Printf("+%v", server);
}