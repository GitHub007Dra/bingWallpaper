package flag

import (
	"flag"
)

const (
	HostNameHolder = "{HOSTNAME}"
)

var (
	ListenPort int
)

func init() {
	flag.IntVar(&ListenPort, "listen-port", 8701, "api server listen port")
}
