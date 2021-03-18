package data

import (
	"fmt"

	"github.com/go-playground/validator"
)

// ErrorInvalidRelationshipType : Invalid RelationshipType specific error
var ErrorInvalidRelationshipType = fmt.Errorf("Invalid RelationshipType")

// ValidateRelationship a relationship with json validation and customer SKU validator
func (relationship *Relationship) ValidateRelationship() error {
	validate := validator.New()

	err := validate.RegisterValidation("isRelationshipType", validateIsRelationshipType)
	if err != nil {
		panic(ErrorInvalidRelationshipType)
	}
	
	return validate.Struct(relationship)
}

// validates the relationship type is valid
func validateIsRelationshipType(fieldLevel validator.FieldLevel) bool {
	relationshipType := int(fieldLevel.Field().Int())

	return !(relationshipType < int(None) || relationshipType > int(PendingOutgoing))
}
