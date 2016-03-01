package main

import "github.com/mitchellh/colorstring"
import "golang.org/x/net/websocket"
import "github.com/codegangsta/cli"
import "strings"
import "time"
import "fmt"
import "log"
import "os"

// import "io"

func main() {
	app := cli.NewApp()
	app.Name = "ws"
	app.Author = "haoxin"
	app.Version = "0.1.0"
	app.Usage = "💓 WebSocket CLI 💓"

	app.Action = func(c *cli.Context) {
		if len(c.Args()) == 0 {
			fmt.Println("invalid args")
			os.Exit(1)
		}

		url := c.Args()[0]

		ori := strings.Replace(url, "http", "ws", 1)
		ws, err := websocket.Dial(url, "", ori)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		wl := newWsLiner()
		defer wl.Close()

		wl.initHistoryFile()

		go func() {
			for {
				var msg = make([]byte, 1024*10)
				if n, e := ws.Read(msg); e != nil {
					log.Fatal(e)
				} else {
					// clear line
					// io.WriteString(os.Stdout, "\033[2K")
					colorstring.Printf("[green] i> %s\n", msg[:n])
					readAndSend(ws, wl)
				}

				time.Sleep(100 * time.Microsecond)
			}
		}()

		go func() {
			for {
				readAndSend(ws, wl)
				time.Sleep(100 * time.Microsecond)
			}
		}()

		quit := make(chan bool)
		if <-quit {
		}
	}

	app.Run(os.Args)
}

func readAndSend(ws *websocket.Conn, wl *wsLiner) {
	i := wl.readInput()
	if i != "" {
		// send data
		ws.Write([]byte(i))
	}
}
