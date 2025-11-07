package html

// A creates a a element with optional children and attributes.
func A(items ...any) Node {
	return Element("a", false, items...)
}

// Abbr creates a abbr element with optional children and attributes.
func Abbr(items ...any) Node {
	return Element("abbr", false, items...)
}

// Address creates a address element with optional children and attributes.
func Address(items ...any) Node {
	return Element("address", false, items...)
}

// Area creates a area element with optional children and attributes.
func Area(items ...any) Node {
	return Element("area", true, items...)
}

// Article creates a article element with optional children and attributes.
func Article(items ...any) Node {
	return Element("article", false, items...)
}

// Aside creates a aside element with optional children and attributes.
func Aside(items ...any) Node {
	return Element("aside", false, items...)
}

// Audio creates a audio element with optional children and attributes.
func Audio(items ...any) Node {
	return Element("audio", false, items...)
}

// B creates a b element with optional children and attributes.
func B(items ...any) Node {
	return Element("b", false, items...)
}

// Base creates a base element with optional children and attributes.
func Base(items ...any) Node {
	return Element("base", true, items...)
}

// Bdi creates a bdi element with optional children and attributes.
func Bdi(items ...any) Node {
	return Element("bdi", false, items...)
}

// Bdo creates a bdo element with optional children and attributes.
func Bdo(items ...any) Node {
	return Element("bdo", false, items...)
}

// Blockquote creates a blockquote element with optional children and attributes.
func Blockquote(items ...any) Node {
	return Element("blockquote", false, items...)
}

// Body creates a body element with optional children and attributes.
func Body(items ...any) Node {
	return Element("body", false, items...)
}

// Br creates a br element with optional children and attributes.
func Br(items ...any) Node {
	return Element("br", true, items...)
}

// Button creates a button element with optional children and attributes.
func Button(items ...any) Node {
	return Element("button", false, items...)
}

// Canvas creates a canvas element with optional children and attributes.
func Canvas(items ...any) Node {
	return Element("canvas", false, items...)
}

// Caption creates a caption element with optional children and attributes.
func Caption(items ...any) Node {
	return Element("caption", false, items...)
}

// Cite creates a cite element with optional children and attributes.
func Cite(items ...any) Node {
	return Element("cite", false, items...)
}

// Code creates a code element with optional children and attributes.
func Code(items ...any) Node {
	return Element("code", false, items...)
}

// Col creates a col element with optional children and attributes.
func Col(items ...any) Node {
	return Element("col", true, items...)
}

// Colgroup creates a colgroup element with optional children and attributes.
func Colgroup(items ...any) Node {
	return Element("colgroup", false, items...)
}

// Data creates a data element with optional children and attributes.
func Data(items ...any) Node {
	return Element("data", false, items...)
}

// Datalist creates a datalist element with optional children and attributes.
func Datalist(items ...any) Node {
	return Element("datalist", false, items...)
}

// Dd creates a dd element with optional children and attributes.
func Dd(items ...any) Node {
	return Element("dd", false, items...)
}

// Del creates a del element with optional children and attributes.
func Del(items ...any) Node {
	return Element("del", false, items...)
}

// Details creates a details element with optional children and attributes.
func Details(items ...any) Node {
	return Element("details", false, items...)
}

// Dfn creates a dfn element with optional children and attributes.
func Dfn(items ...any) Node {
	return Element("dfn", false, items...)
}

// Dialog creates a dialog element with optional children and attributes.
func Dialog(items ...any) Node {
	return Element("dialog", false, items...)
}

// Div creates a div element with optional children and attributes.
func Div(items ...any) Node {
	return Element("div", false, items...)
}

// Dl creates a dl element with optional children and attributes.
func Dl(items ...any) Node {
	return Element("dl", false, items...)
}

// Dt creates a dt element with optional children and attributes.
func Dt(items ...any) Node {
	return Element("dt", false, items...)
}

// Em creates a em element with optional children and attributes.
func Em(items ...any) Node {
	return Element("em", false, items...)
}

// Embed creates a embed element with optional children and attributes.
func Embed(items ...any) Node {
	return Element("embed", true, items...)
}

// Fieldset creates a fieldset element with optional children and attributes.
func Fieldset(items ...any) Node {
	return Element("fieldset", false, items...)
}

// Figcaption creates a figcaption element with optional children and attributes.
func Figcaption(items ...any) Node {
	return Element("figcaption", false, items...)
}

// Figure creates a figure element with optional children and attributes.
func Figure(items ...any) Node {
	return Element("figure", false, items...)
}

// Footer creates a footer element with optional children and attributes.
func Footer(items ...any) Node {
	return Element("footer", false, items...)
}

