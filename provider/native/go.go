package native

// go wasm 中不能调用的:
// import "fmt"

func Fibonacci(in int32) int32 {
	if in <= 1 {
		return in
	}
	return Fibonacci(in-1) + Fibonacci(in-2)
}

func RequestHTTP() int32 {
	return 0
	// fmt.Println("RequestHTTP")
}

/*
func RequestHTTP() {
	httpTestURL := `https://httpbin.org/basic-auth`
	username := "xxx"
	password := "yyy"

	reqURL := fmt.Sprintf("%s/%s/%s", httpTestURL, username, password)

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(username, password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}

	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}
*/

func FileIO() int32 {
	return 0
	// fmt.Println("FileIO")
}

/*
func FileIO() error {
	dir, err := ioutil.TempDir("", "test-*")
	if err != nil {
		return err
	}
	defer os.Remove(dir)

	fmt.Println(dir)

	f, err := ioutil.TempFile(dir, "tmp-*")
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	fmt.Println(f.Name())

	writeN, err := f.WriteAt([]byte("test-content"), 0)
	if err != nil {
		return err
	}

	readDest := make([]byte, writeN)
	readN, err := f.ReadAt(readDest, 0)
	if err != nil {
		return err
	}
	if writeN != readN {
		return fmt.Errorf("read length is %d, written length is %d", readN, writeN)
	}

	return nil
}
*/

func MultiThreads(num int32) int32 {
	return 0
	// fmt.Println("MultiThreads")
}

/*
func MultiThreads(num int32) {
	g := new(errgroup.Group)

	for i := int32(0); i < num; i++ {
		g.Go(func() error {
			Fibonacci(30)
			return nil
		})
	}

	g.Wait()
}
*/
