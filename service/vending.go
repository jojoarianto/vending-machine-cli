package service

import "github.com/jojoarianto/vending-machine-cli/model"

type InsertService interface {
	Insert(newCoin int64)
}

type insertService struct {
	coin model.Coin
}

func NewInsertService(coin model.Coin) InsertService {
	return &insertService{coin}
}

func (svc *insertService) Insert(newCoin int64)  {

}