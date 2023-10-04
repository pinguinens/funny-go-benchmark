# Funny Go Benchmark

---

Synthetic benchmark for fun. It tests whole system memory (CPU cache, RAM, SWAP).

## How to use

Run

``$ go run . -c 2 -b 2048``

### Options

#### -b
Buffer size in megabytes which allocates memory. Default is 1 MB.

#### -c
Routines count. It usually equals cpu cores. Default is all cpu cores.

## Enjoy!