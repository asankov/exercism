package protein

import "errors"

var (
	ErrStop        = errors.New("STOP")
	ErrInvalidBase = errors.New("Invalid Base")
	codons         = map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}
)

func FromCodon(codon string) (string, error) {
	res, ok := codons[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if res == "STOP" {
		return "", ErrStop
	}
	return res, nil
}

func FromRNA(rna string) ([]string, error) {
	res := []string{}
	start, end := 0, 3
	for end <= len(rna) {
		codon, err := FromCodon(rna[start:end])
		if err != nil {
			if err == ErrStop {
				return res, nil
			}
			return res, err
		}
		res = append(res, codon)
		start = end
		end += 3
	}

	return res, nil
}
