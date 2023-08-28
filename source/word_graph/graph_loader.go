/**--------------------------------------------------------------------------------*
 *       ____                             ____                    _      
 *      / __ \                           |  _ \                  (_)     
 *     | |  | |_ __ __ _ _ __   __ _  ___| |_) |_   _  __ _  __ _ _  ___ 
 *     | |  | | '__/ _` | '_ \ / _` |/ _ \  _ <| | | |/ _` |/ _` | |/ _ \
 *     | |__| | | | (_| | | | | (_| |  __/ |_) | |_| | (_| | (_| | |  __/
 *      \____/|_|  \__,_|_| |_|\__, |\___|____/ \__,_|\__, |\__, |_|\___|
 *                              __/ |                  __/ | __/ |       
 *                             |___/                  |___/ |___/        
 * Name  : graph_loader.go
 * Desc  : This file holds the code to load the graph.
 *
 * Author: Peter Antoine
 * Date  : 08/08/2023
 *--------------------------------------------------------------------------------*
 *					   Copyright (c) 2023 Peter Antoine
 *							  All rights Reserved.
 *						Released Under the MIT Licence
 *--------------------------------------------------------------------------------*/

package word_graph

import "os"
import "fmt"
import "encoding/json"

func (w *WordGraph) LoadGraph(filename string) bool {

	result := false
	data, err := os.ReadFile(filename) // filename is the JSON file to read

	if err != nil {
		fmt.Println("Error: failed to read file '", filename , "'")
	} else {
		result = w.readJSONData(data)
	}

	return result
}

func (w *WordGraph) readJSONData(data []byte) bool {
	result := false
	var unpacked_data interface{}

	err := json.Unmarshal(data, &unpacked_data)

	if err != nil {
		fmt.Println("Cannot unmarshal the json ", err)
	} else {
		json_data := unpacked_data.(map[string] interface{})

		dictionary := json_data["dictionary"].(map[string] interface{})

		if (dictionary != nil) {
			language := dictionary["language"].(string)
			words    := dictionary["words"].(map[string] interface{})

			for key, value := range words {
				result = true
				w.decodeWord(key, value.(map[string] interface{}))
			}

			fmt.Println(language)
		}
	}

	return result
}

func (w *WordGraph) decodeWord(word string, details map[string]interface{}) bool {
	result := false

	fmt.Println(word)
	new_word := w.AddWord([]byte(word))

	if details["parts"] != nil {
		new_word.SetParts(decodeParts(details["parts"].([]interface{})))
	}

	fmt.Println(new_word)

	return result
}

func decodeParts(parts []interface{}) uint16 {
	var result uint16 = 0

	for _, part := range(parts) {
		fmt.Println(part.(string))
		switch (part.(string)) {
			case "noun":		result |= 0x01;
			case "preposition":	result |= 0x02;
			case "verb":		result |= 0x04;
		}
	}

	return result
}
