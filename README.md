# Golist
> List of articles, projects, books with small annotations, my comments and (may be) pros and cons.

## Testing
+ [Fakeit](https://github.com/brianvoe/gofakeit) - random data generator with over 160 functions. You can create your own functions, can fire up built in local server to test functions.
+ [Testfixtures](https://github.com/go-testfixtures/testfixtures) - helps with functional testing, keeping fixture data in yml files.

## Slices
+ [Slice tricks](https://ueokande.github.io/go-slice-tricks/) - cheat sheet for common slice operations.
+ [Bad go slice of pointers](https://philpearl.github.io/post/bad_go_slice_of_pointers/) - why you should not use slice of pointers in most cases

## Project Structure
+ [Cli project structure](https://bencane.com/2020/12/29/how-to-structure-a-golang-cli-project/) - describes simple structure of cli project.

## Gotchas
+ [Cmd dev null](https://rohitpaulk.com/articles/cmd-run-dev-null) - if either one of Cmd.Stdout, Cmd.Stderr, Cmd.Stdin are not explicitly set, Go sets them to /dev/null. This can cause failures when using Cmd.Run in Go programs that execute in environments where /dev/null isn’t accessible. To fix this, either assign dummy values to Cmd.Stdout, Cmd.Stderr and Cmd.Stdin, or ensure that /dev/null is accessible.