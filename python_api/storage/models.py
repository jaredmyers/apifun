from sqlmodel import Relationship, SQLModel, Field
from typing import Optional, List


class Store(SQLModel, table=True):
    __tablename__ = "stores"

    id: Optional[int] = Field(None, primary_key=True)
    name: str

    foods: List["Food"] = Relationship(back_populates="stores")

    class Config:
        arbitrary_types_allowed = True


class Nutrition(SQLModel, table=True):
    __tablename__ = "nutrition"

    id: Optional[int] = Field(None, primary_key=True)
    food_id: int = Field(foreign_key="foods.id")
    price: int
    server_per_container: int
    serving_size: int
    serving_unit: str
    cholesterol: int
    sodium: int
    protein: int
    carb: int
    fat: int
    fiber: int
    sugar: int

    class Config:
        arbitrary_types_allowed = True


class GroceryList(SQLModel, table=True):
    __tablename_ = "grocerylist"

    user_id: Optional[int] = Field(default=None, foreign_key="users.id", primary_key=True)
    food_id: Optional[int] = Field(default=None, foreign_key="foods.id", primary_key=True)

    class Config:
        arbitrary_types_allowed = True


class Food(SQLModel, table=True):
    __tablename__ = "foods"

    id: Optional[int] = Field(None, primary_key=True)
    name: str
    store_id: int = Field(foreign_key="stores.id")
    isle_id: str

    stores: Optional["Store"] = Relationship(back_populates="foods")

    class Config:
        arbitrary_types_allowed = True


class User(SQLModel, table=True):
    __tablename__ = "users"

    id: Optional[int] = Field(default=None, primary_key=True)
    username: str
    pw_hash: str

    class Config:
        arbitrary_types_allowed = True
