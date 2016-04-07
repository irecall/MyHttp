package MyHttp

import(
	"fmt"
	"net/http"
)

type myHttp struct{}

func (*myHttp) post(url string,param string,h map){
	return *myHttp.execu("POST",url,param,h))
}

func (*myHttp) get(url string,param string,h map){
	return *myHttp.execu("GET",url,param,h)
}


func (*myHttp) execu(method string,url string,param string,h map) {
	client := &http.Client{}
	req, err := http.NewRequest(method,url,string.NewReader(param))
	if err !=nil {

	}
	
	for k,v := range h{
		req.setCookie(k,v)
	}

	resp,err := client.Do(req)
	
	defer resp.Body.Close()
	
	body,err := ioutil.ReadAll(reqs.Body)
	if err != nil{

	}

	return fmt.PrintLn(string(body))	
}
