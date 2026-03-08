package weather

type apiResponse struct {
	Name    string        `json:"name"`
	Main    mainInfo      `json:"main"`
	Weather []weatherInfo `json:"weather"`
	Wind    windInfo      `json:"wind"`
}

type mainInfo struct {
	Temp     float64 `json:"temp"`
	Humidity int     `json:"humidity"`
}

type weatherInfo struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type windInfo struct {
	Speed float64 `json:"speed"`
}

type Result struct {
	City        string
	Temperature float64
	Condition   string
	Description string
	Humidity    int
	WindSpeed   float64
}
