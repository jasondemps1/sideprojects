package elo

// DetermineElo - Elo system described: https://en.wikipedia.org/wiki/Elo_rating_system
func DetermineElo(opponentElo, wins, losses, games int) int {
	return (opponentElo + 400 * (wins - losses)) / games
}
