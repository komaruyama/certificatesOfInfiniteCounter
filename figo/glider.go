package figo

import (
	. "brick_automaton"
	"fmt"
	"math"
	"oritatami"
	"strconv"
	"svgco"
)

func RunCheckAllGlider(filename string) {
	margin := MakeMargin()
	color := MakeColor()

	outputF := func(terminal oritatami.Conformation, index int) {
		tgraph, _ := PathToGraph(terminal, terminal.LastFormed, Routeview{}, false, "", margin, color)
		transform := svgco.TransformSvg{ /* "translate", []float64{x + width/2, y + height/2 */ }
		body := GraphToStructs(tgraph, margin, transform, "")

		lines := svgco.MakeSvgString(body, 0, -200, 600, 600)

		OutputSvg(lines, filename+strconv.Itoa(index))
	}

	confs := CheckAllGliderNondet(outputF)

	fmt.Println("# of conformation", len(confs))
	//for index, terminal := range confs {
	//	outputF(terminal, index)
	//}
}

func CheckAllLongCounterNondet(outputF func(oritatami.Conformation, int)) []oritatami.Conformation {
	delay := 3
	radix := 3
	expectConf := "embedaembeda"
	//sq := []int{0, 0, 0, 0, 0, 0}
	//finishTranscript := false
	return CheckAllNondet(delay, radix, expectConf, GetGliderSeedAndTranscript, outputF)
}

func CheckAllGliderNondet(outputF func(oritatami.Conformation, int)) []oritatami.Conformation {
	delay := 3
	radix := 3
	expectConf := "embedaembeda"
	//sq := []int{0, 0, 0, 0, 0, 0}
	//finishTranscript := false
	return CheckAllNondet(delay, radix, expectConf, GetGliderSeedAndTranscript, outputF)
}

func CheckAllNondet(delay int, radix int, expectConf string, formationF func(sq ...int) (OritatamiFigure, oritatami.Transcript),
	outputF func(oritatami.Conformation, int)) []oritatami.Conformation {
	ans := []oritatami.Conformation{}
	fmt.Println("start")
	rulesetBinary := uint64(math.Pow(2.0, float64(radix)*float64(radix+1)/2.0))

	filecount := 0
	for rulesetBinary != 0 {
		ruleset := GetRulesetFromBinary(rulesetBinary-1, radix)
		finishTranscript := false
		sq := []int{0, 0, 0, 0, 0, 0}

		fmt.Println("----", rulesetBinary-1)
		for !finishTranscript {
			//fmt.Println(sq)
			seed, transcript := formationF(sq...)

			startConf := oritatami.MakeNewConformationFromDirectionToSeed(seed.Dircarr, seed.SeedW, oritatami.Position{0, 0})
			startConf.ReadyExecuting()
			terminal, ok := oritatami.FoldWithoutArity(startConf, delay, transcript, ruleset)
			if ok {
				if CheckMatchConf(terminal, expectConf) {
					ans = append(ans, terminal)
					outputF(terminal, filecount)
					filecount += 1
				}
			}
			sq, finishTranscript = GetNextNumber(sq, radix)
		}
		rulesetBinary -= 1
	}
	return ans
}

func GetRulesetFromBinary(binary uint64, radix int) oritatami.Ruleset { // binary msb: (0 -- 0)
	//type Ruleset struct {
	//	rule map[string][]string
	rule := oritatami.Ruleset{}
	rule.Initialize()

	count := uint(0)
	for i := radix; i > 0; i-- {
		for j := i; j > 0; j-- {
			if (binary>>count)%2 == 1 {
				rule.Add(strconv.Itoa(i-1), strconv.Itoa(j-1))
			}
			count += 1
		}
	}
	return rule
}

func GetNextNumber(now []int, radix int) ([]int, bool) { // radix 0, 1, ...
	for index, e := range now {
		if e == (radix - 1) {
			now[index] = 0
		} else {
			now[index] = e + 1
			return now, false
		}
	}
	return now, true
}

func GetGliderSeedAndTranscript(sq ...int) (OritatamiFigure, oritatami.Transcript) {
	seq6 := []int{sq[0], sq[1], sq[2], sq[3], sq[4], sq[5]}
	seq12 := append(seq6, seq6...)
	sequence := Ints2Strings(seq12)
	return GliderSeed(sq[0], sq[1], sq[2], sq[3], sq[4], sq[5]), oritatami.Transcript{sequence}
}

func GliderSeed(p1 int, p2 int, p3 int, p4 int, p5 int, p6 int) OritatamiFigure {
	Dircarr := "daembeda"
	sequence := Ints2Strings([]int{p4, p5, p6, p1, p2, p3, p4, p5, p6})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func CheckMatchConf(conf oritatami.Conformation, dirs string) bool {
	pos := conf.LastSeed
	bead, _ := conf.GetBead(pos.GetPos())
	for bead.NextDir != -1 && len(dirs) > 0 {
		if bead.NextDir != oritatami.DirExpChar2Byte(rune(dirs[0])) {
			return false
		}
		dirs = string(dirs[1:])
		pos = bead.GetPosition().NextPosition(bead.NextDir)
		bead, _ = conf.GetBead(pos.GetPos())
	}
	return true
}
