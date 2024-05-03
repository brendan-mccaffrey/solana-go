// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package raydium_amm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// MigrateToOpenBook is the `migrateToOpenBook` instruction.
type MigrateToOpenBook struct {

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
	// [6] = [WRITE] ammTokenCoin
	//
	// [7] = [WRITE] ammTokenPc
	//
	// [8] = [WRITE] ammTargetOrders
	//
	// [9] = [] serumProgram
	//
	// [10] = [WRITE] serumMarket
	//
	// [11] = [WRITE] serumBids
	//
	// [12] = [WRITE] serumAsks
	//
	// [13] = [WRITE] serumEventQueue
	//
	// [14] = [WRITE] serumCoinVault
	//
	// [15] = [WRITE] serumPcVault
	//
	// [16] = [] serumVaultSigner
	//
	// [17] = [WRITE] newAmmOpenOrders
	//
	// [18] = [] newSerumProgram
	//
	// [19] = [] newSerumMarket
	//
	// [20] = [WRITE, SIGNER] admin
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewMigrateToOpenBookInstructionBuilder creates a new `MigrateToOpenBook` instruction builder.
func NewMigrateToOpenBookInstructionBuilder() *MigrateToOpenBook {
	nd := &MigrateToOpenBook{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 21),
	}
	return nd
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *MigrateToOpenBook) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *MigrateToOpenBook) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *MigrateToOpenBook) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *MigrateToOpenBook) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetRentAccount sets the "rent" account.
func (inst *MigrateToOpenBook) SetRentAccount(rent ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *MigrateToOpenBook) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAmmAccount sets the "amm" account.
func (inst *MigrateToOpenBook) SetAmmAccount(amm ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(amm).WRITE()
	return inst
}

// GetAmmAccount gets the "amm" account.
func (inst *MigrateToOpenBook) GetAmmAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAmmAuthorityAccount sets the "ammAuthority" account.
func (inst *MigrateToOpenBook) SetAmmAuthorityAccount(ammAuthority ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(ammAuthority)
	return inst
}

// GetAmmAuthorityAccount gets the "ammAuthority" account.
func (inst *MigrateToOpenBook) GetAmmAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetAmmOpenOrdersAccount sets the "ammOpenOrders" account.
func (inst *MigrateToOpenBook) SetAmmOpenOrdersAccount(ammOpenOrders ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(ammOpenOrders).WRITE()
	return inst
}

// GetAmmOpenOrdersAccount gets the "ammOpenOrders" account.
func (inst *MigrateToOpenBook) GetAmmOpenOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetAmmTokenCoinAccount sets the "ammTokenCoin" account.
func (inst *MigrateToOpenBook) SetAmmTokenCoinAccount(ammTokenCoin ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(ammTokenCoin).WRITE()
	return inst
}

// GetAmmTokenCoinAccount gets the "ammTokenCoin" account.
func (inst *MigrateToOpenBook) GetAmmTokenCoinAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetAmmTokenPcAccount sets the "ammTokenPc" account.
func (inst *MigrateToOpenBook) SetAmmTokenPcAccount(ammTokenPc ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(ammTokenPc).WRITE()
	return inst
}

// GetAmmTokenPcAccount gets the "ammTokenPc" account.
func (inst *MigrateToOpenBook) GetAmmTokenPcAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetAmmTargetOrdersAccount sets the "ammTargetOrders" account.
func (inst *MigrateToOpenBook) SetAmmTargetOrdersAccount(ammTargetOrders ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(ammTargetOrders).WRITE()
	return inst
}

// GetAmmTargetOrdersAccount gets the "ammTargetOrders" account.
func (inst *MigrateToOpenBook) GetAmmTargetOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSerumProgramAccount sets the "serumProgram" account.
func (inst *MigrateToOpenBook) SetSerumProgramAccount(serumProgram ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(serumProgram)
	return inst
}

