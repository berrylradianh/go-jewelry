package transactions

import (
	up "github.com/berrylradianh/go-jewelry/modules/usecase/products"
	ut "github.com/berrylradianh/go-jewelry/modules/usecase/transactions"
)

type Handler struct {
	Usecase *ut.Usecase
}
type ProductHandler struct {
	Usecase *up.Usecase
}
