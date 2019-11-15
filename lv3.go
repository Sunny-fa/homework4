package main
import(
	"fmt"
	"net/http"
	"io/ioutil"
	)
func main(){
	response,err := http.Get("http://www.baidu.com")
	if(err!=nil){
		fmt.Println(err)
	}
	defer response.Body.Close()
	body,err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
