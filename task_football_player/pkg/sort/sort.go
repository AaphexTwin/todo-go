package sort

import (
	"slices"

	"github.com/AaphexTwin/task_football_player/internal/player"
)

func goalsSort(players []player.Player) []player.Player {
	playersCopy := slices.Clone(players)
	slices.SortFunc(playersCopy, func(a, b player.Player) int {
		switch {
		case a.Goals > b.Goals:
			return -1
		case a.Goals < b.Goals:
			return 1
		default:
			if a.Name > b.Name {
				return 1
			} else if a.Name < b.Name {
				return -1
			}
			return 0
		}
	})
	return playersCopy
}

func ratingSort(players []player.Player) []player.Player {
	playersCopy := slices.Clone(players)
	slices.SortFunc(playersCopy, func(a, b player.Player) int {
		switch {
		case a.Rating > b.Rating:
			return -1
		case a.Rating < b.Rating:
			return 1
		default:
			if a.Name > b.Name {
				return 1
			} else if a.Name < b.Name {
				return -1
			}
			return 0
		}
	})
	return playersCopy

}

func gmSort(players []player.Player) []player.Player {
	playersCopy := slices.Clone(players)
	slices.SortFunc(playersCopy, func(a, b player.Player) int {
		if result := misAndGoals(a, b); result != 0 {
			return result
		}
		if a.Name > b.Name {
			return 1
		} else if a.Name < b.Name {
			return -1
		}
		return 0
	})

	return playersCopy
}

func misAndGoals(a, b player.Player) int {
	if a.Misses == 0 && b.Misses == 0 {
		if a.Goals > b.Goals {
			return -1
		} else if a.Goals < b.Goals {
			return 1
		} else {
			return 0
		}
	}
	if a.Misses == 0 {
		return -1
	} else if b.Misses == 0 {
		return 1
	} else {
		if float64(a.Goals)/float64(a.Misses) > float64(b.Goals)/float64(b.Misses) {
			return -1
		} else if float64(a.Goals)/float64(a.Misses) < float64(b.Goals)/float64(b.Misses) {
			return 1
		} else {
			return 0
		}
	}

}
