package brick_automaton

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/komaruyama/certificatesOfInfiniteCounter/pkg/oritatami"
)

const (
	delimiter             = ";" // 区切り文字
	commentout            = "#" //
	separateParamAndConst = ":"
)

func ParseParamAndConstBasedOnDef(lines []string) map[string]string {
	line := ""
	for _, el := range lines {
		line += el
	}
	spt1 := strings.Split(line, delimiter)

	parsed := make(map[string]string)

	for _, onesente := range spt1[:len(spt1)-1] {
		sentence := strings.TrimSpace(onesente)
		if string(sentence[0]) != commentout {
			spt2 := strings.SplitN(sentence, separateParamAndConst, 2)
			if len(spt2) == 2 {
				paramst := strings.TrimSpace(spt2[0])
				constst := strings.TrimSpace(spt2[1])
				if len(paramst) > 0 && len(constst) > 0 {
					parsed[paramst] = constst
				}
			} else {
				paramst := strings.TrimSpace(spt2[0])
				if len(paramst) > 0 {
					parsed[paramst] = ""
				}
			}
		}
	}
	return parsed
}

func ParseIOConformation() []BrickModule {
	modules := []BrickModule{}

	path := "oritatami system/conformation/node"
	files, err := ioutil.ReadDir(path)
	//fmt.Println("l.19 ", len(files))
	if err != nil {
		fmt.Println("err at ParseIOConformation function")
		return nil
	}

	for _, file := range files {
		if !file.IsDir() {
			lines := Readfile(filepath.Join(path, file.Name()))
			if module, ok := parseIOConformationLine(lines); ok {
				modules = append(modules, module)
			}
		}
	}
	return modules
}

func ParseModule() []BrickModule {
	modules := []BrickModule{}

	path := "oritatami system/conformation/direction"
	files, err := ioutil.ReadDir(path)
	//fmt.Println("l.19 ", len(files))
	if err != nil {
		fmt.Println("err at ParseModule function")
		return nil
	}

	for _, file := range files {
		if !file.IsDir() {
			label := strings.Split(file.Name(), ".")[0]
			lines := Readfile(filepath.Join(path, file.Name()))
			if module, ok := parseModuleLine(lines, label); ok {
				modules = append(modules, module)
			}
		}
	}
	return modules
}

func ReadParameter(delay int) map[rune]oritatami.Parameter {
	path := "oritatami system"
	rulepath := filepath.Join(path, "ruleset.aos")
	tpath := filepath.Join(path, "transcript.prm")
	adaptpath := filepath.Join(path, "adapt.prm")

	tline := toOneline(Readfile(tpath))

	//beadtype
	beadtype := ReadAdapter(Readfile(adaptpath))

	//rule
	rule := ReadRule(Readfile(rulepath), beadtype)

	/// t: f -> l -> ha -> r -> ...
	transcriptIndices := strings.Split(tline, ",")
	transcript := beadtype.ApplyIndicesToStrings(transcriptIndices)
	tf := oritatami.Transcript{}
	tl := oritatami.Transcript{}
	th := oritatami.Transcript{}
	tr := oritatami.Transcript{}
	tf.Initialize()
	tl.Initialize()
	th.Initialize()
	tr.Initialize()
	tf.Arr = append(transcript[0:30], transcript[30:32]...)
	tl.Arr = append(transcript[30:66], transcript[66:68]...)
	th.Arr = append(transcript[66:96], transcript[96:98]...)
	tr.Arr = append(transcript[96:132], transcript[0:2]...)

	return map[rune]oritatami.Parameter{
		FM:  oritatami.Parameter{delay, 5, rule, tf},
		LTM: oritatami.Parameter{delay, 5, rule, tl},
		HAM: oritatami.Parameter{delay, 5, rule, th},
		RTM: oritatami.Parameter{delay, 5, rule, tr},
	}
}

func ReadAdapter(readingLines []string) BeadType {
	beadtype := BeadType{}
	beadtype.initialize()

	adaptline := toOneline(readingLines)
	adaptelms := strings.Split(string(adaptline[strings.Index(adaptline, "[")+1:strings.Index(adaptline, "]")]), "(")[1:]
	for _, e := range adaptelms {
		es := strings.Split(e, ",")

		i, err := strconv.Atoi(strings.TrimSpace(es[0]))
		if err == nil {
			s := strings.TrimSpace(strings.Split(es[1], ")")[0])
			beadtype.append(i, s)
		}
	}
	return beadtype
}

func ReadRule(readingLines []string, beadtype BeadType) oritatami.Ruleset {
	rule := oritatami.Ruleset{}
	rule.Initialize()

	for index, l := range readingLines {
		if strings.HasPrefix(strings.TrimSpace(l), "+Rule") {
			oneline := toOneline(readingLines[index-1:])
			oneline = string(oneline[:strings.Index(oneline, ";")])
			parseRule(oneline, beadtype, &rule)
			break
		}
	}
	return rule
}

