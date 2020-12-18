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
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// day5Cmd represents the day5 command
var day5Cmd = &cobra.Command{
	Use:   "day5",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day5 called")

		file, _ := ioutil.ReadFile("./cmd/day5input")
		a := strings.Split(string(file), "\n")

		switch part {
		case 1:
			r := 0
			for _, v := range a {
				if x := Day5Part1(v); x > r {
					r = x
				}
			}
			fmt.Printf("Reuslt: %v\n", r)
		case 2:
			g := []int{}
			for _, v := range a {
				g = append(g, Day5Part1(v))
			}
			sort.Ints(g)
			check := 80
			for _, v := range g {
				if check != v {
					fmt.Printf("%v", check)
					panic("")
				}
				check++
			}
			fmt.Printf("%v", g)
			// fmt.Printf("Reuslt: %v\n", r)
		}
	},
}

func init() {
	rootCmd.AddCommand(day5Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day5Part1 good
func Day5Part1(seatID string) int {
	result := 0

	rowRangeMin, rowRangeMax := 0, 127
	rowSource := seatID[:7]

	for _, v := range rowSource {
		mid := (rowRangeMin + rowRangeMax) / 2
		switch string(v) {
		case "F":
			rowRangeMax = mid
		case "B":
			rowRangeMin = mid
		}
	}

	colRangeMin, colRangeMax := 0, 7
	colSource := seatID[7:]

	for _, v := range colSource {
		mid := (colRangeMin + colRangeMax) / 2
		switch string(v) {
		case "L":
			colRangeMax = mid
		case "R":
			colRangeMin = mid
		}
	}

	result = rowRangeMax*8 + colRangeMax

	return result
}
