QuickSelect
===========

The quickselect package is a Go package which provides primitives for finding the smallest k elements in slices and user-defined collections. The primitives used in the package are modeled off of the standard sort library for Go. Quickselect uses Hoare's Selection Algorithm which finds the smallest k elements in expected O(n) time, and is thus an asymptotically optimal algorithm (and is faster than sorting or heap implementations).

It is on average 20x faster than a naive sorting implementation.

  BenchmarkQuickSelectSize1e2K1e1   100000             27278 ns/op
  BenchmarkQuickSelectSize1e3K1e1    10000            100185 ns/op
  BenchmarkQuickSelectSize1e3K1e2    10000            209630 ns/op
  BenchmarkQuickSelectSize1e4K1e1     5000            694067 ns/op
  BenchmarkQuickSelectSize1e4K1e2     1000           1627887 ns/op
  BenchmarkQuickSelectSize1e4K1e3     1000           2030468 ns/op
  BenchmarkQuickSelectSize1e5K1e1      500           6518500 ns/op
  BenchmarkQuickSelectSize1e5K1e2      200           8595422 ns/op
  BenchmarkQuickSelectSize1e5K1e3      100          16288198 ns/op
  BenchmarkQuickSelectSize1e5K1e4      100          20089387 ns/op
  BenchmarkQuickSelectSize1e6K1e1       50          67049413 ns/op
  BenchmarkQuickSelectSize1e6K1e2       50          70430656 ns/op
  BenchmarkQuickSelectSize1e6K1e3       20         110968263 ns/op
  BenchmarkQuickSelectSize1e6K1e4       10         161747855 ns/op
  BenchmarkQuickSelectSize1e6K1e5       10         189164465 ns/op
  BenchmarkQuickSelectSize1e7K1e1        2         689406050 ns/op
  BenchmarkQuickSelectSize1e7K1e2        2         684015273 ns/op
  BenchmarkQuickSelectSize1e7K1e3        2         737196456 ns/op
  BenchmarkQuickSelectSize1e7K1e4        1        1876629975 ns/op
  BenchmarkQuickSelectSize1e7K1e5        1        1454794314 ns/op
  BenchmarkQuickSelectSize1e7K1e6        1        2102171556 ns/op
  BenchmarkQuickSelectSize1e8K1e1        1        6822528776 ns/op
  BenchmarkQuickSelectSize1e8K1e2        1        6873096754 ns/op
  BenchmarkQuickSelectSize1e8K1e3        1        6922456224 ns/op
  BenchmarkQuickSelectSize1e8K1e4        1        16263664611 ns/op
  BenchmarkQuickSelectSize1e8K1e5        1        16568509572 ns/op
  BenchmarkQuickSelectSize1e8K1e6        1        20671732970 ns/op
  BenchmarkQuickSelectSize1e8K1e7        1        22154108390 ns/op
