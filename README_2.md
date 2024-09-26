## Content
1.[main.go](#main.go)<br>
2.[cycles.go](#cycles.go)

### main.go
Description for the code of `main.go`
    
```go
    ackage main

    import "fmt"

    func main() {
	var age int8 = 24
	fmt.Println(age)
    }
```
---
### cycles.go
Description for `cycles.go` file
``` go
    	nums := []int{1, 2, 3, 4, 5}

	for index, element := range nums {
		fmt.Printf("Index: %d element: %d\n", index, element)
	}

	for _, element := range nums {
		fmt.Printf("Element: %d\n", element)
	}
```
---