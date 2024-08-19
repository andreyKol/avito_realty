package realty

type CreateHouseRequest struct {
	Address   string  `json:"address"`
	Year      int16   `json:"year"`
	Developer *string `json:"developer"`
}

type CreateFlatRequest struct {
	HouseID int64 `json:"house_id"`
	Price   int64 `json:"price"`
	Rooms   int8  `json:"rooms"`
}

type UpdateFlatStatusRequest struct {
	FlatID    int64  `json:"flat_id"`
	NewStatus string `json:"new_status"`
}

type SubscribeToHouseRequest struct {
	Email string `json:"email"`
}
