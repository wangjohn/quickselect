QuickSelect
===========

The quickselect package is a Go package which provides primitives for finding the smallest k elements in slices and user-defined collections. The primitives used in the package are modeled off of the standard sort library for Go. Quickselect uses Hoare's Selection Algorithm which finds the smallest k elements in expected O(n) time, and is thus an asymptotically optimal algorithm (and is faster than sorting or heap implementations).

It is between 3-10x faster than a naive sorting implementation.

    BenchmarkQuickSelectSize1e2K1e1   100000             25604 ns/op
    BenchmarkQuickSelectSize1e4K1e1     1000           1590871 ns/op
    BenchmarkQuickSelectSize1e6K1e1       20         206866094 ns/op
    BenchmarkQuickSelectSize1e8K1e1        1        19385926017 ns/op
    BenchmarkQuickSelectSize1e4K1e2     1000           1641690 ns/op
    BenchmarkQuickSelectSize1e6K1e2       10         147239258 ns/op
    BenchmarkQuickSelectSize1e8K1e2        1        16673482431 ns/op
    BenchmarkQuickSelectSize1e6K1e4       10         163698995 ns/op
    BenchmarkQuickSelectSize1e8K1e4        1        10243557555 ns/op
    BenchmarkQuickSelectSize1e8K1e6        1        15229638769 ns/op
    BenchmarkSortSize1e2K1e1           50000             55482 ns/op
    BenchmarkSortSize1e4K1e1             100          12311364 ns/op
    BenchmarkSortSize1e6K1e1               1        2005734006 ns/op
    BenchmarkSortSize1e8K1e1               1        266734709477 ns/op


