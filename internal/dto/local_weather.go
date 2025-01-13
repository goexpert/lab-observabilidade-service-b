package dto

type LocalWeatherDto struct {
	Locale string  `json:"city"`
	TempC  float64 `json:"temp_c"`
	TempF  float64 `json:"temp_f"`
	TempK  float64 `json:"temp_k"`
}
