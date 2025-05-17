# Go CLI Games Template 🎮

A beginner-friendly Go CLI project template for implementing classic terminal games like "Guess the Number" and "Tic Tac Toe".

This template is designed for **learning Go by building**. You’ll find modular game stubs, clear TODOs, and a working CLI engine ready for you to extend.

---

## ✅ Features

- 🔧 Simple and modular [Cobra](https://github.com/spf13/cobra)-based CLI tool
- 📁 Each game lives in its own folder under `games/`
- 🧩 Game logic is left unimplemented — your task is to fill in the blanks
- 🧠 Encourages Go fundamentals: structs, slices, I/O, flow control, etc.
- 🧰 Easily expandable with new games (e.g., 2048, RPG, Snake...)

---

## 🏁 Getting Started

### 1. Use This Template

Click **“Use this template”** on GitHub and create your own repo — e.g.:
`go-cli-games-implementation`

> You can also fork the repo, but using the template feature is cleaner for solo learners.

---

### 2. Clone Your New Repo

```bash
git clone https://github.com/YOUR_USERNAME/go-cli-games-implementation.git
cd go-cli-games-implementation
```

### 3. Initialize Dependencies

```bash
go mod tidy
```

### 4. Run the CLI
```bash
go run main.go play
```
You’ll be prompted to select a game. The game logic will print TODOs until you implement it.

---

## 🧱 Project Structure

```
go-cli-games/
├── main.go                    # CLI entrypoint
├── go.mod
├── cmd/                       # Cobra commands
│   ├── root.go                # Root command
│   └── play.go                # "play" subcommand (game selector)
├── games/
│   ├── guessnumber/
│   │   └── game.go            # TODO: implement game logic here
│   └── tictactoe/
│       └── game.go            # Includes Game struct + Play stub
├── Makefile                   # (Optional) for quick commands
└── README.md                  # This file
```

---

## 🧪 Example Commands

```bash
make run       # Run the game selector
make tidy      # Clean up modules
make build     # Build binary
```

---

## 🧠 Your Learning Tasks

| Game             | Status           | What to Implement                                      |
|------------------|------------------|--------------------------------------------------------|
| Guess the Number | 🧩 Scaffolding    | Use `rand`, accept user input, give feedback loop      |
| Tic Tac Toe      | 🧩 Partial Struct | Game loop, turn switching, win/draw detection           |
| Add Your Own     | 🔲 Empty          | Create folder, define `Play()` logic, register in CLI  |
| Optional AI      | 🔄 Expandable     | Add random / defensive AI for Tic Tac Toe              |
| UI Enhancements  | ✨ Optional       | Pretty-print board with row/column headers             |


---

## 🧰 How to Add a New Game

1. Create a folder under games/, like games/snake/
2. Add a game.go file with a Play() function
3. Import and register it in cmd/play.go
4. Done!

---

## 🔗 See a Sample Solution?
You can check out go-cli-games(TODO) as a sample reference for finished implementations.

---

## 👋 Author

Created by Jiaxing Yan

For learning, fun, and sharing within teams or beginner Go groups.


