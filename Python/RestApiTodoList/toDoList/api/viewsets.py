from re import S
from rest_framework import viewsets
from toDoList.api import serializers
from toDoList.models import TodoItem


class TodoViewSet(viewsets.ModelViewSet):
    serializer_class = serializers.TodoSerializer
    queryset = TodoItem.objects.all()
