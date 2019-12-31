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

	"github.com/spf13/cobra"
)

// addServiceCmd represents the addService command
var addServiceCmd = &cobra.Command{
	Use:   "addService",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if ok := validate.IsArgumentExist(args); !ok {
			fmt.Println("please input your service name")
			fmt.Println("ex: sgoffold addService sgoffold")
			return
		}
		if ok := validate.IsGolangStandardsProjectLayoutDirExist(); !ok {
			fmt.Println("please init before addService command")
			return
		}
		srvName := args[0]
		if err := layout.GenerateService(srvName);err != nil {
			log.Fatalln(err)
			return
		}
		_ = tree.Exec()
	},
}

func init() {
	rootCmd.AddCommand(addServiceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addServiceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addServiceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
