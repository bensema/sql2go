package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sql2go"
	"strings"
)

type Config struct {
	Active      int
	Idle        int
	IdleTimeout int

	Username string
	Password string
	Host     string
	Port     int
	Database string
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database)
}

type DB struct {
	*sql.DB
	Conf Config
}

func New(c *Config) (db *DB) {
	db = &DB{
		Conf: *c,
	}
	var err error
	db.DB, err = sql.Open("mysql", c.GetDSN())
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxOpenConns(c.Active)
	db.SetMaxIdleConns(c.Idle)
	return
}

type Table struct {
	Name    string
	Comment string
}

// 表结构详情
type TableDesc struct {
	Index            int
	ColumnName       string // 数据库原始字段
	GoColumnName     string // go使用的字段名称
	OriMysqlType     string // 数据库原始类型
	UpperMysqlType   string // 转换大写的类型
	GolangType       string // 转换成golang类型
	MysqlNullType    string // MYSQL对应的空类型
	PrimaryKey       bool   // 是否是主键
	IsNull           string // 是否为空
	DefaultValue     string // 默认值
	ColumnTypeNumber string // 类型(长度)
	ColumnComment    string // 备注
}

type Columns struct {
	TABLECATALOG    string         `json:"TABLE_CATALOG"`
	TABLESCHEMA     string         `json:"TABLE_SCHEMA"`
	TABLENAME       string         `json:"TABLE_NAME"`
	COLUMNNAME      string         `json:"COLUMN_NAME"`
	ORDINALPOSITION int            `json:"ORDINAL_POSITION"`
	COLUMNDEFAULT   sql.NullString `json:"COLUMN_DEFAULT"`
	ISNULLABLE      string         `json:"IS_NULLABLE"`
	DATATYPE        string         `json:"DATA_TYPE"`
	//CHARACTERMAXIMUMLENGTH int64          `json:"CHARACTER_MAXIMUM_LENGTH"`
	//CHARACTEROCTETLENGTH   int64          `json:"CHARACTER_OCTET_LENGTH"`
	//NUMERICPRECISION       int64          `json:"NUMERIC_PRECISION"`
	//NUMERICSCALE           int64          `json:"NUMERIC_SCALE"`
	//DATETIMEPRECISION      int64          `json:"DATETIME_PRECISION"`
	//CHARACTERSETNAME       string         `json:"CHARACTER_SET_NAME"`
	//COLLATIONNAME          string         `json:"COLLATION_NAME"`
	COLUMNTYPE           string `json:"COLUMN_TYPE"`
	COLUMNKEY            string `json:"COLUMN_KEY"`
	EXTRA                string `json:"EXTRA"`
	PRIVILEGES           string `json:"PRIVILEGES"`
	COLUMNCOMMENT        string `json:"COLUMN_COMMENT"`
	GENERATIONEXPRESSION string `json:"GENERATION_EXPRESSION"`
}

func (m *Columns) Table() string {
	return "columns"
}
func (m *Columns) Columns() []string {
	return []string{"TABLE_CATALOG", "TABLE_SCHEMA", "TABLE_NAME", "COLUMN_NAME", "ORDINAL_POSITION", "COLUMN_DEFAULT", "IS_NULLABLE", "DATA_TYPE", "COLUMN_TYPE", "COLUMN_KEY", "EXTRA", "PRIVILEGES", "COLUMN_COMMENT", "GENERATION_EXPRESSION"}
}
func (m *Columns) Fields() []interface{} {
	return []interface{}{&m.TABLECATALOG, &m.TABLESCHEMA, &m.TABLENAME, &m.COLUMNNAME, &m.ORDINALPOSITION, &m.COLUMNDEFAULT, &m.ISNULLABLE, &m.DATATYPE, &m.COLUMNTYPE, &m.COLUMNKEY, &m.EXTRA, &m.PRIVILEGES, &m.COLUMNCOMMENT, &m.GENERATIONEXPRESSION}
}

func (db *DB) FindTables() ([]*Table, error) {
	rows, err := db.DB.Query("SELECT `table_name`, table_comment FROM information_schema.tables WHERE table_schema = ?", db.Conf.Database)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []*Table
	err = nil
	for rows.Next() {
		result := Table{}
		err = rows.Scan(&result.Name, &result.Comment)
		if err != nil {
			continue
		}
		results = append(results, &result)

	}
	return results, err
}

func (db *DB) GetTableColumns(tableName string) ([]*TableDesc, error) {
	columns := Columns{}
	rows, err := db.DB.Query(fmt.Sprintf("select %s from information_schema.columns where table_name = ? and table_schema = ?", strings.Join(columns.Columns(), ",")), tableName, db.Conf.Database)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []*TableDesc
	for rows.Next() {
		data := Columns{}
		err = rows.Scan(data.Fields()...)
		if err != nil {
			continue
		}
		var keyBool bool
		if strings.ToUpper(data.COLUMNKEY) == "PRI" {
			keyBool = true
		}

		results = append(results, &TableDesc{
			Index:            data.ORDINALPOSITION,
			ColumnName:       data.COLUMNNAME,
			GoColumnName:     sql2go.Capitalize(data.COLUMNNAME),
			OriMysqlType:     data.DATATYPE,
			UpperMysqlType:   strings.ToUpper(data.DATATYPE),
			GolangType:       MysqlTypeToGoType[data.DATATYPE],
			MysqlNullType:    MysqlTypeToGoNullType[data.DATATYPE],
			ColumnComment:    data.COLUMNCOMMENT,
			IsNull:           data.ISNULLABLE,
			DefaultValue:     data.COLUMNDEFAULT.String,
			ColumnTypeNumber: data.COLUMNTYPE,
			PrimaryKey:       keyBool,
		})

	}
	return results, err
}
