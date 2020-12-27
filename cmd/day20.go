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
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day20Cmd represents the day20 command
var day20Cmd = &cobra.Command{
	Use:   "day20",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day20 called")
		file, _ := ioutil.ReadFile("./day20input")
		d := strings.Split(string(file), "\n")

		switch part {
		case 1:
			fmt.Printf("Reuslt: %v\n", Day20Part1(d))
		case 2:
			fmt.Printf("Reuslt: %v\n", Day20Part2(d))
		}
	},
}

func init() {
	rootCmd.AddCommand(day20Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day20Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day20Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type tile struct {
	id         string
	t, b, l, r *tile
	image      [][]string
}

func (t tile) display() {
	fmt.Printf("id %v:\n", t.id)

	for i := 0; i < 4; i++ {
		fmt.Printf("before(%v):\n", i)
		fmt.Printf("============\n")
		for x := 0; x < N; x++ {
			fmt.Printf("%#v\n", t.image[x])
		}
		fmt.Printf("============\n")
		t.rotateImage()
		fmt.Printf("after(%v):\n", i)
		fmt.Printf("============\n")
		for x := 0; x < N; x++ {
			fmt.Printf("%#v\n", t.image[x])
		}
		fmt.Printf("============\n")
	}
}

func (t tile) getEdge(direct string) []string {
	result := []string{}
	switch direct {
	case "top":
		result = t.image[0]
	case "bottom":
		result = t.image[N-1]
	case "left":
		for j := 0; j < N; j++ {
			result = append(result, t.image[j][0])
		}
	case "right":
		for j := 0; j < N; j++ {
			result = append(result, t.image[j][N-1])
		}
	default:
		panic("error direct")
	}
	return result
}

func (t *tile) flip(direct int) {
	switch direct {
	case 0:
		// top bottom
		mat := t.image
		// fmt.Printf("%v before:\n", t.id)
		// for i := 0; i < N; i++ {
		// 	fmt.Printf("%#v\n", mat[i])
		// }
		for i := 0; i < N/2; i++ {
			mat[N-i-1], mat[i] = mat[i], mat[N-i-1]
		}
		// fmt.Printf("after:\n")
		// for i := 0; i < N; i++ {
		// 	fmt.Printf("%#v\n", mat[i])
		// }
	case 1:
		// left right
		mat := t.image
		for i := 0; i < N; i++ {
			for j := 0; j < N/2; j++ {
				mat[i][N-j-1], mat[i][j] = mat[i][j], mat[i][N-j-1]
			}
		}
		// for i := 0; i < N; i++ {
		// 	fmt.Printf("%#v\n", mat[i])
		// }
	}
}

func (t *tile) rotateImage(timesOption ...int) {
	times := 1
	if len(timesOption) > 0 {
		times = timesOption[0]
	}

	for i := 0; i < times; i++ {
		mat := t.image
		for x := 0; x < N/2; x++ {
			for y := x; y < N-x-1; y++ {
				temp := mat[x][y]
				mat[x][y] = mat[y][N-1-x]
				mat[y][N-1-x] = mat[N-1-x][N-1-y]
				mat[N-1-x][N-1-y] = mat[N-1-y][x]
				mat[N-1-y][x] = temp
			}
		}
		// t.image = result
	}
}

// N is matrix size
const N int = 10

// Day20Part1 good
func Day20Part1(input []string) int {
	result := 1
	stack := []*tile{}
	for i := 0; i < len(input); i += 12 {
		t := &tile{}
		id := strings.Trim(strings.Trim(input[i], "Tile "), ":")
		t.id = id
		for j := 1; j <= 10; j++ {
			t.image = append(t.image, strings.Split(input[i+j], ""))
		}
		stack = append(stack, t)
	}

	done := []*tile{stack[0]}
	stack = stack[1:]
	stop := 0
	for len(stack) > 0 {
		ele := stack[0]
		stack = stack[1:]
		checkAll := false
		for i := 0; i < len(done); i++ {
			//check done and stack
			if checkEdge(done[i], ele) {
				done = append(done, ele)
				break
			}

			if i == len(done)-1 {
				checkAll = true
			}
		}
		if checkAll {
			stack = append(stack, ele)
		}
		stop++
		if stop++; stop > 5000 {
			panic("fail......")
		}
	}

	for i := range done {
		isCorner := 0
		fmt.Printf("\nID %v :", done[i].id)
		fmt.Printf("\ntop ->")
		if done[i].t != nil {
			fmt.Printf("%v", done[i].t.id)
			isCorner++
		}
		fmt.Printf("\nbottom ->")
		if done[i].b != nil {
			fmt.Printf("%v", done[i].b.id)
			isCorner++
		}
		fmt.Printf("\nleft ->")
		if done[i].l != nil {
			fmt.Printf("%v", done[i].l.id)
			isCorner++
		}
		fmt.Printf("\nright ->")
		if done[i].r != nil {
			fmt.Printf("%v", done[i].r.id)
			isCorner++
		}

		fmt.Printf("\n")
		if isCorner == 2 {
			v, _ := strconv.Atoi(done[i].id)
			// fmt.Printf("%v\n", v)
			result *= v
		}
	}

	return result
}

func deepEqual(a, b []string) bool {
	// fmt.Printf("==\ndeepEqual:\n%#v\n%#v\n==\n", a, b)
	return reflect.DeepEqual(a, b)
}

func checkEdge(done, check *tile) bool {
	for i := 0; i < 2; i++ {
		// check.display()
		for j := 0; j < 4; j++ {
			checkEdge := []string{}
			doneEdge := []string{}

			//t
			// rotate 2 times
			doneEdge = done.getEdge("top")
			checkEdge = check.getEdge("bottom")
			// fmt.Printf("%v ->\n%#v\n", j, checkEdge)
			if done.t == nil && deepEqual(checkEdge, doneEdge) {
				// fmt.Printf("%v ->\n%#v\n--\n%#v\n", j, checkEdge, doneEdge)
				isDouble := false

				if done.l != nil && done.l.t != nil {
					if deepEqual(check.getEdge("left"), done.l.t.getEdge("right")) {
						isDouble = true
						check.b = done
						done.t = check
						check.l = done.l.t
						done.l.t.r = check

						next := done.l.t
						if next.t != nil && next.t.r != nil {
							if deepEqual(check.getEdge("top"), next.t.r.getEdge("bottom")) {
								check.t = next.t.r
								next.t.r.b = check
							} else {
								fmt.Printf("err1-1")
							}
						}
					} else {
						fmt.Printf("err1")
					}
				}

				if done.r != nil && done.r.t != nil {
					if deepEqual(check.getEdge("right"), done.r.t.getEdge("left")) {
						isDouble = true
						check.b = done
						done.t = check
						check.r = done.r.t
						done.r.t.l = check

						next := done.r.t
						if next.t != nil && next.t.l != nil {
							if deepEqual(check.getEdge("top"), next.t.l.getEdge("bottom")) {
								check.t = next.t.l
								next.t.l.b = check
							} else {
								fmt.Printf("err2-2")
							}
						}
					} else {
						fmt.Printf("err2")
					}
				}

				if !isDouble {
					check.b = done
					done.t = check
				}
				return true
			}

			//l
			// rotate 3 times
			doneEdge = done.getEdge("left")
			checkEdge = check.getEdge("right")
			if done.l == nil && deepEqual(checkEdge, doneEdge) {
				isDouble := false
				// check.rotateImage(3)

				if done.t != nil && done.t.l != nil {
					if deepEqual(check.getEdge("top"), done.t.l.getEdge("bottom")) {
						isDouble = true
						check.r = done
						done.l = check
						check.t = done.t.l
						done.t.l.b = check

						next := done.t.l
						if next.l != nil && next.l.b != nil {
							if deepEqual(check.getEdge("left"), next.l.b.getEdge("right")) {
								check.l = next.l.b
								next.l.b.r = check
							} else {
								fmt.Printf("err1-1")
							}
						}
					} else {
						fmt.Printf("err3")
					}

				}

				if done.b != nil && done.b.l != nil {
					if deepEqual(check.getEdge("bottom"), done.b.l.getEdge("top")) {
						isDouble = true
						check.r = done
						done.l = check
						check.b = done.b.l
						done.b.l.t = check

						next := done.b.l
						if next.l != nil && next.l.t != nil {
							if deepEqual(check.getEdge("left"), next.l.t.getEdge("right")) {
								check.l = next.l.t
								next.l.t.r = check
							} else {
								fmt.Printf("err4-1")
							}
						}
					} else {
						fmt.Printf("err4")
					}
				}

				if !isDouble {
					check.r = done
					done.l = check
				}
				return true
			}

			//b
			// rotate 0 times
			doneEdge = done.getEdge("bottom")
			checkEdge = check.getEdge("top")
			if done.b == nil && deepEqual(checkEdge, doneEdge) {
				// check.rotateImage()
				isDouble := false

				if done.l != nil && done.l.b != nil {
					if deepEqual(check.getEdge("left"), done.l.b.getEdge("right")) {
						isDouble = true
						check.t = done
						done.b = check
						check.l = done.l.b
						done.l.b.r = check

						next := done.l.b
						if next.b != nil && next.b.r != nil {
							if deepEqual(check.getEdge("bottom"), next.b.r.getEdge("top")) {
								check.b = next.b.r
								next.b.r.t = check
							} else {
								fmt.Printf("err5-1")
							}
						}
					} else {
						fmt.Printf("err5")
					}
				}

				if done.r != nil && done.r.b != nil {
					if deepEqual(check.getEdge("right"), done.r.b.getEdge("left")) {
						isDouble = true
						check.t = done
						done.b = check
						check.r = done.r.b
						done.r.b.l = check

						next := done.r.b
						if next.b != nil && next.b.l != nil {
							if deepEqual(check.getEdge("bottom"), next.b.l.getEdge("top")) {
								check.b = next.b.l
								next.b.l.t = check
							} else {
								fmt.Printf("err6-1")
							}
						}
					} else {
						fmt.Printf("err6")
					}
				}
				if !isDouble {
					check.t = done
					done.b = check
				}
				return true
			}

			//r
			// rotate 1 times
			doneEdge = done.getEdge("right")
			checkEdge = check.getEdge("left")
			if done.r == nil && deepEqual(checkEdge, doneEdge) {
				isDouble := false
				// check.rotateImage(1)
				check.l = done
				done.r = check
				if done.t != nil && done.t.r != nil {
					if deepEqual(check.getEdge("top"), done.t.r.getEdge("bottom")) {
						isDouble = true
						check.l = done
						done.r = check
						check.t = done.t.r
						done.t.r.b = check

						next := done.t.r
						if next.r != nil && next.r.b != nil {
							if deepEqual(check.getEdge("right"), next.r.b.getEdge("left")) {
								check.r = next.r.b
								next.r.b.l = check
							} else {
								fmt.Printf("err7-1")
							}
						}
					} else {
						fmt.Printf("err7")
					}
				}

				if done.b != nil && done.b.r != nil {
					if deepEqual(check.getEdge("bottom"), done.b.r.getEdge("top")) {
						isDouble = true
						check.l = done
						done.r = check
						check.b = done.b.r
						done.b.r.t = check

						next := done.b.r
						if next.r != nil && next.r.t != nil {
							if deepEqual(check.getEdge("right"), next.r.t.getEdge("left")) {
								check.r = next.r.t
								next.r.t.l = check
							} else {
								fmt.Printf("err8-1")
							}
						}
					} else {
						fmt.Printf("err8")
					}
				}

				if !isDouble {
					check.l = done
					done.r = check
				}
				return true
			}

			check.rotateImage()
		}
		// fmt.Printf("flip %v\n", i)
		check.flip(i)
	}

	return false
}

// for debug
// fmt.Printf("checker %v:\n", check.id)
// 						fmt.Printf("image:\n")
// 						for i := 0; i < N; i++ {
// 							fmt.Printf("%#v\n", check.image[i])
// 						}
// 						fmt.Printf("top :\n%#v\n", check.getEdge("top"))
// 						fmt.Printf("bottom :\n%#v\n", check.getEdge("bottom"))
// 						fmt.Printf("right :\n%#v\n", check.getEdge("right"))
// 						fmt.Printf("left :\n%#v\n", check.getEdge("left"))

// 						fmt.Printf("done %v:\n", done.id)
// 						fmt.Printf("image:\n")
// 						for i := 0; i < N; i++ {
// 							fmt.Printf("%#v\n", done.image[i])
// 						}
// 						fmt.Printf("top :\n%#v\n", done.getEdge("top"))
// 						fmt.Printf("bottom :\n%#v\n", done.getEdge("bottom"))
// 						fmt.Printf("right :\n%#v\n", done.getEdge("right"))
// 						fmt.Printf("left :\n%#v\n", done.getEdge("left"))

// 						fmt.Printf("done.r.t %v:\n", done.r.t.id)
// 						fmt.Printf("image:\n")
// 						for i := 0; i < N; i++ {
// 							fmt.Printf("%#v\n", done.r.t.image[i])
// 						}
// 						fmt.Printf("top :\n%#v\n", done.r.t.getEdge("top"))
// 						fmt.Printf("bottom :\n%#v\n", done.r.t.getEdge("bottom"))
// 						fmt.Printf("right :\n%#v\n", done.r.t.getEdge("right"))
// 						fmt.Printf("left :\n%#v\n", done.r.t.getEdge("left"))

// Day20Part2 good
func Day20Part2(input []string) int {
	result := 0
	return result
}
