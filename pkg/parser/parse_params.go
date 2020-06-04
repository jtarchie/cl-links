//line parse_params.rl:1
package parser

import (
	"fmt"
	"strconv"
	"strings"
)
//line parse_params.rl:8

//line parse_params.go:13
const query_start int = 12
const query_first_final int = 12
const query_error int = 0

const query_en_main int = 12

//line parse_params.rl:9

func ParseParams(data string) (*Params, error) {
	p, cs, pe, eof := 0, 0, len(data), len(data)

	var ts, te, act int

	mark := 0
	params := map[string]interface{}{}
	param := ""

//line parse_params.go:33
	{
		cs = query_start
		ts = 0
		te = 0
		act = 0
	}

//line parse_params.go:41
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 12:
			goto st_case_12
		case 0:
			goto st_case_0
		case 13:
			goto st_case_13
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 14:
			goto st_case_14
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 15:
			goto st_case_15
		case 8:
			goto st_case_8
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 9:
			goto st_case_9
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		}
		goto st_out
	tr0:
//line parse_params.rl:57
		p = (te) - 1

		goto st12
	tr10:
//line parse_params.rl:20
		mark = p
//line parse_params.rl:33

		params[param] = data[mark:p]

//line parse_params.rl:57
		te = p + 1

		goto st12
	tr13:
//line parse_params.rl:33

		params[param] = data[mark:p]

//line parse_params.rl:57
		te = p + 1

		goto st12
	tr24:
//line parse_params.rl:54
		te = p + 1
		{
			return ParseURL(data)
		}
		goto st12
	tr25:
//line parse_params.rl:58
		te = p + 1

		goto st12
	tr29:
//line parse_params.rl:21

		if _, ok := params[param]; !ok && param != "" {
			params[param] = true
		}
		param = strings.ReplaceAll(data[mark:p], "-", "_")

//line parse_params.rl:57
		te = p
		p--

		goto st12
	tr32:
//line parse_params.rl:57
		te = p
		p--

		goto st12
	tr33:
//line parse_params.rl:27

		params[param] = NewRangeFromString(data[mark:p])

//line parse_params.rl:57
		te = p
		p--

		goto st12
	tr34:
//line parse_params.rl:39

		params[param], _ = strconv.Atoi(data[mark:p])

//line parse_params.rl:57
		te = p
		p--

		goto st12
	tr36:
//line parse_params.rl:36

		params[param] = NewRangeFromString(data[mark:p])

//line parse_params.rl:57
		te = p
		p--

		goto st12
	tr37:
//line parse_params.rl:33

		params[param] = data[mark:p]

//line parse_params.rl:57
		te = p
		p--

		goto st12
	tr43:
//line parse_params.rl:30

		params[param] = data[mark:p] == "true"

//line parse_params.rl:57
		te = p
		p--

		goto st12
	st12:
//line NONE:1
		ts = 0

		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line NONE:1
		ts = p

//line parse_params.go:230
		switch data[p] {
		case 32:
			goto tr25
		case 45:
			goto tr27
		case 95:
			goto tr27
		case 104:
			goto tr28
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr25
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr27
			}
		default:
			goto tr27
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	tr27:
//line NONE:1
		te = p + 1

//line parse_params.rl:20
		mark = p
		goto st13
	tr30:
//line NONE:1
		te = p + 1

		goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line parse_params.go:275
		switch data[p] {
		case 45:
			goto tr30
		case 58:
			goto tr31
		case 95:
			goto tr30
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr30
			}
		case data[p] >= 65:
			goto tr30
		}
		goto tr29
	tr31:
//line parse_params.rl:21

		if _, ok := params[param]; !ok && param != "" {
			params[param] = true
		}
		param = strings.ReplaceAll(data[mark:p], "-", "_")

		goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line parse_params.go:307
		switch data[p] {
		case 34:
			goto st2
		case 39:
			goto st5
		case 45:
			goto tr3
		case 60:
			goto tr5
		case 62:
			goto tr5
		case 102:
			goto tr7
		case 116:
			goto tr8
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr4
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr6
			}
		default:
			goto tr6
		}
		goto tr0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 34:
			goto tr10
		case 92:
			goto tr11
		}
		goto tr9
	tr9:
//line parse_params.rl:20
		mark = p
		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line parse_params.go:358
		switch data[p] {
		case 34:
			goto tr13
		case 92:
			goto st4
		}
		goto st3
	tr11:
//line parse_params.rl:20
		mark = p
		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line parse_params.go:375
		switch data[p] {
		case 34:
			goto tr15
		case 92:
			goto st4
		}
		goto st3
	tr15:
//line NONE:1
		te = p + 1

//line parse_params.rl:33

		params[param] = data[mark:p]

		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line parse_params.go:397
		switch data[p] {
		case 34:
			goto tr13
		case 92:
			goto st4
		}
		goto st3
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 39:
			goto tr10
		case 92:
			goto tr17
		}
		goto tr16
	tr16:
//line parse_params.rl:20
		mark = p
		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line parse_params.go:426
		switch data[p] {
		case 39:
			goto tr13
		case 92:
			goto st7
		}
		goto st6
	tr17:
//line parse_params.rl:20
		mark = p
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line parse_params.go:443
		switch data[p] {
		case 39:
			goto tr20
		case 92:
			goto st7
		}
		goto st6
	tr20:
//line NONE:1
		te = p + 1

//line parse_params.rl:33

		params[param] = data[mark:p]

		goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line parse_params.go:465
		switch data[p] {
		case 39:
			goto tr13
		case 92:
			goto st7
		}
		goto st6
	tr3:
