package server

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	default_addr string = "0.0.0.0:8080"
)

type Engine struct {
	http.Server
	handlerEngine *gin.Engine
}

func (ins *Engine) With(srv http.Server,
	handle *gin.Engine) {
	ins.handlerEngine = handle
	ins.Server = srv
	ins.Init()
}

func (ins *Engine) Init() bool {
	changed := false
	if is_nil(ins.handlerEngine) {
		ins.handlerEngine = gin.New()
		changed = true
	}
	if len(ins.Addr) == 0 {
		ins.Addr = default_addr
		ins.Handler = ins.handlerEngine
		changed = true
	} else {
		ins.Handler = ins.handlerEngine
	}
	return changed
}

func (ins *Engine) UseLogWriter(w io.Writer) {
	if not_nil(ins.handlerEngine) || ins.Init() {
		loggerFunc := gin.LoggerWithWriter(w)
		ins.handlerEngine.Use(loggerFunc)
	}
}

func (ins *Engine) UseCors(c cors.Config) {
	if not_nil(ins.handlerEngine) || ins.Init() {
		handlerFunc := cors.New(c)
		ins.handlerEngine.Use(handlerFunc)
	}
}

func (ins *Engine) AddHandler(apply func(engine *gin.Engine)) {
	if not_nil(ins.handlerEngine) || ins.Init() {
		apply(ins.handlerEngine)
	}
}

func (ins *Engine) Run(addr ...string) {
	var (
		interrupt chan os.Signal = make(chan os.Signal)
	)
	if len(addr) > 0 {
		ins.Addr = addr[0]
	}
	if (not_nil(ins.handlerEngine) && len(ins.Addr) > 0) || ins.Init() {
		go func() {
			println("Server startting ... addr:", ins.Addr)
			if err := ins.ListenAndServe(); err != nil {
				println("\r\n", err.Error())
			}
			interrupt <- os.Interrupt
		}()

		signal.Notify(interrupt, os.Interrupt)
		<-interrupt

		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()

		if err := ins.Shutdown(ctx); err != nil {
			println("Shutdown error: ", err.Error())
		}

		os.Exit(0)
	}
}
