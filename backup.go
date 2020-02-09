package main

func (board *Board) createLines() []Line {
	var lines []Line

	for x := 0; x < SIZE; x++ {
		var line Line
		for y := 0; y < SIZE; y++ {
			line = append(line, &board[y][x])
		}
		lines = append(lines, line)
	}

	for y := 0; y < SIZE; y++ {
		var line Line
		for x := 0; x < SIZE; x++ {
			line = append(line, &board[y][x])
		}
		lines = append(lines, line)
	}

	for i := 0; i < SIZE; i++ {
		var line Line
		for j := 0; i+j < SIZE; j++ {
			line = append(line, &board[j][i+j])
		}
		lines = append(lines, line)
	}

	for i := 1; i < SIZE; i++ {
		var line Line
		for j := 0; i+j < SIZE; j++ {
			line = append(line, &board[i+j][j])
		}
		lines = append(lines, line)
	}

	for i := 0; i < SIZE; i++ {
		var line Line
		for j := 0; i-j > -1; j++ {
			line = append(line, &board[j][i-j])
		}
		lines = append(lines, line)
	}

	for i := 1; i < SIZE; i++ {
		var line Line
		for j := 0; i+j < SIZE; j++ {
			line = append(line, &board[i+j][SIZE-1-j])
		}
		lines = append(lines, line)
	}

	return lines
}
