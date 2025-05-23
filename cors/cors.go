package cors

import "strings"

var (
	AllowMethods  = []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"}
	AllowHeaders  = []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization", "Language", "Request-Origin", "X-App-Id", "X-Security-Sign", "X-Security-Iv", "X-Request-Id"}
	ExposeHeaders = []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Cache-Control", "Content-Language", "Content-Type", "X-App-Id", "X-Security-Sign", "X-Security-Iv", "X-Request-Id"}
)

type Config struct {
	AllowOrigin      string   `json:"allow_origin" yaml:"allow_origin" long:"cors-allow-origin" description:"cors allow origin"`
	AllowMethods     []string `json:"allow_methods" yaml:"allow_methods" long:"cors-allow-methods" description:"cors allow methods"`
	AllowHeaders     []string `json:"allow_headers" yaml:"allow_headers" long:"cors-allow-headers" description:"cors allow headers"`
	ExposeHeaders    []string `json:"expose_headers" yaml:"expose_headers" long:"expose-headers" description:"cors expose headers"`
	AllowCredentials bool     `json:"allow_credentials" yaml:"allow_credentials"  long:"cors-allow-credentials" description:"cors allow credentials"`
}

func (c Config) Origin() string {
	if strings.ToLower(c.AllowOrigin) == "all" {
		c.AllowOrigin = "*"
	}
	return c.AllowOrigin
}

func (c Config) Methods() string {
	var m []string
	if len(c.AllowMethods) == 0 {
		m = AllowMethods
	} else {
		m = c.AllowMethods
	}
	return strings.Join(m, ",")
}

func (c Config) ReqHeader() string {
	var m []string
	if len(c.AllowHeaders) == 0 {
		m = AllowHeaders
	} else {
		m = c.AllowHeaders
	}
	return strings.Join(m, ",")
}

func (c Config) RespHeader() string {
	var m []string
	if len(c.ExposeHeaders) == 0 {
		m = ExposeHeaders
	} else {
		m = c.ExposeHeaders
	}
	return strings.Join(m, ",")
}

func (c Config) Credential() string {
	if c.AllowCredentials {
		return "true"
	}
	return "false"
}
