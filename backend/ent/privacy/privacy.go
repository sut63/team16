// Code generated by entc, DO NOT EDIT.

package privacy

import (
	"context"
	"errors"
	"fmt"

	"github.com/G16/app/ent"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with an allow decision.
	Allow = errors.New("ent/privacy: allow rule")

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with an deny decision.
	Deny = errors.New("ent/privacy: deny rule")

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = errors.New("ent/privacy: skip rule")
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

type decisionCtxKey struct{}

// DecisionContext creates a decision context.
func DecisionContext(parent context.Context, decision error) context.Context {
	if decision == nil || errors.Is(decision, Skip) {
		return parent
	}
	return context.WithValue(parent, decisionCtxKey{}, decision)
}

func decisionFromContext(ctx context.Context) (error, bool) {
	decision, ok := ctx.Value(decisionCtxKey{}).(error)
	if ok && errors.Is(decision, Allow) {
		decision = nil
	}
	return decision, ok
}

type (
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy []QueryRule

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule interface {
		EvalQuery(context.Context, ent.Query) error
	}
)

// EvalQuery evaluates a query against a query policy.
func (policy QueryPolicy) EvalQuery(ctx context.Context, q ent.Query) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalQuery(ctx, q); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

type (
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy []MutationRule

	// MutationRule defines the interface deciding whether a
	// mutation is allowed and optionally modify it.
	MutationRule interface {
		EvalMutation(context.Context, ent.Mutation) error
	}
)

// EvalMutation evaluates a mutation against a mutation policy.
func (policy MutationPolicy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalMutation(ctx, m); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// MutationRuleFunc type is an adapter to allow the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// Policy groups query and mutation policies.
type Policy struct {
	Query    QueryPolicy
	Mutation MutationPolicy
}

// EvalQuery forwards evaluation to query policy.
func (policy Policy) EvalQuery(ctx context.Context, q ent.Query) error {
	return policy.Query.EvalQuery(ctx, q)
}

// EvalMutation forwards evaluation to mutation policy.
func (policy Policy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return policy.Mutation.EvalMutation(ctx, m)
}

// QueryMutationRule is the interface that groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AgeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AgeQueryRuleFunc func(context.Context, *ent.AgeQuery) error

// EvalQuery return f(ctx, q).
func (f AgeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AgeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AgeQuery", q)
}

// The AgeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AgeMutationRuleFunc func(context.Context, *ent.AgeMutation) error

// EvalMutation calls f(ctx, m).
func (f AgeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AgeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AgeMutation", m)
}

// The BookcourseQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type BookcourseQueryRuleFunc func(context.Context, *ent.BookcourseQuery) error

// EvalQuery return f(ctx, q).
func (f BookcourseQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BookcourseQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.BookcourseQuery", q)
}

// The BookcourseMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type BookcourseMutationRuleFunc func(context.Context, *ent.BookcourseMutation) error

// EvalMutation calls f(ctx, m).
func (f BookcourseMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.BookcourseMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.BookcourseMutation", m)
}

// The ClassifierQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ClassifierQueryRuleFunc func(context.Context, *ent.ClassifierQuery) error

// EvalQuery return f(ctx, q).
func (f ClassifierQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ClassifierQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ClassifierQuery", q)
}

// The ClassifierMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ClassifierMutationRuleFunc func(context.Context, *ent.ClassifierMutation) error

// EvalMutation calls f(ctx, m).
func (f ClassifierMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ClassifierMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ClassifierMutation", m)
}

// The CourseQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CourseQueryRuleFunc func(context.Context, *ent.CourseQuery) error

// EvalQuery return f(ctx, q).
func (f CourseQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CourseQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CourseQuery", q)
}

// The CourseMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CourseMutationRuleFunc func(context.Context, *ent.CourseMutation) error

// EvalMutation calls f(ctx, m).
func (f CourseMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CourseMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CourseMutation", m)
}

// The EmployeeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EmployeeQueryRuleFunc func(context.Context, *ent.EmployeeQuery) error

// EvalQuery return f(ctx, q).
func (f EmployeeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.EmployeeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.EmployeeQuery", q)
}

// The EmployeeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EmployeeMutationRuleFunc func(context.Context, *ent.EmployeeMutation) error

// EvalMutation calls f(ctx, m).
func (f EmployeeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.EmployeeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.EmployeeMutation", m)
}

// The EquipmentQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EquipmentQueryRuleFunc func(context.Context, *ent.EquipmentQuery) error

// EvalQuery return f(ctx, q).
func (f EquipmentQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.EquipmentQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.EquipmentQuery", q)
}

// The EquipmentMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EquipmentMutationRuleFunc func(context.Context, *ent.EquipmentMutation) error

// EvalMutation calls f(ctx, m).
func (f EquipmentMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.EquipmentMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.EquipmentMutation", m)
}

// The EquipmentrentalQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EquipmentrentalQueryRuleFunc func(context.Context, *ent.EquipmentrentalQuery) error

// EvalQuery return f(ctx, q).
func (f EquipmentrentalQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.EquipmentrentalQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.EquipmentrentalQuery", q)
}

// The EquipmentrentalMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EquipmentrentalMutationRuleFunc func(context.Context, *ent.EquipmentrentalMutation) error

// EvalMutation calls f(ctx, m).
func (f EquipmentrentalMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.EquipmentrentalMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.EquipmentrentalMutation", m)
}

// The EquipmenttypeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EquipmenttypeQueryRuleFunc func(context.Context, *ent.EquipmenttypeQuery) error

// EvalQuery return f(ctx, q).
func (f EquipmenttypeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.EquipmenttypeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.EquipmenttypeQuery", q)
}

// The EquipmenttypeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EquipmenttypeMutationRuleFunc func(context.Context, *ent.EquipmenttypeMutation) error

// EvalMutation calls f(ctx, m).
func (f EquipmenttypeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.EquipmenttypeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.EquipmenttypeMutation", m)
}

// The MemberQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type MemberQueryRuleFunc func(context.Context, *ent.MemberQuery) error

// EvalQuery return f(ctx, q).
func (f MemberQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MemberQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.MemberQuery", q)
}

// The MemberMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type MemberMutationRuleFunc func(context.Context, *ent.MemberMutation) error

// EvalMutation calls f(ctx, m).
func (f MemberMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.MemberMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.MemberMutation", m)
}

// The PaymentQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PaymentQueryRuleFunc func(context.Context, *ent.PaymentQuery) error

// EvalQuery return f(ctx, q).
func (f PaymentQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PaymentQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PaymentQuery", q)
}

// The PaymentMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PaymentMutationRuleFunc func(context.Context, *ent.PaymentMutation) error

// EvalMutation calls f(ctx, m).
func (f PaymentMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PaymentMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PaymentMutation", m)
}

// The PaymenttypeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PaymenttypeQueryRuleFunc func(context.Context, *ent.PaymenttypeQuery) error

// EvalQuery return f(ctx, q).
func (f PaymenttypeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PaymenttypeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PaymenttypeQuery", q)
}

// The PaymenttypeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PaymenttypeMutationRuleFunc func(context.Context, *ent.PaymenttypeMutation) error

// EvalMutation calls f(ctx, m).
func (f PaymenttypeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PaymenttypeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PaymenttypeMutation", m)
}

// The PositionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PositionQueryRuleFunc func(context.Context, *ent.PositionQuery) error

// EvalQuery return f(ctx, q).
func (f PositionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PositionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PositionQuery", q)
}

// The PositionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PositionMutationRuleFunc func(context.Context, *ent.PositionMutation) error

// EvalMutation calls f(ctx, m).
func (f PositionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PositionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PositionMutation", m)
}

// The PromotionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PromotionQueryRuleFunc func(context.Context, *ent.PromotionQuery) error

// EvalQuery return f(ctx, q).
func (f PromotionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PromotionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PromotionQuery", q)
}

// The PromotionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PromotionMutationRuleFunc func(context.Context, *ent.PromotionMutation) error

// EvalMutation calls f(ctx, m).
func (f PromotionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PromotionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PromotionMutation", m)
}

// The PromotionamountQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PromotionamountQueryRuleFunc func(context.Context, *ent.PromotionamountQuery) error

// EvalQuery return f(ctx, q).
func (f PromotionamountQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PromotionamountQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PromotionamountQuery", q)
}

// The PromotionamountMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PromotionamountMutationRuleFunc func(context.Context, *ent.PromotionamountMutation) error

// EvalMutation calls f(ctx, m).
func (f PromotionamountMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PromotionamountMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PromotionamountMutation", m)
}

// The PromotiontimeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PromotiontimeQueryRuleFunc func(context.Context, *ent.PromotiontimeQuery) error

// EvalQuery return f(ctx, q).
func (f PromotiontimeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PromotiontimeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PromotiontimeQuery", q)
}

// The PromotiontimeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PromotiontimeMutationRuleFunc func(context.Context, *ent.PromotiontimeMutation) error

// EvalMutation calls f(ctx, m).
func (f PromotiontimeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PromotiontimeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PromotiontimeMutation", m)
}

// The PromotiontypeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PromotiontypeQueryRuleFunc func(context.Context, *ent.PromotiontypeQuery) error

// EvalQuery return f(ctx, q).
func (f PromotiontypeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PromotiontypeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PromotiontypeQuery", q)
}

// The PromotiontypeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PromotiontypeMutationRuleFunc func(context.Context, *ent.PromotiontypeMutation) error

// EvalMutation calls f(ctx, m).
func (f PromotiontypeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PromotiontypeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PromotiontypeMutation", m)
}

// The SalaryQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SalaryQueryRuleFunc func(context.Context, *ent.SalaryQuery) error

// EvalQuery return f(ctx, q).
func (f SalaryQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SalaryQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.SalaryQuery", q)
}

// The SalaryMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SalaryMutationRuleFunc func(context.Context, *ent.SalaryMutation) error

// EvalMutation calls f(ctx, m).
func (f SalaryMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.SalaryMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.SalaryMutation", m)
}

// The ZoneQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ZoneQueryRuleFunc func(context.Context, *ent.ZoneQuery) error

// EvalQuery return f(ctx, q).
func (f ZoneQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ZoneQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ZoneQuery", q)
}

// The ZoneMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ZoneMutationRuleFunc func(context.Context, *ent.ZoneMutation) error

// EvalMutation calls f(ctx, m).
func (f ZoneMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ZoneMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ZoneMutation", m)
}
