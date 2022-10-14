package sqlLogin

import (
	"github.com/PGSSoft/terraform-provider-mssql/internal/core"
	"github.com/PGSSoft/terraform-provider-mssql/internal/core/datasource"
	"github.com/PGSSoft/terraform-provider-mssql/internal/core/resource"
	sdkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	sdkResource "github.com/hashicorp/terraform-plugin-framework/resource"
)

func Service() core.Service {
	return service{}
}

type service struct{}

func (s service) Name() string {
	return "sql_login"
}

func (s service) Resources() []func() sdkResource.ResourceWithConfigure {
	return []func() sdkResource.ResourceWithConfigure{
		resource.NewResource[resourceData](&res{}),
	}
}

func (s service) DataSources() []func() sdkdatasource.DataSourceWithConfigure {
	return []func() sdkdatasource.DataSourceWithConfigure{
		datasource.NewDataSource[dataSourceData](&dataSource{}),
		datasource.NewDataSource[listDataSourceData](&listDataSource{}),
	}
}

func (s service) Tests() core.AccTests {
	return core.AccTests{
		Resource:       testResource,
		DataSource:     testDataSource,
		ListDataSource: testListDataSource,
	}
}
