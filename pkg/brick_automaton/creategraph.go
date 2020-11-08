package brick_automaton

import (
	"fmt"
	"strings"

	"modules/pkg/oritatami"
	"modules/pkg/svgco"
)

func CreateGraph(seed string, margin TreeMargin, color TreeColor) BrickGraphVE {
	param := Parameters{ReadParameter(3)} // !!

	firstVertex := BrickVertex{RTM, TurnToZig, seed, Empty, Empty, Empty, Empty, "b", nil, nil} // !!
	elements := make(map[string]BrickModule)
	edge := make(map[string][]BrickVertex)
	ioconf := make(map[string]BrickModule)

	graph := BrickGraphVE{map[string]BrickVertex{firstVertex.IDlabel(): firstVertex}, BrickElements{elements}, edge, BrickElements{ioconf}}
	graph.Elements.LoadModules()
	graph.IOConformation.LoadIOConfs()

	graph.viewElements()

	err := []BrickError{}

	queue := []BrickVertex{firstVertex}

	makeEdge := func() {
		pred := queue[0]
		nextenvs := graph.NextEnv(pred)
		//fmt.Println(len(nextenvs), "lennnn")
		// run(sucTemp, pred) -> nexts
		sucTemp := graph.RunOsystemOnEnves(nextenvs, pred, param, margin, color, err)

		//debug
		//for ind, e := range sucTemp {
		//	fmt.Println("DebST i:", ind, "el pr ab abpre", e.Elem, e.Prev, e.Abov, e.Abpr)
		//}

		suc := graph.VerticesWithoutConnecting(pred, sucTemp)
		//debug
		//for ind, e := range suc {
		//	fmt.Println("DebS i:", ind, "el pr ab abpre", e.Elem, e.Prev, e.Abov, e.Abpr)
		//}

		graph.AddVertex(suc)
		graph.AddEdge(pred, suc)
		queue = append(queue[1:], suc...)
	}

	run := func() {
		for len(queue) > 0 {
			makeEdge()
			//fmt.Println("-----tt")
			//graph.viewEdge()
		}
	}

	for len(queue) > 0 {
		run()
		queue, _ = graph.CheckRemain(param, margin, color, err)
	}

	fmt.Println("Finished..")
	graph.viewEdge()

	return graph
}

func (g BrickGraphVE) RunOsystemOnEnves(environments []BrickVertex, previousEnv BrickVertex, osparam Parameters, margin TreeMargin, color TreeColor, errStack []BrickError) []BrickVertex {
	ans := []BrickVertex{}

	for _, env := range environments {
		terminal, ok := g.FoldModule(env.ModuleType, previousEnv.Current, env.Prev, env.Abov, env.Abpr, osparam, margin, color, errStack)
		if ok {
			inputc, okin := g.GetInputConformation(terminal)
			module, okmo := g.GetModule(terminal, env.ModuleType)
			outputc, outOSconf := g.GetOutputConformation(terminal, osparam.Params[NextModuleType(env.ModuleType)])
			if okin && okmo && outOSconf != nil {
				ans = append(ans, BrickVertex{env.ModuleType, previousEnv.NextCurrent(g.IsTurnModule(module)), module, env.Prev, env.Abov, env.Abpr, inputc, outputc, &terminal, outOSconf})
			} else {
				fmt.Println("Exception (Not Found Label)")
			}

		} else {
			fmt.Println("Exception (nondet environment)")
		}
	}
	return ans
}

