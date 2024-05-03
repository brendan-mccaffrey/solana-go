// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package raydium_amm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SimulateInfo is the `simulateInfo` instruction.
type SimulateInfo struct {
	Param            *uint8
	SwapBaseInValue  *SwapInstructionBaseIn  `bin:"optional"`
	SwapBaseOutValue *SwapInstructionBaseOut `bin:"optional"`

	// [0] = [] amm
	//
	// [1] = [] ammAuthority
	//
	// [2] = [] ammOpenOrders
	//
	// [3] = [] poolCoinTokenAccount
	//
	// [4] = [] poolPcTokenAccount
	//
	// [5] = [] lpMintAddress
	//
	// [6] = [] serumMarket
	//
	// [7] = [] serumEventQueue
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSimulateInfoInstructionBuilder creates a new `SimulateInfo` instruction builder.
func NewSimulateInfoInstructionBuilder() *SimulateInfo {
	nd := &SimulateInfo{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetParam sets the "param" parameter.
func (inst *SimulateInfo) SetParam(param uint8) *SimulateInfo {
	inst.Param = &param
	return inst
}

// SetSwapBaseInValue sets the "swapBaseInValue" parameter.
func (inst *SimulateInfo) SetSwapBaseInValue(swapBaseInValue SwapInstructionBaseIn) *SimulateInfo {
	inst.SwapBaseInValue = &swapBaseInValue
	return inst
}

// SetSwapBaseOutValue sets the "swapBaseOutValue" parameter.
func (inst *SimulateInfo) SetSwapBaseOutValue(swapBaseOutValue SwapInstructionBaseOut) *SimulateInfo {
	inst.SwapBaseOutValue = &swapBaseOutValue
	return inst
}

// SetAmmAccount sets the "amm" account.
func (inst *SimulateInfo) SetAmmAccount(amm ag_solanago.PublicKey) *SimulateInfo {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(amm)
	return inst
}

// GetAmmAccount gets the "amm" account.
func (inst *SimulateInfo) GetAmmAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAmmAuthorityAccount sets the "ammAuthority" account.
func (inst *SimulateInfo) SetAmmAuthorityAccount(ammAuthority ag_solanago.PublicKey) *SimulateInfo {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(ammAuthority)
	return inst
}

// GetAmmAuthorityAccount gets the "ammAuthority" account.
func (inst *SimulateInfo) GetAmmAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAmmOpenOrdersAccount sets the "ammOpenOrders" account.
func (inst *SimulateInfo) SetAmmOpenOrdersAccount(ammOpenOrders ag_solanago.PublicKey) *SimulateInfo {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(ammOpenOrders)
	return inst
}

// GetAmmOpenOrdersAccount gets the "ammOpenOrders" account.
func (inst *SimulateInfo) GetAmmOpenOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPoolCoinTokenAccountAccount sets the "poolCoinTokenAccount" account.
func (inst *SimulateInfo) SetPoolCoinTokenAccountAccount(poolCoinTokenAccount ag_solanago.PublicKey) *SimulateInfo {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(poolCoinTokenAccount)
	return inst
}

// GetPoolCoinTokenAccountAccount gets the "poolCoinTokenAccount" account.
func (inst *SimulateInfo) GetPoolCoinTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPoolPcTokenAccountAccount sets the "poolPcTokenAccount" account.
func (inst *SimulateInfo) SetPoolPcTokenAccountAccount(poolPcTokenAccount ag_solanago.PublicKey) *SimulateInfo {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(poolPcTokenAccount)
	return inst
}

// GetPoolPcTokenAccountAccount gets the "poolPcTokenAccount" account.
func (inst *SimulateInfo) GetPoolPcTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetLpMintAddressAccount sets the "lpMintAddress" account.
func (inst *SimulateInfo) SetLpMintAddressAccount(lpMintAddress ag_solanago.PublicKey) *SimulateInfo {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(lpMintAddress)
	return inst
}

// GetLpMintAddressAccount gets the "lpMintAddress" account.
func (inst *SimulateInfo) GetLpMintAddressAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSerumMarketAccount sets the "serumMarket" account.
func (inst *SimulateInfo) SetSerumMarketAccount(serumMarket ag_solanago.PublicKey) *SimulateInfo {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(serumMarket)
	return inst
}

// GetSerumMarketAccount gets the "serumMarket" account.
func (inst *SimulateInfo) GetSerumMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSerumEventQueueAccount sets the "serumEventQueue" account.
func (inst *SimulateInfo) SetSerumEventQueueAccount(serumEventQueue ag_solanago.PublicKey) *SimulateInfo {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(serumEventQueue)
	return inst
}

// GetSerumEventQueueAccount gets the "serumEventQueue" account.
func (inst *SimulateInfo) GetSerumEventQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst SimulateInfo) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SimulateInfo,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SimulateInfo) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SimulateInfo) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Param == nil {
			return errors.New("Param parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Amm is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.AmmAuthority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.AmmOpenOrders is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PoolCoinTokenAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.PoolPcTokenAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.LpMintAddress is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SerumMarket is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SerumEventQueue is not set")
		}
	}
	return nil
}

