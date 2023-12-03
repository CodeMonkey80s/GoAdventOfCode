# Go Advent Of Code Solutions

<p align="center">
    <a href="https://github.com/CodeMonkey80s/GoAdventOfCode"><img src="gopher-advent.png" width="500"/></a>
</p>

Repository with my solutions to **Advent Of Code** problems.

# Solutions // 2023

| #    | Title                                                                                                                                 | Stars  |
|:-----|:--------------------------------------------------------------------------------------------------------------------------------------|:-------|
| 1    | [Day 1: Trebuchet?!, Part 1](2023/day1_part1/day1_part1.go)                                                                           |⭐      |
| 2    | [Day 1: Trebuchet?!, Part 2](2023/day1_part2/day1_part2.go)                                                                           |⭐      |
| 3    | [Day 2: Cube Conundrum, Part 1](2023/day2_part1/day2_part1.go)                                                                        |⭐      |
| 4    | [Day 2: Cube Conundrum, Part 2](2023/day2_part2/day2_part2.go)                                                                        |⭐      |

# Static Analysis

Run `golangci-lint` command. There is a hidden configuration file for this tool inside root directory `.golangci.yml`.

```
$ golangci-lint --color always run --verbose ./...
```

# Testing

```
$ go test -v ./...
```

# Credits

The gopher image at the top is taken from: https://github.com/MariaLetta/free-gophers-pack
