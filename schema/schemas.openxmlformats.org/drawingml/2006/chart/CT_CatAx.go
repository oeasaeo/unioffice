// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_CatAx struct {
	AxId           *CT_UnsignedInt
	Scaling        *CT_Scaling
	Delete         *CT_Boolean
	AxPos          *CT_AxPos
	MajorGridlines *CT_ChartLines
	MinorGridlines *CT_ChartLines
	Title          *CT_Title
	NumFmt         *CT_NumFmt
	MajorTickMark  *CT_TickMark
	MinorTickMark  *CT_TickMark
	TickLblPos     *CT_TickLblPos
	SpPr           *drawingml.CT_ShapeProperties
	TxPr           *drawingml.CT_TextBody
	CrossAx        *CT_UnsignedInt
	Choice         *EG_AxSharedChoice
	Auto           *CT_Boolean
	LblAlgn        *CT_LblAlgn
	LblOffset      *CT_LblOffset
	TickLblSkip    *CT_Skip
	TickMarkSkip   *CT_Skip
	NoMultiLvlLbl  *CT_Boolean
	ExtLst         *CT_ExtensionList
}

func NewCT_CatAx() *CT_CatAx {
	ret := &CT_CatAx{}
	ret.AxId = NewCT_UnsignedInt()
	ret.Scaling = NewCT_Scaling()
	ret.AxPos = NewCT_AxPos()
	ret.CrossAx = NewCT_UnsignedInt()
	return ret
}

func (m *CT_CatAx) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seaxId := xml.StartElement{Name: xml.Name{Local: "c:axId"}}
	e.EncodeElement(m.AxId, seaxId)
	sescaling := xml.StartElement{Name: xml.Name{Local: "c:scaling"}}
	e.EncodeElement(m.Scaling, sescaling)
	if m.Delete != nil {
		sedelete := xml.StartElement{Name: xml.Name{Local: "c:delete"}}
		e.EncodeElement(m.Delete, sedelete)
	}
	seaxPos := xml.StartElement{Name: xml.Name{Local: "c:axPos"}}
	e.EncodeElement(m.AxPos, seaxPos)
	if m.MajorGridlines != nil {
		semajorGridlines := xml.StartElement{Name: xml.Name{Local: "c:majorGridlines"}}
		e.EncodeElement(m.MajorGridlines, semajorGridlines)
	}
	if m.MinorGridlines != nil {
		seminorGridlines := xml.StartElement{Name: xml.Name{Local: "c:minorGridlines"}}
		e.EncodeElement(m.MinorGridlines, seminorGridlines)
	}
	if m.Title != nil {
		setitle := xml.StartElement{Name: xml.Name{Local: "c:title"}}
		e.EncodeElement(m.Title, setitle)
	}
	if m.NumFmt != nil {
		senumFmt := xml.StartElement{Name: xml.Name{Local: "c:numFmt"}}
		e.EncodeElement(m.NumFmt, senumFmt)
	}
	if m.MajorTickMark != nil {
		semajorTickMark := xml.StartElement{Name: xml.Name{Local: "c:majorTickMark"}}
		e.EncodeElement(m.MajorTickMark, semajorTickMark)
	}
	if m.MinorTickMark != nil {
		seminorTickMark := xml.StartElement{Name: xml.Name{Local: "c:minorTickMark"}}
		e.EncodeElement(m.MinorTickMark, seminorTickMark)
	}
	if m.TickLblPos != nil {
		setickLblPos := xml.StartElement{Name: xml.Name{Local: "c:tickLblPos"}}
		e.EncodeElement(m.TickLblPos, setickLblPos)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "c:spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.TxPr != nil {
		setxPr := xml.StartElement{Name: xml.Name{Local: "c:txPr"}}
		e.EncodeElement(m.TxPr, setxPr)
	}
	secrossAx := xml.StartElement{Name: xml.Name{Local: "c:crossAx"}}
	e.EncodeElement(m.CrossAx, secrossAx)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, start)
	}
	if m.Auto != nil {
		seauto := xml.StartElement{Name: xml.Name{Local: "c:auto"}}
		e.EncodeElement(m.Auto, seauto)
	}
	if m.LblAlgn != nil {
		selblAlgn := xml.StartElement{Name: xml.Name{Local: "c:lblAlgn"}}
		e.EncodeElement(m.LblAlgn, selblAlgn)
	}
	if m.LblOffset != nil {
		selblOffset := xml.StartElement{Name: xml.Name{Local: "c:lblOffset"}}
		e.EncodeElement(m.LblOffset, selblOffset)
	}
	if m.TickLblSkip != nil {
		setickLblSkip := xml.StartElement{Name: xml.Name{Local: "c:tickLblSkip"}}
		e.EncodeElement(m.TickLblSkip, setickLblSkip)
	}
	if m.TickMarkSkip != nil {
		setickMarkSkip := xml.StartElement{Name: xml.Name{Local: "c:tickMarkSkip"}}
		e.EncodeElement(m.TickMarkSkip, setickMarkSkip)
	}
	if m.NoMultiLvlLbl != nil {
		senoMultiLvlLbl := xml.StartElement{Name: xml.Name{Local: "c:noMultiLvlLbl"}}
		e.EncodeElement(m.NoMultiLvlLbl, senoMultiLvlLbl)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CatAx) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.AxId = NewCT_UnsignedInt()
	m.Scaling = NewCT_Scaling()
	m.AxPos = NewCT_AxPos()
	m.CrossAx = NewCT_UnsignedInt()
