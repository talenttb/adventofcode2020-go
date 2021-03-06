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
	"sort"

	"github.com/spf13/cobra"
)

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day1 called")
		// day1()
	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day1 anwser
func Day1(arr []int) (int, int) {
	x, y, check := 0, len(arr)-1, 0
	sort.Ints(arr)

	for {
		if x >= y || x > len(arr) {
			return 0, 0
		}

		check = arr[x] + arr[y]
		if check == 2020 {
			return arr[x], arr[y]
		} else if check < 2020 {
			x++
		} else if check > 2020 {
			y--
		}
	}
}

// Day1Part2 ans
func Day1Part2(arr []int) (int, int, int) {
	// x, y, check := 0, len(arr)-1, 0
	sort.Ints(arr)
	// x, y := 0, 0
	for i, v := range arr {
		ok, x, y := day1Part2Help(arr, i)

		if ok {
			return v, x, y
		}
	}
	return 0, 0, 0
	// for {
	// 	if x >= y || x > len(arr) {
	// 		return 0, 0
	// 	}

	// 	check = arr[x] + arr[y]
	// 	if check == 2020 {
	// 		return arr[x], arr[y]
	// 	} else if check < 2020 {
	// 		x++
	// 	} else if check > 2020 {
	// 		y--
	// 	}
	// }
}

func day1Part2Help(arr []int, i int) (bool, int, int) {
	x, y, check := 0, len(arr)-1, 0
	// sort.Ints(arr)
	checkAns := 2020 - arr[i]

	for {
		if x >= y || x > len(arr) {
			return false, 0, 0
		}

		check = arr[x] + arr[y]
		if check == checkAns {
			return true, arr[x], arr[y]
		} else if check < checkAns {
			x++
		} else if check > checkAns {
			y--
		}

		if x == i {
			x++
		}
		if y == i {
			y--
		}
	}
}
