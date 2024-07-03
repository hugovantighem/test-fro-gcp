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

const data = `
[
	{
		"type": "delegation",
		"id": 1098907648,
		"level": 109,
		"timestamp": "2018-06-30T19:30:27Z",
		"block": "BLwRUPupdhP8TyWp9J6TbjLSCxPPW6tyhVPF2KmNAbLPt7thjPw",
		"hash": "ooP37LNma6DiWjVxDbS2XZu4PiNKy7fbHZWSn8Vj8FX1hWfkC3b",
		"counter": 23,
		"sender": {
			"address": "tz1Wit2PqodvPeuRRhdQXmkrtU8e8bRYZecd"
		},
		"gasLimit": 0,
		"gasUsed": 0,
		"storageLimit": 0,
		"bakerFee": 50000,
		"amount": 25079312620,
		"newDelegate": {
			"address": "tz1Wit2PqodvPeuRRhdQXmkrtU8e8bRYZecd"
		},
		"status": "applied"
	},
	{
		"type": "delegation",
		"id": 1649410048,
		"level": 167,
		"timestamp": "2018-06-30T20:29:42Z",
		"block": "BLzCkTwQGUf9ggfk24bW7YFeFzudfncp5zzJSkR6Lf4kSP923PK",
		"hash": "ooedoJWn6fFXaiCkNDftRVCvEbJ855M7fD7gzHryL6x6FXdejP4",
		"counter": 34,
		"sender": {
			"address": "tz1U2ufqFdVkN2RdYormwHtgm3ityYY1uqft"
		},
		"gasLimit": 0,
		"gasUsed": 0,
		"storageLimit": 0,
		"bakerFee": 100,
		"amount": 10199999690,
		"newDelegate": {
			"address": "tz1U2ufqFdVkN2RdYormwHtgm3ityYY1uqft"
		},
		"status": "applied"
	},
	{
		"type": "delegation",
		"id": 1751121920,
		"level": 177,
		"timestamp": "2018-06-30T20:39:42Z",
		"block": "BLNPz3sg78J3Dj2mMThBFPhkNHaKpszQGKVopQ7nST8hRJSU5oP",
		"hash": "onx3SMorQFtaMwCGV9bZuP3VtzyS54jEzTAJBWssMyS7JrWRnX7",
		"counter": 42,
		"sender": {
			"address": "tz1id8EA4PQf3NXeQ2LagJBK5gR8UsCg2fqW"
		},
		"gasLimit": 0,
		"gasUsed": 0,
		"storageLimit": 0,
		"bakerFee": 0,
		"amount": 10000,
		"newDelegate": {
			"address": "tz1id8EA4PQf3NXeQ2LagJBK5gR8UsCg2fqW"
		},
		"status": "applied"
	},
	{
		"type": "delegation",
		"id": 1879048192,
		"level": 190,
		"timestamp": "2018-06-30T20:52:42Z",
		"block": "BLoyJZXRjwXHuAnCjaqZyaJbSs6L8z8Dw7TYHp2hgqEqGDyxwGR",
		"hash": "opFkg3XMnot3eAVordhKHrZqYLfZgBBzzUjjjV8SuCu3rRjXapd",
		"counter": 37,
		"sender": {
			"address": "tz1bakerHqr7tLADpCajhJBCNEBEE6AS73Mz"
		},
		"gasLimit": 0,
		"gasUsed": 0,
		"storageLimit": 0,
		"bakerFee": 0,
		"amount": 1009990000,
		"newDelegate": {
			"address": "tz1bakerHqr7tLADpCajhJBCNEBEE6AS73Mz"
		},
		"status": "applied"
	},
	{
		"type": "delegation",
		"id": 2253389824,
		"level": 228,
		"timestamp": "2018-06-30T21:30:42Z",
		"block": "BL9Qmtx5oq74BHxHtu6EEVAdeJ5rmEFa6wiZWPhwKZn9RZ3oBVo",
		"hash": "oohKb1m1E2o4qjLCewb2Spm5SDPexkvWnsuQ1mfKyDbGjCUGd7r",
		"counter": 67,
		"sender": {
			"alias": "Tezos Istanbul",
			"address": "tz1Vm5cfHncKGBo7YvZfHc4mmudY4qpWzvSB"
		},
		"gasLimit": 0,
		"gasUsed": 0,
		"storageLimit": 0,
		"bakerFee": 50000,
		"amount": 299999950000,
		"newDelegate": {
			"alias": "Tezos Istanbul",
			"address": "tz1Vm5cfHncKGBo7YvZfHc4mmudY4qpWzvSB"
		},
		"status": "applied"
	},
	{
		"type": "delegation",
		"id": 2393899008,
		"level": 242,
		"timestamp": "2018-06-30T21:44:42Z",
		"block": "BMCnStqEhiEddMqMdM3REvuAxjUFN8hJ41XkYWcQTH8wBy4DW3F",
		"hash": "oor46nCMy9WYJj8Q59zpVMjYbuxJLUvLFw9BP5AWJjhAsFBXNw1",
		"counter": 2,
		"sender": {
			"address": "KT1XRAGQEbUqqagjX9j7RbvpUe7ZwF29DnM2"
		},
		"gasLimit": 0,
		"gasUsed": 0,
		"storageLimit": 0,
		"bakerFee": 0,
		"amount": 1193551000000,
		"newDelegate": {
			"alias": "Tezos Istanbul",
			"address": "tz1Vm5cfHncKGBo7YvZfHc4mmudY4qpWzvSB"
		},
		"status": "applied"
	},
	{
		"type": "delegation",
		"id": 2395996160,
		"level": 242,
		"timestamp": "2018-06-30T21:44:42Z",
		"block": "BMCnStqEhiEddMqMdM3REvuAxjUFN8hJ41XkYWcQTH8wBy4DW3F",
		"hash": "oopfZVZKp4Tjibvv9st3vYMMKiSTfUqPCYepeztX7QHSixWPSNc",
		"counter": 73,
		"sender": {
			"address": "tz1TZLETj4eRXzSe8aPEfEqov87bQq9rbTTX"
		},
		"gasLimit": 0,
		"gasUsed": 0,
		"storageLimit": 0,
		"bakerFee": 50000,
		"amount": 2524200000,
		"newDelegate": {
			"address": "tz1TZLETj4eRXzSe8aPEfEqov87bQq9rbTTX"
		},
		"status": "applied"
	},
	{
		"type": "delegation",
		"id": 3191865344,
		"level": 326,
		"timestamp": "2018-06-30T23:08:42Z",
		"block": "BLsr2wYFvfDj5zY85DWnBVCzewnn29GDKLMzhLy3oqE5yhkikeu",
		"hash": "oofDXKD5m8mBUtv6z3mGuqfPKy4weqC75gewEY3sKFjzUsbiDbF",
		"counter": 82,
		"sender": {
			"address": "tz1R4D8nfepxHTDppGzsQqD5CMM4jjmk5o91"
		},
		"gasLimit": 0,
		"gasUsed": 0,
		"storageLimit": 0,
		"bakerFee": 50000,
		"amount": 18374950000,
		"newDelegate": {
			"address": "tz1R4D8nfepxHTDppGzsQqD5CMM4jjmk5o91"
		},
		"status": "applied"
	},
	{
		"type": "delegation",
		"id": 3347054592,
		"level": 341,
		"timestamp": "2018-06-30T23:23:42Z",
		"block": "BMcMaGU4LZKpbNRvR1TFMHr8EubaMyfo3jHfNTSCSeJUycsZVAr",
		"hash": "op5WFJJCJ9eJcpWkBbs5KUoa1Su9atPfQiHECXk5o4KSkLG27di",
		"counter": 106,
		"sender": {
			"address": "tz1TgsrW4uPSiTef9mHP44ta8iMajhYTYPSh"
		},
		"gasLimit": 0,
		"gasUsed": 0,
		"storageLimit": 0,
		"bakerFee": 50000,
		"amount": 499950000,
		"newDelegate": {
			"address": "tz1TgsrW4uPSiTef9mHP44ta8iMajhYTYPSh"
		},
		"status": "applied"
	},
	{
		"type": "delegation",
		"id": 3526361088,
		"level": 359,
		"timestamp": "2018-06-30T23:41:42Z",
		"block": "BLKRzsXYE5pNXWk2HG2TCAX6nEti3QufZMLnezTHMuHPydo5CAb",
		"hash": "op2yUybRTzdrD8vq7zJB1cJuYgmiyUSaKLbfQAJA2jgADVJRhzw",
		"counter": 116,
		"sender": {
			"address": "tz1eewjYhPoR4o1x9pWgNtCU1GUAkp9Qb76q"
		},
		"gasLimit": 0,
		"gasUsed": 0,
		"storageLimit": 0,
		"bakerFee": 0,
		"amount": 100000,
		"newDelegate": {
			"address": "tz1eewjYhPoR4o1x9pWgNtCU1GUAkp9Qb76q"
		},
		"status": "applied"
	}
]`
