package author

import "context"

type IAuthorRepository interface {
	Find(ctx context.Context) (*[]Author, error)
	FindOne(ctx context.Context, filter interface{}) (*Author, error)
	Create(ctx context.Context, data Author) error
	Update(ctx context.Context, data interface{}, filter interface{}) error
	Delete(ctx context.Context, filter interface{}) error
}

type IAuthorService interface {
	Index() (*[]Author, error)
	Show(dto ShowDTO) (*Author, error)
	Store(dto StoreDTO) error
	Update(dto UpdateDTO) error
	Destroy(dto DestroyDTO) error
}
