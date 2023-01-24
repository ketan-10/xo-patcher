{{- $tableNameCamel := camelCase .Table.TableName -}}
{{- $shortName := shortName $tableNameCamel -}}

type I{{ $tableNameCamel }}Repository interface {
    I{{ $tableNameCamel }}RepositoryQueryBuilder
    
    Insert{{ $tableNameCamel }}(ctx context.Context, {{ $shortName }} entities.{{ $tableNameCamel }}Create) (*entities.{{ $tableNameCamel }}, error)
    Insert{{ $tableNameCamel }}WithSuffix(ctx context.Context, {{ $shortName }} entities.{{ $tableNameCamel }}Create, suffix sq.Sqlizer) (*entities.{{ $tableNameCamel }}, error)
    Insert{{ $tableNameCamel }}IDResult(ctx context.Context, {{ $shortName }} entities.{{ $tableNameCamel }}Create, suffix sq.Sqlizer) (int64, error)

    Update{{ $tableNameCamel }}ByFields(ctx context.Context, id int, {{ $shortName }} entities.{{ $tableNameCamel }}Update) (*entities.{{ $tableNameCamel }}, error)
    Update{{ $tableNameCamel }}(ctx context.Context, {{ $shortName }} entities.{{ $tableNameCamel }}) (*entities.{{ $tableNameCamel }}, error)
    
    Delete{{ $tableNameCamel }}(ctx context.Context, {{ $shortName }} entities.{{ $tableNameCamel }}) error
    Delete{{ $tableNameCamel }}ByID(ctx context.Context, id int) (bool, error)
    
    FindAll{{ $tableNameCamel }}(ctx context.Context, {{ $shortName }} entities.{{ $tableNameCamel }}Filter, pagination *entities.Pagination) (entities.List{{ $tableNameCamel }}, error)
    FindAll{{ $tableNameCamel }}WithSuffix(ctx context.Context,{{ $shortName }} entities.{{ $tableNameCamel }}Filter, pagination *entities.Pagination, suffixes ...sq.Sqlizer) (entities.List{{ $tableNameCamel }}, error)

}


type I{{ $tableNameCamel }}RepositoryQueryBuilder interface {
    FindAll{{ $tableNameCamel }}BaseQuery(ctx context.Context, filter *entities.{{ $tableNameCamel }}Filter, fields string, suffix ...sq.Sqlizer) (*sq.SelectBuilder, error)
    AddPagination(ctx contex.Context, qb *sq.SelectBuilder, pagination *entities.Pagination) (*sq.SelectBuilder, error)
}

type {{ $tableNameCamel }}Repository struct {
    DB db_manager.IDb
    QueryBuilder I{{ $tableNameCamel }}RepositoryQueryBuilder
}

type {{ $tableNameCamel }}RepositoryQueryBuilder struct {
}

var New{{ $tableNameCamel }}Repository = wire.WireSet(
    wire.Struct(new({{ $tableNameCamel }}Repository), "*"),
    wire.Struct(new({{ $tableNameCamel }}RepositoryQueryBuilder), "*"),
    wire.Bind(new(I{{ $tableNameCamel }}Repository), new({{ $tableNameCamel }}Repository), "*"),
    wire.Bind(new(I{{ $tableNameCamel }}RepositoryQueryBuilder), new({{ $tableNameCamel }}RepositoryQueryBuilder), "*"),
)

func ({{ $shortName }}r *{{ $tableNameCamel }}Repository) Insert{{ $tableNameCamel }}(ctx context.Context, {{ $shortName }} entities.{{ $tableNameCamel }}Create) (*entities.{{ $tableNameCamel }}, error) {
    return {{ $shortName }}.Insert{{ $tableNameCamel }}WithSuffix(ctx, {{ $shortName }}, nil)
}

