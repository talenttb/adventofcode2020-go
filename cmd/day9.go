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
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day9Cmd represents the day9 command
var day9Cmd = &cobra.Command{
	Use:   "day9",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day9 called")

		file, _ := ioutil.ReadFile("./day9input")
		d := strings.Split(string(file), "\n")

		data := []int{}
		for _, v := range d {
			i, _ := strconv.Atoi(v)
			data = append(data, i)
		}
		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day9Part1(data, 25))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day9Part2(data, 25))
		}
	},
}

func init() {
	rootCmd.AddCommand(day9Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day9Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day9Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day9Part1 good
func Day9Part1(input []int, preambleNO int) int {
	result := 0
	findInvalid := false
	queue := make([]int, preambleNO)
	for i := 0; i < preambleNO; i++ {
		queue[i] = input[i]
	}
	input = input[preambleNO:]

	for _, v := range input {
		queueTmp := make([]int, preambleNO)
		copy(queueTmp, queue)
		sort.Ints(queueTmp)
		i, j, check := 0, len(queueTmp)-1, 0

		for {
			if i >= j {
				findInvalid = true
				result = v
				break
			}

			check = queueTmp[i] + queueTmp[j]
			if check == v {
				break
			} else if check < v {
				i++
			} else if check > v {
				j--
			}
		}

		if findInvalid {
			break
		}

		queue = queue[1:]
		queue = append(queue, v)
	}
	return result
}

// Day9Part2 good
func Day9Part2(input []int, preambleNO int) int {
	target := Day9Part1(input, preambleNO)
	i, j := 0, 1
	for {
		sumary := sum(input[i:j]...)
		if target == sumary {
			break
		} else if target < sumary {
			i++
		} else if target > sumary {
			j++
		}
	}
	min, max := 0, 0
	for idx, v := range input[i:j] {
		if idx == 0 || v < min {
			min = v
		}
		if idx == 0 || v > max {
			max = v
		}
	}

	return min + max
}

func sum(xi ...int) int {
	// fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
