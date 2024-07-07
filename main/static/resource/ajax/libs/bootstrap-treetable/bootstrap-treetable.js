/**
 * 基于bootstrapTreeTable/bootstrap-table-treegrid修改
 * Copyright (c) 2019 yunfan
 */
(function($) {
    "use strict";

    $.fn.bootstrapTreeTable = function(options, param) {
        var target = $(this).data('bootstrap.tree.table');
        target = target ? target : $(this);
        // 如果是调用方法
        if (typeof options == 'string') {
            return $.fn.bootstrapTreeTable.methods[options](target, param);
        }
        // 如果是初始化组件
        options = $.extend({}, $.fn.bootstrapTreeTable.defaults, options || {});
        target.hasSelectItem = false;// 是否有radio或checkbox
        target.data_list = null; //用于缓存格式化后的数据-按父分组
        target.data_obj = null; //用于缓存格式化后的数据-按id存对象
        target.hiddenColumns = []; //用于存放被隐藏列的field
        target.lastAjaxParams; //用户最后一次请求的参数
        target.isFixWidth=false; //是否有固定宽度
        // 初始化
        var init = function() {
            // 初始化容器
            initContainer();
            // 初始化工具栏
            initToolbar();
            // 初始化表头
            initHeader();
            // 初始化表体
            initBody();
            // 初始化数据服务
            initServer();
            // 动态设置表头宽度
            autoTheadWidth(true);
            // 缓存target对象
            target.data('bootstrap.tree.table', target);
        }
        // 初始化容器
        var initContainer = function() {
            // 在外层包装一下div，样式用的bootstrap-table的
            var $main_div = $("<div class='bootstrap-tree-table'></div>");
            var $treetable = $("<div class='treetable-table'></div>");
            target.before($main_div);
            $main_div.append($treetable);
            $treetable.append(target);
            target.addClass("table");
            if (options.striped) {
                target.addClass('table-striped');
            }
            if (options.bordered) {
                target.addClass('table-bordered');
            }
            if (options.hover) {
                target.addClass('table-hover');
            }
            if (options.condensed) {
                target.addClass('table-condensed');
            }
            target.html("");
        }
        // 初始化工具栏
        var initToolbar = function() {
            var $toolbar = $("<div class='treetable-bars'></div>");
            if (options.toolbar) {
                $(options.toolbar).addClass('tool-left');
                $toolbar.append($(options.toolbar));
            }
            var $rightToolbar = $('<div class="btn-group tool-right">');
            $toolbar.append($rightToolbar);
            target.parent().before($toolbar);
            //  是否显示检索信息
            if (options.showSearch) {
                var $searchBtn = $('<button class="btn btn-default btn-outline" type="button" aria-label="lv_sql" title="搜索"><i class="glyphicon glyphicon-lv_sql"></i></button>');
                $rightToolbar.append($searchBtn);
                registerSearchBtnClickEvent($searchBtn);
            }
            // 是否显示刷新按钮
            if (options.showRefresh) {
                var $refreshBtn = $('<button class="btn btn-default btn-outline" type="button" aria-label="refresh" title="刷新"><i class="glyphicon glyphicon-repeat"></i></button>');
                $rightToolbar.append($refreshBtn);
                registerRefreshBtnClickEvent($refreshBtn);
            }
            // 是否显示列选项
            if (options.showColumns) {
                var $columns_div = $('<div class="btn-group pull-right" title="列"><button type="button" aria-label="columns" class="btn btn-default btn-outline dropdown-toggle" data-toggle="dropdown" aria-expanded="false"><i class="glyphicon glyphicon-list"></i> <span class="caret"></span></button></div>');
                var $columns_ul = $('<ul class="dropdown-menu columns" role="menu"></ul>');
                $.each(options.columns, function(i, column) {
                    if (column.field != 'selectItem') {
                        var _li = null;
                        if(typeof column.visible == "undefined"||column.visible==true){
                            _li = $('<li role="menuitem"><label><input type="checkbox" checked="checked" data-field="'+column.field+'" value="'+column.field+'" > '+column.title+'</label></li>');
                        }else{
                            _li = $('<li role="menuitem"><label><input type="checkbox" data-field="'+column.field+'" value="'+column.field+'" > '+column.title+'</label></li>');
                            target.hiddenColumns.push(column.field);
                        }
                        $columns_ul.append(_li);
                    }
                });
                $columns_div.append($columns_ul);
                $rightToolbar.append($columns_div);
                // 注册列选项事件
                registerColumnClickEvent();
            }else{
                $.each(options.columns, function(i, column) {
                    if (column.field != 'selectItem') {
                        if(!(typeof column.visible == "undefined"||column.visible==true)){
                            target.hiddenColumns.push(column.field);
                        }
                    }
                });
            }
        }
        // 初始化隐藏列
        var initHiddenColumns = function(){
            $.each(target.hiddenColumns, function(i, field) {
                target.find("."+field+"_cls").hide();
            });
        }
        // 初始化表头
        var initHeader = function() {
            var $thr = $('<tr></tr>');
            $.each(options.columns, function(i, column) {
                var $th = null;
                // 判断有没有选择列
                if (i == 0 && column.field == 'selectItem') {
                    target.hasSelectItem = true;
                    $th = $('<th style="width:36px"></th>');
                } else {
                    $th = $('<th style="' + ((column.width) ? ('width:' + column.width) : '') + '" class="' + column.field + '_cls"></th>');
                }
                if((!target.isFixWidth)&& column.width){
                    target.isFixWidth = column.width.indexOf("px")>-1?true:false;
                }
                $th.text(column.title);
                $thr.append($th);
            });
            var $thead = $('<thead class="treetable-thead"></thead>');
            $thead.append($thr);
            target.append($thead);
        }
        // 初始化表体
        var initBody = function() {
            var $tbody = $('<tbody class="treetable-tbody"></tbody>');
            target.append($tbody);
            // 默认高度
            if (options.height) {
                $tbody.css("height", options.height);
            }
        }
        // 初始化数据服务
        var initServer = function(parms) {
            // 加载数据前先清空
            target.data_list = {};
            target.data_obj = {};
            var $tbody = target.find("tbody");
            // 添加加载loading
            var $loading = '<tr><td colspan="' + options.columns.length + '"><div style="display: block;text-align: center;">🕗...</div></td></tr>'
            $tbody.html($loading);
            if (options.url) {
                $.ajax({
                    type: options.type,
                    url: options.url,
                    data: parms ? parms : options.ajaxParams,
                    dataType: "JSON",
                    success: function(data, textStatus, jqXHR) {
                    	data = calculateObjectValue(options, options.responseHandler, [data], data);
                        renderTable(data);
                    },
                    error: function(xhr, textStatus) {
                        var _errorMsg = '<tr><td colspan="' + options.columns.length + '"><div style="display: block;text-align: center;">' + xhr.responseText + '</div></td></tr>'
                        $tbody.html(_errorMsg);
                    },
                });
            } else {
                renderTable(options.data);
            }
        }
        // 加载完数据后渲染表格
        var renderTable = function(data) {
            var $tbody = target.find("tbody");
            // 先清空
            $tbody.html("");
            if (!data || data.length <= 0) {
                var _empty = '<tr><td colspan="' + options.columns.length + '"><div style="display: block;text-align: center;">没有找到匹配的记录</div></td></tr>'
                $tbody.html(_empty);
                return;
            }
            // 缓存并格式化数据
            formatData(data);
            // 获取所有根节点
            var rootNode = target.data_list["_root_"];
            // 开始绘制
            if (rootNode) {
                $.each(rootNode, function(i, item) {
                    var _child_row_id = "row_id_" + i
                    recursionNode(item, 1, _child_row_id, "row_root");
                });
            }
            // 下边的操作主要是为了查询时让一些没有根节点的节点显示
            $.each(data, function(i, item) {
                if (!item.isShow) {
                    var tr = renderRow(item, false, 1, "", "");
                    $tbody.append(tr);
                }
            });
            target.append($tbody);
            registerExpanderEvent();
            registerRowClickEvent();
            initHiddenColumns();
            // 动态设置表头宽度
            autoTheadWidth()
        }
        // 动态设置表头宽度
        var autoTheadWidth = function(initFlag) {
            if(options.height>0){
                var $thead = target.find("thead");
                var $tbody = target.find("tbody");
                var borderWidth = parseInt(target.css("border-left-width")) + parseInt(target.css("border-right-width"))
                
                $thead.css("width", $tbody.children(":first").width());
                if(initFlag){
                    var resizeWaiter = false;
                    $(window).resize(function() {
                        if(!resizeWaiter){
                            resizeWaiter = true;
                            setTimeout(function(){
                                if(!target.isFixWidth){
                                    $tbody.css("width", target.parent().width()-borderWidth);
                                }
                                $thead.css("width", $tbody.children(":first").width());
                                resizeWaiter = false;
                            }, 300);
                        }
                    });
                }
            }
        
        }
        // 缓存并格式化数据
        var formatData = function(data) {
            var _root = options.rootIdValue ? options.rootIdValue : null;
            var firstCode = data[0][options.parentCode];
            $.each(data, function(index, item) {
                // 添加一个默认属性，用来判断当前节点有没有被显示
                item.isShow = false;
                // 这里兼容几种常见Root节点写法
                // 默认的几种判断
                var _defaultRootFlag = item[options.parentCode] == '0' ||
                    item[options.parentCode] == 0 ||
                    item[options.parentCode] == null ||
                    item[options.parentCode] == firstCode ||
                    item[options.parentCode] == '';
                if (!item[options.parentCode] || (_root ? (item[options.parentCode] == options.rootIdValue) : _defaultRootFlag)) {
                    if (!target.data_list["_root_"]) {
                        target.data_list["_root_"] = [];
                    }
                    if (!target.data_obj["id_" + item[options.code]]) {
                        target.data_list["_root_"].push(item);
                    }
                } else {
                    if (!target.data_list["_n_" + item[options.parentCode]]) {
                        target.data_list["_n_" + item[options.parentCode]] = [];
                    }
                    if (!target.data_obj["id_" + item[options.code]]) {
                        target.data_list["_n_" + item[options.parentCode]].push(item);
                    }
                }
                target.data_obj["id_" + item[options.code]] = item;
            });
        }
        // 递归获取子节点并且设置子节点
        var recursionNode = function(parentNode, lv, row_id, p_id) {
            var $tbody = target.find("tbody");
            var _ls = target.data_list["_n_" + parentNode[options.code]];
            var $tr = renderRow(parentNode, _ls ? true : false, lv, row_id, p_id);
            $tbody.append($tr);
            if (_ls) {
                $.each(_ls, function(i, item) {
                    var _child_row_id = row_id + "_" + i
                    recursionNode(item, (lv + 1), _child_row_id, row_id)
                });
            }
        };
        // 绘制行
        var renderRow = function(item, isP, lv, row_id, p_id) {
            // 标记已显示
            item.isShow = true;
            item.row_id = row_id;
            item.p_id = p_id;
            item.lv = lv;
            var $tr = $('<tr id="' + row_id + '" pid="' + p_id + '"></tr>');
            var _icon = options.expanderCollapsedClass;
            if (options.expandAll) {
                $tr.css("display", "table");
                _icon = options.expanderExpandedClass;
            } else if (lv == 1) {
                $tr.css("display", "table");
                _icon = (options.expandFirst) ? options.expanderExpandedClass : options.expanderCollapsedClass;
            } else if (lv == 2) {
                if (options.expandFirst) {
                    $tr.css("display", "table");
                } else {
                    $tr.css("display", "none");
                }
                _icon = options.expanderCollapsedClass;
            } else {
                $tr.css("display", "none");
                _icon = options.expanderCollapsedClass;
            }
            $.each(options.columns, function(index, column) {
                // 判断有没有选择列
                if (column.field == 'selectItem') {
                    target.hasSelectItem = true;
                    var $td = $('<td style="text-align:center;width:36px"></td>');
                    if (column.radio) {
                        var _ipt = $('<input name="select_item" type="radio" value="' + item[options.code] + '"></input>');
                        $td.append(_ipt);
                    }
                    if (column.checkbox) {
                        var _ipt = $('<input name="select_item" type="checkbox" value="' + item[options.code] + '"></input>');
                        $td.append(_ipt);
                    }
                    $tr.append($td);
                } else {
                    var $td = $('<td name="' + column.field + '" class="' + column.field + '_cls"></td>');
                    if(column.width){
                        $td.css("width",column.width);
                    }
                    if(column.align){
                        $td.css("text-align",column.align);
                    }
                    if(options.expandColumn == index){
                        $td.css("text-align","left");
                    }
                    if(column.valign){
                        $td.css("vertical-align",column.valign);
                    }
                    if(options.showTitle){
                        $td.addClass("ellipsis");
                    }
                    // 增加formatter渲染
                    if (column.formatter) {
                        $td.html(column.formatter.call(this, getItemField(item, column.field), item, index));
                    } else {
                        if(options.showTitle){
                            // 只在字段没有formatter时才添加title属性
                            $td.attr("title",item[column.field]);
                        }
                        $td.text(getItemField(item, column.field));
                    }
                    if (options.expandColumn == index) {
                        if (!isP) {
                            $td.prepend('<span class="treetable-expander"></span>')
                        } else {
                            $td.prepend('<span class="treetable-expander ' + _icon + '"></span>')
                        }
                        for (var int = 0; int < (lv - 1); int++) {
                            $td.prepend('<span class="treetable-indent"></span>')
                        }
                    }
                    $tr.append($td);
                }
            });
            return $tr;
        }
        // 检索信息按钮点击事件
        var registerSearchBtnClickEvent = function(btn) {
            $(btn).off('click').on('click', function () {
                $(".lv_sql-collapse").slideToggle();
            });
        }
        // 注册刷新按钮点击事件
        var registerRefreshBtnClickEvent = function(btn) {
            $(btn).off('click').on('click', function () {
                target.refresh();
            });
        }
        // 注册列选项事件
        var registerColumnClickEvent = function() {
            $(".bootstrap-tree-table .treetable-bars .columns label input").off('click').on('click', function () {
                var $this = $(this);
                if($this.prop('checked')){
                    target.showColumn($(this).val());
                }else{
                    target.hideColumn($(this).val());
                }
            });
        }
        // 注册行点击选中事件
        var registerRowClickEvent = function() {
            target.find("tbody").find("tr").unbind();
            target.find("tbody").find("tr").click(function() {
                if (target.hasSelectItem) {
                    var _ipt = $(this).find("input[name='select_item']");
                    if (_ipt.attr("type") == "radio") {
                        _ipt.prop('checked', true);
                        target.find("tbody").find("tr").removeClass("treetable-selected");
                        $(this).addClass("treetable-selected");
                    } else if (_ipt.attr("type") == "checkbox") {
                    	if (_ipt.prop('checked')) {
                    		_ipt.prop('checked', true);
                    		target.find("tbody").find("tr").removeClass("treetable-selected");
                    		$(this).addClass("treetable-selected");
                    	} else {
                    		_ipt.prop('checked', false);
                    		target.find("tbody").find("tr").removeClass("treetable-selected");
                    	}
                    } else {
                        if (_ipt.prop('checked')) {
                            _ipt.prop('checked', false);
                            $(this).removeClass("treetable-selected");
                        } else {
                            _ipt.prop('checked', true);
                            $(this).addClass("treetable-selected");
                        }
                    }
                }
            });
        }
        // 注册小图标点击事件--展开缩起
        var registerExpanderEvent = function() {
            target.find("tbody").find("tr").find(".treetable-expander").unbind();
            target.find("tbody").find("tr").find(".treetable-expander").click(function() {
                var _isExpanded = $(this).hasClass(options.expanderExpandedClass);
                var _isCollapsed = $(this).hasClass(options.expanderCollapsedClass);
                if (_isExpanded || _isCollapsed) {
                    var tr = $(this).parent().parent();
                    var row_id = tr.attr("id");
                    var _ls = target.find("tbody").find("tr[id^='" + row_id + "_']"); //下所有
                    if (_isExpanded) {
                        $(this).removeClass(options.expanderExpandedClass);
                        $(this).addClass(options.expanderCollapsedClass);
                        if (_ls && _ls.length > 0) {
                            $.each(_ls, function(index, item) {
                                $(item).css("display", "none");
                            });
                        }
                    } else {
                        $(this).removeClass(options.expanderCollapsedClass);
                        $(this).addClass(options.expanderExpandedClass);
                        if (_ls && _ls.length > 0) {
                            $.each(_ls, function(index, item) {
                                // 父icon
                                var _p_icon = $("#" + $(item).attr("pid")).children().eq(options.expandColumn).find(".treetable-expander");
                                if (_p_icon.hasClass(options.expanderExpandedClass)) {
                                    $(item).css("display", "table");
                                }
                            });
                        }
                    }
                }
            });
        }
        // 刷新数据
        target.refresh = function(parms) {
            if(parms){
                target.lastAjaxParams=parms;
            }
            initServer(target.lastAjaxParams);
        }
        // 添加数据刷新表格
        target.appendData = function(data) {
            // 下边的操作主要是为了查询时让一些没有根节点的节点显示
            $.each(data, function(i, item) {
                var _data = target.data_obj["id_" + item[options.code]];
                var _p_data = target.data_obj["id_" + item[options.parentCode]];
                var _c_list = target.data_list["_n_" + item[options.parentCode]];
                var row_id = ""; //行id
                var p_id = ""; //父行id
                var _lv = 1; //如果没有父就是1默认显示
                var tr; //要添加行的对象
                if (_data && _data.row_id && _data.row_id != "") {
                    row_id = _data.row_id; // 如果已经存在了，就直接引用原来的
                }
                if (_p_data) {
                    p_id = _p_data.row_id;
                    if (row_id == "") {
                        var _tmp = 0
                        if (_c_list && _c_list.length > 0) {
                            _tmp = _c_list.length;
                        }
                        row_id = _p_data.row_id + "_" + _tmp;
                    }
                    _lv = _p_data.lv + 1; //如果有父
                    // 绘制行
                    tr = renderRow(item, false, _lv, row_id, p_id);

                    var _p_icon = $("#" + _p_data.row_id).children().eq(options.expandColumn).find(".treetable-expander");
                    var _isExpanded = _p_icon.hasClass(options.expanderExpandedClass);
                    var _isCollapsed = _p_icon.hasClass(options.expanderCollapsedClass);
                    // 父节点有没有展开收缩按钮
                    if (_isExpanded || _isCollapsed) {
                        // 父节点展开状态显示新加行
                        if (_isExpanded) {
                            tr.css("display", "table");
                        }
                    } else {
                        // 父节点没有展开收缩按钮则添加
                        _p_icon.addClass(options.expanderCollapsedClass);
                    }

                    if (_data) {
                        $("#" + _data.row_id).before(tr);
                        $("#" + _data.row_id).remove();
                    } else {
                        // 计算父的同级下一行
                        var _tmp_ls = _p_data.row_id.split("_");
                        var _p_next = _p_data.row_id.substring(0, _p_data.row_id.length - 1) + (parseInt(_tmp_ls[_tmp_ls.length - 1]) + 1);
                        // 画上
                        $("#" + _p_next).before(tr);
                    }
                } else {
                    tr = renderRow(item, false, _lv, row_id, p_id);
                    if (_data) {
                        $("#" + _data.row_id).before(tr);
                        $("#" + _data.row_id).remove();
                    } else {
                        // 画上
                        var tbody = target.find("tbody");
                        tbody.append(tr);
                    }
                }
                item.isShow = true;
                // 缓存并格式化数据
                formatData([item]);
            });
            registerExpanderEvent();
            registerRowClickEvent();
            initHiddenColumns();
        }

        // 展开/折叠指定的行
        target.toggleRow=function(id) {
            var _rowData = target.data_obj["id_" + id];
            var $row_expander = $("#"+_rowData.row_id).find(".treetable-expander");
            $row_expander.trigger("click");
        }
        // 展开指定的行
        target.expandRow=function(id) {
            var _rowData = target.data_obj["id_" + id];
            var $row_expander = $("#"+_rowData.row_id).find(".treetable-expander");
            var _isCollapsed = $row_expander.hasClass(target.options.expanderCollapsedClass);
            if (_isCollapsed) {
                $row_expander.trigger("click");
            }
        }
        // 折叠 指定的行
        target.collapseRow=function(id) {
            var _rowData = target.data_obj["id_" + id];
            var $row_expander = $("#"+_rowData.row_id).find(".treetable-expander");
            var _isExpanded = $row_expander.hasClass(target.options.expanderExpandedClass);
            if (_isExpanded) {
                $row_expander.trigger("click");
            }
        }
        // 展开所有的行
        target.expandAll=function() {
            target.find("tbody").find("tr").find(".treetable-expander").each(function(i,n){
                var _isCollapsed = $(n).hasClass(options.expanderCollapsedClass);
                if (_isCollapsed) {
                    $(n).trigger("click");
                }
            })
        }
        // 折叠所有的行
        target.collapseAll=function() {
            target.find("tbody").find("tr").find(".treetable-expander").each(function(i,n){
                var _isExpanded = $(n).hasClass(options.expanderExpandedClass);
                if (_isExpanded) {
                    $(n).trigger("click");
                }
            })
        }
        // 显示指定列
        target.showColumn=function(field,flag) {
            var _index = $.inArray(field, target.hiddenColumns);
            if (_index > -1) {
                target.hiddenColumns.splice(_index, 1);
            }
            target.find("."+field+"_cls").show();
            //是否更新列选项状态
            if(flag&&options.showColumns){
                var $input = $(".bootstrap-tree-table .treetable-bars .columns label").find("input[value='"+field+"']")
                $input.prop("checked", 'checked');
            }
        }
        // 隐藏指定列
        target.hideColumn=function(field,flag) {
            target.hiddenColumns.push(field);
            target.find("."+field+"_cls").hide();
            //是否更新列选项状态
            if(flag&&options.showColumns){
                var $input = $(".bootstrap-tree-table .treetable-bars .columns label").find("input[value='"+field+"']")
                $input.prop("checked", '');
            }
        }
        // 解析数据，支持多层级访问
        var getItemField = function (item, field) {
            var value = item;

            if (typeof field !== 'string' || item.hasOwnProperty(field)) {
                return item[field];
            }
            var props = field.split('.');
            for (var p in props) {
                value = value && value[props[p]];
            }
            return value;
        };
        // 发起对目标(target)函数的调用
        var calculateObjectValue = function (self, name, args, defaultValue) {
            var func = name;

            if (typeof name === 'string') {
                var names = name.split('.');

                if (names.length > 1) {
                    func = window;
                    $.each(names, function (i, f) {
                        func = func[f];
                    });
                } else {
                    func = window[name];
                }
            }
            if (typeof func === 'object') {
                return func;
            }
            if (typeof func === 'function') {
                return func.apply(self, args);
            }
            if (!func && typeof name === 'string' && sprintf.apply(this, [name].concat(args))) {
                return sprintf.apply(this, [name].concat(args));
            }
            return defaultValue;
        };
        // 初始化
        init();
        return target;
    };

    // 组件方法封装........
    $.fn.bootstrapTreeTable.methods = {
        // 为了兼容bootstrap-table的写法，统一返回数组，这里返回了表格显示列的数据
        getSelections: function(target, data) {
            // 所有被选中的记录input
            var _ipt = target.find("tbody").find("tr").find("input[name='select_item']:checked");
            var chk_value = [];
            // 如果是radio
            if (_ipt.attr("type") == "radio") {
                var _data = target.data_obj["id_" + _ipt.val()];
                chk_value.push(_data);
            } else {
                _ipt.each(function(_i, _item) {
                    var _data = target.data_obj["id_" + $(_item).val()];
                    chk_value.push(_data);
                });
            }
            return chk_value;
        },
        // 刷新记录
        refresh: function(target, parms) {
            if (parms) {
                target.refresh(parms);
            } else {
                target.refresh();
            }
        },
        // 添加数据到表格
        appendData: function(target, data) {
            if (data) {
                target.appendData(data);
            }
        },
        // 展开/折叠指定的行
        toggleRow: function(target, id) {
            target.toggleRow(id);
        },
        // 展开指定的行
        expandRow: function(target, id) {
            target.expandRow(id);
        },
        // 折叠 指定的行
        collapseRow: function(target, id) {
            target.collapseRow(id);
        },
        // 展开所有的行
        expandAll: function(target) {
            target.expandAll();
        },
        // 折叠所有的行
        collapseAll: function(target) {
            target.collapseAll();
        },
        // 显示指定列
        showColumn: function(target,field) {
            target.showColumn(field,true);
        },
        // 隐藏指定列
        hideColumn: function(target,field) {
            target.hideColumn(field,true);
        }
        // 组件的其他方法也可以进行类似封装........
    };

    $.fn.bootstrapTreeTable.defaults = {
        code: 'code',              // 选取记录返回的值,用于设置父子关系
        parentCode: 'parentCode',  // 用于设置父子关系
        rootIdValue: null,         // 设置根节点id值----可指定根节点，默认为null,"",0,"0"
        data: null,                // 构造table的数据集合
        type: "GET",               // 请求数据的ajax类型
        url: null,                 // 请求数据的ajax的url
        ajaxParams: {},            // 请求数据的ajax的data属性
        expandColumn: 0,           // 在哪一列上面显示展开按钮
        expandAll: false,          // 是否全部展开
        expandFirst: true,         // 是否默认第一级展开--expandAll为false时生效
        striped: false,            // 是否各行渐变色
        bordered: true,            // 是否显示边框
        hover: true,               // 是否鼠标悬停
        condensed: false,          // 是否紧缩表格
        columns: [],               // 列
        toolbar: null,             // 顶部工具条
        height: 0,                 // 表格高度
        showTitle: true,           // 是否采用title属性显示字段内容（被formatter格式化的字段不会显示）
        showSearch: true,          // 是否显示检索信息
        showColumns: true,         // 是否显示内容列下拉框
        showRefresh: true,         // 是否显示刷新按钮
        expanderExpandedClass: 'glyphicon glyphicon-chevron-down', // 展开的按钮的图标
        expanderCollapsedClass: 'glyphicon glyphicon-chevron-right', // 缩起的按钮的图标
        responseHandler: function(res) {
            return false;
        }
    };
})(jQuery);