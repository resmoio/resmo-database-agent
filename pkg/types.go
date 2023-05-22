package pkg

import (
	"database/sql"
)

type Data struct {
	name   string
	values interface{}
	query  string
}

type MySQLColumn struct {
	TableCatalog string `db:"TABLE_CATALOG"`
	TableSchema  string `db:"TABLE_SCHEMA"`
	TableName    string `db:"TABLE_NAME"`
	ColumnName   string `db:"COLUMN_NAME"`
	IsNullable   string `db:"IS_NULLABLE"`
	DataType     string `db:"DATA_TYPE"`
	ColumnType   string `db:"COLUMN_TYPE"`
	Privileges   string `db:"PRIVILEGES"`
}

type MySQLSchema struct {
	CatalogName         string `db:"CATALOG_NAME"`
	SchemaName          string `db:"SCHEMA_NAME"`
	DefaultCharacterSet string `db:"DEFAULT_CHARACTER_SET_NAME"`
	DefaultEncryption   string `db:"DEFAULT_ENCRYPTION"`
}

type MySQLTable struct {
	TableCatalog  string         `db:"TABLE_CATALOG"`
	TableSchema   string         `db:"TABLE_SCHEMA"`
	TableName     string         `db:"TABLE_NAME"`
	TableType     string         `db:"TABLE_TYPE"`
	Engine        string         `db:"ENGINE"`
	Version       string         `db:"VERSION"`
	RowFormat     string         `db:"ROW_FORMAT"`
	TableRows     sql.NullInt64  `db:"TABLE_ROWS"`
	AutoIncrement sql.NullInt64  `db:"AUTO_INCREMENT"`
	CreateTime    string         `db:"CREATE_TIME"`
	UpdateTime    sql.NullString `db:"UPDATE_TIME"`
}

type MySQLTrigger struct {
	TriggerCatalog string `db:"TRIGGER_CATALOG"`
	TriggerSchema  string `db:"TRIGGER_SCHEMA"`
	TriggerName    string `db:"TRIGGER_NAME"`
	Created        string `db:"CREATED"`
}

type MySQLUserPrivilege struct {
	Grantee       string `db:"grantee"`
	TableCatalog  string `db:"table_catalog"`
	PrivilegeType string `db:"privilege_type"`
	IsGrantable   string `db:"is_grantable"`
}

type MySQLIndex struct {
	TableName string `db:"TABLE_NAME"`
	IndexName string `db:"INDEX_NAME"`
}

type MySQLUser struct {
	Host                 string         `db:"host"`
	User                 string         `db:"user"`
	SelectPriv           string         `db:"select_priv"`
	InsertPriv           string         `db:"insert_priv"`
	UpdatePriv           string         `db:"update_priv"`
	DeletePriv           string         `db:"delete_priv"`
	CreatePriv           string         `db:"create_priv"`
	DropPriv             string         `db:"drop_priv"`
	ReloadPriv           string         `db:"reload_priv"`
	ShutdownPriv         string         `db:"shutdown_priv"`
	ProcessPriv          string         `db:"process_priv"`
	FilePriv             string         `db:"file_priv"`
	GrantPriv            string         `db:"grant_priv"`
	ReferencesPriv       string         `db:"references_priv"`
	IndexPriv            string         `db:"index_priv"`
	AlterPriv            string         `db:"alter_priv"`
	ShowDBPriv           string         `db:"show_db_priv"`
	SuperPriv            string         `db:"super_priv"`
	CreateTmpTablePriv   string         `db:"create_tmp_table_priv"`
	LockTablesPriv       string         `db:"lock_tables_priv"`
	ExecutePriv          string         `db:"execute_priv"`
	ReplSlavePriv        string         `db:"repl_slave_priv"`
	ReplClientPriv       string         `db:"repl_client_priv"`
	CreateViewPriv       string         `db:"create_view_priv"`
	ShowViewPriv         string         `db:"show_view_priv"`
	CreateRoutinePriv    string         `db:"create_routine_priv"`
	AlterRoutinePriv     string         `db:"alter_routine_priv"`
	CreateUserPriv       string         `db:"create_user_priv"`
	EventPriv            string         `db:"event_priv"`
	TriggerPriv          string         `db:"trigger_priv"`
	CreateTablespacePriv string         `db:"create_tablespace_priv"`
	CreateRolePriv       string         `db:"create_role_priv"`
	DropRolePriv         string         `db:"drop_role_priv"`
	SSLType              sql.NullString `db:"ssl_type"`
	PasswordExpired      string         `db:"password_expired"`
}

