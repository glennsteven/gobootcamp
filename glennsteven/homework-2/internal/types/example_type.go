package types

import "fmt"

func PointerTypes(x, y int) {
	p := &x // point to var x
	fmt.Println("read x through the pointer:", *p)

	*p = 21 // set x through the pointer
	fmt.Println("see the new value of x:", *p)

	p = &y       // point to y
	*p = *p / 37 // divide y through the pointer
	fmt.Println("see the new value of y:", y)
}

// Wallet struct menyimpan informasi saldo
type Wallet struct {
	Balance int
}

// AddFail menggunakan value receiver
// Artinya method ini menerima salinan Wallet
func (w Wallet) AddFail(amount int) {
	// w di sini adalah COPY dari wallet asli
	// perubahan hanya terjadi pada salinan ini
	w.Balance += amount
	// setelah proses selesai, w akan di buang(trashed)/tidak di simpan
}

// Add menggunakan pointer receiver
// Artinya method ini menerima alamat memori Wallet
func (w *Wallet) Add(amount int) {
	// w di sini menyimpan alamat memori wallet
	// Mengubah w.Balance berarti mengubah Wallet asli (data asli)
	w.Balance += amount
}

// TotalWallet digunakan untuk mendemonstrasikan perbedaan
// value receiver vs pointer receiver
func TotalWallet(val Wallet) {
	// Memory:
	// main wallet ──▶ {Balance: 100}
	// val (copy)   ──▶ {Balance: 100}

	// AddFail hanya mengubah salinan di dalam method
	// val tetap tidak berubah
	val.AddFail(50)
	fmt.Println("Unuse Pointer:", val.Balance) // 100

	// Memanggil method dengan POINTER receiver
	// Go otomatis mengubah ini menjadi: (&val).Add(50)
	val.Add(50)
	fmt.Println("Used Pointer:", val.Balance) // 150
}

func BasicArray() {
	var n [3]int
	n[0] = 1
	n[1] = 2
	n[2] = 3
	fmt.Println("Basic Array:", n)
	ChangeArray(&n)
	fmt.Println("Change Array:", n)
}

func ChangeArray(arr *[3]int) {
	arr[1] = 100
}

func SliceArray(nums []int) {
	s := nums[1:4]
	fmt.Println("SliceArray:", s)

	s = nums[:]
	s[0] = 100
	fmt.Println("SliceArray1:", s)
}

func SliceUsingMake() {
	fmt.Println("SliceUsingMake dengan len 3 dan cap 3")
	s := make([]int, 3)
	fmt.Println(s)      // [0 0 0]
	fmt.Println(len(s)) // 3
	fmt.Println(cap(s)) // 3

	fmt.Println("SliceUsingMake dengan len 2 dan cap 5")
	i := make([]int, 2, 5)

	fmt.Println(i)      // [0 0]
	fmt.Println(len(i)) // 2
	fmt.Println(cap(i)) // 5

	fmt.Println("SliceUsingMake dengan len 2 dan cap 4 dan cara meng-append value ke dalam slice nya")
	a := make([]int, 2, 4)

	a[0] = 10
	a[1] = 20

	a = append(a, 30)
	a = append(a, 40)

	fmt.Println("GO akan me re-allocate memory, karena capacitynya sudah melebihi dari yang di tentukan.")
	fmt.Println("Capacity baru biasanya 2x dari sebelumnya")
	// Analogi:
	// Cap -> Kursi di meja
	// Append -> Nambah orang
	// Artinya: Jika Kursi di meja sudah limit, maka akan di tambah dengan kursi baru sesuai jumlah orangnya.
	a = append(a, 50)

	fmt.Println(a)      // [10 20 30 40]
	fmt.Println(len(a)) // 4
	fmt.Println(cap(a)) // 4
}

type UserMap struct {
	Name string
	Age  int
}

func ExampleMaps() {
	// 1. Membuat Map Value
	users := make(map[string]UserMap)

	// 2. Menambahkan data value ke map
	users["d1"] = UserMap{
		Name: "John Doe",
		Age:  20,
	}

	users["d2"] = UserMap{
		Name: "Jane Saliman",
		Age:  21,
	}

	fmt.Println("Data Awal:", users)

	// 3. Cara mengambil data map
	user, ok := users["d1"]
	if ok {
		fmt.Println("Data1 ditemukan:", user)
	} else {
		fmt.Println("Data1 tidak ditemukan")
	}

	fmt.Println("User u3:", users["u3"]) // zero value UserMap{}

	// 4. Cara update data di map
	tmp := users["d1"] // ambil value (COPY)
	tmp.Age = 25       // ubah copy
	users["d1"] = tmp  // simpan kembali ke map

	fmt.Println("Setelah update data1(d1):", users["d1"])

	// 5. Iterasi value map
	for key, value := range users {
		fmt.Println("Key:", key, "Value:", value)
	}

	// 6. Delete data map
	delete(users, "d2")
	fmt.Println("Setelah delete d2:", users)

	// 7. Cara membuat parameter map
	increaseAge(users, "d1")
	fmt.Println("Setelah function:", users["d1"])

	// 8. Map by reference type
	copyUsers := users
	copyUsers["b1"] = UserMap{Name: "Glenn Steven", Age: 27}

	// users ikut berubah karena menunjuk map yang sama
	fmt.Println("MapUsers:", users["b1"])
	fmt.Println("CopyMapUsers:", copyUsers["b1"])

	// 9. Map Sebagai SET
	set := make(map[string]struct{})

	set["kodok"] = struct{}{}
	set["buaya"] = struct{}{}

	if _, exists := set["kodok"]; exists {
		fmt.Println("kodok ada di set")
	}

	if _, exists := set["kadal"]; !exists {
		fmt.Println("kadal tidak ada di set")
	}
}

// increaseAge akan menerima map sebagai parameter
// Map sifatnya reference type, jadi perubahan akan tersimpan
func increaseAge(users map[string]UserMap, key string) {
	user, ok := users[key]
	if !ok {
		return
	}

	user.Age += 1
	users[key] = user
}
