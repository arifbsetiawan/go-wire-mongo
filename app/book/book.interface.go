package book

import "context"

type IBookRepository interface {
	Find(ctx context.Context, filter interface{}) (*[]Book, error)
	FindOne(ctx context.Context, filter interface{}) (*Book, error)
	Create(ctx context.Context, data Book) error
	Update(ctx context.Context, data interface{}, filter interface{}) error
	Delete(ctx context.Context, filter interface{}) error
}

type IBookService interface {
	Index(dto IndexDTO) (*[]BookDTO, error)
	Show(dto ShowDTO) (*BookDTO, error)
	Store(dto StoreDTO) error
	Update(dto UpdateDTO) error
	Destroy(dto DestroyDTO) error
}
