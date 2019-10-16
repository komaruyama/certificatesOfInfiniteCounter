package brick_automaton

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/komaruyama/certificatesOfInfiniteCounter/pkg/oritatami"
	"github.com/komaruyama/certificatesOfInfiniteCounter/pkg/svgco"
)

type TreeMargin struct {
	SpaceVertical float64
	//oneblockWidth  float64
	BeadRadius float64
	GridSpace  float64
	LabelYDev  float64
	//EdgeThick    float64
	RoutesRadius float64
	//RouteEdgeThick float64
	MultBeadRadius float64
	MultEdgeWidth  float64

	NodeCorner         float64
	DecideNodeThick0   float64
	DecideNodeThick1   float64
	NotDecideNodeThick float64
	//NodeSubFrameThick  float64
	//NodeSubFrameWidth  float64
	//NodeSubFrameHeight float64
	LeftSubAreaWidth float64

	PositionPrec int
	StrokePrec   int
	FontsizePrec int
}

type TreeColor struct {
	FoldingEdge       svgco.ShapeColor
	FormedEdge        map[rune]svgco.ShapeColor
	SeedEdge          map[rune]svgco.ShapeColor
	FoldingBead       svgco.ShapeColor
	FormedBead        map[rune]svgco.ShapeColor
	SeedBead          map[rune]svgco.ShapeColor
	Bond              svgco.ShapeColor
	Font              svgco.FontSvg
	RouteColors       []string
	DecideNode        string
	DefaultSeedEdge   svgco.ShapeColor
	DefaultSeedBead   svgco.ShapeColor
	DefaultFormedEdge svgco.ShapeColor
	DefaultFormedBead svgco.ShapeColor
}

type treegraph struct {
	V        map[string]svgco.GraphVertex
	E        []svgco.GraphEdge
	Determin bool
}

type Routeview struct {
	later0     int // routes have no bond after this number (-1: all beads have no bond)
	routes     []oritatami.Route
	transcript []string
	detDir     int8
	isMaxbond  bool
}

type BeadForView struct {
	Beadtype   string
	pos        oritatami.Position
	PrevDir    int8
	StabconDir uint8
}

type BeadFigure struct {
	Bead       BeadForView
	Delta      FloatPosition
	IDsuffix   string
	BeadColor  svgco.ShapeColor
	BeadRadius float64
	PathColor  svgco.ShapeColor
	BondColor  svgco.ShapeColor
	Labeled    bool
	FontColor  svgco.FontSvg
}

type FloatPosition struct {
	x float64
	y float64
}

type clipAndBody struct {
	clips   []svgco.TagStruct
	bodyies []svgco.TagStruct
}

func (bg BrickGraphVE) OutputProofTrees(margin TreeMargin, color TreeColor) {
	index := 0
	for _, vert := range bg.Vertices {
		if vert.conformation != nil {
			strs := bg.proofTreeSvg(vert, margin, color)
			OutputSvg(strs, strconv.Itoa(index)+vert.IDlabel()+"prooftree")
			index += 1
		}
	}
}

