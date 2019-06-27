package models

var (
	ReviewList map[string]*Review
)

func init() {
	ReviewList = make(map[string]*Review)
	u := Review{"user_11111", "bob", "11111"}
	ReviewList["user_11111"] = &u
}

type Review struct {
	Id       string
	Username string
	Password string
}

func GetAllReviews() map[string]*Review {
	return ReviewList
}

func GetProductReviews() map[string]*Review {
	return ReviewList
}
