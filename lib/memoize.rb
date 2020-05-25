# frozen_string_literal: true

require 'json'
require 'fileutils'

# Populates and load JSON file
# if it already exists, it will load from that
module JSONMemoize
  def capture(filename, &block)
    if File.exist?(filename)
      JSON.parse(File.read(filename))
    else
      results = block.call
      FileUtils.mkdir_p(File.dirname(filename))
      File.write(filename, JSON.pretty_generate(results))
      results
    end
  end
end
