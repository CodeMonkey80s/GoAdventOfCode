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

type Bin struct {
	Value int
}

var listOfBots = map[int]*Bot{}
var listOfBins = map[int]*Bin{}

func NewBot(n int) *Bot {
	return &Bot{
		Number: n,
		Chips:  make([]int, 2),
	}
}

var botNumber = -1
var compareValue1 = -1
var compareValue2 = -1

func GetBot(n int) *Bot {
	b, ok := listOfBots[n]
	if !ok {
		return NewBot(n)
	}
	return b
}

func getAnswer(lines []string, value1 int, value2 int) int {

	listOfBots = make(map[int]*Bot)
	listOfBins = make(map[int]*Bin)

	botNumber = -1
	compareValue1 = value1
	compareValue2 = value2

	for {
		for _, line := range lines {
			parts := strings.Fields(line)
			//fmt.Printf("parts: %v\n", parts)
			switch {
			case parts[0] == "value":
				value := util.ConvertStringToInt(parts[1])
				number := util.ConvertStringToInt(parts[5])
				b := GetBot(number)
				giveChipToBot(b, value)
				listOfBots[number] = b
			case parts[0] == "bot" && parts[5] == "output" && parts[10] == "bot":
				n1 := util.ConvertStringToInt(parts[1])
				out := util.ConvertStringToInt(parts[6])
				n2 := util.ConvertStringToInt(parts[11])
				b1 := GetBot(n1)
				b2 := GetBot(n2)
				giveLowChipToOutputAndHighChipToBot(b1, out, b2)
				listOfBots[n1] = b1
				listOfBots[n2] = b2
			case parts[0] == "bot" && parts[5] == "output" && parts[10] == "output":
				number := util.ConvertStringToInt(parts[1])
				out1 := util.ConvertStringToInt(parts[6])
				out2 := util.ConvertStringToInt(parts[11])
				b := GetBot(number)
				giveLowChipToOutputAndHighChipToOutput(b, out1, out2)
				listOfBots[number] = b
			case parts[0] == "bot" && parts[5] == "bot" && parts[10] == "bot":
				n1 := util.ConvertStringToInt(parts[1])
				n2 := util.ConvertStringToInt(parts[6])
				n3 := util.ConvertStringToInt(parts[11])
				b1 := GetBot(n1)
				b2 := GetBot(n2)
				b3 := GetBot(n3)
				giveLowChipToBotAndHighChipToBot(b1, b2, b3)
				listOfBots[n1] = b1
				listOfBots[n2] = b2
				listOfBots[n3] = b3

			default:
				panic("Unknown line: " + line)
			}
		}
		if botNumber != -1 {
			break
		}
	}

	//util.PrintMap("listOfBots", listOfBots)
	//util.PrintMap("listOfBins", listOfBins)

	/*

		bot 2 => compare 5 with 2

		*** listOfBots ***
		2 => "&{Chips:[0 0]}"
		1 => "&{Chips:[0 0]}"
		0 => "&{Chips:[0 0]}"
		*** listOfBins ***
		2 => "&{Value:3}"
		1 => "&{Value:2}"
		0 => "&{Value:5}"

	*/

	return botNumber
}

func giveChipToBot(b *Bot, value int) {

	if value == compareValue1 && slices.Contains(b.Chips, compareValue2) {
		botNumber = b.Number
		fmt.Printf("FOUND: %+v\n", b)
	}

	if b.Chips[0] == 0 && b.Chips[1] == 0 {
		b.Chips[1] = value
		return
	}

	if b.Chips[0] == 0 && value > b.Chips[1] {
		b.Chips[0] = b.Chips[1]
		b.Chips[1] = value
		return
	}

	if b.Chips[0] == 0 && value <= b.Chips[1] {
		b.Chips[0] = value
		return
	}

}

func giveLowChipToBotAndHighChipToBot(b1 *Bot, b2 *Bot, b3 *Bot) {

	giveChipToBot(b2, b1.Chips[0])
	b1.Chips[0] = 0

	giveChipToBot(b3, b1.Chips[1])
	b1.Chips[1] = 0

}

func giveLowChipToOutputAndHighChipToBot(b1 *Bot, out int, b2 *Bot) {

	bin, ok := listOfBins[out]
	if !ok {
		bin = &Bin{
			Value: b1.Chips[0],
		}
		listOfBins[out] = bin
	} else {
		bin.Value = b1.Chips[0]
	}
	b1.Chips[0] = 0

	giveChipToBot(b2, b1.Chips[1])
	b1.Chips[1] = 0

}
func giveLowChipToOutputAndHighChipToOutput(b *Bot, out1 int, out2 int) {
	bin1, ok1 := listOfBins[out1]
	if !ok1 {
		bin1 = &Bin{
			Value: b.Chips[0],
		}
		listOfBins[out1] = bin1
	} else {
		bin1.Value = b.Chips[0]
	}
	b.Chips[0] = 0

	bin2, ok2 := listOfBins[out2]
	if !ok2 {
		bin2 = &Bin{
			Value: b.Chips[1],
		}
		listOfBins[out2] = bin2
	} else {
		bin2.Value = b.Chips[1]
	}
	b.Chips[1] = 0
}
