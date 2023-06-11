package model

// Member - struct
// By adding struct tags you can control exactly what an how your struct will be marshalled to JSON
type Member struct {
	ProfileId        string `json:"profile_id"`
	Fullname         string `json:"fullname"`
	ProfileLink      string `json:"profile_link"`
	Bio              string `json:"bio"`
	ImageSrc         string `json:"image_src"`
	GroupId          string `json:"group_id"`
	GroupJoiningText string `json:"group_joining_text"`
	ProfileType      string `json:"profile_type"`
}

// Members array
type Members []Member
