package day4_part1

import (
	"fmt"
	"reflect"
	"testing"

	"GoAdventOfCode/util"
)

func Test_sortInput(t *testing.T) {
	var testCases = []struct {
		Input  []string
		Output []string
	}{
		{
			Input: []string{
				"[1518-11-01 00:55] wakes up",
				"[1518-11-04 00:36] falls asleep",
				"[1518-11-01 23:58] Guard #99 begins shift",
				"[1518-11-02 00:40] falls asleep",
				"[1518-11-01 00:05] falls asleep",
				"[1518-11-01 00:30] falls asleep",
				"[1518-11-02 00:50] wakes up",
				"[1518-11-03 00:05] Guard #10 begins shift",
				"[1518-11-04 00:02] Guard #99 begins shift",
				"[1518-11-01 00:00] Guard #10 begins shift",
				"[1518-11-05 00:55] wakes up",
				"[1518-11-01 00:25] wakes up",
				"[1518-11-03 00:24] falls asleep",
				"[1518-11-03 00:29] wakes up",
				"[1518-11-05 00:45] falls asleep",
				"[1518-11-04 00:46] wakes up",
				"[1518-11-05 00:03] Guard #99 begins shift",
			},
			Output: []string{
				"[1518-11-01 00:00] Guard #10 begins shift",
				"[1518-11-01 00:05] falls asleep",
				"[1518-11-01 00:25] wakes up",
				"[1518-11-01 00:30] falls asleep",
				"[1518-11-01 00:55] wakes up",
				"[1518-11-01 23:58] Guard #99 begins shift",
				"[1518-11-02 00:40] falls asleep",
				"[1518-11-02 00:50] wakes up",
				"[1518-11-03 00:05] Guard #10 begins shift",
				"[1518-11-03 00:24] falls asleep",
				"[1518-11-03 00:29] wakes up",
				"[1518-11-04 00:02] Guard #99 begins shift",
				"[1518-11-04 00:36] falls asleep",
				"[1518-11-04 00:46] wakes up",
				"[1518-11-05 00:03] Guard #99 begins shift",
				"[1518-11-05 00:45] falls asleep",
				"[1518-11-05 00:55] wakes up",
			},
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", tc.Input[0], tc.Output[0])
		t.Run(label, func(t *testing.T) {
			output := sortInput(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input  []string
		Output int
	}{
		{
			Input: []string{
				"[1518-11-01 00:55] wakes up",
				"[1518-11-04 00:36] falls asleep",
				"[1518-11-01 23:58] Guard #99 begins shift",
				"[1518-11-02 00:40] falls asleep",
				"[1518-11-01 00:05] falls asleep",
				"[1518-11-01 00:30] falls asleep",
				"[1518-11-02 00:50] wakes up",
				"[1518-11-03 00:05] Guard #10 begins shift",
				"[1518-11-04 00:02] Guard #99 begins shift",
				"[1518-11-01 00:00] Guard #10 begins shift",
				"[1518-11-05 00:55] wakes up",
				"[1518-11-01 00:25] wakes up",
				"[1518-11-03 00:24] falls asleep",
				"[1518-11-03 00:29] wakes up",
				"[1518-11-05 00:45] falls asleep",
				"[1518-11-04 00:46] wakes up",
				"[1518-11-05 00:03] Guard #99 begins shift",
			},
			Output: 240,
		},
	}
	lines := util.LoadInputFile("../inputs/day4_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 14346,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", tc.Input[0], tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
