package main

import "fmt"

type SagaStep struct {
	Action     func() error
	Compensate func()
	DependsOn  []int
}

type Orchestrator struct {
	steps         []SagaStep
	executedSteps []int
}

func (o *Orchestrator) AddStep(action func() error, compensate func(), dependsOn ...int) {
	o.steps = append(o.steps, SagaStep{
		Action:     action,
		Compensate: compensate,
		DependsOn:  dependsOn,
	})
}

func (o *Orchestrator) Execute() error {
	executed := make(map[int]bool)
	for i := 0; i < len(o.steps); i++ {
		step := o.steps[i]
		if !o.canExecute(step, executed) {
			continue
		}
		if err := step.Action(); err != nil {
			fmt.Printf("Step %d failed: %v\n", i+1, err)
			o.rollback()
			return err
		}
		o.executedSteps = append(o.executedSteps, i)
		executed[i] = true
	}
	return nil
}

func (o *Orchestrator) canExecute(step SagaStep, executed map[int]bool) bool {
	for _, dep := range step.DependsOn {
		if !executed[dep] {
			return false
		}
	}
	return true
}

func (o *Orchestrator) rollback() {
	fmt.Println("Rolling back...")
	for i := len(o.executedSteps) - 1; i >= 0; i-- {
		stepIndex := o.executedSteps[i]
		o.steps[stepIndex].Compensate()
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
	orchestrator.AddStep(
		func() error {
			fmt.Println("Step 2: Procesando Orden")
			return nil
		},
		func() {
			fmt.Println("Compensando Step 2: Cancelando Orden")
		},
		0,
	)
	orchestrator.AddStep(
		func() error {
			fmt.Println("Step 3: Confirmando pedido")
			return nil
			// return fmt.Errorf("error al confirmar pedido")
		},
		func() {
			fmt.Println("Compensando Step 3: Cancelando pedido")
		},
		1,
	)
	orchestrator.AddStep(
		func() error {
			fmt.Println("Step 4: Enviando pedido")
			return nil
		},
		func() {
			fmt.Println("Compensando Step 4: Cancelando envío")
		},
		2,
	)
	orchestrator.AddStep(
		func() error {
			fmt.Println("Step 5: Enviar Notificación")
			return nil
		},
		func() {
			fmt.Println("Compensando Step 5: Cancelando Notificación")
		},
		2,
	)
	if err := orchestrator.Execute(); err != nil {
		fmt.Println("Saga terminó con error.")
	}
}
