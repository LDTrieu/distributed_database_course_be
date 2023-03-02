package server

import (
	"context"
	"os"
	"os/signal"
	"time"
)

func Detached(srv ...*Engine) {
	var (
		interrupt chan os.Signal = make(chan os.Signal)
	)
	for i, s := range srv {
		if (not_nil(s.handlerEngine) && len(s.Addr) > 0) || s.Init() {
			go func() {
				println("[", i, "] Server is detached and listening at addr", s.Addr)
				if err := s.ListenAndServe(); err != nil {
					println("\r\n", err.Error())
				}
				interrupt <- os.Interrupt
			}()
			time.Sleep(time.Second)
		}
	}

	signal.Notify(interrupt, os.Interrupt)
	<-interrupt

	for i := range srv {
		func() {
			ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
			defer cancel()
			if err := srv[i].Shutdown(ctx); err != nil {
				println("Shutdown error: ", err.Error())
			}
		}()
	}
	os.Exit(0)
}
