/*
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

// day24Cmd represents the day24 command
var day24Cmd = &cobra.Command{
	Use:   "day24",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day24 called")
		file, _ := ioutil.ReadFile("./day24input")
		d := strings.Split(string(file), "\n")

		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day24Part1(d))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day24Part2(d))
		}
	},
}

func init() {
	rootCmd.AddCommand(day24Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day24Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day24Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day24Part1 good
func Day24Part1(input []string) int {
	result := 0
	for _, line := range input {
		ph := []string{}
		for i := 0; i < len(line); i++ {
			end := i + 2
			if (i + 2) > len(line) {
				end--
			}
			switch line[i:end] {
			case "nw":
				ph = append(ph, "nw")
				i++
			case "ne":
				ph = append(ph, "ne")
				i++
			case "sw":
				ph = append(ph, "sw")
				i++
			case "se":
				ph = append(ph, "se")
				i++
			default:
				switch line[i] {
				case 'e':
					ph = append(ph, "e")
				case 'w':
					ph = append(ph, "w")
				default:
					panic("error")
				}
			}
		}
		fmt.Printf("%#v %v -> %v\n", ph, len(ph), len(ph)%2)
		if len(ph)%2 == 1 {
			result++
		}
	}
	return result
}

// Day24Part2 good
func Day24Part2(input []string) int {
	result := 0
	return result
}
