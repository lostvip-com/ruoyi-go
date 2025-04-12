<!DOCTYPE html>
{{$S := "{{"}}
{{$E := "}}"}}
<html lang="zh">
<head>
   {{$S}}template "header" (OssUrl){{$E}}
</head>
<body class="gray-bg">
<div class="container-div">
    <div class="row">
        <div class="col-sm-12 search-collapse">
            <form id="formId">
                <div class="select-list">
                    <ul>{{range $index, $column := .Columns}} {{if eq $column.IsQuery "1"}} {{if eq $column.HtmlType "input"}}
                        <li>
                            <p>{{$column.ColumnComment}}：</p>
                            <input type="text" name="{{$column.HtmlField}}"/>
                        </li>
                        {{else if eq $column.HtmlType "select" "radio"}} {{if ne $column.DictType ""}}
                            <li>
                                <p>{{$column.ColumnComment}}：</p>  {{$S}}DictSelect "{{$column.HtmlField}}" "{{$column.DictType}}"{{$E}}
                            </li> {{else}} <li>
                                <p>{{$column.ColumnComment}}：</p>
                                <select name="{{$column.HtmlField}}">
                                    <option value="">所有</option>
                                </select>
                            </li> {{end}}
                        {{else if eq $column.HtmlType "datetime"}}
                        <li class="select-time">
                            <p>{{$column.ColumnComment}}：</p>
                            <input type="text" class="time-input" id="startTime" placeholder="开始时间" name="beginTime"/>
                            <span>-</span>
                            <input type="text" class="time-input" id="endTime" placeholder="结束时间" name="endTime"/>
                        </li>  {{end}} {{end}} {{end}}
                        <li>
                            <a class="btn btn-primary btn-rounded btn-sm" onclick="$.table.search()"><i class="fa fa-search"></i>&nbsp;搜索</a>
                            <a class="btn btn-warning btn-rounded btn-sm" onclick="$.form.reset()"><i class="fa fa-refresh"></i>&nbsp;重置</a>
                        </li>
                    </ul>
                </div>
            </form>
        </div>

        <div class="btn-group-sm" id="toolbar" role="group">
            {{$S}}PermButton .uid "{{.PackageName}}:{{.BusinessName}}:add" "add()" "新增" "btn btn-success" "fa fa-plus"{{$E}}
            {{$S}}PermButton .uid "{{.PackageName}}:{{.BusinessName}}:remove" "$.operate.removeAll()" "删除" "btn btn-danger multiple disabled" "fa fa-remove"{{$E}}
            {{$S}}PermButton .uid "{{.PackageName}}:{{.BusinessName}}:export" "$.exportExcel()" "导出" "btn btn-warning" "fa fa-download"{{$E}}
        </div>
        <div class="col-sm-12 select-table table-striped">
            <table id="bootstrap-table"></table>
        </div>
    </div>
</div>
{{$S}}template "footer" (OssUrl){{$E}}
<script type="text/javascript">
    var editFlag   = '{{$S}}HasPerm .uid "{{.PackageName}}:{{.BusinessName}}:edit"{{$E}}';
    var removeFlag = '{{$S}}HasPerm .uid "{{.PackageName}}:{{.BusinessName}}:remove"{{$E}}';
    {{- range $index, $column := .Columns}}
      {{- if ne $column.DictType ""}}
      var {{$column.HtmlField}}Datas = {{$S}}DictType "$column.DictType" {{$E}};
      {{- end -}}
    {{- end}}
    var prefix = ctx + "/{{.PackageName}}/{{.BusinessName}}";

    $(function () {
        var options = {
            uniqueId:"{{.PkColumn.HtmlField}}",
            url: prefix + "/list{{.ClassName}}",
            createUrl: prefix + "/add{{.ClassName}}",
            removeUrl: prefix + "/remove{{.ClassName}}",
            exportUrl: prefix + "/export{{.ClassName}}",
            modalName: "{{.FunctionName}}",
            sortName:"{{.PkColumn.HtmlField}}",
            sortOrder:"desc",
            columns: [ { checkbox: true},
               {{- range $index, $column := .Columns -}}
                 {{ if eq $column.IsPk "1" }}
                       {field: '{{$column.HtmlField}}', title: '{{$column.ColumnComment}}', visible: false},
                 {{else if eq $column.IsList "1" -}}
                    {{if ne $column.DictType ""}}
                       {field: '{{$column.HtmlField}}', title: '{{$column.ColumnComment}}',
                                                formatter: function (value, row, index) {
                                                    return $.selectDictLabel({{$column.HtmlField}}Datas, value);
                                                } },
                    {{else -}}
                        {field: '{{$column.HtmlField}}', title: '{{$column.ColumnComment}}'},
                    {{end}}
                 {{- end -}}
               {{- end -}}
                {title: '操作',align: 'center',
                            formatter: function (value, row, index) {
                                var actions = [];
                                actions.push('<a class="btn btn-success btn-xs ' + editFlag + '" href="javascript:void(0)" onclick="editRight(\'' + row.id + '\')"><i class="fa fa-edit"></i>编辑</a> ');
                                actions.push('<a class="btn btn-danger btn-xs ' + removeFlag + '" href="javascript:void(0)" onclick="$.operate.remove(\'' + row.id + '\')"><i class="fa fa-remove"></i>删除</a>');
                                return actions.join('');
                            }
                }
            ]
        };
        $.table.init(options);
    });
</script>

<script>

   function add() {
           let url = prefix + "/preAdd{{.ClassName}}";
           lv_add_center("添加" + table.options.modalName, url,'800px','600px',function (){
              $.table.search()
           });
   }
   function editRight(id) {
        let url = prefix + "/preEdit{{.ClassName}}?id=" + id
        lv_edit_right("修改"+table.options.modalName ,url,'500px','100%',function (){
           $.table.search()
        });
   }
</script>

</body>
</html>