package coffee

type ICoffeeDB interface {
}

type Service struct {
	db ICoffeeDB
}

func NewService(db ICoffeeDB) *Service {
	return &Service{
		db: db,
	}
}
