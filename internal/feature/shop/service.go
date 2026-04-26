package shop_service

import shop_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/shop/ports/out"

type ShopService struct {
	shopRepository shop_ports_out.ShopRepository
}

func NewShopService(shopRepository shop_ports_out.ShopRepository) ShopService {
	return ShopService{
		shopRepository: shopRepository,
	}
}
