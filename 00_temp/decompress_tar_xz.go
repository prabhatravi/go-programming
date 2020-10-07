package main

import (
	"archive/tar"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"github.com/xi2/xz"
	"path/filepath"
)

//UntarTarXZ will extract xz compressed files to xpath directory
func untarTarXZ(tarFileName, xpath string) (err error) {
	f, err := os.Open(tarFileName)
	if err != nil {
		log.Println("Could not open tarFileName: ", tarFileName)
		return err
	}
	defer f.Close()
	absPath, err := filepath.Abs(xpath)
	if err != nil {
		log.Println("Could not find Absolute path for: ", xpath)
		return err
	}
	r, err := xz.NewReader(f, 0)
	if err != nil {
		log.Println("XZ New reader failed for: ", tarFileName, "Error: ", err)
		return err
	}
	// Create a tar Reader
	tr := tar.NewReader(r)
	// Iterate through the files in the archive.
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}
		if err != nil {
			log.Println("Error in getting the next tar. Error: ", err)
			return err
		}
		finfo := hdr.FileInfo()
		fileName := hdr.Name
		absFileName := filepath.Join(absPath, fileName)
		dir := filepath.Dir(absFileName)

		if _, err := os.Stat(dir); err != nil {
			if err := os.MkdirAll(dir, 0755); err != nil {
				log.Println("Cannot create Dir: ", dir)
				return err
			}
		}
		switch hdr.Typeflag {
		case tar.TypeDir:
			// create a directory
			err = os.MkdirAll(hdr.Name, 0777)
			if err != nil {
				log.Println("Could not create dir: ", hdr.Name, "Error: ", err)
				return err
			}
		case tar.TypeReg, tar.TypeRegA:
			// write a file
			w, err := os.OpenFile(absFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, finfo.Mode().Perm())
			if err != nil {
				log.Println("Error in creating file: ", absFileName, "Error: ", err)
				return err
			}
			_, err = io.Copy(w, tr)
			if err != nil {
				log.Println("Error in writing file: ", absFileName, "Error: ", err)
				return err
			}
			w.Close()
		}
	}

	return nil
}

func main() {
	flag.Parse()

	sourceFile := flag.Arg(0)

	if sourceFile == "" {
		fmt.Println("Dude, you didn't pass in a tar file!")
		os.Exit(1)
	}

	fmt.Println("arg 1: ", flag.Arg(0))

	untarTarXZ(flag.Arg(0), "/tmp")
}
