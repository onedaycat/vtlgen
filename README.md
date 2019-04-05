# Vtlgen

## Installation

```
go get -u github.com/onedaycat/vtlgen/...
```

## Usage

```
vtlgen -dir= <select directory that have datasource_generate.yml and folder mappingTemplates>
or
vtlgen
```

## Example

```
// specific directory
vtlgen -dir=testdata

// default current directory
vtlgen
```

## MappingTemplates folder structure

```
// folder & file structure

mapping-templates
  - datasource_generate.yml
  - resolver
      <type.field>
        req.vtl
        res.vtl
        config.yml
      <type.field>
        before.vtl
        after.vtl
        config.yml
  - function
      <func_name>
        req.vtl
        res.vtl
        config.yml

// result
- resolver.yml
- datasource.yml

```