package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Team int

const (
	TEAM2 = Team(2)
	TEAM3 = Team(3)
	TEAM4 = Team(4)
)

func main() {
	fmt.Println("Starting algorithm")
	fmt.Printf("%v\n", solve())

	/*
		s := &state{}
		s.add(TEAM2, 1)
		s.add(TEAM2, 2)
		fmt.Printf("%v\n", s)
		cs := s.clone()
		cs.add(TEAM3, 3)
		fmt.Printf("%v\n", s)
		fmt.Printf("%v\n", cs)
	*/
}

func keys(m map[int][]string) []int {
	ks := make([]int, len(m))
	for i := range ks {
		ks[i] = i
	}
	return ks
}

// state
type state struct {
	t2, t3, t4    []int
	t2s, t3s, t4s [][]int
}

// add a value to a team
func (s *state) add(t Team, v int) {
	defer s.consistencyCheck()

	switch t {
	case TEAM2:
		s.t2 = append(s.t2, v)
	case TEAM3:
		s.t3 = append(s.t3, v)
	case TEAM4:
		s.t4 = append(s.t4, v)
	default:
		panic("unsupported team size")
	}
}

// consistencyCheck transfers 'full' teams to the corresponding slice
func (s *state) consistencyCheck() {
	if len(s.t2) == 2 {
		cp := make([]int, 2)
		copy(cp, s.t2)
		s.t2s = append(s.t2s, cp)
		s.t2 = []int{}
	}

	if len(s.t3) == 3 {
		cp := make([]int, 3)
		copy(cp, s.t3)
		s.t3s = append(s.t3s, cp)
		s.t3 = []int{}
	}

	if len(s.t4) == 4 {
		cp := make([]int, 4)
		copy(cp, s.t4)
		s.t4s = append(s.t4s, cp)
		s.t4 = []int{}
	}
}

func (s *state) clone() *state {
	t2 := make([]int, len(s.t2))
	t3 := make([]int, len(s.t3))
	t4 := make([]int, len(s.t4))
	t2s := make([][]int, len(s.t2s))
	t3s := make([][]int, len(s.t3s))
	t4s := make([][]int, len(s.t4s))
	copy(t2, s.t2)
	copy(t3, s.t3)
	copy(t4, s.t4)
	copy(t2s, s.t2s)
	copy(t3s, s.t3s)
	copy(t4s, s.t4s)
	return &state{
		t2, t3, t4, t2s, t3s, t4s,
	}
}

func (s *state) print() string {
	totalDelivered := strconv.Itoa(len(s.t2s) + len(s.t3s) + len(s.t4s))
	out := []string{totalDelivered}

	printTeams := func(teams [][]int) (out []string) {
		for _, team := range teams {
			str := []string{strconv.Itoa(len(team))}
			for _, pizza := range team {
				str = append(str, strconv.Itoa(pizza))
			}
			out = append(out, strings.Join(str, " "))
		}
		return
	}

	out = append(out, printTeams(s.t2s)...)
	out = append(out, printTeams(s.t3s)...)
	out = append(out, printTeams(s.t4s)...)

	return strings.Join(out, "\n")
}

// score returns the total score of this state
// this is calculated by looking at the unique ingredients a team obtained, squared
func (s *state) score(m map[int][]string) int {
	countUnique := func(team []int) int {
		ing := map[string]bool{}
		for _, pizza := range team {
			for _, in := range m[pizza] {
				ing[in] = true
			}
		}
		return len(ing)
	}

	countTeamScore := func(teams [][]int) int {
		teamScore := 0
		for _, team := range teams {
			u := countUnique(team)
			teamScore += u * u
		}
		return teamScore
	}

	return countTeamScore(s.t2s) + countTeamScore(s.t3s) + countTeamScore(s.t4s)
}

func solve() int {
	nt2, nt3, nt4, pm := parseInput()

	fmt.Printf("%v\n", keys(pm))

	var backtrack func(pz []int, s *state)

	maxState := state{}
	var max int

	backtrack = func(pz []int, s *state) {
		guard := len(s.t2s) == nt2 || len(s.t3s) == nt3 || len(s.t4s) == nt4

		if guard || len(pz) == 0 {
			if score := s.score(pm); score > max {
				max = score
				maxState = (*s)
			}
			return
		}

		head := pz[0]
		var tail []int
		if len(pz) > 1 {
			tail = pz[1:]
		}

		if len(s.t2s) <= nt2 {
			cs := s.clone()
			cs.add(TEAM2, head)
			backtrack(tail, cs)
		}

		if len(s.t3s) <= nt3 {
			cs := s.clone()
			cs.add(TEAM3, head)
			backtrack(tail, cs)
		}
		if len(s.t4s) <= nt4 {
			cs := s.clone()
			cs.add(TEAM4, head)
			backtrack(tail, cs)
		}

	}

	backtrack(keys(pm), &state{})

	fmt.Printf("%v\n", maxState.print())
	return maxState.score(pm)
}

func parseInput() (int, int, int, map[int][]string) {
	fileName := os.Args[1]

	s, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(s), "\n")
	header := strings.Split(lines[0], " ")

	a, b, c := parseHeader(header)

	m := map[int][]string{}
	for i, line := range lines[1:] {
		if line == "" {
			continue
		}
		m[i] = strings.Split(line, " ")[1:]
	}

	return a, b, c, m
}

func parseHeader(header []string) (int, int, int) {
	a, b, c := header[1], header[2], header[3]

	ia, _ := strconv.Atoi(a)
	ib, _ := strconv.Atoi(b)
	ic, _ := strconv.Atoi(c)

	return ia, ib, ic
}
