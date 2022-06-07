#How to run the benchmarks

```
go test -v -run="none" -bench=. -benchtime="5s" -benchmem

```
##My personal benchmark results

```
goos: linux
goarch: amd64
pkg: github.com/cjodra14/go_microservices/chapter1/example1_4
cpu: AMD Ryzen 7 5700U with Radeon Graphics         
BenchmarkHelloHandlerVariable
BenchmarkHelloHandlerVariable-16                12673378               556.4 ns/op            96 B/op          4 allocs/op
BenchmarkHelloHandlerEncoder
BenchmarkHelloHandlerEncoder-16                 30754070               174.3 ns/op            16 B/op          1 allocs/op
BenchmarkHelloHandlerEncoderReference
BenchmarkHelloHandlerEncoderReference-16        40639287               132.7 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/cjodra14/go_microservices/chapter1/example1_4        18.641s
```