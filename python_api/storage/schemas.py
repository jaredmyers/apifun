from sqlmodel import Relationship, SQLModel, Field
from typing import Optional, List


class Team(SQLModel, table=True):
    id: Optional[int] = Field(default=None, primary_key=True)
    name: str
    headquarters: str

    users: List["User"] = Relationship(back_populates="team")


class User(SQLModel, table=True):
    __tablename__ = "users"

    id: Optional[int] = Field(None, primary_key=True)
    fname: str
    lname: str
    email: str
    team_id: Optional[int] = Field(default=None, foreign_key="team.id")

    team: Optional[Team] = Relationship(back_populates="users")

    class Config:
        arbitrary_types_allowed = True


"""
class Food(SQLModel, table=True):
    __tablename__ = "foods"

    id: Optional[int] = Field(None, primary_key=True)
    name: str

    class Config:
        arbitrary_types_allowed = True
"""
