/**--------------------------------------------------------------------------------*
 *       ____                             ____                    _      
 *      / __ \                           |  _ \                  (_)     
 *     | |  | |_ __ __ _ _ __   __ _  ___| |_) |_   _  __ _  __ _ _  ___ 
 *     | |  | | '__/ _` | '_ \ / _` |/ _ \  _ <| | | |/ _` |/ _` | |/ _ \
 *     | |__| | | | (_| | | | | (_| |  __/ |_) | |_| | (_| | (_| | |  __/
 *      \____/|_|  \__,_|_| |_|\__, |\___|____/ \__,_|\__, |\__, |_|\___|
 *                              __/ |                  __/ | __/ |       
 *                             |___/                  |___/ |___/        
 * Name  : language_loader_test.go
 * Desc  : This is the unit tests for the language loaded.
 *
 * Author: Peter Antoine
 * Date  : 16/08/2023
 *--------------------------------------------------------------------------------*
 *                     Copyright (c) 2023 Peter Antoine
 *                            All rights Reserved.
 *                      Released Under the MIT Licence
 *--------------------------------------------------------------------------------*/

package language_model

import "fmt"
import "testing"

func TestFindTokenSection(t* testing.T) {
	test_data := "# some comment\n# some other comment\n# %token this is a comment\n#%token so is this\n\n\n \n\t\n\n%token\n"
	if i, _, ok := findDirectiveSection([]byte(test_data),0, 0, "token"); !ok || (i != 96 && i != 95) {	// One and Two byte line endings.
		t.Logf("Failed to find a valid token in the correct place.")
	}

	test_array := []string {
				" %tokens\n",
				"%tokens :\n",
				"%tokens:\n",
	}

	for _, line := range(test_array) {
		if _, _, ok := findDirectiveSection([]byte(line),0,0,"token"); ok {
			t.Logf("Failed, as decoded an invalid line")
			t.Fail()
		}
	}
}

func TestTokens(t* testing.T) {
	lm := CreateLanguageModel()

	if lm.IsTokenValid(0) {
		t.Logf("Failed as found a token when it should not have.")
		t.FailNow()
	}

	if lm.IsTokenValid(1) {
		t.Logf("Failed as found a token when it should not have.")
		t.FailNow()
	}

	if _, ok := lm.FindTokenByName("verb"); ok {
		t.Logf("Failed as found a token when it should not have.")
		t.FailNow()
	}

	if first, ok := lm.AddToken("verb", false); !ok {
		t.Logf("Failed to add the first token.")
		t.FailNow()
	} else {
		if !lm.IsTokenValid(first) {
			t.Logf("Failed to find just added token by ID.")
			t.FailNow()
		}

		if _, ok := lm.FindTokenByName("verb"); !ok {
			t.Logf("Failed to find just added token by Name.")
			t.FailNow()
		}
	}

	if _, ok := lm.AddToken("verb", false); ok {
		t.Logf("Failed as added the same token name twice.");
		t.FailNow()
	}

	if first, ok := lm.AddToken("noun", false); !ok {
		t.Logf("Failed to add the next token.")
		t.FailNow()
	} else {
		if first == 0 {
			t.Logf("Failed the index should be greater than zero.")
			t.FailNow()
		}

		if !lm.IsTokenValid(first) {
			t.Logf("Failed to find just added token by ID.")
			t.FailNow()
		}

		if _, ok := lm.FindTokenByName("verb"); !ok {
			t.Logf("Failed to find just added token by Name.")
			t.FailNow()
		}
	}
}

func TestGetNameFromData(t* testing.T) {
	var result string

	index, worked := getNameFromData([]byte("test_one_12 = "), 0, &result)

	fmt.Println(index, worked, result)
}

func TestParseGrammarLineEndings(t *testing.T) {
	lm := CreateLanguageModel()
	lm.AddToken("test", false)

	test_cases := [][]byte{
				[]byte{110,97,109,101,49,48,61,116,101,115,116,0x0a},
				[]byte{110,97,109,101,49,49,61,116,101,115,116,0x0d},
				[]byte{110,97,109,101,49,50,61,116,101,115,116,0x0a, 0x0d},
				[]byte{110,97,109,101,49,51,61,116,101,115,116,0x0d, 0x0a},
				[]byte{110,97,109,101,49,52,61,116,101,115,116,0x0a, 0x0a},
				[]byte{110,97,109,101,49,53,61,116,101,115,116,0x0d, 0x0d},
				[]byte{110,97,109,101,49,54,61,116,101,115,116,0x0a, 0x0a, 0x0a},
			}

	for _, array := range(test_cases) {
		line_number := 0
		if _, clauses, ln, _ := lm.parseClauseLine(array, line_number, 0); len(clauses) == 0 {
			t.Logf("Failed to properly decode line endings. %d", ln)
			t.Logf("test case: %s", array)
			t.Fail()
		}
	}
}

