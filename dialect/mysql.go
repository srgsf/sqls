package dialect

var mysql8Keyword = []string{
	"ACCESSIBLE",
	"ACCOUNT",
	"ACTION",
	"ACTIVE",
	"ADD",
	"ADMIN",
	"AFTER",
	"AGAINST",
	"AGGREGATE",
	"ALGORITHM",
	"ALL",
	"ALTER",
	"ALWAYS",
	"ANALYSE",
	"ANALYZE",
	"AND",
	"ANY",
	"ARRAY",
	"AS",
	"ASC",
	"ASCII",
	"ASENSITIVE",
	"AT",
	"ATTRIBUTE",
	"AUTOEXTEND_SIZE",
	"AUTO_INCREMENT",
	"AVG",
	"AVG_ROW_LENGTH",
	"BACKUP",
	"BEFORE",
	"BEGIN",
	"BETWEEN",
	"BIGINT",
	"BINARY",
	"BINLOG",
	"BIT",
	"BLOB",
	"BLOCK",
	"BOOL",
	"BOOLEAN",
	"BOTH",
	"BTREE",
	"BUCKETS",
	"BY",
	"BYTE",
	"CACHE",
	"CALL",
	"CASCADE",
	"CASCADED",
	"CASE",
	"CATALOG_NAME",
	"CHAIN",
	"CHANGE",
	"CHANGED",
	"CHANNEL",
	"CHAR",
	"CHARACTER",
	"CHARSET",
	"CHECK",
	"CHECKSUM",
	"CIPHER",
	"CLASS_ORIGIN",
	"CLIENT",
	"CLONE",
	"CLOSE",
	"COALESCE",
	"CODE",
	"COLLATE",
	"COLLATION",
	"COLUMN",
	"COLUMNS",
	"COLUMN_FORMAT",
	"COLUMN_NAME",
	"COMMENT",
	"COMMIT",
	"COMMITTED",
	"COMPACT",
	"COMPLETION",
	"COMPONENT",
	"COMPRESSED",
	"COMPRESSION",
	"CONCURRENT",
	"CONDITION",
	"CONNECTION",
	"CONSISTENT",
	"CONSTRAINT",
	"CONSTRAINT_CATALOG",
	"CONSTRAINT_NAME",
	"CONSTRAINT_SCHEMA",
	"CONTAINS",
	"CONTEXT",
	"CONTINUE",
	"CONVERT",
	"CPU",
	"CREATE",
	"CROSS",
	"CUBE",
	"CUME_DIST",
	"CURRENT",
	"CURRENT_DATE",
	"CURRENT_TIME",
	"CURRENT_TIMESTAMP",
	"CURRENT_USER",
	"CURSOR",
	"CURSOR_NAME",
	"DATA",
	"DATABASE",
	"DATABASES",
	"DATAFILE",
	"DATE",
	"DATETIME",
	"DAY",
	"DAY_HOUR",
	"DAY_MICROSECOND",
	"DAY_MINUTE",
	"DAY_SECOND",
	"DEALLOCATE",
	"DEC",
	"DECIMAL",
	"DECLARE",
	"DEFAULT",
	"DEFAULT_AUTH",
	"DEFINER",
	"DEFINITION",
	"DELAYED",
	"DELAY_KEY_WRITE",
	"DELETE",
	"DENSE_RANK",
	"DESC",
	"DESCRIBE",
	"DESCRIPTION",
	"DES_KEY_FILE",
	"DETERMINISTIC",
	"DIAGNOSTICS",
	"DIRECTORY",
	"DISABLE",
	"DISCARD",
	"DISK",
	"DISTINCT",
	"DISTINCTROW",
	"DIV",
	"DO",
	"DOUBLE",
	"DROP",
	"DUAL",
	"DUMPFILE",
	"DUPLICATE",
	"DYNAMIC",
	"EACH",
	"ELSE",
	"ELSEIF",
	"EMPTY",
	"ENABLE",
	"ENCLOSED",
	"ENCRYPTION",
	"END",
	"ENDS",
	"ENFORCED",
	"ENGINE",
	"ENGINES",
	"ENGINE_ATTRIBUTE",
	"ENUM",
	"ERROR",
	"ERRORS",
	"ESCAPE",
	"ESCAPED",
	"EVENT",
	"EVENTS",
	"EVERY",
	"EXCEPT",
	"EXCHANGE",
	"EXCLUDE",
	"EXECUTE",
	"EXISTS",
	"EXIT",
	"EXPANSION",
	"EXPIRE",
	"EXPLAIN",
	"EXPORT",
	"EXTENDED",
	"EXTENT_SIZE",
	"FAILED_LOGIN_ATTEMPTS",
	"FALSE",
	"FAST",
	"FAULTS",
	"FETCH",
	"FIELDS",
	"FILE",
	"FILE_BLOCK_SIZE",
	"FILTER",
	"FIRST",
	"FIRST_VALUE",
	"FIXED",
	"FLOAT",
	"FLOAT4",
	"FLOAT8",
	"FLUSH",
	"FOLLOWING",
	"FOLLOWS",
	"FOR",
	"FORCE",
	"FOREIGN",
	"FORMAT",
	"FOUND",
	"FROM",
	"FULL",
	"FULLTEXT",
	"FUNCTION",
	"GENERAL",
	"GENERATED",
	"GEOMCOLLECTION",
	"GEOMETRY",
	"GEOMETRYCOLLECTION",
	"GET",
	"GET_FORMAT",
	"GET_MASTER_PUBLIC_KEY",
	"GET_SOURCE_PUBLIC_KEY",
	"GLOBAL",
	"GRANT",
	"GRANTS",
	"GROUP",
	"GROUPING",
	"GROUPS",
	"GROUP_REPLICATION",
	"HANDLER",
	"HASH",
	"HAVING",
	"HELP",
	"HIGH_PRIORITY",
	"HISTOGRAM",
	"HISTORY",
	"HOST",
	"HOSTS",
	"HOUR",
	"HOUR_MICROSECOND",
	"HOUR_MINUTE",
	"HOUR_SECOND",
	"IDENTIFIED",
	"IF",
	"IGNORE",
	"IGNORE_SERVER_IDS",
	"IMPORT",
	"IN",
	"INACTIVE",
	"INDEX",
	"INDEXES",
	"INFILE",
	"INITIAL_SIZE",
	"INNER",
	"INOUT",
	"INSENSITIVE",
	"INSERT",
	"INSERT_METHOD",
	"INSTALL",
	"INSTANCE",
	"INT",
	"INT1",
	"INT2",
	"INT3",
	"INT4",
	"INT8",
	"INTEGER",
	"INTERVAL",
	"INTO",
	"INVISIBLE",
	"INVOKER",
	"IO",
	"IO_AFTER_GTIDS",
	"IO_BEFORE_GTIDS",
	"IO_THREAD",
	"IPC",
	"IS",
	"ISOLATION",
	"ISSUER",
	"ITERATE",
	"JOIN",
	"JSON",
	"JSON_TABLE",
	"JSON_VALUE",
	"KEY",
	"KEYRING",
	"KEYS",
	"KEY_BLOCK_SIZE",
	"KILL",
	"LAG",
	"LANGUAGE",
	"LAST",
	"LAST_VALUE",
	"LATERAL",
	"LEAD",
	"LEADING",
	"LEAVE",
	"LEAVES",
	"LEFT",
	"LESS",
	"LEVEL",
	"LIKE",
	"LIMIT",
	"LINEAR",
	"LINES",
	"LINESTRING",
	"LIST",
	"LOAD",
	"LOCAL",
	"LOCALTIME",
	"LOCALTIMESTAMP",
	"LOCK",
	"LOCKED",
	"LOCKS",
	"LOGFILE",
	"LOGS",
	"LONG",
	"LONGBLOB",
	"LONGTEXT",
	"LOOP",
	"LOW_PRIORITY",
	"MASTER",
	"MASTER_AUTO_POSITION",
	"MASTER_BIND",
	"MASTER_COMPRESSION_ALGORITHMS",
	"MASTER_CONNECT_RETRY",
	"MASTER_DELAY",
	"MASTER_HEARTBEAT_PERIOD",
	"MASTER_HOST",
	"MASTER_LOG_FILE",
	"MASTER_LOG_POS",
	"MASTER_PASSWORD",
	"MASTER_PORT",
	"MASTER_PUBLIC_KEY_PATH",
	"MASTER_RETRY_COUNT",
	"MASTER_SERVER_ID",
	"MASTER_SSL",
	"MASTER_SSL_CA",
	"MASTER_SSL_CAPATH",
	"MASTER_SSL_CERT",
	"MASTER_SSL_CIPHER",
	"MASTER_SSL_CRL",
	"MASTER_SSL_CRLPATH",
	"MASTER_SSL_KEY",
	"MASTER_SSL_VERIFY_SERVER_CERT",
	"MASTER_TLS_CIPHERSUITES",
	"MASTER_TLS_VERSION",
	"MASTER_USER",
	"MASTER_ZSTD_COMPRESSION_LEVEL",
	"MATCH",
	"MAXVALUE",
	"MAX_CONNECTIONS_PER_HOUR",
	"MAX_QUERIES_PER_HOUR",
	"MAX_ROWS",
	"MAX_SIZE",
	"MAX_UPDATES_PER_HOUR",
	"MAX_USER_CONNECTIONS",
	"MEDIUM",
	"MEDIUMBLOB",
	"MEDIUMINT",
	"MEDIUMTEXT",
	"MEMBER",
	"MEMORY",
	"MERGE",
	"MESSAGE_TEXT",
	"MICROSECOND",
	"MIDDLEINT",
	"MIGRATE",
	"MINUTE",
	"MINUTE_MICROSECOND",
	"MINUTE_SECOND",
	"MIN_ROWS",
	"MOD",
	"MODE",
	"MODIFIES",
	"MODIFY",
	"MONTH",
	"MULTILINESTRING",
	"MULTIPOINT",
	"MULTIPOLYGON",
	"MUTEX",
	"MYSQL_ERRNO",
	"NAME",
	"NAMES",
	"NATIONAL",
	"NATURAL",
	"NCHAR",
	"NDB",
	"NDBCLUSTER",
	"NESTED",
	"NETWORK_NAMESPACE",
	"NEVER",
	"NEW",
	"NEXT",
	"NO",
	"NODEGROUP",
	"NONE",
	"NOT",
	"NOWAIT",
	"NO_WAIT",
	"NO_WRITE_TO_BINLOG",
	"NTH_VALUE",
	"NTILE",
	"NULL",
	"NULLS",
	"NUMBER",
	"NUMERIC",
	"NVARCHAR",
	"OF",
	"OFF",
	"OFFSET",
	"OJ",
	"OLD",
	"ON",
	"ONE",
	"ONLY",
	"OPEN",
	"OPTIMIZE",
	"OPTIMIZER_COSTS",
	"OPTION",
	"OPTIONAL",
	"OPTIONALLY",
	"OPTIONS",
	"OR",
	"ORDER",
	"ORDINALITY",
	"ORGANIZATION",
	"OTHERS",
	"OUT",
	"OUTER",
	"OUTFILE",
	"OVER",
	"OWNER",
	"PACK_KEYS",
	"PAGE",
	"PARSER",
	"PARTIAL",
	"PARTITION",
	"PARTITIONING",
	"PARTITIONS",
	"PASSWORD",
	"PASSWORD_LOCK_TIME",
	"PATH",
	"PERCENT_RANK",
	"PERSIST",
	"PERSIST_ONLY",
	"PHASE",
	"PLUGIN",
	"PLUGINS",
	"PLUGIN_DIR",
	"POINT",
	"POLYGON",
	"PORT",
	"PRECEDES",
	"PRECEDING",
	"PRECISION",
	"PREPARE",
	"PRESERVE",
	"PREV",
	"PRIMARY",
	"PRIVILEGES",
	"PRIVILEGE_CHECKS_USER",
	"PROCEDURE",
	"PROCESS",
	"PROCESSLIST",
	"PROFILE",
	"PROFILES",
	"PROXY",
	"PURGE",
	"QUARTER",
	"QUERY",
	"QUICK",
	"RANDOM",
	"RANGE",
	"RANK",
	"READ",
	"READS",
	"READ_ONLY",
	"READ_WRITE",
	"REAL",
	"REBUILD",
	"RECOVER",
	"RECURSIVE",
	"REDOFILE",
	"REDO_BUFFER_SIZE",
	"REDUNDANT",
	"REFERENCE",
	"REFERENCES",
	"REGEXP",
	"RELAY",
	"RELAYLOG",
	"RELAY_LOG_FILE",
	"RELAY_LOG_POS",
	"RELAY_THREAD",
	"RELEASE",
	"RELOAD",
	"REMOTE",
	"REMOVE",
	"RENAME",
	"REORGANIZE",
	"REPAIR",
	"REPEAT",
	"REPEATABLE",
	"REPLACE",
	"REPLICA",
	"REPLICAS",
	"REPLICATE_DO_DB",
	"REPLICATE_DO_TABLE",
	"REPLICATE_IGNORE_DB",
	"REPLICATE_IGNORE_TABLE",
	"REPLICATE_REWRITE_DB",
	"REPLICATE_WILD_DO_TABLE",
	"REPLICATE_WILD_IGNORE_TABLE",
	"REPLICATION",
	"REQUIRE",
	"REQUIRE_ROW_FORMAT",
	"RESET",
	"RESIGNAL",
	"RESOURCE",
	"RESPECT",
	"RESTART",
	"RESTORE",
	"RESTRICT",
	"RESUME",
	"RETAIN",
	"RETURN",
	"RETURNED_SQLSTATE",
	"RETURNING",
	"RETURNS",
	"REUSE",
	"REVERSE",
	"REVOKE",
	"RIGHT",
	"RLIKE",
	"ROLE",
	"ROLLBACK",
	"ROLLUP",
	"ROTATE",
	"ROUTINE",
	"ROW",
	"ROWS",
	"ROW_COUNT",
	"ROW_FORMAT",
	"ROW_NUMBER",
	"RTREE",
	"SAVEPOINT",
	"SCHEDULE",
	"SCHEMA",
	"SCHEMAS",
	"SCHEMA_NAME",
	"SECOND",
	"SECONDARY",
	"SECONDARY_ENGINE",
	"SECONDARY_ENGINE_ATTRIBUTE",
	"SECONDARY_LOAD",
	"SECONDARY_UNLOAD",
	"SECOND_MICROSECOND",
	"SECURITY",
	"SELECT",
	"SENSITIVE",
	"SEPARATOR",
	"SERIAL",
	"SERIALIZABLE",
	"SERVER",
	"SESSION",
	"SET",
	"SHARE",
	"SHOW",
	"SHUTDOWN",
	"SIGNAL",
	"SIGNED",
	"SIMPLE",
	"SKIP",
	"SLAVE",
	"SLOW",
	"SMALLINT",
	"SNAPSHOT",
	"SOCKET",
	"SOME",
	"SONAME",
	"SOUNDS",
	"SOURCE",
	"SOURCE_AUTO_POSITION",
	"SOURCE_BIND",
	"SOURCE_COMPRESSION_ALGORITHMS",
	"SOURCE_CONNECT_RETRY",
	"SOURCE_DELAY",
	"SOURCE_HEARTBEAT_PERIOD",
	"SOURCE_HOST",
	"SOURCE_LOG_FILE",
	"SOURCE_LOG_POS",
	"SOURCE_PASSWORD",
	"SOURCE_PORT",
	"SOURCE_PUBLIC_KEY_PATH",
	"SOURCE_RETRY_COUNT",
	"SOURCE_SSL",
	"SOURCE_SSL_CA",
	"SOURCE_SSL_CAPATH",
	"SOURCE_SSL_CERT",
	"SOURCE_SSL_CIPHER",
	"SOURCE_SSL_CRL",
	"SOURCE_SSL_CRLPATH",
	"SOURCE_SSL_KEY",
	"SOURCE_SSL_VERIFY_SERVER_CERT",
	"SOURCE_TLS_CIPHERSUITES",
	"SOURCE_TLS_VERSION",
	"SOURCE_USER",
	"SOURCE_ZSTD_COMPRESSION_LEVEL",
	"SPATIAL",
	"SPECIFIC",
	"SQL",
	"SQLEXCEPTION",
	"SQLSTATE",
	"SQLWARNING",
	"SQL_AFTER_GTIDS",
	"SQL_AFTER_MTS_GAPS",
	"SQL_BEFORE_GTIDS",
	"SQL_BIG_RESULT",
	"SQL_BUFFER_RESULT",
	"SQL_CACHE",
	"SQL_CALC_FOUND_ROWS",
	"SQL_NO_CACHE",
	"SQL_SMALL_RESULT",
	"SQL_THREAD",
	"SQL_TSI_DAY",
	"SQL_TSI_HOUR",
	"SQL_TSI_MINUTE",
	"SQL_TSI_MONTH",
	"SQL_TSI_QUARTER",
	"SQL_TSI_SECOND",
	"SQL_TSI_WEEK",
	"SQL_TSI_YEAR",
	"SRID",
	"SSL",
	"STACKED",
	"START",
	"STARTING",
	"STARTS",
	"STATS_AUTO_RECALC",
	"STATS_PERSISTENT",
	"STATS_SAMPLE_PAGES",
	"STATUS",
	"STOP",
	"STORAGE",
	"STORED",
	"STRAIGHT_JOIN",
	"STREAM",
	"STRING",
	"SUBCLASS_ORIGIN",
	"SUBJECT",
	"SUBPARTITION",
	"SUBPARTITIONS",
	"SUPER",
	"SUSPEND",
	"SWAPS",
	"SWITCHES",
	"SYSTEM",
	"TABLE",
	"TABLES",
	"TABLESPACE",
	"TABLE_CHECKSUM",
	"TABLE_NAME",
	"TEMPORARY",
	"TEMPTABLE",
	"TERMINATED",
	"TEXT",
	"THAN",
	"THEN",
	"THREAD_PRIORITY",
	"TIES",
	"TIME",
	"TIMESTAMP",
	"TIMESTAMPADD",
	"TIMESTAMPDIFF",
	"TINYBLOB",
	"TINYINT",
	"TINYTEXT",
	"TLS",
	"TO",
	"TRAILING",
	"TRANSACTION",
	"TRIGGER",
	"TRIGGERS",
	"TRUE",
	"TRUNCATE",
	"TYPE",
	"TYPES",
	"UNBOUNDED",
	"UNCOMMITTED",
	"UNDEFINED",
	"UNDO",
	"UNDOFILE",
	"UNDO_BUFFER_SIZE",
	"UNICODE",
	"UNINSTALL",
	"UNION",
	"UNIQUE",
	"UNKNOWN",
	"UNLOCK",
	"UNSIGNED",
	"UNTIL",
	"UPDATE",
	"UPGRADE",
	"USAGE",
	"USE",
	"USER",
	"USER_RESOURCES",
	"USE_FRM",
	"USING",
	"UTC_DATE",
	"UTC_TIME",
	"UTC_TIMESTAMP",
	"VALIDATION",
	"VALUE",
	"VALUES",
	"VARBINARY",
	"VARCHAR",
	"VARCHARACTER",
	"VARIABLES",
	"VARYING",
	"VCPU",
	"VIEW",
	"VIRTUAL",
	"VISIBLE",
	"WAIT",
	"WARNINGS",
	"WEEK",
	"WEIGHT_STRING",
	"WHEN",
	"WHERE",
	"WHILE",
	"WINDOW",
	"WITH",
	"WITHOUT",
	"WORK",
	"WRAPPER",
	"WRITE",
	"X509",
	"XA",
	"XID",
	"XML",
	"XOR",
	"YEAR",
	"YEAR_MONTH",
	"ZEROFILL",
	"ZONE",
}

