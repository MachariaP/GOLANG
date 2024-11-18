# Data types in Go
- Specifies the type of data that a Go variable can hold.
- The types can be divided into 4 categories:
    1. Basic type:
        - Numbers
        - Strings (A piece of text surrounded by double quotes)
        - Booleans (true or false)

    2. Aggregate type:
        - Arrays and structs

    3. Reference type:
        - Pointers, slices, maps, functions and channels

    4. Interface type

    ## Integers (int)
    - Can be divided into two :-
        1. Signed integers
        2. Unsigned integers

    ** Note: Use unit only when you need to represent a value as an unsigned number for a certain reason. **


        ### Signed integers
        - This types are ```int8```, ```int16```, ```int32```, ```int64``` and ```int```.
        - They store both negative and positive values
        - The total range of values is split between positive and negative numbers.
        - Can store values ranging from negative values to positive values, but can not store as large a positive number as its unsigned counterpart(uint).
        - The maximum positive value of a signed integer type is half of the maximum value of its corresponding unsigned type.

    ### Rune
    - A rune is an alias for int32 data type.
    - Used to represent a Unicode character (or a Unicode code point).

    ## Floating-point numbers
    - The two size types of floating-point numbers are:
        * float32
        * float64

    - Are used when you need to store large numbers and they don't fit in previously defined integers.

    ## Type Conversion
    - The process of converting one data type into another.
    - Go requires explicit conversion(you have to tell the program what you want to convert the data type to)
    - Unlike other languages where this happens automatically(implicit casting)

    ** Using the '''strconv''' Package for String Conversion
         - Helps with converting strings to integers and vice versa.
         - strconv.Atoi converts strings to integers.
         - strconv.Itoa converts integers to strings.

         *** Note: The underscore character (`-`) is used to ignore the second value returned by strconv.Atoi, which is the error.