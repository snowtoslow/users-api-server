package model

type UserReq struct {
	ID        string  `json:"id"`
	FirstName string  `json:"firstname"`
	LastName  string  `json:"lastname"`
	Age       uint8   `json:"age"`
	RandomKey string  `json:"random_key"`
	LatLong   LatLong `json:"maps"`
}

type UserResponse struct {
	ID          string  `json:"id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Age         uint8   `json:"age"`
	LatLongResp LatLong `json:"lat_long_resp"`
}

type LatLong struct {
	ID        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PaginationResponse struct {
	PageInfo    PageInfo
	PreviewPage []UserResponse `json:"preview_page"`
}

type PageInfo struct {
	TotalItems  int64 `json:"total_items"`
	TotalPage   int64 `json:"total_page"`
	CurrentPage int   `json:"current_page"`
	NextPage    int64 `json:"next_page"`
}