func ({{ $shortName }}r *{{ $tableNameCamel }}Repository) Insert{{ $tableNameCamel }}WithSuffix(ctx context.Context, {{ $shortName }} entities.{{ $tableNameCamel }}Create, suffix sq.Sqlizer) (*entities.{{ $tableNameCamel }}, error) {
    var err error
    
    id, err := {{ $shortName }}r.Insert{{ $tableNameCamel }}IDResult(ctx, {{ $shortName }}, suffix)
    if err != nil {
        return nil, err
    }
    new{{ $shortName }} := entities.{{ $tableNameCamel }}{}
    qb := sq.Select("*").From(` {{ .Table.TableName }}`)

    qb.Where(sq.Eq("`id`": id))
    err = {{ $tableNameCamel }}Repository.DB.Get(ctx, &new{{ $shortName }}, qb)

    if err != nil {
        return nil, err
    }
    return &new{{ $shortName }}, nil
}

func ({{ $shortName }}r *{{ $tableNameCamel }}Repository) Insert{{ $tableNameCamel }}IDResult(ctx context.Context, {{ $shortName }} entities.{{ $tableNameCamel }}Create, suffix sq.Sqlizer) (int64, error) {
    var err error

    qb := sq.Insert("`{{ $tableNameCamel }}`").Columns(
        {{- range .Table.Columns }}
            "`{{ .ColumnName }}`",
        {{- end }}
    ).Values(
        {{- range .Table.Columns }}
            {{ $shortName }}.{{ camelCase .ColumnName }}
        {{- end }}
    )
    if suffix != nil {
        suffixQuery, suffixArgs, suffixErr := suffix.ToSql()
        if suffixErr != nil {
            return 0, suffixErr
        }
        qb.Suffix(suffixQuery, suffixArgs...)
    }

    // run query
	res, err := {{ $shortName }}r.Db.Exec(ctx, qb)
	if err != nil {
		return 0, err
	}

    // retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

    return id, nil
} 

func ({{ $shortName }}r *{{ $tableNameCamel }}Repository) Update{{ $tableNameCamel }}ByFields(ctx context.Context, id int, {{ $shortName }} entities.{{ $tableNameCamel }}Update) (*entities{{ $shortName }}, error) {
    var err error 

    updateMap := map[string]interface{}{}
    {{- range .Table.Columns }}
        if ({{ $shortName }}r.{{ camelCase .ColumnName }} != nil) {
            updatemap["`{{ .ColumnName }}`"] = *{{ $shortName }}r.{{ camelCase .ColumnName }}
        }
    {{- end }}

    qb := sq.Update(`{{ $tableNameCamel  }}`).SetMap(updateMap).Where(sq.Eq{"`id`": id})

    _, err := {{ $shortName }}r.Db.Exec(ctx, qb)
    if err != nil {
        return nil, err
    }

    selectQb := sq.Select("*").From("`{{ $.Table.TableName }}`")

    selectQb = selectQb.Where(sq.Eq{"`id`": id})

    result := entities.{{ $tableNameCamel }}{}
    err = {{ $shortName }}r.Db.Get(ctx, &result, selectQb)
    if err != nil {
        return nil, err
    }

    return &result, nil

}

func ({{ $shortName }}r *{{ $tableNameCamel }}Repository) Delect{{ $tableNameCamel }}(ctx context.Context, {{ $shortName }} entities.{{ $tableNameCamel }}) (error) {
    _, err := {{ $shortName }}r.Delete{{ $tableNameCamel }}ByID(ctx, {{ $shortName }}.ID)
    return err
}


func ({{ $shortName }}r *{{ $tableNameCamel }}Repository) Delect{{ $tableNameCamel }}ByID(ctx context.Context, id int) (bool, error) {
    var err error

    qb := sq.Update("`{{ .Table.TableName }}`").Set("active", false)

    qb = qb.Where(sq.Eq{"`id`": id})

    _, err = {{ $shortName }}r.Db.Exec(ctx, qb)
    if err != nil {
        return false, err
    } 
    return true, nil
}

func ({{ $shortName }} *{{ $tableNameCamel }}Repository) FindAll{{ $tableNameCamel }}BaseQuery(ctx context.Context, filter *entities.{{ $tableNameCamel }}Filter, fields string, suffixes ...sq.Sqlizer) (*sq.SelectBuilder, error) {
    return {{ $shortName }}.QueryBuilder.FindAll{{ $tableNameCamel }}BaseQuery(ctx, filter, fields, suffixes...)
}

