package day2

import (
	"log"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day2", "day-2-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	totalScore := 0
	for _, line := range lines {
		opp := runeToPlay(rune(line[0]))
		play := runeToPlay(rune(line[2]))
		totalScore += score(play, opp)
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 2, Part 1 (%dms): Score = %d", end-start, totalScore)
	return totalScore
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	totalScore := 0
	for _, line := range lines {
		opp := runeToPlay(rune(line[0]))
		outcome := runeToOutcome(rune(line[2]))
		play := playForOutcome(opp, outcome)
		totalScore += score(play, opp)
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 2, Part 2 (%dms): Score = %d", end-start, totalScore)
	return totalScore
}

type Play int
type Outcome int

const (
	Rock     Play = 1
	Paper    Play = 2
	Scissors Play = 3

	Win  Outcome = 6
	Tie  Outcome = 3
	Loss Outcome = 0
)

func score(play Play, opp Play) int {
	switch {
	case play == Rock && opp == Paper:
		return int(Loss) + int(Rock)
	case play == Rock && opp == Scissors:
		return int(Win) + int(Rock)
	case play == Paper && opp == Rock:
		return int(Win) + int(Paper)
	case play == Paper && opp == Scissors:
		return int(Loss) + int(Paper)
	case play == Scissors && opp == Rock:
		return int(Loss) + int(Scissors)
	case play == Scissors && opp == Paper:
		return int(Win) + int(Scissors)
	}
	return int(Tie) + int(play)
}

func runeToPlay(r rune) Play {
	switch r {
	case 'A', 'X':
		return Rock
	case 'B', 'Y':
		return Paper
	default:
		return Scissors
	}
}

func runeToOutcome(r rune) Outcome {
	switch r {
	case 'X':
		return Loss
	case 'Y':
		return Tie
	default:
		return Win
	}
}

func playForOutcome(opp Play, outcome Outcome) Play {
	switch {
	case outcome == Win && opp == Rock:
		return Paper
	case outcome == Win && opp == Paper:
		return Scissors
	case outcome == Win && opp == Scissors:
		return Rock
	case outcome == Loss && opp == Rock:
		return Scissors
	case outcome == Loss && opp == Paper:
		return Rock
	case outcome == Loss && opp == Scissors:
		return Paper
	}
	return opp
}
