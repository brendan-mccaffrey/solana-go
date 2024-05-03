// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package raydium_amm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Initialize is the `initialize` instruction.
type Initialize struct {
	Nonce    *uint8
	OpenTime *uint64

	// [0] = [] tokenProgram
	//
	// [1] = [] systemProgram
	//
	// [2] = [] rent
	//
	// [3] = [WRITE] amm
	//
	// [4] = [] ammAuthority
	//
	// [5] = [WRITE] ammOpenOrders
	//
	// [6] = [WRITE] lpMintAddress
	//
	// [7] = [] coinMintAddress
	//
	// [8] = [] pcMintAddress
	//
	// [9] = [] poolCoinTokenAccount
	//
	// [10] = [] poolPcTokenAccount
	//
	// [11] = [WRITE] poolWithdrawQueue
	//
	// [12] = [WRITE] poolTargetOrdersAccount
	//
	// [13] = [WRITE] userLpTokenAccount
	//
	// [14] = [] poolTempLpTokenAccount
	//
	// [15] = [] serumProgram
	//
	// [16] = [] serumMarket
	//
	// [17] = [WRITE, SIGNER] userWallet
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	nd := &Initialize{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 18),
	}
	return nd
}

// SetNonce sets the "nonce" parameter.
func (inst *Initialize) SetNonce(nonce uint8) *Initialize {
	inst.Nonce = &nonce
	return inst
}

// SetOpenTime sets the "openTime" parameter.
func (inst *Initialize) SetOpenTime(openTime uint64) *Initialize {
	inst.OpenTime = &openTime
	return inst
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *Initialize) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *Initialize) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *Initialize) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *Initialize) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetRentAccount sets the "rent" account.
func (inst *Initialize) SetRentAccount(rent ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *Initialize) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAmmAccount sets the "amm" account.
func (inst *Initialize) SetAmmAccount(amm ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(amm).WRITE()
	return inst
}

// GetAmmAccount gets the "amm" account.
func (inst *Initialize) GetAmmAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAmmAuthorityAccount sets the "ammAuthority" account.
func (inst *Initialize) SetAmmAuthorityAccount(ammAuthority ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(ammAuthority)
	return inst
}

// GetAmmAuthorityAccount gets the "ammAuthority" account.
func (inst *Initialize) GetAmmAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetAmmOpenOrdersAccount sets the "ammOpenOrders" account.
func (inst *Initialize) SetAmmOpenOrdersAccount(ammOpenOrders ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(ammOpenOrders).WRITE()
	return inst
}

// GetAmmOpenOrdersAccount gets the "ammOpenOrders" account.
func (inst *Initialize) GetAmmOpenOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetLpMintAddressAccount sets the "lpMintAddress" account.
func (inst *Initialize) SetLpMintAddressAccount(lpMintAddress ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(lpMintAddress).WRITE()
	return inst
}

// GetLpMintAddressAccount gets the "lpMintAddress" account.
func (inst *Initialize) GetLpMintAddressAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetCoinMintAddressAccount sets the "coinMintAddress" account.
func (inst *Initialize) SetCoinMintAddressAccount(coinMintAddress ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(coinMintAddress)
	return inst
}

// GetCoinMintAddressAccount gets the "coinMintAddress" account.
func (inst *Initialize) GetCoinMintAddressAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetPcMintAddressAccount sets the "pcMintAddress" account.
func (inst *Initialize) SetPcMintAddressAccount(pcMintAddress ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(pcMintAddress)
	return inst
}

// GetPcMintAddressAccount gets the "pcMintAddress" account.
func (inst *Initialize) GetPcMintAddressAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetPoolCoinTokenAccountAccount sets the "poolCoinTokenAccount" account.
func (inst *Initialize) SetPoolCoinTokenAccountAccount(poolCoinTokenAccount ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(poolCoinTokenAccount)
	return inst
}

// GetPoolCoinTokenAccountAccount gets the "poolCoinTokenAccount" account.
func (inst *Initialize) GetPoolCoinTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetPoolPcTokenAccountAccount sets the "poolPcTokenAccount" account.
func (inst *Initialize) SetPoolPcTokenAccountAccount(poolPcTokenAccount ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(poolPcTokenAccount)
	return inst
}

// GetPoolPcTokenAccountAccount gets the "poolPcTokenAccount" account.
func (inst *Initialize) GetPoolPcTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetPoolWithdrawQueueAccount sets the "poolWithdrawQueue" account.
func (inst *Initialize) SetPoolWithdrawQueueAccount(poolWithdrawQueue ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(poolWithdrawQueue).WRITE()
	return inst
}

// GetPoolWithdrawQueueAccount gets the "poolWithdrawQueue" account.
func (inst *Initialize) GetPoolWithdrawQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetPoolTargetOrdersAccountAccount sets the "poolTargetOrdersAccount" account.
func (inst *Initialize) SetPoolTargetOrdersAccountAccount(poolTargetOrdersAccount ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(poolTargetOrdersAccount).WRITE()
	return inst
}

// GetPoolTargetOrdersAccountAccount gets the "poolTargetOrdersAccount" account.
func (inst *Initialize) GetPoolTargetOrdersAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetUserLpTokenAccountAccount sets the "userLpTokenAccount" account.
func (inst *Initialize) SetUserLpTokenAccountAccount(userLpTokenAccount ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(userLpTokenAccount).WRITE()
	return inst
}

