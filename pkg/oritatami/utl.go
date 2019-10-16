package oritatami

const (
	cne    = 'm'
	ceast  = 'e'
	cse    = 'd'
	csw    = 'a'
	cwest  = 'w'
	cnw    = 'b'
	cnodir = 'n'
)

func DirectionwordToPositions(dirction string) []Position {
	pos := []Position{Position{0, 0}}
	focus := Position{0, 0}

	for _, r := range dirction {
		focus := direcToPos(focus, r)
		pos = append(pos, focus)
	}

	return pos
}

func direcToPos(pos Position, dir rune) Position {
	return pos.NextPosition(DirExpChar2Byte(dir))
}

func DirExpChar2Byte(dir rune) int8 {
	switch dir {
	case cne:
		return ne
	case ceast:
		return east
	case cnw:
		return nw
	case cwest:
		return west
	case cse:
		return se
	case csw:
		return sw
	}
	return noDir
}

func DirExpByte2Char(dir int8) rune {
	switch dir {
	case ne:
		return cne
	case east:
		return ceast
	case nw:
		return cnw
	case west:
		return cwest
	case se:
		return cse
	case sw:
		return csw
	}
	return cnodir
}

func FlipDirection(direction string) string {
	ans := []rune{}
	for _, l := range direction {
		d := DirExpChar2Byte(l)
		opp := FlipOneDirection(d)
		ans = append(ans, DirExpByte2Char(opp))
	}
	return string(ans)
}

func FlipOneDirection(direction int8) int8 {
	switch direction {
	case west:
		return east
	case east:
		return west
	case ne:
		return nw
	case se:
		return sw
	case nw:
		return ne
	case sw:
		return se
	}
	return noDir
}
