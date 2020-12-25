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
				if len(dictI[name]) == 0 {
					panic("NIL")
				}
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
			nameO := ""
			switch x := fc.Type.Params.List[1].Type.(*ast.StarExpr).X.(type) {
			case *ast.Ident:
				nameO = x.Name
			case *ast.SelectorExpr:
				nameO = x.Sel.Name
			default:
				panic("objO")
			}
			objO := dictO[nameO]
			if len(objO) == 0 {
				panic("NIL")
			}
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
			ioutil.WriteFile(mdname, []byte(doc), os.ModePerm)
		}

	}
}

func init() {
	flag.StringVar(&filename, "f", "", "filename")
	flag.Parse()
	log.Println(filename)
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
	"TransactionHashLE": {
		"type":        "string",
		"minLength":   64,
		"maxLength":   64,
		"pattern":     "[0-9a-f]{64}",
		"description": "transaction hash in little endian",
	},
	"AssetHash": {
		"type":        "string",
		"minLength":   64,
		"maxLength":   64,
		"pattern":     "[0-9a-f]{64}",
		"description": "asset hash in big endian",
	},
	"AssetHashLE": {
		"type":        "string",
		"minLength":   64,
		"maxLength":   64,
		"pattern":     "[0-9a-f]{64}",
		"description": "asset hash in little endian",
	},
	"BlockHash": {
		"type":        "string",
		"minLength":   64,
		"maxLength":   64,
		"pattern":     "[0-9a-f]{64}",
		"description": "block hash in big endian",
	},
	"BlockHashLE": {
		"type":        "string",
		"minLength":   64,
		"maxLength":   64,
		"pattern":     "[0-9a-f]{64}",
		"description": "block hash in little endian",
	},
	"KeyHash": {
		"type":        "string",
		"minLength":   64,
		"maxLength":   64,
		"pattern":     "[0-9a-f]{64}",
		"description": "sha256^2 hash of key in big endian",
	},
	"KeyHashLE": {
		"type":        "string",
		"minLength":   64,
		"maxLength":   64,
		"pattern":     "[0-9a-f]{64}",
		"description": "sha256^2 hash of key in little endian",
	},
	"ContractHash": {
		"type":        "string",
		"minLength":   40,
		"maxLength":   40,
		"pattern":     "[0-9a-f]{40}",
		"description": "contract hash in big endian",
	},
	"ContractHashLE": {
		"type":        "string",
		"minLength":   40,
		"maxLength":   40,
		"pattern":     "[0-9a-f]{40}",
		"description": "contract hash in little endian",
	},
	"HexKey": {
		"type":        "string",
		"pattern":     "[0-9a-f]*",
		"description": "hex encoded key",
	},
	"BlockHeight": {
		"type":        "integer",
		"minumum":     0,
		"default":     0,
		"description": "block height",
	},
	"OutputIndex": {
		"type":        "integer",
		"minumum":     0,
		"default":     0,
		"description": "output index in a transaction",
	},
}

var filterI = map[string]bool{
	"TransactionHash":   true,
	"TransactionHashLE": true,
	"AssetHash":         true,
	"AssetHashLE":       true,
	"BlockHash":         true,
	"BlockHashLE":       true,
	"ContractHash":      true,
	"ContractHashLE":    true,
	"BlockHeight":       false,
	"OutputIndex":       false,
}

var dictO = map[string]map[string]interface{}{
	"uint64": {
		"type": "integer",
	},
	"string": {
		"type":        "string",
		"pattern":     "[0-9a-f]+",
		"description": "binary data encoded in hex format",
	},
	"RawMessage": {
		"type": "object",
	},
}
