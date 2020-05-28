package parser;

import "strings"
import "fmt"
import "strconv"

%% machine query;
%% write data;

func ParseParams(data string) (*Params, error) {
    p, cs, pe, eof := 0, 0, len(data), len(data)

    var ts, te, act int

    mark   := 0
    params := map[string]interface{}{}
    param  := ""

    %%{
        action mark { mark = p }
        action param {
            if _, ok := params[param]; !ok {
                params[param] = true
            }
            param = strings.ReplaceAll(data[mark:p], "-", "_")
        }
        action range {
            params[param] = NewRangeFromString(data[mark:p])
        }
        action boolean {
            params[param] = data[mark:p] == "true"
        }
        action string {
            params[param] = data[mark+1:p-1]
        }
        action equality {
            params[param] = NewRangeFromString(data[mark:p])
        }
        action integer {
            params[param], _ = strconv.Atoi(data[mark:p])
        }

        param = (alpha | "-" | "_")+ >mark %param;
        range = (digit+ "-" | "-" digit+ | digit+ "-" digit+) >mark %range;
        equality = ((">" | "<") digit+) >mark %equality;
        boolean = ("true" | "false") >mark %boolean;
        string  = ( "'" ([^'] | /\\./)* "'" | '"' ([^"] | /\\./)* '"') >mark %string;
        integer = digit+ >mark %integer;
        value   = param (":" (integer | range | equality | boolean | string))?;
        main := |*
            value => {};
            space => {};
        *|;
        write init;
        write exec;
    }%%

    if cs < query_first_final {
        return nil, fmt.Errorf("Cannot parse: %d", p)
    }

    if _, ok := params[param]; !ok {
        params[param] = true
    }

    ts, act = act, ts

    return NewParams(params), nil
}
