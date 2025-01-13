package entity

import (
	"errors"
	"log/slog"
	"math"
	"strings"

	lab "github.com/goexpert/labobservabilidade"
)

type localeWeatherEntity struct {
	locale string
	tempC  float64
	tempF  float64
	tempK  float64
}

func (w *localeWeatherEntity) TempC() float64 {
	return w.tempC
}

func (w *localeWeatherEntity) Locale() string {
	return w.locale
}

func NewLocaleWeather(locale string, tempC float64) (*lab.LocaleWeatherDto, error) {

	locale = strings.TrimSpace(locale)

	var tc = &localeWeatherEntity{
		locale: locale,
		tempC:  tempC,
		tempF:  0,
		tempK:  0,
	}
	slog.Debug("struct", "localeWeatherEntity", tc)

	err := tc.IsValid()
	if err != nil {
		slog.Error("[invalid locale]", "error", err.Error())
		return nil, err
	}

	return &lab.LocaleWeatherDto{
		Locale: tc.locale,
		TempC:  math.Round((tc.tempC)*10) / 10,
		TempF:  math.Round((tc.tempC*1.8+32)*10) / 10,
		TempK:  math.Round((tc.tempC+273)*10) / 10,
	}, nil
}

func (lw *localeWeatherEntity) IsValid() error {

	if len(lw.Locale()) < 1 {
		return errors.New("local não pode ser vazio")
	}
	return nil
}
