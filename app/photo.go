package app

type PhotoAdded struct {
	Title string `json:"title" valid:"required"`
	Caption string `json:"caption" valid:"required"`
	PhotoUrl string `json:"photo_url" valid:"required"`
}

type PhotoUpdate struct {
	Title string `json:"title" valid:"required"`
	Caption string `json:"caption" valid:"required"`
	PhotoUrl string `json:"photo_url" valid:"required"`
}