package p68

import "strings"

/*
 * @lc app=leetcode id=68 lang=golang
 *
 * [68] Text Justification
 */

// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	if len(words) == 0 {
		return []string{}
	}
	i := 0
	for i < len(words) {
		if lenwords(words[0:i+1])+i > maxWidth {
			break
		}
		i++
	}
	words, left_words := words[0:i], words[i:]
	if len(left_words) == 0 {
		return []string{leftJustify(words, maxWidth)}
	}
	return append([]string{justify(words, maxWidth)}, fullJustify(left_words, maxWidth)...)
}

func justify(words []string, maxwidth int) string {
	if len(words) == 1 {
		return leftJustify(words, maxwidth)
	}
	spaces := make([]string, len(words)-1)
	l := 0
	for _, v := range words {
		l += len(v)
	}
	av := (maxwidth - l) / (len(words) - 1)

	for i := range spaces {
		spaces[i] = gen(' ', av)
	}

	for i := 0; i < maxwidth-l-av*(len(words)-1); i++ {
		spaces[i] = spaces[i] + " "
	}
	var sb strings.Builder
	for i := 0; i < len(spaces); i++ {
		sb.WriteString(words[i])
		sb.WriteString(spaces[i])
	}
	sb.WriteString(words[len(words)-1])
	return sb.String()
}

func leftJustify(words []string, maxwidth int) string {
	var sb strings.Builder
	sb.WriteString(words[0])
	for _, v := range words[1:] {
		sb.WriteRune(' ')
		sb.WriteString(v)
	}
	j :=  maxwidth-sb.Len()
	for i := 0; i < j; i++ {
		sb.WriteRune(' ')
	}
	return sb.String()
}

func gen(c rune, times int) string {
	var sb strings.Builder
	for i := 0; i < times; i++ {
		sb.WriteRune(c)
	}
	return sb.String()
}

func lenwords(words []string) int {
	l := 0
	for _, v := range words {
		l += len(v)
	}
	return l
}

// @lc code=end
