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
	PipelineTask := PipelineTask{name: "Task1"}

	TryPipelineTask := &TryPipelineTask{
		PipelineTask:  PipelineTask,
		runAfter: "Task0",
	}

	FinalPipelineTask := &FinalPipelineTask{
		PipelineTask:  PipelineTask,
	}

	TryPipelineTask.printTaskName()
	check(TryPipelineTask)

	TryPipelineTask.printRunAfter()
//	checkRunAfter(TryPipelineTask)

	FinalPipelineTask.printTaskName()
	check(FinalPipelineTask)

}
