package main

import "golang.org/x/net/websocket"
import "github.com/codegangsta/cli"
import "strings"
import "fmt"
import "log"
import "io"
import "os"

func main() {
	app := cli.NewApp()
	app.Name = "ws"
	app.Author = "haoxin"
	app.Version = "0.0.1"
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

		for {
			var msg = make([]byte, 1024*10)
			if n, e := ws.Read(msg); e != nil {
				log.Fatal(e)
			} else {
				fmt.Printf("i> %s\n", msg[:n])
			}

			i := readInput(wl)
			if i != "" {
				// send data
				ws.Write([]byte(i))
			}
		}
	}

	app.Run(os.Args)
}

func readInput(wl *wsLiner) string {
	i, e := wl.Prompt()
	if e != nil {
		if e == io.EOF {
			// break
		}
		log.Fatal(e)
		os.Exit(1)
	}

	return i
}
