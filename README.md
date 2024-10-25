# go-tutorial

Following video here: https://www.youtube.com/watch?v=8uiZC0l4Ajw + personal notes

- the notes are from the perspective of someone who has worked with C/C++ at a job and has a good understanding of Python/Java/C# as well from university courses
- Go's syntax and features will be compared to these languages
- some things I'll explain like I was talking to a beginner (it's a good way to practice deep understanding of a topic)
- some will be new to me as well (Go-specific stuff)

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

- goroutines - lightweight threads managed by the Go runtime; they're not OS threads, but they're multiplexed onto OS threads; so we can have thousands of goroutines running on a single OS thread; it's way more efficient this way; they're also called green threads from what I've found; Java implemented them kind of recently as well
  - concurrency vs parallel execution - concurrency is about dealing with multiple things at once, while parallel execution is about executing multiple things at once; so concurrency can be parallel, but it doesn't have to be; Go is concurrent by design, so it's easy to write concurrent code; parallelism is achieved by running multiple goroutines on multiple OS threads
- go keyword; this is like async from JS, we don't wait for the function to finish; it starts the goroutine; btw, I like the way we use the keyword, feels witty
- we still need to wait for the goroutines to finish, so we have waitGroups for that;
- and if we modify the same data from multiple goroutines, we need to handle data race conditions; so we use mutexes
- since our program is limited by the delay simulation and not the actual result appending, we still get little overhead; but if we were to move the m.Lock() call above the delay (.Sleep()), then we would destroy our concurrency and get a linear execution time
- we also have RWMutex available, with separate locks for reading and writing; (if a beginner reads this) as a Parallel programming primer, we don't need to lock each reader out since they don't modify the data; we can have multiple readers at the same time; but when a goroutine tries to write, it should lock the readers out; so we use RLock() for reading and Lock() for writing and Go handles the details; keep in mind that time taken to finish a task depends on the task; the constant time scaling of our simulated delay isn't available in real-world scenarios; the flatter the scaling, the better the parallelism of the task/algorithm

### Tutorial 9

- channels - a mechanism to communicate between goroutines
  - they hold data
  - they're thread safe = we avoid data race conditions when r/w from multiple goroutines
  - we can listen for data addition and block execution until it happens
- special syntax: <- to add data to the channel
- simply using a channel on the main thread will block, but Go is smart and it throws a deadlock error; why it happens: writing to the channel blocks the thread execution until something reads from that channel
- we can iterate over a channel with for-range loop; in this case we need the goroutine producing data to close the channel so that the consumer knows no more data is coming
- _defer_ keyword: "do this right before the function exits"; stacking multiple defers will execute them in reverse order; so the last defer will be the first to execute; this is useful for closing resources, like files or channels
- buffer channels - the producer process does not need to wait for the consumers to do their job, it should be able to exit quickly after sending the data; that's why we can have buffer channels, which store more pieces of data; so the producer can send data and fill it up (or not), and the consumer(s) can take data at their own leisure
- select statement - like a switch statement, but for channels; we can listen to multiple channels at once; if multiple channels have data, Go will randomly choose one to read from; if none have data, it will block until one has data; if multiple have data, it will randomly choose one

### Tutorial 10

- generics: anytime we want to apply the same kind of processing behaviour to a number of different types, we need a way to use generics; instead of writing 4 functions to add numbers in an array/slice, depending on their type (int/float/32-64 variants), we can use generics
- syntax is TS-like, with type in angle brackets and we give it a name; so [T int | float32 | float64] will have T represent any of those 3 types; starting with Go 1.18, we can use "any" as a type to represent any type, though addition is only numerical so we'll keep at that
- we can declare a type (generic) that encompasses multiple types, like in TS, but only to be used in generic functions; we cannot declare a variable of that type, since Go is strong-typed and it wants to resolve the type at compile time; the syntax to do this is an interface declaration
  - see example in tutorial 10 which adds of 2 numbers
- sidenote: if we're not interested in a value extracted from a function or in a for-range loop, use \_ to ignore it;
- it makes sense to have an "any" generic when doing operations which are available on any type, like checking if a slice is empty;
- we also have an example using a JSON file here; it's good to practice reading files as well; ioutil is deprecated, I used os instead; it seems that os.ReadFile also looks at the path where the command was executed (I'm executing this from the root of the project so I need to add the subfolders to the path)
- json.Unmarshal() used to decode JSON data
- the guy says he uses generics almost exclusively with functions; indeed, when doing microservice / API work, we mostly work with entities linked to the ORM & db and they're clearly defined; so it's more likely to use generics on functions, rather than define generic struct types to "inherit" from

### Tutorial 11

