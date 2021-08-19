from .model import Field, Mode, Model

import json
class ModelProvider:

    def __init__(self, models) -> None:
        self.models = models

    @staticmethod
    def from_json( models_json):
        models_dict = json.loads(models_json)['models']
        models = []
        for m in models_dict:
            models.append(Model.from_dict(m))
        print("%d model loaded" % len(models))
        return ModelProvider(models)

    @staticmethod
    def from_file( models_path):        
        models_json = open(models_path, encoding="utf-8", mode="r").read()
        return ModelProvider.from_json(models_json)
        
    
    def get_model(self, name: str) -> Model:
        for model in self.models:
            if model.name == name:
                return model
        raise RuntimeError("Unknown model: "+str(name))