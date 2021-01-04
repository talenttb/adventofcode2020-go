/*
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

// day22Cmd represents the day22 command
var day22Cmd = &cobra.Command{
	Use:   "day22",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day22 called")

		file, _ := ioutil.ReadFile("./day22input")
		d := strings.Split(string(file), "\n")

		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day22Part1(d))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day22Part2(d))
		}
	},
}

func init() {
	rootCmd.AddCommand(day22Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day22Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day22Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Day22Part1 good
func Day22Part1(input []string) int {
	result := 0

	cards := []int{}
	players := [][]int{}
	for _, v := range input {
		switch v {
		case "Player 1:", "Player 2:":
			cards = []int{}
		case "":
			players = append(players, cards[:])
		default:
			card, _ := strconv.Atoi(v)
			cards = append(cards, card)
		}
	}
	players = append(players, cards[:])

	round := 1
	p1Deck, p2Deck := players[0], players[1]

	for len(p1Deck) > 0 && len(p2Deck) > 0 {
		fmt.Printf("\n-- Round %v --\n", round)
		fmt.Printf("Player 1's deck: %#v\n", p1Deck)
		fmt.Printf("Player 2's deck: %#v\n", p2Deck)

		p1, p2 := p1Deck[0], p2Deck[0]
		p1Deck, p2Deck = p1Deck[1:], p2Deck[1:]

		fmt.Printf("Player 1 plays: %v\n", p1)
		fmt.Printf("Player 2 plays: %v\n", p2)

		if p1 > p2 {
			fmt.Printf("Player 1 wins the round!\n")
			p1Deck = append(p1Deck, []int{p1, p2}...)
		} else {
			fmt.Printf("Player 2 wins the round!\n")
			p2Deck = append(p2Deck, []int{p2, p1}...)
		}

		round++
	}

	fmt.Printf("\n-- Post-game results --\n")
	fmt.Printf("Player 1's deck: %#v\n", p1Deck)
	fmt.Printf("Player 2's deck: %#v\n", p2Deck)

	cal := []int{}
	if len(p1Deck) == 0 {
		cal = p2Deck
	} else {
		cal = p1Deck
	}

	for i, multiplied := 0, len(cal); i < len(cal); i, multiplied = i+1, multiplied-1 {
		fmt.Printf("%v * %v\n", cal[i], multiplied)
		result += cal[i] * multiplied
	}

	return result
}

// Day22Part2 good
func Day22Part2(input []string) int {
	result := 0
	return result
}
