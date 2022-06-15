package mysql

import (
	"testing"

	"github.com/bytebase/bytebase/common"
	"github.com/bytebase/bytebase/plugin/advisor"
)

func TestWhereRequirement(t *testing.T) {
	tests := []test{
		{
			statement: "DELETE FROM t1",
			want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    common.StatementNoWhere,
					Title:   "statement.where.require",
					Content: "\"DELETE FROM t1\" requires WHERE clause",
				},
			},
		},
		{
			statement: "UPDATE t1 SET a = 1",
			want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    common.StatementNoWhere,
					Title:   "statement.where.require",
					Content: "\"UPDATE t1 SET a = 1\" requires WHERE clause",
				},
			},
		},
		{
			statement: "DELETE FROM t1 WHERE a > 0",
			want: []advisor.Advice{
				{
					Status:  advisor.Success,
					Code:    common.Ok,
					Title:   "OK",
					Content: "",
				},
			},
		},
		{
			statement: "UPDATE t1 SET a = 1 WHERE a > 10",
			want: []advisor.Advice{
				{
					Status:  advisor.Success,
					Code:    common.Ok,
					Title:   "OK",
					Content: "",
				},
			},
		},
		{
			statement: "SELECT a FROM t",
			want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    common.StatementNoWhere,
					Title:   "statement.where.require",
					Content: "\"SELECT a FROM t\" requires WHERE clause",
				},
			},
		},
		{
			statement: "SELECT a FROM t WHERE a > 0",
			want: []advisor.Advice{
				{
					Status:  advisor.Success,
					Code:    common.Ok,
					Title:   "OK",
					Content: "",
				},
			},
		},
		{
			statement: "SELECT a FROM t WHERE a > (SELECT max(id) FROM user)",
			want: []advisor.Advice{
				{
					Status:  advisor.Warn,
					Code:    common.StatementNoWhere,
					Title:   "statement.where.require",
					Content: "\"SELECT a FROM t WHERE a > (SELECT max(id) FROM user)\" requires WHERE clause",
				},
			},
		},
	}

	runSchemaReviewRuleTests(t, tests, &WhereRequirementAdvisor{}, &advisor.SchemaReviewRule{
		Type:    advisor.SchemaRuleStatementRequireWhere,
		Level:   advisor.SchemaRuleLevelWarning,
		Payload: "",
	}, &MockCatalogService{})
}
