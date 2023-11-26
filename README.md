# go-just-options
A simple library for creating options setters in modules within Go projects.

## Requirements
 * Go 1.21

## Usage
```go
// main.go
package main

import (
	"log"
	
	// ...
)

func main() {
	// Create new service with NewService and provide configuration.
	s, err := service.NewService(service.OptionToken("__token"), service.OptionPort(443))
	if err != nil {
		log.Fatalln(err)
	}
	
	// ...
}
```
```go
// service/service.go
package service

import (
	"log"

	options "github.com/SevenOfSpades/go-just-options"
	// ...
)

const (
	optionServiceToken options.OptionKey = `service_token`
	optionServicePort  options.OptionKey = `service_port`
)

type (
	Port  int
	Token string

	Service struct {
		token Token
		port  Port
	}
)

// NewService only needs to accept options.Option from anything that instantiates it.
func NewService(opts ...options.Option) (*Service, error) {
	resolved := options.Resolve(opts)

	token, err := options.Read[Token](resolved, optionServiceToken)
	// or use token := options.ReadOrPanic[Token](opt, optionServiceToken) to panic when there is an error during read
	if err != nil {
		return nil, err
	}

	port, err := options.ReadOrDefault[Port](resolved, optionServicePort, Port(80))
	// or use port, err := options.ReadOrDefaultOrPanic[Port](opt, optionServicePort, Port(80)) to panic when there is an error during read
	if err != nil {
		return nil, err
	}

	return &Service{token: token, port: port}, nil
}

// OptionPort will assign value for Port in options.Resolver. There is no need for specifying custom type. This can be an integer.
// Options can accept multiple values of the same type as long as provided key is unique.
func OptionPort(port Port) options.Option {
	return func(r options.Resolver) {
		if err := options.Write[Port](r, optionServicePort, port); err != nil {
			// detect an error from write operation and do something with it
			log.Fatalln(err)
		}
	}
}

// OptionToken will assign value for Token in options.Resolver. There is no need for specifying custom type. This can be a string.
// Options can accept multiple values of the same type as long as provided key is unique.
func OptionToken(token Token) options.Option {
	return func(r options.Resolver) {
		// it will panic when there is an error during write operation
		options.WriteOrPanic[Token](r, optionServiceToken, token)
	}
}
```