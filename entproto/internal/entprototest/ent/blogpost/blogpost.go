// Code generated by ent, DO NOT EDIT.

package blogpost

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the blogpost type in the database.
	Label = "blog_post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldExternalID holds the string denoting the external_id field in the database.
	FieldExternalID = "external_id"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeCategories holds the string denoting the categories edge name in mutations.
	EdgeCategories = "categories"
	// Table holds the table name of the blogpost in the database.
	Table = "blog_posts"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "blog_posts"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "blog_post_author"
	// CategoriesTable is the table that holds the categories relation/edge. The primary key declared below.
	CategoriesTable = "category_blog_posts"
	// CategoriesInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoriesInverseTable = "categories"
)

// Columns holds all SQL columns for blogpost fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldBody,
	FieldExternalID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "blog_posts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"blog_post_author",
}

var (
	// CategoriesPrimaryKey and CategoriesColumn2 are the table columns denoting the
	// primary key for the categories relation (M2M).
	CategoriesPrimaryKey = []string{"category_id", "blog_post_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Order defines the ordering method for the BlogPost queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByBody orders the results by the body field.
func ByBody(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldBody, opts...).ToFunc()
}

// ByExternalID orders the results by the external_id field.
func ByExternalID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldExternalID, opts...).ToFunc()
}

// ByAuthorField orders the results by author field.
func ByAuthorField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorStep(), sql.OrderByField(field, opts...))
	}
}

// ByCategoriesCount orders the results by categories count.
func ByCategoriesCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCategoriesStep(), opts...)
	}
}

// ByCategories orders the results by categories terms.
func ByCategories(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAuthorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, AuthorTable, AuthorColumn),
	)
}
func newCategoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, CategoriesTable, CategoriesPrimaryKey...),
	)
}
