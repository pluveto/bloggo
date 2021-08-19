from enum import Enum

class Action(Enum):
    create = 'create'
    get = 'get'
    list = 'list'
    update = 'update'
    search = 'search'
    delete = 'delete'
    enable = 'enable'
    disable = 'disable'