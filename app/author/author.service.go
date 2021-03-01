package author

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthorService struct {
	authorRepository IAuthorRepository
	ctx              context.Context
}

func ProvideAuthorService(AuthorRepository AuthorRepository) AuthorService {
	return AuthorService{
		authorRepository: &AuthorRepository,
		ctx:              context.TODO(),
	}
}

func (s *AuthorService) Index() (*[]Author, error) {
	return s.authorRepository.Find(s.ctx)
}

func (s *AuthorService) Show(dto ShowDTO) (*Author, error) {
	ID, _ := primitive.ObjectIDFromHex(dto.ID)
	filter := bson.M{
		"_id": ID,
	}

	return s.authorRepository.FindOne(s.ctx, filter)
}

func (s *AuthorService) Store(dto StoreDTO) error {
	author := ToAuthor(dto)
	return s.authorRepository.Create(s.ctx, author)
}

func (s *AuthorService) Update(dto UpdateDTO) error {
	ID, _ := primitive.ObjectIDFromHex(dto.ID)
	filter := bson.M{
		"_id": ID,
	}
	data := bson.M{
		"$set": bson.M{
			"name":       dto.Name,
			"phone":      dto.Phone,
			"address":    dto.Address,
			"updated_at": time.Now(),
		},
	}

	return s.authorRepository.Update(s.ctx, data, filter)
}

func (s *AuthorService) Destroy(dto DestroyDTO) error {
	ID, _ := primitive.ObjectIDFromHex(dto.ID)
	filter := bson.M{
		"_id": ID,
	}

	return s.authorRepository.Delete(s.ctx, filter)
}
