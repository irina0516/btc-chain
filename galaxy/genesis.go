package galaxy

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/galaxy-digital/lachesis-base/hash"
	"github.com/galaxy-digital/lachesis-base/inter/idx"

	"github.com/galaxy-digital/galaxychain-chain/galaxy/genesis"
	"github.com/galaxy-digital/galaxychain-chain/galaxy/genesis/gpos"
	"github.com/galaxy-digital/galaxychain-chain/inter"
)

type Genesis struct {
	Accounts    genesis.Accounts
	Storage     genesis.Storage
	Delegations genesis.Delegations
	Blocks      genesis.Blocks
	RawEvmItems genesis.RawEvmItems
	Validators  gpos.Validators

	FirstEpoch    idx.Epoch
	PrevEpochTime inter.Timestamp
	Time          inter.Timestamp
	ExtraData     []byte

	TotalSupply *big.Int

	DriverOwner common.Address

	Rules Rules

	Hash func() hash.Hash
}
