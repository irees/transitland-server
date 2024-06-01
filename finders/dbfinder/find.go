package dbfinder

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	sq "github.com/Masterminds/squirrel"
	"github.com/interline-io/transitland-server/model"
)

// Maximum query result limit
var MAXLIMIT = 100_000

// helpers

func checkLimit(limit *int) uint64 {
	return checkRange(limit, 0, MAXLIMIT)
}

func checkRange(limit *int, min, max int) uint64 {
	if limit == nil {
		return uint64(max)
	} else if *limit >= max {
		return uint64(max)
	} else if *limit < min {
		return uint64(min)
	}
	return uint64(*limit)
}

func checkFloat(v *float64, min float64, max float64) float64 {
	if v == nil || *v < min {
		return min
	} else if *v > max {
		return max
	}
	return *v
}

func atoi(v string) int {
	a, _ := strconv.Atoi(v)
	return a
}

// unicode aware remove all non-alphanumeric characters
// this is not for escaping sql; just for preparing to_tsquery
func alphanumeric(v string) string {
	ret := []rune{}
	for _, ch := range v {
		if unicode.IsSpace(ch) {
			ret = append(ret, ' ')
		} else if unicode.IsDigit(ch) || unicode.IsLetter(ch) {
			ret = append(ret, ch)
		}
	}
	return string(ret)
}

// az09 removes any character that is not a a-z or 0-9 or _ or .
func az09(v string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9_\.]+`)
	return reg.ReplaceAllString(v, "")
}

func escapeWordsWithSuffix(v string, sfx string) []string {
	var ret []string
	for _, s := range strings.Fields(v) {
		aa := alphanumeric(s)
		// Minimum length 2 characters
		if len(aa) > 1 {
			ret = append(ret, aa+sfx)
		}
	}
	return ret
}

func tsTableQuery(table string, s string) (rank sq.Sqlizer, wc sq.Sqlizer) {
	s = strings.TrimSpace(s)
	words := append([]string{}, escapeWordsWithSuffix(s, ":*")...)
	wordstsq := strings.Join(words, " & ")
	rank = sq.Expr(
		fmt.Sprintf(`ts_rank_cd("%s".textsearch,to_tsquery('tl',?)) as search_rank`, az09(table)),
		wordstsq,
	)
	wc = sq.Expr(
		fmt.Sprintf(`"%s".textsearch @@ to_tsquery('tl',?)`, az09(table)),
		wordstsq,
	)

	return rank, wc
}

func lateralWrap(q sq.SelectBuilder, outerTable string, outerKey string, innerTable string, innerKey string, outerIds []int) sq.SelectBuilder {
	outerTable = az09(outerTable)
	outerKey = az09(outerKey)
	innerTable = az09(innerTable)
	innerKey = az09(innerKey)
	qInner := q.Where(fmt.Sprintf("%s.%s = out.%s", innerTable, innerKey, outerKey))
	q2 := sq.StatementBuilder.
		Select("t.*").
		From(outerTable + " out").
		JoinClause(qInner.Prefix("JOIN LATERAL (").Suffix(") t on true")).
		Where(sq.Eq{"out." + outerKey: outerIds})
	return q2
}

func quickSelect(table string, limit *int, after *model.Cursor, ids []int) sq.SelectBuilder {
	return quickSelectOrder(table, limit, after, ids, "id")
}

func quickSelectOrder(table string, limit *int, after *model.Cursor, ids []int, order string) sq.SelectBuilder {
	table = az09(table)
	order = az09(order)
	q := sq.StatementBuilder.
		Select("*").
		From(table).
		Limit(checkLimit(limit))
	if order != "" {
		q = q.OrderBy(order)
	}
	if len(ids) > 0 {
		q = q.Where(sq.Eq{"id": ids})
	}
	if after != nil && after.Valid && after.ID > 0 {
		q = q.Where(sq.Gt{"id": after.ID})
	}
	return q
}
