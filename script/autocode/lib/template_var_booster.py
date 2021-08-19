from lib.model_provider import ModelProvider
from .model import Mode, Model
from typing import List
from .api_group import API, APIGroup, Field


class TemplateVarBooster:
    def __init__(self, api_group: APIGroup, model_provider: ModelProvider) -> None:
        self.group = api_group
        self.model_provider = model_provider
        pass

    def boost(self):
        model:Model = self.model_provider.get_model(name=self.group.name)
        self.group.model = model
        self.group.unique_fields = self.get_fields_unique(model.fields)
        for api in self.group.apis:
            api.full_action = api.action + api.action_on
            api.unique_field_names = self.get_field_names_unique(api.req_fields)
        pass

    def get_field_names_unique(self, w: List[Field]) -> List[str]:
        """获取所有 unique 修饰的字段"""
        def f(field: Field) -> bool:
            if(not(field.mode)):
                return False
            if field.column_meta.unique:
                return True
            if field.column_meta.primary_key:
                return True
        fs = list(filter(f, w))
        ret = []
        for item in fs:
            ret.append(item.name)
        return ret

    def get_fields_unique(self, w: List[Field]) -> List[Field]:
        """获取所有 unique 修饰的字段"""
        def f(field: Field) -> bool:
            if(not(field.mode)):
                return False
            if field.column_meta.unique:
                return True
            if field.column_meta.primary_key:
                return True
        fs = list(filter(f, w))
        return fs
        
    def in_modes(self, mode: Mode, fi: List[Mode]):
        for m in fi:
            if(str(m) == str(mode)):
                return True
        return False

    def get_vars(self):
        return self.group
        pass
