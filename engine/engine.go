package engine

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/obnahsgnaw/application/pkg/logging/logger"
	"github.com/obnahsgnaw/application/pkg/logging/writer"
	"github.com/obnahsgnaw/http/cors"
	"github.com/obnahsgnaw/http/corsmid"
	"github.com/obnahsgnaw/http/multiwriter"
	"io"
	"net/http"
	"time"
)

type Config struct {
	Name           string
	DebugMode      bool
	LogDebug       bool
	AccessWriter   io.Writer
	ErrWriter      io.Writer
	TrustedProxies []string
	Cors           *cors.Config
	LogCnf         *logger.Config
	DefFavicon     bool
}

func New(cnf *Config) (*gin.Engine, error) {
	if cnf.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	if cnf.AccessWriter == nil {
		if w, err := newDefAccessWriter(cnf.LogCnf, cnf.LogDebug); err != nil {
			return nil, err
		} else {
			cnf.AccessWriter = w
		}
	}
	if cnf.ErrWriter == nil {
		if w, err := newDefErrorWriter(cnf.LogCnf, cnf.LogDebug); err != nil {
			return nil, err
		} else {
			cnf.ErrWriter = w
		}
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

func newDefAccessWriter(logCnf *logger.Config, debug bool) (io.Writer, error) {
	var wts []io.Writer
	if logCnf != nil {
		if w, err := logger.NewAccessWriter(logCnf, debug); err == nil {
			wts = append(wts, w)
		} else {
			return nil, err
		}
	}
	if debug {
		wts = append(wts, writer.NewStdWriter())
	}

	return multiwriter.New(wts...), nil
}

func newDefErrorWriter(logCnf *logger.Config, debug bool) (io.Writer, error) {
	var wts []io.Writer
	if logCnf != nil {
		if w, err := logger.NewErrorWriter(logCnf, debug); err == nil {
			wts = append(wts, w)
		} else {
			return nil, err
		}

	}
	if debug {
		wts = append(wts, writer.NewStdWriter())
	}

	return multiwriter.New(wts...), nil
}
