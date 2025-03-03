// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/Varissara/app/ent/educationlevel"
	"github.com/Varissara/app/ent/schema"
	"github.com/Varissara/app/ent/scholarship"
	"github.com/Varissara/app/ent/scholarshiptype"
	"github.com/Varissara/app/ent/semester"
	"github.com/Varissara/app/ent/studyplan"
	"github.com/Varissara/app/ent/user"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	educationlevelFields := schema.Educationlevel{}.Fields()
	_ = educationlevelFields
	// educationlevelDescEducationlevelname is the schema descriptor for Educationlevelname field.
	educationlevelDescEducationlevelname := educationlevelFields[0].Descriptor()
	// educationlevel.EducationlevelnameValidator is a validator for the "Educationlevelname" field. It is called by the builders before save.
	educationlevel.EducationlevelnameValidator = educationlevelDescEducationlevelname.Validators[0].(func(string) error)
	scholarshipFields := schema.Scholarship{}.Fields()
	_ = scholarshipFields
	// scholarshipDescSchlolarshipname is the schema descriptor for Schlolarshipname field.
	scholarshipDescSchlolarshipname := scholarshipFields[0].Descriptor()
	// scholarship.SchlolarshipnameValidator is a validator for the "Schlolarshipname" field. It is called by the builders before save.
	scholarship.SchlolarshipnameValidator = scholarshipDescSchlolarshipname.Validators[0].(func(string) error)
	// scholarshipDescOrganization is the schema descriptor for organization field.
	scholarshipDescOrganization := scholarshipFields[1].Descriptor()
	// scholarship.OrganizationValidator is a validator for the "organization" field. It is called by the builders before save.
	scholarship.OrganizationValidator = scholarshipDescOrganization.Validators[0].(func(string) error)
	scholarshiptypeFields := schema.Scholarshiptype{}.Fields()
	_ = scholarshiptypeFields
	// scholarshiptypeDescScholarshiptypename is the schema descriptor for Scholarshiptypename field.
	scholarshiptypeDescScholarshiptypename := scholarshiptypeFields[0].Descriptor()
	// scholarshiptype.ScholarshiptypenameValidator is a validator for the "Scholarshiptypename" field. It is called by the builders before save.
	scholarshiptype.ScholarshiptypenameValidator = scholarshiptypeDescScholarshiptypename.Validators[0].(func(string) error)
	semesterFields := schema.Semester{}.Fields()
	_ = semesterFields
	// semesterDescSemestername is the schema descriptor for Semestername field.
	semesterDescSemestername := semesterFields[0].Descriptor()
	// semester.SemesternameValidator is a validator for the "Semestername" field. It is called by the builders before save.
	semester.SemesternameValidator = semesterDescSemestername.Validators[0].(func(string) error)
	studyplanFields := schema.Studyplan{}.Fields()
	_ = studyplanFields
	// studyplanDescStudyplanname is the schema descriptor for Studyplanname field.
	studyplanDescStudyplanname := studyplanFields[0].Descriptor()
	// studyplan.StudyplannameValidator is a validator for the "Studyplanname" field. It is called by the builders before save.
	studyplan.StudyplannameValidator = studyplanDescStudyplanname.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for Username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "Username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescUseremail is the schema descriptor for Useremail field.
	userDescUseremail := userFields[1].Descriptor()
	// user.UseremailValidator is a validator for the "Useremail" field. It is called by the builders before save.
	user.UseremailValidator = userDescUseremail.Validators[0].(func(string) error)
}
