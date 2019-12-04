const (
    {{range . -}}
    {{.name}} = {{.value}}
    {{end}}
)
