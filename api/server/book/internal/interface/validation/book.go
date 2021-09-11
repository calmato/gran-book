package validation

import (
	pb "github.com/calmato/gran-book/api/server/book/proto/book"
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
	if err != nil {
		if err, ok := err.(pb.ListBookshelfRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *bookRequestValidation) ListBookReview(req *pb.ListBookReviewRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.ListBookReviewRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *bookRequestValidation) ListUserReview(req *pb.ListUserReviewRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.ListUserReviewRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *bookRequestValidation) ListUserMonthlyResult(req *pb.ListUserMonthlyResultRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.ListUserMonthlyResultRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *bookRequestValidation) MultiGetBooks(req *pb.MultiGetBooksRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.MultiGetBooksRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *bookRequestValidation) GetBook(req *pb.GetBookRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.GetBookRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *bookRequestValidation) GetBookByIsbn(req *pb.GetBookByIsbnRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.GetBookByIsbnRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *bookRequestValidation) GetBookshelf(req *pb.GetBookshelfRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.GetBookshelfRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *bookRequestValidation) GetReview(req *pb.GetReviewRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.GetReviewRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *bookRequestValidation) CreateBook(req *pb.CreateBookRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.CreateBookRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *bookRequestValidation) UpdateBook(req *pb.UpdateBookRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.UpdateBookRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *bookRequestValidation) ReadBookshelf(req *pb.ReadBookshelfRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.ReadBookshelfRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *bookRequestValidation) ReadingBookshelf(req *pb.ReadingBookshelfRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.ReadingBookshelfRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *bookRequestValidation) StackedBookshelf(req *pb.StackedBookshelfRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.StackedBookshelfRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *bookRequestValidation) WantBookshelf(req *pb.WantBookshelfRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.WantBookshelfRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *bookRequestValidation) ReleaseBookshelf(req *pb.ReleaseBookshelfRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.ReleaseBookshelfRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *bookRequestValidation) DeleteBook(req *pb.DeleteBookRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.DeleteBookRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}

func (v *bookRequestValidation) DeleteBookshelf(req *pb.DeleteBookshelfRequest) error {
	err := req.Validate()
	if err != nil {
		if err, ok := err.(pb.DeleteBookshelfRequestValidationError); ok {
			return toValidationError(err.Field(), err.Reason())
		}

		return toInternalError()
	}

	return nil
}
