package service

type BuzzService struct{}

func NewBuzzService() *BuzzService {
	return &BuzzService{}
}

func (b *BuzzService) Ping() string {
	return "Pong"
}
