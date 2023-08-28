/**--------------------------------------------------------------------------------*
 * Name  : run_model.go
 * Desc  : This function will run a language model.
 *
 * Author: Peter Antoine
 * Date  : 13/08/2023
 *--------------------------------------------------------------------------------*
 *                     Copyright (c) 2023 Peter Antoine
 *                            All rights Reserved.
 *                      Released Under the MIT Licence
 *--------------------------------------------------------------------------------*/

package main
import "fmt"
import "orange_buggie/source/language_model"

func main() {
	var meh = language_model.CreateLanguageModel()
	fmt.Println(meh.LoadLanguageModel("en_gb.langmod"))
}