- API creation; we use Postman to test things out; the API looks up amount of coins in an imaginary account, with authentication via header token ('Bearer ...' without the Bearer part)
- this part of the tutorial also discusses project structure; check [here](https://github.com/golang-standards/project-layout) for more info on project layout guidance
  - api folder - specs (parameters, response types for endpoint)
  - cmd/api - contains main.go folder
  - internal - contains the actual code; we don't want to expose this to the outside world; this pattern is enforced by the Go compiler (see the link above)
  - pkg - reusable code, shareable with other projects; so a place to put libraries/ other tools; we don't actually use this one in this example
- **naming convention: Capital naming means the function can be imported in other packages; lower case means private function, so only usable within the package**
- middleware - usage via r.Use (r is the router, pointer \*chi.Mux); middlewares are pieces of code we run before the actual path handler for that API endpoint request, or after (this is the case at least with Express/NestJS, though I expect Go to have it as well); so maybe we want to log the request and the user (if authenticated), we would use a middleware to log before the request reaches the actual handler function; or maybe we want to check authentication, that would be a middleware job; an unauthenticated user's request does not need to reach our handler function since it gets denied before that
- we use go-chi for routing; it's lightweight, but it seems to require quite a bit of boilerplate; Express is definitely way faster to setup (5 lines of code); the tradeoff is that Go is much faster since it's compiled and multi-threaded
- the main knowledge that's necessary here for a beginner is first of all how HTTP servers are setup; the Go-related difficulty here is how the project is structured, how the router is setup and how we use middleware and handle errors; once the structure is understood, the rest is just writing the route-handling code; it's quite different from how Express would behave
- I'd recommend a slow and thorough pass through this section of the video and a review after a day to cement the knowledge
- there's lots of concepts to unpack here which are more related to web development, rather than Go itself; the only thing that's Go-specific is the structure (again) and the syntax; so I'll shortly explain concepts just so they're named; middleware is described above
  - mock - a fake implementation of a function or more abstract element; a mock database, for example, would have fake values to test with, so that we don't actually pull from a real database with real values and maybe mess things up in there while we build our features
  - testing - calling our functions with known / random values and expect a certain behaviour in return; we usually use mocked elements while testing; testing is in itself an entire discipline and requires careful study by itself
- a review on project structure:
  - we consider tutorial11 as the root of the project; _go.mod_ defines our module and dependencies; _go.sum_ has checksums of the dependencies and their various versions that are required throughout the project
  - Go projects have typical structure where we put entrypoint commands (think of commands as executables, so a single Go project can have multiple executable files, so different purposes) in the _cmd_ folder; the _cmd/api_ path is a way of telling us that the _main.go_ file in there is the entrypoint for API stuff (meaning our HTTP server);
  - (mention) the commands don't need to be called _main.go_; they could be called _<package>.go_ if the project is more focused on devising a library instead; what matters is they define the main package inside and there's a **main()** function in there
  - the way we import packages so we can use function from there is with _import "module-name/path-to-package"_; so if we have a package in the _internal_ folder (let's pick _handlers_ in our case), we would import it with _import "goapi/internal/handlers"_ we can assign a different name to the package where we import it by using _import h "goapi/internal/handlers"_ and then we can use _h.FunctionName()_ to call the function; the **goapi** part is the module name, which is defined in the _go.mod_ file; the docs say that the module name should be the Github path to the project once uploaded; I'm still confused about that, though I guess it works like this:
    - if our module is called _github.com/username/modulename_, then we would import our "handlers" package with _import handlers "github.com/username/projectname/internal/handlers"_; this would not be permitted if we were to import an external project, since the Go compiler enforces everything under "internal" to be used only within the project;
    - this import would work if our module was called _github.com/username/modulename_; the compiler would look at local files, it would see this and it would know what to do; at the same time, this naming convention lets Go look at the Github repository to find modules and download them if they're not local, so it's like the npm equivalent from Nodejs
    - for a project meant to be kept private (like our backend), I believe this naming convention does not do anything; no one should be able to import our project
    - going back to our _internal/handlers_ folder; there is a convention for all files inside this folder to have the same package name as the folder they're in (so "handlers"); we need at least 1 function with a Capital letter (e.g. Handlers in our case), which will be the entrypoint from other packages; every other function in any file is private to the package and can be used freely within the package (across different files) with no imports; the compiler and the IDE should know how to find them
    - in our case, as we're building an HTTP server (or microservice), almost all of the project should be placed in the _internal_ folder, since we want to share nothing;
    - we do have the _api_ folder, where we put our OpenAPI specs, JSON schema files, protocol definition files (acc. to the Github Go project layout repo); in our case, there are response structs and error handlers; I don't know for sure if it's a good idea to put our error handlers there, but the response structs make sense; Go is often used for microservices, which means multiple HTTP servers need to communicate between themselves; to have type safety, since Go is known for statical and strong typing, we can use this _api/_ folder to define our communication contracts and other modules can import them using that github naming scheme; now I'm wondering how things go when we want to pull from a private repo (update: there are methods to pull this off, so it's ok)

### Personal notes on Go frameworks

- I'm coming from the JS/TS/Nodejs world where we have packages for everything and there's no way you're doing something without a framework if the project is even a tiny bit complex; so I was wondering if there's an industry-standard frameworkd people use for serving HTTP requests in Go; and there seems to be none
- the more I search about this, the more I see that many people and companies simply use net/http to build everything, which means no separate framework is being used; at the same time, there's still a considerable number of people(&copanies) opting for frameworks like fiber, chi, or gin; so my understanding is that Go's landscape is fragmented; I'd say JS is more straightforward in this case, since there are just a few frameworks which are widely used and you find companies using and hiring for them (React/Express for example)
- at the same time, I'm more impressed by Go in the sense that you're not a <framework> developer, but instead you're a Go developer and you're expected to handle any framework cause you have deeper expertise; I wish that were the case with how people view JS frameworks as well
- back to the subject - I see a trend where most people say net/http is enough; Fiber is the Express-like framework and it prouds itself with speed, so I guess I'd recommend that to people coming over from JS; I also see Gin and Echo being mentioned almost equally and some people say they use Chi; maybe these are the main ones; anyway, the syntax looks similar so there's high chance someone who knows the underlying working of any HTTP server can pick any one of them up rather quickly
