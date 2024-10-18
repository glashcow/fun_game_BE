package game

import (
	"github.com/google/uuid"
)

type BingoCardGenerator interface {
	GenerateCard() map[Coordinate]Number
}

type Service struct {
	repo               Repository
	bingoCardGenerator BingoCardGenerator
}

func NewService(repo Repository, bingoCardGenerator BingoCardGenerator) *Service {
	return &Service{
		repo:               repo,
		bingoCardGenerator: bingoCardGenerator,
	}
}

func (s *Service) NewGame(playerIDs []string) (*Game, error) {
	bingoCards := make([]*BingoCard, len(playerIDs))

	for i := 0; i < len(playerIDs); i++ {
		card := &BingoCard{
			ID:       uuid.NewString(),
			PlayerID: playerIDs[i],
			Numbers:  s.bingoCardGenerator.GenerateCard(),
		}
		bingoCards[i] = card
	}

	return &Game{
		ID:         uuid.NewString(),
		PlayerIDs:  playerIDs,
		BingoCards: bingoCards,
		Calls:      nil,
	}, nil
}
