package requestmodel

import responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"

type Order struct {
	UserID      string                        `json:"userid"`
	Address     string                        `json:"address" validate:"required,numeric"`
	Payment     string                        `json:"payment" validate:"required,alpha,uppercase"`
	OrderID     string                        `json:"-"`
	OrderStatus string                        `json:"-"`
	FinalPrice  uint                          `json:"-"`
	Cart        []responsemodel.CartInventory `json:"-"`
}

type OnlinePaymentVerification struct {
	PaymentID string `json:"payment_id" validate:"required"`
	OrderID   string `json:"order_id" validate:"required"`
	Signature string `json:"signature" validate:"required"`
}
