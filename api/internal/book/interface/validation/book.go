package validation

import "github.com/calmato/gran-book/api/proto/book"

type bookRequestValidation struct{}

func NewBookRequestValidation() BookRequestValidation {
	return &bookRequestValidation{}
}

func (v *bookRequestValidation) ListBookshelf(req *book.ListBookshelfRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(book.ListBookshelfRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) ListBookReview(req *book.ListBookReviewRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(book.ListBookReviewRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) ListUserReview(req *book.ListUserReviewRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(book.ListUserReviewRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) ListUserMonthlyResult(req *book.ListUserMonthlyResultRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(book.ListUserMonthlyResultRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) MultiGetBooks(req *book.MultiGetBooksRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(book.MultiGetBooksRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) GetBook(req *book.GetBookRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(book.GetBookRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) GetBookByIsbn(req *book.GetBookByIsbnRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(book.GetBookByIsbnRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) GetBookshelf(req *book.GetBookshelfRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(book.GetBookshelfRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) GetReview(req *book.GetReviewRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(book.GetReviewRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) CreateBook(req *book.CreateBookRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(book.CreateBookRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) UpdateBook(req *book.UpdateBookRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(book.UpdateBookRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) ReadBookshelf(req *book.ReadBookshelfRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(book.ReadBookshelfRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) ReadingBookshelf(req *book.ReadingBookshelfRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(book.ReadingBookshelfRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) StackedBookshelf(req *book.StackedBookshelfRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(book.StackedBookshelfRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) WantBookshelf(req *book.WantBookshelfRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(book.WantBookshelfRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) ReleaseBookshelf(req *book.ReleaseBookshelfRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(book.ReleaseBookshelfRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) DeleteBook(req *book.DeleteBookRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(book.DeleteBookRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}

func (v *bookRequestValidation) DeleteBookshelf(req *book.DeleteBookshelfRequest) error {
	err := req.Validate()
	if err == nil {
		return nil
	}

	validate := err.(book.DeleteBookshelfRequestValidationError)
	return toValidationError(validate.Field(), validate.Reason())
}
