package figo

import (
	. "brick_automaton"
	"oritatami"
	"strconv"
	"svgco"
)

type OritatamiFigure struct {
	Dircarr string
	SeedW   oritatami.Transcript
}

func (of *OritatamiFigure) addpref(prefix string) {
	of.Dircarr = prefix + of.Dircarr
}

func BrickEnvToConf(prevCur string, prev string, above string, abovePrev string, delay int) oritatami.Conformation {
	//(e BrickElements) MakeSeed(prevCur string, prev string, above string, abovePrev string, osparam Parameters)
	belem := BrickElements{make(map[string]BrickModule)}
	belem.LoadModules()

	params := Parameters{ReadParameter(delay)}
	return belem.MakeSeed(prevCur, prev, above, abovePrev, params)
}

//////////////
///////

/*
func OsystemFoldEx1_0() []BeadFigure {
	t := oritatami.Transcript{[]string{"N", "N", "N", "N", "N", "N", "N", "N", "N", "B", "N", "N"}}
	seedconf := oritatami.MakeNewConformationFromDirectionToSeed("wwwwadwdemb", t, oritatami.Position{6, 0})
	bf := beadmapToBeadFigures(seedconf.Seed, "#525266", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})
	//type Bead
	//	Beadtype    string
	//	pos         Position
	//	NextDir     int8
	//	PrevDir     int8
	//	Nowconnect  uint
	//	Stabconnect uint
	//	StabconDir  uint8
	folding := []oritatami.Bead{
		oritatami.NewBeadca("1", 3, 1, 1, 4, 0, 0, 0),
		oritatami.NewBeadca("2", 4, 1, 1, 4, 0, 0, 0),
		oritatami.NewBeadca("3", 5, 1, 1, 4, 0, 0, 0)}
	bf = append(bf, beadsToBeadFigures(folding, "#000000", 4.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})...)

	return bf
}

func OsystemFoldEx1_1() []BeadFigure {
	t := oritatami.Transcript{[]string{"N", "N", "N", "N", "N", "N", "N", "N", "N", "B", "N", "N"}}
	seedconf := oritatami.MakeNewConformationFromDirectionToSeed("wwwwadwdemb", t, oritatami.Position{6, 0})
	bf := beadmapToBeadFigures(seedconf.Seed, "#525266", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})
	folding := []oritatami.Bead{
		oritatami.NewBeadca("1", 3, 1, 1, 4, 0, 0, 0),
		oritatami.NewBeadca("2", 4, 1, 5, 4, 0, 0, 0),
		oritatami.NewBeadca("3", 3, 2, -1, 0, 0, 0, 0)}
	bf = append(bf, beadsToBeadFigures(folding, "#000000", 4.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})...)

	return bf
}

func OsystemFoldEx1_2() []BeadFigure {
	t := oritatami.Transcript{[]string{"N", "N", "N", "N", "N", "N", "N", "N", "N", "B", "N", "N"}}
	seedconf := oritatami.MakeNewConformationFromDirectionToSeed("wwwwadwdemb", t, oritatami.Position{6, 0})
	bf := beadmapToBeadFigures(seedconf.Seed, "#525266", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})
	folding := []oritatami.Bead{
		oritatami.NewBeadca("1", 3, 1, 2, 4, 0, 0, 0),
		oritatami.NewBeadca("2", 3, 2, 5, 3, 0, 0, 0),
		oritatami.NewBeadca("3", 2, 3, -1, 0, 1, 1, 16)}
	bf = append(bf, beadsToBeadFigures(folding, "#000000", 4.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})...)

	return bf
}

func OsystemFoldEx1_3() []BeadFigure {
	t := oritatami.Transcript{[]string{"N", "N", "N", "N", "N", "N", "N", "N", "N", "B", "N", "N"}}
	seedconf := oritatami.MakeNewConformationFromDirectionToSeed("wwwwadwdemb", t, oritatami.Position{6, 0})
	bf := beadmapToBeadFigures(seedconf.Seed, "#525266", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})
	formed := []oritatami.Bead{
		oritatami.NewBeadca("1", 3, 1, 2, 4, 0, 0, 0)}
	bf = append(bf, beadsToBeadFigures(formed, "#f59342", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})...)
	folding := []oritatami.Bead{
		oritatami.NewBeadca("2", 3, 2, 5, 3, 0, 0, 0),
		oritatami.NewBeadca("3", 2, 3, 2, 0, 1, 1, 16),
		oritatami.NewBeadca("4", 2, 4, -1, 3, 0, 0, 0)}
	bf = append(bf, beadsToBeadFigures(folding, "#000000", 4.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})...)

	return bf
}

func OsystemFoldEx1_4() []BeadFigure {
	t := oritatami.Transcript{[]string{"N", "N", "N", "N", "N", "N", "N", "N", "N", "B", "N", "N"}}
	seedconf := oritatami.MakeNewConformationFromDirectionToSeed("wwwwadwdemb", t, oritatami.Position{6, 0})
	bf := beadmapToBeadFigures(seedconf.Seed, "#525266", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})
	formed := []oritatami.Bead{
		oritatami.NewBeadca("1", 3, 1, 2, 4, 0, 0, 0),
		oritatami.NewBeadca("2", 3, 2, 5, 3, 0, 0, 0),
		oritatami.NewBeadca("3", 2, 3, 1, 0, 1, 1, 16),
		oritatami.NewBeadca("4", 3, 3, 0, 4, 0, 0, 0),
		oritatami.NewBeadca("5", 4, 2, 3, 5, 0, 0, 16),
		oritatami.NewBeadca("6", 4, 1, 1, 2, 1, 1, 48),
		oritatami.NewBeadca("7", 5, 1, 2, 4, 0, 0, 0),
		oritatami.NewBeadca("8", 5, 2, 5, 3, 0, 0, 0),
		oritatami.NewBeadca("9", 4, 3, -1, 0, 1, 1, 16)}
	bf = append(bf, beadsToBeadFigures(formed, "#f59342", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})...)

	return bf
}
func OsystemFoldEx2_0() []BeadFigure {
	t := oritatami.Transcript{[]string{"B", "B", "N", "N", "N", "N", "N", "N", "N", "B", "N", "N"}}
	seedconf := oritatami.MakeNewConformationFromDirectionToSeed("wwwwadwdemb", t, oritatami.Position{6, 0})
	bf := beadmapToBeadFigures(seedconf.Seed, "#525266", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})
	adj := []oritatami.Bead{
		oritatami.NewBeadca("B", 6, 0, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("B", 5, 0, -1, -1, 0, 0, 0)}
	bf = append(bf, beadsToBeadFigures(adj, "#b3153d", 3.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#b3153d"})...)

	folding := []oritatami.Bead{
		oritatami.NewBeadca("1", 3, 1, 2, 4, 0, 0, 0),
		oritatami.NewBeadca("2", 3, 2, 5, 3, 0, 0, 0),
		oritatami.NewBeadca("3", 2, 3, -1, 0, 1, 1, 16)}
	bf = append(bf, beadsToBeadFigures(folding, "#000000", 4.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})...)

	return bf
}
func OsystemFoldEx2_1() []BeadFigure {
	t := oritatami.Transcript{[]string{"B", "B", "N", "N", "N", "N", "N", "N", "N", "B", "N", "N"}}
	seedconf := oritatami.MakeNewConformationFromDirectionToSeed("wwwwadwdemb", t, oritatami.Position{6, 0})
	bf := beadmapToBeadFigures(seedconf.Seed, "#525266", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})
	adj := []oritatami.Bead{
		oritatami.NewBeadca("B", 6, 0, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("B", 5, 0, -1, -1, 0, 0, 0)}
	bf = append(bf, beadsToBeadFigures(adj, "#b3153d", 3.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#b3153d"})...)

	folding := []oritatami.Bead{
		oritatami.NewBeadca("1", 3, 1, 1, 4, 0, 0, 0),
		oritatami.NewBeadca("2", 4, 1, 1, 4, 0, 0, 0),
		oritatami.NewBeadca("3", 5, 1, -1, 4, 2, 2, 9)}
	bf = append(bf, beadsToBeadFigures(folding, "#000000", 4.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})...)

	return bf
}

func OsystemFoldEx2_2() []BeadFigure {
	t := oritatami.Transcript{[]string{"B", "B", "N", "N", "N", "N", "N", "N", "N", "B", "N", "N"}}
	seedconf := oritatami.MakeNewConformationFromDirectionToSeed("wwwwadwdemb", t, oritatami.Position{6, 0})
	bf := beadmapToBeadFigures(seedconf.Seed, "#525266", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})
	adj := []oritatami.Bead{
		oritatami.NewBeadca("B", 6, 0, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("B", 5, 0, -1, -1, 0, 0, 0)}
	bf = append(bf, beadsToBeadFigures(adj, "#b3153d", 3.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#b3153d"})...)

	formed := []oritatami.Bead{
		oritatami.NewBeadca("1", 3, 1, 2, 4, 0, 0, 0)}
	bf = append(bf, beadsToBeadFigures(formed, "#f59342", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})...)
	folding := []oritatami.Bead{
		oritatami.NewBeadca("2", 4, 1, 1, 4, 0, 0, 0),
		oritatami.NewBeadca("3", 5, 1, 1, 4, 2, 2, 9),
		oritatami.NewBeadca("4", 6, 1, -1, 4, 0, 0, 0)}
	bf = append(bf, beadsToBeadFigures(folding, "#000000", 4.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})...)

	return bf
}

func OsystemFoldEx2_3() []BeadFigure {
	t := oritatami.Transcript{[]string{"B", "B", "N", "N", "N", "N", "N", "N", "N", "B", "N", "N"}}
	seedconf := oritatami.MakeNewConformationFromDirectionToSeed("wwwwadwdemb", t, oritatami.Position{6, 0})
	bf := beadmapToBeadFigures(seedconf.Seed, "#525266", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})
	adj := []oritatami.Bead{
		oritatami.NewBeadca("B", 6, 0, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("B", 5, 0, -1, -1, 0, 0, 0)}
	bf = append(bf, beadsToBeadFigures(adj, "#b3153d", 3.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#b3153d"})...)

	formed := []oritatami.Bead{
		oritatami.NewBeadca("1", 3, 1, 1, 4, 0, 0, 0),
		oritatami.NewBeadca("2", 4, 1, 1, 4, 0, 0, 0),
		oritatami.NewBeadca("3", 5, 1, 1, 4, 2, 2, 9),
		oritatami.NewBeadca("4", 5, 2, 1, 3, 0, 0, 0),
		oritatami.NewBeadca("5", 4, 2, 1, 1, 2, 2, 8),
		oritatami.NewBeadca("6", 3, 2, 1, 1, 0, 0, 9),
		oritatami.NewBeadca("7", 2, 3, 1, 0, 2, 2, 16),
		oritatami.NewBeadca("8", 3, 3, 1, 4, 0, 0, 0),
		oritatami.NewBeadca("9", 4, 3, -1, 4, 2, 2, 1)}
	bf = append(bf, beadsToBeadFigures(formed, "#f59342", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})...)

	return bf
}
func OsystemFoldEx3_0() []BeadFigure {
	t := oritatami.Transcript{[]string{"B", "N", "N", "N", "N", "N", "N", "N", "N", "B", "N", "N"}}
	seedconf := oritatami.MakeNewConformationFromDirectionToSeed("wwwwadwdemb", t, oritatami.Position{6, 0})
	bf := beadmapToBeadFigures(seedconf.Seed, "#525266", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})
	adj := []oritatami.Bead{
		oritatami.NewBeadca("B", 6, 0, -1, -1, 0, 0, 0)}
	bf = append(bf, beadsToBeadFigures(adj, "#b3153d", 3.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#b3153d"})...)

	folding := []oritatami.Bead{
		oritatami.NewBeadca("1", 3, 1, 2, 4, 0, 0, 0),
		oritatami.NewBeadca("2", 3, 2, 5, 3, 0, 0, 0),
		oritatami.NewBeadca("3", 2, 3, -1, 0, 1, 1, 16)}
	bf = append(bf, beadsToBeadFigures(folding, "#000000", 4.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})...)

	return bf
}
func OsystemFoldEx3_1() []BeadFigure {
	t := oritatami.Transcript{[]string{"B", "N", "N", "N", "N", "N", "N", "N", "N", "B", "N", "N"}}
	seedconf := oritatami.MakeNewConformationFromDirectionToSeed("wwwwadwdemb", t, oritatami.Position{6, 0})
	bf := beadmapToBeadFigures(seedconf.Seed, "#525266", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})
	adj := []oritatami.Bead{
		oritatami.NewBeadca("B", 6, 0, -1, -1, 0, 0, 0)}
	bf = append(bf, beadsToBeadFigures(adj, "#b3153d", 3.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#b3153d"})...)

	folding := []oritatami.Bead{
		oritatami.NewBeadca("1", 3, 1, 1, 4, 0, 0, 0),
		oritatami.NewBeadca("2", 4, 1, 1, 4, 0, 0, 0),
		oritatami.NewBeadca("3", 5, 1, -1, 4, 2, 2, 1)}
	bf = append(bf, beadsToBeadFigures(folding, "#000000", 4.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})...)

	return bf
}
*/
//////////

