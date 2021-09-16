package main

type Language struct {
	name string
	extension string
}

type Languages []Language

var languages = Languages{
	Language{name: "golang", extension: "go"},
	Language{name: "php", extension: "php"},
	Language{name: "javascript", extension: "js"},
	Language{name: "typescript", extension: "ts"},
	Language{name: "ruby", extension: "rb"},
	Language{name: "css", extension: "css"},
	Language{name: "scss", extension: "scss"},
	Language{name: "html", extension: "html"},
	Language{name: "xml", extension: "xml"},
	Language{name: "python", extension: "py"},
	Language{name: "java", extension: "java"},
	Language{name: "c#", extension: "cs"},
	Language{name: "c", extension: "c"},
	Language{name: "c++", extension: "cpp"},
	Language{name: "rust", extension: "rs"},
	Language{name: "scala", extension: "scala"},
}
