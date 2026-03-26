package main

import (
	"flag"
	"fmt"
	"image/color"
	"os"

	"github.com/skip2/go-qrcode"
)

const (
	disableBorder = true
	level         = qrcode.Medium
)

var (
	backgroundColor = color.White
	foregroundColor = color.Black
)

func main() {
	// CLI Arguments
	output := flag.String("o", "", "Chemin du fichier image de sortie (PNG)")
	size := flag.Int("s", 0, "Taille de l'image en pixels")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: qrcode-generator [options] <texte>\n\n")
		fmt.Fprintf(os.Stderr, "Génère un QR Code à partir d'un texte et l'enregistre en PNG.\n\n")
		fmt.Fprintf(os.Stderr, "Exemple:\n")
		fmt.Fprintf(os.Stderr, "  qrcode-generator -o /path/to/qrcode.png -s 512 \"https://example.com\"\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *output == "" || *size <= 0 || flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	text := flag.Arg(0)

	// QR code creation
	qr, err := qrcode.New(text, level)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
		os.Exit(1)
	}
	qr.DisableBorder = disableBorder
	qr.BackgroundColor = backgroundColor
	qr.ForegroundColor = foregroundColor

	if err := qr.WriteFile(*size, *output); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
		os.Exit(1)
	}

	// fmt.Printf("QR Code généré : %s (%dx%dpx)\n", *output, *size, *size)
}
