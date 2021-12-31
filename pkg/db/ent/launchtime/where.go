// Code generated by entc, DO NOT EDIT.

package launchtime

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/innovation-mining/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ProjectID applies equality check predicate on the "project_id" field. It's identical to ProjectIDEQ.
func ProjectID(v uuid.UUID) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProjectID), v))
	})
}

// NetworkName applies equality check predicate on the "network_name" field. It's identical to NetworkNameEQ.
func NetworkName(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNetworkName), v))
	})
}

// NetworkVersion applies equality check predicate on the "network_version" field. It's identical to NetworkVersionEQ.
func NetworkVersion(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNetworkVersion), v))
	})
}

// Introduction applies equality check predicate on the "introduction" field. It's identical to IntroductionEQ.
func Introduction(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIntroduction), v))
	})
}

// LaunchTime applies equality check predicate on the "launch_time" field. It's identical to LaunchTimeEQ.
func LaunchTime(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLaunchTime), v))
	})
}

// Incentive applies equality check predicate on the "incentive" field. It's identical to IncentiveEQ.
func Incentive(v bool) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIncentive), v))
	})
}

// IncentiveTotal applies equality check predicate on the "incentive_total" field. It's identical to IncentiveTotalEQ.
func IncentiveTotal(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIncentiveTotal), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// ProjectIDEQ applies the EQ predicate on the "project_id" field.
func ProjectIDEQ(v uuid.UUID) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProjectID), v))
	})
}

// ProjectIDNEQ applies the NEQ predicate on the "project_id" field.
func ProjectIDNEQ(v uuid.UUID) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProjectID), v))
	})
}

// ProjectIDIn applies the In predicate on the "project_id" field.
func ProjectIDIn(vs ...uuid.UUID) predicate.LaunchTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProjectID), v...))
	})
}

// ProjectIDNotIn applies the NotIn predicate on the "project_id" field.
func ProjectIDNotIn(vs ...uuid.UUID) predicate.LaunchTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProjectID), v...))
	})
}

// ProjectIDGT applies the GT predicate on the "project_id" field.
func ProjectIDGT(v uuid.UUID) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProjectID), v))
	})
}

// ProjectIDGTE applies the GTE predicate on the "project_id" field.
func ProjectIDGTE(v uuid.UUID) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProjectID), v))
	})
}

// ProjectIDLT applies the LT predicate on the "project_id" field.
func ProjectIDLT(v uuid.UUID) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProjectID), v))
	})
}

// ProjectIDLTE applies the LTE predicate on the "project_id" field.
func ProjectIDLTE(v uuid.UUID) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProjectID), v))
	})
}

// NetworkNameEQ applies the EQ predicate on the "network_name" field.
func NetworkNameEQ(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNetworkName), v))
	})
}

// NetworkNameNEQ applies the NEQ predicate on the "network_name" field.
func NetworkNameNEQ(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNetworkName), v))
	})
}

// NetworkNameIn applies the In predicate on the "network_name" field.
func NetworkNameIn(vs ...string) predicate.LaunchTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNetworkName), v...))
	})
}

// NetworkNameNotIn applies the NotIn predicate on the "network_name" field.
func NetworkNameNotIn(vs ...string) predicate.LaunchTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNetworkName), v...))
	})
}

// NetworkNameGT applies the GT predicate on the "network_name" field.
func NetworkNameGT(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNetworkName), v))
	})
}

// NetworkNameGTE applies the GTE predicate on the "network_name" field.
func NetworkNameGTE(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNetworkName), v))
	})
}

// NetworkNameLT applies the LT predicate on the "network_name" field.
func NetworkNameLT(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNetworkName), v))
	})
}

// NetworkNameLTE applies the LTE predicate on the "network_name" field.
func NetworkNameLTE(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNetworkName), v))
	})
}

// NetworkNameContains applies the Contains predicate on the "network_name" field.
func NetworkNameContains(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNetworkName), v))
	})
}

