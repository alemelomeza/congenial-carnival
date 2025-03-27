package main

import "fmt"

type SagaStep struct {
	Action     func() error
	Compensate func()
}

type Orchestrator struct {
	steps []SagaStep
}

func (o *Orchestrator) AddStep(action func() error, compoensate func()) {
	o.steps = append(o.steps, SagaStep{
		Action:     action,
		Compensate: compoensate,
	})
}

func (o *Orchestrator) Execute() error {
	completedSteps := 0

	for i, step := range o.steps {
		if err := step.Action(); err != nil {
			fmt.Printf("Step %d failed: %v\n", i+1, err)
			o.rollback(i)
			return err
		}
		completedSteps++
	}

	fmt.Println("Saga completed successfully.")
	return nil
}

func (o *Orchestrator) rollback(failedStep int) {
	fmt.Println("Rolling back...")
	for i := failedStep; i >= 0; i-- {
		o.steps[i].Compensate()
	}
	fmt.Println("Rollback complete.")
}

func main() {
	orchestrator := &Orchestrator{}
	orchestrator.AddStep(
		func() error {
			fmt.Println("Step 1: Reservando recursos")
			return nil
		},
		func() {
			fmt.Println("Compensando Step 1: Liberando recursos")
		},
	)
	if err := orchestrator.Execute(); err != nil {
		fmt.Println("Saga terminó con error.")
	}
}
