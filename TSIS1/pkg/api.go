package pkg

type Response struct {
	Players []Football json:"players"
}

type Football struct {
	FootballTeam string json:"football_team"
	FullName     string json:"full_name"
	Number       int    json:"number"
}

func PrepareResponse() []Football {
	var players []Football

	var player Football
	player.FootballTeam = "Barcelona"
	player.FullName = "Lionel Messi"
	player.Number = 10
	players = append(players, player)

	player.FootballTeam = "Real Madrid"
	player.FullName = "Luka Modrich"
	player.Number = 10
	players = append(players, player)

	player.FootballTeam = "Juventus"
	player.FullName = "Zinedine Zidane"
	player.Number = 10
	players = append(players, player)
	return players
}