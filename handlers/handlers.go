/*
 * Copyright (C) 2015-2016 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package handlers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func WriteImage(resp http.ResponseWriter, req *http.Request) {
	server_arg := req.URL.Query().Get("server")
	s := strings.Split(server_arg, ":")
	if len(s) != 2 {
		http.Error(resp, "ERROR: Specify server as IP:PORT",
			http.StatusBadRequest)
		return
	}
	ip, port := s[0], s[1]
	dev := req.URL.Query().Get("dev")
	if _, err := os.Stat(dev); os.IsNotExist(err) {
		http.Error(resp, "ERROR: block device: " + dev + " does not exist", http.StatusBadRequest)
		return
	}
	ddcmd := fmt.Sprintf("nc %v %v |gunzip|dd bs=16777216 of=%v", ip, port, dev)
	cmd := "/bin/sh"
	args := []string{"-c", ddcmd}
	output, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		http.Error(resp, string(output), http.StatusBadRequest)
		return
	}
	// Just do a blind call to sync and hdparm -z
	exec.Command("/bin/sync").Run()
	exec.Command("/sbin/hdparm", "-z", dev).Run()
	resp.Write([]byte(string(output)))
	return
}

func Reboot(resp http.ResponseWriter, req *http.Request) {
	go exec.Command("/bin/reboot").Run()
	resp.Write([]byte("Rebooting target device"))
}

func Check(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte(string("Snappy Test Device Imager")))
}

func Runcmd(resp http.ResponseWriter, req *http.Request) {
    cmd_arg := req.URL.Query().Get("cmd")
    cmd := strings.Fields(cmd_arg)
	output, err := exec.Command(cmd[0], cmd[1:len(cmd)]...).CombinedOutput()
	if err != nil {
		http.Error(resp, string(output), http.StatusBadRequest)
		return
	}
	resp.Write([]byte(string(output)))
    return
}
