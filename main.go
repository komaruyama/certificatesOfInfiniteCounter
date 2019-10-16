package main

import (
	"bufio"
	"fmt"
	"oritatami"
	"os"
	"path/filepath"
	"strings"
	"svgco"

	. "brick_automaton"

	. "./figo"
)

//skip sort

type FileStruct struct {
	name string
}

type ModuleStruct struct {
	f  map[string][]string
	l  map[string][]string
	ha map[string][]string
	r  map[string][]string
}

func main() {
	brickgraph := CreateGraph("Rtr", MakeMargin(), MakeColor())
	brickgraph.OutputProofTrees(MakeMargin(), MakeColor())

	//AutomatonSvg()

	//DebugEnv()
	//ExportOverview("Rtr", FigRtr)
	//OutputFigure("tglider", NodeTglider)

	//RunCheckAllGlider("glider")
}

func DebugEnv() {
	//prevCur string, prev string, above string, abovePrev string, delay int
	conf := BrickEnvToConf(Zig, "Rb", "Hn", "Ltrc", 3)
	output := ConformationToSvg(conf, MakeMargin(), MakeColor())
	OutputSvg(output, "testEnv")
}

func OutputFigure(filename string, figure func() []BeadFigure) {
	//func BeadFigureToGraph(beads []BeadFigure, centerPosition oritatami.Position, radius float64,
	//  gridSpace float64, labelYDev float64) treegraph {
	margin := MakeMargin()
	bf1 := figure()
	graph := BeadFigureToGraph(bf1, oritatami.Position{2, -2}, false, margin.GridSpace, margin.LabelYDev, "")
	tagst := GraphToStructs(graph, margin, svgco.TransformSvg{}, "")
	//output := svgco.MakeSvgString(tagst, -100, 30, 250, 140) // glider ex
	output := svgco.MakeSvgString(tagst, -300, -100, 800, 500) // glider ex
	OutputSvg(output, filename)
	fmt.Println("output", filename+".svg")
}

// export an overview svg
func ExportOverview(filename string, figF func() (OritatamiFigure, oritatami.Transcript)) {
	//var output []string
	path := "oritatami system"
	rulepath := filepath.Join(path, "ruleset.aos")
	//tpath := filepath.Join(path, "transcript.prm")
	adaptpath := filepath.Join(path, "adapt.prm")

	//beadtype
	beadtype := ReadAdapter(Readfile(adaptpath))
	//rule
	rule := ReadRule(Readfile(rulepath), beadtype)

	fig, w := figF()
	output := OFigureToSvg(fig, w, rule, 3, MakeMargin(), MakeColor2())

	OutputSvg(output, filename)
}

func ConformationToSvg(conformation oritatami.Conformation, margin TreeMargin, color TreeColor) []string {
	routeview := OsysRoutesToView(conformation.Routes[len(conformation.Routes)-1])
	fmt.Println("main l.80", len(routeview), conformation.LastFormed, len(conformation.Routes)-1)
	tgraph, _ := PathToGraph(conformation, conformation.LastFormed, routeview[0], true, "", margin, color)
	transform := svgco.TransformSvg{ /* "translate", []float64{x + width/2, y + height/2 */ }
	body := GraphToStructs(tgraph, margin, transform, "")

	return svgco.MakeSvgString(body, -1500, -20, 3410, 260) //counter
	//return svgco.MakeSvgString(body, 280, -20, 380, 115) // ha
}

func OFigureToSvg(figure OritatamiFigure, transcript oritatami.Transcript, ruleset oritatami.Ruleset, delay int,
	margin TreeMargin, color TreeColor) []string {

	fmt.Println("debug. seedw= ", figure.SeedW.Arr)
	startConf := oritatami.MakeNewConformationFromDirectionToSeed(figure.Dircarr, figure.SeedW, oritatami.Position{0, 0})
	//b, o := startConf.GetBead(3, 3)
	//if o {
	//	fmt.Println("bead:", b.Beadtype)
	//}
	terminal, ok := oritatami.FoldWithoutArity(startConf, delay, transcript, ruleset)

	if !ok {
		fmt.Println("<------@main l.99----->")
		return ConformationToSvg(terminal, margin, color)
	}

	return ConformationToSvg(terminal, margin, color)
}

