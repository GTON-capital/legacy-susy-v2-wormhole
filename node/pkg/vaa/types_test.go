package vaa

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)


func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handleTestErr(t *testing.T, err error) {
	handleErr(err)
}
func TestSerializeDeserialize(t *testing.T) {
	// module := make([]byte, 32)
	// copy(module, []byte("SuSyBridge"))

	// emitterAddress := common.HexToAddress("0xb6Ecfe5CCB35E9a99aD1538b3748Da36AcC2f8dc")
	// // emmiterAddress := make([]byte, 32)
	// // copy(emmiterAddress, []byte("SuSyBridge"))

	// pk := "CiAtklXWC6R7gkbLtkMYKmnHU6t0Qv6pr+E4aD80fGJX9Q===XEjB"
	// keyBytes, err := base64.StdEncoding.DecodeString(pk)
	// handleErr(err)

	// // keyBytes := common.Hex2Bytes(key)
	// key, err := crypto.ToECDSA(keyBytes)
	// handleErr(err)
	// // if err != nil {
	// // 	return nil, common.Address{}, err
	// // }

	// payload := BridgeStructs_RegisterChain {
	// 	// module: make([]byte, 0),
	// 	module: module,
	// 	action: uint8(1),
	// 	chainId: 1,
	// 	emitterChainID: 4,
	// 	emitterAddress: emitterAddress.Bytes(),
	// }

	// // bytesPayload, _ := SerializeData(payload)
	// // types.SignTx(transaction, types.NewEIP155Signer(cID), key)

	// tests := []struct {
	// 	name string
	// 	vaa  *VAA
	// }{
	// 	{
	// 		name: "NormalVAA",
	// 		vaa: &VAA{
	// 			Version:          1,
	// 			GuardianSetIndex: 9,
	// 			Signatures: []*Signature{
	// 				{
	// 					Index:     1,
	// 					Signature: [65]byte{},
	// 				},
	// 			},
	// 			Timestamp:        time.Unix(2837, 0),
	// 			Nonce:            10,
	// 			Sequence:         3,
	// 			ConsistencyLevel: 5,
	// 			EmitterChain:     8,
	// 			EmitterAddress:   Address{1, 2, 3},
	// 			// Payload:          []byte("abc"),
	// 			Payload:          bytesPayload,
	// 		},
	// 	},
	// }
	// for _, test := range tests {
	// 	t.Run(test.name, func(t *testing.T) {

	// 		vaaData, err := test.vaa.Marshal()
	// 		require.NoError(t, err)

	// 		println(hex.EncodeToString(vaaData))
	// 		vaaParsed, err := Unmarshal(vaaData)
	// 		require.NoError(t, err)

	// 		require.EqualValues(t, test.vaa, vaaParsed)
	// 	})
	// }
}


func TestVerifySignature(t *testing.T) {
	v := &VAA{
		Version:          8,
		GuardianSetIndex: 9,
		Timestamp:        time.Unix(2837, 0),
		Nonce:            5,
		Sequence:         10,
		ConsistencyLevel: 2,
		EmitterChain:     2,
		EmitterAddress:   Address{0, 1, 2, 3, 4},
		Payload:          []byte("abcd"),
	}

	data, err := v.SigningMsg()
	require.NoError(t, err)

	key, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	require.NoError(t, err)

	sig, err := crypto.Sign(data.Bytes(), key)
	require.NoError(t, err)
	sigData := [65]byte{}
	copy(sigData[:], sig)

	v.Signatures = append(v.Signatures, &Signature{
		Index:     0,
		Signature: sigData,
	})
	addr := crypto.PubkeyToAddress(key.PublicKey)
	require.True(t, v.VerifySignatures([]common.Address{
		addr,
	}))
}

func TestVerifySignature_RegisterChain(t *testing.T) {
	module := make([]byte, 32)
	copy(module, []byte("SuSyBridge"))

	emitterAddress := common.HexToAddress("0xb6Ecfe5CCB35E9a99aD1538b3748Da36AcC2f8dc")
	// emmiterAddress := make([]byte, 32)
	// copy(emmiterAddress, []byte("SuSyBridge"))

	pk := "CiAtklXWC6R7gkbLtkMYKmnHU6t0Qv6pr+E4aD80fGJX9Q=="
	keyBytes, err := base64.StdEncoding.DecodeString(pk)
	handleErr(err)

	encodedPK := hexutil.Encode(keyBytes)
	fmt.Println(encodedPK)

	keyBytes = keyBytes[2:]

	// keyBytes := common.Hex2Bytes(key)
	key, err := crypto.ToECDSA(keyBytes)
	handleErr(err)
	// if err != nil {
	// 	return nil, common.Address{}, err
	// }

	payload := BridgeStructs_RegisterChain {
		// module: make([]byte, 0),
		module: module,
		action: uint8(1),
		chainId: 1,
		emitterChainID: 4,
		emitterAddress: emitterAddress.Bytes(),
	}
	payloadBytes, err := SerializeData(payload)
	handleErr(err)


	var emitterAddressCasted Address

	copy(emitterAddressCasted[:], payload.emitterAddress)
	
	v := &VAA{
		Version:          8,
		GuardianSetIndex: 0,
		Timestamp:        time.Unix(2837, 0),
		Nonce:            5,
		Sequence:         10,
		ConsistencyLevel: 2,
		EmitterChain:     ChainID(payload.emitterChainID),
		EmitterAddress:   emitterAddressCasted,
		Payload:          payloadBytes,
	}

	data, err := v.SigningMsg()
	// require.NoError(t, err)

	// key, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	// require.NoError(t, err)

	sig, err := crypto.Sign(data.Bytes(), key)
	handleErr(err)

	sigData := [65]byte{}
	copy(sigData[:], sig)

	v.Signatures = append(v.Signatures, &Signature{
		Index:     0,
		Signature: sigData,
	})

	serializedVaa, err := v.Marshal()
	handleErr(err)

	fmt.Println(serializedVaa)
	fmt.Println(hexutil.Encode(serializedVaa))
}
