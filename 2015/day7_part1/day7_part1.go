package day7_part1

import (
	"fmt"
	"strconv"
	"strings"
)

type connections map[string]uint16

var wires = make(connections)

func getAnswer(lines []string, register string) (connections, uint16) {
	clear(wires)
	output := uint16(0)
outer:
	for {
		for _, line := range lines {
			doCommand(line)
		}

		v, ok := wires[register]
		if ok {
			fmt.Printf("Register: %q, Output: %d\n", register, v)
			output = v
			break outer
		}
	}

	return wires, output
}

func doCommand(line string) {

	idx := strings.Index(line, "->")
	lhs := strings.TrimSpace(line[:idx])
	rhs := strings.TrimSpace(line[idx+3:])

	opNOT := strings.Contains(lhs, "NOT")
	opOR := strings.Contains(lhs, "OR")
	opAND := strings.Contains(lhs, "AND")
	opLSHIFT := strings.Contains(lhs, "LSHIFT")
	opRSHIFT := strings.Contains(lhs, "RSHIFT")

	if opOR || opAND || opLSHIFT || opRSHIFT {

		if _, ok := wires[rhs]; ok {
			return
		}

		parts := []string{}
		if opOR {
			parts = strings.Split(lhs, " OR ")
		}
		if opAND {
			parts = strings.Split(lhs, " AND ")
		}
		if opLSHIFT {
			parts = strings.Split(lhs, " LSHIFT ")
		}
		if opRSHIFT {
			parts = strings.Split(lhs, " RSHIFT ")
		}

		value1, err := getNumberFromWire(parts[0])
		if err != nil {
			return
		}
		value2, err := getNumberFromWire(parts[1])
		if err != nil {
			return
		}

		value := uint16(0)
		switch {
		case opOR:
			value = value1 | value2
		case opAND:
			value = value1 & value2
		case opLSHIFT:
			value = value1 << value2
		case opRSHIFT:
			value = value1 >> value2
		}

		wires[rhs] = value
		return
	}

	if opNOT {
		if _, ok := wires[rhs]; ok {
			return
		}
		parts := strings.Split(lhs, " ")
		register := strings.TrimSpace(parts[1])
		value, err := getNumberFromWire(register)
		if err != nil {
			return
		}

		if opNOT {
			value = ^value
		}
		wires[rhs] = value
		return
	}

	if _, ok := wires[rhs]; ok {
		return
	}
	value, err := getNumberFromWire(lhs)
	if err != nil {
		return
	}
	wires[rhs] = value
}

func getNumberFromWire(value string) (uint16, error) {
	number, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		wire, ok := wires[value]
		if !ok {
			return uint16(number), err
		}
		return wire, nil
	}
	return uint16(number), err
}
