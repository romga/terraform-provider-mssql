---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "mssql_schema Data Source - terraform-provider-mssql"
subcategory: ""
description: |-
  Retrieves information about DB schema.
---

# mssql_schema (Data Source)

Retrieves information about DB schema.

## Example Usage

```terraform
data "mssql_database" "example" {
  name = "example"
}

data "mssql_schema" "by_name" {
  database_id = data.mssql_database.example.id
  name        = "dbo"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `database_id` (String) ID of database. Can be retrieved using `mssql_database` or `SELECT DB_ID('<db_name>')`.
- `id` (String) `<database_id>/<schema_id>`. Schema ID can be retrieved using `SELECT SCHEMA_ID('<schema_name>')`. Either `id` or `name` must be provided.
- `name` (String) Schema name. Either `id` or `name` must be provided.

### Read-Only

- `owner_id` (String) ID of database role or user owning this schema. Can be retrieved using `mssql_database_role`, `mssql_sql_user`, `mssql_azuread_user` or `mssql_azuread_service_principal`

