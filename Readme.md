## How to Run

### Build Production Mode

```sh
go build <name_file>
```

### Run Production Mode

```sh
./<exe_file>
```

### Development Mode Fast Build and Run

```sh
go run <name_file>
```

## Unit Test

- have to use last name `_test`, for example: `hello_world_test.go`

- the function firstname is `Test..`, for Example: `TestHelloWorld()`

- add the parameter `t *testing.T` with no return value.

```go
go test <path>
go test -v <path>   // -v use to show the logs
go test -v -run=<name_function> <path> -count=1   // -count to ignore caching
```

### _Testing.T_

- Parameter for unit testing

- `Fail()` failing the unit test but continue till finish.
- `FailNow()` failing the unit test and stoppped at the same time.
- `Error()` call `Fail()` with print an error as a logging.
- `Fatal()` call `FailNow()` with print an error as a logging.

### _Testing.M_

- Parameter for testing in main

### _Testing.B_

- Parameter for brenchmarking

## Goroutines

- goroutine is pretty small, it work as concurrency instead of parallel.

### channel

- channel is an alternative to receive data by goroutine. it is being a tunnel to communicate with goroutine function. goroutine will be blocked untill the channel receive the data (blocking). The concept is similar with `async await` in javascript.

```go
channel := make(chan int)
channel <- 19               // to assign the data
bufferData <- channel       // to give the data
```

- channel only buffer one data. If want more, need another goroutine to create

- please close channel if it's not used due to memory leak.

- by default, the parameter is parsing the value. It is different with channel. If parsing the channel, it is automatically passing the reference.

- channel is possible to give a sign in and out while parsing as parameter. `chan<-` for in and `<-chan` for out.

- _Buffered Channel_ use to buffer data whenever didn't available or in some case to handle sender faster than receiver.

```go
channel := make(chan string, 3)   // create channel with 3 buffer long.
```
