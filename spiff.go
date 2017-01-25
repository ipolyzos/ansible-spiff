// Copyright (c) 2017 Ioannis Polyzos. All Rights Reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os/exec"
	"io"
	"bytes"
	"os"
	"strings"
	"github.com/lunixbochs/vtclean"
)

var (
	cmd string = "spiff"
)

//implement spiff commands construction and execution
func spiff(moduleArgs ModuleArgs, response *ModuleResponse) bool {
	//validate only one action is provided
	if moduleArgs.Merge != nil && moduleArgs.Diff != nil {
		return fail(response, "Only one action (merge or diff) should be provided!")
	}

	//check if dest file provided
	if moduleArgs.Dest == "" {
		return fail(response, "No dest provided!")
	}

	var args []string
	if moduleArgs.Merge != nil {
		//validate that at least two files provided
		if args = moduleArgs.Merge; len(args) < 1 {
			return fail(response, "At least one files should be provided for the 'merge' operation!")
		}
		args = append([]string{"merge"}, args...)
	} else if moduleArgs.Diff != nil {
		//validate that at only two files provided
		if args = moduleArgs.Diff; len(args) != 2 {
			return fail(response, "Only two files should be provided for the 'diff' operation!")
		}
		args = append([]string{"diff"}, args...)
	} else {
		return fail(response, "Invalid command, or no arguments specified!")
	}

	//creat command and set output to stdout and dest file
	cmd := exec.Command(cmd, args...)

	var outputBuffer bytes.Buffer
	cmd.Stdout = io.Writer(&outputBuffer)

	//execute command
	if err := cmd.Run(); err != nil {
		return fail(response, fmt.Sprintf("There was an error running git spiff command: %v", err))
	}

	//clean contents from
	clnOut := vtclean.Clean(outputBuffer.String(), false)

	//create dest file
	df, err := os.Create(moduleArgs.Dest)
	if err != nil {
		return fail(response, fmt.Sprintf("Cannot create dest file: %s", err))
	}

	//copy contents to file and stdout
	r := strings.NewReader(string(clnOut))
	w := io.MultiWriter(os.Stdout, df)

	if _, err := io.Copy(w, r); err != nil {
		return fail(response, fmt.Sprintf("Cannot copy response to stdout and file: %s", err))
	}

	response.Msg = clnOut + df.Name()
	return true
}