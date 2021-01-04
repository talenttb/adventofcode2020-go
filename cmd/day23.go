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
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day23Cmd represents the day23 command
var day23Cmd = &cobra.Command{
	Use:   "day23",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day23 called")
		file, _ := ioutil.ReadFile("./day23input")
		d := strings.Split(string(file), "\n")

		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day23Part1(d))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day23Part2(d))
		}
	},
}

func init() {
	rootCmd.AddCommand(day23Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day23Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day23Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day23Part1 good
func Day23Part1(input []string) int {
	result := 0
	cups := []int{}

	x := strings.Split(input[0], "")
	for i := 0; i < len(x); i++ {
		v, _ := strconv.Atoi(x[i])
		cups = append(cups, v)
	}

	currIdx, cupCount, max, min := 0, 3, 9, 1

	// fmt.Printf("-- move %v --\n", 0)
	// fmt.Printf("cups: %#v\n", cups)
	// fmt.Printf("current index: %v\n", currIdx)

	// dest := 0
	tmpCups := make([]int, len(cups))
	for round := 1; round <= 10000000; round++ {
		pickUp := []int{}
		// tmpCups = cups[:]
		if currIdx >= len(cups) {
			currIdx = 0
		}
		copy(tmpCups, cups)
		for k, j := currIdx+1, 0; j < cupCount; j, k = j+1, k+1 {
			if k >= len(cups) {
				k = 0
			}
			pickUp = append(pickUp, cups[k])
			cups[k] = 0
		}
		src, stopIdx := cups[currIdx]-1, -1
		for stopIdx < 0 {
			if src < min {
				src = max
				continue
			}

			for idx := range cups {
				if src == cups[idx] {
					stopIdx = idx
					break
				}
			}
			src--
		}

		pickAppendIdx := 0
		i, j := currIdx+1, currIdx+1+cupCount
		for {
			if i >= len(cups) {
				i = 0
			}

			if j >= len(cups) {
				j -= len(cups)
			}

			cups[i], cups[j], pickAppendIdx = cups[j], 0, i+1
			if j == stopIdx {
				break
			}
			i++
			j++
		}

		for i := 0; i < len(pickUp); i++ {
			if pickAppendIdx >= len(cups) {
				pickAppendIdx = 0
			}
			cups[pickAppendIdx] = pickUp[i]
			pickAppendIdx++
		}

		// fmt.Printf("-- move %v --\n", round)
		// fmt.Printf("cups: %#v\n", tmpCups)
		// fmt.Printf("next cups: %#v\n", cups)
		// fmt.Printf("pick up: %#v\n", pickUp)
		// fmt.Printf("current(%v): %v\n", currIdx, tmpCups[currIdx])
		// fmt.Printf("destination(%v): %v\n", stopIdx, tmpCups[stopIdx])

		currIdx++
	}

	preOne, postOne := "", ""
	for _, v := range cups {
		if v == 1 || postOne != "" {
			postOne += strconv.Itoa(v)
		} else {
			preOne += strconv.Itoa(v)
		}
	}
	fmt.Printf("%#v", cups)
	// fmt.Printf("%v%v", postOne[1:], preOne)
	// fmt.Printf("%v%v", postOne[1:], preOne)
	result, _ = strconv.Atoi(fmt.Sprintf("%v%v", postOne[1:], preOne))
	fmt.Printf("%v\n", result)
	return result
}

// Day23Part2 good
func Day23Part2(input []string) int {
	result := 0
	return result
}
