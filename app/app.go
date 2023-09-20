package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	Config ServerConfig
	Engine *gin.Engine
}

type ServerConfig struct {
	HostPort     int
	ReadTimeout  int
	WriteTimeout int
	RunMode      string
}

var Instance *App

// NewApp 构建App
func NewApp(config ServerConfig) *App {
	app := &App{
		Config: config,
	}
	Instance = app
	return Instance
}

// RegisterRouter 注册路由
func (a *App) RegisterRouter(router func(eng *gin.Engine) error) *App {
	a.Engine = gin.Default()
	router(a.Engine)
	return a
}

// Run 启动服务
func (a *App) Run() {
	host := fmt.Sprintf(":%d", a.Config.HostPort)
	gin.SetMode(a.Config.RunMode)
	s := &http.Server{
		Addr:           host,
		Handler:        a.Engine,
		ReadTimeout:    time.Duration(a.Config.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(a.Config.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		log.Println("Server Listen at:" + host)
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen:%s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	s.RegisterOnShutdown(func() {
		log.Println("Server exited")
	})
	log.Println("Server exiting")
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
