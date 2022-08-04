package Async

import (
	"context"
	ApplicationCommonInterfaces "golangcodebase/Application/Common/Interfaces"
)

type async struct {
	await func(ctx context.Context) interface{}
}

func (f async) Await() interface{} {
	return f.await(context.Background())
}

func Execute(function func() interface{}) ApplicationCommonInterfaces.IAsync {
	var result interface{}
	c := make(chan struct{})
	go func() {
		defer close(c)
		result = function()
	}()
	return async{
		await: func(ctx context.Context) interface{} {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				return result
			}
		},
	}
}
