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
	cave := NewCave()
	for _, s := range sensors {
		cave.MarkNoBeacon(s, row)
	}
	for _, s := range sensors {
		cave.MarkBeacon(s.beacon)
	}
	ans := cave.NoBeaconPoints(row)
	end := time.Now().UnixMilli()
	log.Printf("Day 15, Part 1 (%dms): Beacon Not Present = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string, maxCoord int) int {
	start := time.Now().UnixMilli()
	var sensors []Sensor
	for _, line := range lines {
		sensors = append(sensors, *NewSensor(line))
	}

	cave := NewCave()
	cave.maxCoord = maxCoord
	for row := 0; row <= maxCoord; row++ {
		for _, s := range sensors {
			cave.MarkNoBeacon(s, row)
		}
	}

	var beacon *Point
	for row := 0; row <= maxCoord; row++ {
		beacon = cave.PossibleBeacon(row, maxCoord, sensors)
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

type Cave struct {
	maxCoord       int
	noBeaconRanges map[int][]Range
}

func (c *Cave) PossibleBeacon(row int, maxCoord int, sensors []Sensor) *Point {
	nRanges := c.noBeaconRanges[row]
	if nRanges[0].start != 0 {
		p := Point{0, row}
		if !c.DiscoveredBeacon(p, sensors) {
			return &p
		}
	}
	x := nRanges[0].end
	for i := 1; i < len(nRanges); i++ {
		if nRanges[i].start > x+1 {
			p := Point{x + 1, row}
			if !c.DiscoveredBeacon(p, sensors) {
				return &p
			}
		}
		x = nRanges[i].end
	}
	if x+1 <= maxCoord {
		p := Point{x + 1, row}
		if !c.DiscoveredBeacon(p, sensors) {
			return &p
		}
	}
	return nil
}

func (c *Cave) DiscoveredBeacon(p Point, sensors []Sensor) bool {
	for _, s := range sensors {
		if s.beacon == p {
			return true
		}
	}
	return false
}

func (c *Cave) NoBeaconPoints(row int) int {
	points := 0
	for _, r := range c.noBeaconRanges[row] {
		points += r.end - r.start + 1
	}
	return points
}

func (c *Cave) MarkBeacon(p Point) {
	nRanges := c.noBeaconRanges[p.y]
	for i := 0; i < len(c.noBeaconRanges[p.y]); i++ {
		r := c.noBeaconRanges[p.y][i]
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
	c.noBeaconRanges[p.y] = nRanges
}

func (c *Cave) MarkNoBeacon(s Sensor, trackRow int) {
	distance := s.point.ManhattanDistance(s.beacon)
	x := 0
	switch {
	case s.point.y < trackRow:
		x = distance - (trackRow - s.point.y)
	case s.point.y >= trackRow:
		x = distance - (s.point.y - trackRow)
	}
	if x < 0 {
		return
	}
	c.recordRange(trackRow, s.point.x-x, s.point.x+x)
}

func (c *Cave) recordRange(y int, start int, end int) {
	if c.maxCoord > 0 {
		switch {
		case start < 0:
			start = 0
		case start > c.maxCoord:
			return
		case end > c.maxCoord:
			end = c.maxCoord
		case end < 0:
			return
		}
	}

	nRanges := c.noBeaconRanges[y]
	added := false
	for i := 0; i < len(c.noBeaconRanges[y]); i++ {
		r := c.noBeaconRanges[y][i]
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
	c.noBeaconRanges[y] = nRanges
}

func NewCave() *Cave {
	c := Cave{}
	c.noBeaconRanges = make(map[int][]Range)
	return &c
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