////////////////////////////////////////////////
////////////////////////////////////////////////
////////////////////////////////////////////////
////////////////////////////////////////////////
////////////////////////////////////////////////
////////////////////////////////////////////////
/*
func AutomatonSvg() {
	hzig := FileStruct{"horizontal_zig"}
	hzag := FileStruct{"horizontal_zag"}
	vzig := FileStruct{"vertical_zig"}
	vzag := FileStruct{"vertical_zag"}
	vturn := FileStruct{"vertical_turn"}

	hzigStr := StrArrToStr(hzig.LoadFile())
	hzagStr := StrArrToStr(hzag.LoadFile())
	vzigStr := StrArrToStr(vzig.LoadFile())
	vzagStr := StrArrToStr(vzag.LoadFile())
	vturnStr := StrArrToStr(vturn.LoadFile())

	zigModules := ModuleStruct{}
	zagModules := ModuleStruct{}
	verticalZigModules := ModuleStruct{}
	verticalZagModules := ModuleStruct{}
	verticalTurnModules := ModuleStruct{}

	zigModules.ParseModules(hzigStr)
	zagModules.ParseModules(hzagStr)
	verticalZigModules.ParseModules(vzigStr)
	verticalZagModules.ParseModules(vzagStr)
	verticalTurnModules.ParseModules(vturnStr)

	zigBGraph := ParseBrick(zigModules, zagModules, verticalZigModules)
	zagBGraph := ParseBrick(zagModules, zigModules, verticalZagModules)
	turnBGraph := ReparseWithTurn(verticalTurnModules, zigModules, zagModules, &zigBGraph, &zagBGraph)

	fmt.Println(zigModules.f["Fnb"][0])

	fPlt := func(k *BrickEnv, v *BrickEnv, pass bool) {
		fmt.Print("From (e,prev)=(" + k.Elem + "," + k.Prev + "), (Upre, U)=(" + k.Upprev + "," + k.Upper + ") ")
		if pass {
			fmt.Println("-> pass")
		} else {
			fmt.Println("-> (e,prev)=(" + v.Elem + "," + v.Prev + "), (Upre, U)=(" + v.Upprev + "," + v.Upper + ")")
		}
	}

	fmt.Println("---F-----------")
	for k, v := range zigBGraph.Ef {
		if len(v) == 0 {
			fPlt(&k, &k, true)
		}
		for _, e := range v {
			fPlt(&k, &e, false)
		}
	}
	fmt.Println(len(zigBGraph.Ef), len(zigBGraph.Er), len(zigBGraph.Eha), len(zigBGraph.El))
	fmt.Println(len(zagBGraph.Ef), len(zagBGraph.Er), len(zagBGraph.Eha), len(zagBGraph.El))
	fmt.Println(len(turnBGraph.Er), len(turnBGraph.El))

	lssss := ConvertToSvg(zagBGraph, false)
	OutputSvg(lssss, "output")
}
*/
func ReparseWithTurn(turnM ModuleStruct, zigM ModuleStruct, zagM ModuleStruct,
	zigGraph *BrickGraph, zagGraph *BrickGraph) BrickGraph {
	var vl map[BrickEnv][]BrickEnv
	var vr map[BrickEnv][]BrickEnv

	var kl []BrickEnv
	var kr []BrickEnv

	bindFromAndElems := func(fromStArr map[string][]string, turnStArr map[string][]string, fromGraphMap map[BrickEnv][]BrickEnv) []BrickEnv {
		rV := []BrickEnv{}
		for k, v := range fromStArr {
			for _, e := range v {
				for sk, sv := range turnStArr {
					for _, se := range sv {
						if e == se {
							nenv := BrickEnv{e, k, sk, Empty}
							for zk, _ := range fromGraphMap {
								if zk.Elem == k && zk.Upprev == sk {
									fromGraphMap[zk] = append(fromGraphMap[zk], nenv)
								}
							}
							rV = append(rV, nenv)
							break
						}
					}
				}
			}
		}
		return rV
	}

	bindElemsAndTo := func(toGraphMap map[BrickEnv][]BrickEnv, elems []BrickEnv) map[BrickEnv][]BrickEnv {
		rV := make(map[BrickEnv][]BrickEnv)
		for k, _ := range toGraphMap {
			for _, se := range elems {
				if k.Prev == se.Elem && k.Upper == se.Prev {
					rV[se] = append(rV[se], k)
				}
			}
		}
		return rV
	}

	kl = bindFromAndElems(zigM.f, turnM.l, zigGraph.Ef)
	kr = bindFromAndElems(zagM.ha, turnM.r, zagGraph.Eha)

	vl = bindElemsAndTo(zagGraph.Eha, kl)
	vr = bindElemsAndTo(zigGraph.Ef, kr)

	return BrickGraph{make(map[BrickEnv][]BrickEnv), vl, make(map[BrickEnv][]BrickEnv), vr}
}

