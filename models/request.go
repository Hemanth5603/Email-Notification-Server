package models

type EmailRequest struct {
	To          string `json:"to"`
	Weather     string `json:"weather"`
	Temperature string `json:"temperature"`
	Location    string `json:"location"`
	Date        string `json:"date"`
	Time        string `json:"time"`
}
