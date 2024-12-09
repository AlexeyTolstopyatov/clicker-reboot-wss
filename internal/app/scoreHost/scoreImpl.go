package scoreHost

import (
	"clicker/internal/def/models"
	"golang.org/x/net/context"
)

// CountScore
// Counts global teams scores in game, queries statistics from DB
// Uses DBHost API and Postgre server's driver
func (s *ScoreHost) CountScore(ctx context.Context) (models.Score, error) {
	teamWhite, err := s.dbHostApi.SelectTeamByName(ctx, "white")
	if err != nil {
		return models.Score{}, err
	}

	teamBlack, err := s.dbHostApi.SelectTeamByName(ctx, "black")
	if err != nil {
		return models.Score{}, err
	}

	totalClicks := teamWhite.TotalClicks + teamBlack.TotalClicks
	whitePercents := float32(teamWhite.TotalClicks * 100 / totalClicks)
	blackPercents := float32(teamBlack.TotalClicks * 100 / totalClicks)

	return models.Score{
		WhiteScore:    teamWhite.TotalClicks,
		WhitePercents: whitePercents,
		BlackScore:    teamBlack.TotalClicks,
		BlackPercents: blackPercents,
	}, nil
}
