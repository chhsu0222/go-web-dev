The 1st letter of field's name in a strct should be capitalized so it can be used in other package.

e.g. In the following case, p.name would be unexported.

```go
type person struct {
	name string
	Age  int
}
```

