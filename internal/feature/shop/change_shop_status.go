package shop_service

import (
	"context"
	"fmt"

	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	core_errors "github.com/Mirwinli/coffe_plus/internal/core/errors"
	shop_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/shop/ports/in"
	shop_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/shop/ports/out"
)

func (s *ShopService) CangeShopStatus(
	ctx context.Context,
	in shop_ports_in.ChangeShopStatusParams,
) error {
	if in.Status != domain.ShopClosed && in.Status != domain.ShopOpen {
		return fmt.Errorf(
			"status invalid: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	params := shop_ports_out.NewCangeShopStatusParams(in.Status)

	err := s.shopRepository.ChangeShopStatus(ctx, params)
	if err != nil {
		return fmt.Errorf(
			"change shop status: %w", err,
		)
	}

	return nil
}
