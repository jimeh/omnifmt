//
// Copyright (c) 2015 Lorenzo Villani
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.
//

package main

var formatters = []*formatter{
	// C/C++
	{
		Commands: [][]string{
			[]string{"clang-format", "-style=WebKit", "-"},
		},
		Extensions: []string{".c", ".cpp", ".cxx", ".h", ".hpp", ".hxx"},
		Install: installMap{
			"darwin": []string{"brew", "install", "clang-format"},
		},
		Syntaxes: syntaxMap{
			EMACS:   []string{"c-mode", "c++-mode"},
			SUBLIME: []string{"C", "C++"},
		},
	},
	// CSS
	{
		Commands: [][]string{
			[]string{"cssfmt"},
		},
		Extensions: []string{".css"},
		Install: installMap{
			"any": []string{"npm", "install", "-g", "cssfmt"},
		},
		Syntaxes: syntaxMap{
			EMACS:   []string{"css-mode"},
			SUBLIME: []string{"CSS"},
		},
	},
	// Go
	{
		Commands: [][]string{
			[]string{"goimports"},
		},
		Extensions: []string{".go"},
		Install: installMap{
			"any": []string{"go", "get", "golang.org/x/tools/cmd/goimports"},
		},
		Syntaxes: syntaxMap{
			EMACS:   []string{"go-mode"},
			SUBLIME: []string{"GoSublime-Go"},
		},
	},
	// JavaScript
	{
		Commands: [][]string{
			[]string{"semistandard-format", "-"},
		},
		Extensions: []string{".js", ".jsx"},
		Install: installMap{
			"any": []string{"npm", "install", "-g", "semistandard-format"},
		},
		Syntaxes: syntaxMap{
			EMACS:   []string{"js-mode", "js2-mode", "js3-mode"},
			SUBLIME: []string{"JavaScript", "JavaScript (Babel)"},
		},
	},
	// JSON
	{
		Commands: [][]string{
			[]string{"jsonlint", "-"},
		},
		Extensions: []string{".json"},
		Install: installMap{
			"any": []string{"npm", "install", "-g", "jsonlint"},
		},
		Syntaxes: syntaxMap{
			EMACS:   []string{"json-mode"},
			SUBLIME: []string{"JSON", "Sublime Commands"},
		},
	},
	// Python
	{
		Commands: [][]string{
			[]string{"autopep8", "--max-line-length=98", "-"},
			[]string{"isort", "--line-width", "98", "--multi_line", "3", "-"},
		},
		Install: installMap{
			"any": []string{"pip", "install", "autopep8", "isort"},
		},
		Extensions: []string{".py"},
		Syntaxes: syntaxMap{
			EMACS:   []string{"python-mode"},
			SUBLIME: []string{"Python"},
		},
	},
	// Ruby
	{
		Commands: [][]string{
			[]string{"rubocop", "--stdin", "--cache", "false", "--auto-correct", "--format", "emacs", "fake.rb"},
			[]string{"sed", "-n", `/^====================$/,$p`},
			[]string{"tail", "-n", "+2"},
		},
		Install: installMap{
			"any": []string{"gem", "install", "rubocop"},
		},
		Extensions: []string{".rb", ".rake"},
		Syntaxes: syntaxMap{
			EMACS:   []string{"ruby-mode"},
			SUBLIME: []string{"Ruby"},
		},
	},
	// SASS
	{
		Commands: [][]string{
			[]string{"sass-convert", "--no-cache", "--from", "sass", "--to", "sass", "--indent", "4", "--stdin"},
		},
		Install: installMap{
			"any": []string{"gem", "install", "sass"},
		},
		Extensions: []string{".sass"},
		Syntaxes: syntaxMap{
			EMACS:   []string{"sass-mode"},
			SUBLIME: []string{"SASS"},
		},
	},
	// SCSS
	{
		Commands: [][]string{
			[]string{"sass-convert", "--no-cache", "--from", "scss", "--to", "scss", "--indent", "4", "--stdin"},
		},
		Extensions: []string{".scss"},
		Install: installMap{
			"any": []string{"gem", "install", "sass"},
		},
		Syntaxes: syntaxMap{
			EMACS:   []string{"scss-mode"},
			SUBLIME: []string{"SCSS"},
		},
	},
	// Terraform
	{
		Commands: [][]string{
			[]string{"terraform", "fmt", "-"},
		},
		Extensions: []string{".tf"},
		Install: installMap{
			"darwin": []string{"brew", "install", "terraform"},
		},
		Syntaxes: syntaxMap{
			EMACS:   []string{"terraform-mode"},
			SUBLIME: []string{"Terraform"},
		},
	},
	// XML
	{
		Commands: [][]string{
			[]string{"xmllint", "--format", "-"},
		},
		Extensions: []string{".xml"},
		Install:    installMap{},
		Syntaxes: syntaxMap{
			EMACS:   []string{"xml-mode"},
			SUBLIME: []string{"XML"},
		},
	},
}
