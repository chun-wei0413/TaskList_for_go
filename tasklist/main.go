package main

import (
	"TaskList_for_go/tasklist/entity"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	Quit   = "quit"
	prompt = "> "
)

type TaskManager struct {
	in       io.Reader
	out      io.Writer
	todoList *entity.ToDoList
}

func NewTaskManager(in io.Reader, out io.Writer) *TaskManager {
	todoList := entity.NewToDoList(*entity.NewToDoListId(1))
	return &TaskManager{
		in:       in,
		out:      out,
		todoList: todoList,
	}
}

func (m *TaskManager) Run() {
	scanner := bufio.NewScanner(m.in)
	fmt.Fprint(m.out, prompt)
	for scanner.Scan() {
		cmdLine := scanner.Text()
		if cmdLine == Quit {
			return
		}
		m.execute(cmdLine)
		fmt.Fprint(m.out, prompt)
	}
}

func (m *TaskManager) execute(cmdLine string) {
	args := strings.Split(cmdLine, " ")
	command := args[0]
	switch command {
	case "show":
		m.show()
	case "add":
		m.add(args[1:])
	case "check":
		m.check(args[1])
	case "uncheck":
		m.uncheck(args[1])
	case "help":
		m.help()
	default:
		m.error(command)
	}
}

func (m *TaskManager) help() {
	fmt.Fprintln(m.out, `Commands:
  show
  add project <project name>
  add task <project name> <task description>
  check <task ID>
  uncheck <task ID>`)
}

func (m *TaskManager) error(command string) {
	fmt.Fprintf(m.out, "Unknown command \"%s\".\n", command)
}

func (m *TaskManager) show() {
	projects := m.todoList.GetProjects()
	for _, p := range projects {
		name := p.GetName().ToString()
		fmt.Fprintf(m.out, "%s\n", name)
		for _, task := range p.GetTasks() {
			done := ' '
			if task.GetDone() {
				done = 'X'
			}
			id := task.GetId()
			fmt.Fprintf(m.out, "    [%c] %d: %s\n", done, id.GetId(), task.GetDescription())
		}
		fmt.Fprintln(m.out)
	}
}

func (m *TaskManager) add(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(m.out, "Missing parameters for \"add\" command.")
		return
	}
	if args[0] == "project" {
		projectName := entity.NewProjectName(args[1])
		m.todoList.AddProject(*projectName)
	} else if args[0] == "task" {
		if len(args) < 3 {
			fmt.Fprintln(m.out, "Missing task description for \"add task\" command.")
			return
		}
		projectName := entity.NewProjectName(args[1])
		description := strings.Join(args[2:], " ")
		m.todoList.AddTask(*projectName, description, false)
	} else {
		fmt.Fprintf(m.out, "Invalid add command: %s\n", args[0])
	}
}

func (m *TaskManager) check(idString string) {
	m.setDone(idString, true)
}

func (m *TaskManager) uncheck(idString string) {
	m.setDone(idString, false)
}

func (m *TaskManager) setDone(idString string, done bool) {
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		fmt.Fprintf(m.out, "Invalid ID \"%s\".\n", idString)
		return
	}
	taskId := entity.NewTaskId(id)
	m.todoList.SetDone(*taskId, done)
}

func main() {
	NewTaskManager(os.Stdin, os.Stdout).Run()
}
