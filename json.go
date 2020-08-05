package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//Server armazena dados dos servers
type Server struct {
	ServerName string
	ServerIP   string
}

// Serverslice armazena um slice dos itens de Server
type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	//instanciando dados p/criar um json
	s.Servers = append(s.Servers, Server{ServerName: "Shangai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	//criando um json com dados instanciados de s.servers
	sjson, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(sjson))
	fmt.Println(reflect.TypeOf(sjson))

	//outra maneira de se tralhar com json
	b := []byte(`{"Name":"Wedbesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	//json de b e parseando para f (uma interface vazia)
	err = json.Unmarshal(b, &f)
	fmt.Println(f)
	if err != nil {
		fmt.Println(err)
	}

	// manipulando json
	m := f.(map[string]interface{}) //map do tipo interface com chaves do tipo string
	for chave, valor := range m {
		switch vv := valor.(type) {
		case string:
			fmt.Println(chave, "é do tipo string", vv)
		case int:
			fmt.Println(chave, "é do tipo int", vv)
		case float64:
			fmt.Println(chave, "é do tipo float64", vv)
		case []interface{}:
			fmt.Println(chave, "é um array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(chave, "É um tipo que não consigo manipular")
		}
	}
	fmt.Println(reflect.TypeOf(m))
	fmt.Println(reflect.TypeOf(f))
}
