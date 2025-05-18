package main

import "fmt"

type DataUdara struct {
	Kota      string
	Kecamatan string
	AQI       int
	Sumber    string
	Kategori  string
	Tanggal   string
}

var databaseUdara []DataUdara

func main() {
	fmt.Println("==============================================")
	fmt.Println("  SISTEM PEMANTAUAN KUALITAS UDARA LOKAL")
	fmt.Println("==============================================")
	fmt.Println("Dibuat oleh: [Nama Mahasiswa]")
	fmt.Println("NIM: [NIM Mahasiswa]")
	fmt.Println("Mata Kuliah: Algoritma dan Pemrograman 2")
	fmt.Println("==============================================")

	tampilkanMenuUtama()
}

func tampilkanMenuUtama() {
	for {
		fmt.Println("\nMENU UTAMA:")
		fmt.Println("1. Tambah Data Kualitas Udara")
		fmt.Println("2. Lihat Semua Data")
		fmt.Println("3. Edit Data")
		fmt.Println("4. Hapus Data")
		fmt.Println("5. Cari Data")
		fmt.Println("6. Analisis Data")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih menu: ")
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahData()
		case 2:
			tampilkanSemuaData()
		case 3:
			editData()
		case 4:
			hapusData()
		case 5:
			menuPencarian()
		case 6:
			menuAnalisis()
		case 0:
			fmt.Println("Terima kasih telah menggunakan sistem ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func tambahData() {
	var dataBaru DataUdara

	fmt.Println("\n[TAMBAH DATA BARU]")
	fmt.Print("Kota: ")
	fmt.Scanln(&dataBaru.Kota)
	fmt.Print("Kecamatan: ")
	fmt.Scanln(&dataBaru.Kecamatan)
	fmt.Print("Nilai AQI: ")
	fmt.Scanln(&dataBaru.AQI)
	fmt.Print("Sumber polusi: ")
	fmt.Scanln(&dataBaru.Sumber)
	fmt.Print("Tanggal (DD-MM-YYYY): ")
	fmt.Scanln(&dataBaru.Tanggal)

	dataBaru.Kategori = tentukanKategori(dataBaru.AQI)

	databaseUdara = append(databaseUdara, dataBaru)
	fmt.Println("\nData berhasil ditambahkan!")
	fmt.Printf("Kategori: %s\n", dataBaru.Kategori)

	if dataBaru.AQI > 100 {
		fmt.Printf("⚠ PERINGATAN: Kualitas udara di %s %s termasuk %s!\n",
			dataBaru.Kota, dataBaru.Kecamatan, dataBaru.Kategori)
	}
}

func tentukanKategori(aqi int) string {
	switch {
	case aqi <= 50:
		return "Baik"
	case aqi <= 100:
		return "Sedang"
	case aqi <= 150:
		return "Tidak Sehat untuk Sensitif"
	case aqi <= 200:
		return "Tidak Sehat"
	case aqi <= 300:
		return "Sangat Tidak Sehat"
	default:
		return "Berbahaya"
	}
}

func tampilkanSemuaData() {
	if len(databaseUdara) == 0 {
		fmt.Println("\nBelum ada data yang tersimpan.")
		return
	}

	fmt.Println("\n[DATA KUALITAS UDARA]")
	fmt.Println("======================================================================================================")
	fmt.Printf("%-3s | %-15s | %-15s | %-5s | %-20s | %-15s | %-10s\n",
		"No", "Kota", "Kecamatan", "AQI", "Sumber", "Kategori", "Tanggal")
	fmt.Println("======================================================================================================")

	for i, data := range databaseUdara {
		fmt.Printf("%-3d | %-15s | %-15s | %-5d | %-20s | %-15s | %-10s\n",
			i+1, data.Kota, data.Kecamatan, data.AQI, data.Sumber, data.Kategori, data.Tanggal)
	}
	fmt.Println("======================================================================================================")
}

func editData() {
	if len(databaseUdara) == 0 {
		fmt.Println("\nBelum ada data yang bisa diedit.")
		return
	}

	tampilkanSemuaData()
	fmt.Print("\nMasukkan nomor data yang akan diedit: ")
	var nomor int
	fmt.Scanln(&nomor)

	if nomor < 1 || nomor > len(databaseUdara) {
		fmt.Println("Nomor tidak valid!")
		return
	}

	data := &databaseUdara[nomor-1]

	fmt.Println("\n[EDIT DATA] (Kosongkan jika tidak ingin mengubah)")
	fmt.Printf("Kota (%s): ", data.Kota)
	fmt.Scanln(&data.Kota)
	fmt.Printf("Kecamatan (%s): ", data.Kecamatan)
	fmt.Scanln(&data.Kecamatan)

	var aqiStr string
	fmt.Printf("AQI (%d): ", data.AQI)
	fmt.Scanln(&aqiStr)
	if aqiStr != "" {
		fmt.Sscanln(aqiStr, &data.AQI)
		data.Kategori = tentukanKategori(data.AQI)
	}

	fmt.Printf("Sumber (%s): ", data.Sumber)
	fmt.Scanln(&data.Sumber)
	fmt.Printf("Tanggal (%s): ", data.Tanggal)
	fmt.Scanln(&data.Tanggal)

	fmt.Println("\nData berhasil diperbarui!")
}

func hapusData() {
	if len(databaseUdara) == 0 {
		fmt.Println("\nBelum ada data yang bisa dihapus.")
		return
	}

	tampilkanSemuaData()
	fmt.Print("\nMasukkan nomor data yang akan dihapus: ")
	var nomor int
	fmt.Scanln(&nomor)

	if nomor < 1 || nomor > len(databaseUdara) {
		fmt.Println("Nomor tidak valid!")
		return
	}

	data := databaseUdara[nomor-1]
	fmt.Printf("\nAnda akan menghapus data:\nKota: %s\nKecamatan: %s\nAQI: %d\n",
		data.Kota, data.Kecamatan, data.AQI)

	fmt.Print("Yakin ingin menghapus? (y/n): ")
	var konfirmasi string
	fmt.Scanln(&konfirmasi)

	if konfirmasi == "y" || konfirmasi == "Y" {
		databaseUdara = append(databaseUdara[:nomor-1], databaseUdara[nomor:]...)
		fmt.Println("Data berhasil dihapus!")
	} else {
		fmt.Println("Penghapusan dibatalkan.")
	}
}

func menuPencarian() {
	for {
		fmt.Println("\n[MENU PENCARIAN]")
		fmt.Println("1. Cari berdasarkan Kota")
		fmt.Println("2. Cari berdasarkan Kategori")
		fmt.Println("3. Cari AQI > nilai tertentu")
		fmt.Println("0. Kembali")

		fmt.Print("Pilih menu: ")
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			cariBerdasarkanKota()
		case 2:
			cariBerdasarkanKategori()
		case 3:
			cariBerdasarkanAQI()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func cariBerdasarkanKota() {
	fmt.Print("\nMasukkan nama kota yang dicari: ")
	var kota string
	fmt.Scanln(&kota)

	var hasil []DataUdara
	for _, data := range databaseUdara {
		if stringsContains(data.Kota, kota) {
			hasil = append(hasil, data)
		}
	}

	if len(hasil) == 0 {
		fmt.Println("Tidak ditemukan data untuk kota tersebut.")
		return
	}

	fmt.Println("\nHasil pencarian:")
	fmt.Println("==============================================")
	for _, data := range hasil {
		fmt.Printf("Kota: %s\nKecamatan: %s\nAQI: %d (%s)\nSumber: %s\nTanggal: %s\n\n",
			data.Kota, data.Kecamatan, data.AQI, data.Kategori, data.Sumber, data.Tanggal)
	}
}

func cariBerdasarkanKategori() {
	fmt.Println("\nPilih kategori:")
	fmt.Println("1. Baik (0-50)")
	fmt.Println("2. Sedang (51-100)")
	fmt.Println("3. Tidak Sehat untuk Sensitif (101-150)")
	fmt.Println("4. Tidak Sehat (151-200)")
	fmt.Println("5. Sangat Tidak Sehat (201-300)")
	fmt.Println("6. Berbahaya (300+)")

	fmt.Print("Pilihan: ")
	var pilihan int
	fmt.Scanln(&pilihan)

	var kategori string
	switch pilihan {
	case 1:
		kategori = "Baik"
	case 2:
		kategori = "Sedang"
	case 3:
		kategori = "Tidak Sehat untuk Sensitif"
	case 4:
		kategori = "Tidak Sehat"
	case 5:
		kategori = "Sangat Tidak Sehat"
	case 6:
		kategori = "Berbahaya"
	default:
		fmt.Println("Pilihan tidak valid!")
		return
	}

	var hasil []DataUdara
	for _, data := range databaseUdara {
		if data.Kategori == kategori {
			hasil = append(hasil, data)
		}
	}

	if len(hasil) == 0 {
		fmt.Println("Tidak ditemukan data dengan kategori tersebut.")
		return
	}

	fmt.Printf("\nHasil pencarian (%s):\n", kategori)
	fmt.Println("==============================================")
	for _, data := range hasil {
		fmt.Printf("Kota: %s (%s)\nAQI: %d\nTanggal: %s\n\n",
			data.Kota, data.Kecamatan, data.AQI, data.Tanggal)
	}
}

func cariBerdasarkanAQI() {
	fmt.Print("\nMasukkan nilai AQI minimum: ")
	var minAQI int
	fmt.Scanln(&minAQI)

	var hasil []DataUdara
	for _, data := range databaseUdara {
		if data.AQI >= minAQI {
			hasil = append(hasil, data)
		}
	}

	if len(hasil) == 0 {
		fmt.Printf("Tidak ditemukan data dengan AQI ≥ %d.\n", minAQI)
		return
	}

	fmt.Printf("\nData dengan AQI ≥ %d:\n", minAQI)
	fmt.Println("==============================================")
	for _, data := range hasil {
		fmt.Printf("Kota: %s\nAQI: %d (%s)\nTanggal: %s\n\n",
			data.Kota, data.AQI, data.Kategori, data.Tanggal)
	}
}

func menuAnalisis() {
	for {
		fmt.Println("\n[MENU ANALISIS]")
		fmt.Println("1. Tampilkan wilayah terpolusi")
		fmt.Println("2. Statistik kualitas udara")
		fmt.Println("3. Rekomendasi berdasarkan AQI")
		fmt.Println("0. Kembali")

		fmt.Print("Pilih menu: ")
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tampilkanWilayahTerpolusi()
		case 2:
			tampilkanStatistik()
		case 3:
			berikanRekomendasi()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func tampilkanWilayahTerpolusi() {
	if len(databaseUdara) == 0 {
		fmt.Println("\nBelum ada data untuk dianalisis.")
		return
	}

	dataTerurut := make([]DataUdara, len(databaseUdara))
	copy(dataTerurut, databaseUdara)

	for i := 0; i < len(dataTerurut)-1; i++ {
		maxIndex := i
		for j := i + 1; j < len(dataTerurut); j++ {
			if dataTerurut[j].AQI > dataTerurut[maxIndex].AQI {
				maxIndex = j
			}
		}
		dataTerurut[i], dataTerurut[maxIndex] = dataTerurut[maxIndex], dataTerurut[i]
	}

	fmt.Println("\n5 WILAYAH DENGAN POLUSI TERTINGGI:")
	fmt.Println("==============================================")
	for i := 0; i < 5 && i < len(dataTerurut); i++ {
		data := dataTerurut[i]
		fmt.Printf("%d. %s (%s) - AQI: %d (%s)\n",
			i+1, data.Kota, data.Kecamatan, data.AQI, data.Kategori)
	}
	fmt.Println("==============================================")
}

func tampilkanStatistik() {
	if len(databaseUdara) == 0 {
		fmt.Println("\nBelum ada data untuk dianalisis.")
		return
	}

	total := len(databaseUdara)
	var totalAQI int
	kategoriCount := make(map[string]int)

	for _, data := range databaseUdara {
		totalAQI += data.AQI
		kategoriCount[data.Kategori]++
	}

	rataRata := float64(totalAQI) / float64(total)

	fmt.Println("\nSTATISTIK KUALITAS UDARA:")
	fmt.Println("==============================================")
	fmt.Printf("Jumlah data: %d\n", total)
	fmt.Printf("Rata-rata AQI: %.2f\n", rataRata)
	fmt.Println("----------------------------------------------")
	fmt.Println("Distribusi Kategori:")
	for kategori, jumlah := range kategoriCount {
		persentase := float64(jumlah) / float64(total) * 100
		fmt.Printf("- %s: %d (%.1f%%)\n", kategori, jumlah, persentase)
	}
	fmt.Println("==============================================")
}

func berikanRekomendasi() {
	fmt.Println("\nREKOMENDASI BERDASARKAN KUALITAS UDARA:")
	fmt.Println("==============================================")
	fmt.Println("1. Baik (0-50):")
	fmt.Println("   - Aktivitas luar ruangan aman untuk semua orang")
	fmt.Println("   - Pertahankan kualitas udara dengan menanam pohon")

	fmt.Println("\n2. Sedang (51-100):")
	fmt.Println("   - Masih aman untuk kebanyakan orang")
	fmt.Println("   - Orang sensitif mungkin perlu mengurangi aktivitas berat")

	fmt.Println("\n3. Tidak Sehat untuk Sensitif (101-150):")
	fmt.Println("   - Kelompok sensitif (anak, lansia, penderita asma) harus mengurangi aktivitas luar")
	fmt.Println("   - Gunakan masker jika diperlukan")

	fmt.Println("\n4. Tidak Sehat (151-200):")
	fmt.Println("   - Semua orang mungkin mulai merasakan efek kesehatan")
	fmt.Println("   - Hindari aktivitas luar ruangan yang lama")
	fmt.Println("   - Tutup jendela untuk mengurangi paparan")

	fmt.Println("\n5. Sangat Tidak Sehat (201-300):")
	fmt.Println("   - Peringatan kesehatan darurat")
	fmt.Println("   - Hindari semua aktivitas luar ruangan")
	fmt.Println("   - Gunakan air purifier di dalam ruangan")

	fmt.Println("\n6. Berbahaya (300+):")
	fmt.Println("   - Kondisi darurat")
	fmt.Println("   - Tetap di dalam ruangan dengan ventilasi tertutup")
	fmt.Println("   - Pertimbangkan untuk evakuasi jika memungkinkan")
	fmt.Println("==============================================")
}

func stringsContains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		match := true
		for j := 0; j < len(substr); j++ {
			if toLower(s[i+j]) != toLower(substr[j]) {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}