func ({{ $shortName }} *{{ $tableNameCamel }}Repository) FindAll{{ $tableNameCamel }}BaseQuery(ctx context.Context, filter *entities.{{ $tableNameCamel }}Filter, fields string, suffixes ...sq.Sqlizer) (*sq.SelectBuilder, error) {
    var err error
    qb := sq.Select(fields).From("`{{ .Table.TableName }}`")
    if filter != nil {
        if filter.Active == nil {
            if qb, err = addFilter(qb, "{{ `.Table.TableName` }}.`active`", entities.FilterOnField{ {entities.Eq: true} }); err != nil {
                return qb, err
            }
        } else {
            if qb, err = addFilter(qb, "{{ `.Table.TableName` }}.`active`", filter.Active); err != nil {
                return qb, err
            }
        }
    {{- range .Table.Columns }}
        {{- if ne .ColumnName "active" }}
            if qb, err = addFilter(qb, "{{ `$.Table.TableName` }}.`{{ .ColumnName }}`", filter.{{ camelCase .ColumnName }} }); err != nil {
                return qb, err
            }
        {{- end }}
    {{- end }}
    } else {
        if qb, err = addFilter(qb, "{{ `.Table.TableName` }}.`active`", entities.FilterOnField{ {entities.Eq: true} }); err != nil {
            return qb, err
        }
    }

    for _, suffix := range suffixes {
        query, args, err := suffix.ToSql()
        if err != nil {
            return qb, err
        }
        qb.Suffix(query, args...)
    }

}

func ({{ $shortName }}r *{{ $tableNameCamel }}Repository) FindAll{{ $tableNameCamel }}BaseQuery(ctx context.Context, filter *entities.{{ $tableNameCamel }}Filter, fields string, suffixes ...sq.Sqlizer) (*sq.SelectBuilder, error) {
    return {{ $shortName }}.QueryBuilder.FindAll{{ $tableNameCamel }}BaseQuery(ctx, filter, fields, suffixes...)
}


func ({{ $shortName }} *{{ $tableNameCamel }}RepositoryQueryBuilder) AddPagination(ctx context.Context, qb *sq.SelectBuilder, pagination *entities.Pagination) (*sq.SelectBuilder, error) {
    fields := []string {
        {{- range .Table.Columns }}   
            "{{ .ColumnName }}",
        {{- end }}
    }
    return AddPagination(qb, pagination, "{{ .Table.TableName }}", fields)
}

func ({{ $shortName }}r *{{ $tableNameCamel }}Repository) FindAll{{ $tableNameCamel }}(ctx context.Context, filter *entities.{{ $tableNameCamel }}Filter, pagination *entities.Pagination) (list entities.List{{ $tableNameCamel }}, err error) {
    return {{ $shortName }}r.FindAll{{ $tableNameCamel }}WithSuffix(ctx, filter, pagination)
}

func ({{ $shortName }}r *{{ $tableNameCamel }}Repository) FindAll{{ $tableNameCamel }}WithSuffix(ctx context.Context, filter *entities.{{ $tableNameCamel }}Filter, pagination *entities.Pagination, suffixes ...sq.Sqlizer) (list entities.List{{ $tableNameCamel }}, err error) {
    qb, err := {{ $shortName }}r.FindAll{{ $tableNameCamel }}BaseQuery(ctx, filter, "`{{ .Table.TableName }}`.*", suffixes...)
    if err != nil {
        return nil, err
    }
    qb, err = {{ $shortName }}r.AddPagination(ctx, qb, pagination)
    if err != nil {
        return nil, err
    }

    err = {{ $shortName }}.Db.Select(ctx, &list.Data, qb)

    if err != nil {
        return list, err
    }

    if pagination == nil || pagination.PerPage == nil || pagination.Page == nil {
        list.TotalCount = len(list.Data)
        return list, nil
    }

    var listMeta entities.ListMetadata
    if qb, err = {{ $shortName }}.FindAll{{ $tableNameCamel }}BaseQuery(ctx, filter, "COUNT(1) AS count"); err != nil {
        return entities.List{{ $tableNameCamel }}{}, err
    }
    if filter != nil && len(filter.GroupBys) > 0 {
        qb = sq.Select("COUNT(1) AS count").FromSelect(qb, "a")
    }
    err = {{ $shortName }}.Db.Get(ctx, &listMeta, qb)

    list.TotalCount = listMeta.Count

    return list, err
}
