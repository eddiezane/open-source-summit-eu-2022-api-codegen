from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class DeleteImageRequest(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class DeleteImageResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class GenerateImageRequest(_message.Message):
    __slots__ = ["prompt"]
    PROMPT_FIELD_NUMBER: _ClassVar[int]
    prompt: str
    def __init__(self, prompt: _Optional[str] = ...) -> None: ...

class GenerateImageResponse(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class GetImageRequest(_message.Message):
    __slots__ = ["id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class GetImageResponse(_message.Message):
    __slots__ = ["url"]
    URL_FIELD_NUMBER: _ClassVar[int]
    url: str
    def __init__(self, url: _Optional[str] = ...) -> None: ...

class ListImagesRequest(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class ListImagesResponse(_message.Message):
    __slots__ = ["images"]
    IMAGES_FIELD_NUMBER: _ClassVar[int]
    images: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, images: _Optional[_Iterable[str]] = ...) -> None: ...
