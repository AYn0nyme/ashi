package main

func Has(x []string, y string) bool {
	z := false
	for _, v := range x {
		if v == y {
			z = true
		}
	}
	return z
}