/*
func FigH00_m() []BeadFigure {
	t := oritatami.Transcript{[]string{"127", "132", "1", "10", "11", "12", "13"}}
	seedconf := oritatami.MakeNewConformationFromDirectionToSeed("wwwwadwdemb", t, oritatami.Position{6, 0})
	bf := beadmapToBeadFigures(seedconf.Seed, "#525266", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})
	adj := []oritatami.Bead{
		oritatami.NewBeadca("B", 6, 0, -1, -1, 0, 0, 0),
		oritatami.NewBeadca("B", 5, 0, -1, -1, 0, 0, 0)}
	bf = append(bf, beadsToBeadFigures(adj, "#b3153d", 3.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#b3153d"})...)

	formed := []oritatami.Bead{
		oritatami.NewBeadca("1", 3, 1, 1, 4, 0, 0, 0),
		oritatami.NewBeadca("2", 4, 1, 1, 4, 0, 0, 0),
		oritatami.NewBeadca("3", 5, 1, 1, 4, 2, 2, 9),
		oritatami.NewBeadca("4", 5, 2, 1, 3, 0, 0, 0),
		oritatami.NewBeadca("5", 4, 2, 1, 1, 2, 2, 8),
		oritatami.NewBeadca("6", 3, 2, 1, 1, 0, 0, 9),
		oritatami.NewBeadca("7", 2, 3, 1, 0, 2, 2, 16),
		oritatami.NewBeadca("8", 3, 3, 1, 4, 0, 0, 0),
		oritatami.NewBeadca("9", 4, 3, -1, 4, 2, 2, 1)}
	bf = append(bf, beadsToBeadFigures(formed, "#f59342", 2.0, 5.0, "#5e9185", 3.0, svgco.FontSvg{"Super Sans", 8.0, "#000000"})...)

	return bf
}
*/
///////////
func FigCounter1() (OritatamiFigure, oritatami.Transcript) {
	Ltrc := LtrcSeed()
	Hn := HnSeed()
	Hn.addpref("e")
	Rb := RbSeed()
	Rb.addpref("e")
	F0 := F0Seed()
	F0.addpref("e")
	Lbc := LbcSeed()
	Lbc.addpref("e")
	Rtr := RtrSeed()
	Rtr.addpref("e")

	tints := []int{}
	for i := 0; i < 132; i++ {
		tints = append(tints, i+1)
	}

	for i := 0; i < 132; i++ {
		tints = append(tints, i+1)
	}
	for i := 0; i < 132; i++ {
		tints = append(tints, i+1)
	}
	for i := 0; i < 132; i++ {
		tints = append(tints, i+1)
	}
	for i := 0; i < 32; i++ {
		tints = append(tints, i+1)
	}
	for i := 32; i < 132; i++ {
		tints = append(tints, i+1)
	}
	for i := 0; i < 68; i++ {
		tints = append(tints, i+1)
	}

	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return appendfigure(Ltrc, Hn, Rb, F0, Lbc, Hn, Rtr), transcript
}

