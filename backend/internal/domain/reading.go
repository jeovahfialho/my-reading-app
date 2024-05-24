package domain

// Reading represents the reading plan for a specific day
type Reading struct {
	Day           int    `bson:"day"`
	Period        string `bson:"period"`
	FirstReading  string `bson:"first_reading"`
	SecondReading string `bson:"second_reading"`
	ThirdReading  string `bson:"third_reading"`
}
