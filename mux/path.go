package mux

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	pathSeperator    = "/"
	pathVarPrefix    = "{"
	pathVarSuffix    = "}"
	variableSep      = ":"
	patternRegex     = "regex"
	patternValidator = "validator"
)

func comparePath(reqPath string, mockPath Path) error {
	if reqPath == mockPath.Pattern {
		return nil
	}

	reqParts := strings.Split(reqPath, pathSeperator)
	mockParts := strings.Split(mockPath.Pattern, pathSeperator)

	if len(reqParts) != len(mockParts) {
		return fmt.Errorf("request path parts do not match actual[%d] expected[%d]", len(reqParts), len(mockParts))
	}
	for i, mp := range mockParts {
		switch {
		case reqParts[i] == mp:
			continue
		case !strings.HasPrefix(mp, pathVarPrefix), !strings.HasSuffix(mp, pathVarSuffix):
			return fmt.Errorf("request path part do not match actual[%s] expected[%s]", reqParts[i], mp)
		default:
		}
		log.Println(mp, reqParts[i])
		label := mp[1 : len(mockParts)-1]
		for _, v := range mockPath.Variables {
			if v.Label == label {
				pattern, compareFunc, err := retrivePathVariable(v.Value)
				if err != nil {
					log.Println(err.Error())
					return err
				}
				if err := compareFunc(reqParts[i], pattern); err != nil {
					log.Println(err.Error())
					return err
				}
				return nil
			}
		}
		return fmt.Errorf("request path unable to find variable label[%s]", label)
	}
	return nil
}

func retrivePathVariable(part string) (string, func(string, string) error, error) {
	p := strings.Split(part, variableSep)
	if len(p) == 0 {
		return "", nil, errors.New("part has not func")
	}
	pattern := part[len(p[0])+1:]
	switch p[0] {
	case patternRegex:
		return pattern, compareRegex, nil
	case patternValidator:
		return pattern, compareValidator, nil
	default:
		return "", nil, fmt.Errorf("request path func[%s] not supported", p[0])
	}
}

func compareRegex(reqPart string, pattern string) error {
	match, err := regexp.MatchString(pattern, reqPart)
	switch {
	case err != nil:
		return fmt.Errorf("request part regex %w", err)
	case !match:
		return fmt.Errorf("request part regex no match part[%s] pattern[%s]", reqPart, pattern)
	default:
		return nil
	}
}

func compareValidator(reqPart string, pattern string) error {
	if err := validator.New().Var(reqPart, pattern); err != nil {
		return fmt.Errorf("request part validator %w", err)
	}
	return nil
}
