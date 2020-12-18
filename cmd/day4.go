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
	"strings"

	"github.com/spf13/cobra"
)

// day4Cmd represents the day4 command
var day4Cmd = &cobra.Command{
	Use:   "day4",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day4 called")

		file, _ := ioutil.ReadFile("./cmd/day4input")
		a := strings.Split(string(file), "\n")

		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day4Part1(a))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day4Part2(a))
		}
	},
}

func init() {
	rootCmd.AddCommand(day4Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day4Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day4Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	day4Cmd.Flags().IntVarP(&part, "part", "p", 1, "")
}

// Day4Part1 good
func Day4Part1(arr []string) int {
	result := 0
	var passport []string
	for _, v := range arr {

		if v != "" {
			passport = append(passport, strings.Split(v, " ")...)
			continue
		}
		check := checkPp(passport)
		passport = []string{}
		if check {
			result++
		}
	}

	check := checkPp(passport)
	if check {
		result++
	}

	return result
}

func checkPp(passport []string) bool {
	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	pp := make(map[string]string)
	for _, i := range passport {
		kv := strings.Split(i, ":")
		pp[kv[0]] = kv[1]
	}

	check := true
	for _, j := range requiredFields {
		_, ok := pp[j]
		check = check && ok
	}
	return check
}

// Day4Part2 good
func Day4Part2(arr []string) int {
	result := 0
	var passport []string
	for _, v := range arr {

		if v != "" {
			passport = append(passport, strings.Split(v, " ")...)
			continue
		}
		check := checkPpv2(passport)
		passport = []string{}
		if check {
			result++
		}
	}

	check := checkPpv2(passport)
	if check {
		result++
	}

	return result
}

func checkPpv2(passport []string) bool {
	requiredFields := map[string]string{
		"byr": "^(19[2-9][0-9]|200[0-2])$",
		"iyr": "^(201[0-9]|2020)$",
		"eyr": "^(202[0-9]|2030)$",
		"hgt": "^((19[0-3]|1[5-8][0-9])cm)|((59|6[0-9]|7[0-6])in)$",
		"hcl": "^#([a-fA-F0-9]{6})$",
		"ecl": "^(amb|blu|brn|gry|grn|hzl|oth)$",
		// "pid": "^([0-9]{9})$",
		"pid": "^([\\d]{9})$",
	}

	// matched, err := regexp.MatchString(`foo.*`, "seafood")

	pp := make(map[string]string)
	for _, i := range passport {
		kv := strings.Split(i, ":")
		pp[kv[0]] = kv[1]
	}

	check := true
	for k, v := range requiredFields {
		value, ok := pp[k]
		if ok {
			r, _ := regexp.Compile(v)
			b := r.MatchString(value)
			check = check && b
		}
		check = check && ok

	}
	return check
}
