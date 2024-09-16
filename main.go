package main

import (
	"fmt"
	"time"
)

// import logica "github.com/vidalme/gomodoro/controller"

// func timeLeft(t time.Time) int {

// }

// func justPlay() {

// 	wt := convertMinutesToTime(duracao_inicial)

// 	duracao_restante := timeLeft(wt)

// 	fmt.Println(wt.Format(time.TimeOnly))

// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	done := make(chan bool)

// 	go func() {
// 		time.Sleep(time.Second * time.Duration(duracao_inicial))
// 		done <- true
// 	}()
// 	for {
// 		select {
// 		case <-done:
// 			fmt.Println("Done!")
// 			return
// 		case t := <-ticker.C:
// 			tf := fmt.Sprintf("\r Current time: %v", t.Format(time.TimeOnly))
// 			fmt.Print(tf)
// 		}
// 	}

// }

func convertMinutesToTime(mp int) time.Time {
	h := mp / 60
	m := mp % 60
	s := mp * 60
	// fmt.Printf("converteu os %d minutos em %d:%d:%d \n", mp, h, s, m)

	// time.Date(0, 0, 0, h, m, s, 0, time.Local)

	timeString := fmt.Sprintf("%02d:%02d:%02d", h, m, s)

	// t, _ := time.Parse(time.TimeOnly, strconv.Itoa(h)+":"+strconv.Itoa(m)+":"+strconv.Itoa(s))
	t, _ := time.Parse(time.TimeOnly, "05:05:01")
	fmt.Println("------", timeString)
	fmt.Println("------", t)

	return t
}

func main() {
	// logica.Iniciar()
	// justPlay()
	total_minutos := 65
	total_segundos := total_minutos * 60

	countdownDuration := time.Duration(total_segundos) * time.Second
	fmt.Println(countdownDuration)

	countdown := convertMinutesToTime(total_minutos)

	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		time.Sleep(countdownDuration)
		done <- true
	}()

	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case <-ticker.C:
			// tf1 := fmt.Sprintf("\rCurrent time: %v", t.Format(time.TimeOnly))
			// fmt.Print(tf1)

			// tf2 := fmt.Sprintf("Countdown: %v \n", time.ParseDuration( countdownDuration.String() ))
			// fmt.Print(tf2)

			fmt.Println(countdownDuration)
			fmt.Println(countdown.Format(time.TimeOnly))
		}
	}
}
