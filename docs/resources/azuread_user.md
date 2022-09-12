---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "mssql_azuread_user Resource - terraform-provider-mssql"
subcategory: ""
description: |-
  Managed database-level user mapped to Azure AD identity (user or group).
  -> Note When using this resource, Azure SQL server managed identity does not need any AzureAD role assignments https://docs.microsoft.com/en-us/azure/azure-sql/database/authentication-aad-service-principal?view=azuresql.
---

# mssql_azuread_user (Resource)

Managed database-level user mapped to Azure AD identity (user or group).

-> **Note** When using this resource, Azure SQL server managed identity does not need any [AzureAD role assignments](https://docs.microsoft.com/en-us/azure/azure-sql/database/authentication-aad-service-principal?view=azuresql).

## Example Usage

```terraform
data "mssql_database" "example" {
  name = "example"
}

data "azuread_user" "example" {
  user_principal_name = "user@example.com"
}

resource "mssql_azuread_user" "example" {
  name           = "example"
  database_id    = data.mssql_database.example.id
  user_object_id = data.azuread_user.example.object_id
}

output "user_id" {
  value = mssql_azuread_user.example.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `database_id` (String) ID of database. Can be retrieved using `mssql_database` or `SELECT DB_ID('<db_name>')`.
- `name` (String) User name. Cannot be longer than 128 chars.
- `user_object_id` (String) Azure AD object_id of the user. This can be either regular user or a group.

### Read-Only

- `id` (String) `<database_id>/<user_id>`. User ID can be retrieved using `sys.database_principals` view.

## Import

Import is supported using the following syntax:

```shell
# import using <db_id>/<user_id> - can be retrieved using `SELECT CONCAT(DB_ID(), '/', principal_id) FROM sys.database_principals WHERE [name] = '<username>'`
terraform import mssql_azuread_user.example '7/5'
```