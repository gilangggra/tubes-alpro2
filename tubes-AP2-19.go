package main

import "fmt"

const NMAX = 2045

type DataUdara struct {
	Kota      string
	Kecamatan string
	AQI       int
	Sumber    string
	Kategori  string
	Tanggal   string
}

var databaseUdara [NMAX]DataUdara
var nData int

func main() {
	fmt.Println("╔══════════════════════════════════════════════╗")
	fmt.Println("║   SISTEM PEMANTAUAN KUALITAS UDARA LOKAL     ║")
	fmt.Println("╠══════════════════════════════════════════════╣")
	fmt.Println("║ Dibuat oleh:                                 ║")
	fmt.Println("║ • Gilang Ramadan (103012400304)              ║")
	fmt.Println("║ • Muh. Ishaq Afif Ismail (103012400418)      ║")
	fmt.Println("║ Mata Kuliah: Algoritma dan Pemrograman 2     ║")
	fmt.Println("╚══════════════════════════════════════════════╝")

	// Data dummy
	databaseUdara[0] = DataUdara{
		Kota:      "Bekasi",
		Kecamatan: "Bantargebang",
		AQI:       120,
		Sumber:    "Kendaraan Bermotor",
		Kategori:  tentukanKategori(120),
		Tanggal:   "01-06-2025",
	}
	databaseUdara[1] = DataUdara{
		Kota:      "Bandung",
		Kecamatan: "Cicendo",
		AQI:       85,
		Sumber:    "Industri",
		Kategori:  tentukanKategori(85),
		Tanggal:   "02-06-2025",
	}
	databaseUdara[2] = DataUdara{
		Kota:      "Surabaya",
		Kecamatan: "Wonokromo",
		AQI:       45,
		Sumber:    "Rumah Tangga",
		Kategori:  tentukanKategori(45),
		Tanggal:   "03-06-2025",
	}
	databaseUdara[3] = DataUdara{
		Kota:      "Medan",
		Kecamatan: "Medan Baru",
		AQI:       210,
		Sumber:    "Pembakaran Sampah",
		Kategori:  tentukanKategori(210),
		Tanggal:   "04-06-2025",
	}
	databaseUdara[4] = DataUdara{
		Kota:      "Makassar",
		Kecamatan: "Mamajang",
		AQI:       320,
		Sumber:    "Asap Pabrik",
		Kategori:  tentukanKategori(320),
		Tanggal:   "05-06-2025",
	}
	nData = 5

	tampilkanMenuUtama()
}

func stringsEqualFold(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		c1 := s1[i]
		if c1 >= 'A' && c1 <= 'Z' {
			c1 += 32
		}
		c2 := s2[i]
		if c2 >= 'A' && c2 <= 'Z' {
			c2 += 32
		}
		if c1 != c2 {
			return false
		}
	}
	return true
}

func tampilkanMenuUtama() {
	for selesai := false; !selesai; {
		fmt.Println("\n╔═══════════════════════════════╗")
		fmt.Println("║          MENU UTAMA           ║")
		fmt.Println("╠═══════════════════════════════╣")
		fmt.Println("║ 1. Tambah Data                ║")
		fmt.Println("║ 2. Lihat Semua Data           ║")
		fmt.Println("║ 3. Edit Data                  ║")
		fmt.Println("║ 4. Hapus Data                 ║")
		fmt.Println("║ 5. Cari Data                  ║")
		fmt.Println("║ 6. Analisis Data              ║")
		fmt.Println("║ 0. Keluar                     ║")
		fmt.Println("╚═══════════════════════════════╝")
		fmt.Print("Pilih menu: ")
		var pilihan int
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			tambahData()
		} else if pilihan == 2 {
			tampilkanSemuaData()
		} else if pilihan == 3 {
			editData()
		} else if pilihan == 4 {
			hapusData()
		} else if pilihan == 5 {
			menuPencarian()
		} else if pilihan == 6 {
			menuAnalisis()
		} else if pilihan == 0 {
			fmt.Println("\n╔════════════════════════════════╗")
			fmt.Println("║   Terima kasih telah           ║")
			fmt.Println("║   menggunakan sistem ini       ║")
			fmt.Println("╚════════════════════════════════╝")
			selesai = true
		} else {
			fmt.Println("\n⚠ Pilihan tidak valid!")
		}
	}
}

