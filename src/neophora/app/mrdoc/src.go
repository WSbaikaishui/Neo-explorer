package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range f.Decls {
		switch fc := v.(type) {
		case *ast.FuncDecl:
			title := fc.Name.Name
			objI := make(map[string]interface{})
			properties := make(map[string]interface{})
			required := make([]interface{}, 0)
			for _, v := range fc.Type.Params.List[0].Type.(*ast.StructType).Fields.List {
				name := v.Names[0].Name
				properties[name] = dictI[name]
				if filterI[name] {
					required = append(required, name)
				}
			}
			objI["type"] = "object"
			objI["properties"] = properties
			objI["required"] = required
			objI["additionalProperties"] = false
			jsI, _ := json.MarshalIndent(objI, "", "    ")
			input := fmt.Sprintf(jsonschema, string(jsI))
			objO := dictO[fc.Type.Params.List[1].Type.(*ast.StarExpr).X.(*ast.Ident).Name]
			jsO, _ := json.MarshalIndent(objO, "", "    ")
			output := fmt.Sprintf(jsonschema, string(jsO))
			desc := ""
			for _, v := range fc.Doc.List[1:] {
				if len(v.Text) > 3 {
					desc += v.Text[3:] + "\n"
				} else {
					desc += "\n"
				}
			}
			doc := fmt.Sprintf(template, title, desc, input, output)
			base := path.Base(filename)
			mdname := "./" + base[4:len(base)-3] + ".md"
			log.Println(mdname)
			ioutil.WriteFile(mdname, []byte(doc), os.ModePerm)
		}

	}
}

func init() {
	flag.StringVar(&filename, "f", "", "filename")
	flag.Parse()
}

var filename string

const template = `# %s

%s

## Input

%s

## Output

%s

`
const jsonschema = "```json\n%s\n```"

var dictI = map[string]map[string]interface{}{
	"TransactionHash": {
		"type":        "string",
		"minLength":   64,
		"maxLength":   64,
		"pattern":     "[0-9a-f]{64}",
		"description": "transaction hash in big endian",
	},
}

var filterI = map[string]bool{
	"TransactionHash": true,
}

var dictO = map[string]interface{}{
	"uint64": map[string]interface{}{
		"type": "integer",
	},
}