func (g BrickGraphVE) NextEnv(from BrickVertex) []BrickVertex {
	// To give next environments (Verterx.Element is no module label)
	//fmt.Println("NextEnv from el pre ab abpre", from.Elem, from.Prev, from.Abov, from.Abpr, from.Current)

	moduleType := NextModuleType(from.ModuleType)
	prev := from.Elem
	var above string
	//fmt.Println("NNNNNNEw", prev)
	ans := []BrickVertex{}

	if from.IsTurn() {
		above = from.Prev
		if above == Empty {
			//debug
			//fmt.Println("return: after turn & above emp", "pr ab abpr", prev, "e", "e")
			return []BrickVertex{BrickVertex{moduleType, "", "", prev, Empty, Empty, "", "", nil, nil}} // input and output conformation ommit
		} else {
			var current string
			if from.Current == TurnToZag {
				current = Zig
			} else {
				current = Zag
			}
			abprevs := g.PrevModules(above, current)
			for _, abprev := range abprevs {
				ans = append(ans, BrickVertex{moduleType, "", "", prev, above, abprev, "", "", nil, nil}) // input and output conformation ommit
			}
			//debug
			//fmt.Println("return: after turn", "pr ab", prev, above)
			return ans
		}
	} else {
		above = from.Abpr
		if above == Empty {
			//debug
			//fmt.Println("return: after straight & above emp", "pr ab abpr", prev, "e", "e")
			return []BrickVertex{BrickVertex{moduleType, "", "", prev, Empty, Empty, "", "", nil, nil}} // input and output conformation ommit
		}
		if dt, _ := g.getDirctionType(above); dt == TurnToZig || dt == TurnToZag {
			//debug
			//fmt.Println("return: after straight & above turn", "pr ab abpr", prev, above, "e")
			return []BrickVertex{BrickVertex{moduleType, "", "", prev, above, Empty, "", "", nil, nil}} // input and output conformation ommit
		} else {
			var current string
			if from.Current == Zig {
				current = Zag
			} else {
				current = Zig
			}

			abprevs := g.PrevModules(above, current)
			for _, abprev := range abprevs {
				ans = append(ans, BrickVertex{moduleType, "", "", prev, above, abprev, "", "", nil, nil}) // input and output conformation ommit
			}
			//debug
			//fmt.Println("return: after straight & above straight", "pr ab", prev, above)
			return ans
		}
	}
	/*
		if g.IsTurnModule(above) {
			// osystem
			elem, ok := g.FoldModule(moduleType, from.Current, prev, above, Empty, osparam, margin, color, errStack)
			if ok {
				ans = []BrickVertex{BrickVertex{moduleType, from.NextCurrent(g.IsTurnModule(elem)), elem, prev, above, Empty}}
			}
		} else {
			abPrs := g.PrevModules(above)
			abPrs = append(abPrs, Empty)
			ans = []BrickVertex{}
			for _, e := range abPrs {
				// osystem
				elem, ok := g.FoldModule(moduleType, from.Current, prev, above, e, osparam, margin, color, errStack)
				if ok {
					ans = append(ans, BrickVertex{moduleType, from.NextCurrent(g.IsTurnModule(elem)), elem, prev, above, e})
				}
			}
		}
		return ans
	*/
}

func (g BrickGraphVE) PrevModules(successor string, sucCurrent string) []string {
	ans := []string{}

	set := make(map[string]struct{})
	for _, e := range g.Vertices {
		if e.Elem == successor && e.Current == sucCurrent {
			set[e.Prev] = struct{}{}
		}
	}
	for k, _ := range set {
		ans = append(ans, k)
	}
	return ans
}

func (g BrickGraphVE) VerticesWithoutConnecting(pre BrickVertex, sucs []BrickVertex) []BrickVertex {
	ans := []BrickVertex{}
	for _, e := range sucs {
		if !g.IsContainInEdge(pre, e) {
			ans = append(ans, e)
		}
	}
	return ans
}

func (g BrickGraphVE) CheckRemain(param Parameters, margin TreeMargin, color TreeColor, err []BrickError) ([]BrickVertex, bool) {
	for _, e := range g.Vertices {
		nextenvs := g.NextEnv(e)
		suct := g.RunOsystemOnEnves(nextenvs, e, param, margin, color, err)
		suc := g.VerticesWithoutConnecting(e, suct)
		if len(suc) > 0 {
			g.AddVertex(suc)
			g.AddEdge(e, suc)
			return suc, true
		}
	}
	return []BrickVertex{}, false
}

/*
func SettleRouteViews(allRoute []Routeview, terminal oritatami.Conformation) []Routeview {
	// settling all route views
	routeviews := []Routeview{}
	for _, ar := range allRoute {
		//fmt.Print(ar.routes[0].Bonds, "|")
		if ar.routes[0].Bonds == terminal.Routes[len(terminal.Routes)-1].Mbonds {
			routeviews = append(routeviews, ar)
		}
	}
	return routeviews
}
*/

