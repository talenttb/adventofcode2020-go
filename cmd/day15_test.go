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

func TestDay15p1v0(t *testing.T) {
	d := []string{"0,3,6"}
	if got := Day15Part2(d, 7); got != 1 {
		t.Errorf("Day15Part2() = %v", got)
	}
}

func TestDay15p1v1(t *testing.T) {
	d := []string{"0,3,6"}
	if got := Day15Part2(d, 5); got != 3 {
		t.Errorf("Day15Part2() = %v", got)
	}
}
func TestDay15p1v2(t *testing.T) {
	d := []string{"0,3,6"}
	if got := Day15Part2(d, 2020); got != 436 {
		t.Errorf("Day15Part2() = %v", got)
	}
}

func TestDay15p1v3(t *testing.T) {
	d := []string{"1,3,2"}
	if got := Day15Part2(d, 2020); got != 1 {
		t.Errorf("Day15Part2() = %v", got)
	}
}
func TestDay15p1v4(t *testing.T) {
	d := []string{"1,2,3"}
	if got := Day15Part2(d, 2020); got != 27 {
		t.Errorf("Day15Part2() = %v", got)
	}
}
func TestDay15p1v5(t *testing.T) {
	d := []string{"2,3,1"}
	if got := Day15Part2(d, 2020); got != 78 {
		t.Errorf("Day15Part2() = %v", got)
	}
}
func TestDay15p1v6(t *testing.T) {
	d := []string{"3,2,1"}
	if got := Day15Part2(d, 2020); got != 438 {
		t.Errorf("Day15Part2() = %v", got)
	}
}
func TestDay15p1v7(t *testing.T) {
	d := []string{"3,1,2"}
	if got := Day15Part2(d, 2020); got != 1836 {
		t.Errorf("Day15Part2() = %v", got)
	}
}

func TestDay15p2v0(t *testing.T) {
	d := []string{"0,3,6"}
	if got := Day15Part2(d, 10); got != 0 {
		t.Errorf("Day15Part2() = %v", got)
	}
}
func TestDay15p2v1(t *testing.T) {
	d := []string{"0,3,6"}
	if got := Day15Part2(d, 30000000); got != 175594 {
		t.Errorf("Day15Part2() = %v", got)
	}
}

func TestDay15p2v2(t *testing.T) {
	d := []string{"1,3,2"}
	if got := Day15Part2(d, 30000000); got != 2578 {
		t.Errorf("Day15Part2() = %v", got)
	}
}
func TestDay15p2v3(t *testing.T) {
	d := []string{"2,1,3"}
	if got := Day15Part2(d, 30000000); got != 3544142 {
		t.Errorf("Day15Part2() = %v", got)
	}
}
func TestDay15p2v4(t *testing.T) {
	d := []string{"1,2,3"}
	if got := Day15Part2(d, 30000000); got != 261214 {
		t.Errorf("Day15Part2() = %v", got)
	}
}
func TestDay15p2v5(t *testing.T) {
	d := []string{"2,3,1"}
	if got := Day15Part2(d, 30000000); got != 6895259 {
		t.Errorf("Day15Part2() = %v", got)
	}
}
func TestDay15p2v6(t *testing.T) {
	d := []string{"3,2,1"}
	if got := Day15Part2(d, 30000000); got != 18 {
		t.Errorf("Day15Part2() = %v", got)
	}
}
func TestDay15p2v7(t *testing.T) {
	d := []string{"3,1,2"}
	if got := Day15Part2(d, 30000000); got != 362 {
		t.Errorf("Day15Part2() = %v", got)
	}
}
