package apps

import (
	"fmt"
	"net/http/httputil"
	"net/url"

	"github.com/docker/go-connections/nat"
)

func NewProxy(portBinding *nat.PortBinding) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", portBinding.HostIP, portBinding.HostPort),
	})

	return proxy
}
