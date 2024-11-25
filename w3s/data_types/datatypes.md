# Data Types in Go

Go variables are associated with specific data types, which dictate what kind of data they can hold. These types can be categorized into four main categories:

### 1. Basic Types:
   - **Numbers**: Integer or floating-point numbers.
   - **Strings**: Text enclosed in double quotes.
   - **Booleans**: Represent `true` or `false` values.

### 2. Aggregate Types:
   - **Arrays**: A collection of elements of the same type.
   - **Structs**: A custom data type that groups related data.

### 3. Reference Types:
   - **Pointers**: A reference to a memory address.
   - **Slices**: A dynamically-sized array.
   - **Maps**: Key-value pairs (similar to dictionaries in Python).
   - **Functions**: Can be used as types.
   - **Channels**: Used for communication between goroutines.

### 4. Interface Type:
   - Defines a set of methods but doesn't implement them.

## Integers (`int`)
Go offers several integer types, which can be divided into signed and unsigned integers:

### Signed Integers
These types store both negative and positive values, and they include:
- `int8`, `int16`, `int32`, `int64`, and `int`.
- The range is split between negative and positive numbers, with the maximum positive value being half of the unsigned integer's maximum value.

### Unsigned Integers
- Represent only non-negative values (positive numbers and zero).
- Types include `uint8`, `uint16`, `uint32`, `uint64`, and `uint`.

> **Note:** Use `unit` only when you need to represent a value as an unsigned number for a certain reason. 

### Rune
- A **rune** is an alias for the `int32` data type.
- It is used to represent Unicode characters or Unicode code points.

## Floating-Point Numbers
Go supports two floating-point types:
- `float32`: Single-precision floating-point number.
- `float64`: Double-precision floating-point number.

These types are used when you need to store large numbers that don't fit in the previously defined integer types.

## Type Conversion in Go
Type conversion is the process of converting one data type into another. Unlike some other languages, Go requires **explicit type conversion**.

### Implicit vs Explicit Conversion
In Go, you must explicitly tell the compiler how to convert a value from one type to another. There is no implicit casting between different types.

## Using the `strconv` Package for Type Conversion
The `strconv` package in Go provides utilities for converting strings and other types.

- **`strconv.Atoi()`**: Converts a string to an integer.
- **`strconv.Itoa()`**: Converts an integer to a string.

> **Note:** The second return value from `strconv.Atoi()` is an error, which is ignored using the underscore character (`_`) when you don't need to handle errors explicitly.

### Example:
```go
i, err := strconv.Atoi("42") // Convert string to integer
if err != nil {
    fmt.Println("Error:", err)
}
s := strconv.Itoa(42) // Convert integer to string