func (bg BrickGraphVE) proofTreeSvg(vertex BrickVertex, margin TreeMargin, color TreeColor) []string {
	font := color.Font
	font.Font_size = margin.GridSpace // !!
	inputClabel := svgco.TextSvg{0.0, 0.0, "Input conformation: " + vertex.InputConf, font, "start", 0.0, svgco.TransformSvg{}}

	sx, sy := 0.0, margin.GridSpace*2
	width := 2600.0

	positionShift := func() (oritatami.Position, int) {
		fpos := vertex.conformation.LastSeed
		inputModule := bg.IOConformation.Elements[vertex.InputConf]

		for i := 0; i < len(inputModule.Path); i++ {
			bead, ok := vertex.conformation.GetBead(fpos.GetPos())
			if ok {
				fpos = fpos.NextPosition(bead.NextDir)
			} else {
				break
			}
		}
		return fpos, len(inputModule.Path)
	}

	clipanss := []clipAndBody{}

	nowposition, nowindex := positionShift()

	for true {
		routeviews := OsysRoutesToView(vertex.conformation.Routes[nowindex])
		beadtype := vertex.conformation.Routes[nowindex].Transcript[0]

		svg, onesvgheight := bg.OneBeadGraphs(sx, sy, width, routeviews, beadtype, nowposition, vertex.conformation, margin, color)
		sy += (onesvgheight + margin.SpaceVertical)

		clipanss = append(clipanss, svg)

		nbead, ok := vertex.conformation.GetBead(nowposition.GetPos())
		if nowindex == len(vertex.conformation.Routes)-1 {
			fmt.Println("too few routes")
			break
		}
		if !ok || nowposition == vertex.conformation.LastFormed {
			break
		}
		nowposition = nowposition.NextPosition(nbead.NextDir)
		nowindex += 1
	}

	outputModule := bg.IOConformation.Elements[vertex.OutputConf]
	for i := 0; i < len(outputModule.Path); i++ {
		nowindex += 1
		nbead, ok := vertex.OutOSConf.GetBead(nowposition.GetPos())
		if !ok {
			break
		}
		nowposition = nowposition.NextPosition(nbead.NextDir)

		routeviews := OsysRoutesToView(vertex.OutOSConf.Routes[nowindex])
		beadtype := vertex.OutOSConf.Routes[nowindex].Transcript[0]

		svg, onesvgheight := bg.OneBeadGraphs(sx, sy, width, routeviews, beadtype, nowposition, vertex.OutOSConf, margin, color)
		sy += (onesvgheight + margin.SpaceVertical)

		clipanss = append(clipanss, svg)
	}

	clipstrs := []svgco.TagStruct{}
	bodystrs := []svgco.TagStruct{inputClabel.ToStruct(margin.PositionPrec, margin.StrokePrec)}

	for _, ca := range clipanss {
		clipstrs = append(clipstrs, ca.clips...)
		bodystrs = append(bodystrs, ca.bodyies...)
	}

	outputClabel := svgco.TextSvg{0.0, sy, "Output conformation: " + vertex.OutputConf, font, "start", 0.0, svgco.TransformSvg{}}
	sy += margin.GridSpace * 2

	bodystrs = append(bodystrs, outputClabel.ToStruct(margin.PositionPrec, margin.StrokePrec))

	return svgco.MakeSvgString(append(clipstrs, bodystrs...), -50, -50, int(width)+300, int(sy)+50)
}

func (bg BrickGraphVE) OneBeadGraphs(x float64, y float64, maxWidth float64, routes []Routeview, beadtype string, stabilized oritatami.Position, terminal *oritatami.Conformation, margin TreeMargin, color TreeColor) (clipAndBody, float64) {

	// ready to...
	clipsStr := []svgco.TagStruct{}
	bodiesStr := []svgco.TagStruct{}
	xcr, ycr := x+margin.LeftSubAreaWidth, y
	renhei := 0.0

	font := color.Font
	font.Font_size = margin.GridSpace / 2.0 // !!
	labelBtype := svgco.TextSvg{x, y, "beadtype: " + beadtype, font, "start", margin.GridSpace, svgco.TransformSvg{}}

	for index, oneRoute := range routes {
		graph, radius := PathToGraph(*terminal, stabilized, oneRoute, true, "", margin, color)

		if radius != 0.0 {

			bonds := oneRoute.routes[0].Bonds
			routes := len(oneRoute.routes)
			frameWidth := (radius + margin.GridSpace) * 2

			id := beadtype + "-" + strconv.Itoa(index) + "-graphClip"
			if xcr > maxWidth {
				xcr = x + margin.LeftSubAreaWidth
				ycr += renhei
				renhei = 0.0
			}
			if renhei < frameWidth+(margin.GridSpace*3/2) {
				renhei = frameWidth + (margin.GridSpace * 3 / 2) // !!
			}

			cli, body := MakeProofTreeElement(id, xcr, ycr, frameWidth, frameWidth, graph, graph.Determin, bonds, routes, color, margin)
			clipsStr = append(clipsStr, cli)
			bodiesStr = append(bodiesStr, body...)

			xcr += frameWidth + 10 // !!
		}
	}

	return clipAndBody{clipsStr, append([]svgco.TagStruct{labelBtype.ToStruct(margin.PositionPrec, margin.StrokePrec)}, bodiesStr...)}, ycr + renhei - y
}

//////////////////////////
///     Proof Tree     ///
//////////////////////////

