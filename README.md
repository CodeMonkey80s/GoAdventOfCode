### Go Advent Of Code Solutions

Repository with my solutions to [Advent Of Code](https://adventofcode.com) puzzles.

### Progress

The plan is to ultimately solve all the puzzles...

| #  | 2024 | 2023 | 2022 | 2021 | 2020 | 2019 | 2018 | 2017 | 2016 | 2015 |
|:---|:-----|:-----|:-----|:-----|:-----|:-----|:-----|:-----|:-----|:-----|
| 01 | ✅✅   | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 
| 02 | ✅✅   | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 
| 03 | ✅✅   | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 
| 04 | ✅✅   | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 🔲🔲 | 

### Static Analysis

Run `golangci-lint` command. There is a hidden configuration file for this tool inside root directory `.golangci.yml`.

```
$ golangci-lint --color always run --verbose ./...
```

### Testing

```
$ go test -v ./...
```
