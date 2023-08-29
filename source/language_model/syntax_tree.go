/**--------------------------------------------------------------------------------*
 *       ____                             ____                    _      
 *      / __ \                           |  _ \                  (_)     
 *     | |  | |_ __ __ _ _ __   __ _  ___| |_) |_   _  __ _  __ _ _  ___ 
 *     | |  | | '__/ _` | '_ \ / _` |/ _ \  _ <| | | |/ _` |/ _` | |/ _ \
 *     | |__| | | | (_| | | | | (_| |  __/ |_) | |_| | (_| | (_| | |  __/
 *      \____/|_|  \__,_|_| |_|\__, |\___|____/ \__,_|\__, |\__, |_|\___|
 *                              __/ |                  __/ | __/ |       
 *                             |___/                  |___/ |___/        
 * Name  : syntax_tree.go
 * Desc  : This hold the code to handle the syntax parsing.
 *
 * Author: Peter Antoine
 * Date  : 27/08/2023
 *--------------------------------------------------------------------------------*
 *                     Copyright (c) 2023 Peter Antoine
 *                            All rights Reserved.
 *                      Released Under the MIT Licence
 *--------------------------------------------------------------------------------*/

package language_model

type SyntaxTree struct {
}

type SyntaxNode struct {
}

func (s *SyntaxTree) AddNode() *SyntaxNode {
	return new (SyntaxNode)
}

func (l *LanguageModel) buildSyntaxTree(clauses clause_set) bool {
	var node_list []*SyntaxNode

	for _, _ = range(clauses) {
		node_list = append(node_list, l.syntax_tree.AddNode())
	}

	return true
}

