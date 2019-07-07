package model

const (
	BookableGymColumnNum    = 5
	BookableGymColumnName   = 0
	BookableGymColumnRoom   = 1
	BookableGymColumnDate   = 2
	BookableGymColumnTime   = 3
	BookableGymColumnButton = 4

	BookableStatusBookable    = 0
	BookableStatusNotReleased = 1
)

type BookableGymList []BookableGymModel

type BookableGymModel struct {
	Name   string `json:"name"`
	Room   string `json:"room"`
	Date   string `json:"date"`
	Time   string `json:"time"`
	Status int    `json:"status"`
}
