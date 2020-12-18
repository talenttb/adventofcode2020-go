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
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day13Cmd represents the day13 command
var day13Cmd = &cobra.Command{
	Use:   "day13",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day13 called")

		file, _ := ioutil.ReadFile("./day13input")
		d := strings.Split(string(file), "\n")

		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day13Part1(d))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day13Part2(d))
		}

	},
}

func init() {
	rootCmd.AddCommand(day13Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day13Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day13Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day13Part1 good
func Day13Part1(input []string) int {
	result := 0
	ts, _ := strconv.Atoi(input[0])
	busID := []int{}
	for _, item := range strings.Split(input[1], ",") {
		ch := string(item)
		if ch == "x" {
			continue
		}
		v, _ := strconv.Atoi(ch)
		busID = append(busID, v)
	}

	takeBus := 0
	tick := 0
	for {
		for _, bus := range busID {
			if ts%bus == 0 {
				takeBus = bus
				break
			}
		}
		if takeBus != 0 {
			break
		}
		ts++
		tick++
	}
	result = takeBus * tick

	return result
}

// Day13Part2 good
func Day13Part2(input []string) uint64 {
	var result uint64 = 0
	return result
}
