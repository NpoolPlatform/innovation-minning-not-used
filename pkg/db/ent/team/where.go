// Code generated by entc, DO NOT EDIT.

package team

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/innovation-mining/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TeamName applies equality check predicate on the "team_name" field. It's identical to TeamNameEQ.
func TeamName(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTeamName), v))
	})
}

// Logo applies equality check predicate on the "logo" field. It's identical to LogoEQ.
func Logo(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogo), v))
	})
}

// LeaderID applies equality check predicate on the "leader_id" field. It's identical to LeaderIDEQ.
func LeaderID(v uuid.UUID) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLeaderID), v))
	})
}

// Introduction applies equality check predicate on the "introduction" field. It's identical to IntroductionEQ.
func Introduction(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIntroduction), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// TeamNameEQ applies the EQ predicate on the "team_name" field.
func TeamNameEQ(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTeamName), v))
	})
}

// TeamNameNEQ applies the NEQ predicate on the "team_name" field.
func TeamNameNEQ(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTeamName), v))
	})
}

// TeamNameIn applies the In predicate on the "team_name" field.
func TeamNameIn(vs ...string) predicate.Team {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Team(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTeamName), v...))
	})
}

// TeamNameNotIn applies the NotIn predicate on the "team_name" field.
func TeamNameNotIn(vs ...string) predicate.Team {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Team(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTeamName), v...))
	})
}

// TeamNameGT applies the GT predicate on the "team_name" field.
func TeamNameGT(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTeamName), v))
	})
}

// TeamNameGTE applies the GTE predicate on the "team_name" field.
func TeamNameGTE(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTeamName), v))
	})
}

// TeamNameLT applies the LT predicate on the "team_name" field.
func TeamNameLT(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTeamName), v))
	})
}

// TeamNameLTE applies the LTE predicate on the "team_name" field.
func TeamNameLTE(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTeamName), v))
	})
}

// TeamNameContains applies the Contains predicate on the "team_name" field.
func TeamNameContains(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTeamName), v))
	})
}

// TeamNameHasPrefix applies the HasPrefix predicate on the "team_name" field.
func TeamNameHasPrefix(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTeamName), v))
	})
}

// TeamNameHasSuffix applies the HasSuffix predicate on the "team_name" field.
func TeamNameHasSuffix(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTeamName), v))
	})
}

// TeamNameEqualFold applies the EqualFold predicate on the "team_name" field.
func TeamNameEqualFold(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTeamName), v))
	})
}

// TeamNameContainsFold applies the ContainsFold predicate on the "team_name" field.
func TeamNameContainsFold(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTeamName), v))
	})
}

// LogoEQ applies the EQ predicate on the "logo" field.
func LogoEQ(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogo), v))
	})
}

// LogoNEQ applies the NEQ predicate on the "logo" field.
func LogoNEQ(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLogo), v))
	})
}

// LogoIn applies the In predicate on the "logo" field.
func LogoIn(vs ...string) predicate.Team {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Team(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLogo), v...))
	})
}

// LogoNotIn applies the NotIn predicate on the "logo" field.
func LogoNotIn(vs ...string) predicate.Team {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Team(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLogo), v...))
	})
}

// LogoGT applies the GT predicate on the "logo" field.
func LogoGT(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLogo), v))
	})
}

// LogoGTE applies the GTE predicate on the "logo" field.
func LogoGTE(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLogo), v))
	})
}

// LogoLT applies the LT predicate on the "logo" field.
func LogoLT(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLogo), v))
	})
}

// LogoLTE applies the LTE predicate on the "logo" field.
func LogoLTE(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLogo), v))
	})
}

// LogoContains applies the Contains predicate on the "logo" field.
func LogoContains(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLogo), v))
	})
}

// LogoHasPrefix applies the HasPrefix predicate on the "logo" field.
func LogoHasPrefix(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLogo), v))
	})
}

// LogoHasSuffix applies the HasSuffix predicate on the "logo" field.
func LogoHasSuffix(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLogo), v))
	})
}

