# frozen_string_literal: true

# line 1 "lib/parse.rl"

# line 2 "lib/parse.rl"

require_relative 'query'

class Parser
  # line 11 "lib/parse.rb"
  class << self
    attr_accessor :_query_trans_keys
    private :_query_trans_keys, :_query_trans_keys=
  end
  self._query_trans_keys = [
    0, 0, 34, 116, 34, 92,
    34, 92, 39, 92, 39,
    92, 48, 57, 48, 57,
    97, 97, 108, 108, 115, 115,
    101, 101, 114, 114, 117,
    117, 9, 122, 45, 122,
    0, 0, 34, 92, 39, 92,
    48, 57, 45, 57, 48,
    57, 0, 0, 0
  ]

  class << self
    attr_accessor :_query_key_spans
    private :_query_key_spans, :_query_key_spans=
  end
  self._query_key_spans = [
    0, 83, 59, 59, 54, 54, 10, 10,
    1, 1, 1, 1, 1, 1, 114, 78,
    0, 59, 54, 10, 13, 10, 0
  ]

  class << self
    attr_accessor :_query_index_offsets
    private :_query_index_offsets, :_query_index_offsets=
  end
  self._query_index_offsets = [
    0, 0, 84, 144, 204, 259, 314, 325,
    336, 338, 340, 342, 344, 346, 348, 463,
    542, 543, 603, 658, 669, 683, 694
  ]

  class << self
    attr_accessor :_query_indicies
    private :_query_indicies, :_query_indicies=
  end
  self._query_indicies = [
    1, 0, 0, 0, 0, 2, 0, 0,
    0, 0, 0, 3, 0, 0, 4, 4,
    4, 4, 4, 4, 4, 4, 4, 4,
    0, 0, 5, 0, 5, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 6, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 7, 0, 9, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 10, 8,
    11, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 10, 8, 9, 12, 12, 12,
    12, 12, 12, 12, 12, 12, 12, 12,
    12, 12, 12, 12, 12, 12, 12, 12,
    12, 12, 12, 12, 12, 12, 12, 12,
    12, 12, 12, 12, 12, 12, 12, 12,
    12, 12, 12, 12, 12, 12, 12, 12,
    12, 12, 12, 12, 12, 12, 12, 12,
    12, 13, 12, 14, 12, 12, 12, 12,
    12, 12, 12, 12, 12, 12, 12, 12,
    12, 12, 12, 12, 12, 12, 12, 12,
    12, 12, 12, 12, 12, 12, 12, 12,
    12, 12, 12, 12, 12, 12, 12, 12,
    12, 12, 12, 12, 12, 12, 12, 12,
    12, 12, 12, 12, 12, 12, 12, 12,
    13, 12, 15, 15, 15, 15, 15, 15,
    15, 15, 15, 15, 0, 16, 16, 16,
    16, 16, 16, 16, 16, 16, 16, 0,
    17, 0, 18, 0, 19, 0, 20, 0,
    21, 0, 19, 0, 22, 22, 22, 22,
    22, 23, 23, 23, 23, 23, 23, 23,
    23, 23, 23, 23, 23, 23, 23, 23,
    23, 23, 23, 22, 23, 23, 23, 23,
    23, 23, 23, 23, 23, 23, 23, 23,
    24, 23, 23, 23, 23, 23, 23, 23,
    23, 23, 23, 23, 23, 23, 23, 23,
    23, 23, 23, 23, 24, 24, 24, 24,
    24, 24, 24, 24, 24, 24, 24, 24,
    24, 24, 24, 24, 24, 24, 24, 24,
    24, 24, 24, 24, 24, 24, 23, 23,
    23, 23, 24, 23, 24, 24, 24, 24,
    24, 24, 24, 24, 24, 24, 24, 24,
    24, 24, 24, 24, 24, 24, 24, 24,
    24, 24, 24, 24, 24, 24, 23, 26,
    25, 25, 25, 25, 25, 25, 25, 25,
    25, 25, 25, 25, 27, 25, 25, 25,
    25, 25, 25, 26, 26, 26, 26, 26,
    26, 26, 26, 26, 26, 26, 26, 26,
    26, 26, 26, 26, 26, 26, 26, 26,
    26, 26, 26, 26, 26, 25, 25, 25,
    25, 26, 25, 26, 26, 26, 26, 26,
    26, 26, 26, 26, 26, 26, 26, 26,
    26, 26, 26, 26, 26, 26, 26, 26,
    26, 26, 26, 26, 26, 25, 28, 9,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 8, 8, 8, 8, 8, 8, 8,
    8, 10, 8, 9, 12, 12, 12, 12,
    12, 12, 12, 12, 12, 12, 12, 12,
    12, 12, 12, 12, 12, 12, 12, 12,
    12, 12, 12, 12, 12, 12, 12, 12,
    12, 12, 12, 12, 12, 12, 12, 12,
    12, 12, 12, 12, 12, 12, 12, 12,
    12, 12, 12, 12, 12, 12, 12, 12,
    13, 12, 15, 15, 15, 15, 15, 15,
    15, 15, 15, 15, 29, 30, 29, 29,
    31, 31, 31, 31, 31, 31, 31, 31,
    31, 31, 29, 16, 16, 16, 16, 16,
    16, 16, 16, 16, 16, 32, 33, 0
  ]

  class << self
    attr_accessor :_query_trans_targs
    private :_query_trans_targs, :_query_trans_targs=
  end
  self._query_trans_targs = [
    14, 2, 4, 6, 20, 7, 8, 12,
    2, 16, 3, 17, 4, 5, 18, 19,
    21, 9, 10, 11, 22, 13, 14, 0,
    15, 14, 15, 1, 14, 14, 6, 20,
    14, 14
  ]

  class << self
    attr_accessor :_query_trans_actions
    private :_query_trans_actions, :_query_trans_actions=
  end
  self._query_trans_actions = [
    1, 2, 2, 2, 3, 2, 2, 2,
    0, 0, 0, 4, 0, 0, 4, 0,
    0, 0, 0, 0, 0, 0, 7, 0,
    3, 8, 4, 9, 10, 11, 0, 4,
    12, 13
  ]

  class << self
    attr_accessor :_query_to_state_actions
    private :_query_to_state_actions, :_query_to_state_actions=
  end
  self._query_to_state_actions = [
    0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 5, 0,
    0, 0, 0, 0, 0, 0, 0
  ]

  class << self
    attr_accessor :_query_from_state_actions
    private :_query_from_state_actions, :_query_from_state_actions=
  end
  self._query_from_state_actions = [
    0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 6, 0,
    0, 0, 0, 0, 0, 0, 0
  ]

  class << self
    attr_accessor :_query_eof_trans
    private :_query_eof_trans, :_query_eof_trans=
  end
  self._query_eof_trans = [
    0, 1, 1, 1, 1, 1, 1, 1,
    1, 1, 1, 1, 1, 1, 0, 26,
    29, 29, 29, 30, 30, 33, 34
  ]

  class << self
    attr_accessor :query_start
  end
  self.query_start = 14
  class << self
    attr_accessor :query_first_final
  end
  self.query_first_final = 14
  class << self
    attr_accessor :query_error
  end
  self.query_error = 0

  class << self
    attr_accessor :query_en_main
  end
  self.query_en_main = 14

  # line 7 "lib/parse.rl"

  def self.parse(data)
    p = 0
    cs = 0
    pe = data.length
    eof = data.length

    mark = 0
    params = {}
    param = nil

    # line 224 "lib/parse.rb"
    begin
      p ||= 0
      pe ||= data.length
      cs = query_start
      ts = nil
      te = nil
      act = 0
    end

    # line 234 "lib/parse.rb"
    begin
        testEof = false
        _slen, _trans, _keys, _inds, _acts, _nacts = nil
        _goto_level = 0
        _resume = 10
        _eof_trans = 15
        _again = 20
        _test_eof = 30
        _out = 40
        loop do
          if _goto_level <= 0
            if p == pe
              _goto_level = _test_eof
              next
            end
            if cs == 0
              _goto_level = _out
              next
            end
          end
          if _goto_level <= _resume
            case _query_from_state_actions[cs]
            when 6
              # line 1 "NONE"
              begin
                ts = p
              end
              # line 262 "lib/parse.rb"
            end
            _keys = cs << 1
            _inds = _query_index_offsets[cs]
            _slen = _query_key_spans[cs]
            _wide = data[p].ord
            _trans = if _slen > 0 &&
                        _query_trans_keys[_keys] <= _wide &&
                        _wide <= _query_trans_keys[_keys + 1]

                       _query_indicies[_inds + _wide - _query_trans_keys[_keys]]
                     else
                       _query_indicies[_inds + _slen]
               end
          end
          if _goto_level <= _eof_trans
            cs = _query_trans_targs[_trans]
            if _query_trans_actions[_trans] != 0
              case _query_trans_actions[_trans]
              when 2
                # line 15 "lib/parse.rl"
                begin
             mark = p	end
              when 9
                # line 16 "lib/parse.rl"
                begin
                              params[param] = true if param
                              param = data[mark...p].gsub('-', '_')
                            end
              when 4
                # line 1 "NONE"
                begin
                  te = p + 1
                end
              when 7
                # line 52 "lib/parse.rl"
                begin
                  te = p + 1
                end
              when 1
                # line 51 "lib/parse.rl"
                begin
                  begin p = te - 1; end
                end
              when 8 then
                # line 16 "lib/parse.rl"
                begin
                              params[param] = true if param
                              param = data[mark...p].gsub('-', '_')
                            end
                # line 51 "lib/parse.rl"
                begin
            te = p
            p -= 1;	end
              when 11 then
                # line 22 "lib/parse.rl"
                begin
                              left, right = data[mark..p].split('-').map(&:to_i)
                              params[param] = Range.new(left, right)
                              param = nil
                            end
                # line 51 "lib/parse.rl"
                begin
            te = p
            p -= 1;	end
              when 13 then
                # line 27 "lib/parse.rl"
                begin
                              params[param] = data[mark..p] == 'true'
                              param = nil
                            end
                # line 51 "lib/parse.rl"
                begin
            te = p
            p -= 1;	end
              when 10 then
                # line 31 "lib/parse.rl"
                begin
                              params[param] = data[mark + 1..p - 2]
                              param = nil
                            end
                # line 51 "lib/parse.rl"
                begin
            te = p
            p -= 1;	end
              when 12 then
                # line 35 "lib/parse.rl"
                begin
                              left = 0
                              right = data[mark + 1..p].to_i
                              if data[mark] == '>'
                                left = right
                                right = nil
                          end
                              params[param] = Range.new(left, right)
                              param = nil
                            end
                # line 51 "lib/parse.rl"
                begin
            te = p
            p -= 1;	end
              when 3 then
                # line 1 "NONE"
                begin
                  te = p + 1
                end
                # line 15 "lib/parse.rl"
                begin
             mark = p	end
                # line 379 "lib/parse.rb"
              end
            end
          end
          if _goto_level <= _again
            case _query_to_state_actions[cs]
            when 5
              # line 1 "NONE"
              begin
          ts = nil;	end
              # line 389 "lib/parse.rb"
            end

            if cs == 0
              _goto_level = _out
              next
            end
            p += 1
            if p != pe
              _goto_level = _resume
              next
            end
          end
          if _goto_level <= _test_eof
            if p == eof
              if _query_eof_trans[cs] > 0
                _trans = _query_eof_trans[cs] - 1
                _goto_level = _eof_trans
                next
              end
            end

          end
          break if _goto_level <= _out
        end
      end

    # line 56 "lib/parse.rl"

    raise "Cannot parse: #{p}" if cs < query_first_final

    params[param] = true if param

    Query.new(params)
  end
end