// Form creates a form element with optional children and attributes.
func Form(items ...any) Node {
	return Element("form", false, items...)
}

// H1 creates a h1 element with optional children and attributes.
func H1(items ...any) Node {
	return Element("h1", false, items...)
}

// H2 creates a h2 element with optional children and attributes.
func H2(items ...any) Node {
	return Element("h2", false, items...)
}

// H3 creates a h3 element with optional children and attributes.
func H3(items ...any) Node {
	return Element("h3", false, items...)
}

// H4 creates a h4 element with optional children and attributes.
func H4(items ...any) Node {
	return Element("h4", false, items...)
}

// H5 creates a h5 element with optional children and attributes.
func H5(items ...any) Node {
	return Element("h5", false, items...)
}

// H6 creates a h6 element with optional children and attributes.
func H6(items ...any) Node {
	return Element("h6", false, items...)
}

// Head creates a head element with optional children and attributes.
func Head(items ...any) Node {
	return Element("head", false, items...)
}

// Header creates a header element with optional children and attributes.
func Header(items ...any) Node {
	return Element("header", false, items...)
}

// Hgroup creates a hgroup element with optional children and attributes.
func Hgroup(items ...any) Node {
	return Element("hgroup", false, items...)
}

// Hr creates a hr element with optional children and attributes.
func Hr(items ...any) Node {
	return Element("hr", true, items...)
}

// Html creates a html element with optional children and attributes.
func Html(items ...any) Node {
	return Element("html", false, items...)
}

// I creates a i element with optional children and attributes.
func I(items ...any) Node {
	return Element("i", false, items...)
}

// Iframe creates a iframe element with optional children and attributes.
func Iframe(items ...any) Node {
	return Element("iframe", false, items...)
}

// Img creates a img element with optional children and attributes.
func Img(items ...any) Node {
	return Element("img", true, items...)
}

// Input creates a input element with optional children and attributes.
func Input(items ...any) Node {
	return Element("input", true, items...)
}

// Ins creates a ins element with optional children and attributes.
func Ins(items ...any) Node {
	return Element("ins", false, items...)
}

// Kbd creates a kbd element with optional children and attributes.
func Kbd(items ...any) Node {
	return Element("kbd", false, items...)
}

// Label creates a label element with optional children and attributes.
func Label(items ...any) Node {
	return Element("label", false, items...)
}

// Legend creates a legend element with optional children and attributes.
func Legend(items ...any) Node {
	return Element("legend", false, items...)
}

// Li creates a li element with optional children and attributes.
func Li(items ...any) Node {
	return Element("li", false, items...)
}

// Link creates a link element with optional children and attributes.
func Link(items ...any) Node {
	return Element("link", true, items...)
}

// Main creates a main element with optional children and attributes.
func Main(items ...any) Node {
	return Element("main", false, items...)
}

// MapEl creates a map element with optional children and attributes.
func MapEl(items ...any) Node {
	return Element("map", false, items...)
}

// Mark creates a mark element with optional children and attributes.
func Mark(items ...any) Node {
	return Element("mark", false, items...)
}

// Menu creates a menu element with optional children and attributes.
func Menu(items ...any) Node {
	return Element("menu", false, items...)
}

// Meta creates a meta element with optional children and attributes.
func Meta(items ...any) Node {
	return Element("meta", true, items...)
}

// Meter creates a meter element with optional children and attributes.
func Meter(items ...any) Node {
	return Element("meter", false, items...)
}

// Nav creates a nav element with optional children and attributes.
func Nav(items ...any) Node {
	return Element("nav", false, items...)
}

// Noscript creates a noscript element with optional children and attributes.
func Noscript(items ...any) Node {
	return Element("noscript", false, items...)
}

// Object creates a object element with optional children and attributes.
func Object(items ...any) Node {
	return Element("object", false, items...)
}

// Ol creates a ol element with optional children and attributes.
func Ol(items ...any) Node {
	return Element("ol", false, items...)
}

// Optgroup creates a optgroup element with optional children and attributes.
func Optgroup(items ...any) Node {
	return Element("optgroup", false, items...)
}

// Option creates a option element with optional children and attributes.
func Option(items ...any) Node {
	return Element("option", false, items...)
}

// Output creates a output element with optional children and attributes.
func Output(items ...any) Node {
	return Element("output", false, items...)
}

// P creates a p element with optional children and attributes.
func P(items ...any) Node {
	return Element("p", false, items...)
}

// Picture creates a picture element with optional children and attributes.
func Picture(items ...any) Node {
	return Element("picture", false, items...)
}

