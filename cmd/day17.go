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

// day17Cmd represents the day17 command
var day17Cmd = &cobra.Command{
	Use:   "day17",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day17 called")
		file, _ := ioutil.ReadFile("./day17input")
		d := strings.Split(string(file), "\n")

		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day17Part1(d))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day17Part2(d))
		}
	},
}

func init() {
	rootCmd.AddCommand(day17Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day17Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day17Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day17Part1 good
func Day17Part1(input []string) int {
	result := 0

	// maze := [][][]string{}
	// z := []int{0}
	// y := [][]int{z}
	// x := [][][]int{y}
	// maze = append(maze)
	// maze:=[][][]int{
	// 	[[0,0]],
	// }
	maze := [][][]string{}
	for _, itemx := range input {
		y := [][]string{}
		for _, itemy := range itemx {
			// fmt.Printf("%s\n", string(itemy))
			z := []string{string(itemy)}
			y = append(y, z)
		}
		maze = append(maze, y)
	}
	for cycle := 0; cycle < 1; cycle++ {
		for x := 0; x < len(maze); x++ {
			mazeY := [][]string{}
			for y := 0; y < len(maze[x]); y++ {
				mazeZ := []string{}
				for z := 0; z < len(maze[x][y]); z++ {
					for comb := 0; comb < 27; comb++ {
						// mask := fmt.Sprintf("%03b", differ)
						mask := int2TernaryStr(comb)
						// for i := 0; i < 3; i++ {
						// 	switch mask[i] {
						// 	case '0':

						// 	case '1':
						// 		for switchBit := -1; switchBit < 2; switchBit++ {
						// 		}
						// 	}
						// }
					}
				}
				mazeY = append(mazeY, mazeZ)
			}
			maze = append(maze, mazeY)
		}
	}
	fmt.Printf("%#v\n", maze)

	return result
}

func int2TernaryStr(i int) string {
	result := ""
	switch i {
	case 0:
		result = "000"
	case 1:
		result = "001"
	case 2:
		result = "002"
	case 3:
		result = "010"
	case 4:
		result = "011"
	case 5:
		result = "012"
	case 6:
		result = "020"
	case 7:
		result = "021"
	case 8:
		result = "022"
	case 9:
		result = "100"
	case 10:
		result = "101"
	case 11:
		result = "102"
	case 12:
		result = "110"
	case 13:
		result = "111"
	case 14:
		result = "112"
	case 15:
		result = "120"
	case 16:
		result = "121"
	case 17:
		result = "122"
	case 18:
		result = "200"
	case 19:
		result = "201"
	case 20:
		result = "202"
	case 21:
		result = "210"
	case 22:
		result = "211"
	case 23:
		result = "212"
	case 24:
		result = "220"
	case 25:
		result = "221"
	case 26:
		result = "222"
	}
	return result
}

// Day17Part2 good
func Day17Part2(input []string) int {
	result := 0
	return result
}
