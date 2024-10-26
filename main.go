package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	matriz := [5][5]int{
		{0, 1, 0, 1, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 0, 1, 0},
		{0, 1, 0, 0, 1},
		{1, 0, 1, 0, 0},
	}
	var newMatriz [5][5]int
	var wg sync.WaitGroup

	threadTimes := make(map[string]time.Duration)
	var mutex sync.Mutex

	start := time.Now()

	for i := range matriz {
		for j := range matriz[i] {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				threadStart := time.Now()
				newValue := 0
				for x := 0; x < len(matriz); x++ {
					newValue += matriz[i][x] * matriz[x][j]
				}
				newMatriz[i][j] = newValue
				threadEnd := time.Since(threadStart)
				mutex.Lock()
				threadTimes[fmt.Sprintf("Thread[%d,%d]", i, j)] = threadEnd
				mutex.Unlock()
			}(i, j)
		}
	}

	wg.Wait()
	multiplicationTime := time.Since(start)

	fmt.Println("Matriz resultante de A^2:")
	for _, row := range newMatriz {
		fmt.Println(row)
	}

	fmt.Println("\nTempos de execução de cada thread:")
	for key, duration := range threadTimes {
		fmt.Printf("%s: %v\n", key, duration)
	}

	fmt.Println("Matriz resultante de A^2:")
	for _, row := range newMatriz {
		fmt.Println(row)
	}

	startAnalysis := time.Now()

	mostFriends := make(map[int]int)
	for i := range newMatriz {
		for j := range newMatriz[i] {
			mostFriends[i] += newMatriz[i][j]
		}
	}

	influence := make(map[int]int)
	for i := range newMatriz {
		for j := range newMatriz[i] {
			influence[j] += newMatriz[i][j]
		}
	}

	var person1, person2 int
	for p, friends := range mostFriends {
		if friends > mostFriends[person1] {
			person2 = person1
			person1 = p
		} else if friends > mostFriends[person2] {
			person2 = p
		}
	}

	var mostInfluential int
	for p, score := range influence {
		if score > influence[mostInfluential] {
			mostInfluential = p
		}
	}

	analysisTime := time.Since(startAnalysis)

	fmt.Printf("\nPessoas com mais amigos em comum: %d e %d\n", person1+1, person2+1)
	fmt.Printf("Pessoa mais influente: %d\n", mostInfluential+1)

	fmt.Printf("\nTempo de multiplicação: %v\n", multiplicationTime)
	fmt.Printf("Tempo de análise: %v\n", analysisTime)
}
