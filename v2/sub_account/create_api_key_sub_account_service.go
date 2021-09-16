package sub_account

import (
	"context"
	"encoding/json"
)

type CreateAPIKeySubAccountService struct {
	c            *Client
	subAccountId string
	canTrade     bool
	marginTrade  bool
	futuresTrade bool
}

func (s *CreateAPIKeySubAccountService) SubAccountId(subAccountId string) *CreateAPIKeySubAccountService {
	s.subAccountId = subAccountId
	return s
}

func (s *CreateAPIKeySubAccountService) CanTrade(canTrade bool) *CreateAPIKeySubAccountService {
	s.canTrade = canTrade
	return s
}

func (s *CreateAPIKeySubAccountService) MarginTrade(marginTrade bool) *CreateAPIKeySubAccountService {
	s.marginTrade = marginTrade
	return s
}

func (s *CreateAPIKeySubAccountService) FuturesTrade(futuresTrade bool) *CreateAPIKeySubAccountService {
	s.futuresTrade = futuresTrade
	return s
}

func (s *CreateAPIKeySubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *APIKeySubAccount, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccount",
		secType:  secTypeSigned,
	}
	r.setParam("subAccountId", s.subAccountId)
	r.setParam("canTrade", s.canTrade)
	if s.marginTrade {
		r.setParam("marginTrade", s.marginTrade)
	}
	if s.futuresTrade {
		r.setParam("futuresTrade", s.futuresTrade)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &APIKeySubAccount{}, err
	}
	res = new(APIKeySubAccount)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &APIKeySubAccount{}, err
	}
	return res, nil
}

type APIKeySubAccount struct {
	SubAccountId string `json:"subaccountId"`
	ApiKey       string `json:"apiKey"`
	SecretKey    string `json:"secretKey"`
	CanTrade     bool   `json:"canTrade"`
	MarginTrade  bool   `json:"marginTrade"`
	FuturesTrade bool   `json:"futuresTrade"`
}