// NetworkNameHasPrefix applies the HasPrefix predicate on the "network_name" field.
func NetworkNameHasPrefix(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNetworkName), v))
	})
}

// NetworkNameHasSuffix applies the HasSuffix predicate on the "network_name" field.
func NetworkNameHasSuffix(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNetworkName), v))
	})
}

// NetworkNameEqualFold applies the EqualFold predicate on the "network_name" field.
func NetworkNameEqualFold(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNetworkName), v))
	})
}

// NetworkNameContainsFold applies the ContainsFold predicate on the "network_name" field.
func NetworkNameContainsFold(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNetworkName), v))
	})
}

// NetworkVersionEQ applies the EQ predicate on the "network_version" field.
func NetworkVersionEQ(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNetworkVersion), v))
	})
}

// NetworkVersionNEQ applies the NEQ predicate on the "network_version" field.
func NetworkVersionNEQ(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNetworkVersion), v))
	})
}

// NetworkVersionIn applies the In predicate on the "network_version" field.
func NetworkVersionIn(vs ...string) predicate.LaunchTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNetworkVersion), v...))
	})
}

// NetworkVersionNotIn applies the NotIn predicate on the "network_version" field.
func NetworkVersionNotIn(vs ...string) predicate.LaunchTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNetworkVersion), v...))
	})
}

// NetworkVersionGT applies the GT predicate on the "network_version" field.
func NetworkVersionGT(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNetworkVersion), v))
	})
}

// NetworkVersionGTE applies the GTE predicate on the "network_version" field.
func NetworkVersionGTE(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNetworkVersion), v))
	})
}

// NetworkVersionLT applies the LT predicate on the "network_version" field.
func NetworkVersionLT(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNetworkVersion), v))
	})
}

// NetworkVersionLTE applies the LTE predicate on the "network_version" field.
func NetworkVersionLTE(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNetworkVersion), v))
	})
}

// NetworkVersionContains applies the Contains predicate on the "network_version" field.
func NetworkVersionContains(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNetworkVersion), v))
	})
}

// NetworkVersionHasPrefix applies the HasPrefix predicate on the "network_version" field.
func NetworkVersionHasPrefix(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNetworkVersion), v))
	})
}

// NetworkVersionHasSuffix applies the HasSuffix predicate on the "network_version" field.
func NetworkVersionHasSuffix(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNetworkVersion), v))
	})
}

// NetworkVersionEqualFold applies the EqualFold predicate on the "network_version" field.
func NetworkVersionEqualFold(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNetworkVersion), v))
	})
}

// NetworkVersionContainsFold applies the ContainsFold predicate on the "network_version" field.
func NetworkVersionContainsFold(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNetworkVersion), v))
	})
}

// IntroductionEQ applies the EQ predicate on the "introduction" field.
func IntroductionEQ(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIntroduction), v))
	})
}

// IntroductionNEQ applies the NEQ predicate on the "introduction" field.
func IntroductionNEQ(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIntroduction), v))
	})
}

// IntroductionIn applies the In predicate on the "introduction" field.
func IntroductionIn(vs ...string) predicate.LaunchTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIntroduction), v...))
	})
}

// IntroductionNotIn applies the NotIn predicate on the "introduction" field.
func IntroductionNotIn(vs ...string) predicate.LaunchTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIntroduction), v...))
	})
}

// IntroductionGT applies the GT predicate on the "introduction" field.
func IntroductionGT(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIntroduction), v))
	})
}

// IntroductionGTE applies the GTE predicate on the "introduction" field.
func IntroductionGTE(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIntroduction), v))
	})
}

// IntroductionLT applies the LT predicate on the "introduction" field.
func IntroductionLT(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIntroduction), v))
	})
}

// IntroductionLTE applies the LTE predicate on the "introduction" field.
func IntroductionLTE(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIntroduction), v))
	})
}

// IntroductionContains applies the Contains predicate on the "introduction" field.
func IntroductionContains(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIntroduction), v))
	})
}

