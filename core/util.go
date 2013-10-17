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
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(filepath string) []byte {
	data, err := ioutil.ReadFile(filepath)
	checkError(err, "Unable to read file")
	return data
}

func checkError(err error, msg string) {
	if err != nil {
		fmt.Errorf(msg)
		os.Exit(2)
	}
}
