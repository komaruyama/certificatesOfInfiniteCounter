package brick_automaton

import (
	"strconv"
	"strings"

	"github.com/komaruyama/certificatesOfInfiniteCounter/pkg/oritatami"
	"github.com/komaruyama/certificatesOfInfiniteCounter/pkg/svgco"
)

const (
	Empty = "e"
)

const (
	Top    = 't'
	Middle = 'm'
	Bottom = 'b'
)

const (
	TurnToZig = "turntozig"
	TurnToZag = "turntozag"
	Zig       = "zig"
	Zag       = "zag"
)

const (
	HAM = 'h'
	FM  = 'f'
	LTM = 'l'
	RTM = 'r'
	ERR = 'e'
)

type BrickGraph struct {
	Ef  map[BrickEnv][]BrickEnv
	El  map[BrickEnv][]BrickEnv
	Eha map[BrickEnv][]BrickEnv
	Er  map[BrickEnv][]BrickEnv
}

type BrickEnv struct {
	Elem   string
	Prev   string
	Upper  string
	Upprev string
}

type BrickVertex struct {
	ModuleType   rune
	Current      string
	Elem         string
	Prev         string
	Abov         string
	Abpr         string
	InputConf    string
	OutputConf   string
	conformation *oritatami.Conformation
	OutOSConf    *oritatami.Conformation
}

type BrickGraphVE struct {
	Vertices       map[string]BrickVertex
	Elements       BrickElements
	Edges          map[string][]BrickVertex
	IOConformation BrickElements
}

type BrickElements struct {
	Elements map[string]BrickModule
}

type BrickModule struct {
	Label         string
	InputCarry    rune
	OutputCarry   rune
	ComeFrom      rune
	Path          string
	DirectionType string
	//Shape oritatami.Conformation
}

type BrickError struct {
	Dirc string
	env  BrickVertex
}

type Parameters struct {
	Params map[rune]oritatami.Parameter
}

type BeadType struct {
	index2string map[int]string
	string2index map[string]int
}

func NextModuleType(cur rune) rune {
	switch cur {
	case FM:
		return LTM
	case LTM:
		return HAM
	case HAM:
		return RTM
	case RTM:
		return FM
	}
	return ERR
}

/*
func AboveModuleType(cur rune) rune {
	switch cur {
	case FM:
		return HAM
	case LTM:
		return LTM
	case HAM:
		return FM
	case RTM:
		return RTM
	}
	return ERR
}
*/

///////////////////
//  BrickModule
///////////////////
func (m BrickModule) GetPathWithComeFrom() string {
	return string(m.ComeFrom) + m.Path
}

func (m BrickModule) GetType() rune {
	//fmt.Println("GetType", m.ComeFrom, m.Path, m.DirectionType)
	t := m.Label[0]
	switch t {
	case 'F':
		return FM
	case 'H':
		return HAM
	case 'L':
		return LTM
	case 'R':
		return RTM
	}
	return ERR
}

/////////////////
//  BeadType
////////////////

func (bt *BeadType) initialize() {
	bt.index2string = make(map[int]string)
	bt.string2index = make(map[string]int)
}

func (bt *BeadType) append(typeindex int, typelabel string) {
	bt.index2string[typeindex] = typelabel
	bt.string2index[typelabel] = typeindex
}

func (bt BeadType) label(index int) string {
	ans, _ := bt.index2string[index]
	return ans
}

func (bt BeadType) ApplyIndicesToStrings(transcript []string) []string {
	ans := []string{}
	for _, e := range transcript {
		l, err := strconv.Atoi(strings.TrimSpace(e))
		if err == nil {
			s, ok := bt.index2string[l]
			if !ok {
				return nil
			}
			ans = append(ans, s)
		}
	}
	return ans
}

///////////////////
//  BrickGraphVE
///////////////////

func (g BrickGraphVE) getDirctionType(elementLabel string) (string, bool) {
	m, ok := g.Elements.Elements[elementLabel]
	if ok {
		return m.DirectionType, true
	}
	return "", false
}

/////////////////
/// TreeColor ///
/////////////////

func (c TreeColor) GetFormedEdgeColor(label rune) svgco.ShapeColor {
	color, ok := c.FormedEdge[label]
	if ok {
		return color
	} else {
		return c.DefaultFormedEdge
	}
}

func (c TreeColor) GetFormedBeadColor(label rune) svgco.ShapeColor {
	color, ok := c.FormedBead[label]
	if ok {
		return color
	} else {
		return c.DefaultFormedBead
	}
}

func (c TreeColor) GetSeedEdgeColor(label rune) svgco.ShapeColor {
	color, ok := c.SeedEdge[label]
	if ok {
		return color
	} else {
		return c.DefaultSeedEdge
	}
}

func (c TreeColor) GetSeedBeadColor(label rune) svgco.ShapeColor {
	color, ok := c.SeedBead[label]
	if ok {
		return color
	} else {
		return c.DefaultSeedBead
	}
}

/////////////////
/// Routeview ///
/////////////////

func (r Routeview) GetRoutes() []oritatami.Route {
	return r.routes
}

///////////////////
/// BrickVertex ///
///////////////////

func (v BrickVertex) equal(v2 BrickVertex) bool {
	if v.Abov == v2.Abov && v.Abpr == v2.Abpr && v.Current == v2.Current && v.Elem == v2.Elem &&
		v.InputConf == v2.InputConf && v.ModuleType == v2.ModuleType && v.OutputConf == v2.OutputConf && v.Prev == v2.Prev {
		return true
	}
	return false
}

func (v BrickVertex) IDlabel() string {
	return v.Abov + " " + v.Abpr + " " + v.Current + " " + v.Elem + " " + v.InputConf + " " + string(v.ModuleType) + " " + v.OutputConf + " " + v.Prev
}
