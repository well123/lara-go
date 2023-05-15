package config

import (
	"encoding/json"
	"fmt"
	"github.com/koding/multiconfig"
	"os"
	"sync"
)

type Config struct {
	RunMode     string
	PrintConfig bool
	Mysql       Mysql
	Log         Log
	CasBin      CasBin
	WWW         string
	Monitor     Monitor
	Http        Http
	Gorm        Gorm
	JWTAuth     JWTAuth
	Redis       Redis
	Swagger     bool
	CORS        CORS
	GZIP        GZIP
	RateLimiter RateLimiter
}

type RateLimiter struct {
	Enable  bool
	Count   int64
	RedisDB int
}

type GZIP struct {
	Enable             bool
	ExcludedExtensions []string
	ExcludedPaths      []string
}

type CORS struct {
	Enable           bool
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	MaxAge           int
}

type Redis struct {
	Addr     string
	Password string
}

type JWTAuth struct {
	Enable        bool
	SigningMethod string
	SigningKey    string
	Expired       int
	Store         string
	FilePath      string
	RedisDB       int
	RedisPrefix   string
}

type Gorm struct {
	Debug             bool
	DBType            string
	MaxLifetime       int
	MaxOpenConns      int
	MaxIdleConns      int
	TablePrefix       string
	EnableAutoMigrate bool
}

type Http struct {
	Host               string
	Port               int
	ShutdownTimeout    int
	MaxContentLength   int
	MaxReqLoggerLength int
	CertFile           string
	KeyFile            string
	MaxResLoggerLength int
}

type LogHook string

func (h LogHook) isGorm() bool {
	return h == "gorm"
}

func MustLoad(filePath string) {
	once.Do(func() {
		m := multiconfig.NewWithPath(filePath)
		m.MustLoad(C)
	})
}

type Monitor struct {
	Enable    bool
	Addr      string
	ConfigDir string
}

type CasBin struct {
	Enable           bool
	Debug            bool
	Model            string
	AutoLoad         bool
	AutoLoadInternal int
}

type Log struct {
	Level         int
	Format        string
	Output        string
	OutputFile    string
	EnableHook    bool
	HookLevels    []string
	Hook          LogHook
	HookMaxThread int
	HookMaxBuffer int
	RotationCount int
	RotationTime  int
}

type Mysql struct {
	Host       string
	Port       int
	User       string
	Password   string
	DBName     string
	Parameters string
}

func (a Mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}

func PrintConfigWithJson() {
	if C.PrintConfig {
		c, err := json.MarshalIndent(C, "", " ")
		if err != nil {
			_, _ = fmt.Fprintf(os.Stdout, "[CONFIG] JSON marshal error: "+err.Error())
			return
		}
		_, _ = fmt.Fprintf(os.Stdout, string(c)+"\n")
	}
}

var (
	C    = new(Config)
	once sync.Once
)
