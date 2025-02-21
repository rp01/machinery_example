package sender

import (
	"fmt"
	"mchn/server"
	"time"

	"github.com/RichardKnop/machinery/v2/log"
	"github.com/RichardKnop/machinery/v2/tasks"
)

func Send() error {
	// cleanup, err := tracers.SetupTracer("sender")
	// if err != nil {
	// 	log.FATAL.Fatalln("Unable to instantiate a tracer:", err)
	// }
	// defer cleanup()

	serve, err := server.StartServer()
	if err != nil {
		return err
	}

	var (
		addTask0, addTask1, addTask2 tasks.Signature
	)
	var initTasks = func() {
		addTask0 = tasks.Signature{
			Name: "add",
			Args: []tasks.Arg{
				{
					Type:  "int64",
					Value: 1,
				},
				{
					Type:  "int64",
					Value: 1,
				},
			},
		}

		addTask1 = tasks.Signature{
			Name: "add",
			Args: []tasks.Arg{
				{
					Type:  "int64",
					Value: 2,
				},
				{
					Type:  "int64",
					Value: 2,
				},
			},
		}

		addTask2 = tasks.Signature{
			Name: "add",
			Args: []tasks.Arg{
				{
					Type:  "int64",
					Value: 5,
				},
				{
					Type:  "int64",
					Value: 6,
				},
			},
		}
	}
	initTasks()

	// asyncResult, err := serve.SendTask(&addTask0)
	// if err != nil {
	// 	return fmt.Errorf("could not send task: %s", err.Error())
	// }
	// log.INFO.Println("Task UUID:", asyncResult.Signature.UUID)
	// asyncResult, err := serve.SendTask(&addTask0)
	// if err != nil {
	// 	return fmt.Errorf("could not send task: %s", err.Error())
	// }
	// log.INFO.Println("Task UUID:", asyncResult.Signature.UUID)

	log.INFO.Println("Group of tasks (parallel execution):")

	group, err := tasks.NewGroup(&addTask0, &addTask1, &addTask2)
	if err != nil {
		return fmt.Errorf("error creating group: %s", err.Error())
	}

	asyncResults, err := serve.SendGroup(group, 0)
	if err != nil {
		return fmt.Errorf("could not send group: %s", err.Error())
	}

	for _, asyncResult := range asyncResults {
		results, err := asyncResult.Get(time.Millisecond * 5)
		if err != nil {
			return fmt.Errorf("getting task result failed with error: %s", err.Error())
		}
		log.INFO.Printf(
			"%v + %v = %v\n",
			asyncResult.Signature.Args[0].Value,
			asyncResult.Signature.Args[1].Value,
			tasks.HumanReadableResults(results),
		)
	}

	return nil

}
