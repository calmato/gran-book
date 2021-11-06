package cmd

import (
	"github.com/calmato/gran-book/api/internal/information/application"
	"github.com/calmato/gran-book/api/internal/information/infrastructure/repository"
	dv "github.com/calmato/gran-book/api/internal/information/infrastructure/validation"
	"github.com/calmato/gran-book/api/internal/information/interface/server"
	rv "github.com/calmato/gran-book/api/internal/information/interface/validation"
	"github.com/calmato/gran-book/api/pkg/database"
	"github.com/calmato/gran-book/api/pkg/datetime"
	"github.com/calmato/gran-book/api/proto/information"
)

type registry struct {
	inquiry information.InquiryServiceServer
}

func newRegistry(db *database.Client) *registry {
	inquiryApplication, inquiryRequestValidation := inquiryInjector(db)

	return &registry{
		inquiry: server.NewInquiryServer(inquiryRequestValidation, inquiryApplication),
	}
}

func inquiryInjector(db *database.Client) (application.InquiryApplication, rv.InquiryRequestValidation) {
	ir := repository.NewInquiryRepository(db, datetime.Now)
	idv := dv.NewInquiryDomainVaildation()
	ia := application.NewInquiryApplication(idv, ir)

	irv := rv.NewInquiryRequestValidation()

	return ia, irv
}