func (inst *SimulateInfo) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SimulateInfo")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("           Param", *inst.Param))
						paramsBranch.Child(ag_format.Param(" SwapBaseInValue (OPT)", inst.SwapBaseInValue))
						paramsBranch.Child(ag_format.Param("SwapBaseOutValue (OPT)", inst.SwapBaseOutValue))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("            amm", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   ammAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("  ammOpenOrders", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("  poolCoinToken", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("    poolPcToken", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("  lpMintAddress", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("    serumMarket", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("serumEventQueue", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj SimulateInfo) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Param` param:
	err = encoder.Encode(obj.Param)
	if err != nil {
		return err
	}
	// Serialize `SwapBaseInValue` param (optional):
	{
		if obj.SwapBaseInValue == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.SwapBaseInValue)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `SwapBaseOutValue` param (optional):
	{
		if obj.SwapBaseOutValue == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.SwapBaseOutValue)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (obj *SimulateInfo) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Param`:
	err = decoder.Decode(&obj.Param)
	if err != nil {
		return err
	}
	// Deserialize `SwapBaseInValue` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.SwapBaseInValue)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `SwapBaseOutValue` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.SwapBaseOutValue)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewSimulateInfoInstruction declares a new SimulateInfo instruction with the provided parameters and accounts.
func NewSimulateInfoInstruction(
	// Parameters:
	param uint8,
	swapBaseInValue SwapInstructionBaseIn,
	swapBaseOutValue SwapInstructionBaseOut,
	// Accounts:
	amm ag_solanago.PublicKey,
	ammAuthority ag_solanago.PublicKey,
	ammOpenOrders ag_solanago.PublicKey,
	poolCoinTokenAccount ag_solanago.PublicKey,
	poolPcTokenAccount ag_solanago.PublicKey,
	lpMintAddress ag_solanago.PublicKey,
	serumMarket ag_solanago.PublicKey,
	serumEventQueue ag_solanago.PublicKey) *SimulateInfo {
	return NewSimulateInfoInstructionBuilder().
		SetParam(param).
		SetSwapBaseInValue(swapBaseInValue).
		SetSwapBaseOutValue(swapBaseOutValue).
		SetAmmAccount(amm).
		SetAmmAuthorityAccount(ammAuthority).
		SetAmmOpenOrdersAccount(ammOpenOrders).
		SetPoolCoinTokenAccountAccount(poolCoinTokenAccount).
		SetPoolPcTokenAccountAccount(poolPcTokenAccount).
		SetLpMintAddressAccount(lpMintAddress).
		SetSerumMarketAccount(serumMarket).
		SetSerumEventQueueAccount(serumEventQueue)
}
