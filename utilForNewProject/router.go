package utilForNewProject



import (
	 "net/http"
	"fmt"
)
/*set router for basecoin to handle http request
	created by ligang 2017/8/25
*/
type Result struct {
	Data []byte `json:"data"`
}
func init() {
	http.HandleFunc("/sign", sign)
	http.HandleFunc("/verify", verify)
	http.HandleFunc("/register", register)
	http.HandleFunc("/sendTx", sendTx)
	http.HandleFunc("/test", test)
	http.ListenAndServe("localhost:46600", nil)
}

//  "http://localhost:46600/sign?name=ligang&password=1234567890&message=hello"
func sign(w http.ResponseWriter, r *http.Request){
	fmt.Println("sign yes")
	query := r.URL.Query()

	name := query["name"][0]
	message := query["message"][0]
	password := query["password"][0]


	_, sig := Sign(name, message, password)


	fmt.Println(sig)
 	w.Write(sig)
}

// "http://localhost:46600/verify?name=&signature=&message="
func verify(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query()
	name := query["name"][0]
	sig := query["signature"][0]
	message := query["message"][0]

	info, _ := QueryKeyInfo(name)
	pubkey := info.PubKey.Bytes()[1:]
	result := Verify(pubkey, []byte(sig), []byte(message))
	if result {
		w.Write([]byte("true"))
	}else {
		w.Write([]byte("false"))
	}
}

//"http://localhost:46600/register?name=&password="
func register(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query["name"][0]
	password := query["password"][0]

	info, err := NewKeys(name, password)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("error"))
	} else {
		w.Write(info.Address)
	}
}

//"http://localhost:46600/sendTx?userFrom=&password=&money=&userToAddress"
//the value of money should be like "1000mycoin"
func sendTx(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query()
	password := query["password"][0]
	userFrom := query["userFrom"][0]
	money := query["money"][0]
	userToAddress := query["userToAddress"][0]

	err := SendTx(userFrom, password, money, userToAddress)
	if err != nil {
		w.Write([]byte("false"))
	}else {
		w.Write([]byte("true"))
	}
}

func test(w http.ResponseWriter, r *http.Request){
	fmt.Println("hello")
	w.Write([]byte("hello"))
}

func Start(){
	fmt.Println("hello")
}