// GetSerumProgramAccount gets the "serumProgram" account.
func (inst *MigrateToOpenBook) GetSerumProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetSerumMarketAccount sets the "serumMarket" account.
func (inst *MigrateToOpenBook) SetSerumMarketAccount(serumMarket ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(serumMarket).WRITE()
	return inst
}

// GetSerumMarketAccount gets the "serumMarket" account.
func (inst *MigrateToOpenBook) GetSerumMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetSerumBidsAccount sets the "serumBids" account.
func (inst *MigrateToOpenBook) SetSerumBidsAccount(serumBids ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(serumBids).WRITE()
	return inst
}

// GetSerumBidsAccount gets the "serumBids" account.
func (inst *MigrateToOpenBook) GetSerumBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetSerumAsksAccount sets the "serumAsks" account.
func (inst *MigrateToOpenBook) SetSerumAsksAccount(serumAsks ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(serumAsks).WRITE()
	return inst
}

// GetSerumAsksAccount gets the "serumAsks" account.
func (inst *MigrateToOpenBook) GetSerumAsksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetSerumEventQueueAccount sets the "serumEventQueue" account.
func (inst *MigrateToOpenBook) SetSerumEventQueueAccount(serumEventQueue ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(serumEventQueue).WRITE()
	return inst
}

// GetSerumEventQueueAccount gets the "serumEventQueue" account.
func (inst *MigrateToOpenBook) GetSerumEventQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetSerumCoinVaultAccount sets the "serumCoinVault" account.
func (inst *MigrateToOpenBook) SetSerumCoinVaultAccount(serumCoinVault ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(serumCoinVault).WRITE()
	return inst
}

// GetSerumCoinVaultAccount gets the "serumCoinVault" account.
func (inst *MigrateToOpenBook) GetSerumCoinVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetSerumPcVaultAccount sets the "serumPcVault" account.
func (inst *MigrateToOpenBook) SetSerumPcVaultAccount(serumPcVault ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(serumPcVault).WRITE()
	return inst
}

// GetSerumPcVaultAccount gets the "serumPcVault" account.
func (inst *MigrateToOpenBook) GetSerumPcVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetSerumVaultSignerAccount sets the "serumVaultSigner" account.
func (inst *MigrateToOpenBook) SetSerumVaultSignerAccount(serumVaultSigner ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(serumVaultSigner)
	return inst
}

// GetSerumVaultSignerAccount gets the "serumVaultSigner" account.
func (inst *MigrateToOpenBook) GetSerumVaultSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetNewAmmOpenOrdersAccount sets the "newAmmOpenOrders" account.
func (inst *MigrateToOpenBook) SetNewAmmOpenOrdersAccount(newAmmOpenOrders ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(newAmmOpenOrders).WRITE()
	return inst
}

// GetNewAmmOpenOrdersAccount gets the "newAmmOpenOrders" account.
func (inst *MigrateToOpenBook) GetNewAmmOpenOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

// SetNewSerumProgramAccount sets the "newSerumProgram" account.
func (inst *MigrateToOpenBook) SetNewSerumProgramAccount(newSerumProgram ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(newSerumProgram)
	return inst
}

// GetNewSerumProgramAccount gets the "newSerumProgram" account.
func (inst *MigrateToOpenBook) GetNewSerumProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(18)
}

// SetNewSerumMarketAccount sets the "newSerumMarket" account.
func (inst *MigrateToOpenBook) SetNewSerumMarketAccount(newSerumMarket ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[19] = ag_solanago.Meta(newSerumMarket)
	return inst
}

// GetNewSerumMarketAccount gets the "newSerumMarket" account.
func (inst *MigrateToOpenBook) GetNewSerumMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(19)
}

// SetAdminAccount sets the "admin" account.
func (inst *MigrateToOpenBook) SetAdminAccount(admin ag_solanago.PublicKey) *MigrateToOpenBook {
	inst.AccountMetaSlice[20] = ag_solanago.Meta(admin).WRITE().SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *MigrateToOpenBook) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(20)
}

