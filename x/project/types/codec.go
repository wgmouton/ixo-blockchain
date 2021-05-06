package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateProject{}, "project/CreateProject", nil)
	cdc.RegisterConcrete(MsgUpdateProjectStatus{}, "project/UpdateProjectStatus", nil)
	cdc.RegisterConcrete(MsgCreateAgent{}, "project/CreateAgent", nil)
	cdc.RegisterConcrete(MsgUpdateAgent{}, "project/UpdateAgent", nil)
	cdc.RegisterConcrete(MsgCreateClaim{}, "project/CreateClaim", nil)
	cdc.RegisterConcrete(MsgCreateEvaluation{}, "project/CreateEvaluation", nil)
	cdc.RegisterConcrete(MsgWithdrawFunds{}, "project/WithdrawFunds", nil)

	cdc.RegisterConcrete(ProjectDoc{}, "project/ProjectDoc", nil)
	cdc.RegisterConcrete(AccountMap{}, "project/AccountMap", nil)
	cdc.RegisterConcrete(WithdrawalInfoDoc{}, "project/WithdrawalInfo", nil)
}

// ModuleCdc is the codec for the module
var ModuleCdc *codec.LegacyAmino

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}