type PostgresColumn struct {
	TableCatalog      string `db:"table_catalog"`
	TableSchema       string `db:"table_schema"`
	TableName         string `db:"table_name"`
	ColumnName        string `db:"column_name"`
	IsNullable        string `db:"is_nullable"`
	DataType          string `db:"data_type"`
	UdtCatalog        string `db:"udt_catalog"`
	UdtSchema         string `db:"udt_schema"`
	UdtName           string `db:"udt_name"`
	IsSelfReferencing string `db:"is_self_referencing"`
	IsIdentity        string `db:"is_identity"`
	IdentityCycle     string `db:"identity_cycle"`
	IsGenerated       string `db:"is_generated"`
	IsUpdatable       string `db:"is_updatable"`
}

type PostgresDatabase struct {
	DatabaseName     string `db:"datname"`
	IsTemplate       bool   `db:"datistemplate"`
	AllowConnections bool   `db:"datallowconn"`
	ConnLimit        int    `db:"datconnlimit"`
}

type PostgresRole struct {
	RoleName        string `db:"rolname"`
	IsSuperUser     bool   `db:"rolsuper"`
	RoleInherit     bool   `db:"rolinherit"`
	CreateRole      bool   `db:"rolcreaterole"`
	CreateDatabase  bool   `db:"rolcreatedb"`
	CanLogin        bool   `db:"rolcanlogin"`
	IsReplication   bool   `db:"rolreplication"`
	ConnectionLimit int    `db:"rolconnlimit"`
	BypassRls       bool   `db:"rolbypassrls"`
}

type PostgresUser struct {
	UserName       string         `db:"usename"`
	CreateDatabase bool           `db:"usecreatedb"`
	IsSuperUser    bool           `db:"usesuper"`
	IsReplication  bool           `db:"userepl"`
	BypassRls      bool           `db:"usebypassrls"`
	UserConfig     sql.NullString `db:"useconfig"`
}

type PostgresSchema struct {
	CatalogName string `db:"catalog_name"`
	SchemaName  string `db:"schema_name"`
	SchemaOwner string `db:"schema_owner"`
}

type PostgresTable struct {
	TableCatalog string `db:"table_catalog"`
	TableSchema  string `db:"table_schema"`
	TableName    string `db:"table_name"`
	TableType    string `db:"table_type"`
	IsInsertable string `db:"is_insertable_into"`
	IsTyped      string `db:"is_typed"`
}

type PostgresIndex struct {
	IndexName  string `db:"indexname"`
	TableName  string `db:"tablename"`
	SchemaName string `db:"schemaname"`
}

type PostgresTablePrivilege struct {
	Grantor       string `db:"grantor"`
	Grantee       string `db:"grantee"`
	TableCatalog  string `db:"table_catalog"`
	TableSchema   string `db:"table_schema"`
	TableName     string `db:"table_name"`
	PrivilegeType string `db:"privilege_type"`
	IsGrantable   string `db:"is_grantable"`
	WithHierarchy string `db:"with_hierarchy"`
}

type PostgresColumnPrivilege struct {
	Grantor       string `db:"grantor"`
	Grantee       string `db:"grantee"`
	TableCatalog  string `db:"table_catalog"`
	TableSchema   string `db:"table_schema"`
	TableName     string `db:"table_name"`
	ColumnName    string `db:"column_name"`
	PrivilegeType string `db:"privilege_type"`
	IsGrantable   string `db:"is_grantable"`
}

type PostgresTrigger struct {
	TriggerCatalog string `db:"trigger_catalog"`
	TriggerSchema  string `db:"trigger_schema"`
	TriggerName    string `db:"trigger_name"`
}

type PostgresExtension struct {
	ExtName        string `db:"extname"`
	ExtOwner       string `db:"extowner"`
	ExtRelocatable bool   `db:"extrelocatable"`
	ExtVersion     string `db:"extversion"`
}