func MakeProofTreeElement(clipid string, x float64, y float64, width float64, height float64,
	graph treegraph, decide bool, bonds uint, routes int, color TreeColor, margin TreeMargin) (svgco.TagStruct, []svgco.TagStruct) {
	// -f-> (clip, svg)

	// this clip ract has to expand opposite direction of graph translate
	clipTrans := svgco.TransformSvg{"translate", []float64{-x - width/2, -y - height/2}}

	//color fill = false だと fill = None になるけど動く？

	// clip resion
	frameclip := svgco.RectSvg{x, y, width, height, margin.NodeCorner, clipTrans, svgco.ShapeColor{}, ""}
	clipStruct := svgco.MakeClipTagStruct(clipid, []svgco.TagStruct{frameclip.ToStruct(margin.PositionPrec, margin.StrokePrec)})

	// frame rectangle
	frameRects := []svgco.RectSvg{}
	labels := []svgco.TextSvg{}
	if decide {
		frameIn := svgco.RectSvg{x, y, width, height, margin.NodeCorner, svgco.TransformSvg{},
			svgco.ShapeColor{true, "#ffffff", margin.DecideNodeThick0, false, ""}, ""}
		frameOut := svgco.RectSvg{x, y, width, height, margin.NodeCorner, svgco.TransformSvg{},
			svgco.ShapeColor{true, color.DecideNode, margin.DecideNodeThick1, false, ""}, ""}

		frameRects = append(frameRects, frameIn, frameOut)
	} else {
		frame := svgco.RectSvg{x, y, width, height, margin.NodeCorner, svgco.TransformSvg{},
			svgco.ShapeColor{true, "#000000", margin.NotDecideNodeThick, false, ""}, ""}
		frameRects = append(frameRects, frame)
	}

	/*bondsframe := svgco.RectSvg{x, y, margin.NodeSubFrameWidth, margin.NodeSubFrameHeight, margin.NodeCorner, svgco.TransformSvg{},
		svgco.ShapeColor{true, "#000000", margin.NodeSubFrameThick, false, ""}, ""}
	frameRects = append(frameRects, bondsframe)*/
	//TextSvg
	//X         float64
	//Y         float64
	//Text      string
	//Font      FontSvg
	//Anchor    string // start or middle or end
	//DetY      float64
	//Transform TransformSvg
	font := color.Font
	font.Font_size = margin.GridSpace / 2.0 // !!
	labels = append(labels, svgco.TextSvg{x, y + height, "# of bonds: " + strconv.Itoa(int(bonds)), font, "start", margin.GridSpace / 2.0, svgco.TransformSvg{}})

	if routes > 1 {
		labels = append(labels, svgco.TextSvg{x, y + height, "# of paths: " + strconv.Itoa(routes), font, "start", margin.GridSpace, svgco.TransformSvg{}})
	}

	// "items" contains the graph and frames
	params := make(map[string]string)
	transform := svgco.TransformSvg{"translate", []float64{x + width/2, y + height/2}}
	transform.AppendToParam(margin.PositionPrec, params)
	svgco.AppendClip(params, clipid)
	body := GraphToStructs(graph, margin, svgco.TransformSvg{}, "")

	items := []svgco.TagStruct{}
	for _, r := range frameRects {
		items = append(items, r.ToStruct(margin.PositionPrec, margin.StrokePrec))
	}
	for _, l := range labels {
		items = append(items, l.ToStruct(margin.PositionPrec, margin.StrokePrec))
	}
	items = append(items, svgco.NewTagStructure("g", false, params, body))

	return clipStruct, items
}

/////////////
// override
func GraphToStructs(graph treegraph, margin TreeMargin, transform svgco.TransformSvg, clip string) []svgco.TagStruct {
	return svgco.GraphToStructs(graph.E, graph.V, margin.PositionPrec, margin.StrokePrec, margin.FontsizePrec, transform, "")
}

// Routes []map[int][]Route -> SVG []string

// Route map[int][]Route -> Graph (GraphVertex, GraphEdge)

/*
func MakeRouteGraph(route map[int][]oritatami.Route, conformation oritatami.Conformation, sx float64, sy float64, width float64, margin TreeMargin, color TreeColor) []treegraph {
	ans := []treegraph{}
	assort := []routeview{}



	for k, v := range route {
		if k != 0 {
			for _, e := range v {
				assort = frappend(assort, &e)
			}
		}
	}

	// [][]route -> graph (v,e)
	graphs, raduises := []treegraph{}, []float64{}
	for _, rs := range assort {
		g, r := PathToGraph(conformation, rs.routes[0].Detpos.NextPosition(int8(5-rs.routes[0].Path[0])), rs, true, "", margin, color)
		graphs = append(graphs, g)
		raduises = append(raduises, r)
	}
}
*/ // とりあえずコメントアウト↑

