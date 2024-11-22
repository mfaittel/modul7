package main

import "fmt"

func main() {
	var klubA, klubB string
	var pemenang []string

	fmt.Print("Input nama Klub A: ")
	fmt.Scanln(&klubA)

	fmt.Print("Input nama Klub B: ")
	fmt.Scanln(&klubB)

	fmt.Println("Input skor pertandingan. Masukkan skor negatif untuk berhenti.")

	for i := 1; ; i++ {
		var skorA, skorB int
		fmt.Printf("Pertandingan %d: ", i)

		n, err := fmt.Scan(&skorA, &skorB)
		if err != nil || n != 2 {
			fmt.Println("Format input tidak valid. Masukkan dua skor dipisahkan spasi.")
			fmt.Scanln()
			i--
			continue
		}

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}
	}

	fmt.Println("\nHasil Pertandingan:")
	for i, hasil := range pemenang {
		fmt.Printf("Hasil %d: %s\n", i+1, hasil)
	}
	fmt.Println("Pertandingan selesai")
}
