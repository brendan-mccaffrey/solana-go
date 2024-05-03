// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package openbook_v2

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Process the [events](crate::state::AnyEvent) at the given positions.
type ConsumeGivenEvents struct {
	Slots *[]uint64

	// [0] = [SIGNER] consumeEventsAdmin
	//
	// [1] = [WRITE] market
	//
	// [2] = [WRITE] eventHeap
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewConsumeGivenEventsInstructionBuilder creates a new `ConsumeGivenEvents` instruction builder.
func NewConsumeGivenEventsInstructionBuilder() *ConsumeGivenEvents {
	nd := &ConsumeGivenEvents{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetSlots sets the "slots" parameter.
func (inst *ConsumeGivenEvents) SetSlots(slots []uint64) *ConsumeGivenEvents {
	inst.Slots = &slots
	return inst
}

// SetConsumeEventsAdminAccount sets the "consumeEventsAdmin" account.
func (inst *ConsumeGivenEvents) SetConsumeEventsAdminAccount(consumeEventsAdmin ag_solanago.PublicKey) *ConsumeGivenEvents {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(consumeEventsAdmin).SIGNER()
	return inst
}

// GetConsumeEventsAdminAccount gets the "consumeEventsAdmin" account.
func (inst *ConsumeGivenEvents) GetConsumeEventsAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMarketAccount sets the "market" account.
func (inst *ConsumeGivenEvents) SetMarketAccount(market ag_solanago.PublicKey) *ConsumeGivenEvents {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *ConsumeGivenEvents) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetEventHeapAccount sets the "eventHeap" account.
func (inst *ConsumeGivenEvents) SetEventHeapAccount(eventHeap ag_solanago.PublicKey) *ConsumeGivenEvents {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(eventHeap).WRITE()
	return inst
}

// GetEventHeapAccount gets the "eventHeap" account.
func (inst *ConsumeGivenEvents) GetEventHeapAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst ConsumeGivenEvents) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ConsumeGivenEvents,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ConsumeGivenEvents) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ConsumeGivenEvents) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Slots == nil {
			return errors.New("Slots parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.ConsumeEventsAdmin is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Market is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.EventHeap is not set")
		}
	}
	return nil
}

func (inst *ConsumeGivenEvents) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ConsumeGivenEvents")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Slots", *inst.Slots))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("consumeEventsAdmin", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("            market", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         eventHeap", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj ConsumeGivenEvents) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Slots` param:
	err = encoder.Encode(obj.Slots)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ConsumeGivenEvents) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Slots`:
	err = decoder.Decode(&obj.Slots)
	if err != nil {
		return err
	}
	return nil
}

// NewConsumeGivenEventsInstruction declares a new ConsumeGivenEvents instruction with the provided parameters and accounts.
func NewConsumeGivenEventsInstruction(
	// Parameters:
	slots []uint64,
	// Accounts:
	consumeEventsAdmin ag_solanago.PublicKey,
	market ag_solanago.PublicKey,
	eventHeap ag_solanago.PublicKey) *ConsumeGivenEvents {
	return NewConsumeGivenEventsInstructionBuilder().
		SetSlots(slots).
		SetConsumeEventsAdminAccount(consumeEventsAdmin).
		SetMarketAccount(market).
		SetEventHeapAccount(eventHeap)
}
