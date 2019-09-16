package listener

import (
	"github.com/go-ini/ini"
)

//rest and telnet
type input interface {
	start()
	execute(string)
}

var (
	cfg *ini.File
)

func init() {
	cfg, _ = ini.Load("conf.ini")
}

func StartTelnet() {
	telnetSection, _ := cfg.GetSection("telnet")
	if isActive, _ := telnetSection.GetKey("active"); isActive.Value() != "true" {
		return
	}
	host, _ := telnetSection.GetKey("host")
	port, _ := telnetSection.GetKey("port")
	go telnetListener(host.Value(), port.Value())
}

//StartRestServer serve api
func StartRestServer() {
	webSection, _ := cfg.GetSection("web")
	if isActive, _ := webSection.GetKey("active"); isActive.Value() != "true" {
		return
	}
	port, _ := webSection.GetKey("port")
	go StartServer(port.Value())
}
