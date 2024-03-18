package main

import (
	"fmt"
	"time"

	"tawesoft.co.uk/go/log"
	"tawesoft.co.uk/go/log/zerolog"
)

func main() {

	/*ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)


	ctx1, cancel1  := context.WithCancel(ctx)

	select {
	     case <- time.After( 10 * time.Second):
	     	fmt.Println("time after")

	     case <- ctx.Done():
	     	fmt.Println("time cancel")
	     	cancel()

	     case <- ctx1.Done():
	     	fmt.Println("ctx1 done'")

	     case <- time.After(1 * time.Second):
	     	fmt.Println("ctx1 cancel")
	     	cancel1()

	 }


	 time.Sleep(10 * time.Second)*/

	/*for {
		select {
		case <- time.After( 1 * time.Second):
			fmt.Println("time after")

		}

	}*/

	ch := make(chan string)

	go func() {
		select {
		case <-ch:
			fmt.Println("time after 1")

		}
	}()

	go func() {
		select {
		case <-ch:
			fmt.Println("time after 2")

		}
	}()

	time.Sleep(1 * time.Second)
	close(ch)

	time.Sleep(5 * time.Second)

	cfg := log.Config{
		Syslog: log.ConfigSyslog{
			Enabled:  true,
			Network:  "", // local
			Address:  "", // local
			Priority: log.LOG_ERR | log.LOG_DAEMON,
			Tag:      "example",
		},
		File: log.ConfigFile{
			Enabled:          true,
			Mode:             0600,
			Path:             "example.log",
			Rotate:           true,
			RotateCompress:   true,
			RotateMaxSize:    64 * 1024 * 1024, // 64MB
			RotateKeepAge:    30 * 24 * time.Hour,
			RotateKeepNumber: 32, // 32 * 64 MB = 2 GB max storage (before compression)
		},
		Stderr: log.ConfigStderr{
			Enabled: true,
			Color:   true,
		},
	}
	logger, closer, err := zerolog.New(cfg)
	if err != nil {
		panic(err)
	}
	defer closer()

	logger.Info().Msg("Hello world!")

}
