package domain

type Flat struct {
	ID      int64  `json:"id"`
	HouseID int64  `json:"house_id"`
	Price   int64  `json:"price"`
	Rooms   int8   `json:"rooms"`
	Status  string `json:"status"`
}
