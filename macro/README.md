## Setup
```bash
GODEBUG=schedtrace=1000 ./macro > /dev/null

GODEBUG=gctrace=1 ./macro > /dev/null
```

## Load testing
```
hey -m POST -c 100 -n 100000 "http://localhost:1000/search?term=trump&cnn=on&bbc=on&nyt=on"
```

## pprof
```
go tool pprof -alloc_space http://localhost:1000/debug/pprof/allocs
```
