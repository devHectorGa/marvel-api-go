
package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)
type searchByName struct {
	Code int `json:"code"`
	Data	struct{
		Results	[]result `json:"results"`
	}	`json:"data"`
}
type result struct {
	Name	string `json:"name"`
	Description	string	`json:"description"`
	Comics	struct{ Items []items `json:"items"` } `json:"comics"`
	Series struct{ Items []items `json:"items"`  } `json:"series"`
	Stories struct { Items []items `json:"items"` } `json:"stories"`
	Events struct{ Items []items `json:"items"` } `json:"events"`

}
type items struct{
	NameItem	string	`json:"name"`
}
func main(){
	menu :=
`
	Bienvenido, ¿Qué prefieres?
	[ 1 ] Buscar por nombre
	[ 2 ] Listar
	[ 3 ] Salir
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
	case "3":
		break
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

		var record searchByName

		json.Unmarshal(data, &record)
		printInDisplay(record)
	}

}
func listar(auth string){
	resp,err := http.Get("https://gateway.marvel.com/v1/public/characters?orderBy=name&limit=20&" + auth)
	if err != nil {
		fmt.Printf(`Error: %s`, err)
	}else{
		data, _ := ioutil.ReadAll(resp.Body)

		var record searchByName

		json.Unmarshal(data, &record)
		printInDisplay(record)
	}
}

func printInDisplay(record searchByName){
	var comicSearch[]items = record.Data.Results[0].Comics.Items
	var seriesSearch[]items = record.Data.Results[0].Series.Items
	var storiesSearch[]items = record.Data.Results[0].Stories.Items
	var eventsSearch[]items = record.Data.Results[0].Events.Items
	if(len(record.Data.Results) == 1){
		fmt.Println("El nombre del Heroe es: " +record.Data.Results[0].Name)
		fmt.Println("Su descripcion: " +record.Data.Results[0].Description)
		if(len(comicSearch) != 0){
			println("Comics: ")
			for i := 0; i < len(comicSearch); i++ {
				println(" * " + comicSearch[i].NameItem)
			}
		}
		if(len(seriesSearch) != 0){
			println("Series: ")
			for i := 0; i < len(seriesSearch); i++ {
				println(" * " + seriesSearch[i].NameItem)
			}
		}
		if(len(storiesSearch) != 0){
			println("Stories: ")
			for i := 0; i < len(storiesSearch); i++ {
				println(" * " + storiesSearch[i].NameItem)
			}
		}
		if(len(eventsSearch) != 0){
			println("Events: ")
			for i := 0; i < len(eventsSearch); i++ {
				println(" * " + eventsSearch[i].NameItem)
			}
		}
	}else{
		fmt.Println("Lista de Heroes:")

		for i := 0; i < len(record.Data.Results); i++{
			fmt.Println(" * " + record.Data.Results[i].Name)
		}
	}
	main()
}
