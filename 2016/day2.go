package main

import (
	"fmt"
	"strings"
)

type coord struct {
	X, Y int
}

func main() {
	input := strings.Split(`RRLLDDRLLDURLDUUDULDLRLDDDRLDULLRDDRDLUUDLLRRDURRLUDUDULLUUDRUURRDDDDURUULLDULRLRRRDLLLRDLRDULDLRUUDRURLULURUUDRLUUDRDURUUDDDRDLLRDLUDRLUUUUUULDURDRDDURLDDRLUUDLRURRDRLDRDLRRRLURULDLUUURDRLUULRDUDLDRRRDDLDDDRLRLLDRDUURDULUURRRRUDLLUDLDRLLRRULLLRDRDRRDRDRDUULUDLULRLLDRULURLURDLLDDRRLUDRDUULLDRULLLRLULUDDLURLUULDRUDRLDUUDDLLLRURDRLDRLUUUUUUDRUDLDLULULRRURRDDDUDRRRDDDLDDLRLLDDUULLUDRURDDDRDDLURRURULULUUDRLLUUDDDRUULRDLDLLRUUDRRLRRRULLRRURUDDUUDULDUDUUUDURUDUUDUDRULUDULRDUDUUUUDLLURDLRRUDURDDUULLDLLRDUDULRLRDURLRDRDLRDLRURUDURLULDDDLDRLULLRLURRLLDLRLLULLDUURUUDRULDDUDLDDR
LUURULURUURRRDLRDRDDDUDULRDDLUUDUUULRLRRLRUDDLDLURLRULUUDUUDLDRLLUURLUUURDUDLUULLUUDUUDRDUDUDURLLURDUDLDLDDLDUDRLLUULDDRUDDDRLRUDRDUDULLRRDLLDDLRLLLRLRURURLLDULUDDUULULDRRLUURDULRULUDULDULDULRULLLRRDRLDRLDUULLDDULRLUDLLLULDDLULLUULUURRULDLUULDRRUDDDLRDLULDULDRRDLRRRLRUDURUDDDUDDURRRLDUULRDDLRRRLRUULDDURDRDUULDLLULULDRDRUULDLULRUUDUDLUDRLRDURRRRLULURDRLLLUDRRRDRURRLRLLUURDLURLURDULURUDDULLDUUDDLRLUULRDUUDRDRUDRLUUUDURLLRDRRDRURDDDDULLDDUDLDUUDLRLURURLDRRDRDUDLRRDRUDRDLURRLLLULULULRUDDDULULULDDRRLULUUUURDDURURLDLDDDDDRULUUUULLUDDDRRLUULDUULRUUDUURRLLRDDRUURL
RRRLLLLUULLRRLLDRULULLLRDLLDLDDLURUDLULUULLRURLDULLUDULDRURDURLULLDUDDRLLRUURDLLULUURLULLDLRRDDDULUURRUDRDRDURULDLLURUDLLLDDUDLLLLRLDRDRDDRRDLUUDLLLULLLLLDDRDLULLLLRURRRUUUULLDLRDLDLRRRULRRRRLDLLRDURULDDLURURUULULDRDDDURLRDDRDULLUURUDUUUDRDRRLURULULRLUUDDRDULDRLULULUULRLDRLUDRRDDDRUDDRDDRDDRRLRDLRURDULULRRUUURDRRRDURDDRUDULUUDRDDLDRDDDULDLRDUULDUULRUDLRRDDDLLDDLLLRRDLDUULUULULURRULLRRUDUDRUDRRRLDLRLURRLUDLLLUUDDUDRURUUDDURDURULRLDUDRDLULDUDRUDDDR
DRDRRUUUUURRLUDLDLRUUULRLDLRRRDDUDLUUDUDDLRDUDLRRLLURUUDULLUDLULLDLLDDULUDDUDUULURLDLDDUUDDRDDRLUURLUUDUDUDURULLDRLDDRUDLURRRLDRLRULDDLDDLDDDULDUDDLLDULUUDUDDUULDRLDRLRURDULUDDRRLRRLRRDULDRRDUDUDRLDURRLRLRDLRLRRLRDDDRULLULULDUDDLDLULRLLURRRRULUULRUDLDLRDLLURURUUURDRRRLDDRLRLURDDURDDUURUUUDDLRUURRRDLLUULUDRLDDRDDDDUUDRLRRDULDULRDLLLRULULLRDRULLRDLRUUURRRURLRRDLDRDDLURLDLULLDUURRDULUUURRLLDDUUUULDDDDLRDDDRDLDLLUURLDDRULUDDRDDDLRDU
DRRRLURUDUDUULDLLURLUUDDRRRDUDLURLLDRRLLDDURULUDUURURULLRLDLLUURDLLDLLDLDLRUDLLLLDRLLUDLLDULDDRRURDRRLRLRRUURRUDURRLDRDLDURUULUDRLLURLUDURDULLRLLDLURLRRDLLLLUUDRDULLDLURDULDRDURRRLDRLURULULURLLLRRRUULRLRRDRDDDLULRLRRDLUDDUUUUUDRRDLDUDUURLDRRRLRUDRULDRLURUULRDLDDLRURDLURULRRLDURLDUURURULRDUDRRUDUDDLRLUURURULDDLULULULDULRRLRRURUURLRLLDRRLUDLUURDRRURRUUDULRLURRLRLRDDRURDDLRRDUDRLLUUUDULRDRULUURRLRLRDUDULDRDLLUDRDLLDRULDLUUURDDRDDUDDULLLDRRDDUDDDDLDDRLRULRRRURRDRULLDDDURDDLURRDDDUDLURRUDUDLLDDDLRUUDRLRRLRDUUUDDL`, "\n")
	keypadA := map[coord]rune{
		coord{-1, -1}: '7',
		coord{0, -1}:  '8',
		coord{1, -1}:  '9',
		coord{-1, 0}:  '4',
		coord{0, 0}:   '5',
		coord{1, 0}:   '6',
		coord{-1, 1}:  '1',
		coord{0, 1}:   '2',
		coord{1, 1}:   '3',
	}
	fmt.Println(keypadA)

	keypadB := map[coord]rune{
		coord{0, 0}:  '5',
		coord{1, 0}:  '6',
		coord{2, 0}:  '7',
		coord{3, 0}:  '8',
		coord{4, 0}:  '9',
		coord{1, 1}:  '2',
		coord{2, 1}:  '3',
		coord{3, 1}:  '4',
		coord{2, 2}:  '1',
		coord{1, -1}: 'A',
		coord{2, -1}: 'B',
		coord{3, -1}: 'C',
		coord{2, -2}: 'D',
	}
	fmt.Println(keypadB)

	for _, key := range input {
		loc := coord{0, 0}
		for _, inst := range key {
			proposedLoc := loc

			switch inst {
			case 'U':
				proposedLoc.Y += 1
			case 'R':
				proposedLoc.X += 1
			case 'D':
				proposedLoc.Y -= 1
			case 'L':
				proposedLoc.X -= 1
			default:
				fmt.Printf("Invalid direction: %c", inst)
				return
			}
			if keypadB[proposedLoc] != rune(0) {
				loc = proposedLoc
			}
		}
		fmt.Printf("%c", keypadB[loc])
	}
	fmt.Println()
}
