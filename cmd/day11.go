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

// day11Cmd represents the day11 command
var day11Cmd = &cobra.Command{
	Use:   "day11",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day11 called")

		file, _ := ioutil.ReadFile("./day11input")
		d := strings.Split(string(file), "\n")

		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day11Part1(d))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day11Part2(d))
		}
	},
}

func init() {
	rootCmd.AddCommand(day11Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day11Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day11Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day11Part1 good
func Day11Part1(input []string) int {
	result := 0

	seatStr := ""
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			seatStr += string(input[i][j])
		}
	}

	seatState := map[string]bool{
		seatStr: true,
	}

	for {
		seatStr = ""
		newState := make([][]string, len(input))
		for i := 0; i < len(input); i++ {
			newState[i] = make([]string, len(input[i]))

			for j := 0; j < len(input[i]); j++ {
				position := string(input[i][j])
				switch position {
				case "L":
					checkEmpty := true
					if i-1 >= 0 && j-1 >= 0 && string(input[i-1][j-1]) != "." && string(input[i-1][j-1]) == "#" {
						checkEmpty = checkEmpty && false
					}
					if j-1 >= 0 && string(input[i][j-1]) != "." && string(input[i][j-1]) == "#" {
						checkEmpty = checkEmpty && false
					}
					if i+1 < len(input) && j-1 >= 0 && string(input[i+1][j-1]) != "." && string(input[i+1][j-1]) == "#" {
						checkEmpty = checkEmpty && false
					}
					// ===============
					if i-1 >= 0 && string(input[i-1][j]) != "." && string(input[i-1][j]) == "#" {
						checkEmpty = checkEmpty && false
					}
					if i+1 < len(input) && string(input[i+1][j]) != "." && string(input[i+1][j]) == "#" {
						checkEmpty = checkEmpty && false
					}
					// ===============
					if i-1 >= 0 && j+1 < len(input[i]) && string(input[i-1][j+1]) != "." && string(input[i-1][j+1]) == "#" {
						checkEmpty = checkEmpty && false
					}
					if j+1 < len(input[i]) && string(input[i][j+1]) != "." && string(input[i][j+1]) == "#" {
						checkEmpty = checkEmpty && false
					}
					if i+1 < len(input) && j+1 < len(input[i]) && string(input[i+1][j+1]) != "." && string(input[i+1][j+1]) == "#" {
						checkEmpty = checkEmpty && false
					}

					if checkEmpty {
						newState[i][j] = "#"
						seatStr += "#"
					} else {
						newState[i][j] = "L"
						seatStr += "L"
					}

				case "#":
					count := 0
					if i-1 >= 0 && j-1 >= 0 && string(input[i-1][j-1]) != "." && string(input[i-1][j-1]) == "#" {
						count++
					}
					if j-1 >= 0 && string(input[i][j-1]) == "#" {
						count++
					}
					if i+1 < len(input) && j-1 >= 0 && string(input[i+1][j-1]) != "." && string(input[i+1][j-1]) == "#" {
						count++
					}
					// ===============
					if i-1 >= 0 && string(input[i-1][j]) != "." && string(input[i-1][j]) == "#" {
						count++
					}
					if i+1 < len(input) && string(input[i+1][j]) != "." && string(input[i+1][j]) == "#" {
						count++
					}
					// ===============
					if i-1 >= 0 && j+1 < len(input[i]) && string(input[i-1][j+1]) != "." && string(input[i-1][j+1]) == "#" {
						count++
					}
					if j+1 < len(input[i]) && string(input[i][j+1]) != "." && string(input[i][j+1]) == "#" {
						count++
					}
					if i+1 < len(input) && j+1 < len(input[i]) && string(input[i+1][j+1]) != "." && string(input[i+1][j+1]) == "#" {
						count++
					}

					if count >= 4 {
						newState[i][j] = "L"
						seatStr += "L"
					} else {
						newState[i][j] = "#"
						seatStr += "#"
					}

				case ".":
					newState[i][j] = "."
					seatStr += "."
				}
			}
		}

		if _, ok := seatState[seatStr]; ok {
			break
		} else {
			seatState[seatStr] = true
			input = []string{}

			for _, v := range newState {
				input = append(input, strings.Join(v[:], ""))
			}
		}
	}

	result = strings.Count(seatStr, "#")

	return result
}

