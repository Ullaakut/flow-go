package flow

import (
	"fmt"

	ledgerHash "github.com/onflow/flow-go/ledger/common/hash"
)

type RegisterID struct {
	Owner      string
	Controller string
	Key        string
}

func (r *RegisterID) String() string {
	return fmt.Sprintf("%x/%x/%x", r.Owner, r.Controller, r.Key)
}

func NewRegisterID(owner, controller, key string) RegisterID {
	return RegisterID{
		Owner:      owner,
		Controller: controller,
		Key:        key,
	}
}

// RegisterValue (value part of Register)
type RegisterValue = []byte

type RegisterEntry struct {
	Key   RegisterID
	Value RegisterValue
}

//handy container for sorting
type RegisterEntries []RegisterEntry

func (d RegisterEntries) Len() int {
	return len(d)
}

func (d RegisterEntries) Less(i, j int) bool {
	if d[i].Key.Owner != d[j].Key.Owner {
		return d[i].Key.Owner < d[j].Key.Owner
	} else if d[i].Key.Controller != d[j].Key.Controller {
		return d[i].Key.Controller < d[j].Key.Controller
	}
	return d[i].Key.Key < d[j].Key.Key
}

func (d RegisterEntries) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d RegisterEntries) IDs() []RegisterID {
	r := make([]RegisterID, len(d))
	for i, entry := range d {
		r[i] = entry.Key
	}
	return r
}

func (d RegisterEntries) Values() []RegisterValue {
	r := make([]RegisterValue, len(d))
	for i, entry := range d {
		r[i] = entry.Value
	}
	return r
}

// StorageProof (proof of a read or update to the state, Merkle path of some sort)
type StorageProof = []byte

// StateCommitment holds the root hash of the tree (Snapshot)
// TODO: solve the circular dependency and define StateCommitment as ledger.State
type StateCommitment ledgerHash.Hash

// DummyStateCommitment is an arbitrary value used in function failure cases,
// although it can represent a valid state commitment.
var DummyStateCommitment = StateCommitment(ledgerHash.DummyHash)

// ToStateCommitment converts a byte slice into a StateComitment.
// It returns an error if the slice has an invalid length.
func ToStateCommitment(stateBytes []byte) (StateCommitment, error) {
	var state StateCommitment
	if len(stateBytes) != len(state) {
		return DummyStateCommitment, fmt.Errorf("expecting %d bytes but got %d bytes", len(state), len(stateBytes))
	}
	copy(state[:], stateBytes)
	return state, nil
}
