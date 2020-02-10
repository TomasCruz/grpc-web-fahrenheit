package presenter

// ErrStruct contains errorMessage
type ErrStruct struct {
	Msg string `json:"errorMessage" example:"A horrible, terrible, absolutely awful error"`
}

// Health contains status
type Health struct {
	Status string `json:"status" example:"UP"`
}

// DegreePair contains celsius/fahrenheit pair
type DegreePair struct {
	C float64 `json:"celsius" example:"100"`
	F float64 `json:"fahrenheit" example:"212"`
}