//line parse_params.rl:20
		mark = p
		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line parse_params.go:482
		if 48 <= data[p] && data[p] <= 57 {
			goto st16
		}
		goto tr0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if 48 <= data[p] && data[p] <= 57 {
			goto st16
		}
		goto tr33
	tr4:
//line parse_params.rl:20
		mark = p
		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line parse_params.go:505
		if data[p] == 45 {
			goto st16
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st17
		}
		goto tr34
	tr5:
//line parse_params.rl:20
		mark = p
		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line parse_params.go:522
		if 48 <= data[p] && data[p] <= 57 {
			goto st18
		}
		goto tr0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if 48 <= data[p] && data[p] <= 57 {
			goto st18
		}
		goto tr36
	tr6:
//line parse_params.rl:20
		mark = p
		goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line parse_params.go:545
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st19
			}
		case data[p] >= 65:
			goto st19
		}
		goto tr37
	tr7:
//line parse_params.rl:20
		mark = p
		goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line parse_params.go:564
		if data[p] == 97 {
			goto st21
		}
		switch {
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st19
			}
		case data[p] >= 65:
			goto st19
		}
		goto tr37
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if data[p] == 108 {
			goto st22
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st19
			}
		case data[p] >= 65:
			goto st19
		}
		goto tr37
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		if data[p] == 115 {
			goto st23
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st19
			}
		case data[p] >= 65:
			goto st19
		}
		goto tr37
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if data[p] == 101 {
			goto st24
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st19
			}
		case data[p] >= 65:
			goto st19
		}
		goto tr37
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st19
			}
		case data[p] >= 65:
			goto st19
		}
		goto tr43
	tr8:
//line parse_params.rl:20
		mark = p
		goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line parse_params.go:651
		if data[p] == 114 {
			goto st26
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st19
			}
		case data[p] >= 65:
			goto st19
		}
		goto tr37
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		if data[p] == 117 {
			goto st23
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st19
			}
		case data[p] >= 65:
			goto st19
		}
		goto tr37
	tr28:
//line NONE:1
		te = p + 1

//line parse_params.rl:20
		mark = p
		goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line parse_params.go:693
		switch data[p] {
		case 45:
			goto tr30
		case 58:
			goto tr31
		case 95:
			goto tr30
		case 116:
			goto tr45
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr30
			}
		case data[p] >= 65:
			goto tr30
		}
		goto tr29
	tr45:
//line NONE:1
		te = p + 1

		goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line parse_params.go:723
		switch data[p] {
		case 45:
			goto tr30
		case 58:
			goto tr31
		case 95:
			goto tr30
		case 116:
			goto tr46
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr30
			}
		case data[p] >= 65:
			goto tr30
		}
		goto tr29
	tr46:
//line NONE:1
		te = p + 1

		goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line parse_params.go:753
		switch data[p] {
		case 45:
			goto tr30
		case 58:
			goto tr31
		case 95:
			goto tr30
		case 112:
			goto tr47
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr30
			}
		case data[p] >= 65:
			goto tr30
		}
		goto tr29
	tr47:
//line NONE:1
		te = p + 1

		goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line parse_params.go:783
		switch data[p] {
		case 45:
			goto tr30
		case 58:
			goto tr31
		case 95:
			goto tr30
		case 115:
			goto tr48
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr30
			}
		case data[p] >= 65:
			goto tr30
		}
		goto tr29
	tr48:
//line NONE:1
		te = p + 1

		goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line parse_params.go:813
		switch data[p] {
		case 45:
			goto tr30
		case 58:
			goto tr49
		case 95:
			goto tr30
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr30
			}
		case data[p] >= 65:
			goto tr30
		}
		goto tr29
	tr49:
//line parse_params.rl:21

		if _, ok := params[param]; !ok && param != "" {
			params[param] = true
		}
		param = strings.ReplaceAll(data[mark:p], "-", "_")

		goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line parse_params.go:845
		switch data[p] {
		case 34:
			goto st2
		case 39:
			goto st5
		case 45:
			goto tr3
		case 47:
			goto st11
		case 60:
			goto tr5
		case 62:
			goto tr5
		case 102:
			goto tr7
		case 116:
			goto tr8
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr4
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr6
			}
		default:
			goto tr6
		}
		goto tr0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if data[p] == 47 {
			goto tr24
		}
		goto tr0
	st_out:
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof23:
		cs = 23
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof27:
		cs = 27
		goto _test_eof
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof30:
		cs = 30
		goto _test_eof
	_test_eof31:
		cs = 31
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 13:
				goto tr29
			case 1:
				goto tr0
			case 2:
				goto tr0
			case 3:
				goto tr0
			case 4:
				goto tr0
			case 14:
				goto tr32
			case 5:
				goto tr0
			case 6:
				goto tr0
			case 7:
				goto tr0
			case 15:
				goto tr32
			case 8:
				goto tr0
			case 16:
				goto tr33
			case 17:
				goto tr34
			case 9:
				goto tr0
			case 18:
				goto tr36
			case 19:
				goto tr37
			case 20:
				goto tr37
			case 21:
				goto tr37
			case 22:
				goto tr37
			case 23:
				goto tr37
			case 24:
				goto tr43
			case 25:
				goto tr37
			case 26:
				goto tr37
			case 27:
				goto tr29
			case 28:
				goto tr29
			case 29:
				goto tr29
			case 30:
				goto tr29
			case 31:
				goto tr29
			case 10:
				goto tr0
			case 11:
				goto tr0
			}
		}

	_out:
		{
		}
	}

//line parse_params.rl:62

	if cs < query_first_final {
		return nil, fmt.Errorf("Cannot parse: %d", p)
	}

	if _, ok := params[param]; !ok && param != "" {
		params[param] = true
	}

	ts, act = act, ts

	return NewParams(params), nil
}
