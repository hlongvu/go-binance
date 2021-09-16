package sub_account

import (
	"context"
	"encoding/json"
)

type CreateSubAccountService struct {
	c 	*Client
	tag string
}

func (s *CreateSubAccountService) Tag(tag string) *CreateSubAccountService {
	s.tag = tag
	return s
}

func (s *CreateSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccount, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccount",
		secType:  secTypeSigned,
	}
	if s.tag != "" {
		r.setParam("tag", s.tag)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &SubAccount{}, err
	}
	res = new(SubAccount)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &SubAccount{}, err
	}
	return res, nil
}

type SubAccount struct {
	SubAccountId       string `json:"subaccountId"`
	Email              string `json:"email"`
	Tag            	   string `json:"tag"`
}