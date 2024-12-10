package document

import (
	"context"
	"fmt"

	"github.com/cfif1982/eds/internal/models/response"
	"github.com/google/uuid"
)

func (s *Services) GetById(
	ctx context.Context,
	documentID uuid.UUID,
) (*response.Document, error) {

	// находим документ по его id
	doc, err := s.docRepo.GetByID(ctx, documentID)

	if err != nil {
		return nil, fmt.Errorf("GetById() service error: %w", err)
	}

	// готовим модель для ответа
	// узнаем данные создаетля
	creator, err := s.userRepo.GetByID(ctx, doc.CreatorID)
	if err != nil {
		return nil, fmt.Errorf("GetById() service error: %w", err)
	}

	// узнаем данные подписантов
	signers := make([]response.User, 0, len(doc.SignersID))

	for _, id := range doc.SignersID {
		u, err := s.userRepo.GetByID(ctx, id)

		if err != nil {
			return nil, fmt.Errorf("GetById() service error: %w", err)
		}

		signers = append(signers, response.User{
			ID:    id,
			Email: u.Email,
			Name:  u.Name,
		})
	}

	// узнаем данные файлов
	files := make([]string, 0, len(doc.Files))

	for _, f := range doc.Files {
		files = append(files, f.FileName)
	}

	result := &response.Document{
		ID: doc.ID,
		Creator: response.User{
			ID:    creator.ID,
			Email: creator.Email,
			Name:  creator.Name,
		},
		Signers:  signers,
		FileURLs: files,
	}

	return result, nil
}
