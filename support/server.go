package main

import "golang.org/x/net/websocket"
import "net/http"
import "time"
import "fmt"
import "log"

func main() {
	wsHandler := websocket.Handler(func(ws *websocket.Conn) {
		for {
			time.Sleep(1 * time.Second)
			ws.Write([]byte("💓heartbeat💓"))

			msg := make([]byte, 1024*10)
			n, e := ws.Read(msg)
			if e != nil {
				log.Fatal(e)
			} else {
				ws.Write([]byte("Received: " + string(msg[:n])))
			}
		}
	})

	fmt.Println("serve on 3000")
	http.ListenAndServe(":3000", wsHandler)
}
