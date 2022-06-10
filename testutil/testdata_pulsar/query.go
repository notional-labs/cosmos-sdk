//nolint:revive,gocritic // later we could resolve this by changing how we generate things, but was not sure how this file is generated and if that woud pose a breaking change.
package testdata_pulsar

import (
	"context"
	"fmt"
)

type QueryImpl struct {
	UnimplementedQueryServer
}

func (q QueryImpl) Echo(_ context.Context, request *EchoRequest) (*EchoResponse, error) {
	return &EchoResponse{Message: request.Message}, nil
}

func (q QueryImpl) SayHello(_ context.Context, request *SayHelloRequest) (*SayHelloResponse, error) {
	greeting := fmt.Sprintf("Hello %s!", request.Name)
	return &SayHelloResponse{Greeting: greeting}, nil
}

func (q QueryImpl) TestAny(_ context.Context, request *TestAnyRequest) (*TestAnyResponse, error) {
	return &TestAnyResponse{HasAnimal: &HasAnimal{
		Animal: request.AnyAnimal,
		X:      10,
	}}, nil
}

var _ QueryServer = QueryImpl{}
