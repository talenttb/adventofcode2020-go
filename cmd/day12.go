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
	"math"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day12Cmd represents the day12 command
var day12Cmd = &cobra.Command{
	Use:   "day12",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day12 called")

		file, _ := ioutil.ReadFile("./day12input")
		d := strings.Split(string(file), "\n")

		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day12Part1(d))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day12Part2(d))
		}
	},
}

func init() {
	rootCmd.AddCommand(day12Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day12Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day12Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day12Part1 good
func Day12Part1(input []string) int {
	result := 0
	// (e/w, s/n)
	position := []int{0, 0}
	cos := 0
	for _, item := range input {
		v, _ := strconv.Atoi(item[1:])
		switch string(item[0]) {
		case "E":
			position[0] += v
		case "W":
			position[0] -= v
		case "N":
			position[1] += v
		case "S":
			position[1] -= v
		case "L":
			cos += v
			if cos >= 360 {
				cos -= 360
			}
		case "R":
			cos -= v
			if cos < 0 {
				cos += 360
			}
		case "F":
			switch cos {
			case 0:
				position[0] += v
			case 90:
				position[1] += v
			case 180:
				position[0] -= v
			case 270:
				position[1] -= v
			}
		}
		// if string(item[0]) == "R" || string(item[0]) == "L" {
		// 	fmt.Printf("%item\n", item)
		// }
	}
	if position[0] < 0 {
		result -= position[0]
	} else {
		result += position[0]
	}
	if position[1] < 0 {
		result -= position[1]
	} else {
		result += position[1]
	}
	// result = int64(math.Abs(position[0]) + math.Abs(position[1]))
	// result = int32(^uint32(int32(2)-1)) + int32(^uint32(int32(2)-1))
	return result
}

// Day12Part2 good
func Day12Part2(input []string) int {
	result := 0
	// (e/w, s/n)
	position := []int{0, 0}
	waypoint := []int{10, 1}
	var theta float64 = 0
	for _, item := range input {
		v, _ := strconv.Atoi(item[1:])
		switch string(item[0]) {
		case "E":
			waypoint[0] += v
		case "W":
			waypoint[0] -= v
		case "N":
			waypoint[1] += v
		case "S":
			waypoint[1] -= v
		case "L":
			theta = float64(v)

			waypoint[0], waypoint[1] = waypoint[0]*int(math.Cos(theta*math.Pi/180))-waypoint[1]*int(math.Sin(theta*math.Pi/180)), waypoint[0]*int(math.Sin(theta*math.Pi/180))+waypoint[1]*int(math.Cos(theta*math.Pi/180))
		case "R":
			theta = 360.0 - float64(v)

			waypoint[0], waypoint[1] = waypoint[0]*int(math.Cos(theta*math.Pi/180))-waypoint[1]*int(math.Sin(theta*math.Pi/180)), waypoint[0]*int(math.Sin(theta*math.Pi/180))+waypoint[1]*int(math.Cos(theta*math.Pi/180))
		case "F":
			moveX, moveY := waypoint[0]*v, waypoint[1]*v
			position[0] += moveX
			position[1] += moveY
		}
	}

	result = int(math.Abs(float64(position[0])) + math.Abs(float64(position[1])))
	return result
}
