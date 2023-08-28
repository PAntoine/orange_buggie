/**--------------------------------------------------------------------------------*
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
	tokens		[]*TokenItem
	token_map	map[string]*TokenItem
	syntax_tree SyntaxTree
}

type TokenItem struct {
	id		uint16
}

type clause_item struct {flags uint8; item *TokenItem}
type clause_set = map[uint16] []clause_item

func CreateLanguageModel () *LanguageModel {
	result := new (LanguageModel)

	result.token_map = map[string]*TokenItem{}

	return result
}

func (l *LanguageModel) AddToken(name string) (uint16, bool) {
	result := uint16(0)
	worked := false

	if _, ok := l.token_map[name]; !ok {
		result = uint16(len(l.tokens))
		worked = true

		item := new (TokenItem)
		item.id = result

		l.tokens = append(l.tokens, item)
		l.token_map[name] = item
	} else {
		fmt.Println("Warning: Duplicate token:", name)
	}

	return result, worked
}

func (l *LanguageModel) FindTokenByName(name string) *TokenItem {
	return l.token_map[name]
}

func (l *LanguageModel) FindTokenByID(id uint16) *TokenItem {
	var result *TokenItem = nil

	if id < uint16(len(l.tokens)) {
		result =  l.tokens[id]
	}

	return result
}
