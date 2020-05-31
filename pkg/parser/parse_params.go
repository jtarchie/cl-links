
//line parse_params.rl:1
package parser;

import "strings"
import "fmt"
import "strconv"


//line parse_params.rl:8

//line parse_params.go:13
const query_start int = 10
const query_first_final int = 10
const query_error int = 0

const query_en_main int = 10


//line parse_params.rl:9

func ParseParams(data string) (*Params, error) {
    p, cs, pe, eof := 0, 0, len(data), len(data)

    var ts, te, act int

    mark   := 0
    params := map[string]interface{}{}
    param  := ""

    
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
	case 10:
		goto st_case_10
	case 0:
		goto st_case_0
	case 11:
		goto st_case_11
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 12:
		goto st_case_12
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 13:
		goto st_case_13
	case 8:
		goto st_case_8
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 9:
		goto st_case_9
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
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
	}
	goto st_out
tr0:
//line parse_params.rl:54
p = (te) - 1

	goto st10
tr10:
//line parse_params.rl:20
 mark = p 
//line parse_params.rl:33

            params[param] = data[mark:p]
        
//line parse_params.rl:54
te = p+1

	goto st10
tr13:
//line parse_params.rl:33

            params[param] = data[mark:p]
        
//line parse_params.rl:54
te = p+1

	goto st10
tr23:
//line parse_params.rl:55
te = p+1

	goto st10
tr26:
//line parse_params.rl:21

            if _, ok := params[param]; !ok && param != "" {
                params[param] = true
            }
            param = strings.ReplaceAll(data[mark:p], "-", "_")
        
//line parse_params.rl:54
te = p
p--

	goto st10
tr29:
//line parse_params.rl:54
te = p
p--

	goto st10
tr30:
//line parse_params.rl:27

            params[param] = NewRangeFromString(data[mark:p])
        
//line parse_params.rl:54
te = p
p--

	goto st10
tr31:
//line parse_params.rl:39

            params[param], _ = strconv.Atoi(data[mark:p])
        
//line parse_params.rl:54
te = p
p--

	goto st10
tr33:
//line parse_params.rl:36

            params[param] = NewRangeFromString(data[mark:p])
        
//line parse_params.rl:54
te = p
p--

	goto st10
tr34:
//line parse_params.rl:33

            params[param] = data[mark:p]
        
//line parse_params.rl:54
te = p
p--

	goto st10
tr40:
//line parse_params.rl:30

            params[param] = data[mark:p] == "true"
        
//line parse_params.rl:54
te = p
p--

	goto st10
	st10:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line NONE:1
ts = p

//line parse_params.go:209
		switch data[p] {
		case 32:
			goto tr23
		case 45:
			goto tr25
		case 95:
			goto tr25
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr23
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
tr25:
//line NONE:1
te = p+1

//line parse_params.rl:20
 mark = p 
	goto st11
tr27:
//line NONE:1
te = p+1

	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line parse_params.go:252
		switch data[p] {
		case 45:
			goto tr27
		case 58:
			goto tr28
		case 95:
			goto tr27
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr27
			}
		case data[p] >= 65:
			goto tr27
		}
		goto tr26
tr28:
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
//line parse_params.go:284
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
//line parse_params.go:335
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
//line parse_params.go:352
		switch data[p] {
		case 34:
			goto tr15
		case 92:
			goto st4
		}
		goto st3
tr15:
//line NONE:1
te = p+1

//line parse_params.rl:33

            params[param] = data[mark:p]
        
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line parse_params.go:374
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
//line parse_params.go:403
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
//line parse_params.go:420
		switch data[p] {
		case 39:
			goto tr20
		case 92:
			goto st7
		}
		goto st6
tr20:
//line NONE:1
te = p+1

//line parse_params.rl:33

            params[param] = data[mark:p]
        
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line parse_params.go:442
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
//line parse_params.go:459
		if 48 <= data[p] && data[p] <= 57 {
			goto st14
		}
		goto tr0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if 48 <= data[p] && data[p] <= 57 {
			goto st14
		}
		goto tr30
tr4:
//line parse_params.rl:20
 mark = p 
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line parse_params.go:482
		if data[p] == 45 {
			goto st14
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st15
		}
		goto tr31
tr5:
//line parse_params.rl:20
 mark = p 
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line parse_params.go:499
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
tr6:
//line parse_params.rl:20
 mark = p 
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line parse_params.go:522
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st17
			}
		case data[p] >= 65:
			goto st17
		}
		goto tr34
tr7:
//line parse_params.rl:20
 mark = p 
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line parse_params.go:541
		if data[p] == 97 {
			goto st19
		}
		switch {
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st17
			}
		case data[p] >= 65:
			goto st17
		}
		goto tr34
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if data[p] == 108 {
			goto st20
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st17
			}
		case data[p] >= 65:
			goto st17
		}
		goto tr34
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		if data[p] == 115 {
			goto st21
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st17
			}
		case data[p] >= 65:
			goto st17
		}
		goto tr34
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if data[p] == 101 {
			goto st22
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st17
			}
		case data[p] >= 65:
			goto st17
		}
		goto tr34
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st17
			}
		case data[p] >= 65:
			goto st17
		}
		goto tr40
tr8:
//line parse_params.rl:20
 mark = p 
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line parse_params.go:628
		if data[p] == 114 {
			goto st24
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st17
			}
		case data[p] >= 65:
			goto st17
		}
		goto tr34
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if data[p] == 117 {
			goto st21
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st17
			}
		case data[p] >= 65:
			goto st17
		}
		goto tr34
	st_out:
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 11:
			goto tr26
		case 1:
			goto tr0
		case 2:
			goto tr0
		case 3:
			goto tr0
		case 4:
			goto tr0
		case 12:
			goto tr29
		case 5:
			goto tr0
		case 6:
			goto tr0
		case 7:
			goto tr0
		case 13:
			goto tr29
		case 8:
			goto tr0
		case 14:
			goto tr30
		case 15:
			goto tr31
		case 9:
			goto tr0
		case 16:
			goto tr33
		case 17:
			goto tr34
		case 18:
			goto tr34
		case 19:
			goto tr34
		case 20:
			goto tr34
		case 21:
			goto tr34
		case 22:
			goto tr40
		case 23:
			goto tr34
		case 24:
			goto tr34
		}
	}

	_out: {}
	}

//line parse_params.rl:59


    if cs < query_first_final {
        return nil, fmt.Errorf("Cannot parse: %d", p)
    }

    if _, ok := params[param]; !ok && param != "" {
        params[param] = true
    }

    ts, act = act, ts

    return NewParams(params), nil
}
