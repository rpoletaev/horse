package main

import "fmt"
import "bufio"
import "os"

const (
	minLetter = byte(97)
	maxLetter = byte(104)
	minNum    = byte(49)
	maxNum    = byte(56)
)

type point struct {
	X byte
	Y byte
}

func (p point) String() string {
	bts := []byte{p.X, p.Y}
	return fmt.Sprintf("%s", bts)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			input := scanner.Text()
			avPositions, err := getAvailablePositions(input)
			if err != nil {
				println(err)
				continue
			}

			fmt.Printf("%q\n", avPositions)
		}
	}
}

func getAvailablePositions(startPosition string) ([]string, error) {
	if len(startPosition) > 2 || len(startPosition) < 2 {
		return nil, fmt.Errorf("Wrong start position format %s", startPosition)
	}
	if startPosition[0] < minLetter || startPosition[0] > maxLetter {
		return nil, fmt.Errorf("Wrong position letter")
	}

	if startPosition[1] < minNum || startPosition[1] > maxNum {
		return nil, fmt.Errorf("Wrong position number")
	}

	p := point{startPosition[0], startPosition[1]}
	var result []string

	//try go to left
	if lp, err := toLeft(2, p); err == nil {
		if fp, inerr := toForward(1, lp); inerr == nil {
			result = append(result, fp.String())
		}

		if bp, inerr := toBack(1, lp); inerr == nil {
			result = append(result, bp.String())
		}
	}

	//try go to forward
	if fp, err := toForward(2, p); err == nil {
		if lp, inerr := toLeft(1, fp); inerr == nil {
			result = append(result, lp.String())
		}

		if rp, inerr := toRight(1, fp); inerr == nil {
			result = append(result, rp.String())
		}
	}

	//try go to right
	if rp, err := toRight(2, p); err == nil {
		if fp, inerr := toForward(1, rp); inerr == nil {
			result = append(result, fp.String())
		}

		if bp, inerr := toBack(1, rp); inerr == nil {
			result = append(result, bp.String())
		}
	}

	//try go to backward
	if bp, err := toBack(2, p); err == nil {
		if lp, inerr := toLeft(1, bp); inerr == nil {
			result = append(result, lp.String())
		}

		if rp, inerr := toRight(1, bp); inerr == nil {
			result = append(result, rp.String())
		}
	}

	return result, nil
}

func toLeft(stepCount int, position point) (point, error) {
	byteCnt := byte(stepCount)
	if (position.X-byteCnt < minLetter) || (position.X-byteCnt > maxLetter) {
		return position, fmt.Errorf("WrongPosition")
	}

	position.X -= byteCnt
	return position, nil
}

func toRight(stepCount int, position point) (point, error) {
	byteCnt := byte(stepCount)
	if (position.X+byteCnt < minLetter) || (position.X+byteCnt > maxLetter) {
		return position, fmt.Errorf("WrongPosition")
	}

	position.X += byteCnt
	return position, nil
}

func toBack(stepCount int, position point) (point, error) {
	byteCnt := byte(stepCount)
	if (position.Y-byteCnt < minNum) || (position.Y-byteCnt > maxNum) {
		return position, fmt.Errorf("WrongPosition")
	}

	position.Y -= byteCnt
	return position, nil
}

func toForward(stepCount int, position point) (point, error) {
	byteCnt := byte(stepCount)
	if (position.Y+byteCnt < minNum) || (position.Y+byteCnt > maxNum) {
		return position, fmt.Errorf("WrongPosition")
	}

	position.Y += byteCnt
	return position, nil
}
