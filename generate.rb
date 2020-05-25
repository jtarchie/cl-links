# frozen_string_literal: true

require 'json'
require 'ostruct'
require 'erb'
require_relative 'lib/parse'

include ERB::Util

cities = JSON.parse(File.read('us.json')).collect { |c| OpenStruct.new(c) }

queries = [
  'q:"jeep truck" include_nearby has-image bundle-duplicates auto-year:<1985',
  'q:"jeep j10" include_nearby has-image bundle-duplicates auto-year:<1985',
  'q:"jeep j20" include_nearby has-image bundle-duplicates auto-year:<1985'
].map { |q| Parser.parse(q) }

top_cities = cities
             .sort_by { |c| c.nearby_cities.length }
             .reverse
             .uniq(&:region_name)
             .take(50)
             .sort_by(&:region_name)

template = <<~ERB
  <html>
      <head>
          <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/mini.css/3.0.1/mini-default.min.css">
          <meta name="viewport" content="width=device-width, initial-scale=1">
      </head>
      <body>
      <div class="container">
          <% top_cities.each do |city| %>
              <h1><%= city.country_name %> / <%= city.region_name %> / <%= city.name %></h1>
              <ol>
                  <% queries.each do |query| %>
                      <li><a target="_blank" href="<%= query.url(city: city) %>"><%= h query.to_s %></a></li>
                  <% end %>
              </ol>
          <% end %>
      </div>
      </body>
  </html>
ERB
puts ERB.new(template).result(binding)