var mysql57Keyword = []string{
	"ACCOUNT",
	"ACTION",
	"ADD",
	"AFTER",
	"AGGREGATE",
	"ALGORITHM",
	"ALL",
	"ALTER",
	"ANALYSE",
	"ANALYZE",
	"ARCHIVE",
	"AS",
	"ASC",
	"AT",
	"AUTOCOMMIT",
	"AUTOEXTEND_SIZE",
	"AUTO_INCREMENT",
	"AVG_ROW_LENGTH",
	"BEFORE",
	"BEGIN",
	"BINARY",
	"BINLOG",
	"BOOL",
	"BOOLEAN",
	"BTREE",
	"BY",
	"BYTE",
	"CACHE",
	"CALL",
	"CASCADE",
	"CASE",
	"CATALOG_NAME",
	"CHAIN",
	"CHANGE",
	"CHANNEL",
	"CHAR",
	"CHARACTER",
	"CHARSET",
	"CHECK",
	"CHECKSUM",
	"CIPHER",
	"CLASS_ORIGIN",
	"CLIENT",
	"CLOSE",
	"COALESCE",
	"CODE",
	"COLLATE",
	"COLLATION",
	"COLUMN",
	"COLUMNS",
	"COLUMN_NAME",
	"COMMENT",
	"COMMIT",
	"COMMITTED",
	"COMPACT",
	"COMPLETION",
	"COMPRESSED",
	"COMPRESSION",
	"CONCURRENT",
	"CONDITION",
	"CONNECTION",
	"CONSISTENT",
	"CONSTRAINT",
	"CONSTRAINT_CATALOG",
	"CONSTRAINT_NAME",
	"CONSTRAINT_SCHEMA",
	"CONTINUE",
	"COUNT",
	"CREATE",
	"CROSS",
	"CSV",
	"CURRENT_USER",
	"CURSOR",
	"CURSOR_NAME",
	"DATA",
	"DATABASE",
	"DATABASES",
	"DATAFILE",
	"DATE",
	"DEALLOCATE",
	"DEC",
	"DECIMAL",
	"DECLARE",
	"DEFAULT",
	"DEFAULT_AUTH",
	"DEFINER",
	"DELAYED",
	"DELAY_KEY_WRITE",
	"DELETE",
	"DESC",
	"DESCRIBE",
	"DES_KEY_FILE",
	"DIAGNOSTICS",
	"DIRECTORY",
	"DISABLE",
	"DISCARD",
	"DISTINCT",
	"DISTINCTROW",
	"DO",
	"DROP",
	"DUAL",
	"DUMPFILE",
	"DUPLICATE",
	"DYNAMIC",
	"ELSE",
	"ELSEIF",
	"ENABLE",
	"ENCLOSED",
	"ENCRYPTION",
	"END",
	"ENDS",
	"ENGINE",
	"ENGINES",
	"ERROR",
	"ERRORS",
	"ESCAPED",
	"EVENT",
	"EVENTS",
	"EVERY",
	"EXCHANGE",
	"EXECUTE",
	"EXISTS",
	"EXIT",
	"EXPIRE",
	"EXPLAIN",
	"EXPORT",
	"EXTENDED",
	"EXTENT_SIZE",
	"FALSE",
	"FAST",
	"FEDERATED",
	"FETCH",
	"FIELDS",
	"FILE",
	"FILE_BLOCK_SIZE",
	"FILTER",
	"FIRST",
	"FIXED",
	"FLOAT4",
	"FLOAT8",
	"FLUSH",
	"FOR",
	"FORCE",
	"FOREIGN",
	"FORMAT",
	"FROM",
	"FULL",
	"FULLTEXT",
	"FUNCTION",
	"GENERAL",
	"GET",
	"GLOBAL",
	"GRANT",
	"GRANTS",
	"GROUP",
	"HANDLER",
	"HAVING",
	"HEAP",
	"HELP",
	"HELP_DATE",
	"HELP_VERSION",
	"HIGH_PRIORITY",
	"HOST",
	"HOSTS",
	"IDENTIFIED",
	"IF",
	"IGNORE",
	"IGNORE_SERVER_IDS",
	"IMPORT",
	"IN",
	"INDEX",
	"INDEXES",
	"INFILE",
	"INITIAL_SIZE",
	"INNER",
	"INNODB",
	"INSERT",
	"INSERT_METHOD",
	"INSTALL",
	"INSTANCE",
	"INT1",
	"INT2",
	"INT3",
	"INT4",
	"INT8",
	"INTEGER",
	"INTERVAL",
	"INTO",
	"IO_THREAD",
	"ISOLATION",
	"ISSUER",
	"ITERATE",
	"JOIN",
	"JSON",
	"KEY",
	"KEYS",
	"KEY_BLOCK_SIZE",
	"KILL",
	"LAST",
	"LEAVE",
	"LEAVES",
	"LEFT",
	"LEVEL",
	"LIKE",
	"LIMIT",
	"LINES",
	"LOAD",
	"LOCAL",
	"LOCK",
	"LOGFILE",
	"LOGS",
	"LONG",
	"LONGBINARY",
	"LOOP",
	"LOW_PRIORITY",
	"MASTER",
	"MASTER_AUTO_POSITION",
	"MASTER_BIND",
	"MASTER_CONNECT_RETRY",
	"MASTER_HEARTBEAT_PERIOD",
	"MASTER_HOST",
	"MASTER_LOG_FILE",
	"MASTER_LOG_POS",
	"MASTER_PASSWORD",
	"MASTER_PORT",
	"MASTER_RETRY_COUNT",
	"MASTER_SSL",
	"MASTER_SSL_CA",
	"MASTER_SSL_CERT",
	"MASTER_SSL_CIPHER",
	"MASTER_SSL_CRL",
	"MASTER_SSL_CRLPATH",
	"MASTER_SSL_KEY",
	"MASTER_SSL_VERIFY_SERVER_CERT",
	"MASTER_TLS_VERSION",
	"MASTER_USER",
	"MAX_CONNECTIONS_PER_HOUR",
	"MAX_QUERIES_PER_HOUR",
	"MAX_ROWS",
	"MAX_SIZE",
	"MAX_UPDATES_PER_HOUR",
	"MAX_USER_CONNECTIONS",
	"MEDIUM",
	"MEMORY",
	"MERGE",
	"MESSAGE_TEXT",
	"MIDDLEINT",
	"MIN_ROWS",
	"MODE",
	"MODIFY",
	"MRG_MYISAM",
	"MUTEX",
	"MYISAM",
	"MYSQL_ERRNO",
	"NAME",
	"NAMES",
	"NATIONAL",
	"NATURAL",
	"NCHAR",
	"NDB",
	"NDBCLUSTER",
	"NEVER",
	"NEXT",
	"NO",
	"NODEGROUP",
	"NONE",
	"NOT",
	"NO_WRITE_TO_BINLOG",
	"NULL",
	"NUMBER",
	"NUMERIC",
	"NVARCHAR",
	"OFFSET",
	"ON",
	"ONLY",
	"OPEN",
	"OPTIMIZE",
	"OPTIMIZER_COSTS",
	"OPTION",
	"OPTIONALLY",
	"OPTIONS",
	"ORDER",
	"OUTER",
	"OUTFILE",
	"OWNER",
	"PACK_KEYS",
	"PARSER",
	"PARTIAL",
	"PARTITION",
	"PARTITIONING",
	"PARTITIONS",
	"PASSWORD",
	"PLUGIN",
	"PLUGINS",
	"PLUGIN_DIR",
	"PORT",
	"PRECISION",
	"PREPARE",
	"PRESERVE",
	"PREV",
	"PRIMARY",
	"PRIVILEGES",
	"PROCEDURE",
	"PROCESS",
	"PROCESSLIST",
	"PROFILE",
	"PROFILES",
	"PROXY",
	"PURGE",
	"QUERY",
	"QUICK",
	"READ",
	"REAL",
	"REBUILD",
	"RECOVER",
	"REDUNDANT",
	"REFERENCES",
	"RELAY",
	"RELAYLOG",
	"RELAY_LOG_FILE",
	"RELAY_LOG_POS",
	"RELEASE",
	"RELOAD",
	"REMOVE",
	"RENAME",
	"REORGANIZE",
	"REPAIR",
	"REPEAT",
	"REPEATABLE",
	"REPLACE",
	"REPLICATE_DO_DB",
	"REPLICATE_DO_TABLE",
	"REPLICATE_IGNORE_DB",
	"REPLICATE_IGNORE_TABLE",
	"REPLICATE_REWRITE_DB",
	"REPLICATE_WILD_DO_TABLE",
	"REPLICATE_WILD_IGNORE_TABLE",
	"REPLICATION",
	"REQUIRE",
	"RESET",
	"RESIGNAL",
	"RESTRICT",
	"RETURN",
	"RETURNED_SQLSTATE",
	"RETURNS",
	"REVOKE",
	"RIGHT",
	"ROLLBACK",
	"ROWS",
	"ROW_COUNT",
	"ROW_FORMAT",
	"SAVEPOINT",
	"SCHEDULE",
	"SCHEMA",
	"SCHEMAS",
	"SCHEMA_NAME",
	"SECURITY",
	"SELECT",
	"SERIAL",
	"SERIALIZABLE",
	"SERVER",
	"SESSION",
	"SET",
	"SHARE",
	"SHOW",
	"SHUTDOWN",
	"SIGNAL",
	"SLAVE",
	"SLOW",
	"SNAPSHOT",
	"SOCKET",
	"SONAME",
	"SPATIAL",
	"SQLSTATE",
	"SQL_AFTER_GTIDS",
	"SQL_AFTER_MTS_GAPS",
	"SQL_BEFORE_GTIDS",
	"SQL_BIG_RESULT",
	"SQL_BUFFER_RESULT",
	"SQL_CACHE",
	"SQL_CALC_FOUND_ROWS",
	"SQL_LOG_BIN",
	"SQL_NO_CACHE",
	"SQL_SLAVE_SKIP_COUNTER",
	"SQL_SMALL_RESULT",
	"SQL_THREAD",
	"SSL",
	"START",
	"STARTING",
	"STARTS",
	"STATS_AUTO_RECALC",
	"STATS_PERSISTENT",
	"STATS_SAMPLE_PAGES",
	"STATUS",
	"STOP",
	"STORAGE",
	"STORED",
	"STRAIGHT_JOIN",
	"STRING",
	"SUBCLASS_ORIGIN",
	"SUBJECT",
	"SUPER",
	"TABLE",
	"TABLES",
	"TABLESPACE",
	"TABLE_NAME",
	"TEMPORARY",
	"TERMINATED",
	"THEN",
	"TIME",
	"TIMESTAMP",
	"TO",
	"TRADITIONAL",
	"TRANSACTION",
	"TRIGGER",
	"TRIGGERS",
	"TRUE",
	"TRUNCATE",
	"TYPE",
	"UNCOMMITTED",
	"UNDO",
	"UNINSTALL",
	"UNION",
	"UNIQUE",
	"UNLOCK",
	"UNSIGNED",
	"UNTIL",
	"UPDATE",
	"UPGRADE",
	"USAGE",
	"USE",
	"USER",
	"USER_RESOURCES",
	"USE_FRM",
	"USING",
	"VALUE",
	"VALUES",
	"VARCHARACTER",
	"VARIABLE",
	"VARIABLES",
	"VARYING",
	"VIEW",
	"VIRTUAL",
	"WAIT",
	"WARNINGS",
	"WHEN",
	"WHERE",
	"WHILE",
	"WITH",
	"WORK",
	"WRAPPER",
	"WRITE",
	"X509",
	"XA",
	"ZEROFILL",
}