func tambahData() {
	if nData >= NMAX {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║     Database penuh!            ║")
		fmt.Println("║ Tidak bisa menambah data lagi  ║")
		fmt.Println("╚════════════════════════════════╝")
		return
	}

	var dataBaru DataUdara
	fmt.Println("\n╔════════════════════════════════╗")
	fmt.Println("║       TAMBAH DATA BARU         ║")
	fmt.Println("╚════════════════════════════════╝")
	fmt.Print(" Kota: ")
	fmt.Scanln(&dataBaru.Kota)
	fmt.Print(" Kecamatan: ")
	fmt.Scanln(&dataBaru.Kecamatan)
	fmt.Print(" Nilai AQI: ")
	fmt.Scanln(&dataBaru.AQI)
	fmt.Print(" Sumber polusi: ")
	fmt.Scanln(&dataBaru.Sumber)
	fmt.Print(" Tanggal (DD-MM-YYYY): ")
	fmt.Scanln(&dataBaru.Tanggal)

	dataBaru.Kategori = tentukanKategori(dataBaru.AQI)
	databaseUdara[nData] = dataBaru
	nData++

	fmt.Println("\n╔════════════════════════════════╗")
	fmt.Println("║   Data berhasil ditambahkan!   ║")
	fmt.Println("╠════════════════════════════════╣")
	fmt.Printf("║ Kategori: %-20s ║\n", dataBaru.Kategori)
	fmt.Println("╚════════════════════════════════╝")

	if dataBaru.AQI > 100 {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║           PERINGATAN!          ║")
		fmt.Println("╠════════════════════════════════╣")
		fmt.Printf("║ Kualitas udara di %-12s ║\n", dataBaru.Kota)
		fmt.Printf("║ %-s termasuk %-15s ║\n", dataBaru.Kecamatan, dataBaru.Kategori)
		fmt.Println("╚════════════════════════════════╝")
	}
}

func tentukanKategori(aqi int) string {
	switch {
	case aqi <= 50:
		return "Baik"
	case aqi <= 100:
		return "Sedang"
	case aqi <= 150:
		return "Tidak Sehat (Sensitif)"
	case aqi <= 200:
		return "Tidak Sehat"
	case aqi <= 300:
		return "Sangat Tidak Sehat"
	default:
		return "Berbahaya"
	}
}

func tampilkanSemuaData() {
	if nData == 0 {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║   Belum ada data tersimpan!    ║")
		fmt.Println("╚════════════════════════════════╝")
		return
	}

	fmt.Println("\n╔═══════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                           DATA KUALITAS UDARA                                             ║")
	fmt.Println("╠════╦═════════════╦═══════════════════╦═══════╦════════════════════╦══════════════════════════╦════════════╣")
	fmt.Printf("║ %-2s ║ %-10s  ║   %-15s ║ %-5s ║ %-18s ║ %-24s ║ %-10s ║\n", "No", "    Kota", "Kecamatan", " AQI", "       Sumber", "           Kategori", "  Tanggal")
	fmt.Println("╠════╬═════════════╬═══════════════════╬═══════╬════════════════════╬══════════════════════════╬════════════╣")

	for i := 0; i < nData; i++ {
		data := databaseUdara[i]
		fmt.Printf("║ %-2d ║ %-10s  ║ %-15s   ║ %-5d ║ %-18s ║ %-24s ║ %-10s ║\n", i+1, data.Kota, data.Kecamatan, data.AQI, data.Sumber, data.Kategori, data.Tanggal)
	}
	fmt.Println("╚════╩═════════════╩═══════════════════╩═══════╩════════════════════╩══════════════════════════╩════════════╝")
}

