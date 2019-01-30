# Handle

[Source](https://godoc.org/net/http#Handle)

```go
http.Handle("/dog/", d)
```
This route catches "/dog", "/dog/", "/dog/something" and "/dog/some/thing".

```go
http.Handle("/cat", c)
```
However, this route only catches "/cat".
Be careful about the trailing slash.

