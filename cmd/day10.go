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

// day10Cmd represents the day10 command
var day10Cmd = &cobra.Command{
	Use:   "day10",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day10 called")

		file, _ := ioutil.ReadFile("./day10input")
		d := strings.Split(string(file), "\n")

		data := []int{}
		for _, v := range d {
			i, _ := strconv.Atoi(v)
			data = append(data, i)
		}
		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day10Part1(data))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day10Part2(data))
		}
	},
}

func init() {
	rootCmd.AddCommand(day10Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day10Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day10Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day10Part1 good
func Day10Part1(input []int) int {
	result := 0
	sort.Ints(input)
	m := make(map[int]int)
	check := 0
	for i := 0; i < len(input); i++ {
		m[input[i]-check]++
		check = input[i]
	}
	m[3]++

	for k, v := range m {
		fmt.Printf("%v has %v\n", k, v)
	}
	fmt.Printf("1 has %v, 3 has %v\n", m[1], m[3])
	result = m[1] * m[3]
	return result
}

// Day10Part2 good
func Day10Part2(input []int) int {
	result := 1
	sort.Ints(input)
	conti := []int{}
	check := 0
	for i := 0; i < len(input); i++ {
		conti = append(conti, input[i]-check)
		check = input[i]
	}
	conti = append(conti, 3)

	x := ""
	result = 7
	for _, v := range conti[1:] {
		strV := strconv.Itoa(v)
		x += strV

		switch x {
		case "3113":
			result *= 2
		case "31113":
			result *= 4
		case "311113":
			result *= 7
		}

		if strV == "3" {
			x = "3"
		}
	}

	return result
}
