/**--------------------------------------------------------------------------------*
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

func TestTokens(t* testing.T) {
	lm := CreateLanguageModel()

	if lm.FindTokenByID(0) != nil {
		t.Logf("Failed as found a token when it should not have.")
		t.FailNow()
	}

	if found := lm.FindTokenByName("verb"); found != nil {
		t.Logf("Failed as found a token when it should not have.")
		t.FailNow()
	}

	if first, ok := lm.AddToken("verb"); !ok {
		t.Logf("Failed to add the first token.")
		t.FailNow()
	} else {
		if found := lm.FindTokenByID(first); found == nil {
			t.Logf("Failed to find just added token by ID.")
			t.FailNow()
		}

		if found := lm.FindTokenByName("verb"); found == nil {
			t.Logf("Failed to find just added token by Name.")
			t.FailNow()
		}
	}

	if _, ok := lm.AddToken("verb"); ok {
		t.Logf("Failed as added the same token name twice.");
		t.FailNow()
	}

	if first, ok := lm.AddToken("noun"); !ok {
		t.Logf("Failed to add the next token.")
		t.FailNow()
	} else {
		if first == 0 {
			t.Logf("Failed the index should be greater than zero.")
			t.FailNow()
		}

		if found := lm.FindTokenByID(first); found == nil {
			t.Logf("Failed to find just added token by ID.")
			t.FailNow()
		}

		if found := lm.FindTokenByName("verb"); found == nil {
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
		if _, ok := lm.parseGrammar(array); !ok {
			t.Logf("Failed to properly decode line endings.")
			t.Logf("test case: %s", array)
			t.Fail()
		}
	}
}

func TestParseGrammar(t *testing.T) {
	lm := CreateLanguageModel()

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
		}

	for _, line := range(test_cases) {
		if _, ok := lm.parseGrammar([]byte(line)); !ok {
			t.Logf("Failed to parse name, test case: %s", line)
			t.Fail()
			fmt.Println(line)
			fmt.Println(lm.parseGrammar([]byte(line)))
			fmt.Println("----------------------------------")
		}
	}

	negative_test := "name4 name = test\n"
	if _, ok := lm.parseGrammar([]byte(negative_test)); ok {
		t.Logf("Parsed a line it should not, test case: %s", negative_test)
		t.Fail()
	}
}

func TestLoadLanguageModel(t *testing.T) {
	test_model := "name_name = {test} [test1]\n\n\nname_name1=test\n"

	lm := CreateLanguageModel()
	if clauses, worked := lm.parseGrammar([]byte(test_model)); !worked {
		t.Logf("Failed to parse the grammer.")
		t.FailNow()
	} else {
		if !lm.buildParserTree(clauses) {
			t.Logf("Failed to build parser tree.");
			t.FailNow()
		}
	}
}
