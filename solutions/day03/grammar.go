package day03

import (
	"math"
)

type EngineSchematic struct {
	EngineLines []*EngineLine `@@*`
}

type EngineLine struct {
	Elements []*Element `@@* Newline`
}

type Element struct {
	Value    *Value `  @@`
	Dot      string `| @Dot`
	Symbol   string `| @Symbol`
	position *position
}

type Value struct {
	Number int `  @Number`
	length int
}

type position struct {
	x      int
	y      int
	width  int
	height int
}

func (es *EngineSchematic) CalculateValueLengths() {
	for _, el := range es.EngineLines {
		for _, e := range el.Elements {
			if e.Value != nil {
				e.Value.length = int(math.Floor(math.Log10(float64(e.Value.Number))) + 1)
			}
		}
	}
}

func (es *EngineSchematic) CalculatePositions() {
	for x, el := range es.EngineLines {
		y := 0
		for _, e := range el.Elements {
			if e.Value != nil {
				e.position = &position{
					x:      x,
					y:      y,
					width:  e.Value.length,
					height: 1,
				}
				y += e.Value.length
			} else {
				e.position = &position{
					x:      x,
					y:      y,
					width:  1,
					height: 1,
				}
				y++
			}
		}
	}
}

func (p1 *position) isAdjacent(p2 *position) bool {
	isOverlap1D := func(min1, max1, min2, max2 int) bool {
		return max1 >= min2 && max2 >= min1
	}
	return isOverlap1D(p1.x, p1.x+p1.height, p2.x, p2.x+p2.height) && isOverlap1D(p1.y, p1.y+p1.width, p2.y, p2.y+p2.width)
}

func (es *EngineSchematic) GetAdjacentElements(e *Element) []*Element {
	adjacentElements := make([]*Element, 0)
	searchRows := make([]*EngineLine, 0)

	if e.position.x == 0 {
		searchRows = append(searchRows, es.EngineLines[e.position.x], es.EngineLines[e.position.x+1])
	} else if e.position.x == len(es.EngineLines)-1 {
		searchRows = append(searchRows, es.EngineLines[e.position.x-1], es.EngineLines[e.position.x])
	} else {
		searchRows = append(searchRows, es.EngineLines[e.position.x-1], es.EngineLines[e.position.x], es.EngineLines[e.position.x+1])
	}

	for _, row := range searchRows {
		for _, el := range row.Elements {
			if e.position.isAdjacent(el.position) {
				adjacentElements = append(adjacentElements, el)
			}
		}
	}
	return adjacentElements
}
