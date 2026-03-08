package weather

import "fmt"

type Service struct {
	client *Client
}

func NewService(client *Client) *Service {
	return &Service{client: client}
}

func (s *Service) GetCurrentWeather(city string) (*Result, error) {
	data, err := s.client.FetchCurrentWeather(city)
	if err != nil {
		return nil, err
	}

	if len(data.Weather) == 0 {
		return nil, fmt.Errorf("no weather information found for %s", city)
	}

	return &Result{
		City:        data.Name,
		Temperature: data.Main.Temp,
		Condition:   data.Weather[0].Main,
		Description: data.Weather[0].Description,
		Humidity:    data.Main.Humidity,
		WindSpeed:   data.Wind.Speed,
	}, nil
}