func editData() {
	if nData == 0 {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║   Belum ada data tersimpan!    ║")
		fmt.Println("╚════════════════════════════════╝")
		return
	}
	tampilkanSemuaData()
	var nomor int

	fmt.Print("\nMasukkan nomor data yang akan diedit: ")
	fmt.Scanln(&nomor)

	if nomor < 1 || nomor > nData {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║      Nomor tidak valid!        ║")
		fmt.Println("╚════════════════════════════════╝")
		return
	}

	data := &databaseUdara[nomor-1]
	fmt.Println("\n╔════════════════════════════════╗")
	fmt.Println("║          EDIT DATA             ║")
	fmt.Println("╠════════════════════════════════╣")
	fmt.Println("║ Kosongkan jika tidak diubah    ║")
	fmt.Println("╚════════════════════════════════╝")
	fmt.Printf("Kota (%s): ", data.Kota)
	fmt.Scanln(&data.Kota)
	fmt.Printf("Kecamatan (%s): ", data.Kecamatan)
	fmt.Scanln(&data.Kecamatan)
	fmt.Printf("AQI (%d): ", data.AQI)
	fmt.Scanln(&data.AQI)
	data.Kategori = tentukanKategori(data.AQI)
	fmt.Printf("Sumber (%s): ", data.Sumber)
	fmt.Scanln(&data.Sumber)
	fmt.Printf("Tanggal (%s): ", data.Tanggal)
	fmt.Scanln(&data.Tanggal)

	fmt.Println("\n╔════════════════════════════════╗")
	fmt.Println("║   Data berhasil diperbarui!    ║")
	fmt.Println("╚════════════════════════════════╝")
}

func hapusData() {
	if nData == 0 {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║   Belum ada data tersimpan!    ║")
		fmt.Println("╚════════════════════════════════╝")
		return
	}

	tampilkanSemuaData()
	var nomor int

	fmt.Print("\nMasukkan nomor data yang akan dihapus: ")
	fmt.Scanln(&nomor)

	if nomor < 1 || nomor > nData {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║      Nomor tidak valid!        ║")
		fmt.Println("╚════════════════════════════════╝")
		return
	}

	data := databaseUdara[nomor-1]
	fmt.Println("\n╔════════════════════════════════╗")
	fmt.Println("║      KONFIRMASI HAPUS          ║")
	fmt.Println("╠════════════════════════════════╣")
	fmt.Printf("║ Kota: %-24s ║\n", data.Kota)
	fmt.Printf("║ Kecamatan: %-19s ║\n", data.Kecamatan)
	fmt.Printf("║ AQI: %-25d ║\n", data.AQI)
	fmt.Println("╚═════════════════════════════════╝")
	fmt.Print("Yakin ingin menghapus? (ya/tidak): ")
	var konfirmasi string
	fmt.Scanln(&konfirmasi)

	if konfirmasi == "ya" || konfirmasi == "Ya" {
		for i := nomor - 1; i < nData-1; i++ {
			databaseUdara[i] = databaseUdara[i+1]
		}
		nData--
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║   Data berhasil dihapus!       ║")
		fmt.Println("╚════════════════════════════════╝")
	} else {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║   Penghapusan dibatalkan!      ║")
		fmt.Println("╚════════════════════════════════╝")
	}
}

