
## Create Functions

- Functions allow you to group together a set of statements that you can call from other parts of your program.
- A function can take zero or more arguments.
- When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
- Functions can return one or more values.
- Functions can be used to create modular and reusable code.

### Main Function
- It is the entry point of every executable program.
- There can only be one `main()` function in your program, but if you are creating a Go package (a library), you don't need one.
- You can use `os.Args` to interact with values passed when running the program.

### Custom Functions
```go
func name(parameters) (results) {
    body-content
}

```

-   Custom functions can have different types of parameters and return types.
    
-   You can use named return values to improve code readability.
    

### Function Types

-   **Named Return Values**: Functions can have named return values, which act as variables defined at the top of the function.
    
    go
    
    ```
    func split(sum int) (x, y int) {
        x = sum * 4 / 9
        y = sum - x
        return
    }
    
    ```
    
-   **Variadic Functions**: Functions that can take a variable number of arguments.
    
    go
    
    ```
    func sum(nums ...int) int {
        total := 0
        for _, num := range nums {
            total += num
        }
        return total
    }
    fmt.Println(sum(1, 2, 3, 4)) // Output: 10
    
    ```
    
-   **Anonymous Functions**: Functions without a name, often used for short-term use.
    
    go
    
    ```
    func() {
        fmt.Println("Anonymous function called")
    }()
    
    ```
    
-   **First-Class Functions**: Functions can be passed as arguments to other functions, returned as values from functions, and assigned to variables.
    
    go
    
    ```
    func compute(fn func(int, int) int) int {
        return fn(3, 4)
    }
    func add(x, y int) int { return x + y }
    fmt.Println(compute(add)) // Output: 7
    
    ```
    

### Recursive Functions

-   Functions that call themselves to solve smaller instances of the same problem.
    
    go
    
    ```
    func factorial(n int) int {
        if n == 0 {
            return 1
        }
        return n * factorial(n-1)
    }
    fmt.Println(factorial(5)) // Output: 120
    
    ```
    

### Defer, Panic, and Recover

-   **Defer**: Schedules a function call to be run after the function completes.
    
    go
    
    ```
    func main() {
        defer fmt.Println("This will be printed last")
        fmt.Println("Hello")
    }
    // Output:
    // Hello
    // This will be printed last
    
    ```
    
-   **Panic**: Causes a run-time error.
    
    go
    
    ```
    func causePanic() {
        panic("A severe error occurred!")
    }
    
    ```
    
-   **Recover**: Allows you to regain control of a panicking goroutine.
    
    go
    
    ```
    func main() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Println("Recovered from:", r)
            }
        }()
        causePanic()
    }
    // Output:
    // Recovered from: A severe error occurred!
    ```