lCT_CatAx:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "axId":
				if err := d.DecodeElement(m.AxId, &el); err != nil {
					return err
				}
			case "scaling":
				if err := d.DecodeElement(m.Scaling, &el); err != nil {
					return err
				}
			case "delete":
				m.Delete = NewCT_Boolean()
				if err := d.DecodeElement(m.Delete, &el); err != nil {
					return err
				}
			case "axPos":
				if err := d.DecodeElement(m.AxPos, &el); err != nil {
					return err
				}
			case "majorGridlines":
				m.MajorGridlines = NewCT_ChartLines()
				if err := d.DecodeElement(m.MajorGridlines, &el); err != nil {
					return err
				}
			case "minorGridlines":
				m.MinorGridlines = NewCT_ChartLines()
				if err := d.DecodeElement(m.MinorGridlines, &el); err != nil {
					return err
				}
			case "title":
				m.Title = NewCT_Title()
				if err := d.DecodeElement(m.Title, &el); err != nil {
					return err
				}
			case "numFmt":
				m.NumFmt = NewCT_NumFmt()
				if err := d.DecodeElement(m.NumFmt, &el); err != nil {
					return err
				}
			case "majorTickMark":
				m.MajorTickMark = NewCT_TickMark()
				if err := d.DecodeElement(m.MajorTickMark, &el); err != nil {
					return err
				}
			case "minorTickMark":
				m.MinorTickMark = NewCT_TickMark()
				if err := d.DecodeElement(m.MinorTickMark, &el); err != nil {
					return err
				}
			case "tickLblPos":
				m.TickLblPos = NewCT_TickLblPos()
				if err := d.DecodeElement(m.TickLblPos, &el); err != nil {
					return err
				}
			case "spPr":
				m.SpPr = drawingml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case "txPr":
				m.TxPr = drawingml.NewCT_TextBody()
				if err := d.DecodeElement(m.TxPr, &el); err != nil {
					return err
				}
			case "crossAx":
				if err := d.DecodeElement(m.CrossAx, &el); err != nil {
					return err
				}
			case "crosses":
				m.Choice = NewEG_AxSharedChoice()
				if err := d.DecodeElement(&m.Choice.Crosses, &el); err != nil {
					return err
				}
			case "crossesAt":
				m.Choice = NewEG_AxSharedChoice()
				if err := d.DecodeElement(&m.Choice.CrossesAt, &el); err != nil {
					return err
				}
			case "auto":
				m.Auto = NewCT_Boolean()
				if err := d.DecodeElement(m.Auto, &el); err != nil {
					return err
				}
			case "lblAlgn":
				m.LblAlgn = NewCT_LblAlgn()
				if err := d.DecodeElement(m.LblAlgn, &el); err != nil {
					return err
				}
			case "lblOffset":
				m.LblOffset = NewCT_LblOffset()
				if err := d.DecodeElement(m.LblOffset, &el); err != nil {
					return err
				}
			case "tickLblSkip":
				m.TickLblSkip = NewCT_Skip()
				if err := d.DecodeElement(m.TickLblSkip, &el); err != nil {
					return err
				}
			case "tickMarkSkip":
				m.TickMarkSkip = NewCT_Skip()
				if err := d.DecodeElement(m.TickMarkSkip, &el); err != nil {
					return err
				}
			case "noMultiLvlLbl":
				m.NoMultiLvlLbl = NewCT_Boolean()
				if err := d.DecodeElement(m.NoMultiLvlLbl, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_CatAx %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CatAx
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CatAx and its children
func (m *CT_CatAx) Validate() error {
	return m.ValidateWithPath("CT_CatAx")
}

// ValidateWithPath validates the CT_CatAx and its children, prefixing error messages with path
func (m *CT_CatAx) ValidateWithPath(path string) error {
	if err := m.AxId.ValidateWithPath(path + "/AxId"); err != nil {
		return err
	}
	if err := m.Scaling.ValidateWithPath(path + "/Scaling"); err != nil {
		return err
	}
	if m.Delete != nil {
		if err := m.Delete.ValidateWithPath(path + "/Delete"); err != nil {
			return err
		}
	}
	if err := m.AxPos.ValidateWithPath(path + "/AxPos"); err != nil {
		return err
	}
	if m.MajorGridlines != nil {
		if err := m.MajorGridlines.ValidateWithPath(path + "/MajorGridlines"); err != nil {
			return err
		}
	}
	if m.MinorGridlines != nil {
		if err := m.MinorGridlines.ValidateWithPath(path + "/MinorGridlines"); err != nil {
			return err
		}
	}
	if m.Title != nil {
		if err := m.Title.ValidateWithPath(path + "/Title"); err != nil {
			return err
		}
	}
	if m.NumFmt != nil {
		if err := m.NumFmt.ValidateWithPath(path + "/NumFmt"); err != nil {
			return err
		}
	}
	if m.MajorTickMark != nil {
		if err := m.MajorTickMark.ValidateWithPath(path + "/MajorTickMark"); err != nil {
			return err
		}
	}
	if m.MinorTickMark != nil {
		if err := m.MinorTickMark.ValidateWithPath(path + "/MinorTickMark"); err != nil {
			return err
		}
	}
	if m.TickLblPos != nil {
		if err := m.TickLblPos.ValidateWithPath(path + "/TickLblPos"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	if m.TxPr != nil {
		if err := m.TxPr.ValidateWithPath(path + "/TxPr"); err != nil {
			return err
		}
	}
	if err := m.CrossAx.ValidateWithPath(path + "/CrossAx"); err != nil {
		return err
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if m.Auto != nil {
		if err := m.Auto.ValidateWithPath(path + "/Auto"); err != nil {
			return err
		}
	}
	if m.LblAlgn != nil {
		if err := m.LblAlgn.ValidateWithPath(path + "/LblAlgn"); err != nil {
			return err
		}
	}
	if m.LblOffset != nil {
		if err := m.LblOffset.ValidateWithPath(path + "/LblOffset"); err != nil {
			return err
		}
	}
	if m.TickLblSkip != nil {
		if err := m.TickLblSkip.ValidateWithPath(path + "/TickLblSkip"); err != nil {
			return err
		}
	}
	if m.TickMarkSkip != nil {
		if err := m.TickMarkSkip.ValidateWithPath(path + "/TickMarkSkip"); err != nil {
			return err
		}
	}
	if m.NoMultiLvlLbl != nil {
		if err := m.NoMultiLvlLbl.ValidateWithPath(path + "/NoMultiLvlLbl"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
