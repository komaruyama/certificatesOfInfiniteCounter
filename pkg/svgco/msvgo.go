package svgco

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	TAB = '	'
)

type TextSvgStruct struct {
	fontname  string
	fontcolor string
	fontsize  string
}

type LinePathShape struct {
	VertexArr    []string
	Stroke_width string
	Stroke_color string
	Fill         bool
	Fillcolor    string
	Stroke_round bool
	Edge_round   string
}

type RectSvg struct {
	X            float64
	Y            float64
	Width        float64
	Height       float64
	CornerRadius float64
	Transform    TransformSvg
	Color        ShapeColor
	Clip         string
}

type CircleSvg struct {
	Cx        float64
	Cy        float64
	Radius    float64
	Transform TransformSvg
	Color     ShapeColor
	Clip      string
}

type TextSvg struct {
	X         float64
	Y         float64
	Text      string
	Font      FontSvg
	Anchor    string // start or middle or end
	DetY      float64
	Transform TransformSvg
}

type LineSvg struct {
	X1        float64
	Y1        float64
	X2        float64
	Y2        float64
	Color     ShapeColor
	Transform TransformSvg
	Clip      string
}

type FontSvg struct {
	Font_family string
	Font_size   float64
	Fill_color  string
}

type ShapeColor struct {
	Stroke       bool
	Stroke_color string
	Stroke_width float64
	Fill         bool
	Fill_color   string
}

type TransformSvg struct {
	TransType string
	Args      []float64
}

type GraphEdge struct {
	V1    string
	V2    string
	Color ShapeColor
}

type GraphVertex struct {
	ID     string
	Vertex CircleSvg
	Label  TextSvg
}

type TagStruct struct {
	label string
	plain bool
	param map[string]string
	body  []TagStruct
}

////////////

func DrawLinePathAndFill(shape LinePathShape) []string {
	if len(shape.VertexArr) < 2 {
		return []string{}
	}
	fst := "<path d=\"M " + strings.TrimSpace(shape.VertexArr[0])
	rV := []string{fst}
	path := "L "
	for _, e := range shape.VertexArr[1:] {
		path += strings.TrimSpace(e)
		path += " "
	}
	if shape.Stroke_round {
		path += "Z"
	}
	rV = append(rV, path)

	opptions := "\" stroke=\""
	opptions += shape.Stroke_color
	opptions += "\" "
	if len(shape.Stroke_width) != 0 {
		opptions += "stroke-width=\"" + shape.Stroke_width + "\" "
	}
	if shape.Fill {
		opptions += "fill=\"" + shape.Fillcolor + "\" "
	} else {
		opptions += "fill=\"none\" "
	}
	if len(shape.Edge_round) != 0 {
		opptions += "stroke-linejoin=\"" + shape.Edge_round + "\" "
	}
	opptions += "/>"
	rV = append(rV, opptions)
	return rV
}

func WriteText(x string, y string, font_size string, text string, text_anchor string,
	stroke_width string, stroke_color string, others map[string]string) string {
	pref := "<text x=\"" + x + "\" y=\"" + y + "\" font-size=\"" + font_size +
		"\" stroke=\"" + stroke_color + "\" stroke-width=\"" + stroke_width +
		"\" text-anchor=\"" + text_anchor + "\" "
	suff := ">" + text + "</text>"

	for k, v := range others {
		pref += k + "=\"" + v + "\" "
	}
	return pref + suff
}

func PackLinesToSvg(lines []string) []string {
	rV := []string{"<svg xmlns=\"http://www.w3.org/2000/svg\">"}
	rV = append(rV, lines...)
	rV = append(rV, "</svg>")

	return rV
}

