package day21_part1

import (
	"fmt"
	"testing"
)

func Test_fight(t *testing.T) {
	var testCases = []struct {
		InputPlayer Warrior
		InputBoss   Warrior
		Output      bool
	}{
		{
			InputPlayer: Warrior{
				hp:     8,
				damage: 5,
				armor:  5,
			},
			InputBoss: Warrior{
				hp:     12,
				damage: 7,
				armor:  2,
			},
			Output: true,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := fight(tc.InputPlayer, tc.InputBoss)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		InputPlayer Warrior
		InputBoss   Warrior
		Output      int
	}{
		{
			InputPlayer: Warrior{
				hp: 100,
			},
			InputBoss: Warrior{
				hp:     109,
				damage: 8,
				armor:  2,
			},
			Output: 111,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.InputPlayer, tc.InputBoss)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
