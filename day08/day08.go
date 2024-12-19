package day08

import (
	"advent-of-code/utils"
	"fmt"
	"maps"
)

type Position struct {
	i, j int
}

func Solve(inputPath string) {
	antennaGrid := utils.ReadFileTo2D(inputPath, "")
	signalGrid := make(map[Position][]string)

	iMax := len(antennaGrid) - 1
	jMax := len(antennaGrid[0]) - 1

	// Find tower
	for i, row := range antennaGrid {
		for j, antenna := range row {
			if antenna == "." {
				continue
			}

			// Find matching towers
			for mi, mrow := range antennaGrid {
				for mj, mantenna := range mrow {
					if antenna == "." || antenna != mantenna || (mi == i && mj == j) {
						continue
					}

					di, dj := mi-i, mj-j
					signalPos := Position{i - di, j - dj}

					if !utils.IsOutOfBounds2D(signalPos.i, signalPos.j, iMax, jMax) {
						signalGrid[signalPos] = append(signalGrid[signalPos], antenna)
					}

				}
			}
		}
	}

	signalPositions := maps.Keys(signalGrid)
	count := utils.IterLength(signalPositions)
	fmt.Printf("Part 01: %v (number of antinodes)\n", count)

}
