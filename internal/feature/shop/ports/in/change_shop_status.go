package shop_ports_in

type ChangeShopStatusParams struct {
	Status string
}

func NewCangeShopStatusParams(status string) ChangeShopStatusParams {
	return ChangeShopStatusParams{
		Status: status,
	}
}
