# MAIN
## Interface

### How to run

```
go test -benchmem -run=^$ -bench . .
```

### Result

```
goos: windows
goarch: amd64
pkg: github.com/go-practice/example15-PerformenceBenchmark
cpu: 11th Gen Intel(R) Core(TM) i7-1185G7 @ 3.00GHz
BenchmarkPrintInt2String01-8    24117063                47.03 ns/op
BenchmarkPrintInt2String02-8    73777143                14.89 ns/op
BenchmarkPrintInt2String03-8    74606914                14.75 ns/op
PASS
ok      github.com/go-practice/example15-PerformenceBenchmark   3.473s
```