/*Package cmd ...
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
)

// day6Cmd represents the day6 command
var day6Cmd = &cobra.Command{
	Use:   "day6",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day6 called")
		file, _ := ioutil.ReadFile("./cmd/day6input")
		d := strings.Split(string(file), "\n")

		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day6Part1(d))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day6Part2(d))
		}
	},
}

func init() {
	rootCmd.AddCommand(day6Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day6Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day6Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day6Part1 good
func Day6Part1(input []string) int {
	result := 0
	input = append(input, "")

	isGroup := false
	ansMap := make(map[rune]bool)
	for _, v := range input {
		if v != "" {
			isGroup = false
			for _, ans := range v {
				ansMap[ans] = true
			}
		} else {
			isGroup = true
		}

		if isGroup {
			for k := range ansMap {
				result++
				delete(ansMap, k)
			}
		}
	}
	return result
}

// Day6Part2 good
func Day6Part2(input []string) int {
	result := 0
	input = append(input, "")

	isGroup := false
	groupPeople := 0
	ansMap := make(map[rune]int)
	for _, v := range input {
		if v != "" {
			groupPeople++
			isGroup = false
			for _, ans := range v {
				ansMap[ans]++
			}
		} else {
			isGroup = true
		}

		if isGroup {
			for k, p := range ansMap {
				if p == groupPeople {
					result++
				}

				delete(ansMap, k)
			}
			groupPeople = 0
		}
	}
	return result
}
