# Build instructions

```sh
bundle
ragel -R -e -F1 lib/parse.rl
rubocop -a
rspec
ruby build.rb # generates a list of only US cities at the moment
```

# Run

```sh
ruby generate.rb
# load in browser
```
