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
