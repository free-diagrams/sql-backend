package usecase

import "context"

type UseCase interface {
	GetColors(ctx context.Context) ([]GetColorsResponseItem, error)
}

type GetColorsResponseItem struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	HexCode string `json:"hexCode"`
}
