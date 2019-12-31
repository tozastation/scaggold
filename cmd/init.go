/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"fmt"
	"github.com/tozastation/sgoffold/pkg/layout"
	"github.com/tozastation/sgoffold/pkg/tree"
	"github.com/tozastation/sgoffold/pkg/validate"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	arch string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if ok := validate.IsArgumentExist(args); !ok {
			fmt.Println("please input your package name")
			fmt.Println("ex: sgoffold init github.com/tozastation/sgoffold")
			return
		}
		pkgName := args[0]
		goModInitArgs := []string{
			"mod",
			"init",
			pkgName,
		}
		goModInit := exec.Command("go", goModInitArgs...) // ...enable us to pass them slice
		goModInit.Stdout = os.Stdout
		goModInit.Stderr = os.Stderr
		if err := goModInit.Run(); err != nil {
			log.Fatalln(err)
			return
		}
		if err := layout.GenerateGolangStandardsProjectLayout(); err != nil {
			log.Fatalln(err)
			return
		}
		_ = tree.Exec()
		//TODO: Option. architecture
		// arch, _ := cmd.Flags().GetString("arch")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	// localCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
	initCmd.Flags().StringVar(&arch,"arch", "go-standard-package-layout", "architecture option")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
