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

const monitoramento = 2
const delay = 5

func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()
		executaComando(comando)
	}
}

func executaComando(comandoRecebido int) {
	switch comandoRecebido {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")
		imprimeLogs()
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	fmt.Println("Digite seu nome")
	var nome string
	fmt.Scan(&nome)
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando monitoramento...")
	sites := lerArquivo()

	for i := 0; i < monitoramento; i++ {
		for i, cadaSite := range sites {
			fmt.Println("Testando site", i, cadaSite)
			testandoSite(cadaSite)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testandoSite(site string) {
	resp, err := http.Get(site)
	fmt.Println(resp)
	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "carregado com sucesso")
		registraLogs(site, true)
	} else {
		fmt.Println("Eita, o site trashou", err)
		registraLogs(site, false)
	}
}

func lerArquivo() []string {
	var sites []string
	arquivo, err := os.Open("lib/sites.txt")

	if err != nil {
		fmt.Println(err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		if err == io.EOF {
			break
		}
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
	}

	arquivo.Close()
	return sites
}

func imprimeLogs() {
	arquivos, err := ioutil.ReadFile("lib/logs.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivos))
}

func registraLogs(site string, status bool) {
	arquivo, err := os.OpenFile("lib/logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}