// Day11Part2 good
func Day11Part2(input []string) int {
	result := 0

	seatStr := ""
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			seatStr += string(input[i][j])
		}
	}

	seatState := map[string]bool{
		seatStr: true,
	}

	for {
		seatStr = ""
		newState := make([][]string, len(input))
		for i := 0; i < len(input); i++ {
			newState[i] = make([]string, len(input[i]))

			for j := 0; j < len(input[i]); j++ {
				position := string(input[i][j])
				switch position {
				case "L":
					checkEmpty := true
					px, py := 0, 0
					px, py = i-1, j-1
					for px >= 0 && py >= 0 {
						dot := string(input[px][py])
						if dot == "#" {
							checkEmpty = checkEmpty && false
							break
						} else if dot == "L" {
							break
						}

						px--
						py--
					}

					px, py = i, j-1
					for py >= 0 {
						dot := string(input[px][py])
						if dot == "#" {
							checkEmpty = checkEmpty && false
							break
						} else if dot == "L" {
							break
						}

						py--
					}

					px, py = i+1, j-1
					for px < len(input) && py >= 0 {
						dot := string(input[px][py])
						if dot == "#" {
							checkEmpty = checkEmpty && false
							break
						} else if dot == "L" {
							break
						}

						px++
						py--
					}
					// ===============
					px, py = i-1, j
					for px >= 0 {
						dot := string(input[px][py])
						if dot == "#" {
							checkEmpty = checkEmpty && false
							break
						} else if dot == "L" {
							break
						}

						px--
					}

					px, py = i+1, j
					for px < len(input) {
						dot := string(input[px][py])
						if dot == "#" {
							checkEmpty = checkEmpty && false
							break
						} else if dot == "L" {
							break
						}

						px++
					}
					// ===============
					px, py = i-1, j+1
					for px >= 0 && py < len(input[i]) {
						dot := string(input[px][py])
						if dot == "#" {
							checkEmpty = checkEmpty && false
							break
						} else if dot == "L" {
							break
						}

						px--
						py++
					}

					px, py = i, j+1
					for py < len(input[i]) {
						dot := string(input[px][py])
						if dot == "#" {
							checkEmpty = checkEmpty && false
							break
						} else if dot == "L" {
							break
						}

						py++
					}

					px, py = i+1, j+1
					for px < len(input) && py < len(input[i]) {
						dot := string(input[px][py])
						if dot == "#" {
							checkEmpty = checkEmpty && false
							break
						} else if dot == "L" {
							break
						}

						px++
						py++
					}

					if checkEmpty {
						newState[i][j] = "#"
						seatStr += "#"
					} else {
						newState[i][j] = "L"
						seatStr += "L"
					}

				case "#":
					count := 0
					px, py := 0, 0
					px, py = i-1, j-1
					for px >= 0 && py >= 0 {
						dot := string(input[px][py])
						if dot == "#" {
							count++
							break
						} else if dot == "L" {
							break
						}

						px--
						py--
					}

					px, py = i, j-1
					for py >= 0 {
						dot := string(input[px][py])
						if dot == "#" {
							count++
							break
						} else if dot == "L" {
							break
						}

						py--
					}

					px, py = i+1, j-1
					for px < len(input) && py >= 0 {
						dot := string(input[px][py])
						if dot == "#" {
							count++
							break
						} else if dot == "L" {
							break
						}

						px++
						py--
					}
					// ===============
					px, py = i-1, j
					for px >= 0 {
						dot := string(input[px][py])
						if dot == "#" {
							count++
							break
						} else if dot == "L" {
							break
						}

						px--
					}

					px, py = i+1, j
					for px < len(input) {
						dot := string(input[px][py])
						if dot == "#" {
							count++
							break
						} else if dot == "L" {
							break
						}

						px++
					}
					// ===============
					px, py = i-1, j+1
					for px >= 0 && py < len(input[i]) {
						dot := string(input[px][py])
						if dot == "#" {
							count++
							break
						} else if dot == "L" {
							break
						}

						px--
						py++
					}

					px, py = i, j+1
					for py < len(input[i]) {
						dot := string(input[px][py])
						if dot == "#" {
							count++
							break
						} else if dot == "L" {
							break
						}

						py++
					}

					px, py = i+1, j+1
					for px < len(input) && py < len(input[i]) {
						dot := string(input[px][py])
						if dot == "#" {
							count++
							break
						} else if dot == "L" {
							break
						}

						px++
						py++
					}

					if count >= 5 {
						newState[i][j] = "L"
						seatStr += "L"
					} else {
						newState[i][j] = "#"
						seatStr += "#"
					}

				case ".":
					newState[i][j] = "."
					seatStr += "."
				}
			}
		}

		if _, ok := seatState[seatStr]; ok {
			break
		} else {
			seatState[seatStr] = true
			input = []string{}

			for _, v := range newState {
				input = append(input, strings.Join(v[:], ""))
			}
		}
	}

	result = strings.Count(seatStr, "#")

	return result
}
