""" 
通过 https://app.quicktype.io/ 生成
表单规范：
https://jsonforms.io/examples/categorization
"""

from .model import Field
from dataclasses_json import dataclass_json, LetterCase
from dataclasses import dataclass, field
from typing import List, Any, Union, Optional


@dataclass_json(letter_case=LetterCase.CAMEL)
@dataclass
class API:
    form: Union[Any, None] = field(default_factory=dict)
    req_fields: Optional[List[Field]] = None
    ret_fields: Optional[List[Field]] = None
    paging: Optional[bool] = False
    method: Optional[str] = "POST"
    route: Optional[str] = "."
    action: Optional[str] = "."
    action_on: Optional[str] = "."
    base_model: Optional[str] = "."
    permissions: Optional[List[Any]] = field(default_factory=list)
    auto_assign: Optional[Any] = None
    order: Optional[int] = 0


@dataclass_json(letter_case=LetterCase.CAMEL)
@dataclass
class Meta:
    json: Optional[str] = None
    label: Optional[str] = None


@dataclass_json(letter_case=LetterCase.CAMEL)
@dataclass
class APIGroup:
    name: Optional[str] = None
    route: Optional[str] = None
    meta: Optional[Meta] = None
    apis: Optional[List[API]] = None

    