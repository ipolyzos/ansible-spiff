// Copyright (c) 2017 Ioannis Polyzos. All Rights Reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func main() {
	var response ModuleResponse

	// check number or arguments
	if len(os.Args) != 2 {
		fail(&response,"No arguments file provided")
		FailJson(response)
	}

	args := os.Args[1]

	//read provided arguments file contents
	text, err := ioutil.ReadFile(args)
	if err != nil {
		fail(&response,"Could not read configuration file: " + args)
		FailJson(response)
	}

	//unmarshal arguments file contents
	var moduleArgs ModuleArgs
	err = json.Unmarshal(text, &moduleArgs)
	if err != nil {
		fail(&response,"Configuration file not valid JSON: " + args)
		FailJson(response)
	}

	//execute module
	spiff(moduleArgs,&response)

	if (!response.Failed) {
		ExitJson(response)
	}

	FailJson(response)
}
