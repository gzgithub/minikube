// +build integration

/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package integration

import (
	"flag"
	"os"
	"testing"

	"k8s.io/minikube/test/integration/util"
)

var binaryPath = flag.String("binary", "../../out/minikube", "path to minikube binary")

func TestStartStop(t *testing.T) {

	runner := util.MinikubeRunner{T: t, BinaryPath: *binaryPath}
	runner.RunCommand("delete", false)
	runner.CheckStatus("Does Not Exist")

	runner.RunCommand("start", true)
	runner.CheckStatus("Running")

	runner.RunCommand("stop", true)
	runner.CheckStatus("Stopped")

	runner.RunCommand("start", true)
	runner.CheckStatus("Running")

	runner.RunCommand("delete", true)
	runner.CheckStatus("Does Not Exist")
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
