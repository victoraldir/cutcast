package command

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

//go:generate mockgen -destination=../command/mocks/mockCommandBuilder.go -package=command github.com/victoraldir/cutcast/pkg/command CommandBuilder
type CommandBuilder interface {
	Build(command string, args ...string) CommandExecutor
}

type commandBuilder struct{}

func NewCommandBuilder() CommandBuilder {
	return &commandBuilder{}
}

func (c *commandBuilder) Build(command string, args ...string) CommandExecutor {

	ctx := context.Background()

	return NewCommandExecutor(
		exec.CommandContext(ctx, command, args...),
	)
}

//go:generate mockgen -destination=../command/mocks/mockCommandExecutor.go -package=command github.com/victoraldir/cutcast/pkg/command CommandExecutor
type CommandExecutor interface {
	Run() error
	Stop() error
	Signal() error
}

type commadExecutor struct {
	cmd *exec.Cmd
	mu  *sync.Mutex
}

func NewCommandExecutor(cmd *exec.Cmd) CommandExecutor {
	return &commadExecutor{
		cmd: cmd,
		mu:  &sync.Mutex{},
	}
}

func (c *commadExecutor) Run() error {

	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr

	return c.cmd.Run()
}

func (c *commadExecutor) Stop() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.cmd == nil {
		return fmt.Errorf("no download in progress")
	}

	if err := c.cmd.Process.Signal(os.Interrupt); err != nil {
		return err
	}

	c.cmd = nil

	return nil
}

func (c *commadExecutor) Signal() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.cmd.Process.Signal(os.Interrupt)
}
