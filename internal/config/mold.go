package config

import (
	"context"
	"errors"
	"github.com/go-playground/mold/v4"
	"github.com/go-playground/mold/v4/modifiers"
	"reflect"
	"regexp"
	"strings"
)

func NewMold() *mold.Transformer {
	m := modifiers.New()
	m.Register("normalize_spaces", normalizeSpaces)
	return m
}

func normalizeSpaces(ctx context.Context, fl mold.FieldLevel) error {
	field := fl.Field()
	if field.Kind() != reflect.String {
		return errors.New("field is not a string")
	}
	str := field.String()
	trimmed := strings.TrimSpace(str)
	re := regexp.MustCompile(`\s+`)
	normalized := re.ReplaceAllString(trimmed, " ")
	field.SetString(normalized)
	return nil
}
