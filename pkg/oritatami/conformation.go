package oritatami

const (
	ne    = 0
	east  = 1
	se    = 2
	nw    = 3
	west  = 4
	sw    = 5
	noDir = -1
)

type Bead struct {
	Beadtype    string
	pos         Position
	NextDir     int8
	PrevDir     int8
	Nowconnect  uint
	Stabconnect uint
	StabconDir  uint8
}

type Position struct {
	X int64
	Y int64
}

type Conformation struct {
	Seed       BeadMap
	LastSeed   Position
	Formed     BeadMap
	LastFormed Position
	Routes     []Routes
}

type BeadMap struct {
	beadmap map[int64]map[int64]Bead
}

type Transcript struct {
	Arr []string
}

type Ruleset struct {
	rule map[string][]string
}

/*
func ForAllDirection(f, args ...interface{}) interface{} {
	for i := (int8)(0); i < 6; i++ {

	}
}
*/

/////////////////
/// Bead
/////////////////

func (b Bead) HasSuccessor() bool {
	if b.NextDir == noDir {
		return false
	} else {
		return true
	}
}

func (b Bead) GetPosition() Position {
	return b.pos
}

func NewBead(beadtype string, position Position) *Bead {
	return &Bead{beadtype, position, noDir, noDir, 0, 0, 0}
}

func NewBeadca(beadtype string, x int64, y int64, next int8, prev int8, nowconnect uint, stabconnect uint, stabdir uint8) Bead {
	return Bead{beadtype, Position{x, y}, next, prev, nowconnect, stabconnect, stabdir}
}

/////////////////
/// Position
/////////////////

func (p Position) GetPos() (int64, int64) {
	return p.X, p.Y
}

func (p Position) NextPositionConDist(dir int8, dist int64) Position {
	switch dir {
	case ne:
		return Position{p.X + dist, p.Y - dist}
	case east:
		return Position{p.X + dist, p.Y}
	case se:
		return Position{p.X, p.Y + dist}
	case nw:
		return Position{p.X, p.Y - dist}
	case west:
		return Position{p.X - dist, p.Y}
	case sw:
		return Position{p.X - dist, p.Y + dist}
	}
	return Position{}
}

func (p Position) NextPosition(dir int8) Position {
	return p.NextPositionConDist(dir, 1)
}

func (p Position) TransPosition(trans Position) Position {
	return Position{p.X + trans.X, p.Y + trans.Y}
}

func (p Position) Opposite() Position {
	return Position{-p.X, -p.Y}
}

/////////////////
/// BeadMap
/////////////////
func (bm *BeadMap) Initialize() {
	bm.beadmap = make(map[int64]map[int64]Bead)
}

func (bm BeadMap) GetBead(x int64, y int64) (Bead, bool) {
	xval, xok := bm.beadmap[x]
	if xok {
		yval, yok := xval[y]
		if yok {
			return yval, true
		}
	}
	return Bead{}, false
}

func (bm *BeadMap) PutBead(bead Bead) {
	if yv, ok := bm.beadmap[bead.pos.X]; ok {
		yv[bead.pos.Y] = bead
	} else {
		bm.beadmap[bead.pos.X] = map[int64]Bead{bead.pos.Y: bead}
	}
}

func (bm BeadMap) MapMarge(trans Position, maps ...BeadMap) BeadMap {
	marged := BeadMap{}
	marged.Initialize()
	for _, vs := range bm.beadmap {
		for _, b := range vs {
			b.pos = b.pos.TransPosition(trans)
			marged.PutBead(b)
		}
	}

	for _, vmap := range maps {
		for _, vs := range vmap.beadmap {
			for _, b := range vs {
				marged.PutBead(b)
			}
		}
	}
	return marged
}

func (bm BeadMap) Getmap() map[int64]map[int64]Bead {
	return bm.beadmap
}

func (c *Conformation) FromDirectionToSeedWithConformation(dircPrefixIsBridge string, transcript Transcript, start Position) {
	_, ok := c.GetBead(start.GetPos())
	if !ok {
		fpos := start.NextPosition(DirExpChar2Byte(rune(dircPrefixIsBridge[0])))
		c.Seed.PutBead(Bead{transcript.Arr[0], fpos, noDir, noDir, 0, 0, 0})
		nw := Transcript{transcript.Arr[1:]}
		dircPrefixIsBridge = string(dircPrefixIsBridge[1:])

		pos, _ := c.Seed.AddByDirection(dircPrefixIsBridge, nw, fpos)
		c.LastSeed = pos
		c.LastFormed = pos
	} else {
		pos, _ := c.Seed.AddByDirection(dircPrefixIsBridge, transcript, start)
		c.LastSeed = pos
		c.LastFormed = pos
	}
}

func MakeNewConformationFromDirectionToSeed(dirc string, transcript Transcript, start Position) Conformation {
	seed := BeadMap{make(map[int64]map[int64]Bead)}
	form := BeadMap{make(map[int64]map[int64]Bead)}

	seed.PutBead(Bead{transcript.Arr[0], start, noDir, noDir, 0, 0, 0})
	w := Transcript{transcript.Arr[1:]}
	pos, _ := seed.AddByDirection(dirc, w, start)
	return Conformation{seed, pos, form, pos, []Routes{}}
}

