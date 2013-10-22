/*
Copyright © 2013 Jijesh Mohan

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package web

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"github.com/ActiveState/tail"
	"github.com/jijeshmohan/logviewer/core"
	"log"
	"net/http"
	"text/template"
)

type connection struct {
	ws   *websocket.Conn
	send chan string
}

type room struct {
	connections map[*connection]bool
	broadcast   chan string
	register    chan *connection
	unregister  chan *connection
}

var r = room{
	broadcast:   make(chan string),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
	connections: make(map[*connection]bool),
}

func (r *room) run() {
	for {
		select {
		case c := <-r.register:
			r.connections[c] = true
		case c := <-r.unregister:
			delete(r.connections, c)
			close(c.send)
		case m := <-r.broadcast:
			for c := range r.connections {
				select {
				case c.send <- m:
				default:
					delete(r.connections, c)
					close(c.send)
					go c.ws.Close()
				}
			}
		}
	}
}

func (c *connection) reader() {
	for {
		var message string
		err := websocket.Message.Receive(c.ws, &message)
		if err != nil {
			break
		}
		// r.broadcast <- message
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {
		err := websocket.Message.Send(c.ws, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

var (
	homeTemplate = template.Must(template.ParseFiles("pages/home.html"))
)

func StartServer(port int, config *core.Config) {
	go r.run()
	http.HandleFunc("/", homeHandler)
	http.Handle("/ws", websocket.Handler(wsHandler))
	addr := fmt.Sprintf(":%d", port)
	for _, log := range config.Logs {
		go logFile(log)
	}
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

func homeHandler(c http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})
	data["host"] = req.Host
	homeTemplate.Execute(c, data)
}

func wsHandler(ws *websocket.Conn) {
	c := &connection{send: make(chan string, 256), ws: ws}
	r.register <- c
	defer func() { r.unregister <- c }()
	go c.writer()
	c.reader()
}

func logFile(log core.Log) {
	t, _ := tail.TailFile(log.Logpath, tail.Config{Follow: true, ReOpen: true})
	for line := range t.Lines {
		r.broadcast <- "{\"server\":\"" + log.Appname + "\",\"log\":\"" + line.Text + "\"}"
	}
}
