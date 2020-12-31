package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/skyaxl/synack/pkg/maze"
)

func main() {
	errs := make(chan error, 2)
	handler := maze.HttpWarmUp()
	go func() {
		fmt.Printf("[Maze Api] Has started with address localhost:8080\n")
		errs <- http.ListenAndServe(":8080", handler)
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("[Maze Api] Has stoped. \n%s\n", <-errs)
}