func GraphToSvg(edges []GraphEdge, vertices map[string]GraphVertex, transform TransformSvg, clipId string) []string {
	posPrec := 4
	strokePrec := 0
	fontsizePrec := 1

	paramsEdge := []TagStruct{}
	paramsVertex := []TagStruct{}
	paramsLabel := []TagStruct{}

	for _, e := range edges {
		v1 := vertices[e.V1]
		v2 := vertices[e.V2]

		line := LineSvg{v1.Vertex.Cx, v1.Vertex.Cy, v2.Vertex.Cx, v2.Vertex.Cy, e.Color, transform, ""}

		paramsEdge = append(paramsEdge, line.ToStruct(posPrec, strokePrec))
	}

	for _, v := range vertices {
		paramsVertex = append(paramsVertex, v.Vertex.ToStruct(posPrec, strokePrec))
		paramsLabel = append(paramsLabel, v.Label.ToStruct(posPrec, fontsizePrec))
	}

	// if clipId len > 0 ....
	if len(clipId) > 0 {
		for _, e := range paramsEdge {
			e.param["clip-path"] = "url(#" + clipId + ")"
		}
		for _, e := range paramsVertex {
			e.param["clip-path"] = "url(#" + clipId + ")"
		}
		for _, e := range paramsLabel {
			e.param["clip-path"] = "url(#" + clipId + ")"
		}
	}

	ans := []string{}

	for _, e := range paramsEdge {
		ans = append(ans, tagstructToString(e, TAB, 0)...)
	}
	for _, e := range paramsVertex {
		ans = append(ans, tagstructToString(e, TAB, 0)...)
	}
	for _, e := range paramsLabel {
		ans = append(ans, tagstructToString(e, TAB, 0)...)
	}

	return ans
}

/////////////////////
// shape to string
/////////////////////

func GraphToStructs(e []GraphEdge, v map[string]GraphVertex, posPrec int, strokePrec int, fontsizePrec int, transform TransformSvg, clip string) []TagStruct {
	params := []TagStruct{}

	for _, edge := range e {
		v1, ok1 := v[edge.V1]
		v2, ok2 := v[edge.V2]

		if ok1 && ok2 {
			line := LineSvg{v1.Vertex.Cx, v1.Vertex.Cy, v2.Vertex.Cx, v2.Vertex.Cy, edge.Color, transform, ""}

			params = append(params, line.ToStruct(posPrec, strokePrec))
		} else {
			fmt.Println("at svgco.GraphToStruct,", "an edge:", edge.V1, "--", edge.V2, "does Not exist")
		}
	}

	for _, v := range v {
		params = append(params, v.Vertex.ToStruct(posPrec, strokePrec))
		params = append(params, v.Label.ToStruct(posPrec, fontsizePrec))
	}

	/*
		// if clipId len > 0 ....
		if len(clip) > 0 {
			for _, e := range paramsEdge {
				e.param["clip-path"] = "url(#" + clipId + ")"
			}
			for _, e := range paramsVertex {
				e.param["clip-path"] = "url(#" + clipId + ")"
			}
			for _, e := range paramsLabel {
				e.param["clip-path"] = "url(#" + clipId + ")"
			}
		}
	*/
	return params
}

func (r RectSvg) ToStruct(positionPrec int, thickPrec int) TagStruct {
	x := strconv.FormatFloat(r.X, 'f', positionPrec, 32)
	y := strconv.FormatFloat(r.Y, 'f', positionPrec, 32)
	width := strconv.FormatFloat(r.Width, 'f', positionPrec, 32)
	height := strconv.FormatFloat(r.Height, 'f', positionPrec, 32)

	param := map[string]string{"x": x, "y": y, "width": width, "height": height}

	if r.CornerRadius <= 0 {
		rx := strconv.FormatFloat(r.CornerRadius, 'f', positionPrec, 32)
		param["rx"] = rx
	}

	r.Transform.AppendToParam(positionPrec, param)
	r.Color.appendToParam(thickPrec, param)

	AppendClip(param, r.Clip)

	return TagStruct{"rect", false, param, []TagStruct{}}
}

func (l LineSvg) ToStruct(positionPrec int, thickPrec int) TagStruct {
	x1 := strconv.FormatFloat(l.X1, 'f', positionPrec, 32)
	y1 := strconv.FormatFloat(l.Y1, 'f', positionPrec, 32)
	x2 := strconv.FormatFloat(l.X2, 'f', positionPrec, 32)
	y2 := strconv.FormatFloat(l.Y2, 'f', positionPrec, 32)

	param := map[string]string{"x1": x1, "x2": x2, "y1": y1, "y2": y2}
	l.Color.appendToParam(thickPrec, param)
	l.Transform.AppendToParam(positionPrec, param)

	AppendClip(param, l.Clip)

	return TagStruct{"line", false, param, []TagStruct{}}
}

func (c CircleSvg) ToStruct(positionPrec int, thickPrec int) TagStruct {
	lx := strconv.FormatFloat(c.Cx, 'f', positionPrec, 32)
	ly := strconv.FormatFloat(c.Cy, 'f', positionPrec, 32)
	lr := strconv.FormatFloat(c.Radius, 'f', positionPrec, 32)

	param := map[string]string{"cx": lx, "cy": ly, "r": lr}

	c.Color.appendToParam(thickPrec, param)
	c.Transform.AppendToParam(positionPrec, param)

	AppendClip(param, c.Clip)

	return TagStruct{"circle", false, param, []TagStruct{}}
}