type ClickHouseUser struct {
	Name               string   `db:"name"`
	ID                 string   `db:"id"`
	Storage            string   `db:"storage"`
	HostIP             []string `db:"host_ip"`
	HostNames          []string `db:"host_names"`
	DefaultRolesAll    bool     `db:"default_roles_all"`
	DefaultRolesExcept []string `db:"default_roles_except"`
	DefaultRolesList   []string `db:"default_roles_list"`
	GranteesAny        bool     `db:"grantees_any"`
	GranteesList       []string `db:"grantees_list"`
	GranteesExcept     []string `db:"grantees_except"`
}

type ClickHouseSchema struct {
	CatalogName string `db:"catalog_name"`
	SchemaName  string `db:"schema_name"`
	SchemaOwner string `db:"schema_owner"`
}

type ClickHouseTable struct {
	TableCatalog string `db:"table_catalog"`
	TableSchema  string `db:"table_schema"`
	TableName    string `db:"table_name"`
	TableType    string `db:"table_type"`
}

type ClickHouseRole struct {
	Name    string `db:"name"`
	ID      string `db:"id"`
	Storage string `db:"storage"`
}

type ClickHouseColumn struct {
	TableCatalog  string `db:"table_catalog"`
	TableSchema   string `db:"table_schema"`
	TableName     string `db:"table_name"`
	ColumnName    string `db:"column_name"`
	IsNullable    bool   `db:"is_nullable"`
	DataType      string `db:"data_type"`
	ColumnDefault string `db:"column_default"`
}

type MongoCollection struct {
	Name     string `json:"name"`
	Database string `json:"database"`
}

type MongoDatabase struct {
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	Empty bool   `json:"empty"`
}

type MongoRole struct {
	Role           string              `json:"role"`
	DB             string              `json:"db"`
	InheritedRoles []map[string]string `json:"inheritedRoles"`
	Roles          []map[string]string `json:"roles"`
}

type MongoUser struct {
	User  string              `json:"user"`
	DB    string              `json:"db"`
	Roles []map[string]string `json:"roles"`
}

