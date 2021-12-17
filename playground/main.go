package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		taskCount      int
		allTasks       [100]string
		remainderCount int
		taskIdx        int
		remainderTasks [100]string
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Scan(&taskCount)
	//taskCountS, _ := reader.ReadString('\n')
	//taskCountS = strings.Replace(taskCountS, "\n", "", -1)
	//taskCount, _ = strconv.Atoi(taskCountS)

	for i := 1; i <= taskCount; i++ {
		task, _ := reader.ReadString('\n')
		task = strings.Replace(task, "\n", "", -1)
		allTasks[i-1] = task
	}
	//fmt.Println(allTasks)

	fmt.Scan(&remainderCount)
	//remainderCountS, _ := reader.ReadString('\n')
	//remainderCountS = strings.Replace(remainderCountS, "\n", "", -1)
	//remainderCount, _ = strconv.Atoi(remainderCountS)

	for i := 1; i <= remainderCount; i++ {
		fmt.Scan(&taskIdx)
		//taskIdxS, _ := reader.ReadString('\n')
		//taskIdxS = strings.Replace(taskIdxS, "\n", "", -1)
		//taskIdx, _ = strconv.Atoi(remainderCountS)

		remainderTasks[i-1] = allTasks[taskIdx-1]
	}

	for i := 0; i <= remainderCount-1; i++ {
		fmt.Println(remainderTasks[i])
	}
}
