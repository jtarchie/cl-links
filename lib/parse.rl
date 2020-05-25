%% machine query;

require_relative 'query'

class Parser
    %% write data;

    def self.parse(data)
        p, cs, pe, eof = 0, 0, data.length, data.length

        mark = 0
        params, param = {}, nil

        %%{
            action mark { mark = p }
            action param {
                if param
                    params[param] = true
                end
                param = data[mark...p].gsub('-', '_')
            }
            action range {
                left, right = data[mark..p].split('-').map(&:to_i)
                params[param] = Range.new(left, right)
                param = nil
            }
            action boolean {
                params[param] = data[mark..p] == 'true'
                param = nil
            }
            action string {
                params[param] = data[mark+1..p-2]
                param = nil
            }
            action equality {
                left, right = 0, data[mark+1..p].to_i
                if data[mark] == ">"
                    left, right = right, nil
                end
                params[param] = Range.new(left, right)
                param = nil
            }

            param = (alpha | "-" | "_")+ >mark %param;
            range = (digit+ | "-" digit+ | digit+ "-" digit+) >mark %range;
            equality = ((">" | "<") digit+) >mark %equality;
            boolean = ("true" | "false") >mark %boolean;
            string  = ( "'" ([^'] | /\\./)* "'" | '"' ([^"] | /\\./)* '"') >mark %string;
            value   = param (":" (range | equality | boolean | string))?;
            main := |*
                value => {};
                space => {};
            *|;
            write init;
            write exec;
        }%%

        if cs < query_first_final
            raise "Cannot parse: #{p}"
        end

        if param
            params[param] = true
        end

        return Query.new(params)
    end
end
