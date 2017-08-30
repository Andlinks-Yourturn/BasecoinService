package utilForNewProject

import (
	"os"
	"github.com/tendermint/go-crypto"
	"github.com/tendermint/go-crypto/keys/cryptostore"
	"fmt"
	"path/filepath"
	"github.com/tendermint/go-crypto/keys"
	"github.com/tendermint/go-crypto/keys/storage/filestorage"
	ed25519 "golang.org/x/crypto/ed25519"
)
const KeySubdir = "keys"
var signEssue  SignatureEssue

type SignatureEssue struct {
	Publikey crypto.PubKey
	Sig  crypto.Signature
	Message [] byte
}

func (sign *SignatureEssue) SignBytes() []byte{
	// SignBytes is the immutable data, which needs to be signed
	return sign.Message
}
func (sign *SignatureEssue) Sign(pubkey crypto.PubKey, sig crypto.Signature) error{
	//put the pubkey into SignatureEssue
	//sign.Publikeys = append(sign.Publikeys, pubkey)
	//sign.Sigs = append(sign.Sigs, sig)
	sign.Publikey = pubkey
	sign.Sig = sig
	return nil
}
func (sign *SignatureEssue) Signers() ([]crypto.PubKey, error){
	// Signers will return the public key(s) that signed if the signature
	// is valid, or an error if there is any issue with the signature,
	return []crypto.PubKey{sign.Publikey},  nil
}
func (sign * SignatureEssue) TxBytes() ([]byte, error){
	// TxBytes returns the transaction data as well as all signatures
	fmt.Println("txByte")
	return nil, nil
}


//return publickey & signature
//later using these return value to verify
func Sign(name string, message string, password string)([]byte, []byte, error){

	var manager cryptostore.Manager
	rootDir := os.Getenv("BASECLIHOME")
	fmt.Println("rootdir"+rootDir)
	keyDir := filepath.Join(rootDir, KeySubdir)

	signEssue.Message = []byte(message)
	// TODO: smarter loading??? with language and fallback?
	codec := keys.MustLoadCodec("english")
	manager = cryptostore.New(
		cryptostore.SecretBox,
		filestorage.New(keyDir),
		codec,
	)
	var sign keys.Signable
	fmt.Println(&sign)
	sign = &signEssue
	fmt.Println(sign)
	err := manager.Sign(name, password, sign)
	if err != nil{
		return 	nil, nil, err
	}else {
		return signEssue.Publikey.Bytes()[1:], signEssue.Sig.Bytes()[1:], nil
	}
}

func Verify(pubkey []byte , sig  []byte, messsage []byte) bool {
	//pubKey []byte, message , sig []byte

	//if len(signEssue.Sigs) == 0 || len(signEssue.Publikeys)== 0 || len(signEssue.Message) == 0{
	//	return false
	//}
	//fmt.Println("signEssue.Publickeys[0].Bytes")
	//fmt.Println(len(signEssue.Publikeys[0].Bytes()))
	//fmt.Println(string(signEssue.Publikeys[0].Bytes()))
	return ed25519.Verify(pubkey, messsage, sig)
}