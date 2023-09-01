/**--------------------------------------------------------------------------------*
 *       ____                             ____                    _      
 *      / __ \                           |  _ \                  (_)     
 *     | |  | |_ __ __ _ _ __   __ _  ___| |_) |_   _  __ _  __ _ _  ___ 
 *     | |  | | '__/ _` | '_ \ / _` |/ _ \  _ <| | | |/ _` |/ _` | |/ _ \
 *     | |__| | | | (_| | | | | (_| |  __/ |_) | |_| | (_| | (_| | |  __/
 *      \____/|_|  \__,_|_| |_|\__, |\___|____/ \__,_|\__, |\__, |_|\___|
 *                              __/ |                  __/ | __/ |       
 *                             |___/                  |___/ |___/        
 * Name  : syntax_parser.go
 * Desc  : This file holds the syntax parser functions.
 *
 * Author: Peter Antoine
 * Date  : 29/08/2023
 *--------------------------------------------------------------------------------*
 *                     Copyright (c) 2023 Peter Antoine
 *                            All rights Reserved.
 *                      Released Under the MIT Licence
 *--------------------------------------------------------------------------------*/

package language_model

import "fmt"

func (l *LanguageModel) parseGrammar(data []byte) (clause_set, bool) {
	var clauses clause_set

	result := false

	if index, line_number, ok := l.parseTokenDefinitions(data, 0); ok {
		clauses, result = l.parseClauseList(data, line_number, index)
	}

	return clauses, result
}

func (l *LanguageModel) parseTokenDefinitions(data []byte, index int) (int, int, bool) {
	worked := false
	line_number := 0

	if index, line_number, worked = findDirectiveSection(data, index, line_number, "tokens"); worked {
		for {
				var found bool
				index = eatWhiteSpace(data, index)
				for {
					if found, index = isLineEnding(data, index); found {
						index = eatWhiteSpace(data, index)
						line_number++
					} else {
						break
					}
				}

				var token string

				if index, worked = getNameFromData(data, index, &token); worked {
					if _, worked = l.AddToken(token, false); !worked {
						break;
					}
				} else if data[index] == '%' {
					// end of section
					worked = true
					break
				} else {
					fmt.Println(data[index], string(data[index:]))
					break
			}
		}
	}

	return index, line_number, worked
}

func (l *LanguageModel) parseClauseList(data []byte, line_number int, index int) (clause_set, bool) {
	clauses := make(clause_set)
	worked  := true

	if index, line_number, worked = findDirectiveSection(data, index, line_number, "rules"); worked {
		for index < len(data) {

			for {
				var found bool
				index = eatWhiteSpace(data, index)
				if  found, index = isLineEnding(data, index); found {
					// eat white space lines.
					line_number++
				} else {
					break
				}
			}

			if data[index] == '#' {
				// eat comment lines.
				index = eatWholeLine(data, index)
			} else {
				// now parse the clause list
				var tokens []clause_item
				var clause_id uint16
				clause_id, tokens, line_number, index = l.parseClauseLine(data, line_number, index)

				if len(tokens) == 0 {
					fmt.Printf("line %3d: Error clause must have at least one token.\n", line_number)
					worked = false
					break
				} else {
					clauses[clause_id] = tokens
				}

				// start the next parse
				index = eatWhiteSpace(data, index)
			}
		}
	}

	return clauses, worked
}

func (l *LanguageModel) parseClauseLine(data []byte, line_number int, index int) (uint16, []clause_item, int, int) {
	var tokens = []clause_item{}
	var ok bool
	var found bool = true
	var clause_id uint16

	if clause_id, index, ok = l.decodeClauseName(data, index); !ok {
		fmt.Printf("line %3d: Invalid clause name.\n", line_number)
		index = eatWholeLine(data, index)
		line_number++
	} else {
		for found {
			var token string
			var token_flags uint8

			index = eatWhiteSpace(data, index)

			if index, found = getLineToken(data, index, &token, &token_flags); found {
				if token_item, ok := l.FindTokenByName(token); ok {
					tokens = append(tokens, clause_item{token_flags, token_item})
					index = eatWhiteSpace(data, index)
				} else {
					fmt.Printf("line %3d: Error unknown token '%s' found.\n", line_number, token)
				}
			}

			var is_line_ending bool
			if is_line_ending, index = isLineEnding(data, index); is_line_ending {
				for is_line_ending {
					line_number++
					is_line_ending, index = isLineEnding(data, index)
				}
				break
			}
		}
	}

	return clause_id, tokens, line_number, index
}

