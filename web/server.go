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
  "github.com/jijeshmohan/logviewer/core"
)

var (
	homeTemplate = template.Must(template.ParseFiles("pages/home.html"))
)

func StartServer(port int,config *core.Config) {
	http.HandleFunc("/", homeHandler)
	addr := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

func homeHandler(c http.ResponseWriter, req *http.Request) {
	homeTemplate.Execute(c, req.Host)
}