func NondetEx() (OritatamiFigure, oritatami.Transcript) {
	Rb := RbSeed()
	F0 := F0Seed()
	F0.addpref("e")
	Lt := LtSeedPrevZig()
	Lt.addpref("d")

	tints := []int{}
	for i := 66; i < 75; i++ {
		tints = append(tints, i+1)
	}
	tints = append(tints, 1, 2, 3, 4)
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return appendfigure(Rb, F0, Lt), transcript
}

func FigH00() (OritatamiFigure, oritatami.Transcript) {
	Rb := RbSeed()
	F0 := F0Seed()
	F0.addpref("e")
	Lt := LtSeedPrevZig()
	Lt.addpref("d")

	tints := []int{}
	for i := 66; i < 98; i++ {
		tints = append(tints, i+1)
	}
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return appendfigure(Rb, F0, Lt), transcript
}

func FigH01() (OritatamiFigure, oritatami.Transcript) {
	Rb := RbSeed()
	F0 := F0Seed()
	F0.addpref("e")
	Lbc := LbcSeedPrevZig()
	Lbc.addpref("d")

	tints := []int{}
	for i := 66; i < 98; i++ {
		tints = append(tints, i+1)
	}
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return appendfigure(Rb, F0, Lbc), transcript
}

func FigH10() (OritatamiFigure, oritatami.Transcript) {
	Rb := RbSeed()
	F1 := F1Seed()
	F1.addpref("e")
	Lt := LtSeedPrevZig()
	Lt.addpref("d")

	tints := []int{}
	for i := 66; i < 98; i++ {
		tints = append(tints, i+1)
	}
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return appendfigure(Rb, F1, Lt), transcript
}

