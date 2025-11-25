# Development

```sh
mkdir module
cd module
dagger init --sdk=go --source=.
```

```sh
mkdir tests
cd tests
dagger init --sdk=go --source=.
dagger install ..
dagger update
```
