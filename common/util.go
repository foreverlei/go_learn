package common

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func CaptureStopSingal(msg string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	s := <-c
	fmt.Printf("[%s] exit ------- signal:[%v]", msg, s)
}
