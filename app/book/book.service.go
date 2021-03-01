package book

import (
	"context"
	"go-wire-mongo/app/author"
	"go-wire-mongo/shared"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookService struct {
	bookRepository IBookRepository
	authorService  author.IAuthorService
	ctx            context.Context
}

func ProvideBookService(BookRepository BookRepository, AuthorService author.AuthorService) BookService {
	return BookService{
		bookRepository: &BookRepository,
		authorService:  &AuthorService,
		ctx:            context.TODO(),
	}
}

func (s *BookService) Index(dto IndexDTO) (*[]BookDTO, error) {
	ID, _ := primitive.ObjectIDFromHex(dto.AuthorID)
	filter := bson.M{
		"author_id": ID,
	}

	author, err := s.getAuthor(dto.AuthorID)
	if err != nil {
		return nil, err
	}

	data, err := s.bookRepository.Find(s.ctx, filter)
	if err != nil {
		return nil, err
	}

	books := ToBookDTOs(*data, *author)

	return &books, nil
}

func (s *BookService) Show(dto ShowDTO) (*BookDTO, error) {
	ID, _ := primitive.ObjectIDFromHex(dto.ID)
	AuthorID, _ := primitive.ObjectIDFromHex(dto.AuthorID)
	filter := bson.M{
		"_id":       ID,
		"author_id": AuthorID,
	}

	author, err := s.getAuthor(dto.AuthorID)
	if err != nil {
		return nil, err
	}

	data, err := s.bookRepository.FindOne(s.ctx, filter)
	if err != nil {
		return nil, err
	}

	book := ToBookDTO(*data, *author)

	return &book, nil
}

func (s *BookService) Store(dto StoreDTO) error {
	book := ToBook(dto)
	return s.bookRepository.Create(s.ctx, book)
}

func (s *BookService) Update(dto UpdateDTO) error {
	ID, _ := primitive.ObjectIDFromHex(dto.ID)
	AuthorID, _ := primitive.ObjectIDFromHex(dto.AuthorID)
	filter := bson.M{
		"_id": ID,
	}
	data := bson.M{
		"$set": bson.M{
			"author_id":  AuthorID,
			"title":      dto.Title,
			"year":       dto.Year,
			"updated_at": time.Now(),
		},
	}

	return s.bookRepository.Update(s.ctx, data, filter)
}

func (s *BookService) Destroy(dto DestroyDTO) error {
	ID, _ := primitive.ObjectIDFromHex(dto.ID)
	AuthorID, _ := primitive.ObjectIDFromHex(dto.AuthorID)
	filter := bson.M{
		"_id":       ID,
		"author_id": AuthorID,
	}

	return s.bookRepository.Delete(s.ctx, filter)
}

func (s *BookService) getAuthor(authorID string) (*author.Author, error) {
	showDTO := ToAuthorShowDTO(authorID)
	author, err := s.authorService.Show(showDTO)
	if err != nil {
		return nil, shared.ErrNilAuthor
	}

	return author, nil
}