func (g BrickGraphVE) FoldModule(module rune, prevCur string, prev string, above string, abovePrev string,
	osparam Parameters, margin TreeMargin, color TreeColor, errorStack []BrickError) (oritatami.Conformation, bool) {
	seed := g.Elements.MakeSeed(prevCur, prev, above, abovePrev, osparam)

	//debug
	//outputConf(seed, margin, color, "debug"+string(module))

	parameter := osparam.Params[module]
	terminal, det := oritatami.FoldWithoutArity(seed, parameter.Delay, parameter.Transcript, parameter.Rule)

	if !det {
		// Error
		routeviews := OsysRoutesToView(terminal.Routes[len(terminal.Routes)-1])
		fmt.Println("creategraph.go l.133", len(routeviews), terminal.Routes[len(terminal.Routes)-1].Mbonds)

		fmt.Println()
		graph, _ := PathToGraph(terminal, terminal.LastFormed, routeviews[0], true, "", margin, color)
		body := GraphToStructs(graph, margin, svgco.TransformSvg{}, "")

		strs := svgco.MakeSvgString(body, -180, -20, 2010, 260)

		OutputSvg(strs, "nondetTest")

		fmt.Println("nondet...", "creategraph.go", "FoldModule")
		return oritatami.Conformation{}, false
	} else {
		return terminal, true
	}
}

func (g BrickGraphVE) GetInputConformation(terminal oritatami.Conformation) (string, bool) {
	dirc := ""
	pos := terminal.LastSeed
	fbead, _ := terminal.GetBead(pos.GetPos())
	for fbead.HasSuccessor() {
		dir := fbead.NextDir
		dirc += string(oritatami.DirExpByte2Char(dir))
		pos = pos.NextPosition(dir)
		fbead, _ = terminal.GetBead(pos.GetPos())
	}

	labeled, ok := g.GetIOConformationLabel(dirc)
	return labeled, ok
}

func (g BrickGraphVE) GetModule(terminal oritatami.Conformation, moduleType rune) (string, bool) {
	dirc := ""
	pos := terminal.LastSeed
	fbead, _ := terminal.GetBead(pos.GetPos())
	for fbead.HasSuccessor() {
		dir := fbead.NextDir
		dirc += string(oritatami.DirExpByte2Char(dir))
		pos = pos.NextPosition(dir)
		fbead, _ = terminal.GetBead(pos.GetPos())
	}

	labeled, ok := g.GetModuleLabel(dirc, moduleType)
	return labeled, ok
}

func (g BrickGraphVE) GetOutputConformation(terminal oritatami.Conformation, parameter oritatami.Parameter) (string, *oritatami.Conformation) {
	terminal.MoveToSeed()
	outputConf, _ := oritatami.FoldWithoutArity(terminal, parameter.Delay, parameter.Transcript, parameter.Rule)

	dirc := ""
	pos := outputConf.LastSeed
	fbead, _ := outputConf.GetBead(pos.GetPos())
	for fbead.HasSuccessor() {
		dir := fbead.NextDir
		dirc += string(oritatami.DirExpByte2Char(dir))
		pos = pos.NextPosition(dir)
		fbead, _ = outputConf.GetBead(pos.GetPos())
	}

	labeled, ok := g.GetIOConformationLabel(dirc)
	//if ok {
	//	fmt.Println("l.304", len(dirc), labeled, len(outputConf.Routes))
	//	terminal.Routes = outputConf.Routes[:len(termCp.Routes)+len(dirc)]
	//}
	retoutconf := &outputConf
	if !ok {
		retoutconf = nil
	}
	return labeled, retoutconf
}

func (be *BrickElements) LoadModules() {
	modules := ParseModule()
	for _, el := range modules {
		be.Elements[el.Label] = el
	}
}

func (be *BrickElements) LoadIOConfs() {
	modules := ParseIOConformation()
	for _, el := range modules {
		be.Elements[el.Label] = el
	}
}

