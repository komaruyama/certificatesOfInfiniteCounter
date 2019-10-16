package oritatami

import (
	"fmt"
	"sort"
	"strconv"
)

const (
	err_out_of_index      = 1
	err_no_space          = 2
	err_non_deterministic = 3
)

type error_route struct {
	nondet     []Route
	errCode    int
	transcript Transcript
	delay      int
}

type Parameter struct {
	Delay      int
	Arity      int
	Rule       Ruleset
	Transcript Transcript
}

type Route struct {
	Path      []int8
	Bonds     uint
	Eachbonds []uint
	Bonddir   []uint8
	Detpos    Position
}

type Routes struct {
	Routes     []Route
	Transcript []string
	Mbonds     uint
}

func FoldWithoutArity(startConformation Conformation, delay int, transcript Transcript, ruleset Ruleset) (Conformation, bool) {
	startConformation.ReadyExecuting()
	ans := true

	//fmt.Println("os debug t:", transcript.Arr)
	//debug
	//p := Parameter{delay, 5, ruleset, transcript}
	//p.viewParameter(false)

	for len(transcript.Arr) >= 3 { // !!
		startConformation.nextRouteMap(transcript)
		det, err := startConformation.FoldOneSetp(delay, transcript, ruleset)
		if !det {
			ans = false
			fmt.Println("err", err.errCode)
			//debug
			//err.viewError()
			break
		}
		transcript.Arr = transcript.Arr[1:]
	}
	return startConformation, ans
}

func (c *Conformation) FoldOneSetp(delay int, transcript Transcript, ruleset Ruleset) (bool, error_route) {
	// "transcript" has to start from a beadtype of stabilized bead
	var routes []Route
	lw, flag0 := transcript.GetLocalTranscript(0, delay)
	if !flag0 {
		// out of index (transcript)
		return false, error_route{nil, err_out_of_index, transcript, delay}
	}

	searchSpace := func(w Transcript, lastdir int8) (Route, bool) {
		path := []int8{}
		count := uint(0)
		bondslist := []uint{}
		bonddir := []uint8{}
		//var firstPos Position
		ans := Route{path, count, bondslist, bonddir, Position{}}

		//npos := c.LastFormed

		ok0 := c.putNextBead(&ans, lastdir, w.Arr[0], ruleset)
		if ok0 {
			ans.Detpos = c.LastFormed
		} else {
			return Route{}, false
		}

		if len(w.Arr) == 1 {
			return ans, true
		}

		for _, beadtype := range w.Arr[1:] {
			//npos := c.LastFormed
			ok := c.putNextBead(&ans, noDir, beadtype, ruleset)
			if !ok {
				return ans, true
			}
		}
		return ans, true
	}

	nroute, flag1 := searchSpace(lw, noDir)
	c.appendNewRoute(nroute)
	if !flag1 {
		// no space
		return false, error_route{nil, err_no_space, transcript, delay}
	}

	routes = []Route{nroute}

	//↑初期route
	//ここから全方向へ
	nextRouteAndBond := func(prevroute Route, w Transcript) (Route, bool) {
		route := prevroute.clone()
		nextpath := route.Path
		count := route.Bonds
		nextbondslist := route.Eachbonds
		nextbonddir := route.Bonddir

		for i := len(route.Path) - 1; i >= 0; i-- {
			dir := nextpath[i]
			nextpath = nextpath[:i]
			count -= c.RemoveLastFormd(ruleset)
			nextbondslist = nextbondslist[:i]
			nextbonddir = nextbonddir[:i]

			locT, _ := w.GetLocalTranscript(len(nextpath), delay) // "delay" is maximum
			tailRoute, f := searchSpace(locT, dir)
			if f {
				if i == 0 {
					return tailRoute, true
				} else {
					ans := Route{nextpath, count, nextbondslist, nextbonddir, route.Detpos}
					ans.AppendToTail(tailRoute)
					return ans, true
				}
			}
		}

		/* delete 19-06-21
		count -= c.RemoveLastFormd(ruleset)
		dir := route.Path[0]
		locT, _ := w.GetLocalTranscript(0, delay)
		ans, fl := searchSpace(locT, dir)
		if fl {
			return ans, true
		}
		return Route{}, false
		*/
		return Route{}, false
	}

	var cont bool
	//debC := 0
	nroute, cont = nextRouteAndBond(nroute, lw)
	c.appendNewRoute(nroute)
	for cont /* && debC < delay*5 */ {
		//debug
		//nroute.viewRoute()
		//debC++

		if routes[0].Bonds < nroute.Bonds {
			routes = []Route{nroute}
			//fmt.Println("debug.. replace nroute # of bonds =", strconv.Itoa(int(nroute.Bonds)), "dir:", nroute.Path[0], "pos:", nroute.Detpos.X, nroute.Detpos.Y)
		} else if routes[0].Bonds == nroute.Bonds {
			routes = append(routes, nroute)
			//fmt.Println("debug.. add nroute # of bonds =", strconv.Itoa(int(nroute.Bonds)), "dir:", nroute.Path[0], "pos:", nroute.Detpos.X, nroute.Detpos.Y)
		}

		nroute, cont = nextRouteAndBond(nroute, lw)
		c.appendNewRoute(nroute)
	}

	//checking deterministic
	if len(routes) > 1 {
		pos := routes[0].Detpos
		for _, r := range routes[1:] {
			if r.Detpos != pos {
				// delete temp beads
				for i := 0; i < len(nroute.Path); i++ {
					c.RemoveLastFormd(ruleset)
				}
				// non-deterministic
				return false, error_route{routes, err_non_deterministic, transcript, delay}
			}
		}
	}

	//debug
	/*
		stabilizedNotice := func(route Route) {
			fmt.Print("++stabilized ", transcript.Arr[0], " at: (")
			fmt.Print(route.Detpos.GetPos())
			fmt.Println(")")
		}
	*/

	// delete temp beads without stabilized bead

	//for i := 1; i < len(nroute.Path); i++ {
	//	c.RemoveLastFormd(ruleset)
	//}

	//apply
	c.Extend(routes[0].Path[0], transcript.Arr[0], routes[0].Eachbonds[0], routes[0].Bonddir[0])

	//stabilizedNotice(routes[0])
	return true, error_route{}

	//path   []int8
	//bonds  int
	//detpos Position
}