func (bm *BeadMap) AddByDirection(dirc string, transcript Transcript, start Position) (Position, Transcript) {
	if len(dirc) > len(transcript.Arr) {
		dirc = string(dirc[:len(transcript.Arr)])
	}

	putf := func(pos Position, ndir rune, base string) Position {
		bdir := DirExpChar2Byte(ndir)
		opp := (int8)(5 - bdir)
		fb, okfb := bm.GetBead(pos.GetPos())
		spos := pos.NextPosition(bdir)

		sbead := Bead{base, spos, noDir, opp, 0, 0, 0}
		if okfb {
			fb.NextDir = bdir
			bm.PutBead(fb)
		}
		bm.PutBead(sbead)

		//fmt.Print(base, ",", string(ndir), ",", bdir, " ")
		return spos
	}

	ans := transcript
	npos := start

	for _, l := range dirc {
		n := ans.Arr[0]
		npos = putf(npos, l, n)
		ans.Arr = ans.Arr[1:]
	}

	//fmt.Println()
	return npos, ans
}

/////////////////////
///  Conformation
/////////////////////

func (c Conformation) GetBead(x int64, y int64) (Bead, bool) {
	fval, fok := c.Formed.GetBead(x, y)
	if fok {
		return fval, true
	} else {
		sval, sok := c.Seed.GetBead(x, y)
		if sok {
			return sval, true
		}
	}
	return Bead{}, false
}

func (c *Conformation) RemoveBead(x int64, y int64) (Bead, bool) {
	fval, fok := c.Formed.GetBead(x, y)
	if fok {
		v, _ := c.Formed.beadmap[x]
		delete(v, y)
		return fval, true
	} else {
		return Bead{}, false
	}
}

func (c *Conformation) Extend(direction int8, beadtype string, connection uint, connectionDir uint8) bool {
	pos := c.LastFormed.NextPosition(direction)
	_, check := c.Formed.GetBead(pos.GetPos())
	if check {
		return false
	}
	//Bead
	//	Beadtype    string
	//	pos         Position
	//	NextDir     int8
	//	PrevDir     int8
	//	Nowconnect  uint
	//	Stabconnect uint
	//	StabconDir  uint8

	pbead, _ := c.Formed.GetBead(c.LastFormed.GetPos())
	pbead.NextDir = direction
	c.Formed.PutBead(pbead)
	c.Formed.PutBead(Bead{beadtype, pos, noDir, int8(5 - direction), connection, connection, connectionDir})

	c.LastFormed = pos

	return true
}

func (c *Conformation) appendNewRoute(route Route) {
	curmap := c.Routes[len(c.Routes)-1]
	//curmap.appendRoute(route)
	curmap.Routes = append(curmap.Routes, route)

	if route.Bonds > curmap.Mbonds {
		curmap.Mbonds = route.Bonds
	}

	c.Routes[len(c.Routes)-1] = curmap

	//c.Routes[len(c.Routes)-1] = curmap
}

func (c *Conformation) nextRouteMap(localtranscript Transcript) {
	newmap := NewRoutes()
	newmap.Transcript = localtranscript.Arr
	c.Routes = append(c.Routes, *newmap)
}

func (c *Conformation) ChangeLastFormed(newpos Position) {
	/*
		//debug
		fmt.Print("debug: LastFormed (")
		fmt.Print(c.LastFormed.GetPos())
		fmt.Print(") -> (")
		fmt.Print(newpos.GetPos())
		fmt.Println(")")
	*/
	c.LastFormed = newpos
}

/////////////////
/// Ruleset
/////////////////

func (r *Ruleset) Add(b1 string, b2 string) {
	//rule map[int64][]int64
	if r.Contain(b1, b2) {
		return
	}
	arr1, ok1 := r.rule[b1]
	if ok1 {
		r.rule[b1] = append(arr1, b2)
	} else {
		r.rule[b1] = []string{b2}
	}

	arr2, ok2 := r.rule[b2]
	if ok2 {
		r.rule[b2] = append(arr2, b1)
	} else {
		r.rule[b2] = []string{b1}
	}
}

func (r Ruleset) Contain(b1 string, b2 string) bool {
	arr, ok := r.rule[b1]
	if ok {
		for _, e := range arr {
			if e == b2 {
				return true
			}
		}
	}
	return false
}

func (r *Ruleset) Initialize() {
	r.rule = make(map[string][]string)
}

/////////////////
/// Transcript
/////////////////

func (t *Transcript) Initialize() {
	t.Arr = []string{}
}

/*
func (t *Transcript) AddArr(transcript []string) {
	for _, e := range transcript {
		l, err := strconv.Atoi(strings.TrimSpace(e))
		if err == nil {
			t.Arr = append(t.Arr, int64(l))
		}
	}
}
*/

func (t Transcript) GetLocalTranscript(index int, length int) (Transcript, bool) {
	if len(t.Arr) > index+length {
		return Transcript{t.Arr[index:length]}, true
	} else if len(t.Arr) > index {
		return Transcript{t.Arr[index:]}, true
	} else {
		return Transcript{}, false
	}
}

/////////////////
/// Transcript
/////////////////

func NewRoutes() *Routes {
	return &Routes{[]Route{}, []string{}, 0}
}

func (r *Routes) appendRoute(route Route) {
	//fmt.Println("-", route.Path)
	r.Routes = append(r.Routes, route)

	/* //for map
	ros, ok := r.Routes[route.Bonds]
	if ok {
		r.Routes[route.Bonds] = append(ros, route)
	} else {
		if route.Bonds > r.Mbonds {
			r.Mbonds = route.Bonds
		}
		r.Routes[route.Bonds] = []Route{route}
	}
	*/
}

func (r Route) clone() *Route {
	newBonddir := []uint8{}
	newEachbonds := []uint{}
	newPath := []int8{}
	for index, bd := range r.Bonddir {
		newBonddir = append(newBonddir, bd)
		newEachbonds = append(newEachbonds, r.Eachbonds[index])
		newPath = append(newPath, r.Path[index])
	}
	return &Route{newPath, r.Bonds, newEachbonds, newBonddir, r.Detpos}
}
