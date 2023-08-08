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
	fmt.Println("-----")

//	var meh = graph_node.CreateNode(nil)
//
//	var a = meh.AddFindChild(12)
//	var b = a.AddFindChild(12)
//	var c = b.AddFindChild(13)
//	c.AddFindChild(10)
//
//	fmt.Println("==========")
//	x,d := meh.WalkTree([]uint8{12,12,13,10})
//	fmt.Println("====##====")
//
//	fmt.Println(x,d)

	var meh = word_graph.CreateWordGraph()
	meh.AddWord(([]uint8{12,12,13,10}))

	fmt.Println(meh.FindWord([]uint8{12,12,13,10}))
	fmt.Println(meh.FindWord([]uint8{12,12,01,10}))
}
