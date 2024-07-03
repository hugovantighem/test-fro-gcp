package infra_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"myproject/infra"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const host = "https://api.tzkt.io"

func TestGetDelegations(t *testing.T) {
	httpClient := &http.Client{
		Transport: GetDelegationsStub{},
	}
	elmtCount := 8
	client := infra.NewTzktClient(httpClient, host, elmtCount)

	result, err := client.GetDelegations(context.Background(), 0)
	require.NoError(t, err)
	assert.Len(t, result, elmtCount)

}

type GetDelegationsStub struct {
}

func (x GetDelegationsStub) RoundTrip(req *http.Request) (*http.Response, error) {
	limitParam := req.URL.Query().Get("limit")

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return nil, err
	}

	itemFactory := DeletationItemFactory()

	listing := make([]map[string]any, limit)
	for i := 0; i < limit; i++ {
		listing[i] = itemFactory()
	}

	b, err := json.Marshal(listing)
	if err != nil {
		return nil, err
	}

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(b)),
	}, nil
}

func DeletationItemFactory() func() map[string]any {
	id := 0

	return func() map[string]any {
		id++
		return map[string]any{
			"type":      "delegation",
			"id":        id,
			"level":     109,
			"timestamp": "2018-06-30T19:30:27Z",
			"block":     "BLwRUPupdhP8TyWp9J6TbjLSCxPPW6tyhVPF2KmNAbLPt7thjPw",
			"hash":      "ooP37LNma6DiWjVxDbS2XZu4PiNKy7fbHZWSn8Vj8FX1hWfkC3b",
			"counter":   23,
			"sender": map[string]any{
				"address": "tz1Wit2PqodvPeuRRhdQXmkrtU8e8bRYZecd",
			},
			"gasLimit":     0,
			"gasUsed":      0,
			"storageLimit": 0,
			"bakerFee":     50000,
			"amount":       25079312620,
			"newDelegate": map[string]any{
				"address": "tz1Wit2PqodvPeuRRhdQXmkrtU8e8bRYZecd",
			},
			"status": "applied",
		}
	}

}
