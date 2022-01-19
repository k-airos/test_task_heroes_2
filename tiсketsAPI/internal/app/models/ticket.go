package models

//Ticket model defenition
type Ticket struct {
	ID           int    `json:"id"`
	Film         string `json:"film"`
	ScreenWriter string `json:"screen_writer"`
	Content      string `json:"content"`
}
