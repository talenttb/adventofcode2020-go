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

import "testing"

func TestDay16p1(t *testing.T) {
	d := []string{
		"class: 1-3 or 5-7",
		"row: 6-11 or 33-44",
		"seat: 13-40 or 45-50",
		"",
		"your ticket:",
		"7,1,14",
		"",
		"nearby tickets:",
		"7,3,47",
		"40,4,50",
		"55,2,20",
		"38,6,12",
	}
	if got := Day16Part1(d); got != 71 {
		t.Errorf("Day16Part1() = %v", got)
	}
}

func TestDay16p2(t *testing.T) {
	d := []string{
		"class: 0-1 or 4-19",
		"row: 0-5 or 8-19",
		"seat: 0-13 or 16-19",
		"",
		"your ticket:",
		"11,12,13",
		"",
		"nearby tickets:",
		"3,9,18",
		"15,1,5",
		"5,14,9",
	}
	if got := Day16Part2(d); got != 0 {
		t.Errorf("Day16Part2() = %v", got)
	}
}