func FigH11() (OritatamiFigure, oritatami.Transcript) {
	Rb := RbSeed()
	F1 := F1Seed()
	F1.addpref("e")
	Lbc := LbcSeedPrevZig()
	Lbc.addpref("d")

	tints := []int{}
	for i := 66; i < 98; i++ {
		tints = append(tints, i+1)
	}
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return appendfigure(Rb, F1, Lbc), transcript
}

func FigHe1() (OritatamiFigure, oritatami.Transcript) {
	Lbe := LbeSeedPrevZig()

	tints := []int{}
	for i := 66; i < 98; i++ {
		tints = append(tints, i+1)
	}
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return Lbe, transcript
}

func FigHn() (OritatamiFigure, oritatami.Transcript) {
	Fnb := FnbSeedForZagWithRb()
	Lbc := LbcSeedPrevZag()
	Lbc.addpref("a")

	tints := []int{}
	for i := 66; i < 98; i++ {
		tints = append(tints, i+1)
	}
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return appendfigure(Fnb, Lbc), transcript
}

func FigLt() (OritatamiFigure, oritatami.Transcript) {
	F1 := F1Seed()
	Lbc := LbcSeed()
	Lbc.addpref("e")
	Fnt := FntSeedPrevZig()
	Fnt.addpref("d")

	tints := []int{}
	for i := 30; i < 68; i++ {
		tints = append(tints, i+1)
	}
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return appendfigure(F1, Lbc, Fnt), transcript
}

func FigLtrc() (OritatamiFigure, oritatami.Transcript) {
	Ltrc := LtrcSeed()
	Fnt := FntSeedPrevZig2()
	Fnt.addpref("d")

	tints := []int{}
	for i := 30; i < 68; i++ {
		tints = append(tints, i+1)
	}
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return appendfigure(Ltrc, Fnt), transcript
}

