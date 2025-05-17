## 📘 `CONTRIBUTING.md`

# Contributing to Go CLI Games Template

This project is a learning playground. Feel free to fork, extend, or contribute your own games.

## 🕹️ How to Add a New Game

1. Create a new folder under `games/`, like `games/2048/`
2. Add a `game.go` file with a `Play()` function
3. Update `cmd/play.go` to add a new entry to the selection menu
4. Test it with `go run main.go play`

## 📦 Best Practices

- Keep each game self-contained
- Use `Play()` as the game entrypoint
- Write readable and beginner-friendly code
- Use comments to guide the user if logic is complex
- Avoid using 3rd-party dependencies unless absolutely necessary

## 💡 Ideas for New Games

- 2048
- Rock Paper Scissors
- Simple Text RPG
- Snake (ASCII version)
- Memory Match

---

This template is about **learning Go by doing** — enjoy the journey!