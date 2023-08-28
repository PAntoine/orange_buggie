/**--------------------------------------------------------------------------------*
 *       ____                             ____                    _      
 *      / __ \                           |  _ \                  (_)     
 *     | |  | |_ __ __ _ _ __   __ _  ___| |_) |_   _  __ _  __ _ _  ___ 
 *     | |  | | '__/ _` | '_ \ / _` |/ _ \  _ <| | | |/ _` |/ _` | |/ _ \
 *     | |__| | | | (_| | | | | (_| |  __/ |_) | |_| | (_| | (_| | |  __/
 *      \____/|_|  \__,_|_| |_|\__, |\___|____/ \__,_|\__, |\__, |_|\___|
 *                              __/ |                  __/ | __/ |       
 *                             |___/                  |___/ |___/        
 * Name  : graph_node.go
 * Desc  : This is the class that handles the graph mode.
 *
 * Author: Peter Antoine
 * Date  : 30/07/2023
 *--------------------------------------------------------------------------------*
 *                     Copyright (c) 2023 Peter Antoine
 *                            All rights Reserved.
 *                      Released Under the MIT Licence
 *--------------------------------------------------------------------------------*/

package graph_node

type GraphNode struct {
	parent		*GraphNode;
	children	map [uint8] *GraphNode;
	terminates	bool;
	parts		uint16;
}

func CreateNode (parent *GraphNode) *GraphNode {
	var result = new(GraphNode)
	result.parent = parent

	return result
}

func (g *GraphNode) SetTerminates() {
	g.terminates = true
}

func (g *GraphNode) SetParts(parts uint16) {
	g.parts = parts
}

func (g GraphNode) GetParts() uint16 {
	return g.parts
}

func (g GraphNode) IsTerminal() bool {
	return g.terminates
}

func (g GraphNode) FindChild(index uint8) *GraphNode {
	return g.children[index]
}

func (g *GraphNode) AddFindChild(index uint8) *GraphNode {

	var ok bool
	var result *GraphNode

	if result, ok = g.children[index]; !ok {
		// initialize the map if it has not been. Will save space but will
		// make creating nodes slower. Which is better?
		if g.children == nil {
			g.children = make(map[uint8] *GraphNode)
		}

		result = CreateNode(g)
		g.children[index] = result
	}

	return result
}

func (g GraphNode) WalkTree(items []uint8) (bool, []*GraphNode) {

	result := make([]*GraphNode, 0)
	var current = &g

	for _, value := range(items) {
		var ok bool
		if current, ok = current.children[value]; !ok {
			// didn't find the item in the graph.
			break
		} else {
			result = append(result, current)
		}
	}

	return false, result
}
