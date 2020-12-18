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
	"container/list"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day15Cmd represents the day15 command
var day15Cmd = &cobra.Command{
	Use:   "day15",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day15 called")
		file, _ := ioutil.ReadFile("./day15input")
		d := strings.Split(string(file), "\n")

		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day15Part1(d, 2020))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day15Part2(d, 30000000))
		}
	},
}

func init() {
	rootCmd.AddCommand(day15Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day15Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day15Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day15Part1 good
func Day15Part1(input []string, nth int) int {
	result := 0
	spokenOrder := []int{}
	spokenAudit := make(map[int]bool)
	for _, item := range strings.Split(input[0], ",") {
		v, _ := strconv.Atoi(item)
		spokenOrder = append(spokenOrder, v)
		spokenAudit[v] = true
	}

	turn := len(spokenOrder)
	delete(spokenAudit, spokenOrder[len(spokenOrder)-1])
	// spokenOrder = append(spokenOrder, 0)
	// spokenAudit[0] = true

	for turn < nth {
		turn++

		lastSpoken := spokenOrder[len(spokenOrder)-1]
		if _, ok := spokenAudit[lastSpoken]; ok {
			history := []int{}
			for i := len(spokenOrder) - 1; i >= 0; i-- {
				if spokenOrder[i] == lastSpoken {
					history = append(history, i)
				}

				if len(history) == 2 {
					break
				}
			}
			n := history[0] - history[1]
			spokenOrder = append(spokenOrder, n)
		} else {
			spokenOrder = append(spokenOrder, 0)
			spokenAudit[lastSpoken] = true
			// spokenAudit[0] = true
		}
		// break
	}
	result = spokenOrder[turn-1]
	return result
}

// Day15Part2 good
func Day15Part2(input []string, nth int) int {
	result := 0
	spokenOrder := list.New()
	spokenAudit := make(map[int][]int)
	for i, item := range strings.Split(input[0], ",") {
		v, _ := strconv.Atoi(item)
		spokenOrder.PushBack(v)
		spokenAudit[v] = append(spokenAudit[v], i+1)
	}

	turn := spokenOrder.Len()
	delete(spokenAudit, spokenOrder.Back().Value.(int))

	for turn < nth {
		turn++

		lastSpoken := spokenOrder.Back().Value.(int)
		if lastTwo, ok := spokenAudit[lastSpoken]; ok {
			if len(lastTwo) >= 2 {
				lastTwo = lastTwo[1:]
			}
			idx := spokenOrder.Len()
			for i := spokenOrder.Back(); i != nil; i = i.Prev() {
				if i.Value.(int) == lastSpoken {
					lastTwo = append(lastTwo, idx)
					break
				}
				idx--
			}

			n := lastTwo[1] - lastTwo[0]
			spokenOrder.PushBack(n)
			spokenAudit[lastSpoken] = lastTwo
		} else {
			spokenOrder.PushBack(0)
			spokenAudit[lastSpoken] = append(spokenAudit[lastSpoken], spokenOrder.Len()-1)
		}
	}
	result = spokenOrder.Back().Value.(int)
	return result
}
