from fastapi import APIRouter, Depends
from storage.initdb import get_async_session
from storage.models import User, Food
from sqlmodel import select
from fastapi.responses import JSONResponse


router = APIRouter()


# == Food Routes ==

# gets all foods
@router.get("/foods")
async def get_foods(session=Depends(get_async_session)):
    statement = select(Food)
    result = session.execute(statement).all()
    return result


# get food by id
@router.get("/food/{id}")
def get_user(food: int, session=Depends(get_async_session)):
    statement = select(Food).filter(Food.id == id)
    result = session.execute(statement).first()
    return result


# create, update, delete
# create
@router.post("/foods")
async def create_food(new_food: Food, session=Depends(get_async_session)):
    session.add(new_food)
    session.commit()
    return JSONResponse(status_code=200, content={"foods": "created"})


# delete
@router.delete("/food/{id}")
def delete_food(food: int, session=Depends(get_async_session)):
    pass


# update
@router.put("/food/{id}")
def update_food(food: int, session=Depends(get_async_session)):
    pass


# Nutrition for food
# get, create, delete, update
@router.get("/foods/{id}/nutrition}")
async def get_nutrition(session=Depends(get_async_session)):
    pass


@router.post("/foods/{id}/nutrition}")
async def create_nutrition(session=Depends(get_async_session)):
    pass


@router.delete("/foods/{id}/nutrition}")
async def delete_nutrition(session=Depends(get_async_session)):
    pass


@router.put("/foods/{id}/nutrition}")
async def update_nutrition(session=Depends(get_async_session)):
    pass
