<!DOCTYPE html>
{{$S := "{{"}}
{{$E := "}}"}}
<html lang="zh">
<head>
    {{$S}}template "header" (OssUrl){{$E}}
</head>
<body class="white-bg">
<div class="wrapper wrapper-content   ibox-content">
    <form class="form-horizontal m" id="formEdit">
        <input name="{{.PkColumn.HtmlField}}" value="{{$S}}.{{.FunctionName}}.Id{{$E}}" type="hidden" />
        {{- range $index, $column := .Columns -}}
        {{if eq $column.HtmlType "input"}}
            <div class="form-group">
                <label class="col-sm-3 control-label">{{$column.ColumnComment}}:</label>
                <div class="col-sm-8">
                    <input name="{{$column.HtmlField}}" value="{{$S}}.{{$column.HtmlField}}.{{$column.GoField}}{{$E}}" class="form-control"  type="text" {{if eq $column.IsRequired "1"}}required{{end}}>
                </div>
            </div>
        {{else if eq $column.HtmlType "select"}}
            {{- if eq $column.DictType "" -}}
            <div class="form-group">
                <label class="col-sm-3 control-label">{{$column.ColumnComment}}:</label>
                <div class="col-sm-8">
                    <select name="{{$column.HtmlField}}" class="form-control m-b" {{if eq $column.IsRequired "1"}}required{{end}}>
                        <option value="">所有</option>
                    </select>
                    <span class="help-block m-b-none"><i class="fa fa-info-circle"></i> 代码生成请选择字典属性</span>
                </div>
            </div>
            {{- else -}}
            <div class="form-group">
                <label class="col-sm-3 control-label">{{$column.ColumnComment}}:</label>
                <div class="col-sm-8">
                    {{$S}}DictSelect "{{$column.HtmlField}}" "{{$column.DictType}}"{{$E}}
                </div>
            </div>
            {{- end}}
        {{else if eq $column.HtmlType "radio"}}  {{if eq $column.DictType ""}}
            <div class="form-group">
                <label class="col-sm-3 control-label">{{$column.ColumnComment}}:</label>
                <div class="col-sm-8">
                    <div class="radio-box">
                        <input type="radio" name="{{$column.HtmlField}}" value="" {{if eq $column.IsRequired "1"}}required{{end}}>
                        <label for="{{$column.HtmlField}}" text="未知"></label>
                    </div>
                    <span class="help-block m-b-none"><i class="fa fa-info-circle"></i> 代码生成请选择字典属性</span>
                </div>
            </div>  {{else}}
            <div class="form-group">
                <label class="col-sm-3 control-label">{{$column.ColumnComment}}:</label>
                <div class="col-sm-8">
                    {{$S}}DictRadio {{$column.HtmlField}} {{$column.HtmlField}}.{{$column.GoField}} {{$E}}
                </div>
            </div>  {{end}}
        {{else if eq $column.HtmlType "datetime"}}
            <div class="form-group">
                <label class="col-sm-3 control-label">{{$column.ColumnComment}}:</label>
                <div class="col-sm-8">
                    <div class="input-group date">
                        <span class="input-group-addon"><i class="fa fa-calendar"></i></span>
                        <input name="{{$column.HtmlField}}" value=" {{$S}}.{{$column.HtmlField}}.{{$column.GoField}}{{$E}}" class="form-control"  placeholder="yyyy-MM-dd" type="text" {{if eq $column.IsRequired "1"}}required{{end}}>
                    </div>
                </div>
            </div>
        {{else if eq $column.HtmlType "textarea"}}
            <div class="form-group">
                <label class="col-sm-3 control-label">{{$column.ColumnComment}}:</label>
                <div class="col-sm-8">
                    <textarea name="{{$column.HtmlField}}" class="form-control" {{if eq $column.IsRequired "1"}}required{{end}}>
                        {{$S}}.{{$column.HtmlField}}.{{$column.GoField}}{{$E}}
                    </textarea>
                </div>
            </div>
            {{end}}
       {{- end -}}
    </form>
</div>

{{$S}}template "footer" (OssUrl){{$E}}
<script type="text/javascript">
    var prefix = ctx + "/{{.PackageName}}/{{.BusinessName}}";
    $("#formEdit").validate({
        focusCleanup: true
    });

    function submitHandler(index,layero,callback) {
        console.log(callback)
        if ($.validate.form()) {
            let form =  $('#formEdit').serialize()
            $.operate.save(prefix + "/edit",form,function (result){
                $.modal.closeAll()
                console.log(callback)
                callback(result) //通知父页面刷新
            });
        }
    }

    {{range $index, $column := .Columns}}  {{if eq $column.HtmlType "datetime"}}
    $("input[name='{{$column.HtmlField}}']").datetimepicker({
        format: "yyyy-mm-dd",
        minView: "month",
        autoclose: true
    });
    {{end}}

    {{- end }}
    {{if eq .TplCategory "tree"}}
    /** {{.FunctionName}} -新增-选择父部门树 */
    function select{{.ClassName}}Tree() {
        var options = {
            title: '选择',
            width: "380",
            url: prefix + "/select{{.ClassName}}Tree/" + $("#treeId").val(),
            callBack: doSubmit
        };
        $.modal.openOptions(options);
    }

    function doSubmit(index, layero){
            var body = $.modal.getChildFrame(index);
            $("#treeId").val(body.find('#treeId').val());
            $("#treeName").val(body.find('#treeName').val());
            $.modal.close(index);
    }

{{- end -}}
</script>
</body>
</html>