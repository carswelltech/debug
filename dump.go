package debug

import (
	`encoding/json`
	`strings`
	`io`
	`os`
)
// SDump renders a variable as a string, in a readable form.
// Think Data::Dumper in perl.  Clearer than fmt.Sprintf("%+v"...)
// This function is NOT intended to be used in final or production code ... it is simply to visualize the contents of a variable.  The format of the output is not guaranteed to remain the same.
// Note: Only includes public members ...
func SDump(i interface{}) string {
	b := strings.Builder{}
	e := json.NewEncoder(&b)
	e.SetIndent(``,"\t")
	e.Encode(i)
	return b.String()
}
// Dump renders a variable as a string, just like SDump, but then writes it to one or more io.Writers (or, by default,to os.Stderr)
// This function is NOT intended to be used in final or production code ... it is simply to visualize the contents of a variable.  The format of the output is not guaranteed to remain the same.
func Dump(i interface{}, o ...io.Writer) {
	if len(o) == 0 {
		o = []io.Writer{os.Stderr}
	}
	for _,out := range o {
		e := json.NewEncoder(out)
		e.SetIndent(``,"\t")
		e.Encode(i)
	}
}
