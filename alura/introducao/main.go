package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 3
const delay = 1

func main() {
	exibirMenu()
	//registrarLogs("site-falso.com.br", false)
	comando := lerComando()

	if comando == 1 {
		monitorandoPaginas()
	} else if comando == 2 {
		exibindoLogs()
	} else if comando == 3 {
		saindoDoPrograma()
	} else {
		fmt.Println("comando invalido")
		os.Exit(-1)
	}
	fmt.Println("<================>")
}

func exibirMenu() {
	nome := "Elton"
	fmt.Println("Olá, Sr.", nome)
	versao := 1.1
	fmt.Println("A versão do programa", versao)
	fmt.Println("<================>")
}

func lerComando() int {
	var comando int
	fmt.Println("digite qual comando deseja:")
	fmt.Println("1 - monitoramento dos sites")
	fmt.Println("2 - exibição dos logs")
	fmt.Println("3 - saida do programa")
	fmt.Println("<================>")
	fmt.Scan(&comando)
	fmt.Println("Você escolheu o comando:", comando)
	return comando
}

func exibindoLogs() {
	fmt.Println("Exibindo logs...")
}

func saindoDoPrograma() {
	fmt.Println("Saindo do programa")
	os.Exit(0)
}

func monitorandoPaginas() {
	sites := lersitesDeArquivos()

	for i := 0; i < monitoramento; i++ {

		for i, site := range sites {
			fmt.Println("testando site numero:", i, "| nome:", site)
			testeSite(site)
		}

		time.Sleep(delay * time.Second)	
		fmt.Println("")
	}
}

func testeSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("ocorreu um erro no teste do site", err)
	}
	
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println("site,", site, "carregando com sucesso!")
	} else {
		fmt.Println("O site,", site, "está fora do ar")
	}
}

func lersitesDeArquivos() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("ocorreu um erro na leitura do arquivo", err)
	}
	
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
		
	return sites
}

func registrarLogs(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("ocorreu um erro", err)
	}
	fmt.Println(arquivo)

	arquivo.WriteString(site + " - online: " + strconv.FormatBool(status))	

	arquivo.Close()
}
