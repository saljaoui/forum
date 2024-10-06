function crosswordSolver(emptyPuzzle, words) {
    console.log(emptyPuzzle);
    console.log(words
        
    );
    
    
    try {
        // Check for correct types and empty inputs
        if (typeof emptyPuzzle !== 'string' || !Array.isArray(words)) {
            throw new Error('Invalid input types');
        }

        // Check for empty puzzle or words
        if (emptyPuzzle.trim() === '' || words.length === 0) {
            throw new Error('Empty puzzle or words list');
        }

        // Check if all words are strings
        if (words.some(word => typeof word !== 'string')) {
            throw new Error('All words must be strings');
        }

        // Convert the puzzle string to a 2D array
        const grid = emptyPuzzle.split('\n').map(row => row.trim().split(''));
        const rows = grid.length;
        const cols = grid[0].length;

        // Check for wrong puzzle format
        if (rows === 0 || cols === 0 || !grid.every(row => row.length === cols)) {
            throw new Error('Invalid puzzle format');
        }

        // Find all word positions in the grid and validate starting numbers
        const wordPositions = [];
        for (let r = 0; r < rows; r++) {
            for (let c = 0; c < cols; c++) {
                if (grid[r][c] !== '.' && grid[r][c] !== '0') {
                    const num = parseInt(grid[r][c]);
                    if (isNaN(num) || num < 1 || num > 2) {
                        throw new Error(`Invalid starting number ${grid[r][c]} at (${r},${c})`);
                    }
                    let wordCount = 0;
                    // Check horizontal
                    if ((c === 0 || grid[r][c-1] === '.') && c + 1 < cols && grid[r][c+1] !== '.') {
                        let length = 0;
                        while (c + length < cols && grid[r][c + length] !== '.') length++;
                        if (length > 1) {
                            wordPositions.push({ row: r, col: c, length, direction: 'horizontal' });
                            wordCount++;
                        }
                    }
                    // Check vertical
                    if ((r === 0 || grid[r-1][c] === '.') && r + 1 < rows && grid[r+1][c] !== '.') {
                        let length = 0;
                        while (r + length < rows && grid[r + length][c] !== '.') length++;
                        if (length > 1) {
                            wordPositions.push({ row: r, col: c, length, direction: 'vertical' });
                            wordCount++;
                        }
                    }
                    // Validate that the starting number matches the number of words that can start
                    if (wordCount !== num) {
                        throw new Error(`Starting number ${num} at (${r},${c}) does not match the number of words that can start (${wordCount})`);
                    }
                }
            }
        }

        // Check for mismatch between number of input words and puzzle starting cells
        if (wordPositions.length !== words.length) {
            throw new Error(`Mismatch between number of words (${words.length}) and starting positions (${wordPositions.length})`);
        }

        // Check for word repetition
        if (new Set(words).size !== words.length) {
            throw new Error('Duplicate words are not allowed');
        }

        // Check for wrong word format
        if (words.some(word => !/^[a-zA-Z]+$/.test(word))) {
            throw new Error('Words must contain only letters');
        }

        const usedWords = new Set();
        let solutionsFound = 0;
        let solution = null;

        function canPlaceWord(word, row, col, length, direction) {
            if (word.length !== length) return false;
            for (let i = 0; i < length; i++) {
                const r = direction === 'horizontal' ? row : row + i;
                const c = direction === 'horizontal' ? col + i : col;
                if (r >= rows || c >= cols) return false;
                const cellContent = grid[r][c];
                if (cellContent !== '.' && cellContent !== '0' && cellContent !== word[i] && isNaN(parseInt(cellContent))) {
                    return false;
                }
            }
            return true;
        }

        function placeWord(word, row, col, direction) {
            for (let i = 0; i < word.length; i++) {
                const r = direction === 'horizontal' ? row : row + i;
                const c = direction === 'horizontal' ? col + i : col;
                grid[r][c] = word[i];
            }
        }

        function removeWord(word, row, col, direction) {
            for (let i = 0; i < word.length; i++) {
                const r = direction === 'horizontal' ? row : row + i;
                const c = direction === 'horizontal' ? col + i : col;
                if (isNaN(parseInt(grid[r][c]))) {
                    grid[r][c] = '.';
                }
            }
        }

        function solve(posIndex) {
            if (posIndex === wordPositions.length) {
                solutionsFound++;
                if (solutionsFound > 1) {
                    return false; // Multiple solutions found
                }
                return true;
            }

            const { row, col, length, direction } = wordPositions[posIndex];

            for (const word of words) {
                if (!usedWords.has(word) && canPlaceWord(word, row, col, length, direction)) {
                    placeWord(word, row, col, direction);
                    usedWords.add(word);

                    if (!solve(posIndex + 1)) {
                        removeWord(word, row, col, direction);
                        usedWords.delete(word);
                        return false; // Propagate multiple solutions signal
                    }

                    if (solutionsFound > 1) {
                        return false; // Multiple solutions found
                    }

                    removeWord(word, row, col, direction);
                    usedWords.delete(word);
                }
            }

            return true;
        }

        solve(0);

        if (solutionsFound === 0) {
            console.log('Error: No solution found');
        } else if (solutionsFound > 1) {
            console.log('Error: Multiple solutions found');
        } else {
            console.log(grid.map(row => row.join('')).join('\n'));
        }
    } catch (error) {
        console.log('Error:', error.message);
    }
}

// Example usage

const puzzle = '2000\n0...\n0...\n0...'
const words = ['abba', 'assa']


crosswordSolver(puzzle, words)
