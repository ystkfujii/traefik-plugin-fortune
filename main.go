// Package plugindemo a demo plugin.
package traefik_plugin_fortune

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"
)

// Config the plugin configuration.
type Config struct {
	Rate int `json:"rate,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Rate: 0,
	}
}

// Lottery.
type Lottery struct {
	next     http.Handler
	rate     int
	name     string
	template *template.Template
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Lottery{
		rate:     config.Rate,
		next:     next,
		name:     name,
		template: template.New("lottery").Delims("[[", "]]"),
	}, nil
}

func (l *Lottery) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rand.Seed(time.Now().UnixNano())

	fortune := rand.Intn(100) + 1
	os.Stdout.WriteString(fmt.Sprintf("fortune/Rate: %v/%v \n", fortune, l.rate))
	if l.rate < fortune {
		req.Header.Set("X-FORTUNE", "Bad")
	} else {
		req.Header.Set("X-FORTUNE", "Good")
	}

	l.next.ServeHTTP(rw, req)
}