func (t TextSvg) ToStruct(positionPrec int, fontsizePrec int) TagStruct {
	if len(t.Text) == 0 {
		return TagStruct{"", true, nil, nil}
	}

	lx := strconv.FormatFloat(t.X, 'f', positionPrec, 32)
	ly := strconv.FormatFloat(t.Y+t.DetY, 'f', positionPrec, 32)

	param := map[string]string{"x": lx, "y": ly, "text-anchor": t.Anchor}
	t.Font.appendToParm(fontsizePrec, param)
	t.Transform.AppendToParam(positionPrec, param)

	return TagStruct{"text", false, param, []TagStruct{TagStruct{t.Text, true, nil, nil}}}
}

func MakeClipTagStruct(id string, shapes []TagStruct) TagStruct {
	return TagStruct{"clipPath", false, map[string]string{"id": id}, shapes}
}

func MakeSvgString(elements []TagStruct, viewbox ...int) []string {
	// viewbox[0-3] = x1, y1, x2, y2

	//TagStruct
	// label string
	// plain bool
	// param map[string]string
	// body  []TagStruct
	param := map[string]string{"xmlns": "http://www.w3.org/2000/svg"}
	if len(viewbox) == 4 {
		param["viewBox"] = strconv.Itoa(viewbox[0]) + " " + strconv.Itoa(viewbox[1]) + " " + strconv.Itoa(viewbox[2]) + " " + strconv.Itoa(viewbox[3])
	}
	return tagstructToString(TagStruct{"svg", false, param, elements}, TAB, 0)
}

//////////////////////
// element to map
//////////////////////

func (t TransformSvg) AppendToParam(prec int, param map[string]string) {
	//ans := "transform=\""
	if len(t.TransType) == 0 {
		return
	}
	ans := t.TransType
	ans += "("
	for _, e := range t.Args {
		l := strconv.FormatFloat(e, 'f', prec, 32)
		ans += l + " "
	}
	ans += ")"

	par, ok := param["transform"]
	if ok {
		param["transform"] = par + " " + ans
	} else {
		param["transform"] = ans
	}
}

func (c ShapeColor) appendToParam(strokePrec int, param map[string]string) {
	// if the parameter already exists, it is over writen
	if c.Fill {
		param["fill"] = c.Fill_color
	} else {
		param["fill"] = "none"
	}

	if c.Stroke {
		param["stroke"] = c.Stroke_color
		param["stroke-width"] = strconv.FormatFloat(c.Stroke_width, 'f', strokePrec, 32)
	}
}

func (f FontSvg) appendToParm(fontsizePrec int, param map[string]string) {
	param["font-family"] = f.Font_family
	param["font-size"] = strconv.FormatFloat(f.Font_size, 'f', fontsizePrec, 32)
	if len(f.Fill_color) > 0 {
		param["fill"] = f.Fill_color
	}
}

//////////////////
func NewTagStructure(label string, plain bool, param map[string]string, body []TagStruct) TagStruct {
	//label string
	//plain bool
	//param map[string]string
	//body  []TagStruct
	return TagStruct{label, plain, param, body}
}

func AppendClip(param map[string]string, clip string) {
	if len(clip) > 0 {
		param["clip-path"] = "url(#" + clip + ")"
	}
}

func tagstructToString(tag TagStruct, tabChar rune, depth int) []string {
	if len(tag.label) == 0 {
		return nil
	}

	insTab := func() string {
		ans := ""
		for i := 0; i < depth; i++ {
			ans += string(tabChar)
		}
		return ans
	}

	if tag.plain {
		return []string{insTab() + tag.label}
	} else {
		line := insTab()
		line += "<" + tag.label + " "
		for k, v := range tag.param {
			line += k
			line += "=\""
			line += v
			line += "\" "
		}

		if len(tag.body) == 0 {
			line += "/>"
			return []string{line}
		} else {
			line += ">"
			ans := []string{line}
			for _, cot := range tag.body {
				ans = append(ans, tagstructToString(cot, tabChar, depth+1)...)
			}
			ans = append(ans, insTab()+"</"+tag.label+">")
			return ans
		}
	}
}
