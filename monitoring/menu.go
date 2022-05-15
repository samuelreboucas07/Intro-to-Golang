package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 3

func main() {
	for {
		exibeMenu()
		comando := lerComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando.")
			os.Exit(-1)

		}
	}
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando) // or fmt.Scanf("%d", &comando)
	return comando
}

func iniciarMonitoramento() {
	sites := lerSitesArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, elem := range sites { //for i := 0; i < len(sites); i++ {
			testaSite(elem)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)

	}
}

func exibeNomes() {
	nomes := []string{"Dougd las", "daniel", "bernado"} //slice, append, len
	fmt.Println(nomes)
}

func testaSite(site string) {
	response, _ := http.Get(site)
	if response.StatusCode == 200 {
		fmt.Println("Site: ", site, "Carregado com sucesso.")
		escreverLogs(site, true)
	} else {
		fmt.Println("Site: ", site, "Está com problemas. Status:", response.StatusCode)
		escreverLogs(site, false)
	}

}

func lerSitesArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	// arquivo, err := ioutil.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("Ocoreu um erro:", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func escreverLogs(site string, status bool) {
	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocoreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("logs.txt")
	if err != nil {
		fmt.Println("Ocoreu um erro:", err)
	}
	fmt.Println(string(arquivo))
}
