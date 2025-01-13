package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	lab "github.com/goexpert/labobservabilidade"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func GetLogradouro(ctx context.Context, tracer trace.Tracer, cep lab.CepDto, client *http.Client) (*lab.LogradouroDto, error) {

	ctx, span := tracer.Start(ctx, "cepapi")
	defer span.End()

	// _url := "https://viacep.com.br/ws/" + cep.Cep + "/json/"
	_url := "https://opencep.com/v1/" + cep.Cep + ".json"

	wClient, err := lab.NewWebclient(ctx, client, http.MethodGet, _url, nil)
	if err != nil {
		slog.Error("falha na req para o OpenCep", "error", err.Error())
		return nil, err
	}
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(wClient.Request().Header))

	var logradouro lab.LogradouroDto

	err = wClient.Do(func(p []byte) error {
		err = json.Unmarshal(p, &logradouro)
		if err != nil {
			slog.Error("erro no unmarshal do cep", "error", err.Error())
		}
		return err
	})
	if err != nil {
		slog.Error("executa webclient", "error", err.Error())

	}

	if logradouro.Erro != "" {
		return nil, errors.New("cep não encontrado")
	}

	return &logradouro, err
}