func FigLbe() (OritatamiFigure, oritatami.Transcript) {
	Ltrc := LtrcSeed()
	Fnb := FnbSeedPrevZig()
	Fnb.addpref("d")

	tints := []int{}
	for i := 30; i < 68; i++ {
		tints = append(tints, i+1)
	}
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return appendfigure(Ltrc, Fnb), transcript
}

func FigLbc() (OritatamiFigure, oritatami.Transcript) {
	Fnb := FnbSeed()
	Lbc := LbcSeed()
	Lbc.addpref("e")
	Fnbpre := FnbSeedPrevZig()
	Fnbpre.addpref("d")

	tints := []int{}
	for i := 30; i < 68; i++ {
		tints = append(tints, i+1)
	}
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return appendfigure(Fnb, Lbc, Fnbpre), transcript
}

func FigLtre() (OritatamiFigure, oritatami.Transcript) {
	Rb := ConfRb()
	Fnb := ConfFnb()
	Fnb.addpref("w")

	tints := []int{}
	for i := 30; i < 68; i++ {
		tints = append(tints, i+1)
	}
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return appendfigure(Rb, Fnb), transcript
}

func FigF0() (OritatamiFigure, oritatami.Transcript) {
	Lbc := LbcSeed()
	Lbc.Dircarr = "wwwww" + "wwwww" + "w"
	H11 := H11SeedWithRt()
	H11.addpref("w")
	Rb := RbSeedPrevZag()
	Rb.addpref("a")

	tints := []int{}
	for i := 0; i < 32; i++ {
		tints = append(tints, i+1)
	}
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return appendfigure(Lbc, H11, Rb), transcript
}

func FigF1() (OritatamiFigure, oritatami.Transcript) {
	Lbc := LbcSeed()
	Lbc.Dircarr = "wwwww" + "wwwww" + "w"
	H01 := H01SeedWithRt()
	H01.addpref("w")
	Rb := RbSeedPrevZag()
	Rb.addpref("a")

	tints := []int{}
	for i := 0; i < 32; i++ {
		tints = append(tints, i+1)
	}
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return appendfigure(Lbc, H01, Rb), transcript
}

func FigFnb() (OritatamiFigure, oritatami.Transcript) {
	Hn := HnSeedForZigWithLbcRb()
	Rb := RbSeedPrevZig()
	Rb.addpref("d")

	tints := []int{}
	for i := 0; i < 32; i++ {
		tints = append(tints, i+1)
	}
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return appendfigure(Hn, Rb), transcript
}

func FigFnt() (OritatamiFigure, oritatami.Transcript) {
	Hn := HnSeedForZigWithLbcRb()
	Rt := RtSeedPrevZig()
	Rt.addpref("a")

	tints := []int{}
	for i := 0; i < 32; i++ {
		tints = append(tints, i+1)
	}
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return appendfigure(Hn, Rt), transcript
}

func FigRb() (OritatamiFigure, oritatami.Transcript) {
	He1pre := He1SeedPrevZig()

	tints := []int{}
	for i := 96; i < 132; i++ {
		tints = append(tints, i+1)
	}
	tints = append(tints, 1, 2)
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return He1pre, transcript
}

func FigRt() (OritatamiFigure, oritatami.Transcript) {
	Rb := RbSeedForZigWithHnFnb()
	H01pre := H01SeedPrevZig()
	H01pre.addpref("d")

	tints := []int{}
	for i := 96; i < 132; i++ {
		tints = append(tints, i+1)
	}
	tints = append(tints, 1, 2)
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return appendfigure(Rb, H01pre), transcript
}

func FigRtr() (OritatamiFigure, oritatami.Transcript) {
	Rtr := RtrSeed()
	Hn := HnSeedPrevZag()
	Rtr.Dircarr += "w"
	Rtr.SeedW.Arr = append(Rtr.SeedW.Arr, "1")
	Hn.addpref("a")

	tints := []int{}
	for i := 96; i < 132; i++ {
		tints = append(tints, i+1)
	}
	tints = append(tints, 1, 2)
	transcript := oritatami.Transcript{Ints2Strings(tints)}

	return appendfigure(Rtr, Hn), transcript
}

//////////////////////////////////////////

