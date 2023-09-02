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

func (s SyntaxGraph) StartParse(id uint16) *SyntaxNode {
	return s.root.FindChild(id)
}

func (s SyntaxGraph) ParseSyntax(token_list []uint16) uint16 {
	result := uint16(0)

	match_list := [][]uint16{token_list}

	for len(match_list) > 0 {
		if len(match_list[0]) > 0 {
			current := s.root.FindChild(match_list[0][0])

			if current != nil {
				index := 1

				for ; index < len(match_list[0]); index++ {
					if current = current.FindChild(match_list[0][index]); current == nil {
						break
					} else if (current.GetClause() != 0) && index != (len(match_list[0]) - 1) {
						// Ok, we have found a sub-clause in the sentence - let's remember this
						// if the current node is not the last node in the current search.
						match_list = append(match_list, append([]uint16{current.GetClause()}, match_list[0][index+1:]...))
					}
				}

				if index == len(match_list[0]) && current != nil {
					result = current.GetClause()
				}
			}
		}

		if result != 0 {
			break;
		} else {
			// take the first item off the list.
			match_list = match_list[1:]
		}
	}

	return result
}

