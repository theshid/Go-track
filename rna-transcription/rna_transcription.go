package strand

//ToRNA is a function that takes a dna string and return an transcription of its equivalent rna
func ToRNA(dna string) string {
	rna := ""

	for i:=0;i<len(dna);i++{
       rna+= string(transcription[dna[i]])
	}
	return rna
}

var transcription = [256]byte{
	'G' :'C',
	'C' :'G',
	'T' :'A',
	'A' :'U',
}
