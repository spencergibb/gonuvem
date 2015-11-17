package static

import (
	"fmt"
	"github.com/spencergibb/go-nuvem/loadbalancer"
	"github.com/spencergibb/go-nuvem/loadbalancer/serverlist"
	"github.com/spencergibb/go-nuvem/loadbalancer/serverlist/builder"
	"github.com/spf13/viper"
	"net"
	"strconv"
)

type (
	StaticServerList struct {
		Namespace string
		Servers   []string
	}
)

func NewStaticServerList() serverlist.ServerList {
	return &StaticServerList{}
}

func (s *StaticServerList) Init(namespace string) {
	s.Namespace = namespace
	serverConfigs := viper.GetStringSlice(s.GetServerKey())
	fmt.Printf("serverConfigs %+v\n", serverConfigs)
	s.Servers = serverConfigs
}

func (s *StaticServerList) GetServers() []loadbalancer.Server {
	servers := make([]loadbalancer.Server, len(s.Servers))

	for i, config := range s.Servers {
		host, portStr, err := net.SplitHostPort(config)

		port, err := strconv.Atoi(portStr)

		print(err) //TODO: deal with err

		servers[i] = loadbalancer.Server{Host: host, Port: port}
	}

	return servers
}

func (s *StaticServerList) GetServerKey() string {
	key := fmt.Sprintf("loadbalancer.%s.serverlist.static.servers", s.Namespace)
	fmt.Printf("key %+v\n", key)
	return key
}

func init() {
	println("registering static serverlist")
	builder.Register("StaticServerList", NewStaticServerList)
}
