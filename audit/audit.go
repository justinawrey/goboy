package audit

import (
	"html/template"
	"log"
	"os"

	"github.com/justinawrey/goboy/gb"
)

type instructionData struct {
	Instructions8  map[uint8]gb.Instruction
	Instructions16 map[uint16]gb.Instruction
}

func Generate() {
	data := instructionData{
		Instructions8:  gb.InstructionTable8,
		Instructions16: gb.InstructionTable16,
	}

	const tmpl = `
	<h1>CPU instruction completion status</h1>

	<h2>8-bit Instructions</h2>
	<ul>
			{{range $_, $instruction := .Instructions8}}
					{{if $instruction.Implemented }}
							<li style="color: green">{{$instruction.Encoding}}: {{$instruction.Mnemonic}}</li>
					{{else}}
							<li style="color: red">{{$instruction.Encoding}}: {{$instruction.Mnemonic}}</li>
					{{end}}
			{{end}}
	</ul>

	<h2>16-bit Instructions</h2>
	<ul>
			{{range $_, $instruction := .Instructions16}}
					{{if $instruction.Implemented }}
							<li style="color: green">{{$instruction.Encoding}}: {{$instruction.Mnemonic}}</li>
					{{else}}
							<li style="color: red">{{$instruction.Encoding}}: {{$instruction.Mnemonic}}</li>
					{{end}}
			{{end}}
	</ul>`

	t, err := template.New("audit").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("audit.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = t.Execute(f, data)
	if err != nil {
		log.Fatal(err)
	}
}
