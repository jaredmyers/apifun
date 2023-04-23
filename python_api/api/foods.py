from fastapi import APIRouter, Depends
from storage.initdb import get_session
from storage.models import User, Food
from sqlmodel import select
from fastapi.responses import JSONResponse


router = APIRouter()


# == Food Routes ==

# gets all foods
@router.get("/foods")
async def get_foods(session=Depends(get_session)):
    statement = select(Food)
    result = session.execute(statement).all()
    return result


# create new food
@router.post("/foods")
async def create_food(new_food: Food, session=Depends(get_session)):
    session.add(new_food)
    session.commit()
    return JSONResponse(status_code=200, content={"foods": "created"})


# get food by id
@router.get("/food/{id}")
def get_user(food: int, session=Depends(get_session)):
    statement = select(Food).filter(Food.id == id)
    result = session.execute(statement).all()
    return result
