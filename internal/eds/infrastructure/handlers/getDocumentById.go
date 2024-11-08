package handlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	edsv1 "github.com/cfif1982/eds/protos/gen"
)

func (h *Handlers) GetDocumentByID(
	ctx context.Context,
	req *edsv1.GetDocumentByIDRequest,
) (*edsv1.GetDocumentByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}
