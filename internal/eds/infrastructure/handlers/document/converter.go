package document

import (
	"github.com/cfif1982/eds/internal/models/response"
	edsv1 "github.com/cfif1982/eds/protos/gen"
)

func (h *Handlers) documentToGetDocumentByIDResponse(doc *response.Document) *edsv1.GetDocumentByIDResponse {
	// находим данные для creator
	user := h.services.

	// создаем модель 
	creator := &edsv1.User{
		UserId: ,
	}
	// получаем слайс моделей user для подписантов документа
	signers := make([]*models.User, 0, len(doc.Signers))
	for _, id := range doc.Signers {
		// signer, err := h.services.
		signers = append(signers, signer)
	}

	// получаем слайс моделей file для файлов документа
	files := make([]*models.File, 0, len(doc.Files))
	for _, id := range doc.Files {
		// file, err := h.services.
		files = append(files, file)
	}

	result := &edsv1.GetDocumentByIDResponse{
		DocumentId: doc.ID.String(),
		Creator:    creator,
		Signers:    signers,
		FilesUrl:   files,
	}
}
