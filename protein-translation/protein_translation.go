package protein

import "errors"

var ErrStop = errors.New("Stop error")
var ErrInvalidBase = errors.New("Invalid Base")

func FromCodon(s string) (string, error) {
	if val ,ok := codons[s];ok {
		if val == "STOP" {
			return "",ErrStop
		}
		return val,nil
	}
	return "",ErrInvalidBase
}

func FromRNA(s string) ([]string, error) {
	codon := Chunks(s)
	var protein []string

	for _ , val := range codon {
		proteinString,err :=FromCodon(val)
		if err == ErrInvalidBase {
			return protein,ErrInvalidBase
		} else if err == ErrStop {
			return protein,nil
		}
		protein = append(protein,proteinString)
	}

	return protein,nil
}


var codons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func Chunks(s string) []string {
	chunkSize := 3
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string
	chunk := make([]rune, chunkSize)
	len := 0
	for _, r := range s {
		chunk[len] = r
		len++
		if len == chunkSize {
			chunks = append(chunks, string(chunk))
			len = 0
		}
	}
	if len > 0 {
		chunks = append(chunks, string(chunk[:len]))
	}
	return chunks
}
//Best Solution
/*
var (
	STOP           error = errors.New("STOP")
	ErrInvalidBase error = errors.New("Invalid codon")
)

// Convert a codon into a protein
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", STOP
	default:
		return "", ErrInvalidBase
	}
}

// Convert an RNA strand into an array of proteins
func FromRNA(strand string) ([]string, error) {
	proteins := make([]string, 0, len(strand)/3)

	for i := 0; i < len(strand); i += 3 {
		protein, err := FromCodon(strand[i : i+3])

		if err == STOP {
			break
		}

		if err != nil {
			return proteins, err
		}

		proteins = append(proteins, protein)
	}

	return proteins, nil
}
*/