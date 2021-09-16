package sub_account

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type createSubAccountTestSuite struct {
	baseTestSuite
}

func TestAccountService(t *testing.T) {
	suite.Run(t, new(createSubAccountTestSuite))
}

func (s *createSubAccountTestSuite) TestCreateAccount() {
	data := []byte(`{
  		"subaccountId": "1",
  		"email": "vai_42038996_47411276_brokersubuser@lac.info",
  		"tag":"bob123d"
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewCreateSubAccountService().Do(newContext())
	s.r().NoError(err)
	//s.r().Len(res, 1)
	e := &SubAccount{
		SubAccountId:       "1",
		Email:              "vai_42038996_47411276_brokersubuser@lac.info",
		Tag:            	"bob123d",
	}
	s.assertAccountEqual(e, res)
}

func (s *createSubAccountTestSuite) assertAccountEqual(e, a *SubAccount) {
	r := s.r()
	r.Equal(e.SubAccountId, a.SubAccountId, "subaccountId")
	r.Equal(e.Email, a.Email, "email")
	r.Equal(e.Tag, a.Tag, "tag")
}
