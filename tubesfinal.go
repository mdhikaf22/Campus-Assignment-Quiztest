// insertion sort - print_rank
// selection sort - print_soal_termudah
// sequential search - cek_pin,username,pw
// binary search-search_user

package main

import (
	"fmt"
	"math/rand"
)

const dataLimit int = 100

type formatSoal struct {
	question string
	option   [4]string
	correct  int
	kj       int
}

type formatUser struct {
	username     string
	fName, lName string
	id           int
	pw           string
	score        float64
	reward       string
}

type bankSoal struct {
	data  [dataLimit]formatSoal
	n     int
	nShow int //untuk menentukan brp soal yang mau ditunjukan di quiz
}

type userTab struct {
	data [dataLimit]formatUser
	n    int
}

func main() {
	var bank bankSoal
	var user userTab
	var input string
	use_temp(&bank, &user)
	for input != "c" {
		fmt.Println("===========================================================================")
		fmt.Println("|                          Selamat Datang Di                              |")
		fmt.Println("|                        Who One to Be a Millioner                        |")
		fmt.Println("|                 Mahardhika Fernanda // Indah Pratiwi                    |")
		fmt.Println("===========================================================================")
		fmt.Println("a. ADMIN")
		fmt.Println("b. KONTESTAN")
		fmt.Println("c. EXIT")
		fmt.Print("Masukan input (a/b/c) : ")
		fmt.Scan(&input)
		if input == "a" {
			menu_admin(&bank, &user)
		} else if input == "b" {
			menu_guest(&bank, &user)
		} else if input != "c" {
			fmt.Print("error, input tidak dikenali")
		}
	}
}

func menu_admin(bank *bankSoal, user *userTab) {
	var input string
	for input != "i" {
		fmt.Println("============================= ADMIN MENU =============================")
		fmt.Println("a. Tambahkan soal")
		fmt.Println("b. Hapus soal")
		fmt.Println("c. Edit soal")
		fmt.Println("d. Set jumlah pertanyaan")
		fmt.Println("e. Tampilkan soal dengan jawaban benar terbanyak")
		fmt.Println("f. Tampilkan peringkat Terkecil ke Terbesar")
		fmt.Println("g. Tampilkan peringkat Terbesar ke Terkecil")
		fmt.Println("h. Cari kontestan berdasarkan id")
		fmt.Println("i. Log out")
		fmt.Println("----------------------------------------------------------------------")
		fmt.Print("Masukan perintah (a/b/c....g/h/i) : ")
		fmt.Scan(&input)
		if input == "a" {
			add_soal(&*bank)
		} else if input == "b" {
			remove_soal(&*bank)
		} else if input == "c" {
			edit_soal(&*bank)
		} else if input == "d" {
			set_nquestion(&*bank)
		} else if input == "e" {
			print_soal_termudah(*bank)
		} else if input == "f" {
			print_rank_asc(*user)
		} else if input == "g" {
			print_rank_desc(*user)
		} else if input == "h" {
			search_user(*user)
		} else if input != "i" {
			fmt.Print("Input tidak dikenali\n")
		}
	}
}

func menu_user(bank *bankSoal, user *userTab, cur int) { //cur buat ngasi idx current user, jadi nanti kalo ngerjain kuiz tau, yang ngerjain siapa tergantung log in
	var input string
	fmt.Println("============================= MAIN MENU ==============================")
	fmt.Printf("Selamat datang %s, skor quiz Anda saat ini %0.f\n", user.data[cur].lName, user.data[cur].score)
	fmt.Println("Reward Anda Untuk Bulan ini Sebesar ", user.data[cur].reward)
	fmt.Println("a. Coba QUIZ!!!")
	fmt.Println("b. Lihat peringkat")
	fmt.Println("c. Log out")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Print("Masukan perintah (a/b/c) : ")
	fmt.Scan(&input)
	if input == "a" {
		random_soal(&*bank, &*user, cur)
	} else if input == "b" {
		print_rank_desc(*user)
	} else if input != "c" {
		fmt.Print("Input tidak dikenali\n")
	}
}

