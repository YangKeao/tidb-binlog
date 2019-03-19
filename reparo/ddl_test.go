package reparo

import (
	"github.com/pingcap/check"
	"github.com/pingcap/tidb-binlog/pkg/filter"
)

type testDDLSuite struct{}

var _ = check.Suite(&testDDLSuite{})

func (s *testDDLSuite) TestParseDDL(c *check.C) {
	tests := map[string]filter.TableName{
		"create database db1": filter.TableName{"db1", ""},
		"drop database db1":   filter.TableName{"db1", ""},

		"use db1; create table table1(id int)": filter.TableName{"db1", "table1"},
		"create table table1(id int)":          filter.TableName{"", "table1"},

		"use db1; drop table table1": filter.TableName{"db1", "table1"},
		"drop table table1":          filter.TableName{"", "table1"},

		"use db1; alter table table1 drop column v1": filter.TableName{"db1", "table1"},
		"alter table table1 drop column v1":          filter.TableName{"", "table1"},

		"use db1; truncate table table1": filter.TableName{"db1", "table1"},
		"truncate table table1":          filter.TableName{"", "table1"},

		"use db1; create index idx on table1(id)": filter.TableName{"db1", "table1"},
		"create index idx on table1(id)":          filter.TableName{"", "table1"},

		"use db1; alter table table1 drop index index_name": filter.TableName{"db1", "table1"},
		"alter table table1 drop index index_name":          filter.TableName{"", "table1"},

		"use db1;rename table table1 to table2": filter.TableName{"db1", "table1"},
		"rename table table1 to table2":         filter.TableName{"", "table1"},
	}

	for sql, table := range tests {
		_, parseTable, err := parseDDL(sql)
		c.Assert(err, check.IsNil)
		c.Assert(parseTable, check.DeepEquals, table, check.Commentf("sql: %s", sql))
	}
}
