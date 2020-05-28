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
