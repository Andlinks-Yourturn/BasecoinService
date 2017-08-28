package utilForNewProject

import (
	"github.com/tendermint/go-crypto/cmd"
	"fmt"
	"github.com/tendermint/go-crypto/keys"
	"github.com/pkg/errors"
	"time"
	"net/http"
 	gojson "github.com/bitly/go-simplejson"
	keycmd "github.com/tendermint/go-crypto/cmd"
	btypes "github.com/tendermint/basecoin/types"
	sendtx "github.com/tendermint/basecoin/cmd/basecli/commands"
	txcmd "github.com/tendermint/light-client/commands/txs"

	rpcclient "github.com/tendermint/tendermint/rpc/client"
	"encoding/hex"
)

//save the address which can withdraw money from the system
var  addressNeedToValidate = map[string]string{
	"19D4B36BAAA7B203B301CB86F543EB2F49E34D39":"liming",
	 "B23D7891E79656A1A9576BF910028A8749E61C7D":"lihao",
}
var chain_Id = "test_chain_id"
type  QuerySequenceStr struct{
	Address []byte
}
type Json struct {
	data interface{}
}

func NewKeys(userName string, password string) (keys.Info,  error){

	var algo string = "ed25519"
	name := userName
	pass:= password

	info, seed, err := cmd.GetKeyManager().Create(name, pass, algo)
	if err == nil {
		fmt.Println("successful create keys, address is")
		fmt.Println(info.PubKey.Bytes())
		fmt.Println(info.Address.String())
		fmt.Println(seed)
	}
	return info, err
}

func AddAddressToWithDraw(address string, name string) (error ,bool){
	_, e := addressNeedToValidate[address]
	if e {
		//already exist
		return errors.New("address already exist"), false
	}else {
		addressNeedToValidate[address] = name
		return nil, true
	}

}
func IsExitInMap(address string) bool {
	fmt.Println(address)
	_, e := addressNeedToValidate[address]
	fmt.Println(e)
	return e
}
func CheckAuthority(message *chan string)bool{
	fmt.Println("yes you can withdraw the money")
	fmt.Println("sleep  for 4000")
	time.Sleep(12900*time.Millisecond)
	fmt.Println("sleep ok")
	*message <- "yes"
	return true
}
func ReturnSequence(sequence int) int{
	return sequence
}
func Query(address string , ip string, port string, path string)(int, error){
	var url string = "http://"+ip+":"+port+"/abci_query?path=\""+path+"\"&data=\""+address+"\"&prove=true"
	fmt.Println(url)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	http_client := &http.Client{}
	response, err := http_client.Do(request)
	defer  response.Body.Close()
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	jsonInstance, err := gojson.NewFromReader(response.Body)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	index := jsonInstance.Get("result").Get("response").Get("index").MustInt()
	fmt.Println(index)
	return  index, nil
}

func QueryBalance(address string, ip string, port string) (int, error){
	balance,err :=Query(address, ip, port, "balance")
	if err != nil {
		fmt.Println(err)
		return -1,err
	}
	return balance, nil
}
func QuerySequence(address string, ip string, port string) (int, error){
	sequence, err := Query(address, ip, port, "sequence")
	if err != nil {
		fmt.Println(err)
		return -1,err
	}
	return sequence, nil
}

//@Param userToAddress is info.address
func SendTx(userFrom string, password string, money string,userToAddress string)(error){
	fmt.Println("helslo")
	//construct a sendtx construct


	tx := new(btypes.SendTx)
	info, err := QueryKeyInfo(userFrom)
	if err != nil {
		return err
	}
	input_address := info.Address
//	input_pubKey := info.PubKey
	sequence, err := QuerySequence(info.Address.String(), "localhost", "46657")
	if err != nil {
		return err
	}
	amountCoins, err := btypes.ParseCoins(money)
	if err != nil {
		return err
	}
	output_Address, err := hex.DecodeString(userToAddress)
	fmt.Println(string(output_Address))



	tx.Inputs = []btypes.TxInput{{
		Coins:    amountCoins,
		Sequence:  sequence+1,
		Address: input_address,
		//PubKey: input_pubKey,
	}}
	tx.Outputs = []btypes.TxOutput{{
		Address: output_Address,
		Coins:   amountCoins,
	}}


	wrapTx := new(sendtx.SendTx)
	wrapTx.Tx = tx
	wrapTx.SetChianId(chain_Id)

	fmt.Println("wrapTx")
	fmt.Println(wrapTx)
	err = wrapTx.ValidateBasic()
	if err != nil {
		return err
	}
//  sign the transaction data
	manager := keycmd.GetKeyManager()
	manager.Sign(userFrom, password, wrapTx)
	packetBytes, err  := wrapTx.TxBytes()
	fmt.Println(packetBytes)
	if err != nil {
		return err
	}
	node := rpcclient.NewHTTP("tcp://localhost:46657", "/websocket")
	broadTx, err := node.BroadcastTxCommit(packetBytes)
	if err != nil {
		return err
	}
	fmt.Println(broadTx)
	result := txcmd.OutputTx(broadTx)
	fmt.Println(result)
	return result
}
func QueryKeyInfo(name string)(keys.Info,  error){
	info := keys.Info{}
	manager := keycmd.GetKeyManager()
	info, err := manager.Get(name)
	if err != nil {
		return info, err
	}
	return info, nil
}

