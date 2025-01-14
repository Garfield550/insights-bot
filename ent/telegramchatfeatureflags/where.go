// Code generated by ent, DO NOT EDIT.

package telegramchatfeatureflags

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/nekomeowww/insights-bot/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldLTE(FieldID, id))
}

// ChatID applies equality check predicate on the "chat_id" field. It's identical to ChatIDEQ.
func ChatID(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldEQ(FieldChatID, v))
}

// ChatType applies equality check predicate on the "chat_type" field. It's identical to ChatTypeEQ.
func ChatType(v string) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldEQ(FieldChatType, v))
}

// FeatureChatHistoriesRecap applies equality check predicate on the "feature_chat_histories_recap" field. It's identical to FeatureChatHistoriesRecapEQ.
func FeatureChatHistoriesRecap(v bool) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldEQ(FieldFeatureChatHistoriesRecap, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldEQ(FieldUpdatedAt, v))
}

// ChatIDEQ applies the EQ predicate on the "chat_id" field.
func ChatIDEQ(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldEQ(FieldChatID, v))
}

// ChatIDNEQ applies the NEQ predicate on the "chat_id" field.
func ChatIDNEQ(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldNEQ(FieldChatID, v))
}

// ChatIDIn applies the In predicate on the "chat_id" field.
func ChatIDIn(vs ...int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldIn(FieldChatID, vs...))
}

// ChatIDNotIn applies the NotIn predicate on the "chat_id" field.
func ChatIDNotIn(vs ...int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldNotIn(FieldChatID, vs...))
}

// ChatIDGT applies the GT predicate on the "chat_id" field.
func ChatIDGT(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldGT(FieldChatID, v))
}

// ChatIDGTE applies the GTE predicate on the "chat_id" field.
func ChatIDGTE(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldGTE(FieldChatID, v))
}

// ChatIDLT applies the LT predicate on the "chat_id" field.
func ChatIDLT(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldLT(FieldChatID, v))
}

// ChatIDLTE applies the LTE predicate on the "chat_id" field.
func ChatIDLTE(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldLTE(FieldChatID, v))
}

// ChatTypeEQ applies the EQ predicate on the "chat_type" field.
func ChatTypeEQ(v string) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldEQ(FieldChatType, v))
}

// ChatTypeNEQ applies the NEQ predicate on the "chat_type" field.
func ChatTypeNEQ(v string) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldNEQ(FieldChatType, v))
}

// ChatTypeIn applies the In predicate on the "chat_type" field.
func ChatTypeIn(vs ...string) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldIn(FieldChatType, vs...))
}

// ChatTypeNotIn applies the NotIn predicate on the "chat_type" field.
func ChatTypeNotIn(vs ...string) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldNotIn(FieldChatType, vs...))
}

// ChatTypeGT applies the GT predicate on the "chat_type" field.
func ChatTypeGT(v string) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldGT(FieldChatType, v))
}

// ChatTypeGTE applies the GTE predicate on the "chat_type" field.
func ChatTypeGTE(v string) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldGTE(FieldChatType, v))
}

// ChatTypeLT applies the LT predicate on the "chat_type" field.
func ChatTypeLT(v string) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldLT(FieldChatType, v))
}

// ChatTypeLTE applies the LTE predicate on the "chat_type" field.
func ChatTypeLTE(v string) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldLTE(FieldChatType, v))
}

// ChatTypeContains applies the Contains predicate on the "chat_type" field.
func ChatTypeContains(v string) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldContains(FieldChatType, v))
}

// ChatTypeHasPrefix applies the HasPrefix predicate on the "chat_type" field.
func ChatTypeHasPrefix(v string) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldHasPrefix(FieldChatType, v))
}

// ChatTypeHasSuffix applies the HasSuffix predicate on the "chat_type" field.
func ChatTypeHasSuffix(v string) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldHasSuffix(FieldChatType, v))
}

// ChatTypeEqualFold applies the EqualFold predicate on the "chat_type" field.
func ChatTypeEqualFold(v string) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldEqualFold(FieldChatType, v))
}

// ChatTypeContainsFold applies the ContainsFold predicate on the "chat_type" field.
func ChatTypeContainsFold(v string) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldContainsFold(FieldChatType, v))
}

// FeatureChatHistoriesRecapEQ applies the EQ predicate on the "feature_chat_histories_recap" field.
func FeatureChatHistoriesRecapEQ(v bool) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldEQ(FieldFeatureChatHistoriesRecap, v))
}

// FeatureChatHistoriesRecapNEQ applies the NEQ predicate on the "feature_chat_histories_recap" field.
func FeatureChatHistoriesRecapNEQ(v bool) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldNEQ(FieldFeatureChatHistoriesRecap, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TelegramChatFeatureFlags) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TelegramChatFeatureFlags) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(func(s *sql.Selector) {
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
func Not(p predicate.TelegramChatFeatureFlags) predicate.TelegramChatFeatureFlags {
	return predicate.TelegramChatFeatureFlags(func(s *sql.Selector) {
		p(s.Not())
	})
}
