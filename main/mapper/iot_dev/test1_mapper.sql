

-- name: ListTest1
select

from test1 t where 1=1 and t.del_flag=0


    order by
        {{if (eq .SortName "") }}
           
        {{else}}
           
        {{end}}

        {{if (eq .SortOrder "asc") }}
             asc
        {{else}}
             desc
        {{end}}



