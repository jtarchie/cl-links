# frozen_string_literal: true

require_relative 'lib/craigslist'
require 'json'

cl = Craigslist.new
payload = cl.cities.map do |city|
  print "city: #{city.name} - "
  v = {
    name: city.name,
    region_name: city.region_name,
    country_name: city.country_name,
    nearby_cities: city.nearby_cities,
    url: city.url
  }
  puts 'done'
  v
end

File.write(File.join(DATA_DIR, 'processed.json'), JSON.pretty_generate(payload))
