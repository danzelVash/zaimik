package models

type Review struct {
	Id           int    `json:"-" db:"id"`
	ReviewerName string `json:"reviewer_name" binding:"required" db:"reviewer_name"`
	ReviwerPhone string `json:"reviwer_phone" binding:"required" db:"reviewer_phone"`
	Review       string `json:"review" binding:"required" db:"review"`
}

type ReviewAdmin struct {
	Id            int    `json:"id" db:"id"`
	ReviewerName  string `json:"reviewer_name" binding:"required" db:"reviewer_name"`
	ReviewerPhone string `json:"reviewer_phone" binding:"required" db:"reviewer_phone"`
	Review        string `json:"review" binding:"required" db:"review"`
	Moderated     bool   `json:"moderated" binding:"required" db:"moderated"`
}

func (r *ReviewAdmin) Valid() bool {
	if r.ReviewerName == "" || r.ReviewerPhone == "" || r.Review == "" {
		return false
	}
	return true
}
