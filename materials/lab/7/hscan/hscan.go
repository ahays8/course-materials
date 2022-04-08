package hscan

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	//"time"
)

//==========================================================================\\

var shalookup map[string]string
var md5lookup map[string]string

func GuessSingle(sourceHash string, filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var password = ""
	for scanner.Scan() {
		password = scanner.Text()

		// TODONE - From the length of the hash you should know which one of these to check ...
		// add a check and logicial structure
		if len(sourceHash) == 32 {
			hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (MD5): %s\n", password)
				return password
			}
		} else if len(sourceHash) == 64 {
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (SHA-256): %s\n", password)
				return password
			}
		} else {
			log.Fatalln("Hash is incorrect size for md5 or sha256")
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return password
}

func GenHashMaps(filename string) (int,int){//returns sizes of the maps for checking
	//TODONE
	//itterate through a file (look in the guessSingle function above)
	//rather than check for equality add each hash:passwd entry to a map SHA and MD5 where the key = hash and the value = password
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	var MD5 sync.Map
	var MD5count=0
	var SHA sync.Map
	var SHAcount=0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		password := scanner.Text()

		//w/subroutines: 29s, 99 microseconds per password
		
		go md5Map(password, MD5)
		MD5count++
		go shaMap(password, SHA)
		SHAcount++
		/*
		//w/o subroutines: 13s, 43 microseconds per password
		//I do not know why this is so much faster
		//I may try to refactor this on my own time
		hashmd5 := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		hashsha := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
		MD5.Store(hashmd5, password)
		MD5count++
		SHA.Store(hashsha, password)
		SHAcount++*/
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return MD5count, SHAcount
	//TODONE at the very least use go subroutines to generate the sha and md5 hashes at the same time
	//OPTIONAL -- Can you use workers to make this even faster

	//TODO create a test in hscan_test.go so that you can time the performance of your implementation
	//Test and record the time it takes to scan to generate these Maps
	// 1. With and without using go subroutines
	// 2. Compute the time per password (hint the number of passwords for each file is listed on the site...)
}
func md5Map(password string, MD5 sync.Map) {
	var hashmd5 = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	MD5.Store(hashmd5, password)
}
func shaMap(password string, SHA sync.Map) {
	hashsha := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	SHA.Store(hashsha, password)
}
func GetSHA(hash string) (string, error) {
	password, ok := shalookup[hash]
	if ok {
		return password, nil

	} else {

		return "", errors.New("password does not exist")

	}

}

//TODONE
func GetMD5(hash string) (string, error) {
	password, ok := md5lookup[hash]
	if ok {
		return password, nil

	} else {

		return "", errors.New("password does not exist")

	}
}
