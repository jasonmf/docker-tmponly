# docker-tmponly

I'm a newbie to docker and I wanted to build my images `FROM scratch`. I also wanted my contained services to run as `nobody` and be able to write files locally.

This turned out to be tricky because you can only use the `RUN` dockerfile command on executables that are visible from the top layer. `FROM scratch`, there's nothing so you can't `mkdir` or `chown`.

I wrote a small go program to create `/tmp/` world-writable and sticky as is standard. The resulting docker image is completely empty except for the standard `/tmp` directory.

## Usage

```
CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o setup setup.go
docker build -t tmponly:latest .
```
