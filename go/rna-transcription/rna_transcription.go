package strand

func ToRNA(dna string) string {
	replacements := map[rune]string{
		'G': "C",
		'C': "G",
		'T': "A",
		'A': "U",
	}

	result := ""
	for _, c := range dna {
		result += replacements[c]
	}

	return result
}
