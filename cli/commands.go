// Copyright 2020 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package cli

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mendersoftware/mender-shell/app"
	"github.com/mendersoftware/mender-shell/config"
)

type runOptionsType struct {
	config         string
	fallbackConfig string
}

func initDaemon(config *config.MenderConfig) (*app.MenderShellDaemon, error) {
	daemon := app.NewDaemon(config)
	return daemon, nil
}

func runDaemon(d *app.MenderShellDaemon) error {
	// Handle user forcing update check.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM)
		defer signal.Stop(c)

		for {
			s := <-c // Block until a signal is received.
			if s == syscall.SIGTERM {
				d.StopDaemon()
			}
		}
	}()
	return d.Run()
}
