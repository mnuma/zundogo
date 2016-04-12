package main

import (
	"fmt"
	"math/rand"
	"time"
	"reflect"
)

func main() {
	z := "ズン1"
	d := "ドコ"
	zd := []string{z, d}
	zundoko := []string{z, z, z, z, d}
	fmt.Println(zundoko)

	rand.Seed(time.Now().UnixNano())
	var zundokoArry []string

	go func() {
		for {
			randZundoko := zd[rand.Intn(2)]
			zundokoArry = append(zundokoArry, randZundoko)
			if len(zundokoArry) >= 5 {

				fmt.Println(zundokoArry)
				if reflect.DeepEqual(zundoko, zundokoArry) {
					fmt.Println(zundokoArry)
					break
				}

				zundokoArry = []string{}
				continue
			}
		}
	}()
	fmt.Println("キ・ヨ・シ!")
}

