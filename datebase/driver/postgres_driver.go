package driver

const (
	Postgres            = "postgres"
	PostgresColumnQuery = `
		SELECT
			col_description ( A.attrelid, A.attnum ) AS column_comment,
			format_type ( A.atttypid, A.atttypmod ) AS data_type,
			A.attname AS column_name,
			A.attnotnull AS is_nullable,
			concat_ws ( '', T.typname, SUBSTRING ( format_type ( A.atttypid, A.atttypmod ) FROM '\(.*\)' ) ) AS column_type
		FROM
			pg_class AS C,
			pg_attribute AS A,
			pg_type AS T 
		WHERE
			C.relname = ? 
			AND A.attrelid = C.oid 
			AND A.attnum > 0
			AND A.atttypid = T.oid 
`
)
