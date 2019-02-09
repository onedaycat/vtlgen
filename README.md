# Vtlgen

## Installation

```
go get -u github.com/onedaycat/vtlgen/...
```

## Usage

```
vtlgen -dir=<mappingtemplates folder> -out=<specific file generate>
```

## Example

```
vtlgen -dir=testdata/mapping-templates -out=testdata/mappingTemplates.yml
```

## MappingTemplates folder structure

```
mapping-templates
  req.vtl // default req
  res.vtl // default res
  <datasource>
    <type>
      <field>
        req.vtl
        res.vtl
```