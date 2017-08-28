package main

import (
    util "github.com/tendermint/basecoin/utilForNewProject"
	"fmt"
)
func main(){
	fmt.Println("main")
	util.Start()
	//name := "liming"
	//password := "1234567890"
	//message := "iLOVEu"
	//pubkey, sig:= util.Sign(name, message, password)
	//fmt.Println(len(sig))
	//fmt.Println(len(pubkey))
	////pubkey = append(pubkey[:31], '1')
	//fmt.Println(util.Verify(pubkey, sig, []byte(message)))

}