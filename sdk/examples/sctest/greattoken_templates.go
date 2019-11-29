package sctest

import (
	"fmt"

	"github.com/dapperlabs/flow-go/model/flow"
)

// GenerateCreateMinterScript Creates a script that instantiates a new GreatNFTMinter instance
// and stores it in memory.
// Initial ID and special mod are arguments to the GreatNFTMinter constructor.
// The GreatNFTMinter must have been deployed already.
func GenerateCreateMinterScript(nftAddr flow.Address, initialID, specialMod int) []byte {
	template := `
		import 0x%s

		transaction {
		  prepare(acct: Account) {
			let existing <- acct.storage[GreatNFTMinter] <- createGreatNFTMinter(firstID: %d, specialMod: %d)
			if existing != nil {
				panic("existed")
			}
			destroy existing
			acct.storage[&GreatNFTMinter] = &acct.storage[GreatNFTMinter] as GreatNFTMinter
		  }
		  execute {}
		}
	`

	return []byte(fmt.Sprintf(template, nftAddr, initialID, specialMod))
}

// GenerateMintScript Creates a script that mints an NFT and put it into storage.
// The minter must have been instantiated already.
func GenerateMintScript(nftCodeAddr flow.Address) []byte {
	template := `
		import GreatNFTMinter, GreatNFT from 0x%s

		transaction {
		  prepare(acct: Account) {
			let minter = acct.storage[&GreatNFTMinter] ?? panic("missing minter")
			let existing <- acct.storage[GreatNFT] <- minter.mint()
			destroy existing
		  }
		  execute {}
		}
	`

	return []byte(fmt.Sprintf(template, nftCodeAddr.String()))
}

// GenerateInspectNFTScript Creates a script that retrieves an NFT from storage and makes assertions
// about its properties. If these assertions fail, the script panics.
func GenerateInspectNFTScript(nftCodeAddr, userAddr flow.Address, expectedID int, expectedIsSpecial bool) []byte {
	template := `
		import GreatNFT from 0x%s

		pub fun main() {
		  let acct = getAccount(0x%s)
		  let nft <- acct.storage[GreatNFT] ?? panic("missing nft")
		  if nft.id() != %d {
			panic("incorrect id")
		  }
		  if nft.isSpecial() != %t {
			panic("incorrect specialness")
		  }
		  destroy nft
		}
	`

	return []byte(fmt.Sprintf(template, nftCodeAddr, userAddr, expectedID, expectedIsSpecial))
}
