package integration_test

import (
	"time"

	"github.com/stretchr/testify/require"

	"github.com/bianjieai/irita-sdk-go/modules/validator"
	sdk "github.com/bianjieai/irita-sdk-go/types"
)

func (s IntegrationTestSuite) TestValidator() {
	baseTx := sdk.BaseTx{
		From:     s.Account().Name,
		Gas:      200000,
		Memo:     "test",
		Mode:     sdk.Commit,
		Password: s.Account().Password,
	}

	cert := "-----BEGIN CERTIFICATE-----\nMIICkDCCAjegAwIBAgIIICAGFgACiJQwCgYIKoEcz1UBg3UwdjEcMBoGA1UEAwwT\nU21hcnRDQV9UZXN0X1NNMl9DQTEVMBMGA1UECwwMU21hcnRDQV9UZXN0MRAwDgYD\nVQQKDAdTbWFydENBMQ8wDQYDVQQHDAbljZfkuqwxDzANBgNVBAgMBuaxn+iLjzEL\nMAkGA1UEBhMCQ04wHhcNMjAwNjE2MDIyNjQ1WhcNMjEwNjE2MDIyNjQ3WjCBpTFR\nME8GA1UELQxIYXBpX2NhX1RFU1RfVE9fUEhfUkFfMl9VU0wxM1NTMV9hcGlfY2Ff\nMGRmMDkwNTIyZWMzNGMyN2I4OWE3ZGM2NGFhZjBjNWUxMQswCQYDVQQGEwJDTjEL\nMAkGA1UECAwCU0gxCzAJBgNVBAcMAlNIMQswCQYDVQQKDAJCSjELMAkGA1UECwwC\nVEoxDzANBgNVBAMMBnp6cUAzODBZMBMGByqGSM49AgEGCCqBHM9VAYItA0IABLXE\n/d2lBDUiukYTKeyMePQTrxpqoqInpVR4sxZAMRcJ5hPiFQy86cDsuhYRyTP7e5VG\nzpVpKCHZtU/5MVD9ZkmjfzB9MAsGA1UdDwQEAwIGwDAMBgNVHRMEBTADAQH/MB0G\nA1UdDgQWBBTdlZ0yFCT4VVCLglPPfTKFGZxd9jAfBgNVHSMEGDAWgBRc87oljMJl\nDOxn777djWunXq/jrDAgBggqgRzQFAQBAwQUExIzMjAxMjUyMDE4MDkwOTM0MTgw\nCgYIKoEcz1UBg3UDRwAwRAIgCg8nQAK64wWok6v/vwu9VF21UcRyP2X6rsNBZsKd\ny7sCICyeC3c+QjMB6dr/DyJieA6s48tLob4/2z7uMmh7eZZK\n-----END CERTIFICATE-----"

	createReq := validator.CreateValidatorRequest{
		Name:        "test1",
		Certificate: cert,
		Power:       10,
		Details:     "this is a test",
	}
	rs, err := s.Validator.CreateValidator(createReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), rs.Hash)

	validatorID, er := rs.Events.GetValue("create_validator", "validator")
	require.NoError(s.T(), er)

	v, err := s.Validator.QueryValidator(validatorID)
	require.NoError(s.T(), err)

	vs, err := s.Validator.QueryValidators(1, 10)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), vs)

	updateReq := validator.UpdateValidatorRequest{
		ID:          validatorID,
		Name:        "test2",
		Certificate: cert,
		Power:       10,
		Details:     "this is a updated test",
	}
	rs, err = s.Validator.UpdateValidator(updateReq, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), rs.Hash)

	v, err = s.Validator.QueryValidator(validatorID)
	require.NoError(s.T(), err)
	require.Equal(s.T(), updateReq.Name, v.Name)
	require.Equal(s.T(), updateReq.Details, v.Details)

	//wait to be jail
	time.Sleep(10 * time.Second)

	rs, err = s.Validator.Unjail(validatorID, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), rs.Hash)

	rs, err = s.Validator.RemoveValidator(validatorID, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), rs.Hash)

	v, err = s.Validator.QueryValidator(validatorID)
	require.Error(s.T(), err)
}
