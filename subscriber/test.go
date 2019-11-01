package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	test "test/proto/test"
)

type Test struct{}

func (e *Test) Handle(ctx context.Context, msg *test.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *test.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
