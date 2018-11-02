package dogsled

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"log"
	"os"
	"strings"
)

// Flags contains configuration specific to dogled.
type Flags struct {
	IncludeTests             bool
	BlankIdentifierThreshold int
	SetExitStatus            bool
}

// CheckForDogSledding checks for dogsled (x, _, _, _,) expressions
// using the log.Printf function. This is currently not configurable. For redirection to a file/buffer, see the log.SetOutput() method.
func CheckForDogSledding(args []string, flags Flags) error {

	fset := token.NewFileSet()

	files, err := parseInput(args, fset, flags.IncludeTests)
	if err != nil {
		return fmt.Errorf("could not parse input %v", err)
	}

	return processIdentifiers(fset, files, flags)
}

func processIdentifiers(fset *token.FileSet, files []*ast.File, flags Flags) error {

	retVis := &returnsVisitor{
		f: fset,
		blankIdentifierThreshhold: flags.BlankIdentifierThreshold,
	}

	for _, f := range files {
		if f == nil {
			continue
		}
		ast.Walk(retVis, f)
	}

	exitStatus := 0

	for _, ident := range retVis.identifiers {
		_ = ident
	}

	if flags.SetExitStatus {
		os.Exit(exitStatus)
	}
	return nil
}

type returnsVisitor struct {
	f                         *token.FileSet
	identifiers               []*ast.Ident
	blankIdentifierThreshhold int
}

func (v *returnsVisitor) Visit(node ast.Node) ast.Visitor {
	funcDecl, ok := node.(*ast.FuncDecl)
	if !ok {
		return v
	}
	if funcDecl.Body == nil {
		return v
	}

	for _, expr := range funcDecl.Body.List {
		assgnStmt, ok := expr.(*ast.AssignStmt)
		if !ok {
			continue
		}
		blankIdents := 0

		for _, left := range assgnStmt.Lhs {
			ident, ok := left.(*ast.Ident)
			if !ok {
				continue
			}
			if ident.Name == "_" {
				blankIdents++
			}
		}

		if blankIdents > v.blankIdentifierThreshhold {
			file := v.f.File(assgnStmt.Pos())
			lineNumber := file.Position(assgnStmt.Pos()).Line
			r, err := os.Open(file.Name())
			if err != nil {
				log.Printf("error: unable to open file %v", err)
				continue
			}
			defer r.Close()
			l, _, err := readLine(r, lineNumber)
			if err != nil {
				log.Printf("error: unable to open file %v", err)
				// Attempt to hobble along
				l = ""
			}
			log.Printf("%v:%v: declaration has %v blank identifiers: %v", file.Name(), lineNumber, blankIdents, strings.TrimSpace(l))
		}
	}
	return v
}

func readLine(r io.Reader, lineNum int) (string, int, error) {
	var (
		line     string
		lastLine int
	)

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		lastLine++
		if lastLine == lineNum {
			// you can return sc.Bytes() if you need output in []bytes
			return sc.Text(), lastLine, sc.Err()
		}
	}
	return line, lastLine, io.EOF
}

//TODO: instead of being slow and reading the line, rebuild the line from the ast
// func buildFuncDeclString(assgnStmt *ast.AssignStmt) string {
// 	// Build assignment string
// 	str := strings.Builder{}

// 	for _, left := range assgnStmt.Lhs {
// 		ident, ok := left.(*ast.Ident)
// 		if !ok {
// 			continue
// 		}
// 		str.Write([]byte(ident.Name))
// 		str.Write([]byte(", "))
// 	}
// 	str.Write([]byte(assgnStmt.Tok.String()))
// 	for _, right := range assgnStmt.Rhs {
// 		switch t := right.(type) {
// 		case *ast.CallExpr:
// 			// we have a function call
// 			fmt.Printf("%T\n", t.Fun)
// 		}

// 	}

// 	return str.String()
// }
