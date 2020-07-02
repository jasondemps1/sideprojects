package elo

import "testing"

func TestElo(t *testing.T) {
	cases := []struct{
		opponentElo, wins, losses, games, want int
	}{
		{1000, 1, 0, 1, 1400},
		{2000, 2, 0, 2, 1400},
		{1000, 0, 0, 1, 1000},
	}

	for _, c := range cases {
		got := DetermineElo(c.opponentElo, c.wins, c.losses, c.games)

		if got != c.want {
			t.Errorf("DetermineElo(%q, %q, %q, %q) == %q, want %q", c.opponentElo, c.wins, c.losses, c.games, got, c.want)
		}
	}
}
