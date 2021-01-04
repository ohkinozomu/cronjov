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
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/takumakanari/cronv"

	"github.com/ohkinozomu/cronjov/pkg/crontab"
	"github.com/ohkinozomu/cronjov/pkg/fileutil"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := cmd.PersistentFlags().GetString("dir")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if dir == "" {
			fmt.Println("set --dir")
			os.Exit(1)
		}
		opts := cronv.NewCronvCommand()
		ctx, err := cronv.NewCtx(opts)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		files, err := fileutil.GetYamls(dir)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ct, err := crontab.Build(files)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, cron := range strings.Split(ct, "\n") {
			if _, err := ctx.AppendNewLine(cron); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		path, err := ctx.Dump()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("[%s] %d tasks.\n", opts.Title, len(ctx.CronEntries))
		fmt.Printf("[%s] '%s' generated.\n", opts.Title, path)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.PersistentFlags().String("dir", "", "directory")
}
