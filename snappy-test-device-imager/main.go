/*
 * Copyright (C) 2014-2015 Canonical Ltd
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

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	http.HandleFunc("/", handler)
	ch := make(chan error, 2)
	go func() {
		ch <- http.ListenAndServe(":8989", nil)
	}()
	return <-ch
}

func handler(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/writeimage" {
		server_arg := req.URL.Query().Get("server")
		s := strings.Split(server_arg, ":")
		if len(s) != 2 {
			http.Error(resp, "ERROR: Specify server as IP:PORT",
				http.StatusBadRequest)
			return
		}
		ip, port := s[0], s[1]
		dev := req.URL.Query().Get("dev")
		ddcmd := fmt.Sprintf("nc %v %v |gunzip|dd bs=16777216 of=%v", ip, port, dev)
		cmd := "/bin/sh"
		args := []string{"-c", ddcmd}
		output, err := exec.Command(cmd, args...).CombinedOutput()
		if err != nil {
			http.Error(resp, string(output), http.StatusBadRequest)
			return
		}
		// Just do a blind call to sync
		exec.Command("/bin/sync").Run()
		resp.Write([]byte(string(output)))
		return
	}
	if req.URL.Path == "/reboot" {
		output, err := exec.Command("/bin/reboot").CombinedOutput()
		if err != nil {
			http.Error(resp, string(output), http.StatusBadRequest)
			return
		}
	}
}
