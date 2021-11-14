package model

type LinkTrace struct {
	Sc     string `bson:"sc"`
	URL    string `bson:"url"`
	Ip     string `bson:"ip"`
	UA     string `bson:"ua"`
	Cookie string `bson:"cookie"`
}
