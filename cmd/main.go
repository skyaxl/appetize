package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"github.com/skyaxl/appetizer/pkg/maze"
)

func main() {
	errs := make(chan error, 2)
	handler := maze.HttpWarmUp()
	go func() {
		fmt.Printf("[TMS-ORDER_API] Has started with address %s\n", port)
		errs <- http.ListenAndServe(":8080")
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("[TMS-ORDER_API] Has stoped. \n%s\n", <-errs)
}
