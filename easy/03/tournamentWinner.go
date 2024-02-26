package easy03

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	resultsByTeam := map[string]int{}

	for index, teams := range competitions {
		result := results[index]
		homeTeam := teams[0]
		awayTeam := teams[1]

		if result == HOME_TEAM_WON {
			resultsByTeam[homeTeam] += 3
		} else {
			resultsByTeam[awayTeam] += 3
		}
	}

	winner := ""
	maxValue := -1
	for k, v := range resultsByTeam {
		if v > maxValue {
			winner = k
			maxValue = v
		}
	}

	return winner
}
