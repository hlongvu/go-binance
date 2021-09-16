package sub_account

import (
	"context"
	"encoding/json"
)

type EnableSubAccountMarginService struct {
	c            *Client
	subAccountId string
}

func (s *EnableSubAccountMarginService) SubAccountId(subAccountId string) *EnableSubAccountMarginService {
	s.subAccountId = subAccountId
	return s
}

func (s *EnableSubAccountMarginService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountMargin, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccount/futures",
		secType:  secTypeSigned,
	}
	r.setParam("subAccountId", s.subAccountId)
	r.setParam("margin", true)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &SubAccountMargin{}, err
	}
	res = new(SubAccountMargin)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &SubAccountMargin{}, err
	}
	return res, nil
}

type SubAccountMargin struct {
	SubAccountId string `json:"subaccountId"`
	EnableMargin bool   `json:"enableMargin"`
	UpdateTime   int64  `json:"updateTime"`
}
