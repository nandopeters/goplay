package configfile

import (
	"strings"
	"bufio"
	"os"
)


func GetHostPort ( cfgFile string) (host string, port string, errOut error )  {
	file, err := os.Open(cfgFile) // For read access.
	if err != nil {
		return "","", err
	}
	
	r := bufio.NewReader(file)
	line, _, err := r.ReadLine()
	aa:= strings.Split(string(line[:]), "=")
	for i := 1; err == nil ; i++ {
		aa = strings.Split(string(line[:]), "=")
		switch {
		case aa[0] == "PORT":
			port=aa[1]
		case aa[0] == "HOST" :
			host = aa[1]		
		}

		line, _, err = r.ReadLine()
		}
	file.Close();

	return host, port, nil
}

