package protein

import "errors"

var ErrStop = errors.New("STOP")
var ErrInvalidBase = errors.New("invalid codon")

func FromRNA(rna string) ([]string, error) {
	var result []string

	for i := 0; i <= len(rna)-3; i += 3 {
		codon := rna[i : i+3]
		protein, err := FromCodon(codon)
		if errors.Is(err, ErrStop) {
			return result, nil
		}
		if err != nil {
			return nil, err
		}
		result = append(result, protein)
	}

	return result, nil
}

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
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
