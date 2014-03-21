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
	"fmt"
	"log"
	"net/http"
	"text/template"

	"../core"
	"github.com/ActiveState/tail"
	"github.com/gorilla/websocket"
)

type MData struct {
	Server string `json:"server"`
	Msg    string `json:"msg"`
}

func (self *MData) String() string {
	return self.Server + " says " + self.Msg
}

type connection struct {
	ws   *websocket.Conn
	send chan *MData
}

type room struct {
	connections map[*connection]bool
	broadcast   chan *MData
	register    chan *connection
	unregister  chan *connection
}

var r = room{
	broadcast:   make(chan *MData),
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
			if _, ok := r.connections[c]; ok {
				delete(r.connections, c)
				close(c.send)
			}
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
		var message MData
		err := websocket.ReadJSON(c.ws, &message)
		if err != nil {
			break
		}
		// r.broadcast <- message
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {
		err := websocket.WriteJSON(c.ws, message)
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
	http.HandleFunc("/", handleHomePage(config))
	http.HandleFunc("/ws", handleWebsoket)
	addr := fmt.Sprintf(":%d", port)
	for _, log := range config.Logs {
		go logFile(log)
	}
	fmt.Println("Starting server at ", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

func handleHomePage(config *core.Config) func(http.ResponseWriter, *http.Request) {
	return func(c http.ResponseWriter, req *http.Request) {
		data := make(map[string]interface{})
		data["host"] = req.Host
		data["logs"] = config.Logs
		data["server"] = config.ServerName
		homeTemplate.Execute(c, data)
	}
}
func handleWebsoket(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	if req.Header.Get("Origin") != "http://"+req.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}
	ws, err := websocket.Upgrade(w, req, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}
	c := &connection{send: make(chan *MData, 256), ws: ws}
	r.register <- c
	defer func() { r.unregister <- c }()
	go c.writer()
	c.reader()
}

func logFile(log core.Log) {
	t, _ := tail.TailFile(log.Logpath, tail.Config{Follow: true, ReOpen: true})
	for line := range t.Lines {
		r.broadcast <- &MData{Server: log.Appname, Msg: line.Text}
	}
}
