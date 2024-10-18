package game

type Repository interface {
	Save(game *Game) error
	GetByID(id string) (*Game, error)
	Delete(id string) error
}
