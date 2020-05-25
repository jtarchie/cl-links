# Build instructions

```sh
bundle
ragel -R -e -F1 lib/parse.rl
rubocop -a
rspec
```
