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
	"testing"
)

func TestDay13p1(t *testing.T) {
	d := []string{
		"939",
		"7,13,x,x,59,x,31,19",
	}
	if got := Day13Part1(d); got != 295 {
		t.Errorf("Day13Part1() = %v", got)
	}
}

func TestDay13p2v1(t *testing.T) {
	d := []string{"", "17,x,13,19"}
	if got := Day13Part2(d); got != 3417 {
		t.Errorf("Day13Part2() = %v", got)
	}
}
func TestDay13p2v2(t *testing.T) {
	d := []string{"", "67,7,59,61"}
	if got := Day13Part2(d); got != 754018 {
		t.Errorf("Day13Part2() = %v", got)
	}
}
func TestDay13p2v3(t *testing.T) {
	d := []string{"", "67,x,7,59,61"}
	if got := Day13Part2(d); got != 779210 {
		t.Errorf("Day13Part2() = %v", got)
	}
}
func TestDay13p2v4(t *testing.T) {
	d := []string{"", "67,7,x,59,61"}
	if got := Day13Part2(d); got != 1261476 {
		t.Errorf("Day13Part2() = %v", got)
	}
}

func TestDay13p2v5(t *testing.T) {
	d := []string{"", "1789,37,47,1889"}
	if got := Day13Part2(d); got != 1202161486 {
		t.Errorf("Day13Part2() = %v", got)
	}
}

func TestDay13Part1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day13Part1(tt.args.input); got != tt.want {
				t.Errorf("Day13Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay13Part2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day13Part2(tt.args.input); got != tt.want {
				t.Errorf("Day13Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
