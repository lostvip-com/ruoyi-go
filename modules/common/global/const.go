package global

import "os"

const Layout = "2006-01-02 15:04:05" //时间常量
const ErrTimes2Lock = 20

// KeyOssUrl 获取oss地址Key
const KeyOssUrl = "sys.resource.url"

const JQ_BE_NO = "false"
const JQ_BE_OK = "true"

const DIR_DIST_CODE = ""
const DIR_TPL_CODE_GEN = "resources" + string(os.PathSeparator) + "tpl_gen"
