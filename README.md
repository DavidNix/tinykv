# tinykv
A small, in-memory key/value database. Experimental only. Not for production use.

To see useful commands:

```shell
make help
```

## To Run

Prereq: Go 1.18 is installed. (Hint: Try `go env`)

```shell
# Run the cli
make run

# Run test suite
make test
```

## TODOs
* Counts within transactions do not work. There is a failing unit test showing as such.
* Abstract interactive functionality in `main()` into a testable component.
