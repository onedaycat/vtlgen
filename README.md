# Vtlgen

## Installation

```
go get -u github.com/onedaycat/vtlgen/...
```

## Usage

```
vtlgen template <flag: -p, --path> <specific folder contain mapping-templates folder>
vtlgen datasource <flag: -f, --file> <specific file datasource_generate.yml>
```

## Example

```shall

# Template

## Currenct directory
vtlgen template -p ./
vtlgen template

## Not currenct directory
vtlgen template -p testdata


# Datasource

## Currenct directory
vtlgen datasource -f datasource_generate.yml
vtlgen datasource

## Not currenct directory
vtlgen datasource -f testdata/datasource_generate.yml

```

## MappingTemplates folder structure

```
mapping-templates
  - datasource_generate.yml
  - resolver.yml   // result generated
  - datasource.yml // result generated
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
```

## datasource_generate.yml structure

```yml
accountId: FAKE12345678
serviceRoleArn: arn:aws:iam::${accountId}:role/${self:service}-${self:provider.stage}
noneDatasource: local
lambdaDatasources:
- name: accountQuery
  service: sel-account-qry
  handler: qry
- name: accountMutation
  service: sel-account-cmd
  handler: cmd
- name: storeMutation
  service: sel-store-cmd
  handler: cmd
  version: VVVV
  serviceRoleArn: XXXX
```
