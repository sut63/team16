// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// AgesColumns holds the columns for the "ages" table.
	AgesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt, Unique: true},
	}
	// AgesTable holds the schema information for the "ages" table.
	AgesTable = &schema.Table{
		Name:        "ages",
		Columns:     AgesColumns,
		PrimaryKey:  []*schema.Column{AgesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// BookcoursesColumns holds the columns for the "bookcourses" table.
	BookcoursesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "access", Type: field.TypeInt},
		{Name: "phone", Type: field.TypeString, Size: 10},
		{Name: "detail", Type: field.TypeString, Size: 30},
		{Name: "booktime", Type: field.TypeTime},
		{Name: "course_bookcourse", Type: field.TypeInt, Nullable: true},
		{Name: "employee_bookcourse", Type: field.TypeInt, Nullable: true},
		{Name: "member_bookcourse", Type: field.TypeInt, Nullable: true},
	}
	// BookcoursesTable holds the schema information for the "bookcourses" table.
	BookcoursesTable = &schema.Table{
		Name:       "bookcourses",
		Columns:    BookcoursesColumns,
		PrimaryKey: []*schema.Column{BookcoursesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "bookcourses_courses_bookcourse",
				Columns: []*schema.Column{BookcoursesColumns[5]},

				RefColumns: []*schema.Column{CoursesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bookcourses_employees_bookcourse",
				Columns: []*schema.Column{BookcoursesColumns[6]},

				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bookcourses_members_bookcourse",
				Columns: []*schema.Column{BookcoursesColumns[7]},

				RefColumns: []*schema.Column{MembersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ClassifiersColumns holds the columns for the "classifiers" table.
	ClassifiersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "equipmentclassifier", Type: field.TypeString},
	}
	// ClassifiersTable holds the schema information for the "classifiers" table.
	ClassifiersTable = &schema.Table{
		Name:        "classifiers",
		Columns:     ClassifiersColumns,
		PrimaryKey:  []*schema.Column{ClassifiersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CoursesColumns holds the columns for the "courses" table.
	CoursesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "course", Type: field.TypeString},
	}
	// CoursesTable holds the schema information for the "courses" table.
	CoursesTable = &schema.Table{
		Name:        "courses",
		Columns:     CoursesColumns,
		PrimaryKey:  []*schema.Column{CoursesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// EmployeesColumns holds the columns for the "employees" table.
	EmployeesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "employeeid", Type: field.TypeString, Unique: true},
		{Name: "employeename", Type: field.TypeString},
		{Name: "employeeaddress", Type: field.TypeString},
		{Name: "idcardnumber", Type: field.TypeString, Unique: true, Size: 13},
		{Name: "age_employee", Type: field.TypeInt, Nullable: true},
		{Name: "position_employee", Type: field.TypeInt, Nullable: true},
		{Name: "salary_employee", Type: field.TypeInt, Nullable: true},
	}
	// EmployeesTable holds the schema information for the "employees" table.
	EmployeesTable = &schema.Table{
		Name:       "employees",
		Columns:    EmployeesColumns,
		PrimaryKey: []*schema.Column{EmployeesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "employees_ages_employee",
				Columns: []*schema.Column{EmployeesColumns[5]},

				RefColumns: []*schema.Column{AgesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "employees_positions_employee",
				Columns: []*schema.Column{EmployeesColumns[6]},

				RefColumns: []*schema.Column{PositionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "employees_salaries_employee",
				Columns: []*schema.Column{EmployeesColumns[7]},

				RefColumns: []*schema.Column{SalariesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EquipmentColumns holds the columns for the "equipment" table.
	EquipmentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "equipmentname", Type: field.TypeString},
		{Name: "equipmentamount", Type: field.TypeInt},
		{Name: "equipmentdetail", Type: field.TypeString, Size: 100},
		{Name: "equipmentdate", Type: field.TypeTime},
		{Name: "classifier_equipment", Type: field.TypeInt, Nullable: true},
		{Name: "employee_equipment", Type: field.TypeInt, Nullable: true},
		{Name: "equipmenttype_equipment", Type: field.TypeInt, Nullable: true},
		{Name: "zone_equipment", Type: field.TypeInt, Nullable: true},
	}
	// EquipmentTable holds the schema information for the "equipment" table.
	EquipmentTable = &schema.Table{
		Name:       "equipment",
		Columns:    EquipmentColumns,
		PrimaryKey: []*schema.Column{EquipmentColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "equipment_classifiers_equipment",
				Columns: []*schema.Column{EquipmentColumns[5]},

				RefColumns: []*schema.Column{ClassifiersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "equipment_employees_equipment",
				Columns: []*schema.Column{EquipmentColumns[6]},

				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "equipment_equipmenttypes_equipment",
				Columns: []*schema.Column{EquipmentColumns[7]},

				RefColumns: []*schema.Column{EquipmenttypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "equipment_zones_equipment",
				Columns: []*schema.Column{EquipmentColumns[8]},

				RefColumns: []*schema.Column{ZonesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EquipmentrentalsColumns holds the columns for the "equipmentrentals" table.
	EquipmentrentalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "rentalamount", Type: field.TypeInt},
		{Name: "rentaldate", Type: field.TypeTime},
		{Name: "returndate", Type: field.TypeTime},
		{Name: "employee_equipmentrental", Type: field.TypeInt, Nullable: true},
		{Name: "equipment_equipmentrental", Type: field.TypeInt, Nullable: true},
		{Name: "equipmenttype_equipmentrental", Type: field.TypeInt, Nullable: true},
		{Name: "member_equipmentrental", Type: field.TypeInt, Nullable: true},
	}
	// EquipmentrentalsTable holds the schema information for the "equipmentrentals" table.
	EquipmentrentalsTable = &schema.Table{
		Name:       "equipmentrentals",
		Columns:    EquipmentrentalsColumns,
		PrimaryKey: []*schema.Column{EquipmentrentalsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "equipmentrentals_employees_equipmentrental",
				Columns: []*schema.Column{EquipmentrentalsColumns[4]},

				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "equipmentrentals_equipment_equipmentrental",
				Columns: []*schema.Column{EquipmentrentalsColumns[5]},

				RefColumns: []*schema.Column{EquipmentColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "equipmentrentals_equipmenttypes_equipmentrental",
				Columns: []*schema.Column{EquipmentrentalsColumns[6]},

				RefColumns: []*schema.Column{EquipmenttypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "equipmentrentals_members_equipmentrental",
				Columns: []*schema.Column{EquipmentrentalsColumns[7]},

				RefColumns: []*schema.Column{MembersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EquipmenttypesColumns holds the columns for the "equipmenttypes" table.
	EquipmenttypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "equipmenttype", Type: field.TypeString},
	}
	// EquipmenttypesTable holds the schema information for the "equipmenttypes" table.
	EquipmenttypesTable = &schema.Table{
		Name:        "equipmenttypes",
		Columns:     EquipmenttypesColumns,
		PrimaryKey:  []*schema.Column{EquipmenttypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// MembersColumns holds the columns for the "members" table.
	MembersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "memberid", Type: field.TypeString},
		{Name: "membername", Type: field.TypeString},
	}
	// MembersTable holds the schema information for the "members" table.
	MembersTable = &schema.Table{
		Name:        "members",
		Columns:     MembersColumns,
		PrimaryKey:  []*schema.Column{MembersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PaymentsColumns holds the columns for the "payments" table.
	PaymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "paymentamount", Type: field.TypeString},
		{Name: "phonenumber", Type: field.TypeString, Size: 10},
		{Name: "email", Type: field.TypeString},
		{Name: "paymentdate", Type: field.TypeTime},
		{Name: "employee_payment", Type: field.TypeInt, Nullable: true},
		{Name: "member_payment", Type: field.TypeInt, Nullable: true},
		{Name: "paymenttype_payment", Type: field.TypeInt, Nullable: true},
		{Name: "promotion_payment", Type: field.TypeInt, Nullable: true},
	}
	// PaymentsTable holds the schema information for the "payments" table.
	PaymentsTable = &schema.Table{
		Name:       "payments",
		Columns:    PaymentsColumns,
		PrimaryKey: []*schema.Column{PaymentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "payments_employees_payment",
				Columns: []*schema.Column{PaymentsColumns[5]},

				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "payments_members_payment",
				Columns: []*schema.Column{PaymentsColumns[6]},

				RefColumns: []*schema.Column{MembersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "payments_paymenttypes_payment",
				Columns: []*schema.Column{PaymentsColumns[7]},

				RefColumns: []*schema.Column{PaymenttypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "payments_promotions_payment",
				Columns: []*schema.Column{PaymentsColumns[8]},

				RefColumns: []*schema.Column{PromotionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PaymenttypesColumns holds the columns for the "paymenttypes" table.
	PaymenttypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString},
	}
	// PaymenttypesTable holds the schema information for the "paymenttypes" table.
	PaymenttypesTable = &schema.Table{
		Name:        "paymenttypes",
		Columns:     PaymenttypesColumns,
		PrimaryKey:  []*schema.Column{PaymenttypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PositionsColumns holds the columns for the "positions" table.
	PositionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "position", Type: field.TypeString, Unique: true},
	}
	// PositionsTable holds the schema information for the "positions" table.
	PositionsTable = &schema.Table{
		Name:        "positions",
		Columns:     PositionsColumns,
		PrimaryKey:  []*schema.Column{PositionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PromotionsColumns holds the columns for the "promotions" table.
	PromotionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeString, Size: 30},
		{Name: "code", Type: field.TypeString, Size: 5},
		{Name: "date", Type: field.TypeTime},
		{Name: "employee_promotion", Type: field.TypeInt, Nullable: true},
		{Name: "promotionamount_promotion", Type: field.TypeInt, Nullable: true},
		{Name: "promotiontype_promotion", Type: field.TypeInt, Nullable: true},
	}
	// PromotionsTable holds the schema information for the "promotions" table.
	PromotionsTable = &schema.Table{
		Name:       "promotions",
		Columns:    PromotionsColumns,
		PrimaryKey: []*schema.Column{PromotionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "promotions_employees_promotion",
				Columns: []*schema.Column{PromotionsColumns[5]},

				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "promotions_promotionamounts_promotion",
				Columns: []*schema.Column{PromotionsColumns[6]},

				RefColumns: []*schema.Column{PromotionamountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "promotions_promotiontypes_promotion",
				Columns: []*schema.Column{PromotionsColumns[7]},

				RefColumns: []*schema.Column{PromotiontypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PromotionamountsColumns holds the columns for the "promotionamounts" table.
	PromotionamountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "amount", Type: field.TypeInt},
	}
	// PromotionamountsTable holds the schema information for the "promotionamounts" table.
	PromotionamountsTable = &schema.Table{
		Name:        "promotionamounts",
		Columns:     PromotionamountsColumns,
		PrimaryKey:  []*schema.Column{PromotionamountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PromotiontypesColumns holds the columns for the "promotiontypes" table.
	PromotiontypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString},
	}
	// PromotiontypesTable holds the schema information for the "promotiontypes" table.
	PromotiontypesTable = &schema.Table{
		Name:        "promotiontypes",
		Columns:     PromotiontypesColumns,
		PrimaryKey:  []*schema.Column{PromotiontypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SalariesColumns holds the columns for the "salaries" table.
	SalariesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "salary", Type: field.TypeInt, Unique: true},
	}
	// SalariesTable holds the schema information for the "salaries" table.
	SalariesTable = &schema.Table{
		Name:        "salaries",
		Columns:     SalariesColumns,
		PrimaryKey:  []*schema.Column{SalariesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ZonesColumns holds the columns for the "zones" table.
	ZonesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "equipmentzone", Type: field.TypeString},
	}
	// ZonesTable holds the schema information for the "zones" table.
	ZonesTable = &schema.Table{
		Name:        "zones",
		Columns:     ZonesColumns,
		PrimaryKey:  []*schema.Column{ZonesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AgesTable,
		BookcoursesTable,
		ClassifiersTable,
		CoursesTable,
		EmployeesTable,
		EquipmentTable,
		EquipmentrentalsTable,
		EquipmenttypesTable,
		MembersTable,
		PaymentsTable,
		PaymenttypesTable,
		PositionsTable,
		PromotionsTable,
		PromotionamountsTable,
		PromotiontypesTable,
		SalariesTable,
		UsersTable,
		ZonesTable,
	}
)

func init() {
	BookcoursesTable.ForeignKeys[0].RefTable = CoursesTable
	BookcoursesTable.ForeignKeys[1].RefTable = EmployeesTable
	BookcoursesTable.ForeignKeys[2].RefTable = MembersTable
	EmployeesTable.ForeignKeys[0].RefTable = AgesTable
	EmployeesTable.ForeignKeys[1].RefTable = PositionsTable
	EmployeesTable.ForeignKeys[2].RefTable = SalariesTable
	EquipmentTable.ForeignKeys[0].RefTable = ClassifiersTable
	EquipmentTable.ForeignKeys[1].RefTable = EmployeesTable
	EquipmentTable.ForeignKeys[2].RefTable = EquipmenttypesTable
	EquipmentTable.ForeignKeys[3].RefTable = ZonesTable
	EquipmentrentalsTable.ForeignKeys[0].RefTable = EmployeesTable
	EquipmentrentalsTable.ForeignKeys[1].RefTable = EquipmentTable
	EquipmentrentalsTable.ForeignKeys[2].RefTable = EquipmenttypesTable
	EquipmentrentalsTable.ForeignKeys[3].RefTable = MembersTable
	PaymentsTable.ForeignKeys[0].RefTable = EmployeesTable
	PaymentsTable.ForeignKeys[1].RefTable = MembersTable
	PaymentsTable.ForeignKeys[2].RefTable = PaymenttypesTable
	PaymentsTable.ForeignKeys[3].RefTable = PromotionsTable
	PromotionsTable.ForeignKeys[0].RefTable = EmployeesTable
	PromotionsTable.ForeignKeys[1].RefTable = PromotionamountsTable
	PromotionsTable.ForeignKeys[2].RefTable = PromotiontypesTable
}