var (
	mysqlQueries = []Data{
		{name: "mysql_column", values: &[]MySQLColumn{}, query: "SELECT table_catalog, table_schema, table_name, column_name, is_nullable, data_type, column_type, privileges FROM information_schema.columns WHERE table_schema NOT IN ('mysql', 'information_schema', 'sys', 'performance_schema')"},
		{name: "mysql_schema", values: &[]MySQLSchema{}, query: "SELECT catalog_name, schema_name, default_character_set_name, default_encryption FROM information_schema.schemata WHERE schema_name NOT IN ('mysql', 'information_schema', 'sys', 'performance_schema')"},
		{name: "mysql_table", values: &[]MySQLTable{}, query: "SELECT table_catalog, table_schema, table_name, table_type, engine, version, row_format, table_rows, auto_increment, create_time, update_time FROM information_schema.tables WHERE table_schema NOT IN ('mysql', 'information_schema', 'sys', 'performance_schema')"},
		{name: "mysql_trigger", values: &[]MySQLTrigger{}, query: "SELECT trigger_catalog, trigger_schema, trigger_name, created FROM information_schema.triggers WHERE trigger_schema NOT IN ('mysql', 'information_schema', 'sys', 'performance_schema')"},
		{name: "mysql_user_privilege", values: &[]MySQLUserPrivilege{}, query: "SELECT grantee, table_catalog, privilege_type, is_grantable FROM information_schema.user_privileges"},
		{name: "mysql_index", values: &[]MySQLIndex{}, query: "SELECT DISTINCT TABLE_NAME, INDEX_NAME FROM INFORMATION_SCHEMA.STATISTICS WHERE table_schema NOT IN ('mysql', 'information_schema', 'sys', 'performance_schema')"},
		{name: "mysql_user", values: &[]MySQLUser{}, query: "SELECT host, user, select_priv, insert_priv, update_priv, delete_priv, create_priv, drop_priv, reload_priv, shutdown_priv, process_priv, file_priv, grant_priv, references_priv, index_priv, alter_priv, show_db_priv, super_priv, create_tmp_table_priv, lock_tables_priv, execute_priv, repl_slave_priv, repl_client_priv, create_view_priv, show_view_priv, create_routine_priv, alter_routine_priv, create_user_priv, event_priv, trigger_priv, create_tablespace_priv, create_role_priv, drop_role_priv, user_attributes ssl_type, password_expired FROM mysql.user"},
	}

	postgresQueries = []Data{
		{name: "postgres_column", values: &[]PostgresColumn{}, query: "SELECT table_catalog, table_schema, table_name, column_name, is_nullable, data_type, udt_catalog, udt_schema, udt_name, is_self_referencing, is_identity, identity_cycle, is_generated, is_updatable FROM information_schema.columns WHERE table_schema NOT IN ('pg_catalog', 'information_schema')"},
		{name: "postgres_database", values: &[]PostgresDatabase{}, query: "SELECT datname, datistemplate, datallowconn, datconnlimit FROM pg_database"},
		{name: "postgres_role", values: &[]PostgresRole{}, query: "SELECT rolname, rolsuper, rolinherit, rolcreaterole, rolcreatedb, rolcanlogin, rolreplication, rolconnlimit, rolbypassrls FROM pg_roles"},
		{name: "postgres_user", values: &[]PostgresUser{}, query: "SELECT usename, usecreatedb, usesuper, userepl, usebypassrls, useconfig FROM pg_user"},
		{name: "postgres_schema", values: &[]PostgresSchema{}, query: "SELECT catalog_name, schema_name, schema_owner FROM information_schema.schemata"},
		{name: "postgres_table", values: &[]PostgresTable{}, query: "SELECT t.table_catalog, t.table_schema, t.table_name, t.table_type, t.is_insertable_into, t.is_typed FROM information_schema.tables t INNER JOIN information_schema.schemata s ON t.table_schema = s.schema_name WHERE t.table_schema NOT IN ('pg_catalog', 'information_schema')"},
		{name: "postgres_index", values: &[]PostgresIndex{}, query: "SELECT indexname, tablename, schemaname FROM pg_indexes"},
		{name: "postgres_table_privilege", values: &[]PostgresTablePrivilege{}, query: "SELECT grantor, grantee, table_catalog, table_schema, table_name, privilege_type, is_grantable, with_hierarchy FROM information_schema.table_privileges WHERE table_schema NOT IN ('pg_catalog', 'information_schema')"},
		{name: "postgres_column_privilege", values: &[]PostgresColumnPrivilege{}, query: "SELECT grantor, grantee, table_catalog, table_schema, table_name, column_name, privilege_type, is_grantable FROM information_schema.column_privileges WHERE table_schema NOT IN ('pg_catalog', 'information_schema')"},
		{name: "postgres_trigger", values: &[]PostgresTrigger{}, query: "SELECT trigger_catalog, trigger_schema, trigger_name FROM information_schema.triggers WHERE trigger_schema NOT IN ('pg_catalog', 'information_schema')"},
		{name: "postgres_extension", values: &[]PostgresExtension{}, query: "SELECT extname, extowner, extrelocatable, extversion FROM pg_extension"},
	}

	clickhouseQueries = []Data{
		{name: "clickhouse_user", values: &[]ClickHouseUser{}, query: "SELECT name, id, storage, host_ip, host_names, default_roles_all, default_roles_except, default_roles_list, grantees_any, grantees_list, grantees_except FROM system.users"},
		{name: "clickhouse_schema", values: &[]ClickHouseSchema{}, query: "SELECT catalog_name, schema_name, schema_owner FROM information_schema.schemata WHERE schema_name NOT IN ('information_schema', 'INFORMATION_SCHEMA', 'system')"},
		{name: "clickhouse_table", values: &[]ClickHouseTable{}, query: "SELECT table_catalog, table_schema, table_name, table_type FROM information_schema.tables WHERE table_schema NOT IN ('information_schema', 'INFORMATION_SCHEMA', 'system')"},
		{name: "clickhouse_role", values: &[]ClickHouseRole{}, query: "SELECT name, id, storage FROM system.roles"},
		{name: "clickhouse_column", values: &[]ClickHouseColumn{}, query: "SELECT table_catalog, table_schema, table_name, column_name, is_nullable, data_type, column_default FROM information_schema.columns WHERE table_schema NOT IN ('information_schema', 'INFORMATION_SCHEMA', 'system')"},
	}
)
