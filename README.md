# Build instructions

```sh
ragel -R -e -F1 lib/parse.rl;
rubocop -a
rspec
```
