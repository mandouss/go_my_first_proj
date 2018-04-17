## String in go

### string template in go

#### first example

In Go, we use the template package and methods like “Parse”, “ParseFile”, “Execute” to load a template from a string or file and then perform the merge. The content to merge is within a defined type and that has exported fields, i.e. fields within the struct that are used within the template have to start with a capital letter.

```go
package main

import (
        "log"
        "os"
        "text/template"
)

type Student struct {
        //exported field since it begins
        //with a capital letter
        Name string
}

func main() {
        //define an instance
        s := Student{"Satish"}

        //create a new template with some name
        tmpl := template.New("test")

        //parse some content and generate a template
        tmpl, err := tmpl.Parse("Hello {{.Name}}!")
        if err != nil {
                log.Fatal("Parse: ", err)
                return
        }

        //merge template 'tmpl' with content of 's'
        err1 := tmpl.Execute(os.Stdout, s)
        if err1 != nil {
                log.Fatal("Execute: ", err1)
                return
        }
}
```

The output is:

```
Hello Satish!
```

**Note**:

- “New” allocates a new template with the given name.
- “Parse” parses a string into a template.
- To include the content of a field within a template, enclose it within curly braces and add a dot at the beginning. E.g. if Name is a field within a struct and its value needs to be substituted while merging, then include the text “{{.Name}}” in the template. **Do remember that the field name has to be present and it should also be exported (i.e. it should begin with a capital letter in the type definition), or there could be errors.** All text outside “{{.Name}}” is copied to the output unchanged.
- We have used the predefined variable “os.Stdout” which refers to the standard output to print out the merged data — “os.Stdout” implements “io.Writer”.
- “Execute” applies a parsed template to the specified data object, and writes the output to “os.Stdout”.

#### second example

```go
package main

import (
        "log"
        "os"
        "text/template"
)

type Person struct {
        Name   string
        Emails []string
}

const tmpl = `The name is {{.Name}}.
{{range .Emails}}
    His email id is {{.}}
{{end}}
`

func main() {
        person := Person{
                Name:   "Satish",
                Emails: []string{"satish@rubylearning.org", "satishtalim@gmail.com"},
        }

        t := template.New("Person template")

        t, err := t.Parse(tmpl)
        if err != nil {
                log.Fatal("Parse: ", err)
                return
        }

        err = t.Execute(os.Stdout, person)
        if err != nil {
                log.Fatal("Execute: ", err)
                return
        }
}
```

The output is:

```
The name is Satish.
    His email id is satish@rubylearning.org
    His email id is satishtalim@gmail.com
```

In the above program, we have “{{range .Emails}}”. With “range” the current object “.” is set to the successive elements of the array or slice Emails.

**Variables**

The template package allows you to define and use variables. In the above example, how would we print each person’s email address prefixed by their name? Let’s modify the above program.

In the code snippet:

```
{{range .Emails}}
    {{.}}
{{end}}
```

We cannot access the “Name” field as “.” is now traversing the array elements and the “Name” is outside of this scope. The solution is to save the value of the “Name” field in a variable that can be accessed anywhere in its scope. Variables in templates are prefixed by $. So we write:

```
{{$name := .Name}}
{{range .Emails}}
    Name is {{$name}}, email is {{.}}
{{end}}
```

The modified program, named “new_person.go” is:

```go
package main

import (
        "log"
        "os"
        "text/template"
)

type Person struct {
        Name   string
        Emails []string
}

const tmpl = `{{$name := .Name}}
{{range .Emails}}
    Name is {{$name}}, email is {{.}}
{{end}}
`

func main() {
        person := Person{
                Name:   "Satish",
                Emails: []string{"satish@rubylearning.org", "satishtalim@gmail.com"},
        }

        t := template.New("Person template")

        t, err := t.Parse(tmpl)
        if err != nil {
                log.Fatal("Parse: ", err)
                return
        }

        err = t.Execute(os.Stdout, person)
        if err != nil {
                log.Fatal("Execute: ", err)
                return
        }
}
```

The output is:

```
Name is Satish, email is satish@rubylearning.org
Name is Satish, email is satishtalim@gmail.com
```

The Go template package is useful for certain kinds of text transformations involving inserting values of objects. It does not have the power of, say, regular expressions, but is faster and in many cases will be easier to use than regular expressions.

### templ to string

To achieve this we can simply use *any types* that implement io.Writer interface. For this strings case, we can use a buffer's pointer to store html template execution result then parse them to string.

```go
t := template.New("action")

var err error
t, err = t.ParseFiles("path/to/action.html")
if err != nil {
    return err
}

key := "some strings"

data := struct{
    Key string
}{
    Key: key
}

var tpl bytes.Buffer
if err := t.Execute(&tpl, data); err != nil {
    return err
}

result := tpl.String()
```



#### source : [This is source ](https://medium.com/@IndianGuru/understanding-go-s-template-package-c5307758fab0)

#### [This is source for templ to string ](https://coderwall.com/p/ns60fq/simply-output-go-html-template-execution-to-strings)





