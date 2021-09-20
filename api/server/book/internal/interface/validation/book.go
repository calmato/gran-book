package validation

import (
	pb "github.com/calmato/gran-book/api/server/book/proto/service/book"
)

type BookRequestValidation interface {
	ListBookshelf(req *pb.ListBookshelfRequest) error
	ListBookReview(req *pb.ListBookReviewRequest) error
	ListUserReview(req *pb.ListUserReviewRequest) error
	ListUserMonthlyResult(req *pb.ListUserMonthlyResultRequest) error
	MultiGetBooks(req *pb.MultiGetBooksRequest) error
	GetBook(req *pb.GetBookRequest) error
	GetBookByIsbn(req *pb.GetBookByIsbnRequest) error
	GetBookshelf(req *pb.GetBookshelfRequest) error
	GetReview(req *pb.GetReviewRequest) error
	CreateBook(req *pb.CreateBookRequest) error
	UpdateBook(req *pb.UpdateBookRequest) error
	ReadBookshelf(req *pb.ReadBookshelfRequest) error
	ReadingBookshelf(req *pb.ReadingBookshelfRequest) error
	StackedBookshelf(req *pb.StackedBookshelfRequest) error
	WantBookshelf(req *pb.WantBookshelfRequest) error
	ReleaseBookshelf(req *pb.ReleaseBookshelfRequest) error
	DeleteBook(req *pb.DeleteBookRequest) error
	DeleteBookshelf(req *pb.DeleteBookshelfRequest) error
}

type bookRequestValidation struct{}

func NewBookRequestValidation() BookRequestValidation {
	return &bookRequestValidation{}
}

func (v *bookRequestValidation) ListBookshelf(req *pb.ListBookshelfRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.ListBookshelfRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) ListBookReview(req *pb.ListBookReviewRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.ListBookReviewRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) ListUserReview(req *pb.ListUserReviewRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.ListUserReviewRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) ListUserMonthlyResult(req *pb.ListUserMonthlyResultRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.ListUserMonthlyResultRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) MultiGetBooks(req *pb.MultiGetBooksRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.MultiGetBooksRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) GetBook(req *pb.GetBookRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.GetBookRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) GetBookByIsbn(req *pb.GetBookByIsbnRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.GetBookByIsbnRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) GetBookshelf(req *pb.GetBookshelfRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.GetBookshelfRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) GetReview(req *pb.GetReviewRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.GetReviewRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) CreateBook(req *pb.CreateBookRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.CreateBookRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) UpdateBook(req *pb.UpdateBookRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.UpdateBookRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) ReadBookshelf(req *pb.ReadBookshelfRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.ReadBookshelfRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) ReadingBookshelf(req *pb.ReadingBookshelfRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.ReadingBookshelfRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) StackedBookshelf(req *pb.StackedBookshelfRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.StackedBookshelfRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) WantBookshelf(req *pb.WantBookshelfRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.WantBookshelfRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) ReleaseBookshelf(req *pb.ReleaseBookshelfRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.ReleaseBookshelfRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) DeleteBook(req *pb.DeleteBookRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.DeleteBookRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) DeleteBookshelf(req *pb.DeleteBookshelfRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(pb.DeleteBookshelfRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}
