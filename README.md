# Sum Types

This library is an attempt to emulate [sum types](https://en.wikipedia.org/wiki/Algebraic_data_type) using [Go generics](https://go.dev/doc/tutorial/generics).

| Package | Description |
| -- | -- |
| `opt.Optional` | Optional type. |
| `st2.SumType2` | Sum type for 2 types. |
| `st3.SumType3` | Sum type for 3 types. |

## Sum type example

Let's create a sum type variable which can hold an `int` or a `string`. In such case, `int` will be our type 0, and `string` will be our type 1. To initialize the variable with a `string` value, which is our type 1, we call the `New1()` function:

```go
nameOrAge := st2.New1[int, string]("John")
```

Attempting to retrieve the `string` value with the `Get1()` method:

```go
if name, ok := nameOrAge.Get1(); ok {
    fmt.Printf("Name is %s.\n", name)
}
```

To reassign this variable to an `int`, which is our type 0, we call the `New0()` function. And we attempt to read it with the `Get0()` method:

```go
nameOrAge = st2.New0[int, string](40)

if age, ok := nameOrAge.Get0(); ok {
    fmt.Printf("Age is %d.\n", age)
}
```

Exhaustive pattern matching can be performed by passing callback functions to the `Match()` method:

```go
nameOrAge.Match(
    func(age int) {
        fmt.Printf("Age is %d.\n", age)
    },
    func(name string) {
        fmt.Printf("Name is %s.\n", name)
    },
)
```

For sum type structs with more fields, the logic is the same: `st3` exposes functions `New0()`, `New1()` and `New2()`, and so on.

## License

Licensed under [MIT license](https://opensource.org/licenses/MIT), see [LICENSE.md](LICENSE.md) for details.
