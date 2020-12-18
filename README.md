# Solution of advent of code 2020

## How to use

### cmd
```bash
go run main.go -pX dayXX
```

### new day cheat sheet
```bash
cobra add dayXX
```

#### on the cmd/dayXX.go in Run func
```go
file, _ := ioutil.ReadFile("./dayXXinput")
d := strings.Split(string(file), "\n")

switch part {
case 1:
    fmt.Printf("Reuslt: %v\n", DayXXPart1(d))
case 2:
    fmt.Printf("Reuslt: %v\n", DayXXPart2(d))
}
```

#### on the cmd/dayXX.go
```go
// DayXXPart1 good
func DayXXPart1(input []string) int {
	result := 0
	return result
}
// DayXXPart2 good
func DayXXPart2(input []string) int {
	result := 0
	return result
}
```

#### Replace XX->some day

#### create test file
```
touch cmd/dayXX_test.go
```

#### content of test file
```go
package cmd

import "testing"

func TestDayXXp1(t *testing.T) {
	d := []string{
	}
	if got := DayXXPart1(d); got != 0  {
		t.Errorf("DayXXPart1() = %v", got)
	}
}

func TestDayXXp2(t *testing.T) {
	d := []string{
	}
	if got := DayXXPart2(d); got != 0  {
		t.Errorf("DayXXPart2() = %v", got)
	}
}
```