package main

import (
	"fmt"
	"sync"
)

// Мьютекс - средство защиты от "Resource Sharing".
// Во время работы конкурентных программ, главной проблемой является то, что множество горутин
// не должны одновременно использовать какой-то общий ресурс (переменная/файл/бд) для модификации.

// Пример Resource Sharing.
// x += 1
// Если работает одна горутина - проблем нет.
// Если горутины 2 или более - т.к. для обоих горутин данный ресурс будет использоваться конкурентно,
// НЕЛЬЗЯ утверждать, что после выполнения горутин (при начальном x == 0) x == 2.
// 1. Первая и вторая горутина могут начать работать с параметром x == 0. Т.е. вторая не дождется, пока первая увеличит x.
// x == 1
// 2. Первая горутина начнет работать, увеличит x, но не успеет присвоить его.
// x == 1
// 3. Первая горутина начнет работать, завершится, потом стартует вторая. Оптимистичный вариант - его вероятность - 1/3.
// x == 2

// Для решения данной проблемы используют мьютексы. Мьютекс - блокирует ресурс, пока его не освободит одна из горутин.
// mutex.Lock()
// x += 1
// mutex.Unlock()

// Resource Sharing часто называют Race Condition (гонка за ресурс).

var x int = 0

func increment(wg *sync.WaitGroup) {
	x += 1
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("Final value of x:", x) // почти никогда не будет 1000
}