func insertionSortByAQI(arr *[NMAX]DataUdara, jumlah int) {
	for i := 1; i < jumlah; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j].AQI > key.AQI {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func binarySearchAQI(arr *[NMAX]DataUdara, jumlah, target int) int {
	low := 0
	high := jumlah - 1
	idx := -1
	for low <= high && idx == -1 {
		mid := (low + high) / 2
		if arr[mid].AQI == target {
			idx = mid
		} else if arr[mid].AQI < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return idx
}

func menuPencarian() {
	for selesai := false; !selesai; {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║        MENU PENCARIAN          ║")
		fmt.Println("╠════════════════════════════════╣")
		fmt.Println("║ 1. Cari berdasarkan Kota       ║")
		fmt.Println("║ 2. Cari berdasarkan Kategori   ║")
		fmt.Println("║ 3. Cari AQI > nilai tertentu   ║")
		fmt.Println("║ 4. Cari berdasarkan AQI        ║")
		fmt.Println("║ 0. Kembali                     ║")
		fmt.Println("╚════════════════════════════════╝")
		fmt.Print("Pilih menu: ")
		var pilihan int
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			cariBerdasarkanKota()
		} else if pilihan == 2 {
			cariBerdasarkanKategori()
		} else if pilihan == 3 {
			cariBerdasarkanAQI()
		} else if pilihan == 4 {
			cariBerdasarkanAQIBinary()
		} else if pilihan == 0 {
			selesai = true
		} else {
			fmt.Println("\n⚠ Pilihan tidak valid!")
		}
	}
}

func cariBerdasarkanKota() {
	var kota string

	fmt.Print("\nMasukkan nama kota yang dicari: ")
	fmt.Scanln(&kota)

	ada := false
	for i := 0; i < nData; i++ {
		if stringsEqualFold(databaseUdara[i].Kota, kota) {
			if !ada {
				fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
				fmt.Println("║                      HASIL PENCARIAN                           ║")
				fmt.Println("╠════════════════════════════════════════════════════════════════╣")
				ada = true
			}
			data := databaseUdara[i]
			fmt.Printf("║ Kota: %-56s ║\n", data.Kota)
			fmt.Printf("║ Kecamatan: %-51s ║\n", data.Kecamatan)
			fmt.Printf("║ AQI: %-4d (%s)%28s ║\n", data.AQI, data.Kategori, "")
			fmt.Printf("║ Sumber: %-54s ║\n", data.Sumber)
			fmt.Printf("║ Tanggal: %-53s ║\n", data.Tanggal)
			fmt.Println("╚════════════════════════════════════════════════════════════════╝")
		}
	}

	if !ada {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║   Data tidak ditemukan!        ║")
		fmt.Println("╚════════════════════════════════╝")
	}

}

func cariBerdasarkanKategori() {
	fmt.Println("\n╔════════════════════════════════╗")
	fmt.Println("║        PILIH KATEGORI          ║")
	fmt.Println("╠════════════════════════════════╣")
	fmt.Println("║ 1. Baik (0-50)                 ║")
	fmt.Println("║ 2. Sedang (51-100)             ║")
	fmt.Println("║ 3. Tidak Sehat (Sensitif)      ║")
	fmt.Println("║ 4. Tidak Sehat (151-200)       ║")
	fmt.Println("║ 5. Sangat Tidak Sehat (201-300)║")
	fmt.Println("║ 6. Berbahaya (300+)            ║")
	fmt.Println("╚════════════════════════════════╝")
	fmt.Print("Pilihan: ")

	var pilihan int
	fmt.Scanln(&pilihan)

	var kategori string
	if pilihan == 1 {
		kategori = "Baik"
	} else if pilihan == 2 {
		kategori = "Sedang"
	} else if pilihan == 3 {
		kategori = "Tidak Sehat (Sensitif)"
	} else if pilihan == 4 {
		kategori = "Tidak Sehat"
	} else if pilihan == 5 {
		kategori = "Sangat Tidak Sehat"
	} else if pilihan == 6 {
		kategori = "Berbahaya"
	} else {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║      Pilihan tidak valid!      ║")
		fmt.Println("╚════════════════════════════════╝")
		return
	}

	ada := false
	for i := 0; i < nData; i++ {
		if databaseUdara[i].Kategori == kategori {
			if !ada {
				fmt.Printf("\n╔════════════════════════════════════════════════════════╗\n")
				fmt.Printf("║ HASIL PENCARIAN (%s)%18s               ║\n", kategori, "")
				fmt.Println("╠════════════════════════════════════════════════════════╣")
				ada = true
			}
			data := databaseUdara[i]
			fmt.Printf("║ Kota: %-15s (%s)%20s  ║\n", data.Kota, data.Kecamatan, "")
			fmt.Printf("║ AQI: %-3d %-45s ║\n", data.AQI, "")
			fmt.Printf("║ Tanggal: %-45s ║\n", data.Tanggal)
			fmt.Println("╚════════════════════════════════════════════════════════╝")
		}
	}

	if !ada {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║   Data tidak ditemukan!        ║")
		fmt.Println("╚════════════════════════════════╝")
	}
}

func cariBerdasarkanAQI() {
	var minAQI int

	fmt.Print("\nMasukkan nilai AQI minimum: ")
	fmt.Scanln(&minAQI)

	ada := false
	for i := 0; i < nData; i++ {
		if databaseUdara[i].AQI >= minAQI {
			if !ada {
				fmt.Printf("\n╔════════════════════════════════════════════════════════╗\n")
				fmt.Printf("║ DATA DENGAN AQI ≥ %-3d%33s ║\n", minAQI, "")
				fmt.Println("╠════════════════════════════════════════════════════════╣")
				ada = true
			}
			data := databaseUdara[i]
			fmt.Printf("║ Kota: %-48s ║\n", data.Kota)
			fmt.Printf("║ AQI: %-3d (%-10s)%21s ║\n", data.AQI, data.Kategori, "")
			fmt.Printf("║ Tanggal: %-45s ║\n", data.Tanggal)
			fmt.Println("╠════════════════════════════════════════════════════════╣")
		}
	}

	if !ada {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Printf("║ Tidak ada data AQI ≥ %-3d       ║\n", minAQI)
		fmt.Println("╚════════════════════════════════╝")
	} else {
		fmt.Println("║                [Ini Adalah Hasilnya]                   ║")
		fmt.Println("╚════════════════════════════════════════════════════════╝")
	}
}

func cariBerdasarkanAQIBinary() {
	if nData == 0 {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║   Belum ada data tersimpan!    ║")
		fmt.Println("╚════════════════════════════════╝")
		return
	}

	var dataTerurut [NMAX]DataUdara
	for i := 0; i < nData; i++ {
		dataTerurut[i] = databaseUdara[i]
	}
	insertionSortByAQI(&dataTerurut, nData)

	var target int
	fmt.Print("\nMasukkan nilai AQI yang dicari: ")
	fmt.Scanln(&target)

	index := binarySearchAQI(&dataTerurut, nData, target)
	if index != -1 {
		data := dataTerurut[index]
		fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
		fmt.Println("║                      DATA DITEMUKAN                            ║")
		fmt.Println("╠════════════════════════════════════════════════════════════════╣")
		fmt.Printf("║ Kota: %-56s ║\n", data.Kota)
		fmt.Printf("║ Kecamatan: %-51s ║\n", data.Kecamatan)
		fmt.Printf("║ AQI: %-4d (%s)%28s ║\n", data.AQI, data.Kategori, "")
		fmt.Printf("║ Sumber: %-54s ║\n", data.Sumber)
		fmt.Printf("║ Tanggal: %-53s ║\n", data.Tanggal)
		fmt.Println("╚════════════════════════════════════════════════════════════════╝")
	} else {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║   Data tidak ditemukan!        ║")
		fmt.Println("╚════════════════════════════════╝")
	}
}

func menuAnalisis() {
	for selesai := false; !selesai; {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║        MENU ANALISIS           ║")
		fmt.Println("╠════════════════════════════════╣")
		fmt.Println("║ 1. Wilayah Terpolusi           ║")
		fmt.Println("║ 2. Wilayah Terpolusi (Periode) ║")
		fmt.Println("║ 3. Statistik Kualitas Udara    ║")
		fmt.Println("║ 4. Rekomendasi                 ║")
		fmt.Println("║ 5. Data Terurut (AQI)          ║")
		fmt.Println("║ 0. Kembali                     ║")
		fmt.Println("╚════════════════════════════════╝")
		fmt.Print("Pilih menu: ")
		var pilihan int
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			tampilkanWilayahTerpolusi()
		} else if pilihan == 2 {
			tampilkanWilayahTerpolusiPeriode()
		} else if pilihan == 3 {
			tampilkanStatistik()
		} else if pilihan == 4 {
			berikanRekomendasi()
		} else if pilihan == 5 {
			tampilkanDataTerurutInsertion()
		} else if pilihan == 0 {
			selesai = true
		} else {
			fmt.Println("\n⚠ Pilihan tidak valid!")
		}
	}
}