func parseIOConformationLine(lines []string) (BrickModule, bool) {
	var label string
	var path string

	oklabel := false
	okpath := false

	params := ParseParamAndConstBasedOnDef(lines)
	for param, con := range params {
		switch param {
		case "label":
			label = con
			oklabel = true
		case "dir":
			path = con
			okpath = true
		}
	}
	if oklabel && okpath {
		return BrickModule{label, 'e', 'e', 'e', path, ""}, true
	}
	return BrickModule{}, false
}

func parseModuleLine(lines []string, label string) (BrickModule, bool) {
	/*
		reading := []string{}
		for _, li := range lines {
			li = strings.TrimSpace(li)
			if len(li) > 1 && li[0] == '&' {
				reading = append(reading, string(li[1:]))
			}
		}

		var turn bool
		var dirc string
		var in rune
		var out rune

		if _, ok := searchPrefix(reading, "turn"); ok {
			turn = true
		} else {
			turn = false
		}
		first, ok1 := searchPrefix(reading, "in:")
		dir, ok2 := searchPrefix(reading, "dir:")
		ein, ok3 := searchPrefix(reading, "start:")
		eout, ok4 := searchPrefix(reading, "end:")

		if !(ok1 && ok2 && ok3 && ok4) {
			return BrickModule{}, false
		}

		if firsttrim := strings.TrimSpace(first); len(firsttrim) > 0 {
			first = string(firsttrim[0])
		} else {
			fmt.Println("loader.go l.111.. Nothing in:.. label:")
			return BrickModule{}, false
		}
		dir = strings.TrimSpace(strings.Split(dir, ";")[0])
		//dirc = first + dir

		in = []rune(strings.TrimSpace(ein))[0]
		out = []rune(strings.TrimSpace(eout))[0]

		return BrickModule{label, in, out, first, dir, turn}, true
	*/
	var inputc rune
	var outputc rune
	var comefrom rune
	var path string
	var directiontype string

	okin := false
	okout := false
	okcome := false
	okpath := false
	okdirtype := false

	params := ParseParamAndConstBasedOnDef(lines)
	for param, con := range params {
		switch param {
		case "&start":
			inputc = rune(con[0])
			okin = true
		case "&end":
			outputc = rune(con[0])
			okout = true
		case "&in":
			comefrom = rune(con[0])
			okcome = true
		case "&dir":
			path = con
			okpath = true
		case "&turn":
			if con == "toZig" {
				directiontype = TurnToZig
				okdirtype = true
			} else if con == "toZag" {
				directiontype = TurnToZag
				okdirtype = true
			}
		case "&zig":
			directiontype = Zig
			okdirtype = true
		case "&zag":
			directiontype = Zag
			okdirtype = true
		case "&toZig":
			directiontype = TurnToZig
			okdirtype = true
		case "&toZag":
			directiontype = TurnToZag
			okdirtype = true
		}
	}
	if okin && okout && okcome && okpath && okdirtype {
		return BrickModule{label, inputc, outputc, comefrom, path, directiontype}, true
	}
	return BrickModule{}, false
}

func parseRule(line string, adapt BeadType, rule *oritatami.Ruleset) {
	div1 := strings.Split(line, ":")
	for _, w := range div1 {
		ib1s := strings.Index(w, "-")
		ib1e := strings.Index(w, "=")
		ib2ars := strings.Index(w, "[")
		ib2are := strings.Index(w, "]")
		if ib1s != -1 && ib1e != -1 && ib1s < ib1e &&
			ib2ars != -1 && ib2are != -1 && ib2ars < ib2are {
			b1s := w[ib1s+1 : ib1e]
			b2sar := strings.Split(w[ib2ars+1:ib2are], ",")
			b1s = strings.TrimSpace(b1s)

			b1, ok1 := strconv.Atoi(b1s)

			if ok1 == nil {
				for _, b2s := range b2sar {
					b2, ok2 := strconv.Atoi(strings.TrimSpace(b2s))
					if ok2 == nil {
						s1, ok3 := adapt.index2string[b1]
						s2, ok4 := adapt.index2string[b2]
						if ok3 && ok4 {
							rule.Add(s1, s2)
						}
					}
				}
			}
		}
	}
}

//////////////////
//////////////////

/*
func TopORBottom(c oritatami.Conformation, prevOut rune) (rune, rune) {
	fbead, _ := c.Seed.GetBead(c.LastSeed.GetPos())
	fpos := c.LastSeed.NextPosition(fbead.NextDir)

	maxy := fpos.Y
	miny := maxy
}
*/

////////////////

func Readfile(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", path, err)
		os.Exit(1)
	}
	defer f.Close()

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if serr := scanner.Err(); serr != nil {
		fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", path, err)
	}

	return lines
}

func OutputSvg(lines []string, filename string) {
	lfname := "output/" + strings.TrimSpace(filename) + ".svg"
	os.Mkdir("output", 0777)
	file, err := os.Create(lfname)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, e := range lines {
		file.Write(([]byte)(e + "\n"))
	}
}

func toOneline(lines []string) string {
	ans := ""
	for _, l := range lines {
		ans += l
	}
	return ans
}

func searchPrefix(lines []string, prefix string) (string, bool) {
	for _, li := range lines {
		li = strings.TrimSpace(li)
		if strings.HasPrefix(li, prefix) {
			return string(li[len(prefix):]), true
		}
	}
	return "", false
}
