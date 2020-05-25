# frozen_string_literal: true

require 'spec_helper'
require 'ostruct'
require_relative '../lib/query'

RSpec.describe 'query' do
  context '#url' do
    let(:city) do
      OpenStruct.new(
        url: 'https://city.craigslist.org',
        nearby_cities: {
          'a' => '1',
          'b' => '2'
        }
      )
    end

    it 'generates a url for a city' do
      url = p('q:"jeep truck"').url(city: city)
      expect(url).to eq 'https://city.craigslist.org/search/sss?query=jeep%20truck'
    end

    it 'includes nearby cities' do
      url = p('q:"jeep truck" include-nearby:true').url(city: city)
      expect(url).to eq 'https://city.craigslist.org/search/sss?nearbyArea=1&nearbyArea=2&query=jeep%20truck&searchNearby=2'
    end

    it 'can bundle duplicates' do
      url = p('bundle-duplicates').url(city: city)
      expect(url).to eq 'https://city.craigslist.org/search/sss?bundleDuplicates=1'

      url = p('has-image').url(city: city)
      expect(url).to eq 'https://city.craigslist.org/search/sss?hasPic=1'
    end

    it 'includes ranges' do
      url = p('price:1000-2000 auto-year:1980-2000').url(city: city)
      expect(url).to eq 'https://city.craigslist.org/search/sss?max_auto_year=2000&max_price=2000&min_auto_year=1980&min_price=1000'
    end
  end
end
