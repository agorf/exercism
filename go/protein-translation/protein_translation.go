package protein

import (
	"errors"
)

var codonToAminoAcid = map[string]string{
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

var (
	ErrStop        = errors.New("encountered stop")
	ErrInvalidBase = errors.New("invalid base")
)

func FromRNA(rna string) ([]string, error) {
	aminoAcids := make([]string, 0)
	for i := 0; i < len(rna)-2; i += 3 {
		codon := rna[i : i+3]
		aminoAcid, err := FromCodon(codon)
		if err == ErrStop {
			break
		}
		if err != nil {
			return nil, err
		}
		aminoAcids = append(aminoAcids, aminoAcid)
	}
	return aminoAcids, nil
}

func FromCodon(codon string) (string, error) {
	aminoAcid, exists := codonToAminoAcid[codon]
	if !exists {
		return "", ErrInvalidBase
	}
	if aminoAcid == "STOP" {
		return "", ErrStop
	}
	return aminoAcid, nil
}
