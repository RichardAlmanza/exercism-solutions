package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type result map[string]int
type Tournament struct {
	Teams []string
	Results map[string]result
}

var (
	ErrBrokenLine = errors.New("the line does not have 3 sections")
	ErrEmptyInfo = errors.New("there is no information")
	ErrWrongState = errors.New("the state is neither draw, loss, or win")

	States = map[string]int{
		"loss": 0,
		"draw": 1,
		"win": 3,
	}

	resultKeys = []string{
		"matches",
		"win",
		"draw",
		"loss",
		"points",
	}
)

func newResult() result {
	return result{
		"matches": 0, //"MatchesPlayed": 0,
		"win": 0, //"MatchesWon": 0,
		"loss": 0, //"MatchesDrawn": 0,
		"draw": 0, //"MatchesLost": 0,
		"points": 0, //"Points": 0,
	}
}

func (r result) String() string {
	var output strings.Builder

	for _, key := range resultKeys {
		output.WriteString(" | ")
		output.WriteString(fmt.Sprintf("%2s", strconv.Itoa(r[key])))
	}

	return output.String()
}

func NewTournament() Tournament {
	return Tournament{
		Teams: make([]string, 0, 10),
		Results: make(map[string]result),
	}
}

func (t *Tournament) createTeam(team string) {
	t.Teams = append(t.Teams, team)
	t.Results[team] = newResult()
}

func (t *Tournament) updateResult(team, state string) {
	_, exists := t.Results[team]

	if !exists {
		t.createTeam(team)
	}

	t.Results[team]["matches"]++
	t.Results[team][state]++
	t.Results[team]["points"] += States[state]
}

func (t *Tournament) AddRecord(record []string) {
	var team1 string = record[0]
	var team2 string = record[1]
	var state string = record[2]

	t.updateResult(team1, state)
	t.updateResult(team2, negativeState(state))
}

func (t Tournament) String() string {
	var output strings.Builder

	sort.Slice(t.Teams, func(i, j int) bool {
		var score1 int = t.Results[t.Teams[i]]["points"]
		var score2 int = t.Results[t.Teams[j]]["points"]

		if score1 == score2 {
			return t.Teams[i] < t.Teams[j]
		}

		return score1 > score2
	})

	output.WriteString("Team                           | MP |  W |  D |  L |  P\n")

	for _, team := range t.Teams {
		output.WriteString(fmt.Sprintf("%-30s", team))
		output.WriteString(t.Results[team].String())
		output.WriteRune('\n')
	}

	return output.String()
}

func negativeState(state string) string {
	var output string

	switch state {
	case "loss":
		output = "win"
	case "win":
		output = "loss"
	default:
		output = state
	}

	return output
}

func parseLine(line string) ([]string, error) {
	var elements []string = strings.Split(line, ";")

	if len(elements) != 3 {
		return nil, ErrBrokenLine
	}

	if _, exists := States[elements[2]]; !exists {
		return nil, ErrWrongState
	}

	return elements, nil
}

func ignoreLine(line string) bool {
	switch {
	case line == "",
		line[0] == '#':
		return true
	}

	return false
}

func Tally(reader io.Reader, writer io.Writer) error {
	var scanner *bufio.Scanner = bufio.NewScanner(reader)
	var t = NewTournament()

	for scanner.Scan() {
		line := scanner.Text()

		if ignoreLine(line) {
			continue
		}

		elements, err := parseLine(line)

		if err != nil {
			return err
		}

		t.AddRecord(elements)
	}

	if len(t.Teams) == 0 {
		return ErrEmptyInfo
	}

	writer.Write([]byte(t.String()))

	return nil
}
