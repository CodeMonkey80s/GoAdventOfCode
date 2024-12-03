package day21_part2

type Warrior struct {
	hp     int
	damage int
	armor  int
}

type Item struct {
	name   string
	cost   int
	damage int
	armor  int
}

var weapons = []Item{
	{
		name:   "Dagger",
		cost:   8,
		damage: 4,
		armor:  0,
	},
	{
		name:   "Shortsword",
		cost:   10,
		damage: 5,
		armor:  0,
	},
	{
		name:   "Warhammer",
		cost:   25,
		damage: 6,
		armor:  0,
	},
	{
		name:   "Longsword",
		cost:   40,
		damage: 7,
		armor:  0,
	},
	{
		name:   "Greataxe",
		cost:   74,
		damage: 8,
		armor:  0,
	},
}

var armors = []Item{
	{
		name:   "Leather",
		cost:   13,
		damage: 0,
		armor:  1,
	},
	{
		name:   "Chainmail",
		cost:   31,
		damage: 0,
		armor:  2,
	},
	{
		name:   "Splintmail",
		cost:   53,
		damage: 0,
		armor:  3,
	},
	{
		name:   "Bandedmail",
		cost:   75,
		damage: 0,
		armor:  4,
	},
	{
		name:   "Platemail",
		cost:   102,
		damage: 0,
		armor:  5,
	},
	{
		name:   "[None]",
		cost:   0,
		damage: 0,
		armor:  0,
	},
}

var rings = []Item{
	{
		name:   "Damage +1",
		cost:   25,
		damage: 1,
		armor:  0,
	},
	{
		name:   "Damage +2",
		cost:   50,
		damage: 2,
		armor:  0,
	},
	{
		name:   "Damage +3",
		cost:   100,
		damage: 3,
		armor:  0,
	},
	{
		name:   "Defense +1",
		cost:   20,
		damage: 0,
		armor:  1,
	},
	{
		name:   "Defense +2",
		cost:   40,
		damage: 0,
		armor:  2,
	},
	{
		name:   "Defense +3",
		cost:   80,
		damage: 0,
		armor:  3,
	},
	{
		name:   "[None]",
		cost:   0,
		damage: 0,
		armor:  0,
	},
}

func getAnswer(player Warrior, boss Warrior) int {

	ringsIds := []int{0, 1, 2, 3, 4, 5, 6}

	maxCost := 0
	for _, weapon := range weapons {
		for _, armor := range armors {
			allPermutations := getAllPermutations(ringsIds, 3)
			for _, perms := range allPermutations {
				for _, nums := range perms {
					p := Warrior{
						hp:     player.hp,
						armor:  weapon.armor + armor.armor,
						damage: weapon.damage + armor.damage,
					}
					cost := 0
					cost += weapon.cost + armor.cost

					for _, id := range nums {
						p.damage += rings[id].damage
						p.armor += rings[id].armor
						cost += rings[id].cost
					}

					if fight(p, boss) == false {
						maxCost = max(cost, maxCost)
					}
				}
			}

		}
	}

	return maxCost
}

func fight(player Warrior, boss Warrior) bool {
	for {
		if player.damage-boss.armor == 0 {
			boss.hp--
		} else {
			boss.hp -= player.damage - boss.armor
		}

		if boss.hp <= 0 {
			return true
		}

		if boss.damage-player.armor == 0 {
			player.hp--
		} else {
			player.hp -= boss.damage - player.armor
		}

		if player.hp <= 0 {
			return false
		}
	}
}

func permutations(nums []int, size int) [][]int {
	if size == 0 {
		return [][]int{{}}
	}
	if len(nums) < size {
		return nil
	}

	var result [][]int
	used := make([]bool, len(nums))
	var backtrack func(path []int)

	backtrack = func(path []int) {
		if len(path) == size {
			result = append(result, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			backtrack(append(path, nums[i]))
			used[i] = false
		}
	}

	backtrack([]int{})
	return result
}

// getAllPermutations generates permutations from size 0 up to the max size of nums.
func getAllPermutations(nums []int, m int) [][][]int {
	var allPermutations [][][]int
	for size := 0; size <= m; size++ {
		allPermutations = append(allPermutations, permutations(nums, size))
	}
	return allPermutations
}
