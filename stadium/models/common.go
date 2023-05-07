package models

type DefaultError struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Query struct {
	Offset int    `json:"offset" default:"0"`
	Limit  int    `json:"limit" default:"10"`
	Search string `json:"search"`
}

/*
{
    "StadiumName": "MetLife Stadium",
    "AvailableTimeStart": "2022-2",
    "AvailableTimeEnd": "2022-21",
    "PricePerHour": 100,
    "Rating": 4.5,
    "Images": "20140726.jpg",
    "StadiumType": "inside",
    "Location": "jsonb"{
      "longitude": -74.005941,
      "latitude": 40.712783
    }
*/