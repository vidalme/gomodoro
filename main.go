package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const TEMPO_TRABALHO = 1
const TEMPO_REPOUSO = 1

type intervalo struct {
	t         time.Time
	d         time.Duration
	mensagens [2]string
}

func menuMensagem() {
	fmt.Println("+++++++  GOmodoro menu  +++++++")
	fmt.Println("1 - Inicio rapido ( 45 min trabalho / 15 min )")
	fmt.Println("2 - Configurar timer")
	fmt.Println("3 - Sair")
	fmt.Println("Digite uma opção: ")
}

func menuInicial() {

	menuMensagem()
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		fmt.Println("inicio rapido")
		inicioRapido()
	case "2":
		fmt.Println("Configurar timer")
	case "3":
		fmt.Println("Sair do jogo")
	default:
		fmt.Println("opção invalida")
	}
}

func convertMinutesToTime(mp int) time.Time {
	h := mp / 60
	m := mp % 60

	ts := fmt.Sprintf("%02d:%02d:00", h, m)
	t, _ := time.Parse(time.TimeOnly, ts)

	return t
}

func convertMinutesToDuration(mp int) time.Duration {
	s := fmt.Sprintf("%dm", mp)
	mpp, _ := time.ParseDuration(s)
	return mpp
}

func inicio(tt int, tr int) {

	total_segundos := tt * 60
	d := time.Duration(total_segundos) * time.Second
	t := convertMinutesToTime(tt)
	// d := convertMinutesToDuration(tt)

	msgs := [2]string{
		"Foco total ->",
		"Parabens, você conclui um periodo de foco, descanse um pouco agora!",
	}

	i := intervalo{
		d:         d,
		t:         t,
		mensagens: msgs,
	}

	rodaIntervalo(i)
}

func inicioRapido() {
	inicio(TEMPO_TRABALHO, TEMPO_REPOUSO)
}

func inicioConfigurado() {}

func rodaIntervalo(i intervalo) {
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		time.Sleep(i.d)
		done <- true
	}()

	for {
		select {
		case <-done:
			fmt.Printf("\n%v", i.mensagens[1])
			return
		case <-ticker.C:
			i.t = i.t.Add(time.Second * -1)
			tf := fmt.Sprintf("\r%v %v", i.mensagens[0], i.t.Format(time.TimeOnly))
			fmt.Print(tf)
		}
	}
}

func main() {
	// menuInicial()
	inicioRapido()
}
