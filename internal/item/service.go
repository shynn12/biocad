package item

import (
	"context"

	"github.com/shynn12/biocad/pkg/logging"
)

type Service interface {
	GetAllItems(ctx context.Context, page, perPage int) []*Item
	CreateItem(ctx context.Context, dto ItemDTO) string
	GetOneByGuid(ctx context.Context, guid string) *Item
}

type service struct {
	storage Storage
	logger  *logging.Logger
}

func NewService(storage Storage, logger *logging.Logger) Service {
	return &service{storage: storage, logger: logger}
}

func (s *service) GetAllItems(ctx context.Context, page, perPage int) []*Item {
	i, err := s.storage.FindAll(ctx, page, perPage)
	if err != nil {
		s.logger.Error(err)
	}
	return i
}

func (s *service) GetOneByGuid(ctx context.Context, guid string) *Item {
	item, err := s.storage.FindOne(ctx, guid)
	if err != nil {
		s.logger.Error(err)
	}
	return item
}

func (s *service) CreateItem(ctx context.Context, dto ItemDTO) string {
	id, err := s.storage.Create(ctx, dto)
	if err != nil {
		s.logger.Error(err)
	}
	return id
}
