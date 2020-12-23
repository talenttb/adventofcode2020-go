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

// day19Cmd represents the day19 command
var day19Cmd = &cobra.Command{
	Use:   "day19",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day19 called")
		file, _ := ioutil.ReadFile("./day19input")
		d := strings.Split(string(file), "\n")

		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day19Part1(d))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day19Part2(d))
		}
	},
}

func init() {
	rootCmd.AddCommand(day19Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day19Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day19Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day19Part1 good
func Day19Part1(input []string) int {
	result := 0
	ruleMap := make(map[string]string)
	wordList := []string{}
	section := "rules"

	for _, v := range input {
		if v == "" {
			section = "words"
			continue
		}
		switch section {
		case "rules":
			item := strings.Split(v, ": ")
			ruleMap[item[0]] = strings.Trim(item[1], "\"")
		case "words":
			wordList = append(wordList, v)
		}
	}
	r := unfold(ruleMap, "0")
	// fmt.Printf("%#v", r)

	validWords := make(map[string]bool)
	for _, v := range r {
		validWords[v] = true
	}

	for _, v := range wordList {
		if _, ok := validWords[v]; ok {
			result++
		}
	}

	return result
}

func combinate(a, b []string) []string {
	result := []string{}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			result = append(result, a[i]+b[j])
		}
	}
	return result
}

func unfold(rule map[string]string, key string) []string {
	stackL := []string{}
	stackR := []string{}

	if i := strings.Index(rule[key], "|"); i > 0 {
		sp := strings.Split(rule[key], " | ")

		stackL = append(stackL, strings.Split(sp[0], " ")...)

		stackR = append(stackR, strings.Split(sp[1], " ")...)
		l := loopStack(rule, stackL)
		r := loopStack(rule, stackR)
		return append(l, r...)
	}

	stackL = append(stackL, strings.Split(rule[key], " ")...)
	r := loopStack(rule, stackL)
	return r
}

func loopStack(rule map[string]string, s []string) []string {
	tmp := [][]string{}
	for len(s) > 0 {
		ele := s[0]
		s = s[1:]
		if ele == "a" || ele == "b" {
			tmp = append(tmp, []string{ele})
			continue
		}
		if len(tmp) == 2 {
			t := combinate(tmp[0], tmp[1])
			tmp = [][]string{}
			tmp = append(tmp, t)
		}
		tmp = append(tmp, unfold(rule, ele))
	}
	if len(tmp) == 2 {
		return combinate(tmp[0], tmp[1])
	}

	return tmp[0]
}

// Day19Part2 good
func Day19Part2(input []string) int {
	result := 0
	ruleMap := make(map[string]string)
	wordList := []string{}
	section := "rules"
	for _, v := range input {
		if v == "" {
			section = "words"
			continue
		}
		switch section {
		case "rules":
			item := strings.Split(v, ": ")
			ruleMap[item[0]] = strings.Trim(item[1], "\"")
		case "words":
			wordList = append(wordList, v)
		}
	}
	r42 := unfold(ruleMap, "42")
	r31 := unfold(ruleMap, "31")
	hardCodeLen := len(r42[0])
	// fmt.Printf("%#v\n", r)
	validWords42 := make(map[string]bool)
	for _, v := range r42 {
		validWords42[v] = true
	}

	validWords31 := make(map[string]bool)
	for _, v := range r31 {
		validWords31[v] = true
	}

	newWordList := []string{}
	for _, v := range wordList {
		if len(v)%hardCodeLen != 0 {
			continue
		}
		newWord := ""
		for i := 0; i < len(v); i += hardCodeLen {
			if _, ok := validWords42[v[i:i+hardCodeLen]]; ok {
				newWord += "a"
				continue
			}
			if _, ok := validWords31[v[i:i+hardCodeLen]]; ok {
				newWord += "b"
				continue
			}
		}
		newWordList = append(newWordList, newWord)
	}

	re := regexp.MustCompile(`^a+b+$`)
	for _, v := range newWordList {
		if re.MatchString(v) {
			// fmt.Printf("%s\n", v)
			countA := strings.Count(v, "a")
			countB := strings.Count(v, "b")
			if countA > countB {
				result++
			}
		}
	}

	return result
}
