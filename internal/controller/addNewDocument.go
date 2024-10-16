package controller

import (
	edsv1 "github.com/cfif1982/eds/protos/gen"
	"github.com/google/uuid"
)

func (c *Controller) AddNewDocument(creatorID string) (*edsv1.AddNewDocumentResponse, error) {
	// конвертируем данные из grpc в domain
	// это коммент на тот случай, когда будем больше данных получать
	// сейчас просто конвертируем строку в uuid
	creatorUUID, err := uuid.Parse(creatorID)
	if err != nil {
		return nil, err
	}

	qrCode, err := c.docUseCases.Add(creatorUUID)

	if err != nil {
		return nil, err
	}

	// конвертируем данные из domain в grpc для ответа.
	result := edsv1.AddNewDocumentResponse{
		QrCode: qrCode,
	}

	return &result, nil
}
