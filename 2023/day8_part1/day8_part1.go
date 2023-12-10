package day8_part1

type Node struct {
	Name  string
	Left  string
	Right string
}

func howManySteps(lines []string) int {

	steps := lines[0]
	nodes := make(map[string]Node)
	for _, line := range lines[2:] {
		name := line[0:3]
		left := line[7:10]
		right := line[12:15]
		n := Node{
			Name:  name,
			Left:  left,
			Right: right,
		}
		nodes[name] = n
	}

	numberOfSteps := 0
	node := "AAA"
	next := ""
	loops := 1
outer:
	for {
		if loops == 1 {
			next = ""
		}
		for _, dir := range steps {
			switch dir {
			case 'L':
				next = nodes[node].Left
			case 'R':
				next = nodes[node].Right
			}
			numberOfSteps++
			if next == "ZZZ" {
				break outer
			}
			node = next
		}
		loops++
	}

	return numberOfSteps
}
