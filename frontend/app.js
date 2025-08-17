
const API_BASE = 'http://localhost:9090';
const boardDiv = document.getElementById('board');
const messageDiv = document.getElementById('message');

function renderBoard(state) {
    const cells = boardDiv.querySelectorAll('.cell');
    const gameOver = state.status === 'win' || state.status === 'draw';
    // Remove win highlight from all cells first
    cells.forEach(cell => cell.classList.remove('win'));
    for (let i = 0; i < 9; i++) {
        const row = Math.floor(i / 3);
        const col = i % 3;
        const val = state.board[row][col];
        cells[i].textContent = val;
        cells[i].classList.remove('x', 'o', 'disabled');
        if (val === 'X') cells[i].classList.add('x');
        if (val === 'O') cells[i].classList.add('o');
        // Disable cell visually if game is over
        if (gameOver) cells[i].classList.add('disabled');
        // Highlight winning line
        if (state.status === 'win' && Array.isArray(state.winning_line)) {
            for (const [r, c] of state.winning_line) {
                if (row === r && col === c) {
                    cells[i].classList.add('win');
                }
            }
        }
    }
    if (state.status === 'win') {
        messageDiv.textContent = `Player ${state.winner} wins! Click Restart to play again.`;
    } else if (state.status === 'draw') {
        messageDiv.textContent = "It's a draw! Click Restart to play again.";
    } else {
        messageDiv.textContent = `Player ${state.current_player}'s turn`;
    }
}

async function fetchState() {
    try {
        const res = await fetch(`${API_BASE}/state`);
        const state = await res.json();
        renderBoard(state);
    } catch (e) {
        messageDiv.textContent = 'Failed to load game state.';
    }
}


function cellIndexToRowCol(index) {
    return { row: Math.floor(index / 3), col: index % 3 };
}

async function handleCellClick(e) {
    const index = parseInt(e.target.getAttribute('data-index'));
    const { row, col } = cellIndexToRowCol(index);
    // Prevent move if already filled or game over
    if (
        e.target.textContent !== '' ||
        messageDiv.textContent.includes('wins') ||
        messageDiv.textContent.includes('draw') ||
        e.target.classList.contains('disabled')
    ) {
        return;
    }
    try {
        const res = await fetch(`${API_BASE}/move`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ row, col })
        });
        if (!res.ok) {
            const err = await res.json();
            messageDiv.textContent = err.error || 'Invalid move.';
        }
        // Always fetch the latest state after a move
        await fetchState();
    } catch (e) {
        messageDiv.textContent = 'Failed to send move.';
    }
}

window.addEventListener('DOMContentLoaded', () => {
    fetchState();
    document.querySelectorAll('.cell').forEach(cell => {
        cell.addEventListener('click', handleCellClick);
    });
    const restartBtn = document.getElementById('restart');
    if (restartBtn) {
        restartBtn.addEventListener('click', async () => {
            try {
                await fetch(`${API_BASE}/restart`, { method: 'POST' });
                await fetchState();
            } catch (e) {
                messageDiv.textContent = 'Failed to restart game.';
            }
        });
    }
});
