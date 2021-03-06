package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day20.input", "Relative file path to use as input.")

type Tile [10][10]bool
type Edge uint16

const (
	TOP = iota
	RIGHT
	BOTTOM
	LEFT
)

func (t Tile) Crop() Cropped {
	var ret Cropped
	for r := 1; r < 9; r++ {
		for c := 1; c < 9; c++ {
			ret[r-1][c-1] = t[r][c]
		}
	}
	return ret
}

func (c Cropped) Count() int {
	count := 0
	for _, row := range c {
		for _, pixel := range row {
			if pixel {
				count++
			}
		}
	}
	return count
}

func (t Tile) FlipX() Tile {
	var ret Tile
	for r, row := range t {
		ret[9-r] = row
	}
	return ret
}

func (t Tile) RotCW() Tile {
	var ret Tile
	for r := range t {
		for c := range t[r] {
			ret[r][c] = t[9-c][r]
		}
	}
	return ret
}

func (t Tile) Edges() [4]Edge {
	// left to right
	var top, bottom Edge
	for i, v := range t[0] {
		if v {
			top |= 1 << i
		}
	}
	for i, v := range t[9] {
		if v {
			bottom |= 1 << i
		}
	}
	// top to bottom
	var left, right Edge
	for i, r := range t {
		if r[0] {
			left |= 1 << i
		}
		if r[9] {
			right |= 1 << i
		}
	}
	// Return in clockwise orientation.
	return [4]Edge{top, right, bottom.Flip(), left.Flip()}
}

func (e Edge) Flip() Edge {
	var ret Edge
	for i := 0; i < 10; i++ {
		ret |= ((e >> i) & 1) << (9 - i)
	}
	return ret
}

type Cropped [8][8]bool

// This is assumed to be pre-rotated/flipped.
type Mosaic [][]Tile

func (m Mosaic) PixelAtCoord(r, c int) bool {
	mosaicRow := r / 8
	mosaicCol := c / 8
	cropped := m[mosaicRow][mosaicCol].Crop()
	row := r % 8
	col := c % 8
	return cropped[row][col]
}

func (m Mosaic) FlipX() Mosaic {
	var ret Mosaic
	for r := range m {
		var row []Tile
		for c := range m[r] {
			row = append(row, m[len(m)-1-r][c].FlipX())
		}
		ret = append(ret, row)
	}
	return ret
}

func (m Mosaic) RotCW() Mosaic {
	var ret Mosaic
	for r := range m {
		var row []Tile
		for c := range m[r] {
			row = append(row, m[len(m[0])-1-c][r].RotCW())
		}
		ret = append(ret, row)
	}
	return ret
}

func (m Mosaic) CheckSingleEdge(t Tile, r, c, dir int) bool {
	oDir := (dir + 2) % 4
	var otherR, otherC int
	switch dir {
	case TOP:
		if r-1 < 0 {
			return true
		}
		otherC = c
		otherR = r - 1
	case BOTTOM:
		if r+1 >= len(m) {
			return true
		}
		otherC = c
		otherR = r + 1
	case LEFT:
		if c-1 < 0 {
			return true
		}
		otherC = c - 1
		otherR = r
	case RIGHT:
		if c+1 >= len(m[0]) {
			return true
		}
		otherC = c + 1
		otherR = r
	}
	otherTile := m[otherR][otherC]
	var empty Tile
	if otherTile == empty {
		return true
	}
	return otherTile.Edges()[oDir] == t.Edges()[dir].Flip()
}

func (m Mosaic) CheckFit(t Tile, r, c int) bool {
	// Check up, if exists.
	if !m.CheckSingleEdge(t, r, c, TOP) {
		return false
	}
	// Check down, if exists.
	if !m.CheckSingleEdge(t, r, c, BOTTOM) {
		return false
	}
	// Check left, if exists.
	if !m.CheckSingleEdge(t, r, c, LEFT) {
		return false
	}
	// Check right, if exists.
	if !m.CheckSingleEdge(t, r, c, RIGHT) {
		return false
	}

	return true
}

func (m Mosaic) Traverse(allTiles map[Tile]int, used map[int]bool, r, c int, dir int) int {
	var rIncr, cIncr int
	switch dir {
	case RIGHT:
		cIncr = 1
	case BOTTOM:
		rIncr = 1
	case LEFT:
		cIncr = -1
	case TOP:
		rIncr = -1
	}
	i := 1
outer:
	for ; ; i++ {
		row := r + rIncr*i
		col := c + cIncr*i
		if row < 0 || row >= len(m) {
			break
		}
		if row < 0 || col >= len(m[0]) {
			break
		}
		for t, s := range allTiles {
			if used[s] {
				// Don't re-use the same piece twice.
				continue
			}
			if m.CheckFit(t, row, col) {
				used[s] = true
				m[row][col] = t
				continue outer
			}
			// This piece hasn't matched, continue on to other pieces.
		}
		// We didn't match any pieces, abort.
		return 0
	}
	return i - 1
}

