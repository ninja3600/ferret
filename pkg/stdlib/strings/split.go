package strings

import (
	"context"
	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/MontFerret/ferret/pkg/runtime/values"
	"strings"
)

/*
 * Splits the given string value into a list of strings, using the separator.
 * @params text (String) - The string to split.
 * @params separator (String) - The sperator.
 * @params limit (Int) - Limit the number of split values in the result. If no limit is given, the number of splits returned is not bounded.
 * @returns strings (Array<String>) - Array of strings.
 */
func Split(_ context.Context, args ...core.Value) (core.Value, error) {
	err := core.ValidateArgs(args, 2, 3)

	if err != nil {
		return values.None, err
	}

	text := args[0].String()
	separator := args[1].String()
	limit := -1

	if len(args) > 2 {
		if args[2].Type() == core.IntType {
			limit = int(args[2].(values.Int))
		}
	}

	var strs []string

	if limit < 0 {
		strs = strings.Split(text, separator)
	} else {
		strs = strings.SplitN(text, separator, limit)
	}

	arr := values.NewArray(len(strs))

	for _, str := range strs {
		arr.Push(values.NewString(str))
	}

	return arr, nil
}
