package main

import "fmt"

type PipelineTaskInterface interface {
    say()
}

type PipelineTask struct {
	name string
}

func (b *PipelineTask) say() {
	fmt.Println(b.name)
}

type PipelineTryTask struct {
	PipelineTask  //embedding
	runAfter string
}

func check(b PipelineTaskInterface) {
	b.say()
}

func main() {
	PipelineTask := PipelineTask{name: "Task1"}
	PipelineTryTask := &PipelineTryTask{
		PipelineTask:  PipelineTask,
		runAfter: "Task0",
	}
	PipelineTryTask.say()
	check(PipelineTryTask)
}