package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://418678a555a0c4c261aa98c3319768016e78589b6b3865d1010901860aa9bd945f073b5a423f27b3b204710331eb5b611eecf3da367632830d7bb2dabf5dc296@85.206.160.134:15060",
		"enode://d9c487cd09f09990ff8eba01b4879dfa64d8153e218ab492380143f1457e6b05a3e9501c9a1976b3904e1d738980e18e769d7e327fd478239ddba5fa2fcbc8e2@85.206.160.135:15060",
		"enode://9c68b72d73832b37c5853a19ad4ae3355f0fd1a4b094dd27179773080faffb133f95623e8c409d20741415122c58d06a080365b2094ef7f4658f1448ea5d768a@185.25.48.207:15060",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
