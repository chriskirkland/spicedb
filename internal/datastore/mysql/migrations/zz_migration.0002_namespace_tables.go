package migrations

// namespace max size: https://buf.build/authzed/api/file/main/authzed/api/v0/core.proto#L29
const createNamespaceConfig = `CREATE TABLE namespace_config (
	namespace VARCHAR(128) NOT NULL,
	serialized_config BLOB NOT NULL,
	created_transaction BIGINT NOT NULL,
	deleted_transaction BIGINT NOT NULL DEFAULT '9223372036854775807',
	CONSTRAINT pk_namespace_config PRIMARY KEY (namespace, created_transaction),
	CONSTRAINT uq_namespace_living UNIQUE (namespace, deleted_transaction)
);`

// relationship max size: https://buf.build/authzed/api/file/main/authzed/api/v1/core.proto#L33
// object id max size: https://buf.build/authzed/api/file/main/authzed/api/v1/core.proto#L45
const createRelationTuple = `CREATE TABLE relation_tuple (
		id BIGINT NOT NULL AUTO_INCREMENT,
		namespace VARCHAR(128) NOT NULL,
		object_id VARCHAR(128) NOT NULL,
		relation VARCHAR(64) NOT NULL,
		userset_namespace VARCHAR(128) NOT NULL,
		userset_object_id VARCHAR(128) NOT NULL,
		userset_relation VARCHAR(64) NOT NULL,
		created_transaction BIGINT NOT NULL,
		deleted_transaction BIGINT NOT NULL DEFAULT '9223372036854775807',
		PRIMARY KEY (id),
		CONSTRAINT uq_relation_tuple_namespace UNIQUE (namespace, object_id, relation, userset_namespace, userset_object_id, userset_relation, created_transaction, deleted_transaction),
		CONSTRAINT uq_relation_tuple_living UNIQUE (namespace, object_id, relation, userset_namespace, userset_object_id, userset_relation, deleted_transaction)
	);`

const createRelationTupleTransaction = `CREATE TABLE relation_tuple_transaction (
		id BIGINT NOT NULL AUTO_INCREMENT,
		timestamp DATETIME(6) DEFAULT NOW(6) NOT NULL,
		PRIMARY KEY (id)
);`

const insertFirstTransaction = "INSERT INTO relation_tuple_transaction VALUES();"

func init() {
	err := Manager.Register("namespace-tables", "initial",
		newExecutor(
			createNamespaceConfig,
			createRelationTuple,
			createRelationTupleTransaction,
			insertFirstTransaction,
		).migrate,
	)
	if err != nil {
		panic("failed to register migration  " + err.Error())
	}
}