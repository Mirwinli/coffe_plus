package shop_ports_out

type CangeShopStatusParams struct {
	Status string
}

func NewCangeShopStatusParams(status string) CangeShopStatusParams {
	return CangeShopStatusParams{
		Status: status,
	}
}
