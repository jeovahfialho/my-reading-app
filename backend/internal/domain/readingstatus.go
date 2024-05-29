package domain

type ReadingStatus struct {
	UserID string `json:"userId" bson:"userId"`
	Day    int    `json:"day" bson:"day"`
	Status string `json:"status" bson:"status"`
}
