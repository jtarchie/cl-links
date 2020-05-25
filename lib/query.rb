# frozen_string_literal: true

require 'ostruct'
require 'addressable'

class Query < OpenStruct
  def url(city:)
    uri = Addressable::URI.parse(city.url)
    qs = {}

    qs['query'] = q if q
    qs['bundleDuplicates'] = 1 if bundle_duplicates
    qs['hasPic'] = 1 if has_image

    if include_nearby && !city.nearby_cities.empty?
      qs['searchNearby'] = 2
      qs['nearbyArea'] = city.nearby_cities.values
    end

    to_h.each do |key, value|
      if value.is_a?(Range)
        qs["min_#{key}"] = value.min if value.min
        qs["max_#{key}"] = value.max if value.max
      end
    end

    uri.query_values = qs
    uri.path = '/search/sss'
    uri.to_s
  end
end
