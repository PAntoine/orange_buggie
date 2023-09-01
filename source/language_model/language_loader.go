/**--------------------------------------------------------------------------------*
 *       ____                             ____                    _      
 *      / __ \                           |  _ \                  (_)     
 *     | |  | |_ __ __ _ _ __   __ _  ___| |_) |_   _  __ _  __ _ _  ___ 
 *     | |  | | '__/ _` | '_ \ / _` |/ _ \  _ <| | | |/ _` |/ _` | |/ _ \
 *     | |__| | | | (_| | | | | (_| |  __/ |_) | |_| | (_| | (_| | |  __/
 *      \____/|_|  \__,_|_| |_|\__, |\___|____/ \__,_|\__, |\__, |_|\___|
 *                              __/ |                  __/ | __/ |       
 *                             |___/                  |___/ |___/        
 * Name  : language_loader.go
 * Desc  : This will load the language model
 *
 * Author: Peter Antoine
 * Date  : 13/08/2023
 *--------------------------------------------------------------------------------*
 *                     Copyright (c) 2023 Peter Antoine
 *                            All rights Reserved.
 *                      Released Under the MIT Licence
 *--------------------------------------------------------------------------------*/

package language_model

import "os"
import "fmt"

func (l *LanguageModel) LoadLanguageModel(filename string) bool {
	result := false

	data, err := os.ReadFile(filename) // filename is the JSON file to read

	if err != nil {
		fmt.Println("Error: failed to read file '", filename , "'")

	} else if clauses, worked := l.parseGrammar(data); worked {
		result = l.buildSyntaxGraph(clauses)
	}

	return result
}

func fixupOptionalList(optional_list []*SyntaxNode) {
	for index, item := range(optional_list) {
		for _, item2 := range(optional_list[index+1:]) {
			item.AddLink(item2)
		}
	}
}

func (l *LanguageModel) buildSyntaxGraph(clauses clause_set) bool {

	for cid, items := range(clauses) {
		optional_list := []*SyntaxNode{}

		var prev *SyntaxNode = l.syntax_graph.GetRoot()

		for _, item := range(items) {
			if len(optional_list) > 0 && (item.flags & CF_OPTIONAL) != CF_OPTIONAL {
				// Ok, we have a non-optional token after one or more optional tokens, we need
				// to fix-up the connections between all the optional items and the non-optional.
				prev = CreateSyntaxNode(item.id)
				optional_list = append(optional_list, prev)

				fixupOptionalList(optional_list)

				// clear list now that we have handled it.
				optional_list = []*SyntaxNode{}

			} else if (item.flags & CF_OPTIONAL) == CF_OPTIONAL {
				// Ok, we have an optional node, might need to start the optional handling.
				if len(optional_list) == 0 {
					// Need to start the optional list, by adding prev as this will need to be connected.
					optional_list = append(optional_list, prev)
				}
				prev = CreateSyntaxNode(item.id)
				optional_list = append(optional_list, prev)

			} else {
				// Ok, handle normal non-optional entry
				prev = prev.AddOrFindChild(item.id)
			}

			if (item.flags & CF_MULIPLE) == CF_MULIPLE {
				// add back links
				prev.AddLink(prev)
			}
		}

		// At the end of the insert, are there any optional not fixed up?
		fixupOptionalList(optional_list)
		optional_list = []*SyntaxNode{}

		// Set the last node to flag that it is the end of a clause.
		prev.SetClauseID(cid)
	}

	return true
}

