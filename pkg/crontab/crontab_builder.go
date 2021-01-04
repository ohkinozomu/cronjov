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
	"io/ioutil"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
)

func Get(file string, key string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	path, err := yaml.PathString("$." + key)
	if err != nil {
		return "", err
	}

	var value string
	if err := path.Read(strings.NewReader(string(b)), &value); err != nil {
		return "", err
	}
	return value, nil
}

func Build(files []string) (string, error) {
	var crontab string

	for _, file := range files {
		schedule, err := Get(file, "spec.schedule")
		if err != nil {
			return "", err
		}
		name, err := Get(file, "metadata.name")
		if err != nil {
			return "", err
		}
		crontab += schedule + " " + name + "\n"
	}
	return crontab, nil
}
