package main

func calculateNextState(p golParams, world [][]byte) [][]byte {
	for j := 0; j < p.imageHeight; j++ {
		for i := 0; i < p.imageWidth; i++ {
			sum := 0
			if world[i%15][j%15] == 255 {
				sum += 1
			}
			if world[(i+1)%15][j%15] == 255 {
				sum += 1
			}
			if world[(i-1)%15][j%15] == 255 {
				sum += 1
			}
			if world[(i-1)%15][(j+1)%15] == 255 {
				sum += 1
			}
			if world[(i-1)%15][(j-1)%15] == 255 {
				sum += 1
			}
			if world[(i+1)%15][(j+1)%15] == 255 {
				sum += 1
			}
			if world[(i+1)%15][(j-1)%15] == 255 {
				sum += 1
			}
		}
	}
	return world
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	return []cell{}
}
