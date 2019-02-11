# Be careful about the trailing slash.

```go
http.Handle("/dog/", d)
```
This route catches "/dog", "/dog/", "/dog/something" and "/dog/some/thing".

```go
http.Handle("/cat", c)
```
However, this route only catches "/cat".
