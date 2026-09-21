# example-go-hello

A hello-world web app in Go, for FadeHost app hosting.

Deploy it from the panel with no start command and no build command: the
runtime sees the `go.mod`, compiles the module and runs the binary.

It reads the port from the `PORT` variable, which FadeHost sets for every
app, and prints the result of a request to itself once it is listening.

## Running it yourself

```bash
PORT=8080 go run .
curl http://localhost:8080/
```

## Licence

MIT, see [LICENSE](LICENSE).
