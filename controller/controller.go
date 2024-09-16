package logica

import (
	"fmt"

	"github.com/vidalme/gomodoro/relogio"
)

var mensagem_welcome string = "[[ Bem vindo ao Gomodoro, seu pomodoro no terminal ]]"
var mensagem_inicio_rapido string = "[1] Inicio Rapido"
var mensagem_configurar_relogio string = "[2] Configurar Relogio"
var mensagem_sair_do_app string = "[x] Sair do Gomodoro"

// var mensagem_digite_tempo_trabalhando string = "Digite quanto tempo quer ficar em modo foco: "
// var mensagem_digite_tempo_intervalo string = "Digite quanto tempo quer ficar em modo intervalo: "
// var tempo_trabalho int = 45
// var tempo_intervalo int = 15

func menuInicial() {
	fmt.Println(mensagem_inicio_rapido)
	fmt.Println(mensagem_configurar_relogio)
	fmt.Println(mensagem_sair_do_app)
}

func mensagensInicio() {
	fmt.Println(mensagem_welcome)
	menuInicial()
}

func escolherOpcao() string {
	var opcao string
	fmt.Scanln(&opcao)

	switch opcao {
	case "1":
		fmt.Println("1 foi escolhido :", opcao)
		relogio.Inicia()
	case "2":
		fmt.Println("2 foi escolhido :", opcao)
		// relogio.Inicia(20)
	case "x":
		fmt.Println("x foi escolhido :", opcao)
	}

	return opcao
}

func Iniciar() {
	mensagensInicio()
	escolherOpcao()
}