////////////////////////
///     os 2 svg     ///
////////////////////////

func getID(pos oritatami.Position, suffix string) string {
	//to give a vertex id from a position and a suffix.
	return strconv.Itoa(int(pos.X)) + " " + strconv.Itoa(int(pos.Y)) + " " + suffix
}
func pathToEdge(pos1 oritatami.Position, pos2 oritatami.Position, idSuffix string, edgecolor svgco.ShapeColor) svgco.GraphEdge {
	//giving edge from args as two vertices
	v1 := getID(pos1, idSuffix)
	v2 := getID(pos2, idSuffix)
	return svgco.GraphEdge{v1, v2, edgecolor}
}
func beadToVertex(bead *BeadForView, dx float64, dy float64, idSuffix string, center oritatami.Position,
	beadcolor svgco.ShapeColor, radius float64, gridSpace float64, clip string) svgco.GraphVertex {
	//giving a vertex from oritatami bead.
	bp := bead.pos
	id := getID(bp, idSuffix)

	cx, cy := positionToPoint(bp, center, gridSpace)
	//radius := margin.BeadRadius

	circle := svgco.CircleSvg{cx + dx, cy + dy, radius, svgco.TransformSvg{}, beadcolor, clip}

	v := svgco.GraphVertex{}
	v.ID = id
	v.Vertex = circle
	return v
}
func adaptLabelToVertex(v *svgco.GraphVertex, text string, labelDy float64, fontcolor svgco.FontSvg) {
	//Name to a vertex
	tx, ty := v.Vertex.Cx, v.Vertex.Cy
	anchor := "middle"
	label := svgco.TextSvg{tx, ty, text, fontcolor, anchor, labelDy, svgco.TransformSvg{}}
	v.Label = label
}

func getCenterpositionAndViewRadius(conformation *oritatami.Conformation, routes Routeview, stabilized oritatami.Position, margin TreeMargin) (oritatami.Position, float64) {
	routeposition := stabilized
	pss := []oritatami.Position{}
	//fmt.Println("os2svg.go l.158", " len[rview.routes]", len(routes.routes), " len[later0]", routes.later0, len(routes.routes[0].Path)-1, " bonds:", routes.routes[0].Bonds, "|t:", routes.transcript)
	if routes.later0 > -1 {
		for _, e := range routes.routes[0].Path[:routes.later0+1] {
			routeposition = routeposition.NextPosition(e)
			pss = append(pss, routeposition)
		}
	}

	if routes.later0 < len(routes.routes[0].Path)-1 {
		for _, r := range routes.routes {
			routepos2 := routeposition // routeposition is the end of match
			for _, e := range r.Path[routes.later0+1:] {
				routepos2 = routepos2.NextPosition(e)
				pss = append(pss, routepos2)
			}
		}
	}
	return midpointAndMargin(pss, margin)
}

//
func appendBeadToGraph(vs map[string]svgco.GraphVertex, es []svgco.GraphEdge, bead *BeadForView, centerposition oritatami.Position,
	beadcolor svgco.ShapeColor, pathcolor svgco.ShapeColor, font svgco.FontSvg, bondcolor svgco.ShapeColor, radius float64, margin TreeMargin, clip string) []svgco.GraphEdge {
	//appending a vertex and an edge as bead to graph (svg conformation)
	v := beadToVertex(bead, 0, 0, "", centerposition, beadcolor, radius, margin.GridSpace, clip)
	adaptLabelToVertex(&v, bead.Beadtype, margin.LabelYDev, font)
	vs[v.ID] = v
	p0 := bead.pos
	es = append(es, pathToEdge(p0, p0.NextPosition(bead.PrevDir), "", pathcolor))
	for i := uint(0); i < 6; i++ {
		if (bead.StabconDir>>i)%2 == 1 {
			es = append(es, pathToEdge(p0, p0.NextPosition(int8(i)), "", bondcolor))
		}
	}
	return es
}

