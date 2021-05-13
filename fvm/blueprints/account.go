package blueprints

import (
	"fmt"

	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"

	"github.com/onflow/flow-go/model/flow"
)

const initAccountTransactionTemplate = `
import FlowServiceAccount from 0x%s

transaction(restrictedAccountCreationEnabled: Bool) {
  prepare(newAccount: AuthAccount, payerAccount: AuthAccount) {
    if restrictedAccountCreationEnabled && !FlowServiceAccount.isAccountCreator(payerAccount.address) {
	  panic("Account not authorized to create accounts")
    }

    FlowServiceAccount.setupNewAccount(newAccount: newAccount, payer: payerAccount)
  }
}
`

func InitAccountTransaction(
	payerAddress flow.Address,
	accountAddress flow.Address,
	serviceAddress flow.Address,
	restrictedAccountCreationEnabled bool,
) *flow.TransactionBody {
	arg, err := jsoncdc.Encode(cadence.NewBool(restrictedAccountCreationEnabled))
	if err != nil {
		// this should not fail! It simply encodes a boolean
		panic(fmt.Errorf("cannot json encode cadence boolean argument: %w", err))
	}

	return flow.NewTransactionBody().
		SetScript([]byte(fmt.Sprintf(initAccountTransactionTemplate, serviceAddress))).
		AddAuthorizer(accountAddress).
		AddAuthorizer(payerAddress).
		AddArgument(arg)
}

const getFlowTokenBalanceScriptTemplate = `
import FlowServiceAccount from 0x%s

pub fun main(): UFix64 {
  let acct = getAccount(0x%s)
  return FlowServiceAccount.defaultTokenBalance(acct)
}
`

func GetFlowTokenBalanceScript(accountAddress, serviceAddress flow.Address) []byte {
	return []byte(fmt.Sprintf(getFlowTokenBalanceScriptTemplate, serviceAddress, accountAddress))
}

const getFlowTokenAvailableBalanceScriptTemplate = `
import FlowStorageFees from 0x%s

pub fun main(): UFix64 {
  return FlowStorageFees.defaultTokenAvailableBalance(0x%s)
}
`

func GetFlowTokenAvailableBalanceScript(accountAddress, serviceAddress flow.Address) []byte {
	return []byte(fmt.Sprintf(getFlowTokenAvailableBalanceScriptTemplate, serviceAddress, accountAddress))
}

const getStorageCapacityScriptTemplate = `
import FlowStorageFees from 0x%s

pub fun main(): UFix64 {
	return FlowStorageFees.calculateAccountCapacity(0x%s)
}
`

func GetStorageCapacityScript(accountAddress, serviceAddress flow.Address) []byte {
	return []byte(fmt.Sprintf(getStorageCapacityScriptTemplate, serviceAddress, accountAddress))
}