func add_soal(bank *bankSoal) {
	var temp formatSoal
	fmt.Println("============================== ADD MENU ==============================")
	fmt.Print("Masukan pertanyan (spasi dg underscore) : ")
	fmt.Scan(&temp.question)
	for i := 0; i < 4; i++ {
		fmt.Printf("Masukan opsi %d (spasi dg underscore) : ", i+1)
		fmt.Scan(&temp.option[i])
	}
	fmt.Print("Masukan kunci jawaban [1/2/3/4] : ")
	fmt.Scan(&temp.kj)
	bank.data[bank.n] = temp
	bank.n++
	fmt.Print("Soal berhasil ditambahkan ke bank soal . . . \n")
}

func remove_soal(bank *bankSoal) {
	var no int
	fmt.Println("============================ REMOVE MENU =============================")
	Print_soal(*bank)
	fmt.Print("Pilih soal yang ingin dihapus (ketik angka) : ")
	fmt.Scan(&no)
	no--
	if no == bank.n {
		var trash formatSoal
		bank.data[no] = trash
	} else {
		for i := no; i < bank.n; i++ {
			if i != bank.n-1 {
				bank.data[i] = bank.data[i+1]
			}
		}
	}
	bank.n--
}

func edit_soal(bank *bankSoal) {
	var no int
	var temp formatSoal
	fmt.Println("============================= EDIT MENU ==============================")
	Print_soal(*bank)
	fmt.Print("Pilih soal yang ingin diedit (ketik angka) : ")
	fmt.Scan(&no)
	no--
	fmt.Println("Masukan pertanyan baru (spasi dg underscore) : ")
	fmt.Scan(&temp.question)
	for i := 0; i < 4; i++ {
		fmt.Printf("Masukan opsi %d baru (spasi dg underscore) : ", i+1)
		fmt.Scan(&temp.option[i])
	}
	fmt.Print("Masukan kunci jawaban baru [1/2/3/4] : ")
	fmt.Scan(&temp.kj)
	bank.data[no] = temp
	fmt.Print("Soal berhasil diedit . . . \n")
}

func random_soal(bank *bankSoal, user *userTab, cur int) {
	random := make([]int, bank.nShow) //membuat arr untuk randomizer soal, panjang array tergantung dengan bank.nShow
	var answer int
	var count int
	var status bool
	var idx int = 0
	var cek bool
	for !status {
		cek = false
		random[idx] = rand.Intn(bank.n)
		for i := 0; i < idx && !cek; i++ {
			if random[i] == random[idx] {
				cek = true
			}
		}
		if !cek {
			idx++
		}
		if idx == bank.nShow {
			status = true
		}
	}
	for i := 0; i < bank.nShow; i++ {
		x := random[i]
		fmt.Printf("%d. %s\n", i+1, bank.data[x].question)
		for j := 0; j < 4; j++ {
			fmt.Printf("Opsi %d - %s\n", j+1, bank.data[x].option[j])
		}
		fmt.Print("Pilih jawaban Anda (ketik angka) : ")
		fmt.Scan(&answer)
		if answer == bank.data[x].kj {
			bank.data[x].correct++
			count++
		}
	}
	result := float64(count) / float64(bank.nShow) * 100
	user.data[cur].score = result
	if result >= 30 && result <= 35 {
		user.data[cur].reward = "2000"
	} else if result >= 36 && result <= 40 {
		user.data[cur].reward = "4000"
	} else if result >= 41 && result <= 45 {
		user.data[cur].reward = "4000"
	} else if result >= 46 && result <= 50 {
		user.data[cur].reward = "6000"
	} else if result >= 51 && result <= 55 {
		user.data[cur].reward = "15000"
	} else if result >= 56 && result <= 60 {
		user.data[cur].reward = "20000"
	} else if result >= 61 && result <= 75 {
		user.data[cur].reward = "25000"
	} else if result >= 76 && result <= 80 {
		user.data[cur].reward = "30000"
	} else if result >= 81 && result <= 85 {
		user.data[cur].reward = "35000"
	} else if result >= 86 && result <= 90 {
		user.data[cur].reward = "48000"
	} else if result >= 91 && result <= 96 {
		user.data[cur].reward = "55000"
	} else if result >= 96 && result <= 99 {
		user.data[cur].reward = "65000"
	} else if result == 100 {
		user.data[cur].reward = "100000"
	}
	fmt.Printf("Anda berhasil menyelesaikan quiz dengan skor %0.f, selamat!!!\n", result)
}

func set_nquestion(bank *bankSoal) {
	fmt.Print("Set jumlah soal quiz : ")
	fmt.Scan(&bank.nShow)
}

