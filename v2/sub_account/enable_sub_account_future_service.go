package sub_account
import (
"context"
"encoding/json"
)

type EnableSubAccountFutureService struct {
	c 	*Client
	subAccountId string
}

func (s *EnableSubAccountFutureService) SubAccountId(subAccountId string) *EnableSubAccountFutureService {
	s.subAccountId = subAccountId
	return s
}

func (s *EnableSubAccountFutureService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountFuture, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccount/futures",
		secType:  secTypeSigned,
	}
	r.setParam("subAccountId", s.subAccountId)
	r.setParam("futures", true)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &SubAccountFuture{}, err
	}
	res = new(SubAccountFuture)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &SubAccountFuture{}, err
	}
	return res, nil
}

type SubAccountFuture struct {
	SubAccountId       string `json:"subaccountId"`
	EnableFutures      bool `json:"enableFutures"`
	UpdateTime         int64   `json:"updateTime"`
}