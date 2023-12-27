package day10_part1

import (
	"fmt"
	"slices"
	"strings"

	"GoAdventOfCode/util"
)

type Bot struct {
	Number int
	Chips  []int
}

func (b *Bot) receive(value int) {
	b.Chips = append(b.Chips, value)
	if len(b.Chips) == 2 && b.Number > 0 {
		if b.Chips[1] == compareValue1 && b.Chips[0] == compareValue2 || b.Chips[0] == compareValue1 && b.Chips[1] == compareValue2 {
			botNumber = b.Number
		}
	}
	slices.Sort(b.Chips)
}

func (b *Bot) giveLowerChip() int {
	n := b.Chips[0]
	return n
}

func (b *Bot) giveHigherChip() int {
	n := b.Chips[1]
	return n
}

func NewBot(n int) *Bot {
	return &Bot{
		Number: n,
		Chips:  make([]int, 0),
	}
}

func GetBot(n int) *Bot {
	b, ok := listOfBots[n]
	if !ok {
		return NewBot(n)
	}

	return b
}

var listOfBots = map[int]*Bot{}
var listOfBins = map[int][]int{}

var botNumber = -1
var compareValue1 = -1
var compareValue2 = -1

func getAnswer(lines []string, value1 int, value2 int) int {

	botNumber = -1
	compareValue1 = value1
	compareValue2 = value2

	listOfBots = make(map[int]*Bot)
	listOfBins = make(map[int][]int)

	initialize(lines)
	process(lines)

	fmt.Println("BotNumber:", botNumber)
	fmt.Println("Output:", listOfBins[0][0]*listOfBins[1][0]*listOfBins[2][0])
	//util.PrintOrderedMap("listOfBins", listOfBins)
	return botNumber
}

func initialize(lines []string) {
	for _, line := range lines {
		parts := strings.Fields(line)
		if parts[0] == "value" {
			chip := util.ConvertStringToInt(parts[1])
			n := util.ConvertStringToInt(parts[5])
			b := GetBot(n)
			listOfBots[n] = b
			b.receive(chip)
		}
	}
}

func process(lines []string) {
	for {
		for _, line := range lines {
			parts := strings.Fields(line)
			//fmt.Printf("PARTS: %+v\n", parts)
			switch {
			case parts[0] == "value":
			case parts[0] == "bot" && parts[5] == "output" && parts[10] == "bot":
				n1 := util.ConvertStringToInt(parts[1])
				out := util.ConvertStringToInt(parts[6])
				n2 := util.ConvertStringToInt(parts[11])
				b1 := GetBot(n1)
				b2 := GetBot(n2)
				if len(b1.Chips) < 2 {
					continue
				}
				b2.receive(b1.giveHigherChip())
				addToBin(out, b1.giveLowerChip())
				listOfBots[n2] = b2
				delete(listOfBots, n1)
			case parts[0] == "bot" && parts[5] == "output" && parts[10] == "output":
				n1 := util.ConvertStringToInt(parts[1])
				out1 := util.ConvertStringToInt(parts[6])
				out2 := util.ConvertStringToInt(parts[11])
				b := GetBot(n1)
				if len(b.Chips) < 2 {
					continue
				}
				addToBin(out1, b.giveLowerChip())
				addToBin(out2, b.giveHigherChip())
				delete(listOfBots, n1)
			case parts[0] == "bot" && parts[5] == "bot" && parts[10] == "bot":
				n1 := util.ConvertStringToInt(parts[1])
				n2 := util.ConvertStringToInt(parts[6])
				n3 := util.ConvertStringToInt(parts[11])
				b1 := GetBot(n1)
				b2 := GetBot(n2)
				b3 := GetBot(n3)
				if len(b1.Chips) < 2 {
					continue
				}
				b2.receive(b1.giveLowerChip())
				b3.receive(b1.giveHigherChip())
				listOfBots[n2] = b2
				listOfBots[n3] = b3
				delete(listOfBots, n1)
			default:
				panic("Unknown line: " + line)
			}
		}
		if len(listOfBots) == 0 && botNumber != -1 {
			break
		}
	}
}

func addToBin(n int, value int) {
	//fmt.Printf("Bin Number:%v, Value: %v\n", n, value)
	_, ok := listOfBins[n]
	if !ok {
		listOfBins[n] = make([]int, 0)
	}
	listOfBins[n] = append(listOfBins[n], value)
}
