## Setup
```go
go test -run none -bench . -benchtime 3s -benchmem

go test -run none -bench . -benchtime 3s -benchmem -memprofile m.out

go test -gcflags "-m -m" -run none -bench . -benchtime 3s -benchmem -memprofile m.out // produce escape analysis report

go test -run none -bench . -benchtime 3s -cpuprofile c.out
```

```bash
go tool pprof -inuse_space m.out # default view for memory profile

go tool pprof -alloc_space m.out

go tool pprof c.out

go tool pprof -http :3000 c.out
```

## pprof console
```
(pprof) list algOne
16.50MB   144.51MB (flat, cum) 97.63% of Total
         .          .     69:
         .          .     70:// algOne is one way to solve the problem.
         .          .     71:func algOne(data []byte, find []byte, repl []byte, output *bytes.Buffer) {
         .          .     72:
         .          .     73:   // Use a bytes Buffer to provide a stream to process.
         .   128.01MB     74:   input := bytes.NewBuffer(data)
         .          .     75:
         .          .     76:   // The number of bytes we are looking for.
         .          .     77:   size := len(find)
         .          .     78:
         .          .     79:   // Declare the buffers we need to process the stream.
   16.50MB    16.50MB     80:   buf := make([]byte, size)
         .          .     81:   end := size - 1
         .          .     82:
         .          .     83:   // Read in an initial number of bytes we need to get started.
         .          .     84:   if n, err := io.ReadFull(input, buf[:end]); err != nil {
         .          .     85:           output.Write(buf[:n])

// 16.50MB is flat column
// 144.51MB is cummulative column
```
