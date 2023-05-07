package models

import "time"

type Stadium struct {
	ID                 int     `json:"id"`
	StadiumName        string  `json:"stadium_name"`
	AvailableTimeStart string  `json:"available_time_start"`
	AvailableTimeEnd   string  `json:"available_time_end"`
	PricePerHour       int     `json:"price_p_hour"`
	Rating             float64 `json:"rating"`
	Images             string  `json:"images"`
	Location           string
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type CreateStadium struct {
	StadiumName        string  `json:"stadium_name"`
	AvailableTimeStart string  `json:"available_time_start"`
	AvailableTimeEnd   string  `json:"available_time_end"`
	PricePerHour       int     `json:"price_p_hour"`
	Rating             float64 `json:"rating"`
	Images             string  `json:"images"`
	Location           *LocationJ
}

type LocationJ struct {
	Long float64 `json:"long"`
	Lat  float64 `json:"lat"`
}

type UpdateStadium struct {
	StadiumName        string  `json:"stadium_name"`
	AvailableTimeStart string  `json:"available_time_start"`
	AvailableTimeEnd   string  `json:"available_time_end"`
	PricePerHour       int     `json:"price_p_hour"`
	Rating             float64 `json:"rating"`
	Images             string  `json:"images"`
	//StadiumType        string  `json:"stadium_type"`
	Location struct {
		Long float64 `json:"long"`
		Lat  float64 `json:"lat"`
	}
}
