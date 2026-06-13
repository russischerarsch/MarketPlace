package services

import (
	"context"
	"errors"
	"mini-ozon/intern/models/orders"
	"mini-ozon/intern/repositories"
)

type ChangeStatServ struct {
	repo *repositories.ChangeRep
}

func CreateChangeStatServ(repo *repositories.ChangeRep) *ChangeStatServ {
	return &ChangeStatServ{
		repo: repo,
	}
}
func (c ChangeStatServ) ChangeStatus(ctx context.Context, id int, status orders.Status) error {
	current, err := c.repo.GetStatus(id, ctx)
	if err != nil {
		return err
	}

	if !orders.CanTransition(current, status) {
		return errors.New("invalid status transition")
	}

	return c.repo.ChangeStatus(newStatus, id, ctx)
}
