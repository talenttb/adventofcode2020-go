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

// day7Cmd represents the day7 command
var day7Cmd = &cobra.Command{
	Use:   "day7",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day7 called")
		file, _ := ioutil.ReadFile("./day7input")
		d := strings.Split(string(file), "\n")

		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day7Part1(d))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day7Part2(d))
		}
	},
}

func init() {
	rootCmd.AddCommand(day7Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day7Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day7Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day7Part1 good
func Day7Part1(input []string) int {
	roadMap := make(map[string][]string)

	pRE := regexp.MustCompile(`(?P<Parent>[\w\s]+) bags contain`)
	cRE := regexp.MustCompile(`(\d\s(?P<Children>[\w\s]+) bag)`)

	for _, s := range input {

		pArr := pRE.FindStringSubmatch(s)
		// fmt.Printf("%#v\n", pArr[len(pArr)-1])

		for _, c := range cRE.FindAllStringSubmatch(s, -1) {
			childKey := c[len(c)-1]
			roadMap[childKey] = append(roadMap[childKey], pArr[len(pArr)-1])
			// fmt.Printf("%#v\n", childKey)
		}
	}

	target := "shiny gold"
	path := []string{}
	path = append(path, roadMap[target]...)
	r := make(map[string]bool)

	for len(path) > 0 {
		ele := path[0]
		path = path[1:]

		// ele, path := path[0], path[1:]
		r[ele] = true

		if v, ok := roadMap[ele]; ok {
			path = append(path, v...)
		}
	}
	return len(r)
}

type node struct {
	name  string
	value int
}

// Day7Part2 good
func Day7Part2(input []string) int {
	result := 0
	roadMap := make(map[string][]node)

	pRE := regexp.MustCompile(`(?P<Parent>[\w\s]+) bags contain`)
	cRE := regexp.MustCompile(`((\d)\s(?P<Children>[\w\s]+) bag)`)

	for _, s := range input {

		pArr := pRE.FindStringSubmatch(s)
		// fmt.Printf("%#v\n", pArr[len(pArr)-1])
		roadMapKey := pArr[len(pArr)-1]
		for _, c := range cRE.FindAllStringSubmatch(s, -1) {
			i, _ := strconv.Atoi(c[len(c)-2])
			roadMap[roadMapKey] = append(roadMap[roadMapKey], node{name: c[len(c)-1], value: i})
		}
	}

	target := "shiny gold"
	result = myGet(roadMap, target)
	// path := []node{}
	// path = append(path, roadMap[target]...)
	// r := make(map[string]bool)

	// for len(path) > 0 {
	// 	firstNode := path[0]
	// 	path = path[1:]
	// 	result += firstNode.value * myGet(roadMap, firstNode.name)
	// r[ele] = true

	// if v, ok := roadMap[ele]; ok {
	// 	path = append(path, v...)
	// }
	// }
	// result = len(r)
	return result
}

func myGet(r map[string][]node, k string) int {
	result := 0
	for _, v := range r[k] {
		result += v.value + v.value*myGet(r, v.name)
	}
	// if v, ok := r[k]; ok {
	// 	for _, item := range v {

	// 	}
	// }
	return result
}
