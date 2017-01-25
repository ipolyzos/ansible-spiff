// Copyright (c) 2017 Ioannis Polyzos. All Rights Reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//utility function to wrap boilerplate code
func fail(response *ModuleResponse, reason string) bool {
	response.Msg = reason
	response.Failed = true
	return false
}

//wrap successful module response
func ExitJson(responseBody ModuleResponse) {
	rtrn(responseBody)
}

// wrap failure module response
func FailJson(responseBody ModuleResponse) {
	if responseBody.Failed != false {
		panic("Failed response expected in FailJson call!")
	}
	rtrn(responseBody)
}

// marshall and print module response
func rtrn(responseBody ModuleResponse) {
	var response []byte
	var err error

	response, err = json.Marshal(responseBody)
	if err != nil {
		response, _ = json.Marshal(ModuleResponse{Msg: "Invalid response object"})
	}

	//print response to stdout
	fmt.Fprintln(os.Stdout,string(response))
	if responseBody.Failed {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