func (c *Conformation) putNextBead(route *Route, lastdir int8, beadtype string, ruleset Ruleset) bool {
	/*
		b, _ := c.GetBead(npos.GetPos())
		path = append(path, b.NextDir)
		count += n
		bondslist = append(bondslist, n)
		firstPos = c.LastFormed
	*/

	for d := int8(lastdir + 1); d < 6; d++ {
		nextpos := c.LastFormed.NextPosition(d)

		//cont debug
		if /*debug*/ _, fl := c.Formed.GetBead(nextpos.GetPos()); !fl {
			pBead, _ := c.Formed.GetBead(c.LastFormed.GetPos())
			pBead.NextDir = d
			//fmt.Println("pBead:(", pBead.pos.X, pBead.pos.Y, ")next:", d, debug1) //debug
			c.Formed.PutBead(pBead)

			nextBead := Bead{beadtype, nextpos, noDir, int8(5 - d), 0, 0, 0}
			connection, bonddir := c.connectionAdapt(nextBead, ruleset, true)
			nextBead.Nowconnect = connection
			nextBead.Stabconnect = connection

			c.Formed.PutBead(nextBead)
			c.ChangeLastFormed(nextpos)

			route.Path = append(route.Path, d)
			route.Bonds += connection
			route.Eachbonds = append(route.Eachbonds, connection)
			route.Bonddir = append(route.Bonddir, bonddir)
			return true
		} else {
			//debug
			//px := strconv.Itoa(int(debug.GetPosition().X))
			//py := strconv.Itoa(int(debug.GetPosition().Y))
			//fmt.Println("debug: there is a bead", debug.Beadtype, "at", "("+px, py+")")
		}
	}
	return false
}

func (c *Conformation) connectionAdapt(focusbead Bead, ruleset Ruleset, increase bool) (uint, uint8) {
	// if rule contains between focusbead and its neighbors, this function changes # of connection of around beads.
	var pm int
	var dirofBond uint8
	dirofBond = uint8(0)
	nofAdapt := uint(0)
	if increase {
		pm = 1
	} else {
		pm = -1
	}

	//debug
	//fmt.Println("connection search:", focusbead.Beadtype, "increase:", increase)

	for d := int8(0); d < 6; d++ {
		if d != focusbead.PrevDir && d != focusbead.NextDir {
			p2 := focusbead.pos.NextPosition(int8(d))
			b2, f2 := c.Formed.GetBead(p2.GetPos())

			/*
				//debug
				if f2 {
					f1x := strconv.Itoa(int(focusbead.GetPosition().X))
					f1y := strconv.Itoa(int(focusbead.GetPosition().Y))
					f2x := strconv.Itoa(int(b2.GetPosition().X))
					f2y := strconv.Itoa(int(b2.GetPosition().Y))
					fmt.Println("debug", "f1:", focusbead.Beadtype, "at", "("+f1x, f1y+")", "f2", b2.Beadtype, "at", "("+f2x, f2y+")", "::hit l.210")
				} else {
					fmt.Println("debug", "bead2:", f2, "::l.212")
				}
			*/

			if f2 && ruleset.Contain(focusbead.Beadtype, b2.Beadtype) {
				//fmt.Println("debug", "increase:", increase, "::hit l.215")
				b2.Nowconnect = uint(int(b2.Nowconnect) + pm)
				c.Formed.PutBead(b2)
				dirofBond += (1 << uint8(d))
				nofAdapt++
			}
		}
	}
	return nofAdapt, dirofBond
}

