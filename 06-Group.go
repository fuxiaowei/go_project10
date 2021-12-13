package main

import (
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	var g errgroup.Group
	g.Go(func() error {
		println("go....")
		time.Sleep(10000000000)
		return nil
	})
	if err := g.Wait(); err == nil {
		println("执行完成.......")
	}

}
