package api

type Vegetable struct {
	Id              int32  `json:"id"`
	Name            string `json:"string"`
	Amount_per_100  string `json:"status"`
	Depth           string `json:"depth" binding:"required"`
	Distance_row    string `json:"distance_row"`
	Distance_plant  string `json:"distance_plant"`
	Height          string `json:"height"`
	Spring_planting string `json:"spring_planting"`
	Fall_planting   string `json:"fall_planting"`
	Created_at      int32  `json:"created_at"`
	Updated_at      int32  `json:"updated_at"`
}
