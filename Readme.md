### Build Production Mode

```sh
go build <name_file>
```

### Run Production Mode

```sh
./<exe_file>
```

### Development Mode Fast Build abd Run

```sh
go run <name_file>
```

### **Unit Test**

- have to use last name `_test`, for example: `hello_world_test.go`

- the function firstname is `Test..`, for Example: `TestHelloWorld()`

- add the parameter `t *testing.T` with no return value.

```go
go test <path>
go test -v <path>
go test -v run <path> <name_function>
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
