package main

import "fmt"

type PipelineTaskInterface interface {
    printTaskName()
}

type PipelineTask struct {
	name string
}

type TryPipelineTask struct {
	PipelineTask  //embedding
	runAfter string
}

type FinalPipelineTask struct {
	PipelineTask  //embedding
}

func (b *PipelineTask) printTaskName() {
	fmt.Println(b.name)
}

func (b *TryPipelineTask) printRunAfter() {
	fmt.Println(b.runAfter)
}

func check(b PipelineTaskInterface) {
	b.printTaskName()
}

//func checkRunAfter(b *TryPipelineTask) {
//	b.printRunAfter()
//}

func main() {
	pipelineTask := PipelineTask{name: "Task1"}

	tryPipelineTask := &TryPipelineTask{
		PipelineTask:  pipelineTask,
		runAfter: "Task0",
	}

	finalPipelineTask := &FinalPipelineTask{
		PipelineTask:  pipelineTask,
	}

	tryPipelineTask.printTaskName()
	check(tryPipelineTask)

	tryPipelineTask.printRunAfter()
//	checkRunAfter(TryPipelineTask)

	finalPipelineTask.printTaskName()
	check(finalPipelineTask)
}
