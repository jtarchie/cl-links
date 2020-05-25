# frozen_string_literal: true

require_relative 'lib/craigslist'
require 'json'

cl = Craigslist.new
us_cities = cl.cities.by_country(name: 'US')
payload = us_cities.map do |city|
  {
    name: city.name,
    region_name: city.region_name,
    country_name: city.country_name,
    nearby_cities: city.nearby_cities,
    url: city.url
  }
end

File.write('us.json', JSON.pretty_generate(payload))
