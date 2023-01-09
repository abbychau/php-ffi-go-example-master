package main

import (
	"C"
	"fmt"
	"io/ioutil"
	"net/http"
)
import (
	"reflect"
	"sync"
	"unsafe"
)

//export print
func print(out *C.char) {
	fmt.Println("[GO print] " + C.GoString(out))
}

//export sum
func sum(a C.int, b C.int) C.int {
	return a + b
}

//export httpGet
func httpGet(url string) *C.char {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return C.CString(string(body))
}

func httpGetGoString(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)

}

func get(url string, c chan string) {
	res := httpGetGoString(url)
	c <- res
}

//export goroutineRun
func goroutineRun(url1 string, url2 string) **C.char {

	c1 := make(chan string)
	c2 := make(chan string)

	go get(url1, c1)
	go get(url2, c2)

	//get result
	res1 := <-c1
	res2 := <-c2

	cArray := C.malloc(C.size_t(2) * C.size_t(unsafe.Sizeof(uintptr(0))))

	//set result to cArray
	*(*uintptr)(unsafe.Pointer(uintptr(cArray) + 0*unsafe.Sizeof(uintptr(0)))) = uintptr(unsafe.Pointer(C.CString(res1)))
	*(*uintptr)(unsafe.Pointer(uintptr(cArray) + 1*unsafe.Sizeof(uintptr(0)))) = uintptr(unsafe.Pointer(C.CString(res2)))

	return (**C.char)(cArray)

}

//export urlMultiGet
func urlMultiGet(urls []string) **C.char {

	// get all the urls response and put to the return with the same order
	var wg sync.WaitGroup
	wg.Add(len(urls))

	cArray := C.malloc(C.size_t(len(urls)) * C.size_t(unsafe.Sizeof(uintptr(0))))

	for i, url := range urls {
		go func(i int, url string) {
			defer wg.Done()
			*(*uintptr)(unsafe.Pointer(uintptr(cArray) + uintptr(i)*unsafe.Sizeof(uintptr(0)))) = uintptr(unsafe.Pointer(C.CString(httpGetGoString(url))))
		}(i, url)
	}

	wg.Wait()

	return (**C.char)(cArray)

}

func main() {
	// url1 := "http://httpbin.org/get"
	// url2 := "http://httpbin.org/headers"

	// cArray := goroutineRun(url1, url2)

	var slice []*C.char
	// sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&slice)))
	// sliceHeader.Cap = 2
	// sliceHeader.Len = 2
	// sliceHeader.Data = uintptr(unsafe.Pointer(cArray))

	// fmt.Println(C.GoString(slice[0]))
	// fmt.Println(C.GoString(slice[1]))

	//test urlMultiGet
	urls := []string{"http://httpbin.org/get", "http://httpbin.org/headers"}
	cArray := urlMultiGet(urls)

	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&slice)))
	sliceHeader.Cap = len(urls)
	sliceHeader.Len = len(urls)
	sliceHeader.Data = uintptr(unsafe.Pointer(cArray))

	for i := 0; i < len(urls); i++ {
		fmt.Printf("urls: %v, i: %v", urls[i], i)
		fmt.Println(C.GoString(slice[i]))
	}
}
