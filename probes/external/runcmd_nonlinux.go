// Copyright 2022 The Cloudprober Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !linux

package external

import (
	"context"
	"os"
	"os/exec"
)

func (p *Probe) runCommand(ctx context.Context, cmd string, args []string, envVars []string) ([]byte, error) {
	toRun := exec.CommandContext(ctx, cmd, args...)
	if len(envVars) > 0 {
		// make sure to pass down environment variables
		toRun.Env = append(toRun.Env, os.Environ()...)
		toRun.Env = append(toRun.Env, envVars...)
	}
	return toRun.Output()
}
