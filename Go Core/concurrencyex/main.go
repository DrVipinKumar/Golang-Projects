//google.com facebook.com twitter.com github.com quora.com youtube.com yahoo.com
package main
import (
	"log"
	"fmt"
	"net/http"
	"os"
	"sync"
)
var wg sync.WaitGroup
func checkURLConnection(url string){
	defer wg.Done()
	res,err:=http.Get(url)
	if err!=nil {
		panic(err)
	}
     fmt.Printf("Status:[%d], URL:%s\n",res.StatusCode,url)
}
func main(){
	
if len(os.Args)<2 {
  log.Fatalln("Use like: go run main url1 url2 url3 .... urln")
}
for _,url:=range os.Args[1:]{
  wg.Add(1)
  go checkURLConnection("https://"+url)
}
wg.Wait()
}