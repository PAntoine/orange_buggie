/**--------------------------------------------------------------------------------*
 *       ____                             ____                    _      
 *      / __ \                           |  _ \                  (_)     
 *     | |  | |_ __ __ _ _ __   __ _  ___| |_) |_   _  __ _  __ _ _  ___ 
 *     | |  | | '__/ _` | '_ \ / _` |/ _ \  _ <| | | |/ _` |/ _` | |/ _ \
 *     | |__| | | | (_| | | | | (_| |  __/ |_) | |_| | (_| | (_| | |  __/
 *      \____/|_|  \__,_|_| |_|\__, |\___|____/ \__,_|\__, |\__, |_|\___|
 *                              __/ |                  __/ | __/ |       
 *                             |___/                  |___/ |___/        
 * Name  : syntax_node.go
 * Desc  : This function manages the syntax nodes.
 *
 * Author: Peter Antoine
 * Date  : 29/08/2023
 *--------------------------------------------------------------------------------*
 *                     Copyright (c) 2023 Peter Antoine
 *                            All rights Reserved.
 *                      Released Under the MIT Licence
 *--------------------------------------------------------------------------------*/

package language_model

type SyntaxNode struct {
	id			uint16
	clause_id	uint16
	children	map[uint16]*SyntaxNode
}

/* CreateSyntaxNode
 *
 * This will create a syntax node and initialise it.
 */
func CreateSyntaxNode(id uint16) *SyntaxNode {
	result := new (SyntaxNode)
	result.id = id
	result.children = map[uint16]*SyntaxNode{}

	return result
}

/* AddChild
 *
 * This function will add a child to the SyntaxNode. If the id already
 * exists in the parent node then the function will return nil. You
 * should use AddOrFindChild to non-unique add a child node.
 */
func (sn *SyntaxNode) AddChild(id uint16) *SyntaxNode {
	var result *SyntaxNode

	if _, ok := sn.children[id]; !ok {
		result = CreateSyntaxNode(id)
		sn.children[id] = result
	}

	return result
}

/* AddOrFindChild
 *
 * This function is like the AddNode function but will return the already
 * existing child node, if the cid already exists in the tree.
 */
func (sn *SyntaxNode) AddOrFindChild(id uint16) *SyntaxNode {
	var result *SyntaxNode
	ok := false

	if result, ok = sn.children[id]; !ok {
		result = CreateSyntaxNode(id)
		sn.children[id] = result
	}

	return result
}

/* Add Link.
 *
 * This function will link one node to another one. It is allowed for a
 * link to loop back to itself as this is valid for repeating nodes.
 */
func (sn *SyntaxNode) AddLink(node *SyntaxNode) bool {
	if found, ok := sn.children[node.id]; !ok {
		// Only allow for non-connected nodes to be connected.
		sn.children[node.id] = node

		return true
	} else {
		return found == node
	}
}

func (sn *SyntaxNode) SetClauseID (id uint16) {
	sn.clause_id = id
}
