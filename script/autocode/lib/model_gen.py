"""

model_gen.py

本文件负责读取指定目录下的 *.go 文件，然后从中提取形如：

    /*
     * @autocode({"table": "admins", "json": "admin"})
    */
    type TypeName struct {
       ID int64  `json:"id" gorm:   "column=id;primaryKey" auto:"o=0;m=r;"` <--- body
            ^               ^tag_key^-----------------tag_value-----------^
            |     ^-----------------------tags----------------------------^
            kind
    }

的结构体，将其分析后输出如下结构：

{
  "models": [
    {
      "name": "Admin", // 模型名
      "meta": {
        "table": "admins", // 表名
        "json": "admin" // json 名
      },
      "fields": [
        {
          "name": "ID",
          "json": "id",
          "column": "id",
          "kind": "int64",
          "weight": 0, // 权重，决定表单的生成顺序，越小越靠前
          "mode": [
            "r" // r: 可读，w： 可写, i: 对外隐藏
          ],
          "label": null, // 在用户表单显式的名称
          "type": null // HTML 表单类型
        },
        ...
       ]
    }
  ]       
}

"""
from typing import List
from .ac_util import ac_util
from string import Template
import os
import re
import json


def get_files(dir_name, ext):
    _, _, filenames = next(os.walk(dir_name))
    ret = []
    for name in filenames:
        if not name.endswith(ext):
            continue
        ret.append(os.path.join(dir_name, name))
    return ret


def parse_tag_value(key: str, value: str) -> dict:
    """
    input: 
        gorm, "column=email; unique"
    output:
        {
            column: "email",
            unique: true,
        }
    """
    # schemes
    if not key in ["gorm", "auto", "binding"]:
        # for json, value is what is need
        return value

    def to_kv(x: str):
        kv = x.split("=")
        if(len(kv) == 1):
            # 对于 primaryKey; unique 这种
            # 处理成 primaryKey: true
            #      unique: true
            kv.append(True)
        else:
            kv[1] = kv[1].strip()
        kv[0] = kv[0].strip()
        return kv
    if(key in ["gorm", "auto"]):
        items = value.split(";")
    elif(key == "binding"):
        items = value.split(",")
    ret = {}
    for item in items:
        kv = to_kv(item)
        ret[kv[0]] = kv[1]

    return ret


def parse_tags(tagline: str) -> dict:
    """
    input:
        `json:"id"          gorm:"column=id;primaryKey" auto:"m=r;  "`
    output:
        json:"id"
        gorm:"column=id;primaryKey"
        auto:"m=r;  "
    """
    tagline = tagline.strip("`")
    regex = r"(\w+):\"(.*?)\""
    matches = re.finditer(regex, tagline)
    tags = {}
    for match in matches:
        key = match.group(1)
        value = match.group(2).strip()
        if(len(value) == 0):
            continue
        tags[key] = parse_tag_value(key, value)
    return tags


def parse_body(body_str: str) -> list:
    """
    Parse body of go struct
    """
    fields: list = []
    for line in body_str.splitlines():
        line = line.strip()
        if(len(line) == 0):
            continue
        tokens = line.split(maxsplit=2)
        tags = parse_tags(tokens[2])

        def getOrNone(d, k):
            label = None
            if(k in d):
                label = d[k]
            return label
        field = {
            "name": tokens[0],
            "json": tags["json"],
            "column": tags["gorm"]["column"],
            "column_meta": tags["gorm"],
            "kind": tokens[1],
            "weight": int(getOrNone(tags["auto"], "weight")),
            "searchable": bool(getOrNone(tags["auto"], "searchable")),
            "sortable": bool(getOrNone(tags["auto"], "sortable")),
            "mode": list(getOrNone(tags["auto"], "mode")),
            "label": getOrNone(tags["auto"], "label"),
            "type": getOrNone(tags["auto"], "type"),
        }
        fields.append(field)
    return fields


def gen_models(dir_name) -> list:
    gofiles = get_files(dir_name, ".go")
    """
    对于下列例子，将匹配出
     [match1, ...]
     match1 [
         group1 = {"table": "table_name"} -> meta
         group2 = TypeName -> type_name
         group3 = Body -> type_body
     ]
    /*
    * @autocode({"table": "table_name"})
    */
    type TypeName struct {
        Body
    }
    """
    models = []
    regex = r"@autocode\((.*)\)[ \t]*\n[\w\W]*?type (.*?) struct {([\s\S]*?)}"
    for filename in gofiles:
        content = open(filename, encoding="utf-8").read()
        matches = re.finditer(regex, content, re.MULTILINE)
        for match in matches:
            model = {}
            model['name'] = match.group(2)
            model['meta'] = json.loads(match.group(1))
            lazyCheck(model, match.group(3))
            model['fields'] = parse_body(match.group(3))
            models.append(model)
    return models


def lazyCheck(model, body):
    """
    在注解中加上     ↓
    @autocode({..."lazy": true})
    即可享受懒人功能~
    """
    if('lazy' in model['meta'] and model['meta']['lazy']):
        print("~~~~~ copy follow ~~~~~~")
        print("\n".join(complete_body(body)))
        exit(0)


def complete_body(body_str: str) -> List[str]:
    """
    Parse body of go struct
    """
    ret = []
    incr = 0
    for line in body_str.splitlines():
        line = line.strip()
        if(len(line) == 0):
            continue
        tokens = line.split(maxsplit=2)
        fieldName = tokens[0]
        kind = tokens[1]
        tagTpl = Template(
            '`json:"$json" '
            + 'gorm:"column=$column" '
            + 'auto:"weight=$weight;searchable=$searchable;'
            + 'sortable=$sortable;mode=$mode;required=$required;'
            + 'label=$label;type=$type`"')
        mode = "rw"
        underscore = ac_util.pascal_to_underscore(fieldName)
        if ac_util.any_in_str(["id", "password", "secret"],underscore):
            mode = "i"
        tag = tagTpl.substitute(
            json=ac_util.pascal_to_camel(fieldName),
            column=underscore,
            weight=incr,
            searchable="false",
            sortable="false",
            mode=mode,
            required="true",
            label=ac_util.en_to_ch(fieldName),
            type="text", # todo: 智能判断
        )
        incr += 1
        ret.append(fieldName + " " + kind + " " + tag)
    return ret
