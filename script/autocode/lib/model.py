# To use this code, make sure you
#
#     import json
#
# and then, to convert JSON from a string, do
#
#     result = empty_from_dict(json.loads(json_string))

from dataclasses import dataclass
from typing import Optional, Any, List, TypeVar, Callable, Type, cast
from enum import Enum


T = TypeVar("T")
EnumT = TypeVar("EnumT", bound=Enum)


def from_str(x: Any) -> str:
    assert isinstance(x, str)
    return x


def from_none(x: Any) -> Any:
    assert x is None
    return x


def from_union(fs, x):
    for f in fs:
        try:
            return f(x)
        except:
            pass
    assert False


def from_bool(x: Any) -> bool:
    assert isinstance(x, bool)
    return x


def from_int(x: Any) -> int:
    assert isinstance(x, int) and not isinstance(x, bool)
    return x


def from_list(f: Callable[[Any], T], x: Any) -> List[T]:
    assert isinstance(x, list)
    return [f(y) for y in x]


def to_class(c: Type[T], x: Any) -> dict:
    assert isinstance(x, c)
    return cast(Any, x).to_dict()


def to_enum(c: Type[EnumT], x: Any) -> EnumT:
    assert isinstance(x, c)
    return x.value


@dataclass
class ColumnMeta:
    column: Optional[str] = None
    primary_key: Optional[bool] = None
    unique: Optional[bool] = None

    @staticmethod
    def from_dict(obj: Any) -> 'ColumnMeta':
        assert isinstance(obj, dict)
        column = from_union([from_str, from_none], obj.get("column"))
        primary_key = from_union([from_bool, from_none], obj.get("primaryKey"))
        unique = from_union([from_bool, from_none], obj.get("unique"))
        return ColumnMeta(column, primary_key, unique)

    def to_dict(self) -> dict:
        result: dict = {}
        result["column"] = from_union([from_str, from_none], self.column)
        result["primaryKey"] = from_union(
            [from_bool, from_none], self.primary_key)
        result["unique"] = from_union([from_bool, from_none], self.unique)
        return result

@dataclass
class Mode(Enum):
    I = "i"
    R = "r"
    W = "w"


@dataclass
class Field:
    name: Optional[str] = None
    json: Optional[str] = None
    column: Optional[str] = None
    column_meta: Optional[ColumnMeta] = None
    kind: Optional[str] = None
    weight: Optional[int] = None
    mode: Optional[List[Mode]] = None
    label: Optional[str] = None
    type: Optional[str] = None
    searchable:Optional[str] = False
    sortable:Optional[str] = False
    required:Optional[str] = True

    @staticmethod
    def from_dict(obj: Any) -> 'Field':
        assert isinstance(obj, dict)
        name = from_union([from_str, from_none], obj.get("name"))
        json = from_union([from_str, from_none], obj.get("json"))
        column = from_union([from_str, from_none], obj.get("column"))
        column_meta = from_union(
            [ColumnMeta.from_dict, from_none], obj.get("column_meta"))
        kind = from_union([from_str, from_none], obj.get("kind"))
        weight = from_union([from_int, from_none], obj.get("weight"))
        mode = from_union([lambda x: from_list(Mode, x),
                          from_none], obj.get("mode"))
        label = from_union([from_none, from_str], obj.get("label"))
        type = from_union([from_none, from_str], obj.get("type"))
        return Field(name, json, column, column_meta, kind, weight, mode, label, type)

    def to_dict(self) -> dict:
        result: dict = {}
        result["name"] = from_union([from_str, from_none], self.name)
        result["json"] = from_union([from_str, from_none], self.json)
        result["column"] = from_union([from_str, from_none], self.column)
        result["column_meta"] = from_union(
            [lambda x: to_class(ColumnMeta, x), from_none], self.column_meta)
        result["kind"] = from_union([from_str, from_none], self.kind)
        result["weight"] = from_union([from_int, from_none], self.weight)
        result["mode"] = from_union([lambda x: from_list(
            lambda x: to_enum(Mode, x), x), from_none], self.mode)
        result["label"] = from_union([from_none, from_str], self.label)
        result["type"] = from_union([from_none, from_str], self.type)
        return result


@dataclass
class Meta:
    table: Optional[str] = None
    json: Optional[str] = None

    @staticmethod
    def from_dict(obj: Any) -> 'Meta':
        assert isinstance(obj, dict)
        table = from_union([from_str, from_none], obj.get("table"))
        json = from_union([from_str, from_none], obj.get("json"))
        return Meta(table, json)

    def to_dict(self) -> dict:
        result: dict = {}
        result["table"] = from_union([from_str, from_none], self.table)
        result["json"] = from_union([from_str, from_none], self.json)
        return result


@dataclass
class Model:
    name: Optional[str] = None
    meta: Optional[Meta] = None
    fields: Optional[List[Field]] = None

    @staticmethod
    def from_dict(obj: Any) -> 'Model':
        assert isinstance(obj, dict)
        name = from_union([from_str, from_none], obj.get("name"))
        meta = from_union([Meta.from_dict, from_none], obj.get("meta"))
        fields = from_union([lambda x: from_list(
            Field.from_dict, x), from_none], obj.get("fields"))
        return Model(name, meta, fields)

    def to_dict(self) -> dict:
        result: dict = {}
        result["name"] = from_union([from_str, from_none], self.name)
        result["meta"] = from_union(
            [lambda x: to_class(Meta, x), from_none], self.meta)
        result["fields"] = from_union([lambda x: from_list(
            lambda x: to_class(Field, x), x), from_none], self.fields)
        return result

    def get_fields_by_mode(self, mode: Mode) -> List[Field]:
        """
        获取指定模式（可读、可写）的所有字段
        """
        def f(field: Field) -> bool:
            if(not(field.mode)):
                return False
            # 隐藏内部字段
            if self.in_modes(Mode.I, field.mode):
                return False
            if self.in_modes(mode, field.mode):
                return True
        ret = list(filter(f, self.fields))
        
        
        return ret
    def in_modes(self, mode:Mode, fi:List[Mode]):
        for m in fi:
            if(str(m) == str(mode)):
                return True
        return False

    def get_fields_unique(self) -> List[Field]:
        """获取所有 unique 修饰的字段"""
        def f(field: Field) -> bool:
            if(not(field.mode)):
                return False
            # 隐藏内部字段
            if self.in_modes(Mode.I, field.mode):
                return False
            if field.column_meta.unique:
                return True
            if field.column_meta.primary_key:
                return True
        return list(filter(f, self.fields))

    def get_pkey_field(self) -> Field:
        """获取所有 unique 修饰的字段"""
        for field in self.fields:
            if(not(field.mode)):
                continue
            # 隐藏内部字段
            if self.in_modes(Mode.I, field.mode):
                continue
            if field.column_meta.primary_key:
                return field
        return None

    def get_fields_searchable(self) -> List[Field]:
        def f(field: Field) -> bool:
            if(not(field.mode)):
                return False
            # 隐藏内部字段
            if self.in_modes(Mode.I, field.mode):
                return False
            if field.searchable:
                return True
        return list(filter(f, self.fields))

    def get_field(self, fieldName) -> Field :
        try:
            return next(filter(lambda fi: fi.name == fieldName, self.fields))
        except StopIteration:
            return None
