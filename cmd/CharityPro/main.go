package main

import (
	util  "github.com/tendermint/basecoin/utilForNewProject"
	"fmt"
	//eyes "github.com/tendermint/merkleeyes/client"
	//"path"
	"github.com/tendermint/merkleeyes/app"
	abci "github.com/tendermint/abci/types"
	btypes "github.com/tendermint/basecoin/types"
	"net/http"
	//"github.com/tendermint/basecoin/cmd/basecoin/commands"

)

const EyesCacheSize = 10000
func main(){
	//username  := "liyibai"
	//password := "1234567890"
	//flag := "ed25519"
	//
	//info, err := util.NewKeys(username, password)
	//if err != nil{
	//	fmt.Println("util.NewKeys error")
	//	return
	//}
	//fmt.Println(info.Address)
	//fmt.Println("start  ")
	//util.QuerySequence("19D4B36BAAA7B203B301CB86F543EB2F49E34D39", "localhost", "46657")
	//GetSomethingNew([]byte("19D4B36BAAA7B203B301CB86F543EB2F49E34D39"))
	//result2, err :=util.QueryBalance("7716381C2D66A48ABA8C8A2B29E989C49AF2D68C", "localhost", "46657")
	//fmt.Println(result2)
	//result :=util.SendTx("ligang", "1234567890", "1mycoin", "7716381C2D66A48ABA8C8A2B29E989C49AF2D68C")
	//result2, err =util.QueryBalance("7716381C2D66A48ABA8C8A2B29E989C49AF2D68C", "localhost", "46657")
	//fmt.Println(result)
	//fmt.Println(err)
	//fmt.Println(result2)

	//http.HandleFunc("/sign", sign)
	//err := http.ListenAndServe("localhost:46600", nil)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Println("end")

}
func sign(w http.ResponseWriter, r *http.Request){
	fmt.Println("sign yes")
	query := r.URL.Query()

	name := query["name"][0]
	message := query["message"][0]
	password := query["password"][0]

	_, sig, _ := util.Sign(name, message, password)
	fmt.Println(sig)
	w.Write(sig)
}
func  GetSomethingNew(addr  []byte){
	//var eyesCli *localClient
	//eyesCli = eyes.NewLocalClient( "/home/ligang/.basecoin/data/merkleeyes.db", EyesCacheSize)
	//eyesCli.
	//v := eyesCli.Get([]byte("19D4B36BAAA7B203B301CB86F543EB2F49E34D39"))
	eyesApp := app.NewMerkleEyesApp("/home/ligang/.basecoin/data/merkleeyes.db" , EyesCacheSize)
	key := btypes.AccountKey(addr)
	var queryResq = abci.RequestQuery{
		Data:key,
		Path:"/key",
		Prove:true,
	}
	queryRes := eyesApp.Query(queryResq)
	fmt.Println(string(queryRes.Value))
}
