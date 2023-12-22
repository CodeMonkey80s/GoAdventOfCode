package day15_part1

import (
	"strings"

	"GoAdventOfCode/util"
)

type Ingredient struct {
	Cap      int
	Dur      int
	Fla      int
	Tex      int
	Calories int
}

func getAnswerForTwoIngredients(lines []string) int {

	list := make([]*Ingredient, 0)

	score := 0
	for _, line := range lines {
		in := getIngredient(line)
		list = append(list, in)
	}

	m := 100
	for a := 0; a <= m; a++ {
		for b := 0; b <= m-a; b++ {
			capacity := a*list[0].Cap + b*list[1].Cap
			durability := a*list[0].Dur + b*list[1].Dur
			flavor := a*list[0].Fla + b*list[1].Fla
			texture := a*list[0].Tex + b*list[1].Tex
			if capacity > 0 && durability > 0 && flavor > 0 && texture > 0 {
				total := capacity * durability * flavor * texture
				score = max(score, total)
			}
		}
	}

	return score
}

func getAnswerForFourIngredients(lines []string) int {

	list := make([]*Ingredient, 0)

	score := 0
	for _, line := range lines {
		in := getIngredient(line)
		list = append(list, in)
	}

	m := 100
	for a := 0; a <= m; a++ {
		for b := 0; b <= m; b++ {
			for c := 0; c <= m; c++ {
				for d := 0; d <= m; d++ {
					cap := a*list[0].Cap + b*list[1].Cap + c*list[2].Cap + d*list[3].Cap
					dur := a*list[0].Dur + b*list[1].Dur + c*list[2].Dur + d*list[3].Dur
					fla := a*list[0].Fla + b*list[1].Fla + c*list[2].Fla + d*list[3].Fla
					tex := a*list[0].Tex + b*list[1].Tex + c*list[2].Tex + d*list[3].Tex
					if a+b+c+d == 100 && cap > 0 && dur > 0 && fla > 0 && tex > 0 {
						total := cap * dur * fla * tex
						if total > score {
							score = total
						}
					}
				}
			}
		}
	}

	return score
}

func getIngredient(line string) *Ingredient {
	parts := strings.Fields(line)
	capacity := util.ConvertStringToInt(strings.TrimSuffix(parts[2], ","))
	durability := util.ConvertStringToInt(strings.TrimSuffix(parts[4], ","))
	flavor := util.ConvertStringToInt(strings.TrimSuffix(parts[6], ","))
	texture := util.ConvertStringToInt(strings.TrimSuffix(parts[8], ","))
	calories := util.ConvertStringToInt(strings.TrimSuffix(parts[10], ","))
	in := Ingredient{
		Cap:      capacity,
		Dur:      durability,
		Fla:      flavor,
		Tex:      texture,
		Calories: calories,
	}

	return &in
}