func seedBeadFigures(conformation oritatami.Conformation, margin TreeMargin, color TreeColor) []BeadFigure {

	seeds := []BeadFigure{}
	for _, x := range conformation.Seed.Getmap() {
		for _, bead := range x {
			/* // beadtypeごとに色を分ける
			v := beadToVertex(bead, 0, 0, "", centerposition, color.FormedBead[getModuleLabel(bead.Beadtype)], margin.BeadRadius, margin.GridSpace, clip)
			*/
			bcolor := color.GetSeedBeadColor(getModuleLabel(bead.Beadtype))
			pcolor := color.GetSeedEdgeColor(getModuleLabel(bead.Beadtype))
			beadfigure := BeadFigure{*osysBeadToView(bead), FloatPosition{}, "", bcolor, margin.BeadRadius, pcolor, color.Bond, true, color.Font}
			seeds = append(seeds, beadfigure)
		}
	}
	return seeds
}
func formedBeadFigures(conformation oritatami.Conformation, stabilized oritatami.Position, margin TreeMargin, color TreeColor) []BeadFigure {

	formeds := []BeadFigure{}

	FormedBead, _ := conformation.GetBead(conformation.LastSeed.GetPos())
	dir := FormedBead.NextDir

	for FormedBead.GetPosition() != stabilized && dir != int8(-1) {
		FormedBead, _ = conformation.GetBead(FormedBead.GetPosition().NextPosition(dir).GetPos())
		dir = FormedBead.NextDir

		bead := osysBeadToView(FormedBead)
		bcolor := color.GetFormedBeadColor(getModuleLabel(bead.Beadtype))
		pcolor := color.GetFormedEdgeColor(getModuleLabel(bead.Beadtype))
		beadfigure := BeadFigure{*bead, FloatPosition{}, "", bcolor, margin.BeadRadius, pcolor, color.Bond, true, color.Font}
		formeds = append(formeds, beadfigure)
	}
	return formeds
}

////////
func BeadFigureToGraph(beads []BeadFigure, centerPosition oritatami.Position, det bool, gridSpace float64, labelYDev float64, clip string) treegraph {
	//Bead Figures To Graph (svg conformation)

	//type BeadFigure
	//	Bead      oritatami.Bead
	//	BeadColor svgco.ShapeColor
	//	PathColor svgco.ShapeColor
	//	BondColor svgco.ShapeColor

	vertices := make(map[string]svgco.GraphVertex)
	edges := []svgco.GraphEdge{}
	for _, b := range beads {
		v := beadToVertex(&b.Bead, b.Delta.x, b.Delta.y, b.IDsuffix, centerPosition, b.BeadColor, b.BeadRadius, gridSpace, clip)
		if b.Labeled {
			adaptLabelToVertex(&v, b.Bead.Beadtype, labelYDev, b.FontColor)
		}
		vertices[v.ID] = v
		pos := b.Bead.pos
		if b.Bead.PrevDir != int8(-1) {
			edges = append(edges, pathToEdge(pos, pos.NextPosition(b.Bead.PrevDir), b.IDsuffix, b.PathColor))
		}
		for i := uint8(0); i < uint8(6); i++ {
			if (b.Bead.StabconDir>>i)%2 == 1 {
				edges = append(edges, pathToEdge(pos, pos.NextPosition(int8(i)), b.IDsuffix, b.BondColor))
			}
		}
	}
	return treegraph{vertices, edges, det}
}

func PathToGraph(conformation oritatami.Conformation, stabilized oritatami.Position, routes Routeview, containRoute bool,
	clip string, margin TreeMargin, color TreeColor) (treegraph, float64) {

	/*
		createCircle := func(bead oritatami.Bead, idSuffix string, center oritatami.Position, beadcolor svgco.ShapeColor,
			dx float64, dy float64) svgco.GraphVertex {
			bp := bead.GetPosition()
			id := getID(bp, idSuffix)
		}
	*/

	beadsarray := []BeadFigure{}

	// ready to search a center and to make folding vertices

	var centerposition oritatami.Position
	var nodeCircleradius float64

	if containRoute {
		centerposition, nodeCircleradius = getCenterpositionAndViewRadius(&conformation, routes, stabilized, margin)
	} else {
		centerposition, nodeCircleradius = oritatami.Position{0, 0}, 0.0
	}

	// seed
	beadsarray = append(beadsarray, seedBeadFigures(conformation, margin, color)...)

	//formed

	beadsarray = append(beadsarray, formedBeadFigures(conformation, stabilized, margin, color)...)

	determin := false

	//folding1
	lbead, ok := conformation.GetBead(stabilized.GetPos())
	if containRoute && ok {
		// route detdir?
		if routes.detDir == lbead.NextDir && routes.isMaxbond {
			determin = true
		}
		beadsarray = append(beadsarray, nondetRouteBeadFigures(conformation, stabilized, routes, color, margin)...)
	}
	return BeadFigureToGraph(beadsarray, centerposition, determin, margin.GridSpace, margin.LabelYDev, clip), nodeCircleradius
}

