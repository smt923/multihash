package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Error! Not enough arguments given.\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [files to generate checksums]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s firstfile.txt folder/anotherfile.txt /etc/thirdfile.txt\n", os.Args[0])
	}

	filelist := os.Args[1:]

	if err := hashAndPrint(filelist); err != nil {
		fmt.Fprintf(os.Stderr, "Could not hash files! Error:\n%s", err)
		os.Exit(1)
	}

}

func hashAndPrint(filelist []string) error {
	hsha256 := sha256.New()
	hsha1 := sha1.New()
	hmd5 := md5.New()

	for _, file := range filelist {
		// We print a newline at the very beginning to make sure the output is clearly seperated and very readable
		fmt.Print("\n")
		f, err := os.Open(file)
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Could not open file for reading: %s\n", file)
			continue
		} else if err != nil {
			return err
		}
		defer f.Close()

		data := make([]byte, 4096)

		for {
			data = data[:cap(data)]
			n, err := f.Read(data)
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}
			data = data[:n]
			hsha256.Write(data)
			hsha1.Write(data)
			hmd5.Write(data)
		}

		fmt.Println(file)
		fmt.Printf("SHA256: %x\n", hsha256.Sum(nil))
		fmt.Printf("SHA1: %x\n", hsha1.Sum(nil))
		fmt.Printf("MD5: %x\n", hmd5.Sum(nil))

		// Reset our hashes to initial value, this allows us to reuse the same ones we created above
		hsha256.Reset()
		hsha1.Reset()
		hmd5.Reset()
	}

	return nil
}
