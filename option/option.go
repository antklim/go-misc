package option

import (
	"fmt"
	"time"
)

type Logger interface {
	Log(args ...interface{})
}

type options struct {
	addr    string        // service address to listen to in format host:port
	cache   bool          // use cache
	ctout   time.Duration // maximum length of an idle connection
	maxcons int           // maximum number of concurrent connections
	logger  Logger
}

var defaultOptions = options{
	addr:    "localhost:8080",
	ctout:   300 * time.Millisecond,
	maxcons: 100,
}

type Option interface {
	apply(*options)
}

type funcOption struct {
	f func(*options)
}

func (fo *funcOption) apply(o *options) {
	fo.f(o)
}

func newFuncOption(f func(*options)) *funcOption {
	return &funcOption{f}
}

type Server struct {
	opts options
}

func NewServer(opt ...Option) *Server {
	opts := defaultOptions
	for _, o := range opt {
		o.apply(&opts)
	}
	s := &Server{opts}
	return s
}

func (s Server) Info() string {
	loggerProvided := s.opts.logger != nil
	return fmt.Sprintf("address: %s\nwith cache: %t\nclient timeout: %v\nmax connections: %d\nlogger provided: %t",
		s.opts.addr, s.opts.cache, s.opts.ctout, s.opts.maxcons, loggerProvided)
}

func Address(s string) Option {
	return newFuncOption(func(o *options) {
		o.addr = s
	})
}

func WithCache(c bool) Option {
	return newFuncOption(func(o *options) {
		o.cache = c
	})
}

func ClientTimeout(d time.Duration) Option {
	return newFuncOption(func(o *options) {
		o.ctout = d
	})
}

func MaxConcurrentConnections(m int) Option {
	return newFuncOption(func(o *options) {
		o.maxcons = m
	})
}

func WithLogger(l Logger) Option {
	return newFuncOption(func(o *options) {
		o.logger = l
	})
}
