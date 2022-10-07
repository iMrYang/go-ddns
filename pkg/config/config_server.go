package config

import "flag"

var _ Config = (*Server)(nil)

type Server struct {
	// Host is the host to bind to.
	Host string `mapstructure:"host" validate:""`
	// Port is the port to bind to.
	Port int `mapstructure:"port"  validate:"gte=0,lte=65535"`
}

func NewServer() *Server {
	return &Server{
		Host: "localhost",
		Port: 10086,
	}
}

func (s *Server) FlagSet() *flag.FlagSet {
	fs := flag.NewFlagSet("server", flag.ExitOnError)
	fs.StringVar(&s.Host, "server.host", s.Host, "server host")
	fs.IntVar(&s.Port, "server.port", s.Port, "server port")
	return fs
}
