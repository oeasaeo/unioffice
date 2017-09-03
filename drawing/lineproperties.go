// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawing

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/measurement"
	dml "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type LineProperties struct {
	x *dml.CT_LineProperties
}

// X returns the inner wrapped XML type.
func (l LineProperties) X() *dml.CT_LineProperties {
	return l.x
}

// SetWidth sets the line width, MS products treat zero as the minimum width
// that can be displayed.
func (l LineProperties) SetWidth(w measurement.Distance) {
	// TODO:  check these units, can't find documentation on them but this seems
	// to be the right range
	l.x.WAttr = gooxml.Int32(int32(w / measurement.EMU))
}

func (l LineProperties) SetSolidFill(c color.Color) {
	l.x.NoFill = nil
	l.x.GradFill = nil
	l.x.PattFill = nil
	l.x.SolidFill = dml.NewCT_SolidColorFillProperties()
	l.x.SolidFill.SrgbClr = dml.NewCT_SRgbColor()
	l.x.SolidFill.SrgbClr.ValAttr = *c.AsRGBAString()
}

// LineJoin is the type of line join
type LineJoin byte

// LineJoin types
const (
	LineJoinRound LineJoin = iota
	LineJoinBevel
	LineJoinMiter
)

// SetJoin sets the line join style.
func (l LineProperties) SetJoin(e LineJoin) {
	l.x.Round = nil
	l.x.Miter = nil
	l.x.Bevel = nil
	switch e {
	case LineJoinRound:
		l.x.Round = dml.NewCT_LineJoinRound()
	case LineJoinBevel:
		l.x.Bevel = dml.NewCT_LineJoinBevel()
	case LineJoinMiter:
		l.x.Miter = dml.NewCT_LineJoinMiterProperties()
	}
}
