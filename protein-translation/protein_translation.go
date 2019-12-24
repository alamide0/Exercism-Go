package protein

import (
	"errors"
)

var (
	ErrStop        = errors.New("ErrStop")
	ErrInvalidBase = errors.New("ErrInvalidBase")
)

var codons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "", "UAG": "", "UGA": "",
}

func FromCodon(input string) (res string, err error) {
	if res, ok := codons[input]; ok {
		if res == "" {
			return "", ErrStop
		} else {
			return res, nil
		}
	} else {
		return "", ErrInvalidBase
	}
}

func FromRNA(input string) (res []string, err error) {
	for i := 0; i < len(input)-2; i += 3 {
		s, err := FromCodon(input[i : i+3])

		if err != nil {
			if err == ErrStop {
				return res, nil
			}
			if err == ErrInvalidBase {
				return res, err
			}
		} else {
			res = append(res, s)
		}
	}

	return res, err
}
