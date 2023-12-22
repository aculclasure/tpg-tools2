# tpg-tools2
This repo contains the work I did while reading John Arundel's [_The Power of Go: Tools (2024)_](https://bitfieldconsulting.com/books/tools).


### Thoughts I had while reading _The Power of Go: Tools (2024)_
- While reading Chapter 2 _Paperwork_, I came up with some helpful criteria to use when deciding whether to export or hide fields in a struct:
  - If you don't want the user to directly change the state of struct over the course of the program, then make the struct and its fields unexported and instead provide a constructor that sets the fields to acceptable, validated values.
  - If the user does need to change the state of the struct over the course of the program, but you want to make sure the fields get set to acceptable values, then the fields can be hidden and you can use methods to validate and set the fields.
  - If you don't care how the state of the struct changes over the course of the program and you make its fields exported, you should provide a constructor that at least lets you start out with the struct in a known, working, good initial state (e.g. a sane default state).
- After reading the section "The art of judicious logging" in Chapter 5 _Files_, I realized I have spent a lot of development time overly worrying about logging and doing things like exposing command-line flags in many programs to let the user decide what file to write logs to. I like the suggestion to just write logs to the standard output or standard error streams and allow the OS to deal with redirecting that output to files if the users decide to do so!
- I can't recall exactly which chapter this was in, but one of the provided solutions to the end-of-chapter exercise introduced me to the [`io.MultiReader()`](https://pkg.go.dev/io#MultiReader) and [`io.MultiWriter()`](https://pkg.go.dev/io#MultiReader) functions. How useful! If you ever have the need to read from multiple io.Readers, this helps simplify the process and you don't have to deal with reading from each io.Reader individually. Ditto for writing with the `io.MultiWriter()` function.
  