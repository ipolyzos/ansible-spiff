// Copyright (c) 2017 Ioannis Polyzos. All Rights Reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

//module arguments type
type ModuleArgs struct {
	//ansible related args
	AnsibleVersion        string `json:_ansible_version`
	AnsibleNoLog          bool `json:_ansible_no_log`
	AnsibleModuleName     string `json:_ansible_module_name`
	AnsibleVerbosity      int `json:_ansible_verbosity`
	AnsibleSyslogFacility string `json:_ansible_syslog_facility`
	AnsibleDiff           bool `json:_ansible_diff`
	AnsibleDebug          bool `json:_ansible_debug`
	AnsibleCheckMode      bool `json:_ansible_check_mode`
	Ansible               []string `json:_ansible_selinux_special_fs`

	//spiff arguments
	Merge                 []string `json:merge`
	Diff                  []string `json:diff`
	Dest		      string `json:dest`
}

//module output response
type ModuleResponse struct {
	Msg     string `json:"msg"`
	Changed bool   `json:"changed"`
	Failed  bool   `json:"failed"`
}
