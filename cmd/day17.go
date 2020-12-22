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
			fmt.Printf("Reuslt: %v\n", Day17Part1(d, 6))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day17Part2(d, 6))
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
func Day17Part1(input []string, round int) int {
	result := 0

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
	for cycle := 0; cycle < round; cycle++ {
		newMaze := [][][]string{}

		sizeX, sizeY, sizeZ := len(maze), len(maze[0]), len(maze[0][0])
		fmt.Printf("size: %vx%vx%v\n", sizeX, sizeY, sizeZ)
		// printMaze(maze)

		defaultCubesX := [][]string{}
		defaultCubesY := []string{}

		for x := 0; x < sizeX+2; x++ {
			defaultCubesX = [][]string{}
			for y := 0; y < sizeY+2; y++ {
				defaultCubesY = []string{}
				for z := 0; z < sizeZ+2; z++ {
					if x > 0 && x <= sizeX &&
						y > 0 && y <= sizeY &&
						z > 0 && z <= sizeZ {
						defaultCubesY = append(defaultCubesY, maze[x-1][y-1][z-1])
					} else {
						defaultCubesY = append(defaultCubesY, ".")
					}
				}
				defaultCubesX = append(defaultCubesX, defaultCubesY)
			}
			newMaze = append(newMaze, defaultCubesX)
		}

		maze = myClone(newMaze)

		for x := 0; x < len(newMaze); x++ {
			for y := 0; y < len(newMaze[x]); y++ {
				for z := 0; z < len(newMaze[x][y]); z++ {
					activeCount := 0
					for comb := 0; comb < 27; comb++ {
						// mask := fmt.Sprintf("%03b", differ)
						mask := int2TernaryStr(comb)
						dMask := []int{}
						for _, item := range mask {
							v, _ := strconv.Atoi(string(item))
							dMask = append(dMask, v-1)
						}
						dx, dy, dz := dMask[0], dMask[1], dMask[2]
						if dx == 0 && dy == 0 && dz == 0 {
							continue
						}
						if x+dx < 0 || x+dx >= len(newMaze) {
							continue
						}
						if y+dy < 0 || y+dy >= len(newMaze[x]) {
							continue
						}
						if z+dz < 0 || z+dz >= len(newMaze[x][y]) {
							continue
						}

						if newMaze[x+dx][y+dy][z+dz] == "#" {
							activeCount++
						}
					}

					switch newMaze[x][y][z] {
					case "#":
						if activeCount == 2 || activeCount == 3 {
							maze[x][y][z] = "#"
						} else {
							maze[x][y][z] = "."
						}
					case ".":
						if activeCount == 3 {
							maze[x][y][z] = "#"
						} else {
							maze[x][y][z] = "."
						}
					}
				}
			}
		}
		// maze = newMaze
	}
	// for z := 0; z < len(maze[0][0]); z++ {
	// 	fmt.Printf("for z=%v:\n", z)
	// 	fmt.Printf("=======\n")
	// 	for x := 0; x < len(maze); x++ {
	// 		for y := 0; y < len(maze[x]); y++ {
	// 			fmt.Printf("%v", maze[x][y][z])
	// 		}
	// 		fmt.Printf("\n")
	// 	}
	// 	fmt.Printf("=======\n")
	// }
	for x := 0; x < len(maze); x++ {
		for y := 0; y < len(maze[x]); y++ {
			for z := 0; z < len(maze[x][y]); z++ {
				if maze[x][y][z] == "#" {
					result++
				}
			}
		}
	}

	return result
}

func myClone(maze [][][]string) [][][]string {
	result := [][][]string{}
	for i := 0; i < len(maze); i++ {
		arri := [][]string{}
		for j := 0; j < len(maze[i]); j++ {
			arrj := []string{}
			arrj = append(arrj, maze[i][j]...)
			// for k := 0; k < len(maze[j]); k++ {
			// 	arrj
			// }
			arri = append(arri, arrj)
		}
		result = append(result, arri)
	}
	return result
}

func printMaze(maze [][][]string) {
	for z := 0; z < len(maze[0][0]); z++ {
		fmt.Printf("for z=%v:\n", z)
		fmt.Printf("=======\n")
		for x := 0; x < len(maze); x++ {
			for y := 0; y < len(maze[x]); y++ {
				fmt.Printf("%v", maze[x][y][z])
			}
			fmt.Printf("\n")
		}
		fmt.Printf("=======\n")
	}
	// fmt.Printf("*********\n")
	// for i := 0; i < len(maze); i++ {
	// 	for j := 0; j < len(maze[i]); j++ {
	// 		fmt.Printf("%#v\n", maze[i][j])
	// 	}
	// 	fmt.Printf("==============\n")
	// }
	// fmt.Printf("*********\n")
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
func Day17Part2(input []string, round int) int {
	result := 0
	return result
}
