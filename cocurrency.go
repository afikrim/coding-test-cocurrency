package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	carWilBaseUrl := "http://jendela.data.kemdikbud.go.id/api/index.php/CWilayah/wilayahGET"
	// carMusBaseUrl := "http://jendela.data.kemdikbud.go.id/api/index.php/CcariMuseum/searchGET"

	// var prov ProvRes

	// List propinsi
	// testJson := `{"data":[{"kode_wilayah":"010000  ","nama":"Prov. D.K.I. Jakarta","mst_kode_wilayah":"000000  "},{"kode_wilayah":"020000  ","nama":"Prov. Jawa Barat","mst_kode_wilayah":"000000  "},{"kode_wilayah":"030000  ","nama":"Prov. Jawa Tengah","mst_kode_wilayah":"000000  "},{"kode_wilayah":"040000  ","nama":"Prov. D.I. Yogyakarta","mst_kode_wilayah":"000000  "},{"kode_wilayah":"050000  ","nama":"Prov. Jawa Timur","mst_kode_wilayah":"000000  "},{"kode_wilayah":"060000  ","nama":"Prov. Aceh","mst_kode_wilayah":"000000  "},{"kode_wilayah":"070000  ","nama":"Prov. Sumatera Utara","mst_kode_wilayah":"000000  "},{"kode_wilayah":"080000  ","nama":"Prov. Sumatera Barat","mst_kode_wilayah":"000000  "},{"kode_wilayah":"090000  ","nama":"Prov. Riau","mst_kode_wilayah":"000000  "},{"kode_wilayah":"100000  ","nama":"Prov. Jambi","mst_kode_wilayah":"000000  "},{"kode_wilayah":"110000  ","nama":"Prov. Sumatera Selatan","mst_kode_wilayah":"000000  "},{"kode_wilayah":"120000  ","nama":"Prov. Lampung","mst_kode_wilayah":"000000  "},{"kode_wilayah":"130000  ","nama":"Prov. Kalimantan Barat","mst_kode_wilayah":"000000  "},{"kode_wilayah":"140000  ","nama":"Prov. Kalimantan Tengah","mst_kode_wilayah":"000000  "},{"kode_wilayah":"150000  ","nama":"Prov. Kalimantan Selatan","mst_kode_wilayah":"000000  "},{"kode_wilayah":"160000  ","nama":"Prov. Kalimantan Timur","mst_kode_wilayah":"000000  "},{"kode_wilayah":"170000  ","nama":"Prov. Sulawesi Utara","mst_kode_wilayah":"000000  "},{"kode_wilayah":"180000  ","nama":"Prov. Sulawesi Tengah","mst_kode_wilayah":"000000  "},{"kode_wilayah":"190000  ","nama":"Prov. Sulawesi Selatan","mst_kode_wilayah":"000000  "},{"kode_wilayah":"200000  ","nama":"Prov. Sulawesi Tenggara","mst_kode_wilayah":"000000  "},{"kode_wilayah":"210000  ","nama":"Prov. Maluku","mst_kode_wilayah":"000000  "},{"kode_wilayah":"220000  ","nama":"Prov. Bali","mst_kode_wilayah":"000000  "},{"kode_wilayah":"230000  ","nama":"Prov. Nusa Tenggara Barat","mst_kode_wilayah":"000000  "},{"kode_wilayah":"240000  ","nama":"Prov. Nusa Tenggara Timur","mst_kode_wilayah":"000000  "},{"kode_wilayah":"250000  ","nama":"Prov. Papua","mst_kode_wilayah":"000000  "},{"kode_wilayah":"260000  ","nama":"Prov. Bengkulu","mst_kode_wilayah":"000000  "},{"kode_wilayah":"270000  ","nama":"Prov. Maluku Utara","mst_kode_wilayah":"000000  "},{"kode_wilayah":"280000  ","nama":"Prov. Banten","mst_kode_wilayah":"000000  "},{"kode_wilayah":"290000  ","nama":"Prov. Kepulauan Bangka Belitung","mst_kode_wilayah":"000000  "},{"kode_wilayah":"300000  ","nama":"Prov. Gorontalo","mst_kode_wilayah":"000000  "},{"kode_wilayah":"310000  ","nama":"Prov. Kepulauan Riau","mst_kode_wilayah":"000000  "},{"kode_wilayah":"320000  ","nama":"Prov. Papua Barat","mst_kode_wilayah":"000000  "},{"kode_wilayah":"330000  ","nama":"Prov. Sulawesi Barat","mst_kode_wilayah":"000000  "},{"kode_wilayah":"340000  ","nama":"Prov. Kalimantan Utara","mst_kode_wilayah":"000000  "},{"kode_wilayah":"350000  ","nama":"Luar Negeri","mst_kode_wilayah":"000000  "}]}`
	resPropinsi := getData(carWilBaseUrl)
	// json.Unmarshal(resPropinsi, &prov)
	fmt.Print(resPropinsi)
}

func getData(url string) string {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	return string(data)
}

//-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------/
// Model section

type ProvRes struct {
	Data []Prov
}

type Prov struct {
	Kode_wilayah     string
	Nama             string
	Mst_kode_wilayah string
}

type KabKota struct {
	Kode_wilayah string
}

type Museum struct {
	Museum_id          string
	Kode_pengelolaan   string
	Nama               string
	Sdm                string
	Alamat_jalan       string
	Desa_kelurahan     string
	Kecamatan          string
	Kabupaten_kota     string
	Propinsi           string
	Lintang            string
	Bujur              string
	Koleksi            string
	Sumber_dana        string
	Pengelola          string
	Tipe               string
	Standar            string
	Tahun_berdiri      string
	Bangunan           string
	Luas_tanah         string
	Status_kepemilikan string
}

//-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------/
// Worker section

func worker(id int, jobs <-chan int, results chan<- int, url string) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		getData(url)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func job(workerNum int, totalJob int, url string) {
	jobs := make(chan int, totalJob)
	results := make(chan int, totalJob)

	for i := 0; i < workerNum; i++ {
		go worker(i, jobs, results, url)
	}

	for i := 0; i < totalJob; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < totalJob; i++ {
		<-results
	}
}
