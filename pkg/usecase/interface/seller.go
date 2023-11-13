package interfaceUseCase

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type ISellerUseCase interface {
	SellerSignup(*requestmodel.SellerSignup) (*responsemodel.SellerSignupRes, error)
	SellerLogin(*requestmodel.SellerLogin) (*responsemodel.SellerLoginRes, error)
	GetAllSellers(string, string) (*[]responsemodel.SellerDetails, *int, error)
	BlockSeller(string) error
	UnblockSeller(string) error
	GetAllPendingSellers(string, string) (*[]responsemodel.SellerDetails, error)
	FetchSingleVender(string) (*responsemodel.SellerDetails, error)
}
