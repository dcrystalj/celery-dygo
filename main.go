package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	celery "github.com/marselester/gopher-celery"
	credis "github.com/marselester/gopher-celery/redis"
)

func main() {
	broker := credis.NewBroker(
		credis.WithPool(&redis.Pool{
			Dial: func() (redis.Conn, error) {
				return redis.DialURL("redis://localhost:6379/0")
			},
		}),
	)
	app := celery.NewApp(celery.WithBroker(broker))
	app.Register(
		"myproject.apps.myapp.tasks.mytask",
		"celery",
		func(ctx context.Context, p *celery.TaskParam) error {
			p.NameArgs("a", "b")
			// Methods prefixed with Must panic if they can't find an argument name
			// or can't cast it to the corresponding type.
			// The panic doesn't affect other tasks execution; it's logged.
			fmt.Println(p.MustInt("a") + p.MustInt("b"))
			// Non-nil errors are logged.

			app.Register(
				"myproject.apps.myapp.tasks.mytask5",
				"celery",
				func(ctx context.Context, p *celery.TaskParam) error {
					p.NameArgs("a", "b")
					// Methods prefixed with Must panic if they can't find an argument name
					// or can't cast it to the corresponding type.
					// The panic doesn't affect other tasks execution; it's logged.
					fmt.Println(p.MustInt("a") + p.MustInt("b"))
					// Non-nil errors are logged.
					return nil
				},
			)
			return nil
		},
	)
	if err := app.Run(context.Background()); err != nil {
		log.Printf("celery worker error: %v", err)
	}
}
