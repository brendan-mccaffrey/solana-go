// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package raydium_amm

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey

func SetProgramID(pubkey ag_solanago.PublicKey) {
	ProgramID = pubkey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "RaydiumAmm"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_Initialize = ag_binary.TypeID([8]byte{175, 175, 109, 31, 13, 152, 155, 237})

	Instruction_Initialize2 = ag_binary.TypeID([8]byte{9, 203, 254, 64, 89, 32, 179, 159})

	Instruction_MonitorStep = ag_binary.TypeID([8]byte{252, 219, 18, 48, 87, 183, 26, 154})

	Instruction_Deposit = ag_binary.TypeID([8]byte{242, 35, 198, 137, 82, 225, 242, 182})

	Instruction_Withdraw = ag_binary.TypeID([8]byte{183, 18, 70, 156, 148, 109, 161, 34})

	Instruction_MigrateToOpenBook = ag_binary.TypeID([8]byte{207, 98, 243, 89, 114, 174, 205, 20})

	Instruction_SetParams = ag_binary.TypeID([8]byte{27, 234, 178, 52, 147, 2, 187, 141})

	Instruction_WithdrawPnl = ag_binary.TypeID([8]byte{86, 36, 158, 158, 92, 241, 251, 94})

	Instruction_WithdrawSrm = ag_binary.TypeID([8]byte{193, 101, 58, 65, 120, 78, 99, 31})

	Instruction_SwapBaseIn = ag_binary.TypeID([8]byte{42, 236, 72, 162, 242, 24, 39, 84})

	Instruction_PreInitialize = ag_binary.TypeID([8]byte{255, 92, 87, 45, 198, 172, 236, 2})

	Instruction_SwapBaseOut = ag_binary.TypeID([8]byte{163, 210, 155, 208, 175, 146, 213, 150})

	Instruction_SimulateInfo = ag_binary.TypeID([8]byte{195, 75, 104, 72, 253, 176, 183, 160})

	Instruction_AdminCancelOrders = ag_binary.TypeID([8]byte{151, 90, 110, 217, 196, 223, 251, 95})

	Instruction_CreateConfigAccount = ag_binary.TypeID([8]byte{190, 227, 122, 84, 73, 166, 40, 100})

	Instruction_UpdateConfigAccount = ag_binary.TypeID([8]byte{240, 32, 10, 152, 8, 45, 87, 58})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_Initialize:
		return "Initialize"
	case Instruction_Initialize2:
		return "Initialize2"
	case Instruction_MonitorStep:
		return "MonitorStep"
	case Instruction_Deposit:
		return "Deposit"
	case Instruction_Withdraw:
		return "Withdraw"
	case Instruction_MigrateToOpenBook:
		return "MigrateToOpenBook"
	case Instruction_SetParams:
		return "SetParams"
	case Instruction_WithdrawPnl:
		return "WithdrawPnl"
	case Instruction_WithdrawSrm:
		return "WithdrawSrm"
	case Instruction_SwapBaseIn:
		return "SwapBaseIn"
	case Instruction_PreInitialize:
		return "PreInitialize"
	case Instruction_SwapBaseOut:
		return "SwapBaseOut"
	case Instruction_SimulateInfo:
		return "SimulateInfo"
	case Instruction_AdminCancelOrders:
		return "AdminCancelOrders"
	case Instruction_CreateConfigAccount:
		return "CreateConfigAccount"
	case Instruction_UpdateConfigAccount:
		return "UpdateConfigAccount"
	default:
		return ""
	}
}

type Instruction struct {
	ag_binary.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent ag_treeout.Branches) {
	if enToTree, ok := inst.Impl.(ag_text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(ag_spew.Sdump(inst))
	}
}

var InstructionImplDef = ag_binary.NewVariantDefinition(
	ag_binary.AnchorTypeIDEncoding,
	[]ag_binary.VariantType{
		{
			"initialize", (*Initialize)(nil),
		},
		{
			"initialize2", (*Initialize2)(nil),
		},
		{
			"monitor_step", (*MonitorStep)(nil),
		},
		{
			"deposit", (*Deposit)(nil),
		},
		{
			"withdraw", (*Withdraw)(nil),
		},
		{
			"migrate_to_open_book", (*MigrateToOpenBook)(nil),
		},
		{
			"set_params", (*SetParams)(nil),
		},
		{
			"withdraw_pnl", (*WithdrawPnl)(nil),
		},
		{
			"withdraw_srm", (*WithdrawSrm)(nil),
		},
		{
			"swap_base_in", (*SwapBaseIn)(nil),
		},
		{
			"pre_initialize", (*PreInitialize)(nil),
		},
		{
			"swap_base_out", (*SwapBaseOut)(nil),
		},
		{
			"simulate_info", (*SimulateInfo)(nil),
		},
		{
			"admin_cancel_orders", (*AdminCancelOrders)(nil),
		},
		{
			"create_config_account", (*CreateConfigAccount)(nil),
		},
		{
			"update_config_account", (*UpdateConfigAccount)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() ag_solanago.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*ag_solanago.AccountMeta) {
	return inst.Impl.(ag_solanago.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ag_binary.NewBorshEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) TextEncode(encoder *ag_text.Encoder, option *ag_text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst *Instruction) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	err := encoder.WriteBytes(inst.TypeID.Bytes(), false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := ag_binary.NewBorshDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(ag_solanago.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}