func print_rank_desc(user userTab) {
	for i := 1; i < user.n; i++ {
		j := i
		for j != 0 && user.data[j].score > user.data[j-1].score {
			user.data[j], user.data[j-1] = user.data[j-1], user.data[j]
			j--
		}
	}
	fmt.Print("No.\tScore\tNama\n")
	for i := 0; i < user.n; i++ {
		fmt.Printf("%d.\t%0.f\t%s %s\n", i+1, user.data[i].score, user.data[i].fName, user.data[i].lName)
	}
	fmt.Println("tekan enter untuk kembali\n")
	fmt.Scanln()
	fmt.Scanln()
}

func print_rank_asc(user userTab) {
	for i := 1; i < user.n; i++ {
		j := i
		for j != 0 && user.data[j-1].score > user.data[j].score {
			user.data[j], user.data[j-1] = user.data[j-1], user.data[j]
			j--
		}
	}
	fmt.Print("No.\tScore\tNama\n")
	for i := 0; i < user.n; i++ {
		fmt.Printf("%d.\t%0.f\t%s %s\n", i+1, user.data[i].score, user.data[i].fName, user.data[i].lName)
	}
	fmt.Println("tekan enter untuk kembali\n")
	fmt.Scanln()
	fmt.Scanln()
}

func print_soal_termudah(bank bankSoal) {
	var min int
	for i := 0; i < bank.n; i++ {
		min = i
		for j := i; j < bank.n; j++ {
			if bank.data[j].correct > bank.data[min].correct {
				min = j
			}
		}
		bank.data[i], bank.data[min] = bank.data[min], bank.data[i]
	}
	fmt.Print("No.\tCorrect\tJudul\n")
	for i := 0; i < bank.n; i++ {
		fmt.Printf("%d.\t%d\t%s\n", i+1, bank.data[i].correct, bank.data[i].question)
	}
	fmt.Println("tekan enter untuk kembali\n")
	fmt.Scanln()
	fmt.Scanln()
}

func search_user(user userTab) {
	var input int
	var m int
	var status bool = false
	fmt.Print("Masukan id kontestan : ")
	fmt.Scan(&input)
	for i := 1; i < user.n; i++ {
		j := i
		for j != 0 && user.data[j].id < user.data[j-1].id {
			user.data[j], user.data[j-1] = user.data[j-1], user.data[j]
			j--
		}
	}

	kiri := 0
	kanan := user.n - 1
	for kiri <= kanan && !status {
		m = (kanan + kiri) / 2
		if user.data[m].id > input {
			kiri = m + 1
		} else if user.data[m].id < input {
			kanan = m - 1
		} else {
			status = true
		}
	}
	if status {
		fmt.Println("============================= USER INFO ==============================")
		fmt.Printf("Nama		: %s %s\n", user.data[m].fName, user.data[m].lName)
		fmt.Printf("Username	: %s\n", user.data[m].username)
		fmt.Printf("id			: %d\n", user.data[m].id)
		fmt.Printf("Score		: %0.f\n", user.data[m].score)
		fmt.Println("----------------------------------------------------------------------")
		fmt.Print("tekan enter untuk kembali")
		fmt.Scanln()
		fmt.Scanln()
	} else {
		fmt.Print("tidak ada user dengan id tersebut\n")
	}
}

func Print_soal(bank bankSoal) {
	for i := 0; i < bank.n; i++ {
		fmt.Printf("%d. %s\n", i+1, bank.data[i].question)
	}
}

func menu_guest(bank *bankSoal, user *userTab) {
	var input string
	for input != "c" {
		fmt.Println("============================ LOG IN MENU =============================")
		fmt.Print("a. Log in\nb. Sign in\nc. Exit\n")
		fmt.Println("======================================================================")
		fmt.Print("Masukan perintah (a/b/c) : ")
		fmt.Scan(&input)
		if input == "b" {
			sign_in(&*user)
		} else if input == "a" {
			log_in(&*bank, &*user)
		} else {
			fmt.Print("error, input tidak dikenali\n")
		}
	}
}

