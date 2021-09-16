package sub_account

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type enableSubAccountMarginTestSuite struct {
	baseTestSuite
}

func TestAccountServiceMargin(t *testing.T) {
	suite.Run(t, new(enableSubAccountMarginTestSuite))
}

func (s *enableSubAccountMarginTestSuite) TestAccountServiceMargin() {
	data := []byte(`{
    "subaccountId": "1",
    "enableMargin": true,
    "updateTime": 1570801523523
}`)

	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"subAccountId":  "1",
			"margin": true,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewEnableSubAccountMarginService().
		SubAccountId("1").
		Do(newContext())
	s.r().NoError(err)
	//s.r().Len(res, 1)
	e := &SubAccountMargin{
		SubAccountId:       "1",
		EnableMargin:      true,
		UpdateTime:         1570801523523,
	}
	s.assertAccountEqual(e, res)
}

func (s *enableSubAccountMarginTestSuite) assertAccountEqual(e, a *SubAccountMargin) {
	r := s.r()
	r.Equal(e.SubAccountId, a.SubAccountId, "subaccountId")
	r.Equal(e.EnableMargin, a.EnableMargin, "enableMargin")
	r.Equal(e.UpdateTime, a.UpdateTime, "updateTime")
}
