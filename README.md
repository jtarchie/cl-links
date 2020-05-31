# Build instructions

```sh
cd collect
bundle install
rubocop -a
rspec
ruby run.rb # generates a list of only US cities at the moment
```

# Run

```sh
heroku container:push web
heroku container:release web
```

# Usage

1. Go to craigslist, do a search, and copy the URL of the search result page.

   For example, click on "cars+trucks, search for "ranger, the URL is "https://denver.craigslist.org/search/cta?query=ranger".
   
   Goto https://yard-search.herokuapp.com/, copy the URL into the search field, and hit enter.
   
   The page will update with a query (`category:"cta" city:"denver" q:"ranger"`).
   
   The URL has been translated into a query.
   
   Remove the `city:"denver"`, and it shows all cities across craigslist.
   
2. Use the undocumented query language that only JT knows.
   
   The query consists of query language, a syntax for key-value pairs (`key:value`). The supported types are strings (`key:"some string"`, `key:'some string'`, or `key:single-world`), numbers (`key:1000`), ranges (`key:1000-2000`, `key:>2000`, or `key:<2000`), and boolean (`key` same as `key:true` or `key:false`).
   
   Keys that can appear in a query:
   * `q` (_string_) is the keyword search query that will narrow down results on Craigslist. For example, to search trucks use `q:"trucks"`.
   * `city`, `region`, `country` (_string_) can narrow down the area the links are search for. For example, to search USA use `country:"US"`.
   * `price` (_range_) can narrow down the price being search. For example, looking for less than $2000 use `price:<2000`.
   * `include_nearby` (_boolean_) will have the search check nearby cities. For example, Denver will include Boulder, Colorado Springs, etc.
   * `has_image` (_boolean_) will only allow results that have images.
   * `bundle_duplicates` (_boolean_) will ensure that similar postings are bundled together.
   * `top:nearby` (_string_) will filter results by the top result of some type. `nearby` will have filter cities, so only cities in a region with the most nearby cities are shown.

With the query, you now have a set of links to just open in tabs.