// Pre creates a pre element with optional children and attributes.
func Pre(items ...any) Node {
	return Element("pre", false, items...)
}

// Progress creates a progress element with optional children and attributes.
func Progress(items ...any) Node {
	return Element("progress", false, items...)
}

// Q creates a q element with optional children and attributes.
func Q(items ...any) Node {
	return Element("q", false, items...)
}

// Rp creates a rp element with optional children and attributes.
func Rp(items ...any) Node {
	return Element("rp", false, items...)
}

// Rt creates a rt element with optional children and attributes.
func Rt(items ...any) Node {
	return Element("rt", false, items...)
}

// Ruby creates a ruby element with optional children and attributes.
func Ruby(items ...any) Node {
	return Element("ruby", false, items...)
}

// S creates a s element with optional children and attributes.
func S(items ...any) Node {
	return Element("s", false, items...)
}

// Samp creates a samp element with optional children and attributes.
func Samp(items ...any) Node {
	return Element("samp", false, items...)
}

// Script creates a script element with optional children and attributes.
func Script(items ...any) Node {
	return Element("script", false, items...)
}

// Search creates a search element with optional children and attributes.
func Search(items ...any) Node {
	return Element("search", false, items...)
}

// Section creates a section element with optional children and attributes.
func Section(items ...any) Node {
	return Element("section", false, items...)
}

// Select creates a select element with optional children and attributes.
func Select(items ...any) Node {
	return Element("select", false, items...)
}

// Slot creates a slot element with optional children and attributes.
func Slot(items ...any) Node {
	return Element("slot", false, items...)
}

// Small creates a small element with optional children and attributes.
func Small(items ...any) Node {
	return Element("small", false, items...)
}

// Source creates a source element with optional children and attributes.
func Source(items ...any) Node {
	return Element("source", true, items...)
}

// Span creates a span element with optional children and attributes.
func Span(items ...any) Node {
	return Element("span", false, items...)
}

// Strong creates a strong element with optional children and attributes.
func Strong(items ...any) Node {
	return Element("strong", false, items...)
}

// Style creates a style element with optional children and attributes.
func Style(items ...any) Node {
	return Element("style", false, items...)
}

// Sub creates a sub element with optional children and attributes.
func Sub(items ...any) Node {
	return Element("sub", false, items...)
}

// Summary creates a summary element with optional children and attributes.
func Summary(items ...any) Node {
	return Element("summary", false, items...)
}

// Sup creates a sup element with optional children and attributes.
func Sup(items ...any) Node {
	return Element("sup", false, items...)
}

// Table creates a table element with optional children and attributes.
func Table(items ...any) Node {
	return Element("table", false, items...)
}

// Tbody creates a tbody element with optional children and attributes.
func Tbody(items ...any) Node {
	return Element("tbody", false, items...)
}

// Td creates a td element with optional children and attributes.
func Td(items ...any) Node {
	return Element("td", false, items...)
}

// Template creates a template element with optional children and attributes.
func Template(items ...any) Node {
	return Element("template", false, items...)
}

// Textarea creates a textarea element with optional children and attributes.
func Textarea(items ...any) Node {
	return Element("textarea", false, items...)
}

// Tfoot creates a tfoot element with optional children and attributes.
func Tfoot(items ...any) Node {
	return Element("tfoot", false, items...)
}

// Th creates a th element with optional children and attributes.
func Th(items ...any) Node {
	return Element("th", false, items...)
}

// Thead creates a thead element with optional children and attributes.
func Thead(items ...any) Node {
	return Element("thead", false, items...)
}

// Time creates a time element with optional children and attributes.
func Time(items ...any) Node {
	return Element("time", false, items...)
}

// Title creates a title element with optional children and attributes.
func Title(items ...any) Node {
	return Element("title", false, items...)
}

// Tr creates a tr element with optional children and attributes.
func Tr(items ...any) Node {
	return Element("tr", false, items...)
}

// Track creates a track element with optional children and attributes.
func Track(items ...any) Node {
	return Element("track", true, items...)
}

// U creates a u element with optional children and attributes.
func U(items ...any) Node {
	return Element("u", false, items...)
}

// Ul creates a ul element with optional children and attributes.
func Ul(items ...any) Node {
	return Element("ul", false, items...)
}

// Var creates a var element with optional children and attributes.
func Var(items ...any) Node {
	return Element("var", false, items...)
}

// Video creates a video element with optional children and attributes.
func Video(items ...any) Node {
	return Element("video", false, items...)
}

// Wbr creates a wbr element with optional children and attributes.
func Wbr(items ...any) Node {
	return Element("wbr", true, items...)
}

