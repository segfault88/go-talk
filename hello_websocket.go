// SP1 OMIT
package main

import (
	"code.google.com/p/go.net/websocket"
	"net/http"
	"time"
)

func main() {
	http.Handle("/diffuse", websocket.Handler(diffuseHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

// EP1 OMIT
// SP2 OMIT

func diffuseHandler(ws *websocket.Conn) {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(2500 * time.Millisecond)

	for {
		select {
		case <-tick:
			send(ws, "tick.")
		case <-boom:
			send(ws, "BOOM!")
			ws.Close()
			return
		}
	}
}

func send(ws *websocket.Conn, msg string) {
	t := time.Now().Local().Format(time.StampNano)
	str := t + ": " + msg
	websocket.Message.Send(ws, str)
}

// EP2 OMIT
