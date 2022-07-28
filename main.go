package main

import (
	"fmt"
	"strings"
)

func main() {
	// soal nomer 1
	HitungPecahan(145000)

	// soal 2
	str := "telkom"
	str2 := "tlkom"
	CountDifferenctArray(str, str2)
	str3 := "telkom"
	str4 := "telecom"
	CountDifferenctArray(str3, str4)

	// nomer 3
	// kesalahan seharusnya menggunakan EXPOSE bukan listen
	// perlu ditambahkan CMD ["./main"]

}

func HitungPecahan(total int) {
	// pecahan := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	var seratusRibu int
	var limaPuluhRibu int
	var duaPuluhRibu int
	var sepuluhRibu int
	var limaRibu int
	var duaRibu int
	var seribu int
	var limaRatus int
	var duaRatus int
	var seratus int

	dataa := total % 100000
	seratusRibu = (total - dataa) / 100000

	datab := dataa % 50000
	limaPuluhRibu = (dataa - datab) / 50000

	datac := datab % 20000
	duaPuluhRibu = (datab - datac) / 20000

	datad := datac % 10000
	sepuluhRibu = (datac - datad) / 10000

	datae := datad % 5000
	limaRibu = (datad - datae) / 5000

	dataf := datae % 2000
	duaRibu = (datae - dataf) / 2000

	datag := dataf % 1000
	seribu = (dataf - datag) / 1000

	datah := datag % 500
	limaRatus = (datag - datah) / 500

	datai := datah % 200
	duaRatus = (datah - datai) / 200

	dataj := datai % 100
	seratus = (datai - dataj) / 100

	result := make(map[string]int)

	if seratusRibu > 0 {
		// fmt.Println(seratusRibu)
		result["Rp. 100.000"] = seratusRibu
	}
	if limaPuluhRibu > 0 {
		result["Rp. 50.000"] = limaPuluhRibu
	}
	if duaPuluhRibu > 0 {
		result["Rp. 20.000"] = duaPuluhRibu
	}
	if sepuluhRibu > 0 {
		result["Rp. 10.000"] = sepuluhRibu
	}
	if limaRibu > 0 {
		result["Rp. 5.000"] = limaRibu
	}
	if duaRibu > 0 {
		result["Rp. 2.000"] = duaRibu
	}
	if seribu > 0 {
		result["Rp. 1.000"] = seribu
	}
	if limaRatus > 0 {
		result["Rp. 500"] = limaRatus
	}
	if duaRatus > 0 {
		result["Rp. 200"] = duaRatus
	}
	if seratus > 0 {
		result["Rp. 200"] = seratus
	}
	fmt.Println("Hasil soal nomer 1")
	fmt.Println(result)

}

func CountDifferenctArray(str1, str2 string) {
	arrayOfString := strings.Split(str1, "")
	arrayOfString2 := strings.Split(str2, "")
	a := ArrayIntersect(arrayOfString, arrayOfString2)
	fmt.Println("HASIL SOAL NOMER 2", a)
}

func ArrayIntersect(slice1 []string, slice2 []string) bool {
	var diff []string

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	if len(diff) > 1 {
		return false
	}
	return true

	// return diff
}
