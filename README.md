
# Todo list cli made in GO

This is a todo list program made in GO, using cobra-cli and lip gloss packages of go

## Run Locally

Clone the project

```bash
  git clone https://github.com/Kaif-S/cli-taskmanager.git
```

Go to the project directory

```bash
  cd cli-taskmanager
```

Install dependencies

```bash
  go mod download
```

Verify dependencies

```bash
  go mod tidy
```
Build (optional for executable)

```bash
  go build
```

```bash
  go Install
```

```bash
  todo --help
```

```bash
Todo application to manage your tasks in command line interface

Usage:
  todo [command]

Available Commands:
  add         Add task by giving id and Title in argument
  completion  Generate the autocompletion script for the specified shell
  delete      Enter ID for argument to delete that task from tasks
  help        Help about any command
  list        list all of your tasks
  update      Update task to true of false with id

Flags:
  -h, --help     help for todo
  -t, --toggle   Help message for toggle

Use "todo [command] --help" for more information about a command.
```