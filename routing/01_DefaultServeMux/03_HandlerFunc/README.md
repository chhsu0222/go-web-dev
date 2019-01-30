# HandlerFunc

[http.HandlerFunc](https://godoc.org/net/http#HandlerFunc)

``` Go
type HandlerFunc func(ResponseWriter, *Request)
```
If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.

``` Go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
```
ServeHTTP calls f(w, r).

**This is just a nice thing to know about. You wouldn't do this in production code probably.**

***

## Question
Could you get http.Handle to take a func with this signature: func(ResponseWriter, *Request)?

[Source](https://github.com/GoesToEleven/golang-web-dev/tree/master/020_HandlerFunc)
