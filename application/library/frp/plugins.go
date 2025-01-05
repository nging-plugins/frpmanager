package frp

import "github.com/webx-top/echo"

var Plugins = echo.NewKVData().Add(`web`, `内网网站`).
	Add(`ssh`, `SSH`).
	Add("dns", "DNS").
	Add("unix_domain_socket", "Unix域套接字").
	Add("static_file", "文件访问").
	Add("stcp", "STCP").
	Add("p2p", "P2P").
	Add("http_proxy", "HTTP代理").
	Add("socks5", "Socks5代理")