/*
func MakeRoute(last Position, param Parameter, form Conformation, routeTrans list.List) (Route, bool) {
	temp := form.Seed.MapMarge(form.Formed)
	path := list.New()

	var localDelay int
	if (int)(param.Delay) > routeTrans.Len() {
		localDelay = routeTrans.Len()
	} else {
		localDelay = (int)(param.Delay)
	}

	tBead, ok1 := form.GetBead(last.GetPos())
	if !ok1 {
		return Route{}, false
	}
	path.PushFront(tBead)
	temp.PutBead(tBead)

	MakeRouteOneStep := func(nowpos Position, route *list.List, temp *BeadMap) {
		for i := (int8)(0); i < 6; i++ {
			npos := nowpos.NextPosition(i)
			_, exist := temp.GetBead(npos.GetPos())

		}
	}
}
*/

///////////////
///////////////

func (c *Conformation) RemoveLastFormd(ruleset Ruleset) uint {
	lastbead, _ := c.RemoveBead(c.LastFormed.GetPos())
	num, _ := c.connectionAdapt(lastbead, ruleset, false)

	c.ChangeLastFormed(lastbead.pos.NextPosition(lastbead.PrevDir))

	pred, _ := c.GetBead(c.LastFormed.GetPos())
	pred.NextDir = noDir

	c.Formed.PutBead(pred)

	return num
}

func (r *Route) AppendToTail(tail Route) {
	r.Path = append(r.Path, tail.Path...)
	r.Bonds += tail.Bonds
	r.Eachbonds = append(r.Eachbonds, tail.Eachbonds...)
	r.Bonddir = append(r.Bonddir, tail.Bonddir...)
}

func (r Route) PathToBase6() int {
	ans := 0
	pow := 1
	for i := len(r.Path); i > 0; i-- {
		ans += int(r.Path[i-1]) * pow
		pow = pow * 6
	}

	return ans
}

func (r Route) Later0() int {
	ans := -1
	for index, e := range r.Eachbonds {
		if int(e) != 0 {
			ans = index
		}
	}
	/*
		if ans == len(r.Path)-1 {
			ans = -1
		}
	*/
	return ans
}

func (r2 Route) MatchPath(r1 Route, until int) bool {
	if until < len(r1.Path) && until < len(r2.Path) {
		for i := 0; i <= until; i++ {
			if r1.Path[i] != r2.Path[i] {
				return false
			}
		}
		return true
	}
	return false
}

func (c *Conformation) ReadyExecuting() {
	c.Formed = c.Formed.MapMarge(Position{0, 0}, c.Seed)
}

func (c *Conformation) MoveToSeed() {
	c.Seed = c.Seed.MapMarge(Position{0, 0}, c.Formed)
	c.LastSeed = c.LastFormed
	c.Formed.Initialize()
}

///////////////////
//// debug
///////////////////

func (p Parameter) viewParameter(viewrule bool) {
	fmt.Println("Parameter-----")
	fmt.Println("Delay:", p.Delay, "Arity:", p.Arity)

	fmt.Print("Transcript=")
	fmt.Print(p.Transcript.Arr)
	fmt.Println("")

	if viewrule {
		fmt.Println("Ruleset={")
		for b1 := 1; b1 <= 132; b1++ {
			r2s, ok := p.Rule.rule[strconv.Itoa(b1)]
			if ok {
				sort.Strings(r2s)
				for _, r2 := range r2s {
					b2, ok2 := strconv.Atoi(r2)
					if ok2 == nil && b1 <= b2 {
						fmt.Println("(" + strconv.Itoa(b1) + "," + strconv.Itoa(b2) + "),")
					}
				}
			}
		}
		fmt.Println("}----end")
		/*
			fmt.Println("Ruleset={")
			for b1, v := range p.Rule.rule {
				for _, b2 := range v {
					fmt.Println(" -(", b1, b2, ")")
				}
			}
			fmt.Println("}----end")
		*/
	}
}

func (err error_route) viewError() {
	switch err.errCode {
	case err_out_of_index:
		fmt.Println("OS error:", "out of transcript range")
	case err_no_space:
		fmt.Println("OS error:", "no space")
	case err_non_deterministic:
		fmt.Println("OS error:", "non-deterministic")
		fmt.Println(" -number of bonds:", err.nondet[0].Bonds)
		fmt.Println(" -Transcript = ", err.transcript.Arr[0:err.delay])
		fmt.Println(" -nondetroute")
		for _, r := range err.nondet {
			fmt.Println(" --", r.Path)
		}
	}
}

func (r Route) viewRoute() {
	fmt.Println("route----")
	fmt.Println("position=(", r.Detpos.X, r.Detpos.Y, ")")
	fmt.Println("number of bonds =", r.Bonds)
	fmt.Println("route:", r.Path)
}