func sign_in(user *userTab) {
	var temp formatUser
	var status bool = false
	for !status {
		fmt.Print("Masukan username  : ")
		fmt.Scan(&temp.username)
		if !cek_username(*user, temp.username) {
			fmt.Print("Username sudah digunakan, gunakan username lain\n")
		} else {
			status = true
		}
	}
	fmt.Print("Masukan nama depan : ")
	fmt.Scan(&temp.fName)
	fmt.Print("Masukan nama belakang : ")
	fmt.Scan(&temp.lName)
	fmt.Print("Masukan id : ")
	fmt.Scan(&temp.id)
	if !cek_id(*user, temp.id) {
		fmt.Print("Sudah ada user yang terdaftar dengan id tersebut\n")
		return
	}
	fmt.Print("Masukan password : ")
	fmt.Scan(&temp.pw)
	user.data[user.n] = temp
	user.n++
	fmt.Print("Sign in berhasil\n")
}

func log_in(bank *bankSoal, user *userTab) {
	var temp formatUser
	var status bool = false
	var idx int
	if bank.n == 0 {
		fmt.Println("Belum ada kontestan yang terdaftar")
		return
	}
	for !status {
		fmt.Print("Username : ")
		fmt.Scan(&temp.username)
		fmt.Print("Password : ")
		fmt.Scan(&temp.pw)
		if !cek_pw(*user, temp.username, temp.pw, &idx) {
			fmt.Print("Password salah, ulangi\n")
		} else {
			menu_user(&*bank, &*user, idx)
			fmt.Print("Log in berhasil, menuju ke menu utama\n")
			status = true
		}
	}
}

func cek_id(user userTab, cek int) bool {
	for i := 0; i < user.n; i++ {
		if user.data[i].id == cek {
			return false
		}
	}
	return true
}

func cek_username(user userTab, cek string) bool {
	for i := 0; i < user.n; i++ {
		if user.data[i].username == cek {
			return false
		}
	}
	return true
}

func cek_pw(user userTab, username, pw string, idx *int) bool {
	for i := 0; i < user.n; i++ {
		if user.data[i].username == username {
			if user.data[i].pw == pw {
				*idx = i
				return true
			}
		}
	}
	return false
}

func use_temp(bank *bankSoal, user *userTab) {
	bank.data[0] = formatSoal{
		question: "Dimana_ibu_kota_Indonesia?",
		option: [4]string{
			"Jogja",
			"Jakarta",
			"Bandung",
			"Surabaya",
		},
		kj: 2,
	}
	bank.data[1] = formatSoal{
		question: "Dimana_ibu_kota_Malaysia?",
		option: [4]string{
			"Kuala_Lumpur",
			"Jakarta",
			"Bandung",
			"Surabaya",
		},
		kj: 1,
	}
	bank.data[2] = formatSoal{
		question: "Dimana_ibu_kota_Jerman?",
		option: [4]string{
			"Franfurt",
			"Berlin",
			"Leipzig",
			"Munich",
		},
		kj: 2,
	}
	bank.data[3] = formatSoal{
		question: "Dimana_ibu_kota_Ukraina?",
		option: [4]string{
			"Luhansk",
			"Donetsk",
			"Kyiv",
			"Belgorod",
		},
		kj: 3,
	}
	bank.data[4] = formatSoal{
		question: "Dimana_ibu_kota_China?",
		option: [4]string{
			"Guangdong",
			"Hunan",
			"Shanghai",
			"Beijing",
		},
		kj: 4,
	}
	bank.data[5] = formatSoal{
		question: "Dimana_ibu_kota_Jepang?",
		option: [4]string{
			"Tokyo",
			"Kyoto",
			"Osaka",
			"Nagoya",
		},
		kj: 1,
	}
	bank.n = 6
	user.data[0] = formatUser{
		username: "jokowi",
		fName:    "Joko",
		lName:    "Widodo",
		pw:       "1234",
		score:    90,
	}
	user.data[1] = formatUser{
		username: "hendro",
		fName:    "Prabowo",
		lName:    "Subianto",
		pw:       "1234",
		score:    85,
	}
	user.data[2] = formatUser{
		username: "soekarno",
		fName:    "Soekarno",
		lName:    "",
		pw:       "1234",
		score:    95,
	}
	user.data[3] = formatUser{
		username: "Markonah",
		fName:    "Siti",
		lName:    "Markonah",
		pw:       "1234",
		score:    100,
	}
	user.data[4] = formatUser{
		username: "AsepSalep",
		fName:    "Asep",
		lName:    "Salep",
		pw:       "1234",
		score:    99,
	}
	user.data[5] = formatUser{
		username: "udinswalow",
		fName:    "udin",
		lName:    "swalow",
		pw:       "1234",
		score:    70,
	}
	user.n = 6
}
