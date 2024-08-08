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