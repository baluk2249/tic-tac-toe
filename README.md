
# Tic-Tac-Toe Web App

This is a full-stack Tic-Tac-Toe game with a Go backend and a JavaScript frontend.

## Features
- Play Tic-Tac-Toe in your browser
- Real-time board updates and win/draw detection
- Highlights the winning line
- Restart button and turn/win/draw messages
- Modular Go backend with REST API
- Frontend and backend tests

---

## Getting Started

### Prerequisites
- Go 1.18+
- Node.js (for frontend tests)
- Python 3 (for simple static server)

### Backend Setup
1. Open a terminal in the project root.
2. Initialize Go module (if not already):
	```bash
	go mod init tic-tac-toe
	```
3. Start the backend server:
	```bash
	cd backend
	go run main.go
	```
	The backend runs at http://localhost:9090

### Frontend Setup
1. Open a new terminal.
2. Start a static server in the `frontend` folder:
	```bash
	cd frontend
	python3 -m http.server 9010
	```
	The frontend is at http://localhost:9010

### How to Play
- Open the frontend in your browser.
- Click on a cell to make a move.
- The board updates in real time.
- The message area shows whose turn it is, or who won/draw.
- Click Restart to reset the game.

---

## API Endpoints

- `GET /state` — Get current game state (board, status, winner, etc.)
- `POST /move` — Make a move. Body: `{ "row": 0, "col": 1 }`
- `POST /restart` — Reset the game

---

## Testing

### Backend (Go)
```bash
cd backend/game
go test -v
```

### Frontend (Jest)
```bash
cd frontend
npm install
npm test
```

---

## Project Structure

- `backend/` — Go backend (API, game logic)
- `frontend/` — HTML, CSS, JS frontend
- `backend/game/logic.go` — Game logic
- `backend/game/types.go` — Game types
- `frontend/app.js` — Frontend logic
- `frontend/app.test.js` — Frontend tests

---