var SeaMonster = strings.Split(`                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `, "\n")

func (m Mosaic) MonsterTopLeftCoord(r, c int) bool {
	for rOffset, row := range SeaMonster {
		for cOffset, char := range row {
			if char != '#' {
				continue
			}
			if !m.PixelAtCoord(r+rOffset, c+cOffset) {
				return false
			}
		}
	}
	return true
}

func (m Mosaic) FindMonsters() int {
	// Find the sea monster.
	var monstersFound int
	for r := 0; r+len(SeaMonster) < 8*len(m); r++ {
		for c := 0; c+len(SeaMonster[0]) < 8*len(m[0]); c++ {
			if m.MonsterTopLeftCoord(r, c) {
				monstersFound++
			}
		}
	}
	return monstersFound
}

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n")
	split = split[:len(split)-1]
	tiles := make(map[int]Tile)
	for i := 0; i < len(split); i += 12 {
		// First line of metadata describes the tile id.
		// Tile 2477:
		header := split[i]
		n, err := strconv.Atoi(header[5 : len(header)-1])
		if err != nil {
			fmt.Printf("Failed to parse %s\n", header)
			break
		}
		var tile Tile
		for r := 0; r < 10; r++ {
			line := split[i+1+r]
			for c := 0; c < 10; c++ {
				if line[c] == '#' {
					tile[r][c] = true
				}
			}
		}
		tiles[n] = tile
	}
	allEdges := make(map[Edge][]int)
	for i, t := range tiles {
		edges := t.Edges()
		for _, e := range edges {
			f := e.Flip()
			allEdges[e] = append(allEdges[e], i)
			allEdges[f] = append(allEdges[f], i)
		}
	}
	singletonTiles := make(map[int]int)
	for _, v := range allEdges {
		if len(v) < 2 {
			singletonTiles[v[0]]++
		}
	}

	var corners []int
	sides := make(map[int]bool)

	product := 1
	for k, v := range singletonTiles {
		if v <= 2 {
			sides[k] = true
		}
		if v > 2 && v < 5 {
			corners = append(corners, k)
			product *= k
		}
	}
	fmt.Println(product)

	var image Mosaic
	for r := 0; r*r < len(tiles); r++ {
		var line []Tile
		for c := 0; c*c < len(tiles); c++ {
			line = append(line, Tile{})
		}
		image = append(image, line)
	}

	// Fill in the top row.
	var topLeft int
	for i := 0; i < 4; i++ {
		matchingRight := allEdges[tiles[corners[i]].Edges()[RIGHT]]
		matchingDown := allEdges[tiles[corners[i]].Edges()[BOTTOM]]
		if len(matchingRight) < 2 || len(matchingDown) < 2 {
			// This is not the top left corner.
			continue
		}
		topLeft = corners[i]
	}

	image[0][0] = tiles[topLeft]
	used := map[int]bool{topLeft: true}

	// Construct all tile rotations.
	allTiles := make(map[Tile]int)
	for n, t := range tiles {
		for f := 0; f < 2; f++ {
			for r := 0; r < 4; r++ {
				allTiles[t] = n
				t = t.RotCW()
			}
			t = t.FlipX()
		}
	}

	count := image.Traverse(allTiles, used, 0, 0, RIGHT)
	if count != len(image[0])-1 {
		fmt.Println("Failed to traverse right from what should be top left.")
		return
	}

	for c := 0; c < len(image[0]); c++ {
		count = image.Traverse(allTiles, used, 0, c, BOTTOM)
		if count != len(image)-1 {
			fmt.Printf("Failed to traverse down column %d\n", c)
			return
		}
	}

	var monsterCount int
outer:
	for flip := 0; flip < 2; flip++ {
		for rot := 0; rot < 4; rot++ {
			monsterCount = image.FindMonsters()
			if monsterCount != 0 {
				break outer
			}
			image = image.RotCW()
		}
		image = image.FlipX()
	}
	if monsterCount == 0 {
		fmt.Println("Failed to detect any monsters despite rotating and flipping.")
	}

	monsterPixels := monsterCount * 15
	sum := 0
	for _, t := range tiles {
		sum += t.Crop().Count()
	}
	fmt.Println(sum - monsterPixels)
}
