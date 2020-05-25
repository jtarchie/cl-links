# frozen_string_literal: true

require 'capybara/dsl'
require_relative 'memoize'

Capybara.default_driver = :selenium_chrome_headless

CAPTURED_DIR = File.join(__dir__, '..', 'captured')

# Allows a basic interaction with Craigslist
# mainly for getting information about cities
class Craigslist
  include Capybara::DSL
  include JSONMemoize

  def countries
    @countries ||= capture(File.join(CAPTURED_DIR, 'countries.json')) do
      visit 'https://www.craigslist.org/about/sites'
      page.execute_script(<<~JAVASCRIPT)
        return (function() {
            var countries = {};
            $('h1').each(function() {
                var $el = $(this);

                var regions = {};
                $el.find('+ .colmask h4').each(function() {
                    var $el = $(this);

                    var cities = {};
                    $el.find('+ ul li a').each(function() {
                        var $el = $(this);
                        cities[$el.text().trim()] = $el.attr('href');
                    });
                
                    regions[$el.text().trim()] = cities;
                });

                countries[$el.text().trim()] = regions;
            });
            return countries
        })()
      JAVASCRIPT
    end
  end

  City = Struct.new(:name, :region_name, :country_name, :url, keyword_init: true) do
    include Capybara::DSL
    include JSONMemoize

    def nearby_cities
      capture(File.join(CAPTURED_DIR, country_name, region_name, name, 'nearby_cities.json')) do
        visit File.join(url, 'search')

        page.execute_script(<<~JAVASCRIPT)
          return (function() {
              nearby = {};
              $('label.nearby').each(function() {
                  nearby[$(this).text().trim()] = $(this).find('input').val()
              });
              return nearby;
          })()
        JAVASCRIPT
      end
    end
  end

  Cities = Struct.new(:cities) do
    def by_country(name:)
      Cities.new(cities.select { |c| c.country_name == name })
    end

    def by_region(name:)
      Cities.new(cities.select { |c| c.region_name == name })
    end

    def each(&block)
      cities.shuffle.each(&block)
    end
  end

  def cities
    all_cities = []
    countries.each do |country_name, regions|
      regions.each do |region_name, cities|
        cities.each do |city_name, city_url|
          all_cities << City.new(
            name: city_name,
            region_name: region_name,
            country_name: country_name,
            url: city_url
          )
        end
      end
    end
    Cities.new(all_cities)
  end
end
