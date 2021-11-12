package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/radish-miyazaki/go-clean-architecture/config"
)

// NewServer Httpサーバーを開始する
func NewServer(h http.Handler) *http.Server {
	c := config.Config
	return &http.Server{
		Handler:      h,
		Addr:         fmt.Sprintf("0.0.0.0:%s", c.ServerPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
