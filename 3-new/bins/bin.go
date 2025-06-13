package bins

import (
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"dateAndTime"` //дата и время
	Name      string    `json:"name"`
}

func CreateBin(id, name string, private bool) Bin {
	return Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}
