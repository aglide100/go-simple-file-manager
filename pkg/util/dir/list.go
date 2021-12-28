package dir

import (
	"fmt"
	"io/ioutil"
)

func GetDirectoryList(dir string) ([]string,error) {
	if dir == "" {
		dir = "./"
	}

	files, err := ioutil.ReadDir(dir);
	if err != nil {
		return nil, err
	}
	var list []string;
	for _, f := range files {
		list = append(list, f.Name())
		
		fmt.Println(f.Name())
	}

	return list,nil
}
