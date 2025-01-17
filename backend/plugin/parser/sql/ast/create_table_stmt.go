package ast

// CreateTableStmt is the struct for create table stmt.
type CreateTableStmt struct {
	ddl

	IfNotExists    bool
	Name           *TableDef
	ColumnList     []*ColumnDef
	ConstraintList []*ConstraintDef

	// TODO(rebelice): convert the partition definition.
	PartitionDef Node
}
