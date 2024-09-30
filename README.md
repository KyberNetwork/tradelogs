## This is service to indexing tradelogs from event

### Re-generate mock file

First, you need to install `mockery`

- https://vektra.github.io/mockery/latest/installation/
- https://github.com/vektra/mockery

Then you use the `mockery` to generate mock files from interface

```
mockery --dir=<interface directory> --name=<interface name> --output=<output dir> --structname=<name of output struct> --filename=<output filename>
```

Example: generate `Storage` interface to a struct name `MockStorage`

```
mockery --dir=v2/pkg/storage/state --name=Storage --output=v2/mocks/ --structname=MockState --filename=State.go
```

For more information `mockery --help`