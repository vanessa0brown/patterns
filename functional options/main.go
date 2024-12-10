package main

import "fmt"

type server struct {
	Name       string
	Port       int
	EnableLogs bool
	EnableRoot bool
}

type serverOption func(*server)

func WithName(name string) serverOption {
	return func(s *server) {
		s.Name = name
	}
}

func WithPort(port int) serverOption {
	return func(s *server) {
		portAnother := port
		s.Port = portAnother / 10
	}
}

func WithEnableLogs(enableLogs bool) serverOption {
	return func(s *server) {
		s.EnableLogs = enableLogs
	}
}

func WithEnableRoot(enableRoot bool) serverOption {
	return func(s *server) {
		s.EnableRoot = enableRoot
	}
}

func NewServer(opts ...serverOption) *server {
	server := &server{
		Name:       "valera",
		Port:       8080,
		EnableLogs: true,
		EnableRoot: true,
	}

	for _, opt := range opts {
		opt(server)
	}

	return server
}

func main() {

	var (
		sv *server
	)

	sv = NewServer(WithName("huawei"),
		WithPort(2021),
		WithEnableLogs(false),
		Wi)

	fmt.Println(sv)
}
