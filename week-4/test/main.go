package main

import (
	"github.com/go-pdf/fpdf"
)

func main() {
	headerText := "INVOICE"

	marginX := 10.0
	marginY := 20.0
	gapY := 2.0

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(marginX, marginY, marginX)
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 24)

	pageW, _ := pdf.GetPageSize()
	// safeW := pageW - 2*marginX

	_, lineHeight := pdf.GetFontSize()
	currentY := pdf.GetY() + lineHeight + gapY
	lineBreak := lineHeight + 1

	textWidth := pdf.GetStringWidth(headerText)

	pdf.SetXY((pageW-textWidth)/2, currentY)

	pdf.Cell(textWidth, 10, headerText)
	pdf.Ln(lineBreak)
	pdf.Ln(lineBreak)
	pdf.Ln(lineBreak)

	pdf.SetFontSize(14)

	_, lineHeight = pdf.GetFontSize()
	currentY = pdf.GetY() + lineHeight
	pdf.SetXY(marginX, currentY)

	headers := [5]string{"ID", "Name", "Quantity", "Unit Price", "Price"}
	colWidth := [5]float64{10.0, 75.0, 25.0, 40.0, 40.0}

	pdf.SetFillColor(200, 200, 200)
	for col := 0; col < 5; col++ {
		pdf.CellFormat(colWidth[col], 10.0, headers[col], "1", 0, "CM", true, 0, "")
	}

	pdf.Ln(-1)
	pdf.SetFontStyle("")

	lineHeight += gapY

	for row := 0; row < 2; row++ {
		pdf.CellFormat(colWidth[0], lineHeight, "1", "1", 0, "CM", false, 0, "")
		pdf.CellFormat(colWidth[1], lineHeight, "your mom", "1", 0, "CM", false, 0, "")
		pdf.CellFormat(colWidth[2], lineHeight, "2", "1", 0, "CM", false, 0, "")
		pdf.CellFormat(colWidth[3], lineHeight, "25.0", "1", 0, "CM", false, 0, "")
		pdf.CellFormat(colWidth[4], lineHeight, "50.0", "1", 0, "CM", false, 0, "")

		pdf.Ln(-1)
	}

	pdf.SetFontStyle("B")

	leftIndent := 0.0
	for i := 0; i < 3; i++ {
		leftIndent += colWidth[i]
	}

	pdf.SetX(marginX + leftIndent)

	pdf.CellFormat(colWidth[3], lineHeight, "Total", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWidth[4], lineHeight, "50", "1", 0, "CM", false, 0, "")

	pdf.Ln(-1)

	pdf.OutputFileAndClose("test.pdf")
}