func LtrcSeed() OritatamiFigure {
	Dircarr := "eed" + "wwaee"
	sequence := []string{"58", "59", "60", "61", "62", "63", "64", "65", "66"}
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func LtSeedPrevZig() OritatamiFigure {
	Dircarr := "adwbm"
	sequence := Ints2Strings([]int{61, 62, 63, 64, 65, 66})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func LbcSeedPrevZig() OritatamiFigure {
	Dircarr := "waeedww"
	sequence := Ints2Strings([]int{59, 60, 61, 62, 63, 64, 65, 66})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func LbcSeedPrevZag() OritatamiFigure {
	Dircarr := "edwwaee"
	sequence := Ints2Strings([]int{59, 60, 61, 62, 63, 64, 65, 66})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func LbeSeedPrevZig() OritatamiFigure {
	Dircarr := "bmwadwbmwad"
	sequence := Ints2Strings([]int{55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func FntSeedPrevZig() OritatamiFigure {
	Dircarr := "wddwb"
	sequence := Ints2Strings([]int{25, 26, 27, 28, 29, 30})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func FntSeedPrevZig2() OritatamiFigure {
	Dircarr := "wdedwwb"
	sequence := Ints2Strings([]int{25, 26, 27, 23, 22, 28, 29, 30})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func FnbSeedPrevZig() OritatamiFigure {
	Dircarr := "dabbad"
	sequence := Ints2Strings([]int{27, 23, 25, 26, 28, 29, 30})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func HnSeed() OritatamiFigure {
	Dircarr := "eeee" + "eeeee"
	sequence := Ints2Strings([]int{67, 76, 77, 78, 79, 88, 89, 90, 91, 96})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func HnSeedPrevZag() OritatamiFigure {
	Dircarr := "daembeda"
	sequence := Ints2Strings([]int{84, 85, 90, 91, 92, 93, 94, 95, 96})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func He1SeedPrevZig() OritatamiFigure {
	Dircarr := "adwbmwad"
	sequence := Ints2Strings([]int{88, 89, 90, 91, 92, 93, 94, 95, 96})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func H01SeedPrevZig() OritatamiFigure {
	Dircarr := "waedwwbm"
	sequence := Ints2Strings([]int{80, 81, 82, 83, 92, 93, 94, 95, 96})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func H01SeedWithRt() OritatamiFigure {
	Dircarr := "wwwww" + "wwwww"
	sequence := Ints2Strings([]int{67, 72, 73, 88, 89, 90, 91, 92, 93, 94, 99})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func H11SeedWithRt() OritatamiFigure {
	Dircarr := "wwwww" + "wwwww"
	sequence := Ints2Strings([]int{67, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func HnSeedForZigWithLbcRb() OritatamiFigure {
	Dircarr := "eeeee" + "eeeee" + "eeee"
	sequence := Ints2Strings([]int{55, 64, 65, 66, 67, 76, 77, 78, 79, 88, 89, 90, 91, 96, 97})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func RbSeedForZigWithHnFnb() OritatamiFigure {
	Dircarr := "eeeee" + "eeeee" + "eeeee"
	sequence := Ints2Strings([]int{90, 91, 96, 97, 102, 103, 108, 109, 114, 115, 120, 121, 126, 127, 132, 1})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func RbSeed() OritatamiFigure {
	Dircarr := "e" + "eeeee" + "eeeee"
	sequence := Ints2Strings([]int{97, 102, 103, 108, 109, 114, 115, 120, 121, 126, 127, 132})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func RbSeedPrevZag() OritatamiFigure {
	Dircarr := "daembeda"
	sequence := Ints2Strings([]int{124, 125, 126, 127, 128, 129, 130, 131, 132})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func RbSeedPrevZig() OritatamiFigure {
	Dircarr := oritatami.FlipDirection("daembeda")
	sequence := Ints2Strings([]int{124, 125, 126, 127, 128, 129, 130, 131, 132})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func RtrSeed() OritatamiFigure {
	Dircarr := "eeedwwwwd" + "abbad"
	sequence := Ints2Strings([]int{97, 108, 109, 120, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func RtSeedPrevZig() OritatamiFigure {
	Dircarr := "adwbm"
	sequence := Ints2Strings([]int{127, 128, 129, 130, 131, 132})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

/*
func RtSeedPrevZig() OritatamiFigure {
	Dircarr := "adwbm"
	sequence := Ints2Strings([]int{127, 128, 129, 130, 131, 132})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}
*/

func FnbSeed() OritatamiFigure {
	Dircarr := "eeee" + "eeeee"
	sequence := Ints2Strings([]int{1, 6, 7, 12, 13, 18, 19, 24, 25, 30})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func FnbSeedForZagWithRb() OritatamiFigure {
	Dircarr := "wwwww" + "wwwww" + "www"
	sequence := Ints2Strings([]int{121, 126, 127, 132, 1, 6, 7, 12, 13, 18, 19, 24, 25, 30})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func F0Seed() OritatamiFigure {
	Dircarr := "eeee" + "eeeee"
	sequence := Ints2Strings([]int{1, 10, 11, 12, 13, 22, 23, 24, 25, 30})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func F1Seed() OritatamiFigure {
	Dircarr := "eeee" + "eeeee"
	sequence := Ints2Strings([]int{1, 22, 23, 24, 25, 26, 27, 28, 29, 30})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func LbcSeed() OritatamiFigure {
	Dircarr := "e" + "eeeee" + "eeeee"
	sequence := Ints2Strings([]int{31, 36, 37, 42, 43, 48, 49, 54, 55, 64, 65, 66})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func ConfFnb() OritatamiFigure {
	Dircarr := "bmwadw" + "bmwadw" + "bmwadw" + "bmwadw" + "bmwad"
	sequence := Ints2Strings([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30})
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func ConfRb() OritatamiFigure {
	Dircarr := "bmwadw" + "bmwadw" + "bmwadw" + "bmwadw" + "bmwadw" + "bmwad"
	sq := []int{}
	for i := 96; i < 132; i++ {
		sq = append(sq, i+1)
	}
	sequence := Ints2Strings(sq)
	return OritatamiFigure{Dircarr, oritatami.Transcript{sequence}}
}

func Ints2Strings(ints []int) []string {
	ans := []string{}
	for _, e := range ints {
		ans = append(ans, strconv.Itoa(e))
	}
	return ans
}

func appendfigure(fhead OritatamiFigure, elems ...OritatamiFigure) OritatamiFigure {
	ans := OritatamiFigure{fhead.Dircarr, fhead.SeedW}
	for _, el := range elems {
		ans.Dircarr += el.Dircarr
		ans.SeedW.Arr = append(ans.SeedW.Arr, el.SeedW.Arr...)
	}
	return ans
}

func beadmapToBeadFigures(beadmap oritatami.BeadMap, beadcolor string, beadthick float64,
	paththick float64, bondcolor string, bondthick float64, font svgco.FontSvg, radius float64) []BeadFigure {
	beads := []oritatami.Bead{}
	for _, e1 := range beadmap.Getmap() {
		for _, e2 := range e1 {
			//func NewBeadFigure(bead oritatami.Bead, beadcolor string,
			// beadthick float64, paththick float64, bondcolor string, bondthick float64, font svgco.FontSvg) BeadFigure {
			beads = append(beads, e2)
		}
	}
	return beadsToBeadFigures(beads, beadcolor, beadthick, paththick, bondcolor, bondthick, font, radius)
}

func beadsToBeadFigures(beads []oritatami.Bead, beadcolor string, beadthick float64,
	paththick float64, bondcolor string, bondthick float64, font svgco.FontSvg, radius float64) []BeadFigure {
	ans := []BeadFigure{}
	for _, e := range beads {
		ans = append(ans, NewBeadFigure(e, beadcolor, beadthick, paththick, bondcolor, bondthick, font, radius))
	}
	return ans
}

func MakeColor() TreeColor {
	//Stroke       bool
	//Stroke_color string
	//Stroke_width float64
	//Fill         bool
	//Fill_color   string

	//onlyf := func(fill string) svgco.ShapeColor {
	//	return svgco.ShapeColor{false, "", 0, true, fill}
	//}
	mclr := func(stroke string, thick float64, fill string) svgco.ShapeColor {
		if len(fill) == 0 {
			return svgco.ShapeColor{true, stroke, thick, false, ""}
		} else {
			return svgco.ShapeColor{true, stroke, thick, true, fill}
		}
	}
	return TreeColor{
		mclr("#000000", 2.0, ""),
		map[rune]svgco.ShapeColor{RTM: mclr("#ffc14f", 5.0, ""), LTM: mclr("#4fd0ff", 5.0, ""),
			FM: mclr("#b8ed8e", 5.0, ""), HAM: mclr("#ed8e9c", 5.0, "")},
		map[rune]svgco.ShapeColor{}, //seed
		mclr("#000000", 2.0, "#ffffff"),
		map[rune]svgco.ShapeColor{RTM: mclr("#ffc14f", 2.0, "#ffffff"), LTM: mclr("#4fd0ff", 2.0, "#ffffff"),
			FM: mclr("#b8ed8e", 2.0, "#ffffff"), HAM: mclr("#ed8e9c", 2.0, "#ffffff")},
		map[rune]svgco.ShapeColor{},
		mclr("#5e9185", 3.0, ""),
		svgco.FontSvg{"Super Sans", 8.0, "#000000"},
		GetColorList(),
		"#b21e1e",
		mclr("#665782", 5.0, ""),
		mclr("#665782", 2.0, "#ffffff"),
		mclr("#b8ed8e", 5.0, ""),
		mclr("#b8ed8e", 2.0, "#ffffff")}
	/*
	   type FontSvg struct {
	   	Font_family string
	   	Font_size   float64
	   	Fill_color  string
	   }*/
	/*
	  type TreeColor struct {
	  	FoldingEdge svgco.ShapeColor
	  	FormedEdge  map[rune]svgco.ShapeColor
	  	SeedEdge    svgco.ShapeColor
	  	FoldingBead svgco.ShapeColor
	  	FormedBead  map[rune]svgco.ShapeColor
	  	SeedBead    svgco.ShapeColor
	  	Bond        svgco.ShapeColor
	  	Font        svgco.FontSvg
	  	RouteColors []string
	  	DecideNode  string
	*/
}

func MakeColorWithSeedColored() TreeColor {
	//Stroke       bool
	//Stroke_color string
	//Stroke_width float64
	//Fill         bool
	//Fill_color   string

	//onlyf := func(fill string) svgco.ShapeColor {
	//	return svgco.ShapeColor{false, "", 0, true, fill}
	//}
	mclr := func(stroke string, thick float64, fill string) svgco.ShapeColor {
		if len(fill) == 0 {
			return svgco.ShapeColor{true, stroke, thick, false, ""}
		} else {
			return svgco.ShapeColor{true, stroke, thick, true, fill}
		}
	}
	return TreeColor{
		mclr("#000000", 2.0, ""),
		map[rune]svgco.ShapeColor{RTM: mclr("#ffc14f", 5.0, ""), LTM: mclr("#4fd0ff", 5.0, ""),
			FM: mclr("#b8ed8e", 5.0, ""), HAM: mclr("#ed8e9c", 5.0, "")},
		map[rune]svgco.ShapeColor{RTM: mclr("#ffc14f", 5.0, ""), LTM: mclr("#4fd0ff", 5.0, ""),
			FM: mclr("#b8ed8e", 5.0, ""), HAM: mclr("#ed8e9c", 5.0, "")},
		mclr("#000000", 2.0, "#ffffff"),
		map[rune]svgco.ShapeColor{RTM: mclr("#ffc14f", 2.0, "#ffffff"), LTM: mclr("#4fd0ff", 2.0, "#ffffff"),
			FM: mclr("#b8ed8e", 2.0, "#ffffff"), HAM: mclr("#ed8e9c", 2.0, "#ffffff")},
		map[rune]svgco.ShapeColor{RTM: mclr("#ffc14f", 2.0, "#ffffff"), LTM: mclr("#4fd0ff", 2.0, "#ffffff"),
			FM: mclr("#b8ed8e", 2.0, "#ffffff"), HAM: mclr("#ed8e9c", 2.0, "#ffffff")},
		mclr("#5e9185", 3.0, ""),
		svgco.FontSvg{"Super Sans", 8.0, "#000000"},
		GetColorList(),
		"#b21e1e",
		mclr("#665782", 5.0, ""),
		mclr("#665782", 2.0, "#ffffff"),
		mclr("#b8ed8e", 5.0, ""),
		mclr("#b8ed8e", 2.0, "#ffffff")}
	/*
	   type FontSvg struct {
	   	Font_family string
	   	Font_size   float64
	   	Fill_color  string
	   }*/
	/*
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
	*/
}

func MakeColor2() TreeColor {
	mclr := func(stroke string, thick float64, fill string) svgco.ShapeColor {
		if len(fill) == 0 {
			return svgco.ShapeColor{true, stroke, thick, false, ""}
		} else {
			return svgco.ShapeColor{true, stroke, thick, true, fill}
		}
	}
	return TreeColor{
		mclr("#000000", 2.0, ""),
		map[rune]svgco.ShapeColor{RTM: mclr("#ffc14f", 5.0, ""), LTM: mclr("#4fd0ff", 5.0, ""),
			FM: mclr("#b8ed8e", 5.0, ""), HAM: mclr("#ed8e9c", 5.0, "")},
		map[rune]svgco.ShapeColor{}, //seed
		mclr("#000000", 2.0, "#ffffff"),
		map[rune]svgco.ShapeColor{RTM: mclr("#ffc14f", 2.0, "#ffffff"), LTM: mclr("#4fd0ff", 2.0, "#ffffff"),
			FM: mclr("#b8ed8e", 2.0, "#ffffff"), HAM: mclr("#ed8e9c", 2.0, "#ffffff")},
		map[rune]svgco.ShapeColor{},
		mclr("#90d1c2", 1.5, ""),
		svgco.FontSvg{"Super Sans", 8.0, "#000000"},
		GetColorList(),
		"#b21e1e",
		mclr("#665782", 5.0, ""),
		mclr("#665782", 2.0, "#ffffff"),
		mclr("#b8ed8e", 5.0, ""),
		mclr("#b8ed8e", 2.0, "#ffffff")}
}

func MakeMargin() TreeMargin {
	return TreeMargin{
		200.0,
		10.0,
		30.0,
		3.5, //label y dev
		10.0,
		2.0,
		1.0,

		10.0,
		2.5,
		7.0,
		3.0,
		//2.0,
		//25.0,
		//15.0,
		130.0,

		3,
		3,
		2}
	/*
	  type TreeMargin struct {
	  	spaceVertical  float64
	  	beadRadius     float64
	  	gridSpace      float64
	  	edgeThick      float64
	  	routesRadius   float64
	  	routeEdgeThick float64
	  	multBeadRadius float64
	  	multEdgeWidth  float64

	  	nodeCorner         float64
	  	decideNodeThick0   float64
	  	decideNodeThick1   float64
	  	notDecideNodeThick float64
	  	nodeSubFrameThick  float64
	  	nodeSubFrameWidth  float64
	  	nodeSubFrameHeight float64

	  	positionPrec int
	  	strokePrec   int
	  	fontsizePrec int
	  }
	*/
}
