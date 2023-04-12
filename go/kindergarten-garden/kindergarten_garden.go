package kindergarten

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

type Garden struct{
	garden map[string][]string
}

var (
	ErrDiagram error = errors.New("diagram does not follow the defined structure")
	ErrChildren error = errors.New("children are repeated or does not match with the diagram")

	reDiagram *regexp.Regexp = regexp.MustCompile(`^\n([VGRC]{2})+\n([VGRC]{2})+$`)

	plants = map[rune]string{
		'V': "violets",
		'G': "grass",
		'R': "radishes",
		'C': "clover",
	}
)

func isDiagramValid(diagram string) bool {
	if !reDiagram.MatchString(diagram) {
		return false
	}

	rows := strings.Split(diagram, "\n")[1:]

	return len(rows[0]) == len(rows[1])
}

func areChildrenValid(diagram string, children []string) bool {
	var unique map[string]bool = make(map[string]bool)

	for _, child := range children {
		if _, exists := unique[child]; exists {
			return false
		}

		unique[child] = true
	}

	return 4 * len(children) + 2 == len(diagram)
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	if !isDiagramValid(diagram) {
		return nil, ErrDiagram
	}

	if !areChildrenValid(diagram, children) {
		return nil, ErrChildren
	}

	var isSecondRow bool
	var counter int
	var childIndex int
	var sortedChildren []string = make([]string, 0, len(children))
	var newGarden Garden = Garden{make(map[string][]string)}

	sortedChildren = append(sortedChildren, children...)
	sort.Strings(sortedChildren)

	for _, plant := range diagram[1:] {
		if !isSecondRow && plant == '\n' {
			isSecondRow = true
			childIndex = 0
			counter = 0
			continue
		}

		if counter == 2 {
			counter = 0
			childIndex++
		}

		counter++
		newGarden.garden[sortedChildren[childIndex]] =  append(newGarden.garden[sortedChildren[childIndex]], plants[plant])
	}

	return &newGarden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants, exists := g.garden[child]

	return plants, exists
}