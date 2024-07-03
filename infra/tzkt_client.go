package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"myproject/app"
	"net/http"
	"time"
)

const endpoint = "/v1/operations/delegations"

type TzktClient struct {
	client      *http.Client
	host        string
	resultLimit int
}

func NewTzktClient(
	client *http.Client,
	host string,
	resultLimit int,
) app.ThezosSvc {
	return TzktClient{
		client:      client,
		host:        host,
		resultLimit: resultLimit,
	}
}

func (x TzktClient) GetDelegations(ctx context.Context, id int) ([]app.DelegationDto, error) {
	req, err := http.NewRequest(http.MethodGet, x.host+endpoint, nil)
	if err != nil {
		return nil, app.NewTechnicalError(err, "NewRequest failed")
	}

	q := req.URL.Query()
	q.Add("limit", fmt.Sprintf("%d", x.resultLimit))
	q.Add("sort.asc", "id")

	req.URL.RawQuery = q.Encode()

	res, err := x.client.Do(req)
	if err != nil {
		return nil, app.NewTechnicalError(err, "Do failed")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, app.NewTechnicalError(err, "ReadAll failed")
	}

	dtos := []DelegationDto{}
	err = json.Unmarshal(body, &dtos)
	if err != nil {
		return nil, app.NewTechnicalError(err, "Unmarshal failed")
	}

	result := make([]app.DelegationDto, len(dtos))
	for idx, v := range dtos {
		result[idx] = app.DelegationDto{
			Id:          v.Id,
			Amount:      v.Amount,
			SenderAddr:  v.Sender.Address,
			BlockHeight: v.Level,
			Timestamp:   v.Timestamp,
		}
	}

	return result, nil
}

type DelegationDto struct {
	Id        int
	Level     int
	Amount    int
	Timestamp time.Time
	Sender    SenderDto
}

type SenderDto struct {
	Address string
}
