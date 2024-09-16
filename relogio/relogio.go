package relogio

import (
	"fmt"
	"time"
)

type relogio struct {
	minutos_trabalho int
	minutos_repouso  int
	loops            int
}

func Relogio() relogio {
	var r relogio
	return r
}

func Inicia() {
	r := Relogio()
	r.minutos_trabalho = 1
	r.minutos_repouso = 1
	r.loops = 3
	cronometro(r)
}

func cronometro(r relogio) {

}

func iniciaTimer(r relogio) {
	ciclo := 0

	if ciclo < r.loops {
		// contaTempo
	}
	contaTempo(r)
}

func formatarTempo(h int, m int, s int) string {
	tempo_formatado := fmt.Sprintf("\r Periodo de foco - %02d:%02d:%02d", h, m, s)
	return tempo_formatado
}

func contaTempo(r relogio) bool {

	horas_trabalho := r.minutos_trabalho / 60
	minutos_trabalho := r.minutos_trabalho % 60
	segundos_trabalho := r.minutos_trabalho / 600

	ticker := time.NewTicker(1 * time.Second)

	ticks_trabalho := r.minutos_trabalho * 60
	ticks_repouso := 50
	ticks_total := ticks_trabalho + ticks_repouso

	fmt.Println(ticks_total)

	fim_periodo := false

	go func() {
		for range ticker.C {
			if horas_trabalho == 0 && minutos_trabalho == 0 && segundos_trabalho == 0 {
				fim_periodo = true
			}
			if !fim_periodo {
				if horas_trabalho < 24 {
					if minutos_trabalho > 0 {
						if segundos_trabalho > 0 {
							segundos_trabalho--
						} else {
							segundos_trabalho = 59
							minutos_trabalho--
						}
					} else {
						minutos_trabalho = 59
						horas_trabalho--
					}
				} else {
					fmt.Errorf("Duração maior que 24 horas. \n Por favor escolha ua duração menor que 24 horas")
				}
				tempo_formatado := formatarTempo(horas_trabalho, minutos_trabalho, segundos_trabalho)
				fmt.Print(tempo_formatado)
			}
		}
	}()

	duracao := time.Duration(ticks_total) * time.Second
	time.Sleep(duracao)
	ticker.Stop()
	return true
}
