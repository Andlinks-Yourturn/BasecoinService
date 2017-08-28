package utilForNewProject

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"net/http"
	"strconv"
	"os"
	"os/exec"
	"encoding/json"
)




type transaction struct {
	Input    string
	Output     string
	Amount   int
}


func getHeight(url string) int{

	res, err := http.Get(url)
	js, err := simplejson.NewFromReader(res.Body) //反序列化
	if err != nil {
		panic(err.Error())
	}

	height := js.Get("result").Get("response").Get("last_block_height").MustInt()

	return height
}

func getDataHash(num int) string{

	url := "http://localhost:46657/block?height="
	u := url +strconv.Itoa(num)
	//fmt.Println("----",u)
	res, err := http.Get(u)
	js, err := simplejson.NewFromReader(res.Body) //反序列化

	if err != nil {
		panic(err.Error())
	}

	hash := js.Get("result").Get("block_meta").Get("header").Get("data_hash").MustString()

	return hash

}

func getRelatedEntity(transhash string,target string) []byte{

	var (
		cmdOut []byte
		err    error
		inputAddr string
		outputAddr string
		amount int
		b []byte
	)
	cmd := "basecli"

	args := []string{"query","tx",transhash}

	if cmdOut, err = exec.Command(cmd, args...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
		os.Exit(1)
	}
	js, err :=simplejson.NewJson(cmdOut)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(js)
	fmt.Println("---",js.Get("height").MustInt())

	inputArr, err := js.Get("data").Get("data").Get("inputs").Array()
	outputArr, err := js.Get("data").Get("data").Get("outputs").Array()
	for i, _ := range outputArr {
		input :=	js.Get("data").Get("data").Get("outputs").GetIndex(i)
		add:= input.Get("address").MustString()

		//fmt.Println("address= ", outputAddr)
		if add == target{
			fmt.Println("match~~~~~~~~~")
			outputAddr = add
			fmt.Println(len(inputArr))
			for i, _ := range inputArr {
				input :=	js.Get("data").Get("data").Get("inputs").GetIndex(i)
				add:= input.Get("address").MustString()
				// age := person.Get("age").MustInt()
				// email := person.Get("email").MustString()
				coinsArr,_ := input.Get("coins").Array()
				// fmt.Println("coins ",len(coinsArr))
				for i, _ := range coinsArr {
					coin :=input.Get("coins").GetIndex(i)
					amount = coin.Get("amount").MustInt()
					fmt.Println("amount: ",amount)
				}
				inputAddr = add
			}

			//创建返回对象

			tx := transaction {
				Input: inputAddr,
				Output:outputAddr,
				Amount: amount,
			}

			b, err = json.Marshal(tx)
			if err != nil {
				fmt.Println("encoding faild")
			} else {
				fmt.Println("encoded data success ")
			}
		}
	}

	return b


}

func (s * transaction) ShowTransaction() {
	fmt.Println("show Transaction :")
	fmt.Println("\tInput\t:", s.Input)
	fmt.Println("\tOutput\t:", s.Output)
	fmt.Println("\tAmount\t:", s.Amount)
}


func SearchForRelate(address string){

	url := "http://localhost:46657/abci_info"
	height := getHeight(url)


	var transArray []string


	goroutines := (height/3000)
	fmt.Println(goroutines)
	for i:=1;  i<=goroutines; i++  {
		go func(index int){
			for j:=0; j< index*3000; j++{
				h:= getDataHash(j)
				result := getRelatedEntity(h,address)
				transArray = append(transArray, string(result))
			}
		}(i)
	}

	fmt.Println("len: ",len(transArray))
}