package listener

import (
	"fmt"
	"io"
	"../logger"
	"strings"
	"time"

	telnet "github.com/reiver/go-telnet"

	"../command"
)

type caller struct{}

func (c caller) CallTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	waitForMessage(r)
}

func telnetListener(host string, port string) {
	log := logger.GetInstance()
	log.Printf("telnet connecting to %v:%v\n", host, port)
	err := telnet.DialToAndCall(fmt.Sprintf("%s:%s", host, port), caller{})

	if err != nil {
		log.Fatal(err)
	}
}

func waitForMessage(r telnet.Reader) {
	var buffer [1]byte // Seems like the length of the buffer needs to be small, otherwise will have to wait for buffer to fill up.
	p := buffer[:]
	var message strings.Builder
	periodEnd := time.Now().Add(time.Second * 1) // end after 1 sec
	for {
		command.Add(message.String())
		//reading buffer
		n, err := r.Read(p)
		//reset for next message
		if time.Now().After(periodEnd) {
			periodEnd = time.Now().Add(time.Second * 1)
			message.Reset()
		}
		message.WriteString(string(p[:n]))
		if err == io.EOF {
			break
		}
	}
}
