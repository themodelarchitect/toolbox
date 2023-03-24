package toolbox

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

func TempDir() string {
	uniqName := uuid.New().String()
	tmpDirLocation := os.TempDir()
	tempFolder := filepath.Join(tmpDirLocation, uniqName)
	_, err := os.Stat(tempFolder)
	if os.IsNotExist(err) {
		return "/tmp"
	} else {
		_, err = os.Stat(tempFolder)
		if os.IsNotExist(err) {
			if err := os.Mkdir(tempFolder, os.ModePerm); err != nil {
				panic(err)
			}
		}
		return tempFolder
	}
}

func MakeTempFile(tempFileDir, prefix string) (f string, err error) {
	file, err := ioutil.TempFile(tempFileDir, prefix)
	if err != nil {
		return "", err
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
		}
	}(file.Name())
	return file.Name(), nil
}

func DeleteFile(filepath string) error {
	return os.Remove(filepath)
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func DownloadFile(filepath string, url string) (err error) {
	client := http.DefaultClient

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			panic(err)
		}
	}(out)

	// Get the data
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "QA_Automation/1.0")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s => Non 200 Response Code: %s ", url, resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func CopyFile(sourcePath, destinationPath string) error {
	srcFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer func(srcFile *os.File) {
		err := srcFile.Close()
		if err != nil {
		}
	}(srcFile)

	destFile, err := os.Create(destinationPath) // creates if file doesn't exist
	if err != nil {
		return err
	}
	defer func(destFile *os.File) {
		err := destFile.Close()
		if err != nil {
		}
	}(destFile)

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	if err != nil {
		return err
	}

	err = destFile.Sync()
	if err != nil {
		return err
	}
	return nil
}

func ReadFileToBytes(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

	// Get the file size
	stat, err := file.Stat()
	if err != nil {
		return []byte{}, err
	}

	// Read the file into a byte slice
	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		return []byte{}, err
	}
	return bs, nil
}

func ReadFileToString(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	return text, nil
}

func CombineFiles(files []string, filename string) error {
	var buf bytes.Buffer
	for _, file := range files {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		buf.Write(b)
	}

	err := ioutil.WriteFile(filename, buf.Bytes(), 0644)
	return err
}

func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func FindMatch(filePath, pattern string) ([]string, error) {
	var data []string
	file, err := os.Open(filePath)
	if err != nil {
		return data, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile(pattern)

	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			match := scanner.Text()
			data = append(data, match)
		}
	}

	if err := scanner.Err(); err != nil {
		return data, err
	}

	return data, nil
}

func HashFile(filePath string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string

	//Open the passed argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}

	//Tell the program to call the following function when the current function returns
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	//Open a new hash interface to write to
	hash := md5.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]

	//Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil

}
