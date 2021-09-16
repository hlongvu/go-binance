package sub_account

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type createSubAccountAPIKeyTestSuite struct {
	baseTestSuite
}

func TestAccountAPIKeyService(t *testing.T) {
	suite.Run(t, new(createSubAccountAPIKeyTestSuite))
}

func (s *createSubAccountAPIKeyTestSuite) TestCreateAccount() {
	data := []byte(`{
    "subaccountId": "1",
    "apiKey":"vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A",
    "secretKey":"NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0",
    "canTrade": true,
    "marginTrade": false,
    "futuresTrade": false
}`)

	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"subAccountId": "1",
			"canTrade":     true,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCreateAPIKeySubAccountService().
		SubAccountId("1").
		CanTrade(true).
		Do(newContext())

	s.r().NoError(err)
	//s.r().Len(res, 1)
	e := &APIKeySubAccount{
		SubAccountId: "1",
		ApiKey:       "vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A",
		SecretKey:    "NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0",
		CanTrade:     true,
		MarginTrade:  false,
		FuturesTrade: false,
	}
	s.assertAccountEqual(e, res)
}

func (s *createSubAccountAPIKeyTestSuite) assertAccountEqual(e, a *APIKeySubAccount) {
	r := s.r()
	r.Equal(e.SubAccountId, a.SubAccountId, "subaccountId")
	r.Equal(e.ApiKey, a.ApiKey, "apiKey")
	r.Equal(e.SecretKey, a.SecretKey, "secretKey")
	r.Equal(e.CanTrade, a.CanTrade, "canTrade")
	r.Equal(e.MarginTrade, a.MarginTrade, "marginTrade")
	r.Equal(e.FuturesTrade, a.FuturesTrade, "futuresTrade")

}
