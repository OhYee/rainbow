const (
    {{range . -}}
    {{.name}} = {{.value}}
    {{end}}
    {{range . -}}
    light{{upperFirstChar .name}} = {{.value}} + 60
    {{end}}
)
