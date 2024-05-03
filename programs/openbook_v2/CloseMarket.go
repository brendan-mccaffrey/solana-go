// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package openbook_v2

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Close a [`Market`](crate::state::Market) (only
// [`close_market_admin`](crate::state::Market::close_market_admin)).
type CloseMarket struct {

	// [0] = [SIGNER] closeMarketAdmin
	//
	// [1] = [WRITE] market
	//
	// [2] = [WRITE] bids
	//
	// [3] = [WRITE] asks
	//
	// [4] = [WRITE] eventHeap
	//
	// [5] = [WRITE] solDestination
	//
	// [6] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCloseMarketInstructionBuilder creates a new `CloseMarket` instruction builder.
func NewCloseMarketInstructionBuilder() *CloseMarket {
	nd := &CloseMarket{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetCloseMarketAdminAccount sets the "closeMarketAdmin" account.
func (inst *CloseMarket) SetCloseMarketAdminAccount(closeMarketAdmin ag_solanago.PublicKey) *CloseMarket {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(closeMarketAdmin).SIGNER()
	return inst
}

// GetCloseMarketAdminAccount gets the "closeMarketAdmin" account.
func (inst *CloseMarket) GetCloseMarketAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMarketAccount sets the "market" account.
func (inst *CloseMarket) SetMarketAccount(market ag_solanago.PublicKey) *CloseMarket {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *CloseMarket) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetBidsAccount sets the "bids" account.
func (inst *CloseMarket) SetBidsAccount(bids ag_solanago.PublicKey) *CloseMarket {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(bids).WRITE()
	return inst
}

// GetBidsAccount gets the "bids" account.
func (inst *CloseMarket) GetBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAsksAccount sets the "asks" account.
func (inst *CloseMarket) SetAsksAccount(asks ag_solanago.PublicKey) *CloseMarket {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(asks).WRITE()
	return inst
}

// GetAsksAccount gets the "asks" account.
func (inst *CloseMarket) GetAsksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetEventHeapAccount sets the "eventHeap" account.
func (inst *CloseMarket) SetEventHeapAccount(eventHeap ag_solanago.PublicKey) *CloseMarket {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(eventHeap).WRITE()
	return inst
}

// GetEventHeapAccount gets the "eventHeap" account.
func (inst *CloseMarket) GetEventHeapAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSolDestinationAccount sets the "solDestination" account.
func (inst *CloseMarket) SetSolDestinationAccount(solDestination ag_solanago.PublicKey) *CloseMarket {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(solDestination).WRITE()
	return inst
}

// GetSolDestinationAccount gets the "solDestination" account.
func (inst *CloseMarket) GetSolDestinationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *CloseMarket) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *CloseMarket {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *CloseMarket) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst CloseMarket) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CloseMarket,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CloseMarket) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CloseMarket) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CloseMarketAdmin is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Market is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Bids is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Asks is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.EventHeap is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.SolDestination is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *CloseMarket) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CloseMarket")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("closeMarketAdmin", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          market", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("            bids", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("            asks", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       eventHeap", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("  solDestination", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("    tokenProgram", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj CloseMarket) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *CloseMarket) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewCloseMarketInstruction declares a new CloseMarket instruction with the provided parameters and accounts.
func NewCloseMarketInstruction(
	// Accounts:
	closeMarketAdmin ag_solanago.PublicKey,
	market ag_solanago.PublicKey,
	bids ag_solanago.PublicKey,
	asks ag_solanago.PublicKey,
	eventHeap ag_solanago.PublicKey,
	solDestination ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *CloseMarket {
	return NewCloseMarketInstructionBuilder().
		SetCloseMarketAdminAccount(closeMarketAdmin).
		SetMarketAccount(market).
		SetBidsAccount(bids).
		SetAsksAccount(asks).
		SetEventHeapAccount(eventHeap).
		SetSolDestinationAccount(solDestination).
		SetTokenProgramAccount(tokenProgram)
}