package strand

func ToRNA(dna string) string {
	rna := ""
	for _, n := range dna {
		switch n {
		case 'G':
			rna += "C"
		case 'C':
			rna += "G"
		case 'T':
			rna += "A"
		case 'A':
			rna += "U"
		}
	}
	return rna
}