// LogoEqualFold applies the EqualFold predicate on the "logo" field.
func LogoEqualFold(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLogo), v))
	})
}

// LogoContainsFold applies the ContainsFold predicate on the "logo" field.
func LogoContainsFold(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLogo), v))
	})
}

// LeaderIDEQ applies the EQ predicate on the "leader_id" field.
func LeaderIDEQ(v uuid.UUID) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLeaderID), v))
	})
}

// LeaderIDNEQ applies the NEQ predicate on the "leader_id" field.
func LeaderIDNEQ(v uuid.UUID) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLeaderID), v))
	})
}

// LeaderIDIn applies the In predicate on the "leader_id" field.
func LeaderIDIn(vs ...uuid.UUID) predicate.Team {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Team(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLeaderID), v...))
	})
}

// LeaderIDNotIn applies the NotIn predicate on the "leader_id" field.
func LeaderIDNotIn(vs ...uuid.UUID) predicate.Team {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Team(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLeaderID), v...))
	})
}

// LeaderIDGT applies the GT predicate on the "leader_id" field.
func LeaderIDGT(v uuid.UUID) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLeaderID), v))
	})
}

// LeaderIDGTE applies the GTE predicate on the "leader_id" field.
func LeaderIDGTE(v uuid.UUID) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLeaderID), v))
	})
}

// LeaderIDLT applies the LT predicate on the "leader_id" field.
func LeaderIDLT(v uuid.UUID) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLeaderID), v))
	})
}

// LeaderIDLTE applies the LTE predicate on the "leader_id" field.
func LeaderIDLTE(v uuid.UUID) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLeaderID), v))
	})
}

// IntroductionEQ applies the EQ predicate on the "introduction" field.
func IntroductionEQ(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIntroduction), v))
	})
}

// IntroductionNEQ applies the NEQ predicate on the "introduction" field.
func IntroductionNEQ(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIntroduction), v))
	})
}

// IntroductionIn applies the In predicate on the "introduction" field.
func IntroductionIn(vs ...string) predicate.Team {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Team(func(s *sql.Selector) {
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
func IntroductionNotIn(vs ...string) predicate.Team {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Team(func(s *sql.Selector) {
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
func IntroductionGT(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIntroduction), v))
	})
}

// IntroductionGTE applies the GTE predicate on the "introduction" field.
func IntroductionGTE(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIntroduction), v))
	})
}

// IntroductionLT applies the LT predicate on the "introduction" field.
func IntroductionLT(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIntroduction), v))
	})
}

// IntroductionLTE applies the LTE predicate on the "introduction" field.
func IntroductionLTE(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIntroduction), v))
	})
}

// IntroductionContains applies the Contains predicate on the "introduction" field.
func IntroductionContains(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIntroduction), v))
	})
}

// IntroductionHasPrefix applies the HasPrefix predicate on the "introduction" field.
func IntroductionHasPrefix(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIntroduction), v))
	})
}

// IntroductionHasSuffix applies the HasSuffix predicate on the "introduction" field.
func IntroductionHasSuffix(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIntroduction), v))
	})
}

// IntroductionEqualFold applies the EqualFold predicate on the "introduction" field.
func IntroductionEqualFold(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIntroduction), v))
	})
}

// IntroductionContainsFold applies the ContainsFold predicate on the "introduction" field.
func IntroductionContainsFold(v string) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIntroduction), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...uint32) predicate.Team {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Team(func(s *sql.Selector) {
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
func CreateAtNotIn(vs ...uint32) predicate.Team {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Team(func(s *sql.Selector) {
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
func CreateAtGT(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...uint32) predicate.Team {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Team(func(s *sql.Selector) {
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
func UpdateAtNotIn(vs ...uint32) predicate.Team {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Team(func(s *sql.Selector) {
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
func UpdateAtGT(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...uint32) predicate.Team {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Team(func(s *sql.Selector) {
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
func DeleteAtNotIn(vs ...uint32) predicate.Team {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Team(func(s *sql.Selector) {
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
func DeleteAtGT(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v uint32) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Team) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Team) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
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
func Not(p predicate.Team) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		p(s.Not())
	})
}
