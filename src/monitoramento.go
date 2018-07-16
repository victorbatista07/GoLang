package main

import (
"fmt";
"os";
"net/http"
"time"
)
func main() {

    exibeIntroducao()
    for {
        exibeMenu()
        comando := leComando()
        executaComando(comando)
        time.Sleep (5 * time.Minute)
    }
}

func executaComando (comandoRecebido int) {
        switch comandoRecebido {
        case 1:
            iniciarMonitoramento()
        case 2:
            fmt.Println("Exibindo Logs...")
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
    site := "https://www.google.com.br"
    resp, _ := http.Get(site)
    fmt.Println(resp)

	if resp.StatusCode == 200 {
        fmt.Println("Site carregado com sucesso")
    } else {
	    fmt.Println("Eita, o site trashou")
    }
}
