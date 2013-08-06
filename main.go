/**
 * Created by Cosmin on 06.08.2013.
 */
package main

import "fmt"
import "io/ioutil"
import "strings"

func printDir(dirPath string, recursive bool, level int) {
	finfo, err := ioutil.ReadDir(dirPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	for findex := 0; findex < len(finfo); findex ++ {
		displayName := finfo[findex].Name();
		if (finfo[findex].IsDir()) {
			displayName = "[" + displayName + "]"
		}
		fmt.Println(strings.Repeat(" ", level), "- " , displayName)
		if (recursive && finfo[findex].IsDir()) {
			printDir(dirPath + "/" + finfo[findex].Name(), true, level + 1)
		}
	}


}

func main() {
	printDir(".", true, 0)
}
