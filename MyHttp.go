package MyHttp

import(
	"fmt"
	"net/http"
)

type myHttp struct{}

func New()*myHttp{
	return &myHttp{}
}

func (m *myHttp) post(url string,param string,h map){
	return m.execu("POST",url,param,h))
}

func (m *myHttp) get(url string,param string,h map){
	return m.execu("GET",url,param,h)
}


func (m *myHttp) execu(method string,url string,param string,h map) string {
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

	return string(body)	
}
