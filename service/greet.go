package service

import (
	"github.com/kataras/iris/v12"
	"github.com/renjingneng/a_simple_go_project/core/config"
)

// Greet example service.
type Greet interface {
	Say(input string) (string, error)
}

// NewGreetService is
func NewGreetService(iris iris.Context) Greet {
	switch config.Config["Env"] {
	case "prod":
		return &greeter{"Hello"}
	case "dev":
		return &greeterWithLogging{"greeterWithLogging"}
	default:
		return &greeter{"Hello"}
	}
}

type greeter struct {
	prefix string
}

func (s *greeter) Say(input string) (string, error) {
	result := s.prefix + " " + input
	return result, nil
}

type greeterWithLogging struct {
	prefix string
}

func (s *greeterWithLogging) Say(input string) (string, error) {
	result := s.prefix + " " + input
	return result, nil
}
