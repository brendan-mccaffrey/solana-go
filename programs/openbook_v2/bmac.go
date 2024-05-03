package openbook_v2

import (
	"github.com/gagliardetto/solana-go"
)

// ConvertTransactionToCreateMarket converts a solana.Transaction to a CreateMarket struct.
func ConvertTransactionToCreateMarket(tx *solana.Transaction) *CreateMarket {
	cm := NewCreateMarketInstructionBuilder()
	ml, err := tx.AccountMetaList()
	if err != nil {
		return nil
	}
	cm.AccountMetaSlice = ml

	// // Map the transaction's account metas to the CreateMarket's AccountMetaSlice
	// for i, meta := range tx.Message.AccountKeys {
	// 	if i < len(cm.AccountMetaSlice) {
	// 		cm.AccountMetaSlice[i] = ml[i]
	// 		//  ag_solanago.AccountMeta{
	// 		// 	PublicKey:  meta,
	// 		// 	IsWritable: tx.Message.IsAccountWritable(i),
	// 		// 	IsSigner:   tx.Message.IsAccountSigner(i),
	// 		// }
	// 	}
	// }

	return cm
}

// func DecodeCr(tx *solana.Transaction) (*CreateMarket, error) {
// 	for _, i := range tx.Message.Instructions {

// 		// Get (for example) the first instruction of this transaction
// 		// which we know is a `system` program instruction:

// 		// Find the program address of this instruction:
// 		progKey, err := tx.ResolveProgramIDIndex(i.ProgramIDIndex)
// 		if err != nil {
// 			panic(err)
// 		}
// 		_ = progKey

// 		// Find the accounts of this instruction:
// 		accounts, err := i.ResolveInstructionAccounts(&tx.Message)
// 		if err != nil {
// 			panic(err)
// 		}

// 		// Feed the accounts and data to the system program parser
// 		// OR see below for alternative parsing when you DON'T know
// 		// what program the instruction is for / you don't have a parser.
// 		inst, err := DecodeInstruction(accounts, i.Data)
// 		if err != nil {
// 			panic(err)
// 		}

// 		// inst.Impl contains the specific instruction type (in this case, `inst.Impl` is a `*system.Transfer`)
// 		spew.Dump(inst)
// 		if _, ok := inst.Impl.(*CreateMarket); !ok {
// 			fmt.Println("the instruction is not a *openbook.CreateMarket")
// 		}
// 	}
// }
