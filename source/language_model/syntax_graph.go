/**--------------------------------------------------------------------------------*
 *       ____                             ____                    _      
 *      / __ \                           |  _ \                  (_)     
 *     | |  | |_ __ __ _ _ __   __ _  ___| |_) |_   _  __ _  __ _ _  ___ 
 *     | |  | | '__/ _` | '_ \ / _` |/ _ \  _ <| | | |/ _` |/ _` | |/ _ \
 *     | |__| | | | (_| | | | | (_| |  __/ |_) | |_| | (_| | (_| | |  __/
 *      \____/|_|  \__,_|_| |_|\__, |\___|____/ \__,_|\__, |\__, |_|\___|
 *                              __/ |                  __/ | __/ |       
 *                             |___/                  |___/ |___/        
 * Name  : syntax_graph.go
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

type SyntaxGraph struct {
	root *SyntaxNode
}

func (s *SyntaxGraph) Initialise () {
	s.root = CreateSyntaxNode(0)
}

func (s *SyntaxGraph) GetRoot() *SyntaxNode {
	return s.root
}

func (s *SyntaxGraph) AddChild (cid uint16) *SyntaxNode {
	return s.root.AddChild(cid)
}

func (s *SyntaxGraph) AddOrFindChild (cid uint16) *SyntaxNode {
	return s.root.AddOrFindChild(cid)
}