func (inst MigrateToOpenBook) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_MigrateToOpenBook,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MigrateToOpenBook) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MigrateToOpenBook) Validate() error {
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
			return errors.New("accounts.AmmTokenCoin is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.AmmTokenPc is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.AmmTargetOrders is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.SerumProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.SerumMarket is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.SerumBids is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.SerumAsks is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.SerumEventQueue is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.SerumCoinVault is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.SerumPcVault is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.SerumVaultSigner is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.NewAmmOpenOrders is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.NewSerumProgram is not set")
		}
		if inst.AccountMetaSlice[19] == nil {
			return errors.New("accounts.NewSerumMarket is not set")
		}
		if inst.AccountMetaSlice[20] == nil {
			return errors.New("accounts.Admin is not set")
		}
	}
	return nil
}

func (inst *MigrateToOpenBook) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MigrateToOpenBook")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=21]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    tokenProgram", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   systemProgram", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("            rent", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("             amm", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("    ammAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("   ammOpenOrders", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("    ammTokenCoin", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("      ammTokenPc", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta(" ammTargetOrders", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("    serumProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("     serumMarket", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("       serumBids", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("       serumAsks", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta(" serumEventQueue", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("  serumCoinVault", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("    serumPcVault", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("serumVaultSigner", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("newAmmOpenOrders", inst.AccountMetaSlice.Get(17)))
						accountsBranch.Child(ag_format.Meta(" newSerumProgram", inst.AccountMetaSlice.Get(18)))
						accountsBranch.Child(ag_format.Meta("  newSerumMarket", inst.AccountMetaSlice.Get(19)))
						accountsBranch.Child(ag_format.Meta("           admin", inst.AccountMetaSlice.Get(20)))
					})
				})
		})
}

func (obj MigrateToOpenBook) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *MigrateToOpenBook) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewMigrateToOpenBookInstruction declares a new MigrateToOpenBook instruction with the provided parameters and accounts.
func NewMigrateToOpenBookInstruction(
	// Accounts:
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	amm ag_solanago.PublicKey,
	ammAuthority ag_solanago.PublicKey,
	ammOpenOrders ag_solanago.PublicKey,
	ammTokenCoin ag_solanago.PublicKey,
	ammTokenPc ag_solanago.PublicKey,
	ammTargetOrders ag_solanago.PublicKey,
	serumProgram ag_solanago.PublicKey,
	serumMarket ag_solanago.PublicKey,
	serumBids ag_solanago.PublicKey,
	serumAsks ag_solanago.PublicKey,
	serumEventQueue ag_solanago.PublicKey,
	serumCoinVault ag_solanago.PublicKey,
	serumPcVault ag_solanago.PublicKey,
	serumVaultSigner ag_solanago.PublicKey,
	newAmmOpenOrders ag_solanago.PublicKey,
	newSerumProgram ag_solanago.PublicKey,
	newSerumMarket ag_solanago.PublicKey,
	admin ag_solanago.PublicKey) *MigrateToOpenBook {
	return NewMigrateToOpenBookInstructionBuilder().
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent).
		SetAmmAccount(amm).
		SetAmmAuthorityAccount(ammAuthority).
		SetAmmOpenOrdersAccount(ammOpenOrders).
		SetAmmTokenCoinAccount(ammTokenCoin).
		SetAmmTokenPcAccount(ammTokenPc).
		SetAmmTargetOrdersAccount(ammTargetOrders).
		SetSerumProgramAccount(serumProgram).
		SetSerumMarketAccount(serumMarket).
		SetSerumBidsAccount(serumBids).
		SetSerumAsksAccount(serumAsks).
		SetSerumEventQueueAccount(serumEventQueue).
		SetSerumCoinVaultAccount(serumCoinVault).
		SetSerumPcVaultAccount(serumPcVault).
		SetSerumVaultSignerAccount(serumVaultSigner).
		SetNewAmmOpenOrdersAccount(newAmmOpenOrders).
		SetNewSerumProgramAccount(newSerumProgram).
		SetNewSerumMarketAccount(newSerumMarket).
		SetAdminAccount(admin)
}
