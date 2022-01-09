package main

func notaParaConceito (n float64) string {
	var nota = int(n)

	// a switch uses an integer value
	// switch statements only fallthrough
	// with the use of the fallthrough keyword
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8,7:
		return "B"
	case 6,5:
		return "C"
	case 4,3:
		return "D"
	case 2,1,0:
		return "E"
	default: 
		return "Nota inválida"
	}
}