var mysql57function = []string{
	"AES_DECRYPT",
	"AES_ENCRYPT",
	"AGAINST",
	"AND",
	"ANY_VALUE",
	"AREA",
	"ASBINARY",
	"ASC",
	"ASTEXT",
	"ASWKB",
	"ASWKT",
	"BETWEEN",
	"BIGINT",
	"BINARY",
	"BOOLEAN",
	"BOTH",
	"BUFFER",
	"BY",
	"CASE",
	"CEIL",
	"CEILING",
	"CENTROID",
	"CHAR",
	"CONTAINS",
	"CONVERT",
	"CONVEXHULL",
	"COUNT",
	"CREATE_DH_PARAMETERS",
	"CROSSES",
	"DATE",
	"DATETIME",
	"DATE_ADD",
	"DATE_SUB",
	"DAY",
	"DAY_HOUR",
	"DAY_MINUTE",
	"DAY_SECOND",
	"DECIMAL",
	"DESC",
	"DIMENSION",
	"DISJOINT",
	"DISTANCE",
	"DISTINCT",
	"ELSE",
	"END",
	"ENDPOINT",
	"ENVELOPE",
	"EQUALS",
	"ESCAPE",
	"EXPANSION",
	"EXTERIORRING",
	"EXTRACTION)",
	"FLOOR",
	"FROM",
	"GEOMCOLLFROMTEXT",
	"GEOMCOLLFROMWKB",
	"GEOMETRYCOLLECTION",
	"GEOMETRYCOLLECTIONFROMTEXT",
	"GEOMETRYCOLLECTIONFROMWKB",
	"GEOMETRYFROMTEXT",
	"GEOMETRYFROMWKB",
	"GEOMETRYN",
	"GEOMETRYTYPE",
	"GEOMFROMTEXT",
	"GEOMFROMWKB",
	"GLENGTH",
	"HOUR",
	"HOUR_MINUTE",
	"HOUR_SECOND",
	"IF",
	"IN",
	"INLINE",
	"INSERT",
	"INTEGER",
	"INTERIORRINGN",
	"INTERSECTS",
	"INTERVAL",
	"IS",
	"ISCLOSED",
	"ISEMPTY",
	"ISSIMPLE",
	"JSON",
	"JSON_APPEND",
	"JSON_ARRAY",
	"JSON_ARRAYAGG",
	"JSON_ARRAY_APPEND",
	"JSON_ARRAY_INSERT",
	"JSON_CONTAINS",
	"JSON_CONTAINS_PATH",
	"JSON_DEPTH",
	"JSON_EXTRACT",
	"JSON_INSERT",
	"JSON_KEYS",
	"JSON_LENGTH",
	"JSON_MERGE",
	"JSON_MERGE_PATCH",
	"JSON_MERGE_PRESERVE",
	"JSON_OBJECT",
	"JSON_OBJECTAGG",
	"JSON_PRETTY",
	"JSON_QUOTE",
	"JSON_REMOVE",
	"JSON_REPLACE",
	"JSON_SEARCH",
	"JSON_SET",
	"JSON_STORAGE_SIZE",
	"JSON_TYPE",
	"JSON_UNQUOTE",
	"JSON_VALID",
	"LEADING",
	"LIKE",
	"LINEFROMTEXT",
	"LINEFROMWKB",
	"LINESTRING",
	"LINESTRINGFROMTEXT",
	"LINESTRINGFROMWKB",
	"MATCH",
	"MBRCONTAINS",
	"MBRDISJOINT",
	"MBREQUAL",
	"MBRINTERSECTS",
	"MBROVERLAPS",
	"MBRTOUCHES",
	"MBRWITHIN",
	"MINUTE",
	"MINUTE_SECOND",
	"MLINEFROMTEXT",
	"MLINEFROMWKB",
	"MOD",
	"MODE",
	"MONTH",
	"MPOINTFROMTEXT",
	"MPOINTFROMWKB",
	"MPOLYFROMTEXT",
	"MPOLYFROMWKB",
	"MULTILINESTRING",
	"MULTILINESTRINGFROMTEXT",
	"MULTILINESTRINGFROMWKB",
	"MULTIPOINT",
	"MULTIPOINTFROMTEXT",
	"MULTIPOINTFROMWKB",
	"MULTIPOLYGON",
	"MULTIPOLYGONFROMTEXT",
	"MULTIPOLYGONFROMWKB",
	"NOT",
	"NULL",
	"NUMGEOMETRIES",
	"NUMINTERIORRINGS",
	"NUMPOINTS",
	"OR",
	"ORDER",
	"OVERLAPS",
	"PATH)",
	"POINT",
	"POINTFROMTEXT",
	"POINTFROMWKB",
	"POINTN",
	"POLYFROMTEXT",
	"POLYFROMWKB",
	"POLYGON",
	"POLYGONFROMTEXT",
	"POLYGONFROMWKB",
	"POW",
	"POWER",
	"QUERY",
	"RANDOM_BYTES",
	"RLIKE",
	"SECOND",
	"SEPARATOR",
	"SHA",
	"SHA1",
	"SHA2",
	"SIGNED",
	"SOUNDS",
	"SRID",
	"STARTPOINT",
	"STD",
	"STDDEV",
	"ST_AREA",
	"ST_ASBINARY",
	"ST_ASGEOJSON",
	"ST_ASTEXT",
	"ST_ASWKB",
	"ST_ASWKT",
	"ST_BUFFER",
	"ST_BUFFER_STRATEGY",
	"ST_CENTROID",
	"ST_CONTAINS",
	"ST_CONVEXHULL",
	"ST_CROSSES",
	"ST_DIFFERENCE",
	"ST_DIMENSION",
	"ST_DISJOINT",
	"ST_DISTANCE",
	"ST_DISTANCE_SPHERE",
	"ST_ENDPOINT",
	"ST_ENVELOPE",
	"ST_EQUALS",
	"ST_EXTERIORRING",
	"ST_GEOHASH",
	"ST_GEOMCOLLFROMTEXT",
	"ST_GEOMCOLLFROMWKB",
	"ST_GEOMETRYCOLLECTIONFROMTEXT",
	"ST_GEOMETRYCOLLECTIONFROMWKB",
	"ST_GEOMETRYFROMTEXT",
	"ST_GEOMETRYFROMWKB",
	"ST_GEOMETRYN",
	"ST_GEOMETRYTYPE",
	"ST_GEOMFROMGEOJSON",
	"ST_GEOMFROMTEXT",
	"ST_GEOMFROMWKB",
	"ST_INTERIORRINGN",
	"ST_INTERSECTION",
	"ST_INTERSECTS",
	"ST_ISCLOSED",
	"ST_ISEMPTY",
	"ST_ISSIMPLE",
	"ST_ISVALID",
	"ST_LATFROMGEOHASH",
	"ST_LINEFROMTEXT",
	"ST_LINEFROMWKB",
	"ST_LINESTRINGFROMTEXT",
	"ST_LINESTRINGFROMWKB",
	"ST_LONGFROMGEOHASH",
	"ST_MAKEENVELOPE",
	"ST_MLINEFROMTEXT",
	"ST_MLINEFROMWKB",
	"ST_MPOINTFROMTEXT",
	"ST_MPOINTFROMWKB",
	"ST_MPOLYFROMTEXT",
	"ST_MPOLYFROMWKB",
	"ST_MULTILINESTRINGFROMTEXT",
	"ST_MULTILINESTRINGFROMWKB",
	"ST_MULTIPOINTFROMTEXT",
	"ST_MULTIPOINTFROMWKB",
	"ST_MULTIPOLYGONFROMTEXT",
	"ST_MULTIPOLYGONFROMWKB",
	"ST_NUMGEOMETRIES",
	"ST_NUMINTERIORRING",
	"ST_NUMINTERIORRINGS",
	"ST_NUMPOINTS",
	"ST_OVERLAPS",
	"ST_POINTFROMGEOHASH",
	"ST_POINTFROMTEXT",
	"ST_POINTFROMWKB",
	"ST_POINTN",
	"ST_POLYFROMTEXT",
	"ST_POLYFROMWKB",
	"ST_POLYGONFROMTEXT",
	"ST_POLYGONFROMWKB",
	"ST_SIMPLIFY",
	"ST_SRID",
	"ST_STARTPOINT",
	"ST_SYMDIFFERENCE",
	"ST_TOUCHES",
	"ST_UNION",
	"ST_VALIDATE",
	"ST_WITHIN",
	"ST_X",
	"ST_Y",
	"THEN",
	"TIME",
	"TIMESTAMP",
	"TOUCHES",
	"TRAILING",
	"UNSIGNED",
	"WHEN",
	"WITH",
	"WITHIN",
	"X",
	"Y",
	"YEAR",
	"YEAR_MONTH",
}

