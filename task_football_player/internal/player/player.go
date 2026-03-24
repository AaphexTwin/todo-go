package player

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func NewPlayer(name string, goals, misses, assists int) Player {
	player := Player{
		Name:    name,
		Goals:   goals,
		Misses:  misses,
		Assists: assists,
	}
	player.calculateRating()
	return player
}

func (p *Player) calculateRating() {
	if p.Misses == 0 {
		p.Rating = float64(p.Goals) + float64(p.Assists)/2.0
	} else {
		p.Rating = (float64(p.Goals) + float64(p.Assists)/2.0) / float64(p.Misses)
	}
}

// type Team struct {
//     players []Player
// }
