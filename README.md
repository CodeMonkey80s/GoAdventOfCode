### Go Advent Of Code Solutions

Repository with my solutions to [Advent Of Code](https://adventofcode.com) puzzles.

### Progress

The plan is to ultimately solve all the puzzles...

| Year    | 01 | 02 | 03 | 04 | 05 | 06 | 07 | 08 | 09 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 | 24 | 25 |                   
|:--------|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|:---|
| 2024 p1 | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | 🔲 | ✅  | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 | ✅  |
| 2024 p2 | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 |
| 2023 p1 | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | 🔲 | ✅  | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 | ✅  |
| 2023 p2 | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 | 🔲 |

### Static Analysis

Run `golangci-lint` command. There is a hidden configuration file for this tool inside root directory `.golangci.yml`.

```
$ golangci-lint --color always run --verbose ./...
```

### Testing

```
$ go test -v ./...
```
