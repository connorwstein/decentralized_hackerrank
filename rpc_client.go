package main
// 
// import (
//     "context"
// 	"crypto/ecdsa"
//     "fmt"
//     "log"
//     "math/big"
// 
//     "github.com/ethereum/go-ethereum/common"
//     "github.com/ethereum/go-ethereum/core/types"
//     "github.com/ethereum/go-ethereum/common/hexutil"
//     "github.com/ethereum/go-ethereum/ethclient"
//     "github.com/ethereum/go-ethereum/crypto"
// //     "github.com/ethereum/go-ethereum/crypto/sha3"
// )
// 

// func main() {
// 	// Connect to ganache-cli 
//     conn, err := ethclient.Dial("http://127.0.0.1:8545")
//     if err != nil {
//         log.Fatal("Whoops something went wrong!", err)
//     }
// 
//     ctx := context.Background()
// 	block, err := conn.BlockByNumber(ctx, nil)
// 	if err != nil {
// 		fmt.Println("Error getting block")
// 	}
// 	fmt.Println(block)
	// Get all the accounts
// func (ec *Client) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
// 	res, err := conn.BalanceAt(ctx,  common.HexToAddress("0x209f97047869bbd1a5d6bed03a064e6607c89222"), block.Number())
// 	fmt.Println(res)

// 	privateKey, err := crypto.GenerateKey()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	privateKeyBytes := crypto.FromECDSA(privateKey)
// 	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19
// 
// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		log.Fatal("error casting public key to ECDSA")
// 	}
// 	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
// 	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934f
// 
// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
// 	fmt.Println(fromAddress.Hex()) // 0x96216849c49358B10257cb55b28eA603c874b05E
// 	hash := sha3.NewKeccak256()
// 	hash.Write(publicKeyBytes[1:])
// 	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
	
	// Send some eth to from our generated wallet to 0x87abebfcea3c37cf61a2ca6610194613d84da223
// 	nonce, err := conn.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}	
// 	fmt.Printf("Pending nonce for %v is %v", fromAddress, nonce)
// 
// 	
// 	value := big.NewInt(1000000000000000000) // in wei (1 eth)
// 	gasLimit := uint64(21000) // in units
// 	gasPrice := big.NewInt(30000000000) // in wei (30 gwei)
// 	toAddress := common.HexToAddress("0x87abebfcea3c37cf61a2ca6610194613d84da223")
// 	var data []byte
// 	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
// 	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = conn.SendTransaction(context.Background(), signedTx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	

// 	toAddress := common.HexToAddress("0x87abebfcea3c37cf61a2ca6610194613d84da223")

// }