// GetUserLpTokenAccountAccount gets the "userLpTokenAccount" account.
func (inst *Initialize) GetUserLpTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetPoolTempLpTokenAccountAccount sets the "poolTempLpTokenAccount" account.
func (inst *Initialize) SetPoolTempLpTokenAccountAccount(poolTempLpTokenAccount ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(poolTempLpTokenAccount)
	return inst
}

// GetPoolTempLpTokenAccountAccount gets the "poolTempLpTokenAccount" account.
func (inst *Initialize) GetPoolTempLpTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetSerumProgramAccount sets the "serumProgram" account.
func (inst *Initialize) SetSerumProgramAccount(serumProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(serumProgram)
	return inst
}

// GetSerumProgramAccount gets the "serumProgram" account.
func (inst *Initialize) GetSerumProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetSerumMarketAccount sets the "serumMarket" account.
func (inst *Initialize) SetSerumMarketAccount(serumMarket ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(serumMarket)
	return inst
}

// GetSerumMarketAccount gets the "serumMarket" account.
func (inst *Initialize) GetSerumMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetUserWalletAccount sets the "userWallet" account.
func (inst *Initialize) SetUserWalletAccount(userWallet ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(userWallet).WRITE().SIGNER()
	return inst
}

// GetUserWalletAccount gets the "userWallet" account.
func (inst *Initialize) GetUserWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

func (inst Initialize) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Initialize,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Initialize) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Initialize) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Nonce == nil {
			return errors.New("Nonce parameter is not set")
		}
		if inst.OpenTime == nil {
			return errors.New("OpenTime parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Amm is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.AmmAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.AmmOpenOrders is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.LpMintAddress is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.CoinMintAddress is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.PcMintAddress is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.PoolCoinTokenAccount is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.PoolPcTokenAccount is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.PoolWithdrawQueue is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.PoolTargetOrdersAccount is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.UserLpTokenAccount is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.PoolTempLpTokenAccount is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.SerumProgram is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.SerumMarket is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.UserWallet is not set")
		}
	}
	return nil
}

func (inst *Initialize) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Initialize")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("   Nonce", *inst.Nonce))
						paramsBranch.Child(ag_format.Param("OpenTime", *inst.OpenTime))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=18]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     tokenProgram", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    systemProgram", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("             rent", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("              amm", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("     ammAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("    ammOpenOrders", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("    lpMintAddress", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("  coinMintAddress", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("    pcMintAddress", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("    poolCoinToken", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("      poolPcToken", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("poolWithdrawQueue", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta(" poolTargetOrders", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("      userLpToken", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("  poolTempLpToken", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("     serumProgram", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("      serumMarket", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("       userWallet", inst.AccountMetaSlice.Get(17)))
					})
				})
		})
}

func (obj Initialize) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Nonce` param:
	err = encoder.Encode(obj.Nonce)
	if err != nil {
		return err
	}
	// Serialize `OpenTime` param:
	err = encoder.Encode(obj.OpenTime)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Initialize) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Nonce`:
	err = decoder.Decode(&obj.Nonce)
	if err != nil {
		return err
	}
	// Deserialize `OpenTime`:
	err = decoder.Decode(&obj.OpenTime)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeInstruction declares a new Initialize instruction with the provided parameters and accounts.
func NewInitializeInstruction(
	// Parameters:
	nonce uint8,
	openTime uint64,
	// Accounts:
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	amm ag_solanago.PublicKey,
	ammAuthority ag_solanago.PublicKey,
	ammOpenOrders ag_solanago.PublicKey,
	lpMintAddress ag_solanago.PublicKey,
	coinMintAddress ag_solanago.PublicKey,
	pcMintAddress ag_solanago.PublicKey,
	poolCoinTokenAccount ag_solanago.PublicKey,
	poolPcTokenAccount ag_solanago.PublicKey,
	poolWithdrawQueue ag_solanago.PublicKey,
	poolTargetOrdersAccount ag_solanago.PublicKey,
	userLpTokenAccount ag_solanago.PublicKey,
	poolTempLpTokenAccount ag_solanago.PublicKey,
	serumProgram ag_solanago.PublicKey,
	serumMarket ag_solanago.PublicKey,
	userWallet ag_solanago.PublicKey) *Initialize {
	return NewInitializeInstructionBuilder().
		SetNonce(nonce).
		SetOpenTime(openTime).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent).
		SetAmmAccount(amm).
		SetAmmAuthorityAccount(ammAuthority).
		SetAmmOpenOrdersAccount(ammOpenOrders).
		SetLpMintAddressAccount(lpMintAddress).
		SetCoinMintAddressAccount(coinMintAddress).
		SetPcMintAddressAccount(pcMintAddress).
		SetPoolCoinTokenAccountAccount(poolCoinTokenAccount).
		SetPoolPcTokenAccountAccount(poolPcTokenAccount).
		SetPoolWithdrawQueueAccount(poolWithdrawQueue).
		SetPoolTargetOrdersAccountAccount(poolTargetOrdersAccount).
		SetUserLpTokenAccountAccount(userLpTokenAccount).
		SetPoolTempLpTokenAccountAccount(poolTempLpTokenAccount).
		SetSerumProgramAccount(serumProgram).
		SetSerumMarketAccount(serumMarket).
		SetUserWalletAccount(userWallet)
}
