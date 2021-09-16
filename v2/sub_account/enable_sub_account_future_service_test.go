package sub_account

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type enableSubAccountFutureTestSuite struct {
	baseTestSuite
}

func TestAccountServiceFuture(t *testing.T) {
	suite.Run(t, new(enableSubAccountFutureTestSuite))
}

func (s *enableSubAccountFutureTestSuite) TestSubAccountFuture() {
	data := []byte(`{
    	"subaccountId": "1",
    	"enableFutures": true,
    	"updateTime": 1570801523523
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"subAccountId":  "1",
			"futures": true,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewEnableSubAccountFutureService().
		SubAccountId("1").
		Do(newContext())
	s.r().NoError(err)
	//s.r().Len(res, 1)
	e := &SubAccountFuture{
		SubAccountId:       "1",
		EnableFutures:      true,
		UpdateTime:         1570801523523,
	}
	s.assertAccountEqual(e, res)
}

func (s *enableSubAccountFutureTestSuite) assertAccountEqual(e, a *SubAccountFuture) {
	r := s.r()
	r.Equal(e.SubAccountId, a.SubAccountId, "subaccountId")
	r.Equal(e.EnableFutures, a.EnableFutures, "enableFutures")
	r.Equal(e.UpdateTime, a.UpdateTime, "updateTime")
}
