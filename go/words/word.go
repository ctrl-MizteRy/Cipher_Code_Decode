package words

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func createFile(path string) {
	resp, err := http.Get("https://www.mit.edu/~ecprice/wordlist.10000")
	if err != nil {
		panic(fmt.Sprintf("Error: could not access to the webite, error type: %v", err))
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("Error, could not read the body of the website, err: %v ", err))

	}

	err1 := os.WriteFile(path, body, 0644)

	if err1 != nil {
		panic(fmt.Sprintf("Something wrong, could not write into the file, err: %v", err1))
	}
}

func ReadWords() map[string]byte {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintf("Could not get home dir, err: %v", err))
	}

	path := filepath.Join(home, "Desktop/Projects/HTML/Cipher/go/words/dictionary.txt")

	var file *os.File
	if a, err := os.Open(path); err != nil {
		createFile(path)
	} else {
		file = a
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	dict := make(map[string]byte)
	for scan.Scan() {
		dict[scan.Text()] = 1
	}
	return dict
}
