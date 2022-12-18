package day15

import (
	"log"
	"regexp"
	"time"

	"advent-of-code-2022/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day15", "day-15-input.txt")
	solvePart1(input, 2000000)
	solvePart2(input, 4000000)
}

func solvePart1(lines []string, row int) int {
	start := time.Now().UnixMilli()
	var sensors []Sensor
	for _, line := range lines {
		sensors = append(sensors, *NewSensor(line))
	}

	var ranges []Range
	for _, s := range sensors {
		ranges = MarkNoBeacon(s, row, 0, ranges)
	}
	for _, s := range sensors {
		ranges = MarkBeacon(s.beacon, row, ranges)
	}
	ans := NoBeaconPoints(ranges)
	end := time.Now().UnixMilli()
	log.Printf("Day 15, Part 1 (%dms): Beacon Not Present = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string, max int) int {
	start := time.Now().UnixMilli()
	var sensors []Sensor
	for _, line := range lines {
		sensors = append(sensors, *NewSensor(line))
	}

	rows := make([][]Range, max+1)
	for row := 0; row <= max; row++ {
		for _, s := range sensors {
			rows[row] = MarkNoBeacon(s, row, max, rows[row])
		}
	}

	var beacon *Point
	for row := 0; row <= max; row++ {
		beacon = PossibleBeacon(rows[row], row, max, sensors)
		if beacon != nil {
			break
		}
	}

	ans := 0
	if beacon != nil {
		ans = (4000000 * beacon.x) + beacon.y
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 15, Part 2 (%dms): Beacon Frequency = %d", end-start, ans)
	return ans
}

type Range struct {
	start int
	end   int
}

func PossibleBeacon(ranges []Range, row int, max int, sensors []Sensor) *Point {
	nRanges := ranges
	if nRanges[0].start != 0 {
		p := Point{0, row}
		if !DiscoveredBeacon(p, sensors) {
			return &p
		}
	}
	x := nRanges[0].end
	for i := 1; i < len(nRanges); i++ {
		if nRanges[i].start > x+1 {
			p := Point{x + 1, row}
			if !DiscoveredBeacon(p, sensors) {
				return &p
			}
		}
		x = nRanges[i].end
	}
	if x+1 <= max {
		p := Point{x + 1, row}
		if !DiscoveredBeacon(p, sensors) {
			return &p
		}
	}
	return nil
}

func DiscoveredBeacon(p Point, sensors []Sensor) bool {
	for _, s := range sensors {
		if s.beacon == p {
			return true
		}
	}
	return false
}

func NoBeaconPoints(ranges []Range) int {
	points := 0
	for _, r := range ranges {
		points += r.end - r.start + 1
	}
	return points
}

func MarkBeacon(p Point, trackRow int, ranges []Range) []Range {
	if p.y != trackRow {
		return ranges
	}
	nRanges := ranges
	for i := 0; i < len(ranges); i++ {
		r := ranges[i]
		if p.x < r.start || p.x > r.end {
			continue
		}
		switch {
		case p.x == r.start:
			nRanges[i].start++
		case p.x == r.end:
			nRanges[i].end--
		default:
			nRanges = append(nRanges[:i+1], nRanges[i:]...)
			nRanges[i] = Range{r.start, p.x - 1}
			nRanges[i+1] = Range{p.x + 1, r.end}
		}
	}
	return nRanges
}

func MarkNoBeacon(s Sensor, trackRow int, max int, ranges []Range) []Range {
	distance := s.point.ManhattanDistance(s.beacon)
	x := 0
	switch {
	case s.point.y < trackRow:
		x = distance - (trackRow - s.point.y)
	case s.point.y >= trackRow:
		x = distance - (s.point.y - trackRow)
	}
	if x < 0 {
		return ranges
	}
	return recordRange(ranges, max, s.point.x-x, s.point.x+x)
}

func recordRange(ranges []Range, max int, start int, end int) []Range {
	if max > 0 {
		if start > max || end < 0 {
			return ranges
		}
		if start < 0 {
			start = 0
		}
		if end > max {
			end = max
		}
	}

	nRanges := ranges
	added := false
	for i := 0; i < len(ranges); i++ {
		r := ranges[i]
		switch {
		case end < r.start:
			nRanges = append(nRanges[:i+1], nRanges[i:]...)
			nRanges[i] = Range{start, end}
			added = true
		case start <= r.start && end >= r.start && end <= r.end:
			nRanges[i] = Range{start, r.end}
			added = true
		case start <= r.start && end >= r.start && end > r.end:
			e := end
			j := i + 1
			for ; j < len(nRanges); j++ {
				if end < nRanges[j].start {
					break
				}
				if nRanges[j].end > e {
					e = nRanges[j].end
				}
			}
			nRanges = append(nRanges[:i+1], nRanges[j:]...)
			nRanges[i] = Range{start, e}
			added = true
		case start >= r.start && start <= r.end && end <= r.end:
			added = true
		case start >= r.start && start <= r.end && end > r.end:
			e := end
			j := i + 1
			for ; j < len(nRanges); j++ {
				if end < nRanges[j].start {
					break
				}
				if nRanges[j].end > e {
					e = nRanges[j].end
				}
			}
			nRanges = append(nRanges[:i+1], nRanges[j:]...)
			nRanges[i] = Range{r.start, e}
			added = true
		}
		if added {
			break
		}
	}
	if !added {
		nRanges = append(nRanges, Range{start, end})
	}
	return nRanges
}

type Point struct {
	x int
	y int
}

func (p Point) ManhattanDistance(p2 Point) int {
	xDistance := p.x - p2.x
	if xDistance < 0 {
		xDistance *= -1
	}
	yDistance := p.y - p2.y
	if yDistance < 0 {
		yDistance *= -1
	}
	return xDistance + yDistance
}

type Sensor struct {
	point  Point
	beacon Point
}

var sensorParse = regexp.MustCompile(`Sensor at x=([-\d]+), y=([-\d]+): closest beacon is at x=([-\d]+), y=([-\d]+)`)

func NewSensor(s string) *Sensor {
	match := sensorParse.FindStringSubmatch(s)
	if match == nil {
		return nil
	}
	return &Sensor{
		point:  Point{utils.Number(match[1]), utils.Number(match[2])},
		beacon: Point{utils.Number(match[3]), utils.Number(match[4])},
	}
}
