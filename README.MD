

# 📝 CLI Tasker

A simple command-line task manager built with Go.

## 🚀 Features

- Add new tasks
- List tasks with status
- Mark tasks as completed
- Cancel tasks

Tasks are stored locally in a `tasks.json` file.

## 📦 Requirements

- Go 1.18 or higher
- Dependency: [urfave/cli](https://github.com/urfave/cli)

## 🛠️ Installation

Clone the project and run it with:

```bash
go run main.go [command]
```

Or build a binary:

```bash
go build -o tasker main.go
./tasker [command]
```

## 📋 Available Commands

| Command    | Description                  | Flags      |
|------------|------------------------------|------------|
| `add`      | Adds a new task              | `--name`   |
| `list`     | Lists all tasks              | *(none)*   |
| `done`     | Marks a task as completed    | `--id`     |
| `cancel`   | Cancels an existing task     | `--id`     |

## 💡 Examples

```bash
./tasker add --name "Study Go"
./tasker list
./tasker done --id 1
./tasker cancel --id 2
```

## 📂 Project Structure

```
cli-tasker/
├── app/         # CLI and command definitions
├── task/        # Task logic and persistence
├── tasks.json   # Automatically created
└── main.go
```

## ✅ Status

✅ Fully functional and ready for local use.
