package main

import (
	"haologs.com/gondb/db"
	"fmt"
)
func main()  {

	for {
		result,c,err := db.Client.Scan(0,"::*::",0).Result()
		if err != nil {
			fmt.Printf("%v\n",err)
		} else {
			fmt.Printf("%v\n",result)
		}
		if c == 0 {
			break
		}
	}


}
