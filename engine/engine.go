package engine

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/obnahsgnaw/http/cors"
	"github.com/obnahsgnaw/http/corsmid"
	"io"
	"net/http"
	"time"
)

type Config struct {
	Name           string
	DebugMode      bool
	AccessWriter   io.Writer
	ErrWriter      io.Writer
	TrustedProxies []string
	Cors           *cors.Config
	DefFavicon     bool
}

func New(cnf *Config) (*gin.Engine, error) {
	if cnf.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	if cnf.AccessWriter != nil {
		gin.DisableConsoleColor()
	} else {
		gin.ForceConsoleColor()
	}

	r := gin.New()
	if cnf.DefFavicon {
		r.GET("/favicon.ico", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})
	}
	if len(cnf.TrustedProxies) > 0 {
		if err := r.SetTrustedProxies(cnf.TrustedProxies); err != nil {
			return nil, err
		}
	}
	if cnf.AccessWriter != nil {
		r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
			Formatter: func(param gin.LogFormatterParams) string {
				return fmt.Sprintf("[ %s ] - %s %s %s %s %d %s %v %s %s\n",
					param.TimeStamp.Format(time.RFC3339),
					param.ClientIP,
					param.Method,
					cnf.Name,
					param.Path,
					param.StatusCode,
					param.Latency,
					param.Request.Body,
					param.Request.UserAgent(),
					param.ErrorMessage,
				)
			},
			Output: cnf.AccessWriter,
		}))
	} else {
		r.Use(gin.Logger())
	}
	if cnf.ErrWriter != nil {
		r.Use(gin.RecoveryWithWriter(cnf.ErrWriter))
	} else {
		r.Use(gin.Recovery())
	}
	if cnf.Cors != nil {
		r.Use(corsmid.New(func() *cors.Config {
			return cnf.Cors
		}))
	}

	return r, nil
}
