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
		result = l.buildSyntaxTree(clauses)
	}

	return result
}