func findDirectiveSection(data []byte, index int, line_number int, directive string) (int, int, bool) {
	found := false
	i := index
	dir_len := len(directive) + 1

	for {
		if data[i] == '#' {
			i = eatWholeLine(data, i)

		} else if data[i] == '%' && (i + dir_len) < len(data) && string(data[i+1:i+dir_len]) == directive {
			i = eatWhiteSpace(data, i + dir_len)
			if found, i = isLineEnding(data, i); found {
				line_number++
			} else {
				fmt.Printf("line %3d: directive '%s' must be followed by a line ending and no other chars.\n", line_number, directive)
			}
			break
		} else {
				i = eatWhiteSpace(data, i)
				if  found, i = isLineEnding(data, i); found {
					line_number++
				} else {
				fmt.Printf("line %3d: directive '%s' must be at the begining of a line.\n", line_number, directive)
				break
			}
		}
	}

	return i, line_number, found
}

func (l *LanguageModel) decodeClauseName(data []byte, index int) (uint16, int, bool) {
	// Get clause name
	var found bool
	var name string
	var worked bool = false
	var result uint16 = 0

	index, found = getNameFromData(data, index, &name)
	index = eatWhiteSpace(data, index)

	if !found {
		fmt.Println("clause must start with a name.")

	} else if index == len(data) || data[index] != '=' {

		fmt.Println("Invalid clause, '=' must follow name", data[index])

	} else {
		// skip the '='
		index++
		index = eatWhiteSpace(data, index)

		// add the clause name token
		result, worked = l.AddToken(name, true)
	}

	return result, index, worked
}

func isLineEnding(data []byte, index int) (bool, int) {
	offset := index
	is_le  := false

	if index >= len(data) {
		offset++

	} else if data[index] == 0x0a {
		offset++

		if offset < len(data) && data[offset] == 0x0d {
			offset++
		}

		is_le = true
	} else if data[index] == 0x0d {
		offset++

		if offset < len(data) && data[offset] == 0x0a {
			offset++
		}

		is_le = true
	}

	return is_le, offset
}

func eatWhiteSpace(data []byte, index int) int {
	i := index

	for i = index; i < len(data); i++ {
		if !(data[i] == ' ' || data[i] == '\t') {
			break
		}
	}

	return i
}

func eatWholeLine(data[]byte, index int) int {
	i := index
	var fini bool

	for {
		if fini, i = isLineEnding(data, i); fini {
			break
		}
		i++
	}

	return i
}

func getNameFromData(data []byte, index int, name *string) (int, bool) {

	result := false
	i := index
	var new_name []byte

	for i = index; i < len(data); i++ {
		if (data[i] >= 65 && data[i] <= 90)   ||
		   (data[i] >= 97 && data[i] <= 122)  ||
		   (data[i] >= 48 && data[i] <= 57)   ||
		   data[i] == '_' {
			// Is a valid character for a name
			new_name = append(new_name, data[i])

		} else if data[i] == '%' || data[i] == ' ' || data[i] == '\t' || data[i] == ']' ||
				  data[i] == '}' || data[i] == ':' || data[i] == '=' || data[i] == 0x0a || data[i] == 0x0d {
			// found  delimiter of the name
			*name = string(new_name)
			result = len(*name) > 0

			break
		} else {
			fmt.Println("Invalid character in name:", data[i])
			break
		}
	}

	return i, result
}

func getLineToken(data []byte, index int, token *string, flags *uint8) (int, bool) {

	result := false
	worked := false
	offset := index

	*flags = 0

	if index < len(data) && data[index] == '[' {
		*flags |= CF_OPTIONAL
		index++
	}

	if index < len(data) && data[index] == '{' {
		*flags |= CF_MULIPLE
		index++
	}

	offset, worked = getNameFromData(data, index, token)

	offset = eatWhiteSpace(data, offset)

	if worked {
		if index < len(data) && (*flags & CF_MULIPLE) == CF_MULIPLE {
			if data[offset] == '}' {
				offset++
				offset = eatWhiteSpace(data, offset)
				result = true
			}

		}

		if index < len(data) && (*flags & CF_OPTIONAL) == CF_OPTIONAL {
			if data[offset] == ']' {
				offset++
				offset = eatWhiteSpace(data, offset)
				result = true
			}

		}

		result = true
	}

	return offset, result
}

