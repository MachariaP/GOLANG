# Go Variables

A `Variable` is a name given to a storage area that programs can manipulate.

## Go Variable Types

There are different types of variables in Go:

1. **Integer** `(int)`
   - Stores integers
   - Whole numbers such as `123` or `-123`.

2. **Float32** `(float)`
   - Stores floating point numbers with decimals
   - Example: `19.99` or `-19.99`.

3. **String** `(string)`
   - Stores text
   - Example: `"Hello World"`
   - String values are surrounded by double quotes.

4. **Boolean** `(bool)`
   - Stores values with two states
   - Example: `true` or `false`.

## Declaring / Creating Variables

There are two ways to declare variables in Go:

1. **With the `var` keyword:**
   - Syntax: `var variablename type = value`
   - You must specify either `type` or `value` (or both).

2. **With the `:=` sign:**
   - Syntax: `variablename := value`
   - In this case, the type of the variable is *inferred* from the value. The compiler decides the type based on the value.

## Difference between `:=` and `var`

1. `:=`
   - Can only be used inside functions
   - Variable declaration and value assignment must be on the same line.

2. `var`
   - Can be used both inside and outside functions
   - Variable declaration and value assignment can be done separately.
