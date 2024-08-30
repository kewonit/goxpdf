package main

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

func main() {
	// Configuration
	templatePath := "certificate_template.png"
	csvPath := "names.csv"
	outputFolder := "output_certificates"
	fontPath := "Augillion.otf"
	fontSize := float64(100) // Font size in px
	textColor := color.Black // Colour (Black, White, Transparent & Opaque)
	textY := float64(800)    // Y coordinate for text placement

	// Load the certificate template
	template, err := loadImage(templatePath)
	if err != nil {
		fmt.Printf("Error loading template: %v\n", err)
		return
	}

	// Load the font
	fontBytes, err := os.ReadFile(fontPath)
	if err != nil {
		fmt.Printf("Error loading font: %v\n", err)
		return
	}
	f, err := opentype.Parse(fontBytes)
	if err != nil {
		fmt.Printf("Error parsing font: %v\n", err)
		return
	}

	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		fmt.Printf("Error creating font face: %v\n", err)
		return
	}

	// Create output folder if it doesn't exist
	err = os.MkdirAll(outputFolder, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating output folder: %v\n", err)
		return
	}

	// Open the CSV file
	file, err := os.Open(csvPath)
	if err != nil {
		fmt.Printf("Error opening CSV file: %v\n", err)
		return
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading CSV: %v\n", err)
			continue
		}

		name := record[0] // Assuming the name is in the first column

		// Create a new image for each certificate
		dest := image.NewRGBA(template.Bounds())
		draw.Draw(dest, dest.Bounds(), template, image.Point{}, draw.Src)

		// Add the name to the certificate
		d := &font.Drawer{
			Dst:  dest,
			Src:  image.NewUniform(textColor),
			Face: face,
		}

		// Measure string width to center it horizontally
		width := d.MeasureString(name).Round()
		x := (fixed.I(dest.Bounds().Dx()) - fixed.I(width)) / 2

		// Set the dot position with centered X and configurable Y
		d.Dot = fixed.Point26_6{
			X: x,
			Y: fixed.Int26_6(textY * 64),
		}

		d.DrawString(name)

		// Save the new certificate
		outputPath := filepath.Join(outputFolder, fmt.Sprintf("%s_certificate.png", name))
		err = saveImage(outputPath, dest)
		if err != nil {
			fmt.Printf("Error saving certificate for %s: %v\n", name, err)
			continue
		}

		fmt.Printf("Created certificate for %s\n", name)
	}

	fmt.Println("Certificate generation complete!")
}

func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func saveImage(path string, img image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}
