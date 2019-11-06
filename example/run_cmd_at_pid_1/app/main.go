package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ListenAllSignal()

	// ListenSpecificSignal()
}

func ListenAllSignal() {
	c := make(chan os.Signal)
	signal.Notify(c)
	fmt.Println("Start......")
	go func() {
		for {
			fmt.Println("I am alive...")
			time.Sleep(time.Second * 1)
		}
	}()
	s := <-c
	fmt.Println("Signal: ", s)
}

func ListenSpecificSignal() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		str := "helloworld"
		fileName := "test.max"
		filepPath := "./" + fileName
		ioutil.WriteFile(filepPath, []byte(str), 0644)
		os.Exit(0)
	}()

	for {
		fmt.Println("I am alive....")
		time.Sleep(time.Second * 1)
	}
}