func ParseBrick(currentM ModuleStruct, oppositeM ModuleStruct, verticalM ModuleStruct) BrickGraph {
	vf := []BrickEnv{}
	vl := []BrickEnv{}
	vha := []BrickEnv{}
	vr := []BrickEnv{}

	for k, _ := range currentM.f {
		vf = append(vf, MakeBrickEnv(k, currentM.r, verticalM.ha, oppositeM.l)...)
	}
	for k, _ := range currentM.l {
		vl = append(vl, MakeBrickEnv(k, currentM.f, verticalM.l, oppositeM.f)...)
	}
	for k, _ := range currentM.ha {
		vha = append(vha, MakeBrickEnv(k, currentM.l, verticalM.f, oppositeM.r)...)
	}
	for k, _ := range currentM.r {
		vr = append(vr, MakeBrickEnv(k, currentM.ha, verticalM.r, oppositeM.ha)...)
	}

	ef := make(map[BrickEnv][]BrickEnv)
	el := make(map[BrickEnv][]BrickEnv)
	eha := make(map[BrickEnv][]BrickEnv)
	er := make(map[BrickEnv][]BrickEnv)

	for _, e := range vf {
		ef[e] = MakeEdgeFromBrickEnv(e, vl)
	}
	for _, e := range vl {
		el[e] = MakeEdgeFromBrickEnv(e, vha)
	}
	for _, e := range vha {
		eha[e] = MakeEdgeFromBrickEnv(e, vr)
	}
	for _, e := range vr {
		er[e] = MakeEdgeFromBrickEnv(e, vf)
	}

	return BrickGraph{ef, el, eha, er}
}

func StrArrToStr(lines []string) string {
	line := ""
	for _, l := range lines {
		line += l
	}
	return line
}

func MakeEdgeFromBrickEnv(focusEnv BrickEnv, nextEnvs []BrickEnv) []BrickEnv {
	rV := []BrickEnv{}
	for _, e := range nextEnvs {
		var fupprev string
		if focusEnv.Upprev == "" {
			fupprev = Empty
		} else {
			fupprev = focusEnv.Upprev
		}
		if focusEnv.Elem == e.Prev && fupprev == e.Upper {
			rV = append(rV, e)
		}
	}
	return rV
}

func MakeBrickEnv(focusE string, prevM map[string][]string, verticalM map[string][]string, upprevM map[string][]string) []BrickEnv {

	temp := []BrickEnv{}
	prevArr := []string{}
	for k, v := range prevM {
		for _, e := range v {
			if e == focusE {
				prevArr = append(prevArr, k)
				break
			}
		}
	}

	for k, v := range verticalM {
		for _, e := range v {
			if e == focusE {
				if k == Empty {
					temp = append(temp, BrickEnv{focusE, "", k, ""})
				}
				for sk, sv := range upprevM {
					for _, se := range sv {
						if k == se {
							temp = append(temp, BrickEnv{focusE, "", se, sk})
							break
						}
					}
				}
				break
			}
		}
	}

	rV := []BrickEnv{}

	for _, tv := range temp {
		for _, rv := range prevArr {
			rV = append(rV, BrickEnv{tv.Elem, rv, tv.Upper, tv.Upprev})
		}
	}
	return rV
}

func (file FileStruct) LoadFile() []string {
	f, err := os.Open(file.name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", file.name, err)
		os.Exit(1)
	}
	defer f.Close()

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if serr := scanner.Err(); serr != nil {
		fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", file.name, err)
	}

	return lines
}

func (modules *ModuleStruct) ParseModules(stream string) {
	spF := strings.SplitN(stream, "F:", 2)
	spL := strings.SplitN(stream, "L:", 2)
	spHA := strings.SplitN(stream, "HA:", 2)
	spR := strings.SplitN(stream, "R:", 2)

	if len(spF) > 1 {
		senF := strings.SplitN(strings.SplitN(spF[1], "{", 2)[1], "}", 2)
		modules.f = modules.SubParseSemicoron(senF[0])
	}
	if len(spL) > 1 {
		senL := strings.SplitN(strings.SplitN(spL[1], "{", 2)[1], "}", 2)
		modules.l = modules.SubParseSemicoron(senL[0])
	}
	if len(spHA) > 1 {
		senHA := strings.SplitN(strings.SplitN(spHA[1], "{", 2)[1], "}", 2)
		modules.ha = modules.SubParseSemicoron(senHA[0])
	}
	if len(spR) > 1 {
		senR := strings.SplitN(strings.SplitN(spR[1], "{", 2)[1], "}", 2)
		modules.r = modules.SubParseSemicoron(senR[0])
	}
}

func (m ModuleStruct) SubParseSemicoron(line string) map[string][]string {
	rV := make(map[string][]string)
	els := strings.Split(line, ";")
	for _, l := range els[:len(els)-1] {
		ftel := strings.SplitN(l, "->", 2)
		if len(ftel) > 1 {
			fel := strings.Split(ftel[0], ",")
			tel := strings.Split(ftel[1], ",")
			for index, e := range tel {
				tel[index] = strings.TrimSpace(e)
			}
			//fmt.Println(strings.TrimSpace(fel[0]), len(tel))
			for _, e := range fel {
				trimE := strings.TrimSpace(e)
				rV[trimE] = append(rV[trimE], tel...)
			}
		}
	}
	return rV
}

/*
func (modules ModuleStruct) SubParser1(lines []string) []string {

}
*/
