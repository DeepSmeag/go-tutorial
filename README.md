# go-tutorial

Following video here: https://www.youtube.com/watch?v=8uiZC0l4Ajw + personal notes

## Notes

### Tutorial 1

- nothing much, just how a go program is structured
- package = generally a folder with multiple go files
- module = collection of packages
- folder structure and package names don't need to match
- package main is used by the compiler to get the entry point; it looks for the main function there

### Tutorial 2

- all about types, variables, and constants
- exemplified strong typing with required explicit casting/type conversion
- go is pretty strict about types; but you get the control of choosing how many bytes each variables takes
- careful with strings; each character takes 1 byte unless it's outside the original ASCII range; len(string) gives no. of bytes, not of characters
- character = rune in GO; utf8.RuneCountInString(string) gives the actual no. of characters
- type can be omitted if assigned at declaration and it's obvious; but it's good practice to have it there in case we assign from function; instead of wasting time hovering to see type (and some IDEs don't have that), it's better to just have it there
- can omit var (and type) with := operator, but since it requires omitting type I don't like it
- also good practice (if using VSCode with Go extension) don't manually import, just use whatever is needed and go will auto import; the alternative is to have the import get auto removed when saving if doing import->usage, which gets annoying fast
- const needs assignment at declaration, obviously;
- (personal note) int vs int32 types (and 64 variants as well): int is platform dependent, "at least 32 bits"; so depending on the architecture, it can also be 64; so if we need to be specific about our sizes, and also efficient, we need to declare them explicitly

### Tutorial 3

- about functions and control flow
- nil is the null of Go
- functions are enforced with the params and return contract; if a function returns 2 params, you need to get them both in the caller; there's an error otherwise
- I see as a general design pattern the error handling to be C-like (old-style with returning error codes); so we return an error type with a function if that function can have errors; and it's the caller's job to verify it and handle it
- in this case with error handling we cannot have the type written, since some variables will be int for example and the final one will be "error" type; Go does not support inline type declaration for multiple variables
- for && and || operators, Go uses short-circuit evaluation; so if the first condition is false, it won't evaluate the second one; same for ||, if the first is true, it won't evaluate the second one
- printf doesn't end the line; so add \n at the end manually
- bitwise operators also work, they're there; as always, even checks more efficient bitwise
- for switch, we don't necessarily need to specificy a variable to check; we can literally chain ifs; also, the break is implied
- we can also do the usual switch and specify the variable to check against for our cases

### Tutorial 4

- arrays - fixed length, same type, indexable, contiguous in memory (so pretty efficient, but needs to be known at compile time); default value of elements based on default value of type
- pointers / memory; & to get address of where the value is stored; it's getting pretty C-like in here
- unusual, but doing &arr doesn't print the address of the array; so doing &arr[0] would be the first element's address and the array's address as well
  - UPDATE (from tutorial7): this is because we printf %v which interprets as value; do %p for pointer and we get address
- slices - wrappers around arrays (acc. to docs); so arrays with extra functionality
  - they're arrays with dynamic length; declared using [] instead of [n] (or letting the compiler infer the size with [...] and assigning values)
  - the capacity (cap) and length (len) vary, as the cap will double when we try to append to this slice and it's full; it reminds me of dynamic arrays in C++, when I had to create one as homework
  - can use make(type, len, cap) to create a slice with a specific length and capacity
  - is there a reason to work with arrays instead of slices? think C vectors, like for games where we'd want a 3-value vector for coordinates
  - (if a beginner reads this) slices use arrays under the hood; they allocate contiguous memory and use some extra checks on that to give you the functions; when the slice's capacity increases, it creates a new array, copies the elements, and deletes (or rather we let the garbage collector do it) it; so if you get to 100k's of elements, having a slice double itself is expensive; if still confused, read about Data Structures and Algorithms (DSA, any serious university degree has this course on it)
  - there's a speed benchmark at the end of the section to check the cost of reallocation (spoiler: 3x longer)
- maps - key-value pair, looking at value by its key; python's dictionaries, or C/C#/Java's hashmaps fit here
  - declared with map[keyType]valueType
  - if we try to retrieve non-existent key, we get default value of the type (mostly 0 or ""); so map always returns something
  - can also return optional second value, a boolean, to check if the key exists
  - to delete values, delete(map, key) is used; deletes by reference, no return value
  - to add values, map[key] = value; if key exists, it updates the value; if not, it adds it
- for loop - can iterate over the map's keys (and values optionally); order is not guaranteed; in an array, we can have the same syntax and we get index & value
  - traditional for loop is also available; the variable definition is local; it doesn't reassign/affect outside values;
  - if I do var i = 15 outside the for, then do a for loop with i:=0;i<10;i++ it won't affect the i outside the loop
  - we also don't need "var" to declare
- while loop - technically doesn't exist; but we can do a for loop with the condition and it's essentially a while

### Tutorial 5

- more on strings; using e with sharp accent from French as example
- utf8 - variable length encoding; first few bits are used to encode how many bytes the character uses; original ASCII characters are 1 byte, but the rest can be 2, 3, or 4 bytes; this is a more efficient way of encoding characters, since having a fixed length would waste a lot of space
- iterating upon a string with utf8 encoding gives the actual characters; indexing by traditional [index] method gives the bytes; so we can access a 3-byte character byte-by-byte, but it wouldn't be that helpful
- we can cast the string to an array of runes; then we can access indexes and get the actual characters
- string-building: if we use += to concatenate strings, it's inefficient; it creates a new string every time; so we should use strings.Builder instead; this creates an internal array where we append strings; at the end we call .String() to consolidate this array into a string

### Tutorial 6

- structs - user-defined types with multiple fields; default values are the original types' default values
  - we can define & populate inline as well
  - I wonder if Go structs have the same padding problem as C struct - according to an [answer on stacoverflow](https://stackoverflow.com/questions/73211746/does-go-use-something-like-space-padding-for-structs), it does; so sort structs from big to small fields to avoid wasting space; this depends on the system it's running on; if it's a 64-bit system (I assume deployment scenarios are 64-bit), then we have 8-byte word so anything below 8 bytes which changes data type is padded; so a struct with 4 int32 values will have 16 bytes, but if we have 3 int32 values and 1 int8, we would have (int32+int32 = 8 bytes) (int32 + int8 + padding = 5 + 3 = 8 bytes); in this situation we can't do much else, but 16 total bytes are better than 24 or anything else
  - we also have methods = functions tied to structs
- interfaces - Java-like, provides the ability to extract common function and generalize
  - syntax is like struct, but we use interface; functions just use their signature
  - (if a beginner reads this) interfaces tell us that anything that builds upon the interface needs to have the functions the interface defines implemented; like a struct for a Dog and a Cat adhere to an interface Animal, which has a function for sleep(); so both Dog and Cat need to have a sleep() function implementation
  - this is OOP 101
  - though from what I see in the syntax there is no explicit "implements" keyword; so it's a bit tougher to understand when a struct implements an interface, since we have to check the functions
  - interfaces can also include other interfaces
  - there is also a special empty interface (interface{} or "any" starting froom 1.18) used for generic type accepting

### Tutorial 7

- pointers - a special kind of variable that holds the address of another variable; declared as \*type, where type is the type of the variable it points to
  - a lot of what's being talked about here only matters for those who haven't worked extensively with C; some exercise is good anytime, though
    - from what I understand in Go we don't have to worry about freeing memory, we just need to worry about pointer = nil at the end, so that we tell the garbage collector that it can free the memory
    - to allocate memory we use new(type) or make(type, len, cap) for slices, maps, and channels; there are differences, however; make ensures initialization with non-zero values (that's what the internet says, though in my test I see 0-value initialization on slices)
    - from what I gather, _new_ returns the pointer to the allocated memory, while _make_ returns the initialized object; so if we do _new_ on a slice, we get a pointer to a slice, while if we do _make_ on a slice, we get a slice
    - dereferencing is done with \*, same as C; and accessing address is done with &
    - we get runtime error if we try to dereference a nil pointer
    - we now see that slices are actually by reference, so there's underlying pointers
- passing static arrays to functions copies them (so it's pass by value)
- passing slices to functions passes them by reference (so var slicey []type; slicey is actually a pointer to the actual slice data under the hood)
- passing maps to functions passes them by references as well (so var mappy map[type]type; mappy is actually a pointer to the map under the hood)
- we can work with pointers when passing big parameters to avoid copying huge chunks of data
- what if we want to pass by reference, like in C++? This would mean we call functions without worrying about &, yet the function receives the address of that variable; in Go, this isn't possible; everything is passed by value (slices/maps are values which hold references under the hood, so they work like references but are still values)

### Tutorial 8

-
