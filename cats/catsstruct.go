package cats

type Cat struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Breed   string `json:"breed"`
	Age     int    `json:"age"`
	OwnerId int    `json:"owner_id"`
	About   string `json:"about"`
}

type CatCreateRequest struct { // DTO
	Name    string `json:"name"`
	Breed   string `json:"breed"`
	Age     int    `json:"age"`
	OwnerId int    `json:"owner_id"`
	About   string `json:"about"`
}
