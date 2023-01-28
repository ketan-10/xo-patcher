package wire

var RepositorySet = wire.NewSet(
    {{ range . }}
        repo.New{{ camelCase .Table.TableName }}Repository,
    {{- end }}
)