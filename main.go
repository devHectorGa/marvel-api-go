
package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main(){
	menu :=
`
Bienvenido, ¿Qué prefieres?
[ 1 ] Buscar por nombre
[ 2 ] Listar	
`
	fmt.Print(menu)

	reader := bufio.NewReader(os.Stdin)
	entrada, _ := reader.ReadString('\n')
	eleccion := strings.TrimRight(entrada, "\n\r")


	public := "7e6d2e682ce9a165d3d05beb8295f4d4"
	private := "fc8afae2a11baa56780068c16be44d6f50b0c7ea"
	ts := time.Now().Format("2006-01-02T15:04:05.999999-07:00")
	var hash = md5.Sum([]byte(ts + private + public))
	auth := "ts="+ts+"&apikey="+public+"&hash="+hex.EncodeToString(hash[:])
	switch eleccion {
	case "1":
		buscarPorNombre(auth)
	case "2":
		listar(auth)
	default:
		main()
	}

}
func buscarPorNombre(auth string){
	fmt.Println("¿Que Heroe deseas buscar?")
	reader := bufio.NewReader(os.Stdin)
	entrada, _ := reader.ReadString('\n')
	eleccion := strings.TrimRight(entrada, "\n\r")
	hero := strings.Replace(eleccion, " ", "%20", -1)
	resp,err := http.Get("https://gateway.marvel.com/v1/public/characters?"+ "name="+ hero + "&" + auth)
	if err != nil {
		fmt.Printf(`Error: %s`, err)
	}else {
		data, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("%s", data)
	}

}
func listar(auth string){
	resp,err := http.Get("https://gateway.marvel.com/v1/public/characters?" + auth)
	if err != nil {
		fmt.Printf(`Error: %s`, err)
	}
	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Printf("%s", data)
}
/*
[ base url: https://gateway.marvel.com , api version: Cable ]

BUSCAR POR NOMBRE
/v1/public/characters/{characterId}

LISTAR
/v1/public/characters
*/
