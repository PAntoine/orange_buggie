/**--------------------------------------------------------------------------------*
 * Name  : build_graph.go
 * Desc  : This application will build the graph.
 *
 * Author: Peter Antoine
 * Date  : 24/07/2023
 *--------------------------------------------------------------------------------*
 *                     Copyright (c) 2023 Peter Antoine
 *                            All rights Reserved.
 *                      Released Under the MIT Licence
 *--------------------------------------------------------------------------------*/

package main
import "fmt"
import "orange_buggie/source/word_graph"

func main() {
	var meh = word_graph.CreateWordGraph()

	meh.LoadGraph("test_data/test_dictionary.json")

	var parts uint16
	fmt.Println(meh.FindWord([]byte("dog"), &parts), parts)
	fmt.Println(meh.FindWord([]byte("cat"), &parts), parts)
}
