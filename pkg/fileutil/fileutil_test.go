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
package fileutil

import (
	"testing"
)

func TestGetYamls(t *testing.T) {
	files, err := GetYamls("../../test")
	if err != nil {
		t.Fatal(err)
	}
	if files[0] != "../../test/environment/namespace1/test_cronjob1.yaml" {
		t.Log(files[0])
		t.Fatal("failed test")
	}
	if files[1] != "../../test/environment/namespace2/test_cronjob2.yml" {
		t.Log(files[1])
		t.Fatal("failed test")
	}
}
