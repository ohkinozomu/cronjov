/*
Copyright © 2021 Nozomu Ohki

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package crontab

import (
	"testing"
)

func TestGetSchedule(t *testing.T) {
	schedule, err := Get("../../test/environment/namespace1/test_cronjob1.yaml", "spec.schedule")
	if err != nil {
		t.Fatal(err)
	}
	if schedule != "0 * * * *" {
		t.Fatal("failed test")
	}
}

func TestGetName(t *testing.T) {
	name, err := Get("../../test/environment/namespace1/test_cronjob1.yaml", "metadata.name")
	if err != nil {
		t.Fatal(err)
	}
	if name != "hello" {
		t.Fatal("failed test")
	}
}

func TestGetKind(t *testing.T) {
	kind, err := Get("../../test/environment/namespace1/test_deployment.yaml", "kind")
	if err != nil {
		t.Fatal(err)
	}
	if kind != "Deployment" {
		t.Fatal("failed test")
	}
}

func TestBuild(t *testing.T) {
	files := []string{"../../test/environment/namespace1/test_cronjob1.yaml", "../../test/environment/namespace2/test_cronjob2.yml"}
	crontab, err := Build(files)
	if err != nil {
		t.Fatal(err)
	}
	if crontab != "0 * * * * hello\n30 0 * * * world\n" {
		t.Fatal("failed test")
	}
}