func TestParseGrammar(t *testing.T) {
	lm := CreateLanguageModel()

	lm.AddToken("test", false)
	lm.AddToken("test1", false)

	test_cases := []string {
			"name1= test\n",
			"name2 = test\n",
			"name3\t= test\n",
			"name_name5 = test\n",
			"name_name6 = test test1\n",
			"name_name7 = test test1\nname_nameb=test\n",
			"name_name8 = test test1\n\nname_namec=test\n",
			"name_name9 = test test1\n\n\nname_named=test\n",
			"name_namea = {test} [test1]\n\n\nname_namee=test\n",
			"name_namef = [{test}] [test1]\n\n\nname_nameh=test\n",
		}

	for _, line := range(test_cases) {
		line_number := 0
		if _, clauses, _, _ := lm.parseClauseLine([]byte(line), line_number, 0); len(clauses) < 1 {
			t.Logf("Failed to parse name, test case: '%s' %d", line, len(clauses))
			t.Fail()
		}
	}

	line_number := 0
	negative_test := "name4 name = test\n"
	if _, clauses, _, _ := lm.parseClauseLine([]byte(negative_test), line_number, 0); len(clauses) != 0 {
		t.Logf("Parsed a line it should not, test case: %s", negative_test)
		t.Fail()
	}
}

func TestParseTokenSection(t *testing.T) {
	lm := CreateLanguageModel()

	test_model := "%tokens\none two three\nfour\nfive six\nseven eight nine ten\n%rules\n"

	if _,_,ok := lm.parseTokenDefinitions([]byte(test_model), 0); !ok {
		t.Logf("Failed to parse the token defintions")
		t.Fail()
	}

	// duplicate token test
	test_model = "%tokens\none two three\nfour\none six\nseven eight nine ten\n%rules\n"

	if _,_,ok := lm.parseTokenDefinitions([]byte(test_model), 0); ok {
		t.Logf("Failed to detect the duplicate token")
		t.Fail()
	}
}

func TestLoadLanguageModel(t *testing.T) {
	test_model := "%rules\nname_name1 = {test} [test1]\n\n\nname_name2=test2\n"

	lm := CreateLanguageModel()
	lm.AddToken("test", false)
	lm.AddToken("test1", false)
	lm.AddToken("test2", false)

	line_number := 0
	if clauses, worked := lm.parseClauseList([]byte(test_model), line_number, 0); !worked {
		t.Logf("Failed to parse the grammer.")
		t.FailNow()
	} else {
		if !lm.buildSyntaxGraph(clauses) {
			t.Logf("Failed to build parser tree.");
			t.FailNow()
		} else {
			item0, _ := lm.FindTokenByName("test")
			item1, _ := lm.FindTokenByName("test1")
			item2, _ := lm.FindTokenByName("test2")
			clause_id1, _ := lm.FindTokenByName("name_name1")
			clause_id2, _ := lm.FindTokenByName("name_name2")

			item_list := []uint16{item0, item1}
			fmt.Println(lm.ParseSyntax(item_list), clause_id1)

			item_list = []uint16{item0}
			fmt.Println(lm.ParseSyntax(item_list), clause_id1)

			item_list = []uint16{item1}
			fmt.Println(lm.ParseSyntax(item_list), clause_id1)

			item_list = []uint16{item0, item0, item0}
			fmt.Println(lm.ParseSyntax(item_list), clause_id1)

			item_list = []uint16{item2}
			fmt.Println(lm.ParseSyntax(item_list), clause_id2)
		}
	}
}

func TestLoadLanguageModelFile(t *testing.T) {
	lm := CreateLanguageModel()

	if !lm.LoadLanguageModel("../../test_data/english_grammar.grammar") {
		t.Logf("Failed to load language model file.");
		t.FailNow()
	}
}