func nondetRouteBeadFigures(conformation oritatami.Conformation, stabilized oritatami.Position, routes Routeview,
	color TreeColor, margin TreeMargin) []BeadFigure {

	//type BeadFigure struct
	//	Bead      oritatami.Bead
	//	BeadColor svgco.ShapeColor
	//	PathColor svgco.ShapeColor
	//	BondColor svgco.ShapeColor
	//	FontColor svgco.FontSvg

	answer := []BeadFigure{}

	//folding1
	nowposition := stabilized

	//fmt.Println("os2svg l.317", routes.later0)
	if routes.later0 > -1 || len(routes.routes) == 1 {
		route := routes.routes[0]
		for index, nextdir := range route.Path[:routes.later0+1] {
			nowposition = nowposition.NextPosition(nextdir)
			bead := BeadForView{routes.transcript[index], nowposition, int8(5 - nextdir), route.Bonddir[index]}

			beadfigure := BeadFigure{bead, FloatPosition{}, "", color.FoldingBead, margin.RoutesRadius, color.FoldingEdge, color.Bond, true, color.Font}

			answer = append(answer, beadfigure)
		}
	}

	//folding2
	matchingposition := nowposition
	if routes.later0 < len(routes.routes[0].Path)-1 {
		for index, r := range routes.routes {
			nowposition := matchingposition
			rand.Seed(time.Now().UnixNano())
			ang := rand.Float64() * 2 * math.Pi
			dist := rand.Float64() * margin.BeadRadius * 0.9 // !!
			delta := FloatPosition{math.Cos(ang) * dist, math.Sin(ang) * dist}
			rcolor := color.getRouteColor()
			vcolor := svgco.ShapeColor{false, "", 0, true, rcolor}
			ecolor := svgco.ShapeColor{true, rcolor, margin.MultEdgeWidth, false, ""}

			idsuffix := "r" + strconv.Itoa(index)

			lastFormedBead := BeadForView{"", nowposition, int8(-1), 0}
			lastFormedBeadfigure := BeadFigure{lastFormedBead, delta, idsuffix, vcolor, margin.MultBeadRadius, ecolor, color.Bond, false, color.Font}

			answer = append(answer, lastFormedBeadfigure)

			for _, nextdir := range r.Path[routes.later0+1:] {
				nowposition = nowposition.NextPosition(nextdir)

				bead := BeadForView{"", nowposition, int8(5 - nextdir), 0}
				beadfigure := BeadFigure{bead, delta, idsuffix, vcolor, margin.MultBeadRadius, ecolor, color.Bond, false, color.Font}

				answer = append(answer, beadfigure)
			}
		}
	}
	return answer
}

////////////////////////////////////////////
// oritatami parameters to view parameters
////////////////////////////////////////////

func osysBeadToView(bead oritatami.Bead) *BeadForView {
	return &BeadForView{bead.Beadtype, bead.GetPosition(), bead.PrevDir, bead.StabconDir}
}

func OsysRoutesToView(routes oritatami.Routes) []Routeview {
	//type Routeview struct {
	//	later0   int
	//	routes     []*oritatami.Route
	//	transcript []string

	//for _, rs := range routes.Routes {
	//	fmt.Println(rs.Path)
	//}

	answer := []Routeview{}

	typeOfLater0 := make(map[int][]oritatami.Route)

	for _, route := range routes.Routes {
		typeOfLater0[route.Later0()] = append(typeOfLater0[route.Later0()], route)
	}

	for later0, rs := range typeOfLater0 {
		answer = append(answer, convertToRouteview(later0, rs, routes.Transcript, routes.Mbonds)...)
	}
	return answer
}

