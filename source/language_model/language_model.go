/**--------------------------------------------------------------------------------*
 *       ____                             ____                    _      
 *      / __ \                           |  _ \                  (_)     
 *     | |  | |_ __ __ _ _ __   __ _  ___| |_) |_   _  __ _  __ _ _  ___ 
 *     | |  | | '__/ _` | '_ \ / _` |/ _ \  _ <| | | |/ _` |/ _` | |/ _ \
 *     | |__| | | | (_| | | | | (_| |  __/ |_) | |_| | (_| | (_| | |  __/
 *      \____/|_|  \__,_|_| |_|\__, |\___|____/ \__,_|\__, |\__, |_|\___|
 *                              __/ |                  __/ | __/ |       
 *                             |___/                  |___/ |___/        
 * Name  : language_model.go
 * Desc  : This is the main interface class with the language models.
 *
 * Author: Peter Antoine
 * Date  : 13/08/2023
 *--------------------------------------------------------------------------------*
 *                     Copyright (c) 2023 Peter Antoine
 *                            All rights Reserved.
 *                      Released Under the MIT Licence
 *--------------------------------------------------------------------------------*/

package language_model

import "fmt"

const (
	CF_MULIPLE	uint8 = 0x01
	CF_OPTIONAL	uint8 = 0x02
)

type LanguageModel struct {
	token_map		map[string]uint16
	syntax_graph	SyntaxGraph
}

type TokenItem struct {
	id		uint16
}

type clause_item struct {flags uint8; id uint16}
type clause_set = map[uint16] []clause_item

/* Create Language Model.
 *
 * This function will create the language model and initialise
 * the components.
 */
func CreateLanguageModel () *LanguageModel {
	result := new (LanguageModel)

	result.token_map = map[string]uint16{}
	result.syntax_graph.Initialise()

	return result
}

/* Add Token.
 *
 * This function will add a token to the language model. Tokens
 * are used to connect the word types in the dictionary to the
 * syntax graph. It is used in parsing the utterance.
 *
 * Tokens must be unique.
 */
func (l *LanguageModel) AddToken(name string) (uint16, bool) {
	result := uint16(0)
	worked := false

	if _, ok := l.token_map[name]; !ok {
		result = uint16(len(l.token_map) + 1)
		worked = true

		l.token_map[name] = result
	} else {
		fmt.Println("Warning: Duplicate token:", name)
	}

	return result, worked
}

/* Find Token By Name.
 *
 * This function will return the token ID for the token from the map
 * of tokens. Tokens are referenced by the token ID, which is 1 indexed.
 *
 * Zero is an invalid token.
 */
func (l *LanguageModel) FindTokenByName(name string) (uint16, bool) {
	result, ok := l.token_map[name]
	return result, ok
}

/* Is Token Valid.
 *
 * This is a simple test for the validity of the token. It will check that
 * the token is not zero and exists within the range of known tokens.
 */
func (l *LanguageModel) IsTokenValid(id uint16) bool {
	return id != 0 && int(id) <= len(l.token_map)
}
