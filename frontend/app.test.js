const { fireEvent } = require('@testing-library/dom');
require('@testing-library/jest-dom');

/**
 * @jest-environment jsdom
 */

// Example: test that the board renders 9 cells
test('renders 9 cells', () => {
  document.body.innerHTML = `
    <div id="board">
      ${Array.from({ length: 9 }).map((_, i) => `<div class="cell" data-index="${i}"></div>`).join('')}
    </div>
  `;
  const cells = document.querySelectorAll('.cell');
  expect(cells.length).toBe(9);
});

// Example: test clicking a cell
test('cell click updates content', () => {
  document.body.innerHTML = `
    <div id="board">
      <div class="cell" data-index="0"></div>
    </div>
  `;
  const cell = document.querySelector('.cell');
  cell.textContent = '';
  fireEvent.click(cell);
  // Simulate your click logic here, e.g.:
  cell.textContent = 'X';
  expect(cell).toHaveTextContent('X');
});