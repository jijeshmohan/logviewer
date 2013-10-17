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

package core

import (
	"testing"
)

func TestConfigDecoding(t *testing.T) {

	json := []byte(`{"serverName": "QA","logs":[{"appName":"app1", "logPath": "/var/log/app1.log"},{"appName":"app2", "logPath": "/var/log/app2.log"}]}`)
	c := parseConfigData(json)
	if c.ServerName != "QA" {
		t.Fatal("Invalid Server name")
	}
	if len(c.Logs) != 2 {
		t.Fatal("Invalid Logs count")
	}

	if c.Logs[0].Appname != "app1" {
		t.Fatalf("Invalid app name : %s", c.Logs[0].Appname)
	}
}
