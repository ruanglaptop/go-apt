package fazzdb

import (
	"fmt"
)

type Condition struct {
	Table      string
	Key        string
	Prefix     string
	Operator   Operator
	Connector  Connector
	Conditions []Condition
}

func (c *Condition) QueryString() string {
	if len(c.Conditions) > 0 {
		var query = fmt.Sprintf("%s (", c.Connector)
		for _, cond := range c.Conditions {
			query = fmt.Sprintf("%s %s", query, cond.NamedString())
		}
		query = fmt.Sprintf("%s )", query)
		return query
	}

	return c.NamedString()
}

func (c *Condition) NamedString() string {
	query := ""
	switch c.Operator {
	case OP_IS_NOT_NULL:
		fallthrough
	case OP_IS_NULL:
		query = fmt.Sprintf("%s \"%s\".\"%s\" %s", c.Connector, c.Table, c.Key, c.Operator)
	case OP_IN:
		query = fmt.Sprintf("%s \"%s\".\"%s\" %s (:%s)", c.Connector, c.Table, c.Key, c.Operator, c.Prefix)
	default:
		query = fmt.Sprintf("%s \"%s\".\"%s\" %s :%s", c.Connector, c.Table, c.Key, c.Operator, c.Prefix)
	}
	return query
}