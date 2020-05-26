# frozen_string_literal: true

require 'json'
require 'ostruct'
require 'sinatra'
require_relative 'lib/parse'

include ERB::Util

cities     = JSON.parse(File.read('us.json')).collect { |c| OpenStruct.new(c) }
TOP_CITIES = cities
             .sort_by { |c| c.nearby_cities.length }
             .reverse
             .uniq(&:region_name)
             .take(50)
             .sort_by(&:region_name)

get '/' do
  @parsed_query = Parser.parse(@query) if @query = params[:query]

  @example_query = 'q:"jeep truck" include_nearby has-image bundle-duplicates auto-year:<1985'
  @top_cities    = TOP_CITIES

  erb :queries
end

__END__

@@ queries
<html>
  <head>
      <script src="https://cdnjs.cloudflare.com/ajax/libs/turbolinks/5.2.0/turbolinks.js"></script>
      <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/mini.css/3.0.1/mini-default.min.css">
      <meta name="viewport" content="width=device-width, initial-scale=1">
      <style>
        form input {
            width: 100%;
        }
      </style>
  </head>
  <body>
      <div class="container">
          <div class="row">
              <div class="col-sm-12 col-md-12 col-lg-12">
                  <form action="/">
                      <input type="text" value="<%= h @query %>" placeholder="q:'classic guitar' year:<1995 include-pics" name="query">
                      <p>Example: <a href="/?query=<%= u 'q:"jeep truck" include_nearby has-image bundle-duplicates auto-year:<1985' %>">q:"jeep truck" include_nearby has-image bundle-duplicates auto-year:<1985</a></p>
                  </form>
              </div>
          </div>
          <div class="row">
              <div class="col-sm-12 col-md-12 col-lg-12">
              <% if @query %>
                  <ul>
                  <% @top_cities.each do |city| %>
                    <li><a target="_blank" href="<%= @parsed_query.url(city: city) %>"><%= city.country_name %> / <%= city.region_name %> / <%= city.name %></a></li>
                  <% end %>
                  </ul>
              <% end %>
              </div>
          </div>
      </div>
  </body>
</html>
