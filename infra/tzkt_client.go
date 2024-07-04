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
	client *http.Client
	host   string
}

func NewTzktClient(
	client *http.Client,
	host string,

) app.ThezosSvc {
	return TzktClient{
		client: client,
		host:   host,
	}
}

func (x TzktClient) GetDelegations(ctx context.Context, id int, resultLimit int) ([]app.DelegationDto, error) {
	req, err := http.NewRequest(http.MethodGet, x.host+endpoint, nil)
	if err != nil {
		return nil, app.ErrTechnical
	}

	q := req.URL.Query()
	q.Add("limit", fmt.Sprintf("%d", resultLimit))
	q.Add("sort.asc", "id")

	if id != -1 {
		q.Add("id.gt", fmt.Sprintf("%d", id))
	}

	req.URL.RawQuery = q.Encode()

	res, err := x.client.Do(req)
	if err != nil {
		return nil, app.ErrTechnical
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, app.ErrTechnical
	}

	dtos := []DelegationDto{}
	err = json.Unmarshal(body, &dtos)
	if err != nil {
		return nil, app.ErrTechnical
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
