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

// day8Cmd represents the day8 command
var day8Cmd = &cobra.Command{
	Use:   "day8",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day8 called")
		file, _ := ioutil.ReadFile("./day8input")
		d := strings.Split(string(file), "\n")

		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day8Part1(d))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day8Part2(d))
		}
	},
}

func init() {
	rootCmd.AddCommand(day8Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day8Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day8Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day8Part1 good
func Day8Part1(input []string) int {
	result := 0
	step := 0
	path := make(map[int]bool)
	for {
		op := strings.Split(input[step], " ")
		if _, ok := path[step]; ok {
			break
		}
		path[step] = true

		switch string(op[0]) {
		case "nop":
			step++
		case "jmp":
			i, _ := strconv.Atoi(string(op[1]))
			step += i
		case "acc":
			i, _ := strconv.Atoi(string(op[1]))
			result += i
			step++
		}
	}
	return result
}

// Day8Part2 good
func Day8Part2(input []string) int {
	result := 0
	step := 0
	pathMap := make(map[int]bool)
	path := []int{}
	for {
		path = append(path, step)

		op := strings.Split(input[step], " ")
		if _, ok := pathMap[step]; ok {
			break
		}
		pathMap[step] = true

		switch string(op[0]) {
		case "nop":
			step++
		case "jmp":
			i, _ := strconv.Atoi(string(op[1]))
			step += i
		case "acc":
			i, _ := strconv.Atoi(string(op[1]))
			result += i
			step++
		}
	}
	path = path[:len(path)-1]
	for i := len(path) - 1; i >= 0; i-- {
		t := make([]string, len(input))
		copy(t, input)
		op := strings.Split(t[path[i]], " ")
		switch string(op[0]) {
		case "nop":
			t[path[i]] = strings.Replace(t[path[i]], "nop", "jmp", 1)
		case "jmp":
			t[path[i]] = strings.Replace(t[path[i]], "jmp", "nop", 1)
		default:
			continue
		}

		if result, s := day8Part2(t); s == len(input) {
			return result
		}
	}

	// return result, step

	return 0
}

func day8Part2(input []string) (int, int) {
	result := 0
	step := 0
	pathMap := make(map[int]bool)
	path := []int{}
	for {
		path = append(path, step)

		if _, ok := pathMap[step]; ok || step == len(input) {
			// fmt.Printf("current: %v", step)
			break
		}

		op := strings.Split(input[step], " ")
		pathMap[step] = true

		switch string(op[0]) {
		case "nop":
			step++
		case "jmp":
			i, _ := strconv.Atoi(string(op[1]))
			step += i
		case "acc":
			i, _ := strconv.Atoi(string(op[1]))
			result += i
			step++
		}
	}

	return result, step
}
