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
	"github.com/bmatcuk/doublestar/v3"
	"github.com/ohkinozomu/cronjov/pkg/crontab"
	"github.com/thoas/go-funk"
)

func GetYamls(dir string) ([]string, error) {
	files, err := doublestar.Glob("{" + dir + "/**/*.yaml," + dir + "/**/*.yml}")
	if err != nil {
		return nil, nil
	}
	r := funk.FilterString(files, func(file string) bool {
		// TODO: error handling
		kind, _ := crontab.Get(file, "kind")
		return kind == "CronJob"
	})
	return r, nil
}
