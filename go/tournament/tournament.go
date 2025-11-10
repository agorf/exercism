package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"
)

type tallyRow struct {
	teamName      string
	matchesPlayed int
	matchesWon    int
	matchesDrawn  int
	matchesLost   int
	points        int
}

func Tally(reader io.Reader, writer io.Writer) error {
	tally := make(map[string]*tallyRow)
	br := bufio.NewReader(reader)
	for {
		line, readErr := br.ReadString('\n')
		if readErr != nil && readErr != io.EOF {
			return readErr
		}
		line = strings.TrimSpace(line) // For trailing newline
		if line == "" {                // Ignore empty lines
			if readErr == io.EOF {
				break
			}
			continue
		}
		if strings.HasPrefix(line, "#") { // Ignore comments
			if readErr == io.EOF {
				break
			}
			continue
		}
		match := strings.SplitN(line, ";", 3)
		if len(match) != 3 {
			return errors.New("invalid line: " + line)
		}
		teamA, teamB, result := match[0], match[1], match[2]
		if _, ok := tally[teamA]; !ok {
			tally[teamA] = &tallyRow{teamName: teamA} // Allocate
		}
		if _, ok := tally[teamB]; !ok {
			tally[teamB] = &tallyRow{teamName: teamB} // Allocate
		}
		tally[teamA].matchesPlayed++
		tally[teamB].matchesPlayed++
		switch result {
		case "win":
			tally[teamA].matchesWon++
			tally[teamB].matchesLost++
			tally[teamA].points += 3
		case "draw":
			tally[teamA].matchesDrawn++
			tally[teamB].matchesDrawn++
			tally[teamA].points++
			tally[teamB].points++
		case "loss":
			tally[teamA].matchesLost++
			tally[teamB].matchesWon++
			tally[teamB].points += 3
		default:
			return errors.New("invalid match result: " + result)
		}
		if readErr == io.EOF {
			break
		}
	}
	var tallyRows []*tallyRow
	for _, tallyRow := range tally {
		tallyRows = append(tallyRows, tallyRow)
	}
	slices.SortFunc(tallyRows, func(a, b *tallyRow) int {
		switch {
		case a.points < b.points:
			return 1 // Sort descending
		case a.points > b.points:
			return -1 // Sort descending
		case a.points == b.points:
			switch {
			case a.teamName < b.teamName:
				return -1
			case a.teamName > b.teamName:
				return 1
			case a.teamName == b.teamName:
				return 0
			}
		}
		panic("should never happen")
	})
	_, err := fmt.Fprintf(writer, "%-30s | MP |  W |  D |  L |  P\n", "Team")
	if err != nil {
		return err
	}
	for _, tallyRow := range tallyRows {
		_, err := fmt.Fprintf(
			writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			tallyRow.teamName,
			tallyRow.matchesPlayed,
			tallyRow.matchesWon,
			tallyRow.matchesDrawn,
			tallyRow.matchesLost,
			tallyRow.points,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
