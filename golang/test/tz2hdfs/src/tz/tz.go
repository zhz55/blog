package tz

import (
    "fmt"
    "os"
    "io"
    "archive/tar"
    "compress/gzip"
    //"github.com/Unknown/cae"
)

func main() {
	fr, err := os.Open("320000.tar.gz");
	if err != nil {
		panic(err)
	}
	defer fr.Close()

	gr, err := gzip.NewReader(fr)
	if err != nil {
		panic(err)
	}
	defer gr.Close()

	tr := tar.NewReader(gr)

	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Println(h.Name)

		fw, err := os.OpenFile(""+h.Name, os.O_CREATE | os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer fw.Close()

		_, err = io.Copy(fw, tr)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("un tar.gz ok")

}