func tampilkanWilayahTerpolusi() {
	if nData == 0 {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║   Belum ada data tersimpan!    ║")
		fmt.Println("╚════════════════════════════════╝")
		return
	}

	var dataTerurut [NMAX]DataUdara
	for i := 0; i < nData; i++ {
		dataTerurut[i] = databaseUdara[i]
	}

	for i := 0; i < nData-1; i++ {
		maxIdx := i
		for j := i + 1; j < nData; j++ {
			if dataTerurut[j].AQI > dataTerurut[maxIdx].AQI {
				maxIdx = j
			}
		}
		temp := dataTerurut[i]
		dataTerurut[i] = dataTerurut[maxIdx]
		dataTerurut[maxIdx] = temp
	}

	fmt.Println("\n╔══════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║         5 WILAYAH DENGAN POLUSI TERTINGGI                                    ║")
	fmt.Println("╠════╦═════════════════════════════════════════════════════════════════════════╣")
	for i := 0; i < 5 && i < nData; i++ {
		data := dataTerurut[i]
		fmt.Printf("║ %-2d ║ %-10s (%-15s) - AQI: %-3d (%-30s)║\n", i+1, data.Kota, data.Kecamatan, data.AQI, data.Kategori)
	}
	fmt.Println("╚════╩═════════════════════════════════════════════════════════════════════════╝")
}

