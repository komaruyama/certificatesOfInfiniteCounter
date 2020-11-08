package figo

import (
	"modules/pkg/oritatami"
	"modules/pkg/svgco"
	. "modules/pkg/brick_automaton"
)

func NodeBglider() []BeadFigure {
	radius := 10.0

	//func beadmapToBeadFigures(beadmap oritatami.BeadMap, beadcolor string, beadthick float64,
	//	paththick float64, bondcolor string, bondthick float64, font svgco.FontSvg, radius float64) []BeadFigure
	Seed := []oritatami.Bead{
		oritatami.NewBeadca("", 0, 0, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("", 0, -1, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("", 1, -2, -1, -1, 0, 0, 0)}

	bf := beadsToBeadFigures(Seed, "#78bbcc", 2.0, 5.0, "#78bbcc", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)
	//type Bead
	//	Beadtype    string
	//	pos         Position
	//	NextDir     int8
	//	PrevDir     int8
	//	Nowconnect  uint
	//	Stabconnect uint
	//	StabconDir  uint8

	tsunagi := []oritatami.Bead{oritatami.NewBeadca("a", -1, 0, -1, 1, 0, 0, 0)}

	t := oritatami.Transcript{[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}}
	nodeconf := oritatami.MakeNewConformationFromDirectionToSeed("bmwadwbmwad", t, oritatami.Position{-1, 0})
	bf = append(bf, beadsToBeadFigures(tsunagi, "#000000", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)...)
	bf = append(bf, beadmapToBeadFigures(nodeconf.Seed, "#000000", 4.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)...)

	return bf
}

func NodeCliff() []BeadFigure {
	radius := 10.0

	//func beadmapToBeadFigures(beadmap oritatami.BeadMap, beadcolor string, beadthick float64,
	//	paththick float64, bondcolor string, bondthick float64, font svgco.FontSvg, radius float64) []BeadFigure
	Seed := []oritatami.Bead{
		oritatami.NewBeadca("", 0, 0, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("", 0, -1, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("", 1, -2, -1, -1, 0, 0, 0)}

	bf := beadsToBeadFigures(Seed, "#78bbcc", 2.0, 5.0, "#78bbcc", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)
	//type Bead
	//	Beadtype    string
	//	pos         Position
	//	NextDir     int8
	//	PrevDir     int8
	//	Nowconnect  uint
	//	Stabconnect uint
	//	StabconDir  uint8

	tsunagi := []oritatami.Bead{oritatami.NewBeadca("a", -1, 0, -1, 1, 0, 0, 0)}

	t := oritatami.Transcript{[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}}
	nodeconf := oritatami.MakeNewConformationFromDirectionToSeed(oritatami.FlipDirection("mbeeedwwaee"), t, oritatami.Position{-1, 0})
	bf = append(bf, beadsToBeadFigures(tsunagi, "#000000", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)...)
	bf = append(bf, beadmapToBeadFigures(nodeconf.Seed, "#000000", 4.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)...)

	return bf
}

func NodeFault() []BeadFigure {
	radius := 10.0

	//func beadmapToBeadFigures(beadmap oritatami.BeadMap, beadcolor string, beadthick float64,
	//	paththick float64, bondcolor string, bondthick float64, font svgco.FontSvg, radius float64) []BeadFigure
	Seed := []oritatami.Bead{
		oritatami.NewBeadca("", 0, 0, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("", 0, -1, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("", 1, -2, -1, -1, 0, 0, 0)}

	bf := beadsToBeadFigures(Seed, "#78bbcc", 2.0, 5.0, "#78bbcc", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)
	//type Bead
	//	Beadtype    string
	//	pos         Position
	//	NextDir     int8
	//	PrevDir     int8
	//	Nowconnect  uint
	//	Stabconnect uint
	//	StabconDir  uint8

	tsunagi := []oritatami.Bead{oritatami.NewBeadca("a", -1, 0, -1, 1, 0, 0, 0)}

	t := oritatami.Transcript{[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}}
	nodeconf := oritatami.MakeNewConformationFromDirectionToSeed(oritatami.FlipDirection("mbeeeee"), t, oritatami.Position{-1, 0})
	bf = append(bf, beadsToBeadFigures(tsunagi, "#000000", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)...)
	bf = append(bf, beadmapToBeadFigures(nodeconf.Seed, "#000000", 4.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)...)

	return bf
}

func NodeHorn() []BeadFigure {
	radius := 10.0

	//func beadmapToBeadFigures(beadmap oritatami.BeadMap, beadcolor string, beadthick float64,
	//	paththick float64, bondcolor string, bondthick float64, font svgco.FontSvg, radius float64) []BeadFigure
	Seed := []oritatami.Bead{
		oritatami.NewBeadca("", 0, 0, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("", 0, -1, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("", 1, -2, -1, -1, 0, 0, 0)}

	bf := beadsToBeadFigures(Seed, "#78bbcc", 2.0, 5.0, "#78bbcc", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)
	//type Bead
	//	Beadtype    string
	//	pos         Position
	//	NextDir     int8
	//	PrevDir     int8
	//	Nowconnect  uint
	//	Stabconnect uint
	//	StabconDir  uint8

	tsunagi := []oritatami.Bead{oritatami.NewBeadca("a", -1, 0, -1, 1, 0, 0, 0)}

	t := oritatami.Transcript{[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}}
	nodeconf := oritatami.MakeNewConformationFromDirectionToSeed(oritatami.FlipDirection("mbmmeeawa"), t, oritatami.Position{-1, 0})
	bf = append(bf, beadsToBeadFigures(tsunagi, "#000000", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)...)
	bf = append(bf, beadmapToBeadFigures(nodeconf.Seed, "#000000", 4.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)...)

	return bf
}

func NodeKar() []BeadFigure {
	radius := 10.0

	//func beadmapToBeadFigures(beadmap oritatami.BeadMap, beadcolor string, beadthick float64,
	//	paththick float64, bondcolor string, bondthick float64, font svgco.FontSvg, radius float64) []BeadFigure
	Seed := []oritatami.Bead{
		oritatami.NewBeadca("", 0, 0, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("", 0, -1, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("", 1, -2, -1, -1, 0, 0, 0)}

	bf := beadsToBeadFigures(Seed, "#78bbcc", 2.0, 5.0, "#78bbcc", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)
	//type Bead
	//	Beadtype    string
	//	pos         Position
	//	NextDir     int8
	//	PrevDir     int8
	//	Nowconnect  uint
	//	Stabconnect uint
	//	StabconDir  uint8

	tsunagi := []oritatami.Bead{oritatami.NewBeadca("a", 1, -3, -1, 2, 0, 0, 0)}

	t := oritatami.Transcript{[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}}
	nodeconf := oritatami.MakeNewConformationFromDirectionToSeed("wbmwadde", t, oritatami.Position{1, -3})
	bf = append(bf, beadsToBeadFigures(tsunagi, "#000000", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)...)
	bf = append(bf, beadmapToBeadFigures(nodeconf.Seed, "#000000", 4.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)...)

	return bf
}

func NodePlane() []BeadFigure {
	radius := 10.0

	//func beadmapToBeadFigures(beadmap oritatami.BeadMap, beadcolor string, beadthick float64,
	//	paththick float64, bondcolor string, bondthick float64, font svgco.FontSvg, radius float64) []BeadFigure
	Seed := []oritatami.Bead{
		oritatami.NewBeadca("", 0, 0, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("", 0, -1, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("", 1, -2, -1, -1, 0, 0, 0)}

	bf := beadsToBeadFigures(Seed, "#78bbcc", 2.0, 5.0, "#78bbcc", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)
	//type Bead
	//	Beadtype    string
	//	pos         Position
	//	NextDir     int8
	//	PrevDir     int8
	//	Nowconnect  uint
	//	Stabconnect uint
	//	StabconDir  uint8

	tsunagi := []oritatami.Bead{oritatami.NewBeadca("a", 0, -2, -1, 1, 0, 0, 0)}

	t := oritatami.Transcript{[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}}
	nodeconf := oritatami.MakeNewConformationFromDirectionToSeed("wwwww", t, oritatami.Position{0, -2})
	bf = append(bf, beadsToBeadFigures(tsunagi, "#000000", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)...)
	bf = append(bf, beadmapToBeadFigures(nodeconf.Seed, "#000000", 4.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)...)

	return bf
}

func NodeRoof() []BeadFigure {
	radius := 10.0

	//func beadmapToBeadFigures(beadmap oritatami.BeadMap, beadcolor string, beadthick float64,
	//	paththick float64, bondcolor string, bondthick float64, font svgco.FontSvg, radius float64) []BeadFigure
	Seed := []oritatami.Bead{
		oritatami.NewBeadca("", 0, 0, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("", 0, -1, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("", 1, -2, -1, -1, 0, 0, 0)}

	bf := beadsToBeadFigures(Seed, "#78bbcc", 2.0, 5.0, "#78bbcc", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)
	//type Bead
	//	Beadtype    string
	//	pos         Position
	//	NextDir     int8
	//	PrevDir     int8
	//	Nowconnect  uint
	//	Stabconnect uint
	//	StabconDir  uint8

	tsunagi := []oritatami.Bead{oritatami.NewBeadca("a", -1, 0, -1, 1, 0, 0, 0)}

	t := oritatami.Transcript{[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}}
	nodeconf := oritatami.MakeNewConformationFromDirectionToSeed("bmmeeee", t, oritatami.Position{-1, 0})
	bf = append(bf, beadsToBeadFigures(tsunagi, "#000000", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)...)
	bf = append(bf, beadmapToBeadFigures(nodeconf.Seed, "#000000", 4.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)...)

	return bf
}

func NodeSpiral() []BeadFigure {
	radius := 10.0

	//func beadmapToBeadFigures(beadmap oritatami.BeadMap, beadcolor string, beadthick float64,
	//	paththick float64, bondcolor string, bondthick float64, font svgco.FontSvg, radius float64) []BeadFigure
	Seed := []oritatami.Bead{
		oritatami.NewBeadca("", 0, 0, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("", 0, -1, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("", 1, -2, -1, -1, 0, 0, 0)}

	bf := beadsToBeadFigures(Seed, "#78bbcc", 2.0, 5.0, "#78bbcc", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)
	//type Bead
	//	Beadtype    string
	//	pos         Position
	//	NextDir     int8
	//	PrevDir     int8
	//	Nowconnect  uint
	//	Stabconnect uint
	//	StabconDir  uint8

	tsunagi := []oritatami.Bead{oritatami.NewBeadca("a", -1, 0, -1, 1, 0, 0, 0)}

	t := oritatami.Transcript{[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}}
	nodeconf := oritatami.MakeNewConformationFromDirectionToSeed("bmwadwbmwww", t, oritatami.Position{-1, 0})
	bf = append(bf, beadsToBeadFigures(tsunagi, "#000000", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)...)
	bf = append(bf, beadmapToBeadFigures(nodeconf.Seed, "#000000", 4.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)...)

	return bf
}

func NodeTglider() []BeadFigure {
	radius := 10.0

	//func beadmapToBeadFigures(beadmap oritatami.BeadMap, beadcolor string, beadthick float64,
	//	paththick float64, bondcolor string, bondthick float64, font svgco.FontSvg, radius float64) []BeadFigure
	Seed := []oritatami.Bead{
		oritatami.NewBeadca("", 0, 0, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("", 0, -1, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("", 1, -2, -1, -1, 0, 0, 0)}

	bf := beadsToBeadFigures(Seed, "#78bbcc", 2.0, 5.0, "#78bbcc", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)
	//type Bead
	//	Beadtype    string
	//	pos         Position
	//	NextDir     int8
	//	PrevDir     int8
	//	Nowconnect  uint
	//	Stabconnect uint
	//	StabconDir  uint8

	tsunagi := []oritatami.Bead{oritatami.NewBeadca("a", 0, -2, -1, 1, 0, 0, 0)}

	t := oritatami.Transcript{[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}}
	nodeconf := oritatami.MakeNewConformationFromDirectionToSeed("adwbmwadwbm", t, oritatami.Position{0, -2})
	bf = append(bf, beadsToBeadFigures(tsunagi, "#000000", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)...)
	bf = append(bf, beadmapToBeadFigures(nodeconf.Seed, "#000000", 4.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"}, radius)...)

	return bf
}
