package vaa

import (
	"encoding/hex"
	"fmt"
	"os"
	"testing"
)




func TestRegisterChainPayload_Ser(t *testing.T) {

	module := make([]byte, 32)
	emmiterAddress := make([]byte, 32)

	payload := BridgeStructs_RegisterChain {
		// module: make([]byte, 0),
		module: module,
		action: 1,
		chainId: 4,
		emitterChainID: 1,
		emitterAddress: emmiterAddress,
	}

	bytesPayload, err := SerializeData(payload)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	hexStr := hex.EncodeToString(bytesPayload)
	fmt.Println("hex is here: ")
	fmt.Println(hexStr)
}

