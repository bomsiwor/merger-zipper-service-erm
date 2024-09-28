package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"

	"github.com/jung-kurt/gofpdf"
	"golang.org/x/image/draw"
)

const (
	a4WidthMM  = 210.0
	a4HeightMM = 297.0
	dpi        = 72
)

func ConvertImageToPdf(src []string, output string) error {
	err := converter(src, output)

	if err != nil {
		fmt.Printf("Error converting to PDF : %v", err.Error())
		return err
	}

	return nil
}

func converter(imgPaths []string, output string) error {
	// Create a new PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	for _, imagePath := range imgPaths {
		file, err := os.Open(imagePath)
		if err != nil {
			return fmt.Errorf("failed to open image: %w", err)
		}
		defer file.Close()

		// Validate MIME type
		mimeType, err := GetFileMimeType(file)
		if err != nil {
			return fmt.Errorf("failed to get MIME type: %w", err)
		}

		if !IsValidImageMime(mimeType) {
			return fmt.Errorf("invalid MIME type: %s", mimeType)
		}

		// Decode image
		img, format, err := decodeImage(file, mimeType)
		if err != nil {
			return fmt.Errorf("failed to decode image: %w", err)
		}

		// Get image bounds and check for landscape
		width := img.Bounds().Dx()
		height := img.Bounds().Dy()

		// Determine if the image is landscape
		isLandscape := width > height

		// Resize image if it exceeds A4 dimensions
		maxWidthPx, maxHeightPx := mmToPx(a4WidthMM, dpi), mmToPx(a4HeightMM, dpi)

		if isLandscape {
			maxWidthPx, maxHeightPx = maxHeightPx, maxWidthPx
		}

		resizedImg := resizeImage(img, maxWidthPx, maxHeightPx)

		// Create a new page for each image
		if isLandscape {
			pdf.AddPageFormat("L", gofpdf.SizeType{Wd: a4HeightMM, Ht: a4WidthMM})
		} else {
			pdf.AddPageFormat("P", gofpdf.SizeType{Wd: a4WidthMM, Ht: a4HeightMM})
		}

		// Save the resized image to a buffer
		var imgBuf bytes.Buffer

		err = encodeImage(&imgBuf, resizedImg, format)

		if err != nil {
			return fmt.Errorf("failed to encode resized image: %w", err)
		}

		// Add image to PDF
		opt := gofpdf.ImageOptions{
			ImageType: format,
			ReadDpi:   false,
		}

		pdf.RegisterImageOptionsReader("img", opt, &imgBuf)

		pdf.ImageOptions("img", 0, 0, a4WidthMM, 0, false, opt, 0, "")
	}

	// Output the PDF
	err := pdf.OutputFileAndClose(output)

	if err != nil {
		return fmt.Errorf("failed to output PDF: %w", err)
	}

	return nil
}

// decodeImage decodes an image based on MIME type
func decodeImage(r io.Reader, mimeType string) (image.Image, string, error) {
	switch mimeType {
	case "image/png":
		img, err := png.Decode(r)
		return img, "png", err

	case "image/jpeg":
		img, err := jpeg.Decode(r)
		return img, "jpeg", err

	default:
		return nil, "", fmt.Errorf("unsupported image type: %s", mimeType)
	}
}

// encodeImage encodes an image to the specified format
func encodeImage(w io.Writer, img image.Image, format string) error {
	switch format {
	case "png":
		return png.Encode(w, img)

	case "jpeg":
		return jpeg.Encode(w, img, nil)

	default:
		return fmt.Errorf("unsupported image format: %s", format)

	}
}

// resizeImage resizes the image to fit within the specified maximum width and height
func resizeImage(img image.Image, maxWidth, maxHeight int) image.Image {
	srcBounds := img.Bounds()
	srcWidth := srcBounds.Dx()
	srcHeight := srcBounds.Dy()

	// Calculate aspect ratio
	ratio := float64(srcWidth) / float64(srcHeight)
	targetWidth := maxWidth
	targetHeight := maxHeight

	if ratio > 1.0 { // Landscape
		if srcWidth > maxWidth {
			targetWidth = maxWidth
			targetHeight = int(float64(maxWidth) / ratio)
		}
	} else { // Portrait
		if srcHeight > maxHeight {
			targetHeight = maxHeight
			targetWidth = int(float64(maxHeight) * ratio)
		}
	}

	dst := image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))
	draw.BiLinear.Scale(dst, dst.Bounds(), img, srcBounds, draw.Over, nil)

	return dst
}

// mmToPx converts millimeters to pixels based on DPI
func mmToPx(mm float64, dpi int) int {
	return int((mm / 25.4) * float64(dpi))
}
