// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/G16/app/ent/bookcourse"
	"github.com/G16/app/ent/employee"
	"github.com/G16/app/ent/equipment"
	"github.com/G16/app/ent/equipmentrental"
	"github.com/G16/app/ent/member"
	"github.com/G16/app/ent/promotion"
	"github.com/G16/app/ent/promotionamount"
	"github.com/G16/app/ent/promotiontype"
	"github.com/G16/app/ent/schema"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	bookcourseFields := schema.Bookcourse{}.Fields()
	_ = bookcourseFields
	// bookcourseDescACCESS is the schema descriptor for ACCESS field.
	bookcourseDescACCESS := bookcourseFields[0].Descriptor()
	// bookcourse.ACCESSValidator is a validator for the "ACCESS" field. It is called by the builders before save.
	bookcourse.ACCESSValidator = bookcourseDescACCESS.Validators[0].(func(int) error)
	// bookcourseDescPHONE is the schema descriptor for PHONE field.
	bookcourseDescPHONE := bookcourseFields[1].Descriptor()
	// bookcourse.PHONEValidator is a validator for the "PHONE" field. It is called by the builders before save.
	bookcourse.PHONEValidator = func() func(string) error {
		validators := bookcourseDescPHONE.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(_PHONE string) error {
			for _, fn := range fns {
				if err := fn(_PHONE); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// bookcourseDescDETAIL is the schema descriptor for DETAIL field.
	bookcourseDescDETAIL := bookcourseFields[2].Descriptor()
	// bookcourse.DETAILValidator is a validator for the "DETAIL" field. It is called by the builders before save.
	bookcourse.DETAILValidator = bookcourseDescDETAIL.Validators[0].(func(string) error)
	employeeFields := schema.Employee{}.Fields()
	_ = employeeFields
	// employeeDescEMPLOYEEID is the schema descriptor for EMPLOYEEID field.
	employeeDescEMPLOYEEID := employeeFields[0].Descriptor()
	// employee.EMPLOYEEIDValidator is a validator for the "EMPLOYEEID" field. It is called by the builders before save.
	employee.EMPLOYEEIDValidator = employeeDescEMPLOYEEID.Validators[0].(func(string) error)
	// employeeDescEMPLOYEENAME is the schema descriptor for EMPLOYEENAME field.
	employeeDescEMPLOYEENAME := employeeFields[1].Descriptor()
	// employee.EMPLOYEENAMEValidator is a validator for the "EMPLOYEENAME" field. It is called by the builders before save.
	employee.EMPLOYEENAMEValidator = employeeDescEMPLOYEENAME.Validators[0].(func(string) error)
	// employeeDescEMPLOYEEADDRESS is the schema descriptor for EMPLOYEEADDRESS field.
	employeeDescEMPLOYEEADDRESS := employeeFields[2].Descriptor()
	// employee.EMPLOYEEADDRESSValidator is a validator for the "EMPLOYEEADDRESS" field. It is called by the builders before save.
	employee.EMPLOYEEADDRESSValidator = employeeDescEMPLOYEEADDRESS.Validators[0].(func(string) error)
	// employeeDescIDCARDNUMBER is the schema descriptor for IDCARDNUMBER field.
	employeeDescIDCARDNUMBER := employeeFields[3].Descriptor()
	// employee.IDCARDNUMBERValidator is a validator for the "IDCARDNUMBER" field. It is called by the builders before save.
	employee.IDCARDNUMBERValidator = employeeDescIDCARDNUMBER.Validators[0].(func(string) error)
	equipmentFields := schema.Equipment{}.Fields()
	_ = equipmentFields
	// equipmentDescEQUIPMENTAMOUNT is the schema descriptor for EQUIPMENTAMOUNT field.
	equipmentDescEQUIPMENTAMOUNT := equipmentFields[1].Descriptor()
	// equipment.EQUIPMENTAMOUNTValidator is a validator for the "EQUIPMENTAMOUNT" field. It is called by the builders before save.
	equipment.EQUIPMENTAMOUNTValidator = equipmentDescEQUIPMENTAMOUNT.Validators[0].(func(int) error)
	equipmentrentalFields := schema.Equipmentrental{}.Fields()
	_ = equipmentrentalFields
	// equipmentrentalDescRENTALAMOUNT is the schema descriptor for RENTALAMOUNT field.
	equipmentrentalDescRENTALAMOUNT := equipmentrentalFields[0].Descriptor()
	// equipmentrental.RENTALAMOUNTValidator is a validator for the "RENTALAMOUNT" field. It is called by the builders before save.
	equipmentrental.RENTALAMOUNTValidator = equipmentrentalDescRENTALAMOUNT.Validators[0].(func(int) error)
	memberFields := schema.Member{}.Fields()
	_ = memberFields
	// memberDescMEMBERID is the schema descriptor for MEMBERID field.
	memberDescMEMBERID := memberFields[0].Descriptor()
	// member.MEMBERIDValidator is a validator for the "MEMBERID" field. It is called by the builders before save.
	member.MEMBERIDValidator = memberDescMEMBERID.Validators[0].(func(string) error)
	// memberDescMEMBERNAME is the schema descriptor for MEMBERNAME field.
	memberDescMEMBERNAME := memberFields[1].Descriptor()
	// member.MEMBERNAMEValidator is a validator for the "MEMBERNAME" field. It is called by the builders before save.
	member.MEMBERNAMEValidator = memberDescMEMBERNAME.Validators[0].(func(string) error)
	promotionFields := schema.Promotion{}.Fields()
	_ = promotionFields
	// promotionDescNAME is the schema descriptor for NAME field.
	promotionDescNAME := promotionFields[0].Descriptor()
	// promotion.NAMEValidator is a validator for the "NAME" field. It is called by the builders before save.
	promotion.NAMEValidator = promotionDescNAME.Validators[0].(func(string) error)
	// promotionDescDESC is the schema descriptor for DESC field.
	promotionDescDESC := promotionFields[1].Descriptor()
	// promotion.DESCValidator is a validator for the "DESC" field. It is called by the builders before save.
	promotion.DESCValidator = promotionDescDESC.Validators[0].(func(string) error)
	// promotionDescCODE is the schema descriptor for CODE field.
	promotionDescCODE := promotionFields[2].Descriptor()
	// promotion.CODEValidator is a validator for the "CODE" field. It is called by the builders before save.
	promotion.CODEValidator = promotionDescCODE.Validators[0].(func(string) error)
	promotionamountFields := schema.Promotionamount{}.Fields()
	_ = promotionamountFields
	// promotionamountDescAMOUNT is the schema descriptor for AMOUNT field.
	promotionamountDescAMOUNT := promotionamountFields[0].Descriptor()
	// promotionamount.AMOUNTValidator is a validator for the "AMOUNT" field. It is called by the builders before save.
	promotionamount.AMOUNTValidator = promotionamountDescAMOUNT.Validators[0].(func(int) error)
	promotiontypeFields := schema.Promotiontype{}.Fields()
	_ = promotiontypeFields
	// promotiontypeDescTYPE is the schema descriptor for TYPE field.
	promotiontypeDescTYPE := promotiontypeFields[0].Descriptor()
	// promotiontype.TYPEValidator is a validator for the "TYPE" field. It is called by the builders before save.
	promotiontype.TYPEValidator = promotiontypeDescTYPE.Validators[0].(func(string) error)
}
