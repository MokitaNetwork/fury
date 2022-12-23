package types_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	"github.com/crescent-network/crescent/v4/x/claim/types"
)

type keysTestSuite struct {
	suite.Suite
}

func TestKeysTestSuite(t *testing.T) {
	suite.Run(t, new(keysTestSuite))
}

func (suite *keysTestSuite) TestGetAirdropKey() {
	suite.Require().Equal([]byte{0xd5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetAirdropKey(0))
	suite.Require().Equal([]byte{0xd5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5}, types.GetAirdropKey(5))
	suite.Require().Equal([]byte{0xd5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}, types.GetAirdropKey(10))
}

func (s *keysTestSuite) TestGetClaimRecordIndexKey() {
	testCases := []struct {
		airdropId     uint64
		recipientAddr sdk.AccAddress
		expected      []byte
	}{
		{
			uint64(1),
			sdk.AccAddress(crypto.AddressHash([]byte("recipient1"))),
			[]byte{0xd6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x14, 0x8f, 0x8b, 0xd, 0x57,
				0x5, 0x97, 0xbd, 0x4, 0x45, 0x83, 0x45, 0x14, 0x91, 0x94, 0x96, 0xf8, 0x5c, 0xda, 0x75, 0xfb},
		},
		{
			uint64(3),
			sdk.AccAddress(crypto.AddressHash([]byte("recipient2"))),
			[]byte{0xd6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0x14, 0xe8, 0x67, 0x48, 0x2e,
				0x7d, 0x8a, 0xa3, 0x64, 0x51, 0x28, 0x10, 0x50, 0x5b, 0xe4, 0x38, 0x56, 0x35, 0x81, 0xdc, 0xe0},
		},
		{
			uint64(10),
			sdk.AccAddress(crypto.AddressHash([]byte("recipient3"))),
			[]byte{0xd6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x14, 0x3a, 0x9d, 0xab, 0xd,
				0x59, 0xb1, 0x8e, 0xdc, 0xea, 0x29, 0x8e, 0x9c, 0xf8, 0xb, 0x14, 0x6, 0x95, 0xbc, 0x39, 0x95},
		},
	}

	for _, tc := range testCases {
		key := types.GetClaimRecordKey(tc.airdropId, tc.recipientAddr)
		s.Require().Equal(tc.expected, key)
	}
}
