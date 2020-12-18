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

// day2Cmd represents the day2 command
var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day2 called")
		file, _ := ioutil.ReadFile("./cmd/day2input")
		a := strings.Split(string(file), "\n")
		// fmt.Printf("Reuslt: %v", Day2(a))
		fmt.Printf("Reuslt: %v", Day2Part2(a))
	},
}

func init() {
	rootCmd.AddCommand(day2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day2 anwser
func Day2(arr []string) int {
	result := 0
	for _, v := range arr {
		d := strings.Split(v, " ")
		minMax := strings.Split(d[0], "-")
		str := strings.Split(d[1], ":")

		check := 0
		for _, ch := range d[2] {
			if ch == ([]rune(str[0])[0]) {
				check++
			}
		}

		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])

		if min <= check && check <= max {
			result++
		}
	}
	return result
}

// Day2Part2 anwser
func Day2Part2(arr []string) int {
	result := 0
	for _, v := range arr {
		d := strings.Split(v, " ")
		minMax := strings.Split(d[0], "-")
		str := strings.Split(d[1], ":")
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])
		p := ([]rune(d[2])[min-1]) == ([]rune(str[0])[0])
		q := ([]rune(d[2])[max-1]) == ([]rune(str[0])[0])
		if p != q {
			result++
		}
		//  else {
		// 	fmt.Printf("%v, %s[%v] is %v == %v, %s[%v] is %v != %v\n", d, d[2], min-1, []rune(d[2])[min-1], []rune(str[0])[0], d[2], max-1, []rune(d[2])[max-1], []rune(str[0])[0])
		// }
	}
	return result
}
