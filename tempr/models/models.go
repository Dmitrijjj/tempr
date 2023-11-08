package models

type Responce struct {
	Main  Main   `json:"main"`
	CName string `json:"name"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	Humidity int32   `json:"humidity"`
}

type Memury struct {
	Temp     float64
	Humidity int32
}
