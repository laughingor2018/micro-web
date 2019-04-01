package subscriber

import (
	"context"

	"github.com/micro/go-log"

	example "github.com/laughingor2018/micro-web/fnc/proto/example"
)

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *example.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}
