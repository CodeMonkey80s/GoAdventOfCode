package day5_part1

import (
	"fmt"
	"strings"

	"GoAdventOfCode/util"
)

const (
	_ = iota
	scanSeedToSoil
	scanSoilToFertilizer
	scanFertilizerToWater
	scanWaterToLight
	scanLightToTemperature
	scanTemperatureToHumidity
	scanHumidityToLocation
)

type Mapping struct {
	DstRangeStart int
	SrcRangeStart int
	RangeLength   int
}

func lowestLocationNumber(lines []string) int {

	scannerFlag := 0

	seeds := make([]int, 0)
	seedToSoil := make([]Mapping, 0)
	soilToFertilizer := make([]Mapping, 0)
	fertilizerToWater := make([]Mapping, 0)
	waterToLight := make([]Mapping, 0)
	lightToTemperature := make([]Mapping, 0)
	temperatureToHumidity := make([]Mapping, 0)
	humidityToLocation := make([]Mapping, 0)

loop:
	for _, line := range lines {
		if line == "" {
			continue
		}
		switch {
		case strings.HasPrefix(line, "seeds: "):
			ids := strings.Fields(line[7:])
			for _, id := range ids {
				val := util.ConvertStringToInt(id)
				seeds = append(seeds, val)
			}
			fmt.Printf("ids: %v\n", ids)
		case strings.HasPrefix(line, "seed-to-soil map:"):
			scannerFlag = scanSeedToSoil
			continue loop
		case strings.HasPrefix(line, "soil-to-fertilizer map:"):
			scannerFlag = scanSoilToFertilizer
			continue loop
		case strings.HasPrefix(line, "fertilizer-to-water map:"):
			scannerFlag = scanFertilizerToWater
			continue loop
		case strings.HasPrefix(line, "water-to-light map:"):
			scannerFlag = scanWaterToLight
			continue loop
		case strings.HasPrefix(line, "light-to-temperature map:"):
			scannerFlag = scanLightToTemperature
			continue loop
		case strings.HasPrefix(line, "temperature-to-humidity map:"):
			scannerFlag = scanTemperatureToHumidity
			continue loop
		case strings.HasPrefix(line, "humidity-to-location map:"):
			scannerFlag = scanHumidityToLocation
			continue loop
		default:
		}
		switch scannerFlag {
		case scanSeedToSoil:
			seedToSoil = fillMapping(line, seedToSoil)
		case scanSoilToFertilizer:
			soilToFertilizer = fillMapping(line, soilToFertilizer)
		case scanFertilizerToWater:
			fertilizerToWater = fillMapping(line, fertilizerToWater)
		case scanWaterToLight:
			waterToLight = fillMapping(line, waterToLight)
		case scanLightToTemperature:
			lightToTemperature = fillMapping(line, lightToTemperature)
		case scanTemperatureToHumidity:
			temperatureToHumidity = fillMapping(line, temperatureToHumidity)
		case scanHumidityToLocation:
			humidityToLocation = fillMapping(line, humidityToLocation)
		default:
		}
	}

	lowest := 1 << 32
	for _, seed := range seeds {
		soil := corresponds(seed, seedToSoil)
		fertilizer := corresponds(soil, soilToFertilizer)
		water := corresponds(fertilizer, fertilizerToWater)
		light := corresponds(water, waterToLight)
		temperature := corresponds(light, lightToTemperature)
		humidity := corresponds(temperature, temperatureToHumidity)
		location := corresponds(humidity, humidityToLocation)
		lowest = min(location, lowest)
	}

	return lowest
}

func corresponds(n int, mapping []Mapping) int {
	for _, m := range mapping {
		srcStart := m.SrcRangeStart
		srcEnd := m.SrcRangeStart + m.RangeLength
		dstStart := m.DstRangeStart
		if n >= srcStart && n <= srcEnd {
			d := n - srcStart
			return dstStart + d
		}
	}
	return n
}

func fillMapping(line string, sl []Mapping) []Mapping {
	ids := strings.Fields(line)
	item := Mapping{
		DstRangeStart: util.ConvertStringToInt(ids[0]),
		SrcRangeStart: util.ConvertStringToInt(ids[1]),
		RangeLength:   util.ConvertStringToInt(ids[2]),
	}
	return append(sl, item)
}