func (e BrickElements) MakeSeed(prevCur string, prev string, above string, abovePrev string, osparam Parameters) oritatami.Conformation {
	//fmt.Println("__Seed...(prev, above, above prev) =", prev, above, abovePrev)
	var start oritatami.Position
	var aboveEnd oritatami.Position
	var abpreTransAdj oritatami.Position

	var seed oritatami.BeadMap
	prevModule := e.Elements[prev]
	aboveModule := e.Elements[above]
	abvePrevModule := e.Elements[abovePrev]

	if prevCur == Zig || prevCur == TurnToZig {
		if prty := prevModule.DirectionType; prty == Zag || prty == TurnToZag {
			prevModule.Path = oritatami.FlipDirection(prevModule.Path)
		}
		start = PositionShift(oritatami.Position{0, 0}, prevModule.OutputCarry, true)

		if above != Empty {
			aboveEnd = PositionShift(oritatami.Position{2, -3}, aboveModule.OutputCarry, false)
			abpreTransAdj = oritatami.Position{-1, 0}
			if abty := aboveModule.DirectionType; abty == Zig || abty == TurnToZig {
				aboveModule.Path = oritatami.FlipDirection(aboveModule.Path)
			}

			if abovePrev != Empty {
				if abprty := abvePrevModule.DirectionType; abprty == Zig || abprty == TurnToZig {
					abvePrevModule.Path = oritatami.FlipDirection(abvePrevModule.Path)
				}
			}
		}
	} else {
		if prty := prevModule.DirectionType; prty == Zig || prty == TurnToZig {
			prevModule.Path = oritatami.FlipDirection(prevModule.Path)
		}
		start = PositionShift(oritatami.Position{0, 0}, prevModule.OutputCarry, false)

		if above != Empty {
			aboveEnd = PositionShift(oritatami.Position{1, -3}, aboveModule.OutputCarry, true)
			if apmodule, _ := e.Elements[abovePrev]; apmodule.OutputCarry == Middle { //!
				abpreTransAdj = oritatami.Position{0, 1}
			} else {
				abpreTransAdj = oritatami.Position{1, 0}
			}
			if abty := aboveModule.DirectionType; abty == Zag || abty == TurnToZag {
				aboveModule.Path = oritatami.FlipDirection(aboveModule.Path)
			}
			if abovePrev != Empty {
				if abprty := abvePrevModule.DirectionType; abprty == Zag || abprty == TurnToZag {
					abvePrevModule.Path = oritatami.FlipDirection(abvePrevModule.Path)
				}
			}
		}
	}
	prconf := oritatami.MakeNewConformationFromDirectionToSeed(prevModule.Path, osparam.Params[prevModule.GetType()].Transcript, oritatami.Position{0, 0})
	seed = prconf.Seed.MapMarge(transPosition(prconf.LastFormed, start))

	if above != Empty {
		abconfStartPos := oritatami.Position{0, 0}
		abconf := oritatami.MakeNewConformationFromDirectionToSeed(aboveModule.Path, osparam.Params[aboveModule.GetType()].Transcript, abconfStartPos)
		transab := transPosition(abconf.LastFormed, aboveEnd)
		seed = abconf.Seed.MapMarge(transab, seed)

		if abovePrev != Empty {
			abprconf := oritatami.MakeNewConformationFromDirectionToSeed(abvePrevModule.Path, osparam.Params[abvePrevModule.GetType()].Transcript, oritatami.Position{0, 0})
			seed = abprconf.Seed.MapMarge(transPosition(abprconf.LastFormed, abconfStartPos.TransPosition(transab).TransPosition(abpreTransAdj)), seed)
		}
	}

	plainBead := oritatami.BeadMap{}
	plainBead.Initialize()
	return oritatami.Conformation{seed, start, plainBead, start, []oritatami.Routes{}}
}

////////////////////////////////////
////////////////////////////////////
////////////////////////////////////

func (g BrickGraphVE) IsTurnModule(label string) bool {
	v, ok := g.Elements.Elements[label]
	if ok {
		if v.DirectionType == TurnToZig || v.DirectionType == TurnToZag {
			return true
		}
	}
	return false
}

func (g BrickGraphVE) IsContainInEdge(p BrickVertex, s BrickVertex) bool {
	sl, ok := g.Edges[p.IDlabel()]
	if ok {
		for _, e := range sl {
			if e.equal(s) {
				return true
			}
		}
	}
	return false
}
func (g *BrickGraphVE) AddVertex(vertices []BrickVertex) {
	for _, v := range vertices {
		g.Vertices[v.IDlabel()] = v
	}
}

func (g *BrickGraphVE) AddEdge(v1 BrickVertex, v2s []BrickVertex) {
	g.Edges[v1.IDlabel()] = append(g.Edges[v1.IDlabel()], v2s...)
}

func (g BrickGraphVE) GetIOConformationLabel(direction string) (string, bool) {
	// Searching same direction IOconformation from IOConformation in brick graph VE
	check := func(dir string) (string, bool) {
		for _, v := range g.IOConformation.Elements {
			if strings.HasPrefix(dir[1:], v.Path[1:]) {
				return v.Label, true
			}
		}
		return "", false
	}

	ans1, ok1 := check(direction)
	if ok1 {
		return ans1, true
	}

	ans2, ok2 := check(oritatami.FlipDirection(direction))
	if ok2 {
		return ans2, true
	}

	fmt.Println("------- Not Found:", direction, oritatami.FlipDirection(direction))
	return "", false
}