func convertToRouteview(later0 int, routes []oritatami.Route, transcript []string, mbond uint) []Routeview {

	// bind routes
	type TRoutes struct {
		next     map[int8]*TRoutes
		terminal []oritatami.Route
	}

	detpath := func(terminals []oritatami.Route) int8 {
		firstpath := terminals[0].Path
		if len(firstpath) == 0 {
			return int8(-1)
		}
		detdir := firstpath[0]
		for _, term := range terminals[1:] {
			if len(term.Path) == 0 || term.Path[0] != detdir {
				return int8(-1)
			}
		}
		return detdir
	}
	//fmt.Println("os2svg l.442", len(routes[0].Path))
	root := TRoutes{make(map[int8]*TRoutes), nil}

	for _, route := range routes {
		pointer := &root
		for i := -1; i <= later0; i++ {
			if i == later0 {
				//fmt.Println("os2svg l.449", len(route.Path))
				pointer.terminal = append(pointer.terminal, route)
			} else {
				now, ok := pointer.next[route.Path[i+1]]
				if !ok {
					now = &TRoutes{make(map[int8]*TRoutes), nil}
					pointer.next[route.Path[i+1]] = now
				}
				pointer = now
			}
		}
	}

	ans := []Routeview{}

	queue := []*TRoutes{&root}

	for len(queue) > 0 {
		focus := queue[0]
		if focus.terminal == nil {
			tails := []*TRoutes{}
			for b := int8(0); b < 6; b++ {
				el, ok := focus.next[b]
				if ok {
					tails = append(tails, el)
				}
			}
			queue = append(queue[1:], tails...)
		} else {
			//fmt.Println("os2svg l.478", len(focus.terminal[2].Path))
			ismaxbond := false
			if len(focus.terminal) > 0 && focus.terminal[0].Bonds == mbond {
				ismaxbond = true
			}
			ans = append(ans, Routeview{later0, focus.terminal, transcript, detpath(focus.terminal), ismaxbond})
			queue = queue[1:]
		}
	}
	//fmt.Println("os2svg l.481", len(ans[0].routes[2].Path))
	return ans
}

/////////////
/////////////

func (t *TreeMargin) initialize(SpaceVertical float64, BeadRadius float64, grid float64,
	routeRadius float64, routeThich float64) {
	t.SpaceVertical = SpaceVertical
	//t.oneblockWidth = blockWidth
	t.BeadRadius = BeadRadius
	t.GridSpace = grid
	//t.EdgeThick = EdgeThick
	t.RoutesRadius = routeRadius
	//t.RouteEdgeThick = routeThich
}

/*
func (t *TreeColor) initialize(folding string, formed map[rune]string, seed string, edgewidth float64, vertexwidth float64, bond string) {
	//ShapeColor struct
	//	Stroke       bool
	//	Stroke_color string
	//	Stroke_width float64
	//	Fill         bool
	//	Fill_color   string
	t.FoldingEdge = svgco.ShapeColor{true, folding, edgewidth, false, ""}
	t.FoldingBead = svgco.ShapeColor{true, folding, vertexwidth, true, "#ffffff"}
	t.FormedEdge = make(map[rune]svgco.ShapeColor)
	t.FormedBead = make(map[rune]svgco.ShapeColor)
	for label, colorst := range formed {
		t.FormedEdge[label] = svgco.ShapeColor{true, colorst, edgewidth, false, ""}
		t.FormedBead[label] = svgco.ShapeColor{true, colorst, vertexwidth, true, "#ffffff"}
	}
	t.SeedEdge = svgco.ShapeColor{true, seed, edgewidth, false, ""}
	t.SeedBead = svgco.ShapeColor{true, seed, vertexwidth, true, "#ffffff"}
	t.Bond = svgco.ShapeColor{true, bond, edgewidth, false, ""}
	t.RouteColors = GetColorList()
}
*/

