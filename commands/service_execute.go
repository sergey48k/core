package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/logrusorgru/aurora"
	"github.com/mesg-foundation/core/service"
	"github.com/mesg-foundation/core/utils/xpflag"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

type serviceExecuteCmd struct {
	baseCmd

	executeData map[string]string
	taskKey     string
	jsonFile    string

	e ServiceExecutor
}

func newServiceExecuteCmd(e ServiceExecutor) *serviceExecuteCmd {
	c := &serviceExecuteCmd{e: e}
	c.cmd = newCommand(&cobra.Command{
		Use:     "execute",
		Short:   "Execute a task of a service",
		Example: `mesg-core service execute SERVICE`,
		Args:    cobra.ExactArgs(1),
		PreRunE: c.preRunE,
		RunE:    c.runE,
	})
	c.cmd.Flags().StringVarP(&c.taskKey, "task", "t", c.taskKey, "Run the given task")
	c.cmd.Flags().VarP(xpflag.NewStringToStringValue(&c.executeData, nil), "data", "d", "data required to run the task")
	c.cmd.Flags().StringVarP(&c.jsonFile, "json", "j", c.jsonFile, "Path to a JSON file containing the data required to run the task")
	return c
}

func (c *serviceExecuteCmd) preRunE(cmd *cobra.Command, args []string) error {
	if cmd.Flag("data").Changed && cmd.Flag("json").Changed {
		return errors.New("Ynu can specify only one of '--data' or '--json' options")
	}
	return nil
}

func (c *serviceExecuteCmd) runE(cmd *cobra.Command, args []string) error {
	s, err := c.e.ServiceByID(args[0])
	if err != nil {
		return err
	}

	if err := c.getTaskKey(s); err != nil {
		return err
	}

	inputData, err := c.getData(c.taskKey, s, c.executeData)
	if err != nil {
		return err
	}

	// Create an unique tag that will be used to listen to the result of this exact execution
	tags := []string{uuid.NewV4().String()}
	listenResults, err := c.e.ServiceExeucteTask(args[0], c.taskKey, inputData, tags)
	if err != nil {
		return err
	}

	result := <-listenResults
	fmt.Println("Task " + aurora.Green(c.taskKey).String() + " returned output " + aurora.Blue(result.OutputKey).String() + " with data:")
	fmt.Println(aurora.Bold(result.OutputData).String())
	return nil
}

func (c *serviceExecuteCmd) getTaskKey(s *service.Service) error {
	if c.taskKey == "" {
		if survey.AskOne(&survey.Select{
			Message: "Select the task to execute",
			Options: taskKeysFromService(s),
		}, &c.taskKey, nil) != nil {
			return errors.New("no task to execute")
		}
	}
	return nil
}

func (c *serviceExecuteCmd) getData(taskKey string, s *service.Service, dataStruct map[string]string) (string, error) {
	if dataStruct != nil {
		castData, err := s.Cast(taskKey, dataStruct)
		if err != nil {
			return "", err
		}

		b, err := json.Marshal(castData)
		return string(b), err
	}

	if c.jsonFile == "" {
		if survey.AskOne(&survey.Input{Message: "Enter the filepath to the inputs"}, &c.jsonFile, nil) != nil {
			return "", errors.New("no filepath given")
		}
	}
	return readJSONFile(c.jsonFile)
}

func readJSONFile(path string) (string, error) {
	if path == "" {
		return "{}", nil
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func taskKeysFromService(s *service.Service) []string {
	var taskKeys []string
	for key := range s.Tasks {
		taskKeys = append(taskKeys, key)
	}
	return taskKeys
}
