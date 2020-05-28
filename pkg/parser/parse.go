//line parse.rl:1
package parser

import "strings"
import "fmt"
import "strconv"

//line parse.rl:8

//line parse.go:13
const query_start int = 14
const query_first_final int = 14
const query_error int = 0

const query_en_main int = 14

//line parse.rl:9

func ParseParams(data string) (*Params, error) {
	p, cs, pe, eof := 0, 0, len(data), len(data)

	var ts, te, act int

	mark := 0
	params := map[string]interface{}{}
	param := ""

//line parse.go:33
	{
		cs = query_start
		ts = 0
		te = 0
		act = 0
	}

//line parse.go:41
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 14:
			goto st_case_14
		case 0:
			goto st_case_0
		case 15:
			goto st_case_15
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 16:
			goto st_case_16
		case 3:
			goto st_case_3
		case 17:
			goto st_case_17
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 18:
			goto st_case_18
		case 6:
			goto st_case_6
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 7:
			goto st_case_7
		case 21:
			goto st_case_21
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 22:
			goto st_case_22
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		}
		goto st_out
	tr0:
//line parse.rl:51
		p = (te) - 1

		goto st14
	tr22:
//line parse.rl:52
		te = p + 1

		goto st14
	tr25:
//line parse.rl:21

		if _, ok := params[param]; !ok {
			params[param] = true
		}
		param = strings.ReplaceAll(data[mark:p], "-", "_")

//line parse.rl:51
		te = p
		p--

		goto st14
	tr28:
//line parse.rl:33

		params[param] = data[mark+1 : p-1]

//line parse.rl:51
		te = p
		p--

		goto st14
	tr29:
//line parse.rl:27

		params[param] = NewRangeFromString(data[mark:p])

//line parse.rl:51
		te = p
		p--

		goto st14
	tr30:
//line parse.rl:39

		params[param], _ = strconv.Atoi(data[mark:p])

//line parse.rl:51
		te = p
		p--

		goto st14
	tr32:
//line parse.rl:36

		params[param] = NewRangeFromString(data[mark:p])

//line parse.rl:51
		te = p
		p--

		goto st14
	tr33:
//line parse.rl:30

		params[param] = data[mark:p] == "true"

//line parse.rl:51
		te = p
		p--

		goto st14
	st14:
//line NONE:1
		ts = 0

		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line NONE:1
		ts = p

//line parse.go:179
		switch data[p] {
		case 32:
			goto tr22
		case 45:
			goto tr24
		case 95:
			goto tr24
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr22
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	tr26:
//line NONE:1
		te = p + 1

		goto st15
	tr24:
//line NONE:1
		te = p + 1

//line parse.rl:20
		mark = p
		goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line parse.go:222
		switch data[p] {
		case 45:
			goto tr26
		case 58:
			goto tr27
		case 95:
			goto tr26
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr26
			}
		case data[p] >= 65:
			goto tr26
		}
		goto tr25
	tr27:
//line parse.rl:21

		if _, ok := params[param]; !ok {
			params[param] = true
		}
		param = strings.ReplaceAll(data[mark:p], "-", "_")

		goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line parse.go:254
		switch data[p] {
		case 34:
			goto tr1
		case 39:
			goto tr2
		case 45:
			goto tr3
		case 60:
			goto tr5
		case 62:
			goto tr5
		case 102:
			goto tr6
		case 116:
			goto tr7
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr4
		}
		goto tr0
	tr1:
//line parse.rl:20
		mark = p
		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line parse.go:284
		switch data[p] {
		case 34:
			goto st16
		case 92:
			goto st3
		}
		goto st2
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		goto tr28
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 34:
			goto tr11
		case 92:
			goto st3
		}
		goto st2
	tr11:
//line NONE:1
		te = p + 1

		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line parse.go:320
		switch data[p] {
		case 34:
			goto st16
		case 92:
			goto st3
		}
		goto st2
	tr2:
//line parse.rl:20
		mark = p
		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line parse.go:337
		switch data[p] {
		case 39:
			goto st16
		case 92:
			goto st5
		}
		goto st4
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 39:
			goto tr14
		case 92:
			goto st5
		}
		goto st4
	tr14:
//line NONE:1
		te = p + 1

		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line parse.go:367
		switch data[p] {
		case 39:
			goto st16
		case 92:
			goto st5
		}
		goto st4
	tr3:
//line parse.rl:20
		mark = p
		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line parse.go:384
		if 48 <= data[p] && data[p] <= 57 {
			goto st19
		}
		goto tr0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if 48 <= data[p] && data[p] <= 57 {
			goto st19
		}
		goto tr29
	tr4:
//line parse.rl:20
		mark = p
		goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line parse.go:407
		if data[p] == 45 {
			goto st19
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st20
		}
		goto tr30
	tr5:
//line parse.rl:20
		mark = p
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line parse.go:424
		if 48 <= data[p] && data[p] <= 57 {
			goto st21
		}
		goto tr0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if 48 <= data[p] && data[p] <= 57 {
			goto st21
		}
		goto tr32
	tr6:
//line parse.rl:20
		mark = p
		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line parse.go:447
		if data[p] == 97 {
			goto st9
		}
		goto tr0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 108 {
			goto st10
		}
		goto tr0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if data[p] == 115 {
			goto st11
		}
		goto tr0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if data[p] == 101 {
			goto st22
		}
		goto tr0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		goto tr33
	tr7:
//line parse.rl:20
		mark = p
		goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line parse.go:494
		if data[p] == 114 {
			goto st13
		}
		goto tr0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if data[p] == 117 {
			goto st11
		}
		goto tr0
	st_out:
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 15:
				goto tr25
			case 1:
				goto tr0
			case 2:
				goto tr0
			case 16:
				goto tr28
			case 3:
				goto tr0
			case 17:
				goto tr28
			case 4:
				goto tr0
			case 5:
				goto tr0
			case 18:
				goto tr28
			case 6:
				goto tr0
			case 19:
				goto tr29
			case 20:
				goto tr30
			case 7:
				goto tr0
			case 21:
				goto tr32
			case 8:
				goto tr0
			case 9:
				goto tr0
			case 10:
				goto tr0
			case 11:
				goto tr0
			case 22:
				goto tr33
			case 12:
				goto tr0
			case 13:
				goto tr0
			}
		}

	_out:
		{
		}
	}

//line parse.rl:56

	if cs < query_first_final {
		return nil, fmt.Errorf("Cannot parse: %d", p)
	}

	if _, ok := params[param]; !ok {
		params[param] = true
	}

	ts, act = act, ts

	return NewParams(params), nil
}
