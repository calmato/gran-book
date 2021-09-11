package entity

import pb "github.com/calmato/gran-book/api/gateway/native/proto"

type Review struct {
	*pb.Review
}

type Reviews []*Review

func NewReview(r *pb.Review) *Review {
	return &Review{r}
}

func NewReviews(rs []*pb.Review) Reviews {
	res := make(Reviews, len(rs))
	for i := range rs {
		res[i] = NewReview(rs[i])
	}
	return res
}

func (rs Reviews) UserIDs() []string {
	userIDs := []string{}
	reviews := map[string]bool{}
	for _, r := range rs {
		if _, ok := reviews[r.UserId]; ok {
			continue
		}

		reviews[r.UserId] = true
		userIDs = append(userIDs, r.UserId)
	}
	return userIDs
}

func (rs Reviews) BookIDs() []int64 {
	bookIDs := []int64{}
	reviews := map[int64]bool{}
	for _, r := range rs {
		if _, ok := reviews[r.BookId]; ok {
			continue
		}

		reviews[r.BookId] = true
		bookIDs = append(bookIDs, r.BookId)
	}
	return bookIDs
}