func tampilkanDataTerurutInsertion() {
	if nData == 0 {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║   Belum ada data tersimpan!    ║")
		fmt.Println("╚════════════════════════════════╝")
		return
	}
	var dataTerurut [NMAX]DataUdara
	for i := 0; i < nData; i++ {
		dataTerurut[i] = databaseUdara[i]
	}
	insertionSortByAQI(&dataTerurut, nData)

	fmt.Println("\n╔════╦════════════╦═════════════════╦═══════╦════════════════════════╦═════════════════════════╦════════════════╗")
	fmt.Println("║ No ║    Kota    ║    Kecamatan    ║  AQI  ║         Sumber         ║        Kategori         ║     Tanggal    ║")
	fmt.Println("╠════╬════════════╬═════════════════╬═══════╬════════════════════════╬═════════════════════════╬════════════════╣")

	for i := 0; i < nData; i++ {
		data := dataTerurut[i]
		fmt.Printf("║ %2d ║ %-10s ║ %-15s ║ %5d ║ %-22s ║ %-23s ║ %-14s ║\n", i+1, data.Kota, data.Kecamatan, data.AQI, data.Sumber, data.Kategori, data.Tanggal)
	}
	fmt.Println("╚════╩════════════╩═════════════════╩═══════╩════════════════════════╩═════════════════════════╩════════════════╝")
}

func tampilkanStatistik() {
	if nData == 0 {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║   Belum ada data tersimpan!    ║")
		fmt.Println("╚════════════════════════════════╝")
		return
	}

	total := nData
	var totalAQI int
	var kategoriCount [6]int
	var kategoriList = [6]string{
		"Baik", "Sedang", "Tidak Sehat (Sensitif)", "Tidak Sehat", "Sangat Tidak Sehat", "Berbahaya",
	}

	for i := 0; i < nData; i++ {
		totalAQI += databaseUdara[i].AQI
		for j := 0; j < 6; j++ {
			if databaseUdara[i].Kategori == kategoriList[j] {
				kategoriCount[j]++
			}
		}
	}
	rataRata := float64(totalAQI) / float64(total)

	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║               STATISTIK KUALITAS UDARA                 ║")
	fmt.Println("╠════════════════════════════════════════════════════════╣")
	fmt.Printf("║ Jumlah data: %-41d ║\n", total)
	fmt.Printf("║ Rata-rata AQI: %-39.2f ║\n", rataRata)
	fmt.Println("╠════════════════════════════════════════════════════════╣")
	fmt.Println("║              DISTRIBUSI KATEGORI                       ║")
	fmt.Println("╠════════════════════════════════════════════════════════╣")

	for j := 0; j < 6; j++ {
		jumlah := kategoriCount[j]
		if jumlah > 0 {
			persentase := float64(jumlah) / float64(total) * 100
			fmt.Printf("║ - %-23s: %3d (%5.1f%%)%15s ║\n", kategoriList[j], jumlah, persentase, "")
		}
	}
	fmt.Println("╚════════════════════════════════════════════════════════╝")
}

func bandingkanTanggal(tanggal, tanggalAwal, tanggalAkhir string) bool {

	day := tanggal[:2]
	month := tanggal[3:5]
	year := tanggal[6:]

	dayAwal := tanggalAwal[:2]
	monthAwal := tanggalAwal[3:5]
	yearAwal := tanggalAwal[6:]

	dayAkhir := tanggalAkhir[:2]
	monthAkhir := tanggalAkhir[3:5]
	yearAkhir := tanggalAkhir[6:]

	if year < yearAwal || year > yearAkhir {
		return false
	}

	if year == yearAwal && month < monthAwal {
		return false
	}
	if year == yearAkhir && month > monthAkhir {
		return false
	}

	if year == yearAwal && month == monthAwal && day < dayAwal {
		return false
	}
	if year == yearAkhir && month == monthAkhir && day > dayAkhir {
		return false
	}

	return true
}