func (g BrickGraphVE) GetModuleLabel(direction string, moduleType rune) (string, bool) {

	check := func(dir string) (string, bool) {
		for _, v := range g.Elements.Elements {
			if v.GetType() == moduleType && dir == string(v.ComeFrom)+v.Path {
				return v.Label, true
			}
		}
		return "", false
	}

	ans1, ok1 := check(direction)
	if ok1 {
		return ans1, true
	}

	ans2, ok2 := check(oritatami.FlipDirection(direction))
	if ok2 {
		return ans2, true
	}

	fmt.Println("------- Not Found (module label):", string(moduleType), direction, oritatami.FlipDirection(direction))
	return "", false
}

func (v BrickVertex) IsTurn() bool {
	if v.Current == TurnToZag || v.Current == TurnToZig {
		return true
	}
	return false
}

func (v BrickVertex) NextCurrent(isTurn bool) string {
	if isTurn {
		switch v.Current {
		case TurnToZag:
			return TurnToZig
		case TurnToZig:
			return TurnToZag
		case Zig:
			return TurnToZag
		case Zag:
			return TurnToZig
		}
	} else {
		switch v.Current {
		case TurnToZag:
			return Zag
		case TurnToZig:
			return Zig
		case Zig:
			return Zig
		case Zag:
			return Zag
		}
	}
	return TurnToZig
}

func PositionShift(topPosition oritatami.Position, posRune rune, isZig bool) oritatami.Position {
	switch posRune {
	case Top:
		return topPosition
	case Middle:
		if isZig {
			return oritatami.Position{topPosition.X - 1, topPosition.Y + 1}
		} else {
			return oritatami.Position{topPosition.X, topPosition.Y + 1}
		}
	case Bottom:
		return oritatami.Position{topPosition.X - 1, topPosition.Y + 2}
	}
	return topPosition
}

func transPosition(from oritatami.Position, to oritatami.Position) oritatami.Position {
	return oritatami.Position{to.X - from.X, to.Y - from.Y}
}

///////////////////
//// debug
///////////////////
func (g BrickGraphVE) viewElements() {
	fmt.Println("print graph elements----")
	for _, v := range g.Elements.Elements {
		viewBrickModule(&v)
	}
}

func (g BrickGraphVE) viewEdge() {
	fmt.Println("print graph edges----")

	for e1, v := range g.Edges {
		fmt.Println("(", vertexToString(g.Vertices[e1]), ")")
		for _, e2 := range v {
			fmt.Println(" --(", vertexToString(e2), ")")
		}
	}

}

func vertexToString(v BrickVertex) string {
	ans := "Type: "
	ans += string(v.ModuleType)
	ans += ", Module: "
	ans += v.Elem
	ans += " ["
	ans += v.Prev
	ans += ", "
	ans += v.Abov
	ans += ", "
	ans += v.Abpr
	ans += "] : prev, above, above prev"
	ans += " ["
	ans += v.InputConf
	ans += ","
	ans += v.OutputConf
	ans += "] : InC OutC"
	return ans
}

func viewBrickModule(v *BrickModule) {
	fmt.Println(v.Label, v.ComeFrom, v.Path, " (in: ", v.InputCarry, " out:", v.OutputCarry, "), dir type=", v.DirectionType)
}

func outputConf(conformation oritatami.Conformation, margin TreeMargin, color TreeColor, filename string) {
	tgraph, _ := PathToGraph(conformation, conformation.LastFormed, Routeview{}, false, "", margin, color)
	transform := svgco.TransformSvg{ /* "translate", []float64{x + width/2, y + height/2 */ }
	body := GraphToStructs(tgraph, margin, transform, "")

	bodystrings := svgco.MakeSvgString(body, -1000, 0, 3000, 800)

	OutputSvg(bodystrings, filename)
	fmt.Print("debug.. outputConf at creategraph LastFormed:")
	fmt.Print(conformation.LastFormed.GetPos())
	bead, ok := conformation.GetBead(conformation.LastFormed.GetPos())
	if ok {
		fmt.Println("last bead:", bead.Beadtype)
	} else {
		fmt.Println()
	}
}
