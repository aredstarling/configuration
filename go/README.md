# configuration/go

Provides a go client for handling configuration information.

## Usage

Add it to Gopkg.toml

```toml
[[constraint]]
  branch = "master"
  name = "gitlab.com/sellernomics/configuration"
  source = "https://<GITHUB_TOKEN>@gitlab.com/sellernomics/configuration.git"
```

## Running Tests

```bash
make env=test spec
```
