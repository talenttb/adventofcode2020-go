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

// day18Cmd represents the day18 command
var day18Cmd = &cobra.Command{
	Use:   "day18",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day18 called")
		file, _ := ioutil.ReadFile("./day18input")
		d := strings.Split(string(file), "\n")

		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day18Part1(d))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day18Part2(d))
		}
	},
}

func init() {
	rootCmd.AddCommand(day18Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day18Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day18Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day18Part1 good
func Day18Part1(input []string) int {
	result := 0
	for _, v := range input {
		result += calculateLine(v)
	}
	return result
}

func calculateLine(input string) int {
	result := 0
	stack := []string{}
	parenthesesLeft := []int{}
	for _, v := range strings.Split(input, "") {
		switch v {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			stack = append(stack, v)
		case "+", "*":
			stack = append(stack, v)
		case "(":
			parenthesesLeft = append(parenthesesLeft, len(stack))
		case ")":
			exParenthesesL := parenthesesLeft[len(parenthesesLeft)-1]
			parenthesesLeft = parenthesesLeft[:len(parenthesesLeft)-1]
			stack = append(stack, calculate(stack, exParenthesesL, len(stack)))
		}
	}
	result, _ = strconv.Atoi(calculate(stack, 0, len(stack)))
	return result
}

func calculate(stack []string, f, t int) string {
	op := []string{}
	for i := f; i < t; i++ {
		if stack[i] == "" {
			continue
		}

		if len(op) != 2 {
			op = append(op, stack[i])
			stack[i] = ""
			continue
		}
		a, _ := strconv.Atoi(op[0])
		b, _ := strconv.Atoi(stack[i])
		switch op[1] {
		case "+":
			op = []string{}
			op = append(op, fmt.Sprintf("%v", a+b))
		case "*":
			op = []string{}
			op = append(op, fmt.Sprintf("%v", a*b))
		}
		stack[i] = ""
	}
	return op[0]
}

// Day18Part2 good
func Day18Part2(input []string) int {
	result := 0
	for _, v := range input {
		result += calculateLineP2(v)
	}
	return result
}

func calculateLineP2(input string) int {
	result := 0
	stack := []string{}
	parenthesesLeft := []int{}
	for _, v := range strings.Split(input, "") {
		switch v {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			stack = append(stack, v)
		case "+", "*":
			stack = append(stack, v)
		case "(":
			parenthesesLeft = append(parenthesesLeft, len(stack))
		case ")":
			exParenthesesL := parenthesesLeft[len(parenthesesLeft)-1]
			parenthesesLeft = parenthesesLeft[:len(parenthesesLeft)-1]
			stack = append(stack, calculateP2(stack, exParenthesesL, len(stack)))
		}
	}
	result, _ = strconv.Atoi(calculateP2(stack, 0, len(stack)))
	return result
}

func calculateP2(stack []string, f, t int) string {
	op := []string{}
	for i := f; i < t; i++ {
		if stack[i] == "" {
			continue
		}

		if len(op) != 2 {
			op = append(op, stack[i])
			stack[i] = ""
			continue
		}
		a, _ := strconv.Atoi(op[0])
		b, _ := strconv.Atoi(stack[i])
		switch op[1] {
		case "+":
			op = []string{}
			op = append(op, fmt.Sprintf("%v", a+b))
		case "*":
			for j := i - 1; len(op) > 0; j-- {
				if stack[j] == "" {
					stack[j] = op[len(op)-1]
					op = op[:len(op)-1]
				}
			}
			op = append(op, fmt.Sprintf("%v", stack[i]))
		}
		stack[i] = ""
	}
	for j := t - 1; len(op) > 0; j-- {
		if stack[j] == "" {
			stack[j] = op[len(op)-1]
			op = op[:len(op)-1]
		}
	}

	for i := f; i < t; i++ {
		if stack[i] == "" {
			continue
		}

		if len(op) != 2 {
			op = append(op, stack[i])
			stack[i] = ""
			continue
		}
		a, _ := strconv.Atoi(op[0])
		b, _ := strconv.Atoi(stack[i])
		switch op[1] {
		case "*":
			op = []string{}
			op = append(op, fmt.Sprintf("%v", a*b))
		}
		stack[i] = ""
	}
	return op[0]
}
