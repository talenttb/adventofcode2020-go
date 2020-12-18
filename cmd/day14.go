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
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day14Cmd represents the day14 command
var day14Cmd = &cobra.Command{
	Use:   "day14",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day14 called")

		file, _ := ioutil.ReadFile("./day14input")
		d := strings.Split(string(file), "\n")

		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day14Part1(d))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day14Part2(d))
		}

	},
}

func init() {
	rootCmd.AddCommand(day14Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day14Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day14Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day14Part1 good
func Day14Part1(input []string) int {
	result := 0
	// 	mask = 010X11000X010001X1XXX0110011X1X1110X
	// mem[43479] = 61
	// mem[47199] = 15617564
	// mem[18265] = 6027808

	// max idx: 65089
	memRE := regexp.MustCompile(`mem\[(.*)\]\s=\s(\d+)`)
	mask := ""
	memMap := make(map[int]int)

	for i := 0; i < len(input); i++ {
		cmd := string(input[i][:3])
		switch cmd {
		case "mas":
			mask = strings.Split(input[i], "= ")[1]
			// fmt.Printf("mask: %s\n", mask)
			break
		case "mem":
			filterValue := memRE.FindStringSubmatch(input[i])
			idx, _ := strconv.Atoi(filterValue[1])
			value, _ := strconv.Atoi(filterValue[2])
			memMap[idx] = decoderV1(mask, value)
			// fmt.Printf("mem: %v,%v\n", idx, value)
		}
	}

	for _, v := range memMap {
		result += v
	}

	return result
}

func decoderV1(mask string, value int) int {
	maskLen := len(mask)
	bValue := fmt.Sprintf("%0*b", maskLen, value)
	xx := [36]float64{}
	var result float64 = 0
	for i := maskLen - 1; i >= 0; i-- {
		positionOf2 := maskLen - i - 1
		switch mask[i] {
		case 'X':
			v, _ := strconv.ParseFloat(string(bValue[i]), 64)
			result += v * math.Pow(2, float64(positionOf2))
			xx[i] = v
		case '0':
			xx[i] = 0
			// result += float64(bValue[i]) * math.Pow(2, float64(i))
			break
		case '1':
			xx[i] = 1
			result += float64(1) * math.Pow(2, float64(positionOf2))
		}
	}
	return int(result)
}

// Day14Part2 good
func Day14Part2(input []string) int {
	result := 0
	memRE := regexp.MustCompile(`mem\[(.*)\]\s=\s(\d+)`)
	mask := ""
	memMap := make(map[int]int)

	for i := 0; i < len(input); i++ {
		cmd := string(input[i][:3])
		switch cmd {
		case "mas":
			mask = strings.Split(input[i], "= ")[1]
			break
		case "mem":
			filterValue := memRE.FindStringSubmatch(input[i])
			idx, _ := strconv.Atoi(filterValue[1])
			value, _ := strconv.Atoi(filterValue[2])
			combIdx := decoderV2(mask, idx)
			iter := float64(strings.Count(combIdx, "X"))
			maxIter := int(math.Pow(2, iter)) - 1
			leftPad := fmt.Sprintf("%b", maxIter)
			for i := 0; i <= maxIter; i++ {
				cond := fmt.Sprintf("%0*b", len(leftPad), i)
				combIdxSub := combIdx
				for j := 0; j < len(cond); j++ {
					index := strings.Index(combIdxSub, "X")
					combIdxSub = combIdxSub[:index] + string(cond[j]) + combIdxSub[index+1:]
				}
				memMap[int2ToInt10(combIdxSub)] = value
			}
		}
	}

	for _, v := range memMap {
		result += v
	}

	return result
}

func int2ToInt10(b string) int {
	var result float64 = 0
	maskLen := len(b)
	for i := maskLen - 1; i >= 0; i-- {
		positionOf2 := maskLen - i - 1
		v, _ := strconv.ParseFloat(string(b[i]), 64)
		result += v * math.Pow(2, float64(positionOf2))
	}
	return int(result)
}

func decoderV2(mask string, value int) string {
	maskLen := len(mask)
	bValue := fmt.Sprintf("%0*b", maskLen, value)
	result := ""
	for i := 0; i < maskLen; i++ {
		switch mask[i] {
		case 'X':
			result += "X"
		case '0':
			result += string(bValue[i])
		case '1':
			result += "1"
		}
	}
	return result
}
