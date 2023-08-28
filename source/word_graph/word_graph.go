/**--------------------------------------------------------------------------------*
 *       ____                             ____                    _      
 *      / __ \                           |  _ \                  (_)     
 *     | |  | |_ __ __ _ _ __   __ _  ___| |_) |_   _  __ _  __ _ _  ___ 
 *     | |  | | '__/ _` | '_ \ / _` |/ _ \  _ <| | | |/ _` |/ _` | |/ _ \
 *     | |__| | | | (_| | | | | (_| |  __/ |_) | |_| | (_| | (_| | |  __/
 *      \____/|_|  \__,_|_| |_|\__, |\___|____/ \__,_|\__, |\__, |_|\___|
 *                              __/ |                  __/ | __/ |       
 *                             |___/                  |___/ |___/        
 * Name  : word_graph.go
 * Desc  : This class is the graph of words.
 *
 * Author: Peter Antoine
 * Date  : 07/08/2023
 *--------------------------------------------------------------------------------*
 *                     Copyright (c) 2023 Peter Antoine
 *                            All rights Reserved.
 *                      Released Under the MIT Licence
 *--------------------------------------------------------------------------------*/

package word_graph

import "orange_buggie/source/graph_node"

// The main class
//
// This holds the graph of the words that have been added and the list of
// the last nodes in the tree of the word end nodes. This can be used to
// tie the end of the words together.
type WordGraph struct {
	ends	map [uint8] graph_list		// The list that has the ends of the words - this is for tail connection.
	tree	*graph_node.GraphNode		// The node of links that exist in the tree.
}

// utility type to make the code more readable.
type graph_list = []*graph_node.GraphNode

func CreateWordGraph () *WordGraph {
	var result = new(WordGraph)

	result.ends = make(map [uint8] graph_list)
	result.tree = graph_node.CreateNode(nil)

	return result
}

func (w *WordGraph) AddWord(word []uint8) *graph_node.GraphNode {
	var result = w.tree

	// add the bytes from the word to the tree.
	for _, value := range(word) {
		result = result.AddFindChild(value)
	}

	// Ok, this is a terminating node.
	result.SetTerminates()

	// If the entry does not exist, will need to add a list there
	if w.ends[word[len(word)-1]] == nil {
		w.ends[word[len(word)-1]] = make(graph_list, 0)
	}

	// Add the new node to the end of the array
	w.ends[word[len(word)-1]] = append(w.ends[word[len(word)-1]], result)

	return result
}

func (w WordGraph) FindWord(word []uint8, parts *uint16) bool {
	var current = w.tree

	for _, value := range(word) {
		current = current.FindChild(value)

		if current == nil {
			break;
		}
	}

	// if current is not nil and the node is a terminating node, then 
	// we found the word.
	if current != nil && current.IsTerminal() {
		*parts = current.GetParts()
	}

	return current != nil && current.IsTerminal()
}
