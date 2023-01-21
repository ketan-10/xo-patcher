{{- $name := joinWith "_" .TableName .ColumnName -}}
{{ $camelName := camelCase $name }}
{{- $shortName := shortName $camelName -}}

type {{ $camelName }} uint16

const (
{{ range .Values }}
    {{ $camelName -}}{{- camelCase . }} {{ $camelName }} = iota 
{{ end }}
)

func ({{ $shortName }} {{ $camelName }}) String() string {
    var value string

    switch {{ $shortName }} {
        {{ range .Values }}
        case {{ $camelName -}}{{- camelCase . }}:
            value = "{{.}}" 
        {{ end }}
    }

    return value
}

func ({{ $shortName }} {{ $camelName }}) GoString() string {
    return {{ $shortName }}.String()
}