// IntroductionHasPrefix applies the HasPrefix predicate on the "introduction" field.
func IntroductionHasPrefix(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIntroduction), v))
	})
}

// IntroductionHasSuffix applies the HasSuffix predicate on the "introduction" field.
func IntroductionHasSuffix(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIntroduction), v))
	})
}

// IntroductionEqualFold applies the EqualFold predicate on the "introduction" field.
func IntroductionEqualFold(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIntroduction), v))
	})
}

// IntroductionContainsFold applies the ContainsFold predicate on the "introduction" field.
func IntroductionContainsFold(v string) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIntroduction), v))
	})
}

// LaunchTimeEQ applies the EQ predicate on the "launch_time" field.
func LaunchTimeEQ(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLaunchTime), v))
	})
}

// LaunchTimeNEQ applies the NEQ predicate on the "launch_time" field.
func LaunchTimeNEQ(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLaunchTime), v))
	})
}

// LaunchTimeIn applies the In predicate on the "launch_time" field.
func LaunchTimeIn(vs ...uint32) predicate.LaunchTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLaunchTime), v...))
	})
}

// LaunchTimeNotIn applies the NotIn predicate on the "launch_time" field.
func LaunchTimeNotIn(vs ...uint32) predicate.LaunchTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLaunchTime), v...))
	})
}

// LaunchTimeGT applies the GT predicate on the "launch_time" field.
func LaunchTimeGT(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLaunchTime), v))
	})
}

// LaunchTimeGTE applies the GTE predicate on the "launch_time" field.
func LaunchTimeGTE(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLaunchTime), v))
	})
}

// LaunchTimeLT applies the LT predicate on the "launch_time" field.
func LaunchTimeLT(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLaunchTime), v))
	})
}

// LaunchTimeLTE applies the LTE predicate on the "launch_time" field.
func LaunchTimeLTE(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLaunchTime), v))
	})
}

// IncentiveEQ applies the EQ predicate on the "incentive" field.
func IncentiveEQ(v bool) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIncentive), v))
	})
}

// IncentiveNEQ applies the NEQ predicate on the "incentive" field.
func IncentiveNEQ(v bool) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIncentive), v))
	})
}

// IncentiveTotalEQ applies the EQ predicate on the "incentive_total" field.
func IncentiveTotalEQ(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIncentiveTotal), v))
	})
}

// IncentiveTotalNEQ applies the NEQ predicate on the "incentive_total" field.
func IncentiveTotalNEQ(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIncentiveTotal), v))
	})
}

// IncentiveTotalIn applies the In predicate on the "incentive_total" field.
func IncentiveTotalIn(vs ...uint32) predicate.LaunchTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIncentiveTotal), v...))
	})
}

// IncentiveTotalNotIn applies the NotIn predicate on the "incentive_total" field.
func IncentiveTotalNotIn(vs ...uint32) predicate.LaunchTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIncentiveTotal), v...))
	})
}

// IncentiveTotalGT applies the GT predicate on the "incentive_total" field.
func IncentiveTotalGT(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIncentiveTotal), v))
	})
}

// IncentiveTotalGTE applies the GTE predicate on the "incentive_total" field.
func IncentiveTotalGTE(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIncentiveTotal), v))
	})
}

// IncentiveTotalLT applies the LT predicate on the "incentive_total" field.
func IncentiveTotalLT(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIncentiveTotal), v))
	})
}

// IncentiveTotalLTE applies the LTE predicate on the "incentive_total" field.
func IncentiveTotalLTE(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIncentiveTotal), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...uint32) predicate.LaunchTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...uint32) predicate.LaunchTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...uint32) predicate.LaunchTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...uint32) predicate.LaunchTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...uint32) predicate.LaunchTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...uint32) predicate.LaunchTime {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LaunchTime(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v uint32) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LaunchTime) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LaunchTime) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.LaunchTime) predicate.LaunchTime {
	return predicate.LaunchTime(func(s *sql.Selector) {
		p(s.Not())
	})
}
