package common

import (
	"time"
	"strings"
	"errors"
	"os"
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type Credentials struct {
	AccessKeyID     string
	SecretAccessKey string
	SecurityToken   string
	Expiration      time.Time
}

/**
	credentials.csv file is download from amazon console
	when you create an account.
 */
func NewCredentialsFromCSV(filename string) (*Credentials, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	info := make(map[int]string)

	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		sps := bytes.Split(line, []byte(","))

		for _, spb := range (sps) {
			str := string(spb)
			str = strings.TrimSpace(str)
			if DEBUG_VERBOSE!=0 {
				fmt.Printf("%d. %s\n", len(info), str)
			}
			info[len(info)] = str
		}

		//strings.Split()
	}

	if len(info) != 6 || info[0] != "User Name" || info[1] != "Access Key Id" || info[2] != "Secret Access Key" {
		return nil, errors.New("The format of credentials csv file is not correct.")
	}

	cred := new(Credentials)
	cred.AccessKeyID = info[4]
	cred.SecretAccessKey = info[5]
	cred.SecurityToken = ""
	//cred.Expiration = time.Now()

	return cred, nil
}

func (cred *Credentials) Expired() bool {
	if cred.Expiration.IsZero() {
		// Credentials with no expiration can't expire
		return false
	}
	expireTime := cred.Expiration.Add(-4 * time.Minute)
	// if t - 4 mins is before now, true
	if expireTime.Before(time.Now()) {
		return true
	} else {
		return false
	}
}





