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

package main

import (
	"github.com/plars/snappy-test-device-imager/handlers"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/writeimage", handlers.WriteImage)
	http.HandleFunc("/reboot", handlers.Reboot)
	http.HandleFunc("/check", handlers.Check)
	http.HandleFunc("/run", handlers.Runcmd)

	log.Fatal(http.ListenAndServe(":8989", nil))
}

