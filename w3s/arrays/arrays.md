### **Go Arrays**

Arrays in Go are used to store multiple values of the same type in a single variable. They are fixed in size and provide efficient storage and retrieval of elements.

----------

### **Declaring an Array**

There are two ways to declare an array:

#### 1. Using the `var` Keyword:



`var array_name = [length]datatype{values} // Length is explicitly defined` 

OR


`var array_name = [...]datatype{values} // Length is inferred from the number of values` 

#### 2. Using the `:=` Operator:




`array_name := [length]datatype{values} // Length is explicitly defined` 

OR



`array_name := [...]datatype{values} // Length is inferred from the number of values` 

> **Note:**
> 
> -   The `length` specifies the number of elements the array can store.
> -   Arrays in Go have a fixed length that can either be explicitly defined or inferred by the compiler based on the number of values provided.

----------

### **Accessing Array Elements**

Array elements are accessed by their index, which starts at `0`.

Example:




`fmt.Println(array_name[0]) // Access the first element` 

----------

### **Modifying Array Elements**

You can update a specific array element by using its index.

Example:




`array_name[1] = 42 // Update the second element` 

----------

### **Array Initialization**

If an array or one of its elements is not explicitly initialized, it will be assigned the **default value** for its type.

-   **Default Values:**
    -   `int`: `0`
    -   `string`: `""` (empty string)

Example:


`var arr [3]int // All elements are set to 0` 

----------

### **Initializing Specific Elements**

You can initialize only specific elements in an array using the index:value format.

Example:



`arr := [5]int{1: 10, 2: 40}` 

-   `1:10` assigns `10` to index `1` (second element).
-   `2:40` assigns `40` to index `2` (third element).

----------

### **Finding the Length of an Array**

The `len()` function returns the length of an array.

Example:


`fmt.Println(len(array_name)) // Prints the length of the array` 

----------