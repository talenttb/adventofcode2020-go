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
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day16Cmd represents the day16 command
var day16Cmd = &cobra.Command{
	Use:   "day16",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day16 called")

		file, _ := ioutil.ReadFile("./day16input")
		d := strings.Split(string(file), "\n")

		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day16Part1(d))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day16Part2(d))
		}

	},
}

func init() {
	rootCmd.AddCommand(day16Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day16Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day16Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day16Part1 good
func Day16Part1(input []string) int {
	result := 0
	validCond := make(map[int]bool)
	invalid := []int{}
	condRE := regexp.MustCompile(`\d+-\d+`)
	steps := 1
	for _, item := range input {
		cmd := strings.Split(item, ":")
		switch cmd[0] {
		case "departure location", "departure station", "departure platform", "departure track", "departure date", "departure time", "arrival location", "arrival station", "arrival platform", "arrival track", "class", "duration", "price", "route", "row", "seat", "train", "type", "wagon", "zone":
			steps = 1
		case "your ticket":
			steps = 2
		case "nearby tickets":
			steps = 3
		}
		switch steps {
		case 1:

			for _, rangeStr := range condRE.FindAllStringSubmatch(item, -1) {
				s, _ := strconv.Atoi(strings.Split(rangeStr[0], "-")[0])
				e, _ := strconv.Atoi(strings.Split(rangeStr[0], "-")[1])
				for i := s; i < e+1; i++ {
					validCond[i] = true
				}
				// fmt.Printf("%#v\n", rangeStr)
			}
			// x := strings.Split(item, "-")
			// fs := x[0][len(x[0]):]

		case 3:
			if item == "nearby tickets:" {
				break
			}

			for _, value := range strings.Split(item, ",") {
				v, _ := strconv.Atoi(value)
				if _, ok := validCond[v]; !ok {
					invalid = append(invalid, v)
				}
			}
		}
	}

	for _, item := range invalid {
		result += item
	}
	return result
}

// Day16Part2 good
// departure time
func Day16Part2(input []string) int {
	result := 1
	validCond := make(map[int][]string)
	// validValue := []int{}
	myTicket := []int{}
	validTickets := [][]int{}
	possibleFields := [][]string{}
	condRE := regexp.MustCompile(`\d+-\d+`)
	steps := 1
	for _, item := range input {
		cmd := strings.Split(item, ":")
		switch cmd[0] {
		case "departure location", "departure station", "departure platform", "departure track", "departure date", "departure time", "arrival location", "arrival station", "arrival platform", "arrival track", "class", "duration", "price", "route", "row", "seat", "train", "type", "wagon", "zone":
			steps = 1
		case "your ticket":
			steps = 2
		case "nearby tickets":
			steps = 3
		case "":
			steps = 0
		}
		switch steps {
		case 1:
			for _, rangeStr := range condRE.FindAllStringSubmatch(item, -1) {
				s, _ := strconv.Atoi(strings.Split(rangeStr[0], "-")[0])
				e, _ := strconv.Atoi(strings.Split(rangeStr[0], "-")[1])
				for i := s; i < e+1; i++ {
					validCond[i] = append(validCond[i], cmd[0])
				}
			}
		case 2:
			if item == "your ticket:" {
				break
			}

			for _, value := range strings.Split(item, ",") {
				v, _ := strconv.Atoi(value)
				myTicket = append(myTicket, v)
			}

		case 3:
			if item == "nearby tickets:" {
				break
			}

			check := true
			ticketValues := strings.Split(item, ",")
			// maze := [][]string{}
			validTicket := []int{}
			for _, value := range ticketValues {
				v, _ := strconv.Atoi(value)
				if _, ok := validCond[v]; !ok {
					check = false
					break
				}
			}

			if check {
				for _, value := range strings.Split(item, ",") {
					v, _ := strconv.Atoi(value)
					validTicket = append(validTicket, v)
				}
				validTickets = append(validTickets, validTicket)
			}
		}
	}

	for j := 0; j < len(validTickets[0]); j++ {
		possibleField := []string{}
		for i := 0; i < len(validTickets); i++ {
			v := validTickets[i][j]
			f, _ := validCond[v]
			if len(possibleField) == 0 {
				possibleField = append(possibleField, f...)
			} else {
				possibleField = intersection(possibleField, f)
			}
		}
		possibleFields = append(possibleFields, possibleField)
	}

	count := 0
	for count < len(possibleFields) {
		// fmt.Printf("%#v\n", possibleFields)
		count = 0
		for i := 0; i < len(possibleFields); i++ {
			if len(possibleFields[i]) == 1 {
				count++
				for j := 0; j < len(possibleFields); j++ {
					if i != j {
						idx := -1
						for ii, v := range possibleFields[j] {
							if v == possibleFields[i][0] {
								idx = ii
								break
							}
						}
						if idx == -1 {
							continue
						}

						if idx+1 >= len(possibleFields[j]) {
							possibleFields[j] = possibleFields[j][:idx]
						} else {
							possibleFields[j] = append(possibleFields[j][:idx], possibleFields[j][idx+1:]...)
						}
						break
					}
				}
			}
		}
	}
	// fmt.Printf("%#v\n", possibleFields)
	for i, v := range possibleFields {
		if strings.Contains(v[0], "departure") {
			result *= myTicket[i]
		}
	}

	return result
}

func intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}
