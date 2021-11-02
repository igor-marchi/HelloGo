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

const monitoramentos = 5
const delay = 2 //seconds

func main() {
	printIntro()
	sites := readFile()
	fmt.Println(sites)

	for {
		printMenu()
		readCommand := scanAndPrintCommand()

		switch readCommand {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Vazei..")
			os.Exit(0)
		default:
			fmt.Println("Desconheço..")
			os.Exit(-1)
		}

	}

}

func printIntro() {
	var name string
	version := 1.1

	fmt.Println("Qual é seu nome?")
	fmt.Scan(&name)
	fmt.Println("Olá", name)
	fmt.Println("Este program está na versão", version)
}

func printMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair")
}

func scanAndPrintCommand() int {
	var readCommand int
	fmt.Scan(&readCommand)
	fmt.Println("O comando escolhido foi", readCommand)

	return readCommand
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	sites := readFile()
	fmt.Println(sites)
	for i := 0; i < monitoramentos; i++ {
		fmt.Println("Teste", i+1)
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n\nsite:", site)

	if resp.StatusCode == 200 {
		fmt.Println("Positivo e operante", resp.Status)
		fmt.Println("\t status", resp.Status)
		registerLog(site, true)
	} else {
		fmt.Println("Site apresentando indisponibilidade", resp.Status)
		fmt.Println("\t status", resp.Status)
		registerLog(site, false)
	}
}

func readFile() []string {
	var sites []string

	// arquivo, err := ioutil.ReadFile("sites.txt")
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error:", err)
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

func registerLog(site string, status bool) {
	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