func NewBeadFigure(bead oritatami.Bead, beadcolor string, beadthick float64, paththick float64, bondcolor string, bondthick float64, font svgco.FontSvg, radius float64) BeadFigure {
	//Stroke       bool
	//Stroke_color string
	//Stroke_width float64
	//Fill         bool
	//Fill_color   string

	//Bead       BeadForView
	//Delta      FloatPosition
	//IDsuffix   string
	//BeadColor  svgco.ShapeColor
	//BeadRadius float64
	//PathColor  svgco.ShapeColor
	//BondColor  svgco.ShapeColor
	//Labeled    bool
	//FontColor  svgco.FontSvg
	return BeadFigure{*osysBeadToView(bead), FloatPosition{}, "", svgco.ShapeColor{true, beadcolor, beadthick, true, "#ffffff"}, radius,
		svgco.ShapeColor{true, beadcolor, paththick, false, ""}, svgco.ShapeColor{true, bondcolor, bondthick, false, ""}, true, font}
}

func GetColorList() []string {
	return []string{
		"#ff9393",
		"#ff93c9",
		"#ff93ff",
		"#c993ff",
		"#9393ff",
		"#93c9ff",
		"#93ffff",
		"#93ffc9",
		"#93ff93",
		"#c9ff93",
		"#ffff93",
		"#ffc993",
		"#ff2828",
		"#ff2893",
		"#ff28ff",
		"#9328ff",
		"#2828ff",
		"#2893ff",
		"#28ffff",
		"#28ff93",
		"#28ff28",
		"#93ff28",
		"#ffff28",
		"#ff9328"}
}

func (t *TreeColor) getRouteColor() string {
	ans := t.RouteColors[0]
	t.RouteColors = append(t.RouteColors[1:], t.RouteColors[0])
	return ans
}

func midpointAndMargin(positions []oritatami.Position, margin TreeMargin) (oritatami.Position, float64) {
	var miny int64
	var maxy int64
	var minx oritatami.Position
	var maxx oritatami.Position

	if len(positions) == 0 {
		fmt.Println("os2svg", "func", "midpointAndMargin", "len of positions is 0")
		return oritatami.Position{}, 0.0
	}

	maxy = positions[0].Y
	miny = maxy
	maxx = positions[0]
	minx = maxx

	detx := func(y int64) float64 {
		return float64(y) / math.Sqrt(3)
	}

	for _, p := range positions[1:] {
		if p.Y > maxy {
			maxy = p.Y
		}
		if p.Y < miny {
			miny = p.Y
		}
		if float64(p.X)+detx(p.Y) > float64(maxx.X)+detx(maxx.Y) {
			maxx = p
		}
		if float64(p.X)+detx(p.Y) < float64(minx.X)+detx(minx.Y) {
			minx = p
		}
	}

	midy := int64((miny + maxy) / 2)
	midx := int64((float64(minx.X+maxx.X)+float64(minx.Y)/2.0+float64(maxx.Y)/2.0)/2.0 - float64(midy)/2.0)

	marginRadius := max(max(max(math.Abs(float64(maxy-midy))*margin.GridSpace*math.Sqrt(3.0)/2.0,
		math.Abs(float64(miny-midy))*margin.GridSpace*math.Sqrt(3.0)/2.0),
		math.Abs(float64(minx.X-midx)+float64(minx.Y)/2.0)*margin.GridSpace),
		math.Abs(float64(maxx.X-midx)+float64(maxx.Y)/2.0)*margin.GridSpace)

	return oritatami.Position{midx, midy}, marginRadius
}

func max(arg1 float64, arg2 float64) float64 {
	if arg1 > arg2 {
		return arg1
	} else {
		return arg2
	}
}

func positionToPoint(position oritatami.Position, center oritatami.Position, gridSpace float64) (float64, float64) {
	// coordinate (x,y) -> view (w,h)
	px := (float64(position.X) + float64(position.Y)/2.0) * gridSpace
	py := float64(position.Y) * math.Sqrt(3.0) * gridSpace / 2.0
	cx := (float64(center.X) + float64(center.Y)/2.0) * gridSpace
	cy := float64(center.Y) * math.Sqrt(3.0) * gridSpace / 2.0

	return px - cx, py - cy
}

func getModuleLabel(beadtype string) rune {
	btint, err := strconv.Atoi(beadtype)
	if err == nil {
		if btint > 0 && btint <= 30 {
			return FM
		} else if btint > 30 && btint <= 66 {
			return LTM
		} else if btint > 66 && btint <= 96 {
			return HAM
		} else if btint > 96 && btint <= 132 {
			return RTM
		}
	}
	return ERR
}
