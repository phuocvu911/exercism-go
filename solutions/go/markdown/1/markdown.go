package markdown

// implementation to refactor

import (
	"fmt"
	"strings"
)

type renderState struct {
	pos        int
	header     int
	he         bool
	list       int
	listOpened bool
}

// Render translates markdown to HTML
func Render(markdown string) string {
	markdown = applyInlineTags(markdown)
	if len(markdown) == 0 {
		return "<p></p>"
	}

	var st renderState
	var b strings.Builder

	for st.pos < len(markdown) {
		char := markdown[st.pos]
		if char == '#' {
			handleHeader(markdown, &st, &b)
			continue
		}

		st.he = true

		if isListStart(markdown, char, st.header) {
			handleListStart(markdown, &st, &b)
			continue
		}

		if char == '\n' {
			handleNewline(markdown, &st, &b)
			continue
		}

		b.WriteByte(char)
		st.pos++
	}

	html := b.String()
	return finalizeHTML(markdown, html, st.header, st.list)
}

func applyInlineTags(markdown string) string {
	// Preserve the original behavior: only replace the first occurrence of each marker.
	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)
	return markdown
}

func handleHeader(markdown string, st *renderState, b *strings.Builder) {
	char := markdown[st.pos]
	for char == '#' {
		st.header++
		st.pos++
		if st.pos >= len(markdown) {
			break
		}
		char = markdown[st.pos]
	}

	switch {
	case st.header == 7:
		b.WriteString(fmt.Sprintf("<p>%s ", strings.Repeat("#", st.header)))
	case st.he:
		b.WriteString("# ")
		st.header--
	default:
		b.WriteString(fmt.Sprintf("<h%d>", st.header))
	}

	// Match original: consume one extra char after the header marker run (typically a space).
	st.pos++
}

func isListStart(markdown string, char byte, header int) bool {
	return char == '*' && header == 0 && strings.Contains(markdown, "\n")
}

func handleListStart(markdown string, st *renderState, b *strings.Builder) {
	if st.list == 0 {
		b.WriteString("<ul>")
	}
	st.list++

	if !st.listOpened {
		b.WriteString("<li>")
		st.listOpened = true
	} else {
		b.WriteString(string(markdown[st.pos]))
		b.WriteString(" ")
	}

	// Match original: skip "* " (2 chars).
	st.pos += 2
}

func handleNewline(markdown string, st *renderState, b *strings.Builder) {
	if st.listOpened && strings.LastIndex(markdown, "\n") == st.pos && strings.LastIndex(markdown, "\n") > strings.LastIndex(markdown, "*") {
		b.WriteString("</li></ul><p>")
		st.listOpened = false
		st.list = 0
	}
	if st.list > 0 && st.listOpened {
		b.WriteString("</li>")
		st.listOpened = false
	}
	if st.header > 0 {
		b.WriteString(fmt.Sprintf("</h%d>", st.header))
		st.header = 0
	}
	st.pos++
}

func finalizeHTML(markdown, html string, header, list int) string {
	switch {
	case header == 7:
		return html + "</p>"
	case header > 0:
		return html + fmt.Sprintf("</h%d>", header)
	}
	if list > 0 {
		return html + "</li></ul>"
	}
	if strings.Contains(html, "<p>") {
		return html + "</p>"
	}
	return "<p>" + html + "</p>"
}