var mysql56Keyword = []string{
	"ACCESSIBLE",
	"ACTION",
	"ADD",
	"AFTER",
	"AGAINST",
	"AGGREGATE",
	"ALGORITHM",
	"ALL",
	"ALTER",
	"ANALYSE",
	"ANALYZE",
	"AND",
	"ANY",
	"AS",
	"ASC",
	"ASCII",
	"ASENSITIVE",
	"AT",
	"AUTHORS",
	"AUTOEXTEND_SIZE",
	"AUTO_INCREMENT",
	"AVG",
	"AVG_ROW_LENGTH",
	"BACKUP",
	"BEFORE",
	"BEGIN",
	"BETWEEN",
	"BIGINT",
	"BINARY",
	"BINLOG",
	"BIT",
	"BLOB",
	"BLOCK",
	"BOOL",
	"BOOLEAN",
	"BOTH",
	"BTREE",
	"BY",
	"BYTE",
	"CACHE",
	"CALL",
	"CASCADE",
	"CASCADED",
	"CASE",
	"CATALOG_NAME",
	"CHAIN",
	"CHANGE",
	"CHANGED",
	"CHAR",
	"CHARACTER",
	"CHARSET",
	"CHECK",
	"CHECKSUM",
	"CIPHER",
	"CLASS_ORIGIN",
	"CLIENT",
	"CLOSE",
	"COALESCE",
	"CODE",
	"COLLATE",
	"COLLATION",
	"COLUMN",
	"COLUMNS",
	"COLUMN_FORMAT",
	"COLUMN_NAME",
	"COMMENT",
	"COMMIT",
	"COMMITTED",
	"COMPACT",
	"COMPLETION",
	"COMPRESSED",
	"CONCURRENT",
	"CONDITION",
	"CONNECTION",
	"CONSISTENT",
	"CONSTRAINT",
	"CONSTRAINT_CATALOG",
	"CONSTRAINT_NAME",
	"CONSTRAINT_SCHEMA",
	"CONTAINS",
	"CONTEXT",
	"CONTINUE",
	"CONTRIBUTORS",
	"CONVERT",
	"CPU",
	"CREATE",
	"CROSS",
	"CUBE",
	"CURRENT",
	"CURRENT_DATE",
	"CURRENT_TIME",
	"CURRENT_TIMESTAMP",
	"CURRENT_USER",
	"CURSOR",
	"CURSOR_NAME",
	"DATA",
	"DATABASE",
	"DATABASES",
	"DATAFILE",
	"DATE",
	"DATETIME",
	"DAY",
	"DAY_HOUR",
	"DAY_MICROSECOND",
	"DAY_MINUTE",
	"DAY_SECOND",
	"DEALLOCATE",
	"DEC",
	"DECIMAL",
	"DECLARE",
	"DEFAULT",
	"DEFAULT_AUTH",
	"DEFINER",
	"DELAYED",
	"DELAY_KEY_WRITE",
	"DELETE",
	"DESC",
	"DESCRIBE",
	"DES_KEY_FILE",
	"DETERMINISTIC",
	"DIAGNOSTICS",
	"DIRECTORY",
	"DISABLE",
	"DISCARD",
	"DISK",
	"DISTINCT",
	"DISTINCTROW",
	"DIV",
	"DO",
	"DOUBLE",
	"DROP",
	"DUAL",
	"DUMPFILE",
	"DUPLICATE",
	"DYNAMIC",
	"EACH",
	"ELSE",
	"ELSEIF",
	"ENABLE",
	"ENCLOSED",
	"END",
	"ENDS",
	"ENGINE",
	"ENGINES",
	"ENUM",
	"ERROR",
	"ERRORS",
	"ESCAPE",
	"ESCAPED",
	"EVENT",
	"EVENTS",
	"EVERY",
	"EXCHANGE",
	"EXECUTE",
	"EXISTS",
	"EXIT",
	"EXPANSION",
	"EXPIRE",
	"EXPLAIN",
	"EXPORT",
	"EXTENDED",
	"EXTENT_SIZE",
	"FALSE",
	"FAST",
	"FAULTS",
	"FETCH",
	"FIELDS",
	"FILE",
	"FIRST",
	"FIXED",
	"FLOAT",
	"FLOAT4",
	"FLOAT8",
	"FLUSH",
	"FOR",
	"FORCE",
	"FOREIGN",
	"FORMAT",
	"FOUND",
	"FROM",
	"FULL",
	"FULLTEXT",
	"FUNCTION",
	"GENERAL",
	"GEOMETRY",
	"GEOMETRYCOLLECTION",
	"GET",
	"GET_FORMAT",
	"GLOBAL",
	"GRANT",
	"GRANTS",
	"GROUP",
	"HANDLER",
	"HASH",
	"HAVING",
	"HELP",
	"HIGH_PRIORITY",
	"HOST",
	"HOSTS",
	"HOUR",
	"HOUR_MICROSECOND",
	"HOUR_MINUTE",
	"HOUR_SECOND",
	"IDENTIFIED",
	"IF",
	"IGNORE",
	"IGNORE_SERVER_IDS",
	"IMPORT",
	"IN",
	"INDEX",
	"INDEXES",
	"INFILE",
	"INITIAL_SIZE",
	"INNER",
	"INOUT",
	"INSENSITIVE",
	"INSERT",
	"INSERT_METHOD",
	"INSTALL",
	"INT",
	"INT1",
	"INT2",
	"INT3",
	"INT4",
	"INT8",
	"INTEGER",
	"INTERVAL",
	"INTO",
	"INVOKER",
	"IO",
	"IO_AFTER_GTIDS",
	"IO_BEFORE_GTIDS",
	"IO_THREAD",
	"IPC",
	"IS",
	"ISOLATION",
	"ISSUER",
	"ITERATE",
	"JOIN",
	"KEY",
	"KEYS",
	"KEY_BLOCK_SIZE",
	"KILL",
	"LANGUAGE",
	"LAST",
	"LEADING",
	"LEAVE",
	"LEAVES",
	"LEFT",
	"LESS",
	"LEVEL",
	"LIKE",
	"LIMIT",
	"LINEAR",
	"LINES",
	"LINESTRING",
	"LIST",
	"LOAD",
	"LOCAL",
	"LOCALTIME",
	"LOCALTIMESTAMP",
	"LOCK",
	"LOCKS",
	"LOGFILE",
	"LOGS",
	"LONG",
	"LONGBLOB",
	"LONGTEXT",
	"LOOP",
	"LOW_PRIORITY",
	"MASTER",
	"MASTER_AUTO_POSITION",
	"MASTER_BIND",
	"MASTER_CONNECT_RETRY",
	"MASTER_DELAY",
	"MASTER_HEARTBEAT_PERIOD",
	"MASTER_HOST",
	"MASTER_LOG_FILE",
	"MASTER_LOG_POS",
	"MASTER_PASSWORD",
	"MASTER_PORT",
	"MASTER_RETRY_COUNT",
	"MASTER_SERVER_ID",
	"MASTER_SSL",
	"MASTER_SSL_CA",
	"MASTER_SSL_CAPATH",
	"MASTER_SSL_CERT",
	"MASTER_SSL_CIPHER",
	"MASTER_SSL_CRL",
	"MASTER_SSL_CRLPATH",
	"MASTER_SSL_KEY",
	"MASTER_SSL_VERIFY_SERVER_CERT",
	"MASTER_USER",
	"MATCH",
	"MAXVALUE",
	"MAX_CONNECTIONS_PER_HOUR",
	"MAX_QUERIES_PER_HOUR",
	"MAX_ROWS",
	"MAX_SIZE",
	"MAX_UPDATES_PER_HOUR",
	"MAX_USER_CONNECTIONS",
	"MEDIUM",
	"MEDIUMBLOB",
	"MEDIUMINT",
	"MEDIUMTEXT",
	"MEMORY",
	"MERGE",
	"MESSAGE_TEXT",
	"MICROSECOND",
	"MIDDLEINT",
	"MIGRATE",
	"MINUTE",
	"MINUTE_MICROSECOND",
	"MINUTE_SECOND",
	"MIN_ROWS",
	"MOD",
	"MODE",
	"MODIFIES",
	"MODIFY",
	"MONTH",
	"MULTILINESTRING",
	"MULTIPOINT",
	"MULTIPOLYGON",
	"MUTEX",
	"MYSQL_ERRNO",
	"NAME",
	"NAMES",
	"NATIONAL",
	"NATURAL",
	"NCHAR",
	"NDB",
	"NDBCLUSTER",
	"NEW",
	"NEXT",
	"NO",
	"NODEGROUP",
	"NONE",
	"NOT",
	"NO_WAIT",
	"NO_WRITE_TO_BINLOG",
	"NULL",
	"NUMBER",
	"NUMERIC",
	"NVARCHAR",
	"OFFSET",
	"OLD_PASSWORD",
	"ON",
	"ONE",
	"ONE_SHOT",
	"ONLY",
	"OPEN",
	"OPTIMIZE",
	"OPTION",
	"OPTIONALLY",
	"OPTIONS",
	"OR",
	"ORDER",
	"OUT",
	"OUTER",
	"OUTFILE",
	"OWNER",
	"PACK_KEYS",
	"PAGE",
	"PARSER",
	"PARTIAL",
	"PARTITION",
	"PARTITIONING",
	"PARTITIONS",
	"PASSWORD",
	"PHASE",
	"PLUGIN",
	"PLUGINS",
	"PLUGIN_DIR",
	"POINT",
	"POLYGON",
	"PORT",
	"PRECISION",
	"PREPARE",
	"PRESERVE",
	"PREV",
	"PRIMARY",
	"PRIVILEGES",
	"PROCEDURE",
	"PROCESSLIST",
	"PROFILE",
	"PROFILES",
	"PROXY",
	"PURGE",
	"QUARTER",
	"QUERY",
	"QUICK",
	"RANGE",
	"READ",
	"READS",
	"READ_ONLY",
	"READ_WRITE",
	"REAL",
	"REBUILD",
	"RECOVER",
	"REDOFILE",
	"REDO_BUFFER_SIZE",
	"REDUNDANT",
	"REFERENCES",
	"REGEXP",
	"RELAY",
	"RELAYLOG",
	"RELAY_LOG_FILE",
	"RELAY_LOG_POS",
	"RELAY_THREAD",
	"RELEASE",
	"RELOAD",
	"REMOVE",
	"RENAME",
	"REORGANIZE",
	"REPAIR",
	"REPEAT",
	"REPEATABLE",
	"REPLACE",
	"REPLICATION",
	"REQUIRE",
	"RESET",
	"RESIGNAL",
	"RESTORE",
	"RESTRICT",
	"RESUME",
	"RETURN",
	"RETURNED_SQLSTATE",
	"RETURNS",
	"REVERSE",
	"REVOKE",
	"RIGHT",
	"RLIKE",
	"ROLLBACK",
	"ROLLUP",
	"ROUTINE",
	"ROW",
	"ROWS",
	"ROW_COUNT",
	"ROW_FORMAT",
	"RTREE",
	"SAVEPOINT",
	"SCHEDULE",
	"SCHEMA",
	"SCHEMAS",
	"SCHEMA_NAME",
	"SECOND",
	"SECOND_MICROSECOND",
	"SECURITY",
	"SELECT",
	"SENSITIVE",
	"SEPARATOR",
	"SERIAL",
	"SERIALIZABLE",
	"SERVER",
	"SESSION",
	"SET",
	"SHARE",
	"SHOW",
	"SHUTDOWN",
	"SIGNAL",
	"SIGNED",
	"SIMPLE",
	"SLAVE",
	"SLOW",
	"SMALLINT",
	"SNAPSHOT",
	"SOCKET",
	"SOME",
	"SONAME",
	"SOUNDS",
	"SOURCE",
	"SPATIAL",
	"SPECIFIC",
	"SQL",
	"SQLEXCEPTION",
	"SQLSTATE",
	"SQLWARNING",
	"SQL_AFTER_GTIDS",
	"SQL_AFTER_MTS_GAPS",
	"SQL_BEFORE_GTIDS",
	"SQL_BIG_RESULT",
	"SQL_BUFFER_RESULT",
	"SQL_CACHE",
	"SQL_CALC_FOUND_ROWS",
	"SQL_NO_CACHE",
	"SQL_SMALL_RESULT",
	"SQL_THREAD",
	"SQL_TSI_DAY",
	"SQL_TSI_HOUR",
	"SQL_TSI_MINUTE",
	"SQL_TSI_MONTH",
	"SQL_TSI_QUARTER",
	"SQL_TSI_SECOND",
	"SQL_TSI_WEEK",
	"SQL_TSI_YEAR",
	"SSL",
	"START",
	"STARTING",
	"STARTS",
	"STATS_AUTO_RECALC",
	"STATS_PERSISTENT",
	"STATS_SAMPLE_PAGES",
	"STATUS",
	"STOP",
	"STORAGE",
	"STRAIGHT_JOIN",
	"STRING",
	"SUBCLASS_ORIGIN",
	"SUBJECT",
	"SUBPARTITION",
	"SUBPARTITIONS",
	"SUPER",
	"SUSPEND",
	"SWAPS",
	"SWITCHES",
	"TABLE",
	"TABLES",
	"TABLESPACE",
	"TABLE_CHECKSUM",
	"TABLE_NAME",
	"TEMPORARY",
	"TEMPTABLE",
	"TERMINATED",
	"TEXT",
	"THAN",
	"THEN",
	"TIME",
	"TIMESTAMP",
	"TIMESTAMPADD",
	"TIMESTAMPDIFF",
	"TINYBLOB",
	"TINYINT",
	"TINYTEXT",
	"TO",
	"TRAILING",
	"TRANSACTION",
	"TRIGGER",
	"TRIGGERS",
	"TRUE",
	"TRUNCATE",
	"TYPE",
	"TYPES",
	"UNCOMMITTED",
	"UNDEFINED",
	"UNDO",
	"UNDOFILE",
	"UNDO_BUFFER_SIZE",
	"UNICODE",
	"UNINSTALL",
	"UNION",
	"UNIQUE",
	"UNKNOWN",
	"UNLOCK",
	"UNSIGNED",
	"UNTIL",
	"UPDATE",
	"UPGRADE",
	"USAGE",
	"USE",
	"USER",
	"USER_RESOURCES",
	"USE_FRM",
	"USING",
	"UTC_DATE",
	"UTC_TIME",
	"UTC_TIMESTAMP",
	"VALUE",
	"VALUES",
	"VARBINARY",
	"VARCHAR",
	"VARCHARACTER",
	"VARIABLES",
	"VARYING",
	"VIEW",
	"WAIT",
	"WARNINGS",
	"WEEK",
	"WEIGHT_STRING",
	"WHEN",
	"WHERE",
	"WHILE",
	"WITH",
	"WORK",
	"WRAPPER",
	"WRITE",
	"X509",
	"XA",
	"XML",
	"XOR",
	"YEAR",
	"YEAR_MONTH",
	"ZEROFILL",
}
