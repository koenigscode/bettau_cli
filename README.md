# bettaU - AI Language Learning CLI

bettaU is a simple CLI that uses AI to help you learn a new language.

It is more of a proof of concept than a ready-to-use tool, but if you're not looking for much, it works.

You provide it with sentences in English - as a yaml file - and you're prompted to translate them into the language you're learning.

An example yaml file can look like this:

```yaml
name: Swedish
contents:
  - She reads books every night.
  - He plays guitar after school.
```

`name` should be the name of the language you're learning - it's used by ChatGPT. Don't just set it to anything.

To learn a deck - which is simply a yaml file, as shown above - run

```bash
export TOKEN=<your openai token>

go run bettau.go learn <path to deck> # e.g. test-decks/swedish.yaml
```
