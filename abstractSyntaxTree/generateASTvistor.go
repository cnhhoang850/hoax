package abstractSyntaxTree

import (
	"fmt"
	"os"
	"strings"
)

func GenerateAST(path string, baseName string, types []string) {
	file, err := os.Create(path + "/" + baseName + ".go")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	imports := `import (
	"fmt"
	"hoax/token"
)
`

	basenameContent := fmt.Sprintf("package parser\n\n%s\n type %s interface {\n\tAccept(visitor VisitorInterface)\n}\n\n", imports, baseName)

	for _, typeStr := range types {
		parts := strings.Split(typeStr, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid type string", typeStr)
			return
		}

		className := strings.TrimSpace(parts[0])
		fields := strings.TrimSpace(parts[1])

		basenameContent += defineSubClass(baseName, className, fields)
	}

	// Add the VisitorInterface and a default Visitor to the generated code
	basenameContent += `
type VisitorInterface interface {` + visitorInterface(types) + `
}

type Visitor struct{}

` + defaultVisitorImplementations(types)

	_, err = file.WriteString(basenameContent)
	if err != nil {
		fmt.Println("Error writing to file", err)
	}
}

func defineSubClass(structName string, className string, fields string) string {
	fieldSplits := strings.Split(fields, ",")
	structString := fmt.Sprintf(
		"type %s struct {\n\t%s %s\n\t%s\n}\n\n",
		className,
		structName,
		structName,
		strings.Join(fieldSplits, "\n\t"),
	)
	acceptMethod := fmt.Sprintf(
		`func (x *%s) Accept(visitor VisitorInterface) {
	visitor.Visit%s(x)
}
`, className, className)

	return structString + acceptMethod
}

func visitorInterface(types []string) string {
	methods := ""
	for _, typeStr := range types {
		parts := strings.Split(typeStr, ":")
		className := strings.TrimSpace(parts[0])
		methods += fmt.Sprintf("\n\tVisit%s(expr *%s)", className, className)
	}
	return methods
}

func defaultVisitorImplementations(types []string) string {
	methods := ""
	for _, typeStr := range types {
		parts := strings.Split(typeStr, ":")
		className := strings.TrimSpace(parts[0])
		methods += fmt.Sprintf(`
func (v *Visitor) Visit%s(expr *%s) {
	fmt.Println("Visiting %s")
}
`, className, className, className)
	}
	return methods
}
