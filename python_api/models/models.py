from pydantic import BaseModel


class CreateUserRequest(BaseModel):
    username: str
    pw_hash: str


class DeleteUserRequest(BaseModel):
    username: str
