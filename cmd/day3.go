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

// day3Cmd represents the day3 command
var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day3 called")
		file, _ := ioutil.ReadFile("./cmd/day3input")
		a := strings.Split(string(file), "\n")
		// fmt.Printf("Reuslt: %v", Day2(a))
		// if part == 1 {
		// 	fmt.Printf("Reuslt: %v", Day3Part1(a))
		// }

		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day3Part1(a))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day3Part2(a))
		}
	},
}

func init() {
	rootCmd.AddCommand(day3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	day3Cmd.Flags().IntVarP(&part, "part", "p", 1, "")
}

// Day3Part1 good
func Day3Part1(arr []string) int {
	trees := 0
	heigh := len(arr) - 1
	weight := len(arr[0])

	right, down := 0, 0
	for {
		right += 3
		down++

		row, col := right, down
		for row >= weight {
			row -= weight
		}

		if string(arr[col][row]) == "#" {
			trees++
		}

		if down == heigh {
			break
		}
	}

	return trees
}

type steps struct {
	right, down int
}

//Day3Part2 a
func Day3Part2(arr []string) int {
	result := 1
	s := []steps{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}

	for _, v := range s {
		x := day3Part2(arr, v)
		fmt.Printf("%v => %v\n", v, x)
		result *= x
	}

	return result
}

func day3Part2(arr []string, s steps) int {
	trees := 0
	heigh := len(arr) - 1
	weight := len(arr[0])

	right, down := 0, 0
	for {
		right += s.right
		down += s.down

		row, col := right, down
		for row >= weight {
			row -= weight
		}

		if string(arr[col][row]) == "#" {
			trees++
		}

		if down == heigh {
			break
		}
	}

	return trees
}