func tampilkanWilayahTerpolusiPeriode() {
	if nData == 0 {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║   Belum ada data tersimpan!    ║")
		fmt.Println("╚════════════════════════════════╝")
		return
	}

	var tanggalAwal, tanggalAkhir string
	fmt.Print("\nMasukkan tanggal awal (DD-MM-YYYY): ")
	fmt.Scanln(&tanggalAwal)
	fmt.Print("Masukkan tanggal akhir (DD-MM-YYYY): ")
	fmt.Scanln(&tanggalAkhir)

	var dataDalamPeriode []DataUdara
	for i := 0; i < nData; i++ {
		if bandingkanTanggal(databaseUdara[i].Tanggal, tanggalAwal, tanggalAkhir) {
			dataDalamPeriode = append(dataDalamPeriode, databaseUdara[i])
		}
	}

	if len(dataDalamPeriode) == 0 {
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║ Tidak ada data dalam periode   ║")
		fmt.Println("║ yang ditentukan!              ║")
		fmt.Println("╚════════════════════════════════╝")
		return
	}

	for i := 0; i < len(dataDalamPeriode)-1; i++ {
		maxIdx := i
		for j := i + 1; j < len(dataDalamPeriode); j++ {
			if dataDalamPeriode[j].AQI > dataDalamPeriode[maxIdx].AQI {
				maxIdx = j
			}
		}
		dataDalamPeriode[i], dataDalamPeriode[maxIdx] = dataDalamPeriode[maxIdx], dataDalamPeriode[i]
	}

	jumlahTampil := 5
	if len(dataDalamPeriode) < 5 {
		jumlahTampil = len(dataDalamPeriode)
	}

	fmt.Printf("\n╔══════════════════════════════════════════════════════════════════════════════╗\n")
	fmt.Printf("║   %d WILAYAH DENGAN POLUSI TERTINGGI (%s - %s)%15s ║\n", jumlahTampil, tanggalAwal, tanggalAkhir, "")
	fmt.Println("╠════╦═════════════════════════════════════════════════════════════════════════╣")
	for i := 0; i < jumlahTampil; i++ {
		data := dataDalamPeriode[i]
		fmt.Printf("║ %-2d ║ %-10s (%-15s) - AQI: %-3d (%-30s)║\n", i+1, data.Kota, data.Kecamatan, data.AQI, data.Kategori)
	}
	fmt.Println("╚════╩═════════════════════════════════════════════════════════════════════════╝")
}

func berikanRekomendasi() {
	fmt.Println("\n╔═════════════════════════════════════════════════════════╗")
	fmt.Println("║               REKOMENDASI KUALITAS UDARA                ║")
	fmt.Println("╠═════════════════════════════════════════════════════════╣")
	fmt.Println("║ 1. Baik (0-50):                                         ║")
	fmt.Println("║    • Aktivitas luar ruangan aman untuk semua orang      ║")
	fmt.Println("║    • Pertahankan kualitas udara dengan menanam pohon    ║")
	fmt.Println("╠═════════════════════════════════════════════════════════╣")
	fmt.Println("║ 2. Sedang (51-100):                                     ║")
	fmt.Println("║    • Masih aman untuk kebanyakan orang                  ║")
	fmt.Println("║    • Orang sensitif kurangi aktivitas berat             ║")
	fmt.Println("╠═════════════════════════════════════════════════════════╣")
	fmt.Println("║ 3. Tidak Sehat (Sensitif):                              ║")
	fmt.Println("║    • Kelompok sensitif kurangi aktivitas luar           ║")
	fmt.Println("║    • Gunakan masker jika diperlukan                     ║")
	fmt.Println("╠═════════════════════════════════════════════════════════╣")
	fmt.Println("║ 4. Tidak Sehat (151-200):                               ║")
	fmt.Println("║    • Semua orang mungkin mulai merasakan efek           ║")
	fmt.Println("║    • Hindari aktivitas luar ruangan yang lama           ║")
	fmt.Println("╠═════════════════════════════════════════════════════════╣")
	fmt.Println("║ 5. Sangat Tidak Sehat (201-300):                        ║")
	fmt.Println("║    • Peringatan kesehatan darurat                       ║")
	fmt.Println("║    • Hindari semua aktivitas luar ruangan               ║")
	fmt.Println("╠═════════════════════════════════════════════════════════╣")
	fmt.Println("║ 6. Berbahaya (300+):                                    ║")
	fmt.Println("║    • Kondisi darurat                                    ║")
	fmt.Println("║    • Tetap di dalam ruangan dengan ventilasi tertutup   ║")
	fmt.Println("╚═════════════════════════════════════════════════════════╝")
}
