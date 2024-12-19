package day08

import (
	"advent-of-code/utils"
	"fmt"
	"maps"
)

type Position struct {
	i, j int
}

func SolvePart01(inputPath string) {
	antennaGrid := utils.ReadFileTo2D(inputPath, "")
	signalGrid := make(map[Position]struct{})

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
						signalGrid[signalPos] = struct{}{}
					}

				}
			}
		}
	}

	signalPositions := maps.Keys(signalGrid)
	count := utils.IterLength(signalPositions)
	fmt.Printf("Part 01: %v (number of antinodes)\n", count)
}

func SolvePart02(inputPath string) {
	antennaGrid := utils.ReadFileTo2D(inputPath, "")
	signalGrid := make(map[Position]struct{})

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

					signalPos := Position{mi, mj} // The first signal is on the antenna itself
					di, dj := mi-i, mj-j
					n := 0
					for !utils.IsOutOfBounds2D(signalPos.i, signalPos.j, iMax, jMax) {
						signalGrid[signalPos] = struct{}{}
						signalPos.i = i - n*di
						signalPos.j = j - n*dj
						n++
					}
				}
			}
		}
	}

	signalPositions := maps.Keys(signalGrid)
	count := utils.IterLength(signalPositions)
	fmt.Printf("Part 02: %v (number of antinodes considering harmonics)\n", count)
}
