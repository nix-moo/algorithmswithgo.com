# Syntax   

- A helpful reference: [Ofiicial tour of go](https://go.dev/tour)

## loop
```go
func LoopExamples(){
	// go only has for loops
	// 	:= only works INSIDE a func
	
	for init:= 0; init < condition; post{
		// for loop takes 3 args but only the condition is required
		// leaving out the other two args creates a pseudo while loop
	}
	for i := range list {
		// i == index of array
	}
	
	for _, i := range list {
		// i == list[i] aka actual value
	}
	
	for{
		// omit all args for infinite loop
	}
}
```

## if
```go
func IfExamples(){
	// standard
	if booleanExpression {
		//then
	}
	if newStatement:= 0; newStatement == boolean {
		// you can add a statement to run before the expression evaluates
		// newStatement scope is only in if block
	} else {
		// newStatement works in any else blocks
	}
}
```

# Patterns

## Templates

### Two pointers
- One input opposite ends 
```go
func fn(a []int) int{
	ans := 0
	left :=0
	right := len(a)-1

	for left < right {
		// SOME LOGIC HERE
		if CONDITION {
			left++
		} else {
			right--
		}
	}
	return ans
}
```

- Two inputs, exhaust both
```go
func fn(a, b []int) int {
	ans, i, j := 0
	for i <  len(a) && j < len(b){
		// SOME LOGIC HERE
		if CONDITION {
			i++
		} else {
			j++
		}
	}
	for i < len(a) {
		// SOME LOGIC HERE
		i++
	}
	for j < len(b) {
		// SOME LOGIC HERE
		j++
	}
}
```

### Sliding window
```go
func fn(a []int) int {
	left, curr, ans := 0
	for right:=0; right < len(a); right++ {
		// do logic here to add arr[right] to curr
		for (WINDOW_CONDITION_FALSE) {
			left++
		}
		//  LOGIC TO UPDATE ans GOES HERE
	}
	return ans
}
```

### Efficient string building
```go
import "strings"

func fn(w []string) string {
	var built strings.Builder
	for _, i := range w {
		built.WriteString(i)
	}
	return built.String()
}
```

Ø
## Test Edge Cases
- [ ] Empty list
- [ ] Missing first or last
- [ ] Only one element
- [ ] Only two elements
- [ ] 














